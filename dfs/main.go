package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var cmdFlags struct {
	help    bool
	version bool
	run     bool
}

func main() {
	flag.BoolVar(&cmdFlags.help, "help", false, "help message")
	flag.BoolVar(&cmdFlags.version, "version", false, "app version")
	flag.BoolVar(&cmdFlags.run, "run", false, "run app")

	flag.Parse()
	if len(os.Args[1:]) > 2 {
		flag.Usage()
		os.Exit(-1)
	}
	if cmdFlags.help {
		flag.Usage()
		os.Exit(-2)
	}
	if cmdFlags.version {
		fmt.Println("App Version: 0.0.1")
		os.Exit(-3)
	}
	if !cmdFlags.run {
		flag.Usage()
		os.Exit(-4)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("route", slog.String("uri", r.URL.Path))
		fmt.Fprintf(w, "hello world")
	})
	slog.Info("App Run", slog.String("addr", ":8022"))
	if err := http.ListenAndServe(":8022", nil); err != nil {
		panic(err)
	}
}

