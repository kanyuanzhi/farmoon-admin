package rpcServer

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
	pb "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/rpc/rpcv1"
	"gorm.io/gorm/clause"
)

type DataUpdate struct {
	pb.DataUpdaterServer
}

func (server *DataUpdate) FetchOfficialDishes(c context.Context, req *pb.FetchOfficialDishesRequest) (*pb.FetchOfficialDishesResponse, error) {
	var localDishesInfo map[uuid.UUID]int64
	if err := json.Unmarshal(req.GetLocalDishesInfoJson(), &localDishesInfo); err != nil {
		return nil, err
	}

	var localDishesStatus = make(map[uuid.UUID]bool)
	for key, _ := range localDishesInfo {
		localDishesStatus[key] = false
	}

	var localNeedAddDishUUIDs []uuid.UUID
	var localNeedUpdateDishUUIDs []uuid.UUID
	var localNeedDeleteDishUUIDs []uuid.UUID

	var remoteDishes []model.SysDish
	if err := global.FXDb.Where("is_official = ?", true).Select("uuid", "updated_at").Find(&remoteDishes).Error; err != nil {
		return nil, err
	}

	for _, remoteDish := range remoteDishes {
		if _, has := localDishesInfo[remoteDish.UUID]; has {
			localDishesStatus[remoteDish.UUID] = true
			if localDishesInfo[remoteDish.UUID] < remoteDish.UpdatedAt {
				localNeedUpdateDishUUIDs = append(localNeedUpdateDishUUIDs, remoteDish.UUID)
			}
		} else {
			localNeedAddDishUUIDs = append(localNeedAddDishUUIDs, remoteDish.UUID)
		}
	}

	for key, val := range localDishesStatus {
		if !val {
			localNeedDeleteDishUUIDs = append(localNeedDeleteDishUUIDs, key)
		}
	}

	var localNeedAddDishes []model.SysDish
	var localNeedUpdateDishes []model.SysDish

	if err := global.FXDb.Where("uuid in ?", localNeedAddDishUUIDs).Find(&localNeedAddDishes).Error; err != nil {
		return nil, err
	}
	if err := global.FXDb.Where("uuid in ?", localNeedUpdateDishUUIDs).Find(&localNeedUpdateDishes).Error; err != nil {
		return nil, err
	}

	localNeedAddDishesBytes, _ := json.Marshal(localNeedAddDishes)
	localNeedUpdateDishesBytes, _ := json.Marshal(localNeedUpdateDishes)
	localNeedDeleteDishUUIDsBytes, _ := json.Marshal(localNeedDeleteDishUUIDs)

	var cuisines []model.SysCuisine
	if err := global.FXDb.Find(&cuisines).Error; err != nil {
		return nil, err
	}
	cuisinesBytes, _ := json.Marshal(cuisines)

	res := &pb.FetchOfficialDishesResponse{
		LocalNeedAddDishesJson:         localNeedAddDishesBytes,
		LocalNeedUpdateDishesJson:      localNeedUpdateDishesBytes,
		LocalNeedDeleteDishesUuidsJson: localNeedDeleteDishUUIDsBytes,
		CuisinesJson:                   cuisinesBytes,
	}

	return res, nil
}

func (server *DataUpdate) FetchCuisines(c context.Context, req *pb.FetchIngredientsRequest) (*pb.FetchIngredientsResponse, error) {
	var ingredients []model.SysIngredient
	if err := global.FXDb.Find(&ingredients).Error; err != nil {
		return nil, err
	}

	var ingredientTypes []model.SysIngredientType
	if err := global.FXDb.Find(&ingredientTypes).Error; err != nil {
		return nil, err
	}

	var ingredientShapes []model.SysIngredientShape
	if err := global.FXDb.Find(&ingredientShapes).Error; err != nil {
		return nil, err
	}

	ingredientsBytes, _ := json.Marshal(ingredients)
	ingredientTypesBytes, _ := json.Marshal(ingredientTypes)
	ingredientShapesBytes, _ := json.Marshal(ingredientShapes)

	res := &pb.FetchIngredientsResponse{
		IngredientsJson:      ingredientsBytes,
		IngredientTypesJson:  ingredientTypesBytes,
		IngredientShapesJson: ingredientShapesBytes,
	}

	return res, nil
}

func (server *DataUpdate) SynchronizePersonalDishes(c context.Context, req *pb.SynchronizePersonalDishesRequest) (*pb.SynchronizePersonalDishesResponse, error) {
	// 先删除已经在本地删除的菜谱，并记录
	var localDeletedDishUUIDs []uuid.UUID
	if err := json.Unmarshal(req.GetLocalDeletedDishUuidsJson(), &localDeletedDishUUIDs); err != nil {
		return nil, err
	}

	var remoteNeedDeleteDishesNumber int64
	if err := global.FXDb.Where("uuid in ?", localDeletedDishUUIDs).Delete(&model.SysDish{}).Count(&remoteNeedDeleteDishesNumber).Error; err != nil {
		return nil, err
	}

	var userDeletedDishes []model.SysUserDeletedDish
	for _, dishUUID := range localDeletedDishUUIDs {
		userDeletedDishes = append(userDeletedDishes, model.SysUserDeletedDish{
			UUID:  dishUUID,
			Owner: req.GetUserSerialNumber(),
		})
	}

	if len(userDeletedDishes) != 0 {
		if err := global.FXDb.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "uuid"}}, DoNothing: true}).
			Create(&userDeletedDishes).Error; err != nil {
			return nil, err
		}
	}
	// 同步本地菜谱：
	// 本地有远程无的菜谱，上传到远程
	// 本地有远程有的菜谱，依据时间确认谁需要更新
	// 本地无远程有的菜谱，下载到本地
	// bytes remote_need_add_dish_uuids_json = 1;  // 远程需要更新的菜谱uuid
	// bytes remote_need_update_dish_uuids_json = 2;// 远程需要新增的菜谱uuid
	// bytes local_need_add_dishes_json = 3;  // 本地需要新增的菜谱
	// bytes local_need_update_dishes_json = 4;  // 本地需要更新的菜谱
	// bytes local_need_delete_dishes_uuids_json = 5;  // 本地需要删除的菜谱uuid
	var localDishesInfo map[uuid.UUID]int64
	if err := json.Unmarshal(req.GetLocalDishesInfoJson(), &localDishesInfo); err != nil {
		return nil, err
	}

	var localDishesStatus = make(map[uuid.UUID]bool)
	for key, _ := range localDishesInfo {
		localDishesStatus[key] = false
	}

	var remoteNeedAddDishUUIDs []uuid.UUID
	var remoteNeedUpdateDishUUIDs []uuid.UUID
	var localNeedAddDishes []model.SysDish
	var localNeedAddDishUUIDs []uuid.UUID
	var localNeedUpdateDishes []model.SysDish
	var localNeedUpdateDishUUIDs []uuid.UUID
	var localNeedDeleteDishUUIDs []uuid.UUID

	var remoteDishes []model.SysDish
	if err := global.FXDb.Where("owner = ?", req.GetUserSerialNumber()).Select("uuid", "updated_at").
		Find(&remoteDishes).Error; err != nil {
		return nil, err
	}

	for _, remoteDish := range remoteDishes {
		if _, has := localDishesInfo[remoteDish.UUID]; has { // 本地有远程有的菜谱
			localDishesStatus[remoteDish.UUID] = true
			if localDishesInfo[remoteDish.UUID] < remoteDish.UpdatedAt { // 本地需要更新
				localNeedUpdateDishUUIDs = append(localNeedUpdateDishUUIDs, remoteDish.UUID)
			} else if localDishesInfo[remoteDish.UUID] > remoteDish.UpdatedAt { // 远程需要更新
				remoteNeedUpdateDishUUIDs = append(remoteNeedUpdateDishUUIDs, remoteDish.UUID)
			} //等于的情况不需要更新
		} else { // 本地无远程有的菜谱，本地需要新增
			localNeedAddDishUUIDs = append(localNeedAddDishUUIDs, remoteDish.UUID)
		}
	}

	var userAllDeletedDishes []model.SysUserDeletedDish
	if err := global.FXDb.Where("owner = ?", req.GetUserSerialNumber()).Find(&userAllDeletedDishes).Error; err != nil {
		return nil, err
	}
	userAllDeletedDishUUIDs := make(map[uuid.UUID]bool)
	for _, userDeletedDish := range userAllDeletedDishes {
		userAllDeletedDishUUIDs[userDeletedDish.UUID] = true
	}

	for key, val := range localDishesStatus {
		if !val { // 本地有远程无的菜谱
			if _, has := userAllDeletedDishUUIDs[key]; has { // 在其他本地端已经删除，本地需要删除
				localNeedDeleteDishUUIDs = append(localNeedDeleteDishUUIDs, key)
			} else { // 未被本地删除过，远程需要新增
				remoteNeedAddDishUUIDs = append(remoteNeedAddDishUUIDs, key)
			}
		}
	}

	if err := global.FXDb.Where("uuid in ?", localNeedAddDishUUIDs).Find(&localNeedAddDishes).Error; err != nil {
		return nil, err
	}
	if err := global.FXDb.Where("uuid in ?", localNeedUpdateDishUUIDs).Find(&localNeedUpdateDishes).Error; err != nil {
		return nil, err
	}

	localNeedAddDishesBytes, _ := json.Marshal(localNeedAddDishes)
	localNeedUpdateDishesBytes, _ := json.Marshal(localNeedUpdateDishes)
	localNeedDeleteDishUUIDsBytes, _ := json.Marshal(localNeedDeleteDishUUIDs)
	remoteNeedAddDishUUIDsBytes, _ := json.Marshal(remoteNeedAddDishUUIDs)
	remoteNeedUpdateDishUUIDsBytes, _ := json.Marshal(remoteNeedUpdateDishUUIDs)

	res := &pb.SynchronizePersonalDishesResponse{
		RemoteNeedAddDishUuidsJson:    remoteNeedAddDishUUIDsBytes,
		RemoteNeedUpdateDishUuidsJson: remoteNeedUpdateDishUUIDsBytes,
		RemoteNeedDeleteDishesNumber:  remoteNeedDeleteDishesNumber,
		LocalNeedAddDishesJson:        localNeedAddDishesBytes,
		LocalNeedUpdateDishesJson:     localNeedUpdateDishesBytes,
		LocalNeedDeleteDishUuidsJson:  localNeedDeleteDishUUIDsBytes,
	}

	return res, nil
}

func (server *DataUpdate) UploadPersonalDishes(c context.Context, req *pb.UploadPersonalDishesRequest) (*pb.UploadPersonalDishesResponse, error) {
	var remoteNeedAddDishes []model.SysDish
	var remoteNeedUpdateDishes []model.SysDish

	if err := json.Unmarshal(req.GetRemoteNeedAddDishesJson(), &remoteNeedAddDishes); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(req.GetRemoteNeedUpdateDishesJson(), &remoteNeedUpdateDishes); err != nil {
		return nil, err
	}

	tx := global.FXDb.Begin()
	if len(remoteNeedAddDishes) != 0 {
		if err := tx.Create(remoteNeedAddDishes).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if len(remoteNeedUpdateDishes) != 0 {
		for _, dish := range remoteNeedUpdateDishes {
			if err := tx.Model(&model.SysDish{}).Where("uuid = ?", dish.UUID).Omit("id", "uuid").Updates(dish).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
			// 将自动更新的updated_at字段更新为本地上传过来的值
			if err := tx.Model(&model.SysDish{}).Where("uuid = ?", dish.UUID).Update("updated_at", dish.UpdatedAt).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	res := &pb.UploadPersonalDishesResponse{
		Empty: true,
	}

	return res, nil
}
