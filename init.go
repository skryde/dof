package main

import (
	"fmt"
	"github.com/skryde/jsconf"
	"os"
)

func init() {
	if res := jsconf.Exist(configFile); res == jsconf.IsFile {
		err := jsconf.LoadFromFile(configFile, &conf)
		if err != nil {
			panic(err)
		}

	} else if res == jsconf.NotExist {
		conf = append(conf, path{
			Path:         "/path/to/example",
			DaysToDelete: 20,
			SafeDelete:   false,
		},
			path{
				Path:         "/tmp/prueba/**",
				DaysToDelete: 15,
				SafeDelete:   false,
			})

		err := jsconf.SaveToFile(configFile, conf)
		if err != nil {
			panic(err)
		}

		fmt.Println("dof.json file created... You should add your own files/folder to it and runme again")
		os.Exit(0)
	}
}
