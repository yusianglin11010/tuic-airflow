package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/tuic-airflow/internal/repository/postgres"
	"go.uber.org/zap"
)

type Handler struct {
	dbRepo postgres.DBRepo
}

func NewHandler(db postgres.DBRepo) *Handler {
	return &Handler{
		dbRepo: db,
	}
}

type getProjectsReq struct {
	ProjectID string `form:"id"`
}

type getProjectResp struct {
	Projects []projectResp `json:"projects"`
}

type projectResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}

func (h *Handler) GetProjects(ctx *gin.Context) {
	logger := ctx.MustGet("logger").(*zap.Logger)

	req := &getProjectsReq{}
	ctx.Bind(req)

	res, err := h.dbRepo.GetMarkersByID(logger, nil, req.ProjectID)
	if err != nil {
		logger.Error("get data fail", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, "get data failed")
		return
	}

	fmt.Println(res)

	resp := getProjectResp{}
	projects := []projectResp{}

	for _, p := range res {
		project := projectResp{
			ID:   p.ProjectID,
			Name: p.Name,
			Lat:  p.Lat,
			Lng:  p.Lng,
		}
		projects = append(projects, project)
	}
	resp.Projects = projects

	ctx.JSON(http.StatusOK, resp)
	return

}

func (h *Handler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alive",
	})
}
