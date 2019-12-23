package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

type paths []path
type path struct {
	Path         string  `json:"path"`
	DaysToDelete float64 `json:"days_to_delete"`
	SafeDelete   bool    `json:"safe_delete"`
}

const day = 24 // Day duration in hours
const configFile = "dof.json"

var conf paths

func main() {
	for _, cp := range conf {
		files, err := filepath.Glob(filepath.FromSlash(cp.Path))
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			fInfo, err := os.Stat(f)
			if err != nil {
				panic(err)
			}

			lastAccess := time.Unix(fInfo.Sys().(*syscall.Stat_t).Atim.Unix())
			fmt.Printf("Ultimo acceso > %.0f dias: %v ----> Path: %s\n",
				cp.DaysToDelete, time.Since(lastAccess).Hours() > cp.DaysToDelete*day, f)

			if time.Since(lastAccess).Hours() > cp.DaysToDelete*day {
				err = os.Remove(f)
				if err != nil {
					fmt.Printf("No se pudo eliminar el archivo %s", f)
				}
			}
		}
	}
}
