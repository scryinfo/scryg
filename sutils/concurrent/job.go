package concurrent

import (
	"context"
	"sync/atomic"
	"time"
)

/// 做一个job并等待超时，  如果duration时到还没有完成，会直接返回（但job依然在go routine中运行）
/// 返回值 true, done, false time out
/// duration 超时时间
func DoJobWaitTimeOut(job func(), duration time.Duration) bool {
	done := false
	closed := uint32(0)
	doneChannel := make(chan bool)
	closeChannel := func() {
		if atomic.CompareAndSwapUint32(&closed, 0, 1) {
			close(doneChannel)
			doneChannel = nil
		}
	}

	go func() {
		defer closeChannel()
		job()
	}()

	select {
	case <-doneChannel:
		done = true
	case <-time.After(duration):
		done = false
		closeChannel()
	}

	return done
}

// 做一个job，等待超时或ctx被取消
// 返回值 true done/ the ctx done, false time out
// duration 超时时间
func DoJobWaitContext(ctx context.Context, job func(), duration time.Duration) bool {
	done := false
	closed := uint32(0)
	doneChannel := make(chan bool)
	closeChannel := func() {
		if atomic.CompareAndSwapUint32(&closed, 0, 1) {
			close(doneChannel)
			doneChannel = nil
		}
	}

	go func() {
		defer closeChannel()
		job()
	}()

	select {
	case <-doneChannel:
		done = true
	case <-time.After(duration):
		done = false
		closeChannel()
	case <-ctx.Done():
		done = true
		closeChannel()
	}

	return done
}
