package main

import (
	"github.com/hulunbao/golang/learngo/common"
	"sync"
)

const (
	// 模拟最大的goroutine
	maxGoroutine = 5
	// 资源池大小
	poolRes = 2
)

func main() {
	// 等待任务完成
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p, err := common.New1(2)
}
