package ssignal

import (
	"os"
	"os/signal"
)

//返回值表示是否再次等待信号，如果为true表示等待，为false表示不再等待
type HandleSignal func(s os.Signal) bool

func WaitSignal(handle HandleSignal, sig ...os.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, sig...)
	for s := range c {
		if handle == nil || !handle(s) {
			signal.Stop(c)
			close(c)
			break
		}
	}
}

func WatiCtrlC(handle HandleSignal) {
	WaitSignal(handle, os.Interrupt, os.Kill)
}
