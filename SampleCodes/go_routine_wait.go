package main

import (
	"log/slog"
	"sync"
)

func goRoutineWait() {
	slog.Info("main routine start")
	var wg sync.WaitGroup

	wg.Go(func() {
		printNumbers()
	})

	wg.Wait()

	slog.Info("main routine end")
}

func printNumbers() {
	for i := 0; i < 100; i++ {
		slog.Info("go routine", "index", i)
	}
}
