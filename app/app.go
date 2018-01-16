package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/plagiari-sm/api-gin-couchbase-boilerplate/controllers"
	"github.com/plagiari-sm/api-gin-couchbase-boilerplate/db"
	"github.com/plagiari-sm/api-gin-couchbase-boilerplate/middleware"
	psmconfig "github.com/plagiari-sm/psm-config"
)

// APP : Application Structure
type APP struct {
	Router *gin.Engine
	Server *http.Server
}

// Initialize : Initialize Application Components
func (app *APP) Initialize() {
	gin.SetMode(gin.ReleaseMode)

	cfg := psmconfig.Config
	// CouchBase
	bucket := db.ConnectCouchBase(cfg.CouchBase.Host,
		cfg.CouchBase.User, cfg.CouchBase.Pass, cfg.CouchBase.Path)
	// Assigin the router
	app.Router = gin.Default()
	// Inital router settings
	app.Router.RedirectTrailingSlash = true
	app.Router.RedirectFixedPath = true

	// Basic router middlewares
	if cfg.Env == "development" {
		app.Router.Use(middleware.EnableCORS())
	}
	app.Router.Use(middleware.ErrorHandler)

	// Demo Endpoints
	v1 := app.Router.Group("/v1/articles")
	{
		articleCtrl := &controllers.ArticleCTRL{
			Bucket: bucket,
		}
		// CRUD
		v1.POST("/", articleCtrl.Create)
		v1.GET("/", articleCtrl.Read)
		v1.GET("/:id", articleCtrl.ReadOne)
		v1.PUT("/:id", articleCtrl.Update)
		v1.DELETE("/:id", articleCtrl.Delete)
	}

	// Assigin the http server
	app.Server = &http.Server{
		Addr:    cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Handler: app.Router,
	}
}

// Serve : Serve the Application with Error Channels
func (app *APP) Serve() {
	errChan := make(chan error, 10)

	go func() {
		errChan <- app.Server.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured message %v. Exiting...", s))
			os.Exit(0)
		}
	}
}
