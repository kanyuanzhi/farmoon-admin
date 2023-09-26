package private

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/request"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/response"
)

type SeasoningApi struct {
}

func (api *SeasoningApi) List(c *gin.Context) {
	var seasonings []model.SysSeasoning
	if err := global.FXDb.Order("pump").Find(&seasonings).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	seasoningsInfo := []model.SeasoningInfo{}
	for _, seasoning := range seasonings {
		seasoningsInfo = append(seasoningsInfo, model.SeasoningInfo{
			Id:    seasoning.Id,
			Name:  seasoning.Name,
			Pump:  seasoning.Pump,
			Ratio: seasoning.Ratio,
		})
	}

	listSeasoningsResponse := response.ListSeasonings{
		Seasonings: seasoningsInfo,
	}

	response.SuccessData(c, listSeasoningsResponse)
}

func (api *SeasoningApi) Update(c *gin.Context) {
	var updateSeasoningRequest request.UpdateSeasoning
	if err := request.ShouldBindJSON(c, &updateSeasoningRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	seasoning := model.SysSeasoning{
		FXModel: global.FXModel{
			Id: updateSeasoningRequest.Id,
		},
		Name:  updateSeasoningRequest.Name,
		Pump:  updateSeasoningRequest.Pump,
		Ratio: updateSeasoningRequest.Ratio,
	}

	if err := global.FXDb.Model(&seasoning).Select("name", "pump", "ratio").Updates(seasoning).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "更新成功")
}

func (api *SeasoningApi) Delete(c *gin.Context) {
	var deleteSeasoningRequest request.DeleteSeasoning
	if err := request.ShouldBindJSON(c, &deleteSeasoningRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	seasoning := model.SysSeasoning{
		FXModel: global.FXModel{
			Id: deleteSeasoningRequest.Id,
		},
	}

	if err := global.FXDb.Delete(&seasoning).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "删除成功")
}

func (api *SeasoningApi) Add(c *gin.Context) {
	var addSeasoningRequest request.AddSeasoning
	if err := request.ShouldBindJSON(c, &addSeasoningRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	seasoning := model.SysSeasoning{
		Name:  addSeasoningRequest.Name,
		Pump:  addSeasoningRequest.Pump,
		Ratio: addSeasoningRequest.Ratio,
	}

	if err := global.FXDb.Create(&seasoning).Error; err != nil {
		global.FXLogger.Error(err.Error())
		response.ErrorMessage(c, err.Error())
		return
	}

	addSeasoningResponse := response.AddSeasoning{
		Seasoning: model.SeasoningInfo{
			Id:    seasoning.Id,
			Name:  seasoning.Name,
			Pump:  seasoning.Pump,
			Ratio: seasoning.Ratio,
		},
	}

	response.SuccessMessageData(c, addSeasoningResponse, "添加菜系成功")
}
