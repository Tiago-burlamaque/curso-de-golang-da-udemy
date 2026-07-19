package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4) // adiciona 4 goroutines para esperar

	go func() {
		escrever("Gorountine 1!")
		waitGroup.Done() // sinaliza que a goroutine terminou = -1
	}()

	go func() {
		escrever("Gorountine 2!")
		waitGroup.Done() // sinaliza que a goroutine terminou = -1
	}()

	go func() {
		escrever("Gorountine 3!")
		waitGroup.Done() // sinaliza que a goroutine terminou = -1
	}()

	go func() {
		escrever("Gorountine 4!")
		waitGroup.Done() // sinaliza que a goroutine terminou = -1
	}()

	waitGroup.Wait() // espera todas as goroutines terminarem 
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
