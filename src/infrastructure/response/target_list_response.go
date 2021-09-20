package response

import "github.com/mihnealun/commentix/domain/entity"

type TargetListResponse struct {
	Targets []TargetResponse `json:"targets"`
}

func NewTargetListResponse(targets []*entity.Target) TargetListResponse {
	result := TargetListResponse{
		Targets: []TargetResponse{},
	}

	for _, t := range targets {
		result.Targets = append(result.Targets, NewTargetResponse(t))
	}

	return result
}
