// Package service
// Date: 2024/06/10 18:34:36
// Author: Amu
// Description:
package service

import "log/slog"

func Run(configFile string) (func(), error) {
	injector, clearFunc, err := BuildInjector(configFile)
	if err != nil {
		slog.Error("build injector failed: %v", "err", err)
		return nil, err
	}

	// 初始化日志
	slog.SetDefault(injector.Logger.Logger)

	// 定时任务
	timedTask := injector.Task
	go timedTask.Run()

	return func() {
		timedTask.Stop()
		clearFunc()
	}, nil
}
