package common

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// ErrTimeOut 执行超时
var ErrTimeOut = errors.New("执行者执行超时")

// ErrInterrupt 执行中断
var ErrInterrupt = errors.New("执行者被中断")

// Runner 一个执行者
type Runner struct {
	tasks     []func(int)      //要执行的任务
	complete  chan error       //用于通知任务全部完成
	timeout   <-chan time.Time //这些任务多久完成
	interrupt chan os.Signal   //可以控制强制终止信号
}

// New 返回Runner
func New(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}

// Add 给Runner添加任务
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

// Start 开始任务
func (r *Runner) Start() error {
	// 接受系统信号
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}
