package api

import (
	"net/http"
	"fmt"
	db "github.com/Pizhlo/go-shop/db/sqlc"
	"github.com/Pizhlo/go-shop/token"
	"github.com/Pizhlo/go-shop/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main Shop"})
	})

	router.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{"title": "Регистрация"})
	})
	router.POST("/users/login", server.loginUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth.html", gin.H{"title": "Авторизация"})
	})

	router.POST("/users", server.createUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Регистрация"})
	})

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/accounts/:username", server.getUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "account.html", gin.H{"title": "Личный кабинет"})
	})

	router.Run()

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
