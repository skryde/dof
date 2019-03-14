package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

const day = 24  // Day duration in hours

func main() {
	confPath := []string{
		filepath.FromSlash("/home/sebiitta/Gitlab 2 - Seba 0.png"),
		filepath.FromSlash("/tmp/prueba/**"),
	}

	for _, cp := range confPath {
		files, err := filepath.Glob(cp)
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			fInfo, err := os.Stat(f)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Ultimo acceso > 2 dias: %v ----> Path: %s\n", time.Since(time.Unix(fInfo.Sys().(*syscall.Stat_t).Atim.Unix())).Hours() > 2*day, f)
		}
	}
}
