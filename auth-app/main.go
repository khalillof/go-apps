package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"html/template"
	//"strings"

	globals "github.com/khalillof/go-apps/auth-app/globals"
	middleware "github.com/khalillof/go-apps/auth-app/middleware"
	routes "github.com/khalillof/go-apps/auth-app/routes"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "auth-app/assets")
	router.LoadHTMLGlob("auth-app/templates/**")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run("localhost:8080")
}