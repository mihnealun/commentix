package response

import "github.com/mihnealun/commentix/domain/entity"

type AppResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func NewAppResponse(app *entity.App) AppResponse {
	return AppResponse{
		ID:   app.UUID,
		Name: app.Name,
		Slug: app.Slug,
	}
}
