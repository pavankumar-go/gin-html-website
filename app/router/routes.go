package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/handler"
	"github.com/gin-html-website/app/handler/api"
	"github.com/gin-html-website/app/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartServer registers routes and starts server
func StartServer(ctx context.Context, appPath string) {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(middleware.CORS())

	app.Static("static/css", appPath+"/static/css")
	app.Static("static/fonts", appPath+"/static/fonts")
	app.Static("scripts", appPath+"/scripts")
	app.Static("static/assets/images", appPath+"/static/assets/images")
	app.Static("static/assets/sounds", appPath+"/static/assets/sounds")
	app.LoadHTMLGlob("templates/*/*.html")

	// prometheus handler
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// healthz handler
	app.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "up"})
	})

	app.GET("/", handler.HomePage())
	app.GET("/home", handler.HomePage())
	app.GET("/about", handler.About())
	// app.GET("/blogs", handler.Blogs())
	app.GET("/gallery", handler.Gallery())
	app.GET("/wildlife/places", handler.WildlifePlaces())
	app.GET("/landscape/places", handler.LandscapePlaces())

	// to add new route for a place
	wildlifePlaces := app.Group("/wildlife/places")
	wildlifePlaces.GET("/blr", handler.W_Places(1))
	wildlifePlaces.GET("/mandya", handler.W_Places(2))
	wildlifePlaces.GET("/jbr", handler.W_Places(3))
	
	// to add new route for a place
	landscapePlaces := app.Group("/landscape/places")
	landscapePlaces.GET("/blr", handler.L_Places(1))
	landscapePlaces.GET("/mandya", handler.L_Places(2))
	landscapePlaces.GET("/jbr", handler.L_Places(3))


	app.GET("/admin/wildlife/upload", handler.AdminAPIWildlifeUpload())
	app.GET("/admin/wildlife/place/upload", handler.AdminAPIWildlifePlaceUpload())
	app.GET("/admin/wildlife/place/update", handler.AdminAPIWildlifePlacePatch())

	app.GET("/admin/landscape/upload", handler.AdminAPILandscapeUpload())
	app.GET("/admin/landscape/place/upload", handler.AdminAPILandscapePlaceUpload())
	app.GET("/admin/landscape/place/update", handler.AdminAPILandscapePlacePatch())

	adminAPI := app.Group("/v1")
	adminAPI.Use(middleware.Auth())
	adminAPI.POST("/wildlife/upload", api.AddBird())
	adminAPI.POST("/landscape/upload", api.AddLandscape())

	// adminAPI.DELETE("/bird/:birdId/place/:placeId", api.AddPlace()) - refer handler comments frontend.go

	adminAPI.POST("/place/wildlife/create", api.AddWildlifePlace())
	adminAPI.POST("/place/wildlife/delete", api.DeleteWildlifePlace())
	adminAPI.POST("/place/wildlife/update", api.UpdateWildlifePlace())

	adminAPI.POST("/place/landscape/create", api.AddLandscapePlace())
	adminAPI.POST("/place/landscape/delete", api.DeleteLandscapePlace())
	adminAPI.POST("/place/landscape/update", api.UpdateLandscapePlace())
	// 404 route

	server := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Server closed under request or interrupt")
				return
			}
			log.Fatalln("Unexpected error: ", err)
		}
	}()

	<-ctx.Done()
	shtdwnCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if err := server.Shutdown(shtdwnCtx); err != nil {
		log.Fatalln("Server forced to shutdown: ", err)
	}
	defer cancel()
	log.Println("Server shutdown complete")
}
