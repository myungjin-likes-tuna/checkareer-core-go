package skills

import (
	"checkareer-core/modules/skills"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	// Handler 스킬 핸들러
	Handler interface {
		Get(context echo.Context) error
		Post(context echo.Context) error
		Patch(context echo.Context) error
		Delete(context echo.Context) error
	}
	handler struct {
		skills.Creater
		skills.Reader
		skills.Updater
		skills.Deleter
	}
)

// NewHandler 스킬 핸들러 생성
func NewHandler(
	creater skills.Creater,
	reader skills.Reader,
	updater skills.Updater,
	deleter skills.Deleter,
) Handler {
	return &handler{
		creater,
		reader,
		updater,
		deleter,
	}
}

// BindRoutes 스킬 핸들러에 라우팅 정보 바인딩
func BindRoutes(server *echo.Echo, handler Handler) {
	group := server.Group("/v1/skills")
	group.GET("", handler.Get)
	group.POST("", handler.Post)
	group.PATCH("/:id", handler.Patch)
	group.DELETE("/:id", handler.Delete)
}

// Get 스킬 조회
// @ID skills-get
// @Tags skills
// @Summary skills get
// @Description skills get
// @Router /v1/skills [get]
// @Param _ query Params true "params"
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} ManyItemResponse
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h handler) Get(context echo.Context) error {
	var param Params
	if err := context.Bind(&param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(&param); err != nil {
		return err
	}
	_skills, err := h.Read(skills.WithID(param.ID), skills.WithLimit(param.Limit))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response := ManyItemResponse{
		Skills: _skills,
		Limit:  param.Limit,
	}
	return context.JSON(http.StatusOK, response)
}

// Post 스킬 생성
func (h handler) Post(context echo.Context) error {
	return nil
}

// Patch 스킬 수정
func (h handler) Patch(context echo.Context) error {
	return nil
}

// Delete 스킬 삭제
func (h handler) Delete(context echo.Context) error {
	return nil
}
