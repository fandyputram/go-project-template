package http

import (
	"net/http"

	"github.com/fandyputram/go-project-template/internal/usecase"
	"github.com/fandyputram/go-project-template/pkg/middleware"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *gin.Engine {
	h := &Handler{usecase: uc}
	r := gin.Default()

	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.GET("/user/:id", h.GetUser)

	r.POST("/login", h.Login)
	r.GET("/user/:id", h.GetUser)

	// Apply JWT middleware
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	auth.GET("/protected", h.ProtectedEndpoint)

	return r
}

func (h *Handler) Login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.usecase.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Register(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.usecase.Register(user.Username, user.Password, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
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

func (h *Handler) ProtectedEndpoint(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	c.JSON(200, gin.H{"message": "Hello " + userID})
}
