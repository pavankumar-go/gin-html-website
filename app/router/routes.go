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
	app.GET("/blogs", handler.Blogs())
	app.GET("/places", handler.Places())
	app.GET("/admin/bird/upload", handler.AdminAPIBirdUpload())
	app.GET("/admin/place/upload", handler.AdminAPIPlaceUpload())
	app.GET("/admin/place/update", handler.AdminAPIPlacePatch())

	places := app.Group("/places")
	places.GET("/blr", handler.Bangalore())
	places.GET("/mandya", handler.Mandya())

	// places.GET("/gaganachukki", handler.Gaganachukki())
	// places.GET("/valleySchool", handler.Ganeshgudi())
	// places.GET("/ganeshGudi", handler.ValleySchool())

	adminAPI := app.Group("/v1")
	adminAPI.Use(middleware.Auth())
	adminAPI.POST("/bird/upload", api.AddBird())
	// adminAPI.DELETE("/bird/:birdId/place/:placeId", api.AddPlace()) - refer handler comments frontend.go

	adminAPI.POST("/place/create", api.AddPlace())
	adminAPI.DELETE("/place/:id", api.DeletePlace())
	adminAPI.PATCH("/place/:id", api.UpdatePlace())

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
