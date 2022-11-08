package rest

import (
	"backend/base/log"
	"backend/web/filter"
	"fmt"
	h "net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type GETHandler interface {
	GET(context *gin.Context) (interface{}, *HttpError)
}
type POSTHandler interface {
	POST(context *gin.Context) (interface{}, *HttpError)
}
type PUTHandler interface {
	PUT(context *gin.Context) (interface{}, *HttpError)
}
type DELETEHandler interface {
	DELETE(context *gin.Context) (interface{}, *HttpError)
}
type PATCHHandler interface {
	PATCH(context *gin.Context) (interface{}, *HttpError)
}
type HEADEHandler interface {
	HEAD(context *gin.Context) (interface{}, *HttpError)
}
type OPTIONSHandler interface {
	OPTIONS(context *gin.Context) (interface{}, *HttpError)
}

type Application struct {
	router    *gin.Engine
	port      int32
	routerMap map[string]interface{}
}

func (app *Application) Run() error {
	app.configLog()
	_ = app.router.SetTrustedProxies([]string{"localhost"})
	for k := range app.routerMap {
		app.get(k, app.routerMap[k])
		app.post(k, app.routerMap[k])
		app.put(k, app.routerMap[k])
		app.delete(k, app.routerMap[k])
		app.patch(k, app.routerMap[k])
		app.head(k, app.routerMap[k])
		app.options(k, app.routerMap[k])
	}
	return app.router.Run(fmt.Sprint(":", app.port))
}

func (app *Application) configLog() {
	lg := log.GetLog()
	_ = app.router.Use(filter.JWT())
	_ = app.router.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}
		lg.WithFields(logrus.Fields{
			"TimeStamp":    time.Now(),
			"Latency":      time.Since(start),
			"ClientIP":     c.ClientIP(),
			"Method":       c.Request.Method,
			"StatusCode":   c.Writer.Status(),
			"ErrorMessage": c.Errors.ByType(1).String(),
			"BodySize":     c.Writer.Size(),
			"Path":         path,
		}).Debug("finish")
	})
}

func (app *Application) Mapping(url string, handler interface{}) {
	app.routerMap[url] = handler
}

func (app *Application) get(url string, handler interface{}) {
	if h, ok := handler.(GETHandler); ok {
		app.router.GET(url, func(context *gin.Context) {
			res, err := h.GET(
				context)
			app.write(context, res, err)
		})
	}
}

func (app *Application) post(url string, handler interface{}) {
	if h, ok := handler.(POSTHandler); ok {
		app.router.POST(url, func(context *gin.Context) {
			res, err := h.POST(context)
			app.write(context, res, err)
		})
	}
}

func (app *Application) put(url string, handler interface{}) {
	if h, ok := handler.(PUTHandler); ok {
		app.router.PUT(url, func(context *gin.Context) {
			res, err := h.PUT(context)
			app.write(context, res, err)
		})
	}
}

func (app *Application) delete(url string, handler interface{}) {
	if h, ok := handler.(DELETEHandler); ok {
		app.router.DELETE(url, func(context *gin.Context) {
			res, err := h.DELETE(context)
			app.write(context, res, err)
		})
	}
}

func (app *Application) patch(url string, handler interface{}) {
	if h, ok := handler.(PATCHHandler); ok {
		app.router.PATCH(url, func(context *gin.Context) {
			res, err := h.PATCH(context)
			app.write(context, res, err)
		})
	}
}
func (app *Application) head(url string, handler interface{}) {
	if h, ok := handler.(HEADEHandler); ok {
		app.router.HEAD(url, func(context *gin.Context) {
			res, err := h.HEAD(context)
			app.write(context, res, err)
		})
	}
}

func (app *Application) options(url string, handler interface{}) {
	if h, ok := handler.(OPTIONSHandler); ok {
		app.router.OPTIONS(url, func(context *gin.Context) {
			res, err := h.OPTIONS(context)
			app.write(context, res, err)
		})
	}
}

func (app *Application) write(context *gin.Context, res interface{}, err *HttpError) {
	if err == nil {
		if res == nil {
			return
		}
		context.JSON(h.StatusOK, gin.H{
			"isSuccess": true,
			"data":      res,
		})
	} else {
		context.JSON(err.code, gin.H{
			"isSuccess": false,
			"error":     err.Error(),
		})
	}
}

func New(port int32) *Application {
	gin.SetMode(gin.ReleaseMode)
	return &Application{
		port:      port,
		router:    gin.New(),
		routerMap: make(map[string]interface{}),
	}
}
