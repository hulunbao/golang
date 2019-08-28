package common

import (
	"errors"
	"io"
	"log"
	"sync"
)

// ErrPoolClosed 资源池已经关闭
var ErrPoolClosed = errors.New("资源池已经关闭")

// Pool 资源池结构体
type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

//New 创建一个资源池
func New1(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size 的值太小了。")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

// Acquire 从资源池里获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

// Close 关闭资源池,释放资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}

	p.closed = true

	// 关闭通道，不让写入
	close(p.res)
	for r := range p.res {
		r.Close()
	}
}

// Release 释放资源
func (p *Pool) Release(r io.Closer) {
	// 保证该操作和Close方法的操作是安全的

	p.m.Lock()
	defer p.m.Unlock()

	// 资源池都关闭了,就剩这一个没有释放的资源了，释放即可
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了,释放这个资源吧")
		r.Close()
	}
}
