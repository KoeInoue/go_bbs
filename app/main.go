package main

import (
	"go_bbs/db"
	"go_bbs/server"
	"net/http"
	"os"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	db.Init()
	server.GetEnv()

	if os.Getenv("ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	// セッションの設定
	store := cookie.NewStore([]byte(os.Getenv("ACCESS_SECRET")))
	r.Use(sessions.Sessions("auth", store))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},

		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Origin",
		},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
	}))

	r = server.Router(r)
	// show Landing Page
	r.Use(static.Serve("/", BinaryFileSystem("assets")))

	// Not Found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not Found"})
	})
	r.Run()
	db.Close()
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
