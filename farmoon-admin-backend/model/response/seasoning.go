package response

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"

type ListSeasonings struct {
	Seasonings []model.SeasoningInfo `json:"seasonings"`
}

type AddSeasoning struct {
	Seasoning model.SeasoningInfo `json:"seasoning"`
}
