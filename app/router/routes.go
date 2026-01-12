package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

	app.StaticFile("/sitemap.xml", appPath+"/sitemap.xml")

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

	app.GET("/", handler.HomePage(appPath))
	app.GET("/home", handler.HomePage(appPath))
	app.GET("/about", handler.About())
	// app.GET("/blogs", handler.Blogs())
	app.GET("/gallery", handler.Gallery())
	// app.GET("/wildlife/places", handler.WildlifePlaces())
	// app.GET("/landscape/places", handler.LandscapePlaces())

	wildlifePlacesGrp := app.Group("/wildlife/places")
	// landscapePlacesGrp := app.Group("/landscape/places")

	// to add new route for a place
	places := strings.Split(os.Getenv("PLACES"), ",")
	if len(places) == 1 {
		log.Fatalln("PLACES unset, no routes will be exposed.")
	}

	// to add new route for a place
	for index, p := range places {
		wildlifePlacesGrp.GET(fmt.Sprintf("/%s", p), handler.W_Places(index+1))
		// landscapePlacesGrp.GET(fmt.Sprintf("/%s", p), handler.L_Places(index+1))
	}

	app.GET("/admin/wildlife/upload", handler.AdminAPIWildlifeUpload())
	app.GET("/admin/wildlife/place/upload", handler.AdminAPIWildlifePlaceUpload())
	app.GET("/admin/wildlife/place/update", handler.AdminAPIWildlifePlacePatch())

	// app.GET("/admin/landscape/upload", handler.AdminAPILandscapeUpload())
	// app.GET("/admin/landscape/place/upload", handler.AdminAPILandscapePlaceUpload())
	// app.GET("/admin/landscape/place/update", handler.AdminAPILandscapePlacePatch())

	adminAPI := app.Group("/v1")
	adminAPI.Use(middleware.Auth())
	adminAPI.POST("/wildlife/upload", api.AddBird())
	// adminAPI.POST("/landscape/upload", api.AddLandscape())

	// adminAPI.DELETE("/bird/:birdId/place/:placeId", api.AddPlace()) - refer handler comments frontend.go

	adminAPI.POST("/place/wildlife/create", api.AddWildlifePlace())
	adminAPI.POST("/place/wildlife/delete", api.DeleteWildlifePlace())
	// adminAPI.POST("/place/wildlife/update", api.UpdateWildlifePlace())
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
