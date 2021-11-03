package main

import (
	"runtime"
	"sync"

	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/di"
)

func main() {
	meowSeederWorker, cleanup, err := di.InitializeMeowSeederWorker()
	if err != nil {
		panic("cannot initialize meow seeder worker: " + err.Error())
	}
	defer cleanup()

	wg := &sync.WaitGroup{}
	for i := 0; i < runtime.NumCPU(); i++ {
		go meowSeederWorker.Run()
		wg.Add(1)
	}
	wg.Wait()
}
