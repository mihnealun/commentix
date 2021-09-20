package response

import "github.com/mihnealun/commentix/domain/entity"

type TargetResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

func NewTargetResponse(target *entity.Target) TargetResponse {
	return TargetResponse{
		ID:   target.UUID,
		Name: target.Name,
		Type: target.Type,
		Url:  target.Url,
	}
}
