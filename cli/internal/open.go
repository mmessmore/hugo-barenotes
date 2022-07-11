package internal

import (
	"fmt"
	"os"
)

// TODO: this shouldn't be hardcoded
var URL = "http://localhost:1313"

func Open() {

	path, err := GetBrowser()
	if err != nil {
		fmt.Println("ERROR: could not find command to open browser")
		os.Exit(0)
	}

	args := []string{path, URL}

	proc, err := os.StartProcess(path, args, &os.ProcAttr{})
	if err != nil {
		fmt.Println("ERROR: could not open browser")
		fmt.Println(err)
		os.Exit(1)
	}
	proc.Release()
}
