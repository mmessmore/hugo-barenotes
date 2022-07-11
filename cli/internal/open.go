package internal

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

// TODO: this shouldn't be hardcoded
var URL = "http://localhost:1313"

// Fall through config, system defaults, then look at random browsers
var Browsers = []string{
	viper.GetString("browser"),
	"open",
	"xdg-open",
	"x-www-browser",
	"firefox",
	"google-chrome-stable",
	"chromium-browser",
}

func Open() {
	var path = ""
	var err error
	for _, cmd := range Browsers {
		path, err = exec.LookPath(cmd)
		if err == nil {
			break
		}
	}
	if path == "" {
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
