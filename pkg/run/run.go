package run

import (
	"mini_go_docker/pkg/container"
	"os"

	log "github.com/sirupsen/logrus"
)

// Run    run -it /bin/bash
/*
这里的Start方法是真正开始前面创建好的 command 的调用，
它首先会clone出来一个namespace隔离的进程，然后在子进程中，调用/proc/self/exe,也就是自己调用自己
发送 init 参数，调用我们写的 init 方法，去初始化容器的一些资源
*/
func Run(tty bool, cmd string) {
	//这里进程已经启动完成
	parent := container.NewParentProcess(tty, cmd)
	if err := parent.Start(); err != nil {
		log.Errorf("cmd start err: %v", err)
	}
	if err := parent.Wait(); err != nil {
		log.Errorf("cmd wait err: %v", err)
	}
	os.Exit(-1)
}
