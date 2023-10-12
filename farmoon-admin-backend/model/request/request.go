package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

func ShouldBindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		global.FXLogger.Error(global.FXConfig.System.BindErrorMessage, zap.Any("err", err))
		return err
	}
	return nil
}

func ShouldBindQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		global.FXLogger.Error(global.FXConfig.System.BindErrorMessage, zap.Any("err", err))
		return err
	}
	return nil
}

func GenerateUserQueryCondition(filter string, enableGenderFilter bool, genderFilter string, enableRolesFilter bool, rolesFilter []string) (*gorm.DB, error) {
	filterDb := global.FXDb.Model(&model.SysUser{})

	if enableGenderFilter {
		filterDb = filterDb.Where("gender = ?", genderFilter)
	}

	if enableRolesFilter {
		if rolesFilter[0] == "" {
			// 空字符串split后得到数组长度为1！
			filterDb = filterDb.Where("array_length(roles, 1) is NULL")
		} else {
			switch len(rolesFilter) {
			case 1:
				filterDb = filterDb.Where("? LIKE ANY(roles)", rolesFilter[0])
			case 2:
				filterDb = filterDb.Where("? LIKE ANY(roles)", rolesFilter[0]).
					Where("? LIKE ANY(roles)", rolesFilter[1])
			case 3:
				filterDb = filterDb.Where("? LIKE ANY(roles)", rolesFilter[0]).
					Where("? LIKE ANY(roles)", rolesFilter[1]).
					Where("? LIKE ANY(roles)", rolesFilter[2])
			default:
				return nil, errors.New("权限数量错误")
			}
		}
	}

	if filter != "" {
		likeParam := "%" + filter + "%"
		filterDb = filterDb.Where("username LIKE ? or nickname LIKE ? or real_name LIKE ? or mobile LIKE ? or email LIKE ?",
			likeParam, likeParam, likeParam, likeParam, likeParam)
	}

	return filterDb, nil
}

func GenerateDishQueryCondition(filter string, enableCuisineFilter bool, cuisineFilter []string,
	enableOwnerFilter bool, ownerFilter string, isOfficial bool) (*gorm.DB, error) {
	filterDb := global.FXDb.Model(&model.SysDish{}).Where("is_official = ?", isOfficial)

	if enableCuisineFilter {
		var cuisineFilterUint []uint
		for _, cuisine := range cuisineFilter {
			cuisineId, _ := strconv.ParseUint(cuisine, 10, 32)
			cuisineFilterUint = append(cuisineFilterUint, uint(cuisineId))
		}
		filterDb = filterDb.Where("cuisine in ?", cuisineFilterUint)
	}

	if enableOwnerFilter {
		filterDb = filterDb.Where("owner = ?", ownerFilter)
	}

	if filter != "" {
		likeParam := "%" + filter + "%"
		filterDb = filterDb.Where("name LIKE ? or owner LIKE ?", likeParam, likeParam)
	}

	return filterDb, nil
}

func GenerateIngredientQueryCondition(enableTypeFilter bool, typeFilter []string) (*gorm.DB, error) {
	filerDb := global.FXDb.Model(&model.SysIngredient{})

	if enableTypeFilter {
		var typeFilterUint []uint
		for _, ingredientType := range typeFilter {
			ingredientTypeId, _ := strconv.ParseUint(ingredientType, 10, 32)
			typeFilterUint = append(typeFilterUint, uint(ingredientTypeId))
		}
		filerDb = filerDb.Where("type in ?", typeFilterUint)
	}

	return filerDb, nil
}
