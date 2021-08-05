package main

import (
	"net/http"
	"os"
	"strings"
	"viet_college_api/infrastructure"
	"viet_college_api/interfaces/controllers"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	infrastructure.GetEnv()

	if os.Getenv("ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.Use(cors.Default())

	Router = route(r)
	Router.Run()
}

func route(r *gin.Engine) *gin.Engine {
	// apiルーティング
	api := r.Group("/api")
	{
		userController := controllers.NewUserController(infrastructure.NewSqlHandler())
		api.GET("/users", func(c *gin.Context) { c.JSON(200, userController.GetUser()) })
	}

	// show Landing Page
	r.Use(static.Serve("/", BinaryFileSystem("assets")))

	// Not Found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not Found"})
	})

	return r
}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: root}
	return &binaryFileSystem{
		fs,
	}
}
