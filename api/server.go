package api

import (
	"fmt"
	"github.com/Pizhlo/go-shop/token"
	"github.com/Pizhlo/go-shop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Server serves HTTP requests for banking service
type Server struct {
	config     util.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	//router.Use(setUserStatus())

	router.LoadHTMLGlob("templates/*")
	router.Static("static/", "./static/")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Главная страница",
		"auth": c.Keys["authorization_payload"]})
		fmt.Println("1 keys = ", c.Keys)
	})

	router.POST("/", server.createUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Главная страница",
			"auth": c.Keys["authorization_payload"]})
		fmt.Println("2 keys = ", c.Keys)
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{"title": "Регистрация"})
	})

	router.POST("/register", server.createUser)

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth.html", gin.H{"title": "Авторизация"})
	})
	router.POST("/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/account/:id", server.getUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "account.html", gin.H{"title": "Личный кабинет",
		"user": c.Keys["user"]})
	})

	authRoutes.GET("/account/:id/orders", server.getUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "account.html", gin.H{"title": "Личный кабинет",
		"user": c.Keys["user"]})
	})

	authRoutes.GET("/account/:id/favourites", server.getUser, func(c *gin.Context) {
		c.HTML(http.StatusOK, "account.html", gin.H{"title": "Личный кабинет",
		"user": c.Keys["user"]})
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

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
