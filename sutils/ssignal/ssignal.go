// Scry Info.  All rights reserved.
// license that can be found in the license file.

package ssignal

import (
	"os"
	"os/signal"
)

//返回值表示是否再次等待信号，如果为true表示等待，为false表示不再等待
type HandleSignal func(s os.Signal) bool

// 如果handle为空或 handle的返回值为false，那么退出等待
func WaitSignal(handle HandleSignal, sig ...os.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, sig...)

	//另外一种实现是使用两个defer，也都可以的
	defer func() {
		signal.Stop(c)
		close(c)
	}()

	for s := range c {
		if handle == nil || !handle(s) {
			break
		}
	}
}

// 等待程序终止信息 ctrl + c
func WatiCtrlC(handle HandleSignal) {
	WaitSignal(handle, os.Interrupt, os.Kill)
}
