package skills

// Params of skills
type Params struct {
	ID    int64 `query:"id"    json:"id,omitempty"    validate:"gte=0"`
	Limit uint  `query:"limit" json:"limit,omitempty" validate:"gte=1" default:"25" example:"25"`
}
