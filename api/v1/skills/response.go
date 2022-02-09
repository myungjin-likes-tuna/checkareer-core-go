package skills

import "checkareer-core/modules/skills"

// ManyItemResponse 여러 개의 아이템을 반환하는 응답
type ManyItemResponse struct {
	Skills []skills.Node `json:"skills"`
	Limit  uint          `json:"limit"`
}
