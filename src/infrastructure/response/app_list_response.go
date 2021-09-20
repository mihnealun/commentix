package response

import "github.com/mihnealun/commentix/domain/entity"

type AppListResponse struct {
	Apps []AppResponse `json:"apps"`
}

func NewAppListResponse(apps []*entity.App) AppListResponse {
	result := AppListResponse{
		Apps: []AppResponse{},
	}

	for _, app := range apps {
		result.Apps = append(result.Apps, NewAppResponse(app))
	}

	return result
}
