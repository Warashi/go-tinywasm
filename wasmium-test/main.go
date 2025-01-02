package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"runtime/pprof"
	"strconv"

	"github.com/Warashi/wasmium/runtime"
	typesRuntime "github.com/Warashi/wasmium/types/runtime"
)

type JSONWast struct {
	SourceFilename string     `json:"source_filename"`
	Commands       []Commands `json:"commands"`
}
type Args struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Action struct {
	Type  string `json:"type"`
	Field string `json:"field"`
	Args  []Args `json:"args"`
}
type Expected struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Commands struct {
	Type       string     `json:"type"`
	Line       int        `json:"line"`
	Filename   string     `json:"filename,omitempty"`
	Action     Action     `json:"action,omitempty"`
	Expected   []Expected `json:"expected,omitempty"`
	Text       string     `json:"text,omitempty"`
	ModuleType string     `json:"module_type,omitempty"`
}

type Result = Expected

func action(r *runtime.Runtime, a Action) ([]Result, error) {
	switch a.Type {
	case "invoke":
		return invoke(r, a)
	}
	return nil, fmt.Errorf("unsupported action type %s", a.Type)
}

func invoke(r *runtime.Runtime, a Action) ([]Result, error) {
	field := a.Field
	args := make([]typesRuntime.Value, len(a.Args))

	for i, arg := range a.Args {
		switch arg.Type {
		case "i32":
			a, err := strconv.ParseInt(arg.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			args[i] = typesRuntime.ValueI32(a)
		case "i64":
			a, err := strconv.ParseInt(arg.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			args[i] = typesRuntime.ValueI64(a)
		case "f32":
			a, err := strconv.ParseFloat(arg.Value, 32)
			if err != nil {
				return nil, err
			}
			args[i] = typesRuntime.ValueF32(a)
		case "f64":
			a, err := strconv.ParseFloat(arg.Value, 64)
			if err != nil {
				return nil, err
			}
			args[i] = typesRuntime.ValueF64(a)
		}
	}

	v, err := r.Call(field, args...)
	if err != nil {
		return nil, err
	}

	var result []Result
	for _, v := range v {
		switch v := v.(type) {
		case typesRuntime.ValueI32:
			result = append(result, Result{Type: "i32", Value: strconv.FormatInt(int64(v), 10)})
		case typesRuntime.ValueI64:
			result = append(result, Result{Type: "i64", Value: strconv.FormatInt(int64(v), 10)})
		case typesRuntime.ValueF32:
			result = append(result, Result{Type: "f32", Value: strconv.FormatFloat(float64(v), 'f', -1, 32)})
		case typesRuntime.ValueF64:
			result = append(result, Result{Type: "f64", Value: strconv.FormatFloat(float64(v), 'f', -1, 64)})
		}
	}

	return result, nil
}

func _main() (exitCode int) {
	var prof bool
	flag.BoolVar(&prof, "prof", false, "record cpuprofile with profile.out")
	flag.Parse()

	baseDir := filepath.Dir(flag.Arg(0))

	if prof {
		f, err := os.Create("profile.out")
		if err != nil {
			slog.Error("failed to create profile.out", slog.Any("error", err))
			return 1
		}
		defer func() {
			if err := f.Close(); err != nil {
				slog.Error("failed to close profile.out", slog.Any("error", err))
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

	var wast JSONWast
	if err := json.NewDecoder(f).Decode(&wast); err != nil {
		slog.Error("failed to decode json", slog.Any("error", err))
		return 1
	}

	var r *runtime.Runtime

	var current Commands
	defer func() {
		if r := recover(); r != nil {
			slog.Error("panic", slog.Any("command", current), slog.Any("recover", r))
			exitCode = 1
		}
	}()
	for _, cmd := range wast.Commands {
		current = cmd
		switch cmd.Type {
		case "module":
			f, err := os.Open(filepath.Join(baseDir, cmd.Filename))
			if err != nil {
				slog.Error("failed to open file", slog.Any("error", err))
				return 1
			}
			r, err = runtime.New(f)
			if err != nil {
				slog.Error("failed to create runtime", slog.Any("error", err))
				return 1
			}
		case "assert_return":
			got, err := action(r, cmd.Action)
			if err != nil {
				slog.Warn("failed to execute action", slog.Any("command", cmd), slog.Any("error", err))
				continue
			}
			if !reflect.DeepEqual(cmd.Expected, got) {
				slog.Warn("assertion failed", slog.Any("command", cmd), slog.Any("expected", cmd.Expected), slog.Any("got", got))
			}
		}

	}

	return 0
}

func main() {
	os.Exit(_main())
}
