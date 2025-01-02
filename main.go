package main

import (
	"context"
	_ "embed"
	"log/slog"
	"os/signal"

	"flag"
	"os"
	"runtime/pprof"

	"github.com/Warashi/wasmium/runtime"
	"github.com/Warashi/wasmium/wasip1"
)

func main() {
	os.Exit(_main())
}

func _main() (exitCode int) {
	var prof bool
	flag.BoolVar(&prof, "prof", false, "record cpuprofile with profile.out")
	flag.Parse()

	if prof {
		f, err := os.Create("profile.out")
		if err != nil {
			slog.Error("failed to create profile.out", slog.Any("error", err))
			return 1
		}
		defer func() {
			if err := f.Close(); err != nil {
				slog.Error("failed to close profile.out", slog.Any("error", err))
				exitCode = 1
			}
		}()
		if err := pprof.StartCPUProfile(f); err != nil {
			slog.Error("failed to start cpuprofile", slog.Any("error", err))
			return 1
		}
		defer pprof.StopCPUProfile()
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		slog.Error("failed to open file", slog.Any("error", err))
		return 1
	}
	defer f.Close()

	r, err := runtime.New(f)
	if err != nil {
		slog.Error("failed to create runtime", slog.Any("error", err))
		return 1
	}

	wasip1.NewWasiPreview1().Register(r)

	sigCh, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	done := make(chan int)
	go func() {
		defer close(done)
		if _, err := r.Call("_start"); err != nil {
			slog.Error("failed to call _start", slog.Any("error", err))
			done <- 1
		}
	}()

	select {
	case <-sigCh.Done():
		slog.Info("interrupted")
		return 1
	case c := <-done:
		return c
	}
}
