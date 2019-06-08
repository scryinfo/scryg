// Scry Info.  All rights reserved.
// license that can be found in the license file.

package ssignal

import (
	"os"
	"os/signal"
)

//the return value indicates wether to wait for the signal agin. if it is true, it means waiting; if it is false, it means no waiting
type HandleSignal func(s os.Signal) bool

//if the handle is empty or the return value of handle is false, please exsit waiting
func WaitSignal(handle HandleSignal, sig ...os.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, sig...)

	//the other implementation is to use two defers which are possible
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

// waiting for the program termination information ctrl + c
func WaitCtrlC(handle HandleSignal) {
	WaitSignal(handle, os.Interrupt, os.Kill)
}
