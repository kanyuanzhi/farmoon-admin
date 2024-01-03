package private

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/request"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/response"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/utils"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm/clause"
	"image/png"
	"strconv"
	"strings"
)

type DishApi struct{}

func (api *DishApi) Count(c *gin.Context) {
	var countDishesRequest request.CountDishes
	if err := request.ShouldBindQuery(c, &countDishesRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	filterDb, err := request.GenerateDishQueryCondition(countDishesRequest.Filter, countDishesRequest.EnableCuisineFilter,
		strings.Split(countDishesRequest.CuisineFilter, ","), countDishesRequest.EnableOwnerFilter, countDishesRequest.OwnerFilter, countDishesRequest.IsOfficial)
	if err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	var count int64
	if err := filterDb.Count(&count).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	countDishesResponse := response.CountDishes{
		Count: count,
	}

	response.SuccessData(c, countDishesResponse)
}

func (api *DishApi) List(c *gin.Context) {
	var listDishesRequest request.ListDishes
	if err := request.ShouldBindQuery(c, &listDishesRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	filterDb, err := request.GenerateDishQueryCondition(listDishesRequest.Filter, listDishesRequest.EnableCuisineFilter,
		strings.Split(listDishesRequest.CuisineFilter, ","), listDishesRequest.EnableOwnerFilter, listDishesRequest.OwnerFilter, listDishesRequest.IsOfficial)
	if err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	var dishes []model.SysDish
	if err := filterDb.Limit(listDishesRequest.PageSize).Offset((listDishesRequest.PageIndex - 1) * listDishesRequest.PageSize).
		Order("sort, updated_at asc").Find(&dishes).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	dishesInfo := []model.DishInfo{}
	for _, dish := range dishes {
		dishesInfo = append(dishesInfo, model.DishInfo{
			Id:              dish.Id,
			Image:           "data:image/png;base64," + base64.StdEncoding.EncodeToString(dish.Image),
			Name:            dish.Name,
			UUID:            dish.UUID,
			Steps:           dish.Steps,
			CustomStepsList: dish.CustomStepsList,
			Cuisine:         dish.Cuisine,
			Owner:           dish.Owner,
		})
	}

	listDishesResponse := response.ListDishes{
		Dishes: dishesInfo,
	}

	response.SuccessData(c, listDishesResponse)
}

// 仅更新名称和所属菜系
func (api *DishApi) Update(c *gin.Context) {
	var updateDishRequest request.UpdateDish
	if err := request.ShouldBindJSON(c, &updateDishRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	dish := model.SysDish{
		FXModel: global.FXModel{
			Id: updateDishRequest.Id,
		},
		Name:    updateDishRequest.Name,
		Cuisine: updateDishRequest.Cuisine,
	}

	if err := global.FXDb.Model(&dish).Select("name", "cuisine").Updates(dish).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "更新成功")
}

func (api *DishApi) UpdateWithSteps(c *gin.Context) {
	var updateDishWithStepsRequest request.UpdateDishWithSteps
	if err := request.ShouldBindJSON(c, &updateDishWithStepsRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	dish := model.SysDish{
		FXModel: global.FXModel{
			Id: updateDishWithStepsRequest.Id,
		},
		Name:    updateDishWithStepsRequest.Name,
		Cuisine: updateDishWithStepsRequest.Cuisine,
		Steps:   updateDishWithStepsRequest.Steps,
		CustomStepsList: map[string][]map[string]interface{}{
			uuid.New().String(): updateDishWithStepsRequest.Steps,
			uuid.New().String(): updateDishWithStepsRequest.Steps,
			uuid.New().String(): updateDishWithStepsRequest.Steps,
		},
	}

	if err := global.FXDb.Model(&dish).Select("name", "cuisine", "steps", "custom_steps_list").Updates(dish).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	updateDishWithStepsResponse := response.UpdateDishWithSteps{
		Dish: model.DishInfo{
			Id:              dish.Id,
			Image:           "data:image/png;base64," + base64.StdEncoding.EncodeToString(dish.Image),
			Name:            dish.Name,
			UUID:            dish.UUID,
			Steps:           dish.Steps,
			CustomStepsList: dish.CustomStepsList,
			Cuisine:         dish.Cuisine,
		},
	}

	response.SuccessMessageData(c, updateDishWithStepsResponse, "更新成功")
}

func (api *DishApi) Delete(c *gin.Context) {
	var deleteDishesRequest request.DeleteDishes
	if err := request.ShouldBindJSON(c, &deleteDishesRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var dishes []model.SysDish
	for _, id := range deleteDishesRequest.Ids {
		dishes = append(dishes, model.SysDish{
			FXModel: global.FXModel{
				Id: id,
			},
		})
	}

	if err := global.FXDb.Delete(&dishes).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "删除成功")
}

func (api *DishApi) DeletePersonals(c *gin.Context) {
	var deleteDishesRequest request.DeleteDishes
	if err := request.ShouldBindJSON(c, &deleteDishesRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var deletedDishes []model.SysDish
	for _, id := range deleteDishesRequest.Ids {
		deletedDishes = append(deletedDishes, model.SysDish{
			FXModel: global.FXModel{
				Id: id,
			},
		})
	}

	if err := global.FXDb.Clauses(clause.Returning{Columns: []clause.Column{{Name: "uuid"}, {Name: "owner"}}}).
		Delete(&deletedDishes).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	var deletedPersonalDishes []model.SysUserDeletedDish
	for _, dish := range deletedDishes {
		deletedPersonalDishes = append(deletedPersonalDishes, model.SysUserDeletedDish{
			UUID:  dish.UUID,
			Owner: dish.Owner,
		})
	}

	if err := global.FXDb.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "uuid"}}, DoNothing: true}).
		Create(&deletedPersonalDishes).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "删除成功")
}

func (api *DishApi) Add(c *gin.Context) {
	var addDishRequest request.AddDish
	if err := request.ShouldBindJSON(c, &addDishRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	imageData, err := utils.LoadLocalImage("./assets/default_dish_image.png")
	if err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	dish := model.SysDish{
		FXModel: global.FXModel{
			Sort: 1,
		},
		Name:    addDishRequest.Name,
		Cuisine: addDishRequest.Cuisine,
		UUID:    uuid.New(),
		Steps:   addDishRequest.Steps,
		CustomStepsList: map[string][]map[string]interface{}{
			uuid.New().String(): addDishRequest.Steps,
			uuid.New().String(): addDishRequest.Steps,
			uuid.New().String(): addDishRequest.Steps,
		},
		Image:      imageData,
		IsOfficial: true,
		IsShared:   true,
		IsMarked:   false,
	}

	if err := global.FXDb.Create(&dish).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	addDishResponse := response.AddDish{
		Dish: model.DishInfo{
			Id:              dish.Id,
			Image:           "data:image/png;base64," + base64.StdEncoding.EncodeToString(dish.Image),
			Name:            dish.Name,
			UUID:            dish.UUID,
			Steps:           dish.Steps,
			CustomStepsList: dish.CustomStepsList,
			Cuisine:         dish.Cuisine,
			IsOfficial:      dish.IsOfficial,
			IsShared:        dish.IsShared,
			IsMarked:        dish.IsMarked,
		},
	}

	response.SuccessMessageData(c, addDishResponse, "添加成功")
}

func (api *DishApi) Copy(c *gin.Context) {
	var copyDishRequest request.CopyDish
	if err := request.ShouldBindJSON(c, &copyDishRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var dish model.SysDish
	if err := global.FXDb.Where("id = ?", copyDishRequest.Id).First(&dish).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	dish.Id = 0
	dish.CreatedAt = 0
	dish.UpdatedAt = 0
	dish.Name = dish.Name + "-副本"
	dish.UUID = uuid.New()
	dish.CustomStepsList = map[string][]map[string]interface{}{
		uuid.New().String(): dish.Steps,
		uuid.New().String(): dish.Steps,
		uuid.New().String(): dish.Steps,
	}

	if err := global.FXDb.Create(&dish).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	copyDishResponse := response.CopyDish{
		Dish: model.DishInfo{
			Id:              dish.Id,
			Image:           "data:image/png;base64," + base64.StdEncoding.EncodeToString(dish.Image),
			Name:            dish.Name,
			UUID:            dish.UUID,
			Steps:           dish.Steps,
			CustomStepsList: dish.CustomStepsList,
			Cuisine:         dish.Cuisine,
			IsOfficial:      dish.IsOfficial,
			IsShared:        dish.IsShared,
			IsMarked:        dish.IsMarked,
		},
	}

	response.SuccessMessageData(c, copyDishResponse, "复制成功")
}

func (api *DishApi) Topping(c *gin.Context) {
	var toppingDishRequest request.ToppingDish
	if err := request.ShouldBindJSON(c, &toppingDishRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	topDish := model.SysDish{
		FXModel: global.FXModel{
			Id: toppingDishRequest.Id,
		},
	}

	if err := global.FXDb.Model(&topDish).Update("sort", 1).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	var dishes []model.SysDish
	if err := global.FXDb.Where("is_official", true).Not("id = ?", toppingDishRequest.Id).
		Order("sort, updated_at asc").Select("id", "sort").Find(&dishes).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
	}

	for index, dish := range dishes {
		if err := global.FXDb.Model(&dish).Update("sort", int64(index)+2).Error; err != nil {
			global.FXLogger.Error(err.Error())
			response.ErrorMessage(c, err.Error())
			return
		}
	}

	response.SuccessMessage(c, "置顶成功")
}

func (api *DishApi) UpdateImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}
	defer src.Close()

	// Read the file content into a byte slice
	fileData := make([]byte, file.Size)
	_, err = src.Read(fileData)
	if err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}
	idStr := c.PostForm("id")

	id, _ := strconv.Atoi(idStr)
	dish := model.SysDish{
		FXModel: global.FXModel{Id: uint(id)},
		Image:   fileData,
	}

	if err := global.FXDb.Model(&dish).Update("image", dish.Image).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	updatedImage := "data:image/png;base64," + base64.StdEncoding.EncodeToString(dish.Image)

	response.SuccessMessageData(c, updatedImage, "更新菜谱图片成功")
}

func (api *DishApi) getOfficialDishNumber(c *gin.Context) (int64, error) {
	var officialDishNumber int64
	if err := global.FXDb.Model(&model.SysDish{}).Where("is_official", true).Count(&officialDishNumber).Error; err != nil {
		response.ErrorMessage(c, err.Error())
		return 0, err
	}
	return officialDishNumber, nil
}

func (api *DishApi) ExportSteps(c *gin.Context) {
	var exportStepsRequest request.ExportSteps
	if err := request.ShouldBindJSON(c, &exportStepsRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var dishes []model.SysDish
	if err := global.FXDb.Where("id in (?)", exportStepsRequest.Ids).Select("name", "steps").Find(&dishes).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	var dishInfos []model.DishInfo
	for _, dish := range dishes {
		dishInfos = append(dishInfos, model.DishInfo{Name: dish.Name, Steps: dish.Steps})
	}

	exportStepsResponse := response.ExportSteps{
		Dishes: dishInfos,
	}

	response.SuccessData(c, exportStepsResponse)
}

func (api *DishApi) ListOwners(c *gin.Context) {
	var listOwnersRequest request.ListOwners
	if err := request.ShouldBindQuery(c, &listOwnersRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var dishes []model.SysDish
	if err := global.FXDb.Where("owner != ''").Select("owner").Group("owner").Find(&dishes).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	var owners []string
	for _, dish := range dishes {
		owners = append(owners, dish.Owner)
	}

	listOwnersResponse := response.ListOwners{
		Owners: owners,
	}

	response.SuccessData(c, listOwnersResponse)
}

func (api *DishApi) AddToOfficials(c *gin.Context) {
	var addToOfficialsRequest request.AddToOfficials
	if err := request.ShouldBindJSON(c, &addToOfficialsRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var dish model.SysDish
	if err := global.FXDb.Where("id = ?", addToOfficialsRequest.Id).Omit("id", "updated_at", "created_at").
		First(&dish).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	dish.UUID = uuid.New()
	dish.IsOfficial = true
	dish.IsShared = true
	dish.IsMarked = false
	dish.Sort = 1

	if err := global.FXDb.Create(&dish).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "已添加至官方菜品")
}

func (api *DishApi) GetQrCode(c *gin.Context) {
	var getQrCodeRequest request.GetDishQrCode
	if err := request.ShouldBindQuery(c, &getQrCodeRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var dish model.SysDish
	if err := global.FXDb.Where("id = ?", getQrCodeRequest.Id).Select("uuid").First(&dish).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	qr, err := qrcode.New("dishUUID::"+dish.UUID.String(), qrcode.Medium)
	if err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, qr.Image(256)); err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	getQrCodeResponse := response.GetQrCode{QrCode: buf.Bytes()}

	response.SuccessMessageData(c, getQrCodeResponse, "获取二维码成功")
}