package api

import (
	"fmt"
	db "github.com/Pizhlo/go-shop/db/sqlc"
	"github.com/Pizhlo/go-shop/token"
	"github.com/Pizhlo/go-shop/util"
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.Static("static/", "./static/")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Главная страница"})
	})

	router.POST("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Главная страница"})
	})

	router.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{"title": "Регистрация"})
	})

	router.POST("/users", server.createUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Регистрация"})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth.html", gin.H{"title": "Авторизация"})
	})
	router.POST("/login", server.loginUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Главная страница"})
	})

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/account/:username", server.getUser, func(c *gin.Context) {
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
