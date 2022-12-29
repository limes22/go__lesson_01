package main

import (
	"fmt"
	"time"
)

//Goroutines : 기본적으로 다른 함수와 동시에 실행시키는 함수
//Channel : goroutine 이랑 메인함수 사이에 정보를 전달하기 위한 방법

func main() {
	c := make(chan string)
	people := [5]string{"sue", "jin", "oh", "j", "h"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

//func sexyCount(person string) {
//	for i := 0; i < 10; i++ {
//		fmt.Println(person, "is sexy", i)
//		time.Sleep(time.Second)
//	}
//}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	fmt.Println(person)
	c <- person + " is sexy"
}
