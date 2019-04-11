package main

import (
	"fmt"
	"golang/learngo/retriever/mock"
	"golang/learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,map[string]string{
		"name": "ccmouse",
		"course": "golang",
	})
}

type RetriverPoster interface {
	Retriever
	Poster
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}
func main() {
	var r Retriever
	fmt.Printf("%T %v",r,r)
	fmt.Println()

	retriver := &mock.Retriever{"this is a fake imooc.com"}
	r = retriver

	fmt.Printf("%T %v",r,r)
	fmt.Println()

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v",r,r)

	fmt.Println()
	fmt.Println("Try a session")
	fmt.Println(session(retriver))

	// fmt.Println(download(r))

	// fmt.Println(download(mock.Retriever{"this is fake imooc.com"}))

}
