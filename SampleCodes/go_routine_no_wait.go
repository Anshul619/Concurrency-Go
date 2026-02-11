package main

import (
	"log/slog"
)

func goRoutineNoWait() {
	slog.Info("main routine start")
	go printNumbers1()
	slog.Info("main routine end")
}

func printNumbers1() {
	for i := 0; i < 100; i++ {
		slog.Info("go routine %v", i)
	}
}
