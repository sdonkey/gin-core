/*
 * Copyright(C),2021-2026
 * @Author: wangjr<wangjr12@163.com>
 * @Date: 2021/9/8 3:10 下午
 */
package runtime

import (
	"gin-core/logger"
	"gorm.io/gorm"
	"net/http"
)

type Runtime interface {

	SetDb(key string, db *gorm.DB)
	GetDb() map[string]*gorm.DB
	GetDbByKey(key string) *gorm.DB

	// SetEngine 使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler

	GetRouter() []Router

	SetLogger(logger logger.Logger)
	GetLogger() logger.Logger

	// SetMiddleware middleware
	SetMiddleware(string, interface{})
	GetMiddleware() map[string]interface{}
	GetMiddlewareKey(key string) interface{}

}
