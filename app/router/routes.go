package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/handler"
	"github.com/gin-html-website/app/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartServer registers routes and starts server
func StartServer(ctx context.Context) {

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(middleware.CORS())

	app.Static("static/css", "/Users/pavan/go/src/github.com/gin-html-website/static/css")
	app.Static("static/assets/images", "/Users/pavan/go/src/github.com/gin-html-website/static/assets/images")
	app.Static("static/assets/sounds", "/Users/pavan/go/src/github.com/gin-html-website/static/assets/sounds")
	app.LoadHTMLGlob("templates/*.html")

	// prometheus handler
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// healthz handler
	app.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "up"})
	})

	app.GET("/", handler.HomePage())

	// gin.SetMode()
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
