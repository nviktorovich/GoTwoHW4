package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Написать программу, которая при получении в канал сигнала SIGTERM
//останавливается не позднее, чем за одну секунду (установить таймаут).

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		time.Sleep(time.Second * 3)
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
