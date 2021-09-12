/*
 * Copyright(C),2021-2026
 * @Author: wangjr<wangjr12@163.com>
 * @Date: 2021/9/9 3:35 下午
1. 语言，
 */
package logger

import (
	"fmt"
	"go.uber.org/zap"
)


func SetupLogger(appName string, fileName string) Logger{
	return NewJSONLogger(
		WithDisableConsole(),
		WithField("domain", fmt.Sprintf("%s", appName)),
		WithTimeLayout("2006-01-02 15:04:05"),
		WithFileP(fileName),
		)
}