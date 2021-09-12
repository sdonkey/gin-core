package runtime

import (
	"gin-core/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

type Application struct {
	dbs map[string]*gorm.DB
	engine http.Handler
	handler map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)
	mux sync.RWMutex
	routers     []Router
	middlewares map[string]interface{}
}

func NewConfig() *Application {
	return &Application{
		dbs: make(map[string]*gorm.DB),
		handler: make(map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)),
		routers: make([]Router,0),
		middlewares: make(map[string]interface{}),
	}
}



type Router struct {
	HttpMethod string
	RelativePath string
	Handler string
}

type Routers struct {
	List []Router
}

func (e *Application) SetDb(key string, db *gorm.DB){
	e.mux.Lock()
	defer e.mux.Unlock()
	e.dbs[key] = db
}

// GetDB 获取所有map里的db数据
func (e *Application) GetDb() map[string]*gorm.DB{
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.dbs
}

func (e *Application) GetDbByKey(key string) *gorm.DB  {
	e.mux.Lock()
	defer e.mux.Unlock()
	if db, ok := e.dbs["*"]; ok {
		return db
	}
	return e.dbs[key]
}

// SetEngine 设置路由引擎
func (e *Application) SetEngine(engine http.Handler) {
	e.engine = engine
}

// GetEngine 获取路由引擎
func (e *Application) GetEngine() http.Handler {
	return e.engine
}

// GetRouter 获取路由表
func (e *Application) GetRouter() []Router {
	return e.setRouter()
}


// setRouter 设置路由表
func (e *Application) setRouter() []Router {
	switch e.engine.(type) {
	case *gin.Engine:
		routers := e.engine.(*gin.Engine).Routes()
		for _, router := range routers {
			e.routers = append(e.routers, Router{RelativePath: router.Path, Handler: router.Handler, HttpMethod: router.Method})
		}
	}
	return e.routers
}


// SetLogger 设置日志组件
func (e *Application) SetLogger(log logger.Logger) {
	logger.DefaultLogger = log
}

// GetLogger 获取日志组件
func (e *Application) GetLogger() logger.Logger{
	return logger.DefaultLogger
}

// SetMiddleware 设置中间件
func (e *Application) SetMiddleware(key string, middleware interface{}) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.middlewares[key] = middleware
}

// GetMiddleware 获取所有中间件
func (e *Application) GetMiddleware() map[string]interface{} {
	return e.middlewares
}

// GetMiddlewareKey 获取对应key的中间件
func (e *Application) GetMiddlewareKey(key string) interface{} {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.middlewares[key]
}