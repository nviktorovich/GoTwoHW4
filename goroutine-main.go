package main

import (
	"fmt"
	"sync"
)

//С помощью пула воркеров написать программу, которая запускает 1000 горутин,
//каждая из которых увеличивает число на 1. Дождаться завершения всех горутин и
//убедиться, что при каждом запуске программы итоговое число равно 1000.

var cnt int
var wg sync.WaitGroup

func main() {
	c1 := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			c <- cnt
			cnt++
			fmt.Println(<-c)
		}(c1)
	}
	wg.Wait()

}
