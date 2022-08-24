package ginSwaggerBootstrap

import (
	"gin-swagger-bootstrap/file"
	"github.com/CloverOS/swag-gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

func WrapHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodGet {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)

			return
		}
		i := strings.LastIndex(ctx.Request.URL.Path, "/")
		if i == -1 {
			return
		}
		path := ctx.Request.URL.Path[i+1:]
		if path == "" {
			path = "index.html"
		}
		switch filepath.Ext(path) {
		case ".html":
			ctx.Header("Content-Type", "text/html; charset=utf-8")
		case ".css":
			ctx.Header("Content-Type", "text/css; charset=utf-8")
		case ".js":
			ctx.Header("Content-Type", "application/javascript")
		case ".png":
			ctx.Header("Content-Type", "image/png")
		case ".json":
			ctx.Header("Content-Type", "application/json; charset=utf-8")
		}
		switch path {
		case "swagger.json":
			doc, err := swag.ReadDoc("swagger")
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			ctx.String(http.StatusOK, doc)
		default:
			readFile, err := file.ReadFile(path)
			if err != nil {
				file.Handler.ServeHTTP(ctx.Writer, ctx.Request)
				return
			}
			ctx.String(200, string(readFile))
		}
	}
}
