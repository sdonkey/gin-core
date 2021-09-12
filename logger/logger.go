/*
 * Copyright(C),2021-2026
 * @Author: wangjr<wangjr12@163.com>
 * @Date: 2021/9/10 2:32 上午
 */
package logger

var (
	DefaultLogger Logger
)

type Logger interface {
	Options() Options
}


