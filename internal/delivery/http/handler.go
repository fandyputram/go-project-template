package http

import (
	"net/http"

	"github.com/fandyputram/go-project-template/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *gin.Engine {
	h := &Handler{usecase: uc}
	r := gin.Default()

	r.GET("/user/:id", h.GetUser)

	return r
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.usecase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
