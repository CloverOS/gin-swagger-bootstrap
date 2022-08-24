package ginSwaggerBootstrap

import (
	"github.com/CloverOS/gin-swagger-bootstrap/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
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
		case "group.json":
			readFile, err := bootstrap.ReadFile(path)
			if err != nil {
				bootstrap.Handler.ServeHTTP(ctx.Writer, ctx.Request)
				return
			}
			ctx.String(200, string(readFile))
		case "index.html":
			readFile, err := bootstrap.ReadFile(path)
			if err != nil {
				bootstrap.Handler.ServeHTTP(ctx.Writer, ctx.Request)
				return
			}
			ctx.String(200, string(readFile))
		default:
			tmp := strings.Replace(ctx.Request.URL.Path, "/swagger/", "", 1)
			readFile, err := bootstrap.ReadFile(tmp)
			if err != nil {
				bootstrap.Handler.ServeHTTP(ctx.Writer, ctx.Request)
				return
			}
			ctx.String(200, string(readFile))
		}
	}
}
