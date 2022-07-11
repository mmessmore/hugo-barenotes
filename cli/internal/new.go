package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/spf13/viper"
)

func CD() {
	dir := viper.GetString("root")
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("ERROR: Could not change directory to %s\n", dir)
		fmt.Println(err)
		os.Exit(1)
	}
}

func Create(title string) {
	filename := fmt.Sprintf("notes/%s.md",
		strings.ReplaceAll(strings.ToLower(title), " ", "-"))

	hugo, err := exec.LookPath("hugo")
	if err != nil {
		fmt.Println("ERROR: Could not find hugo executable")
		os.Exit(1)
	}

	cmd := exec.Command(hugo, "new", filename)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("ERROR: Failed to run hugo")
		fmt.Println(out)
	}

	editor, err := GetEditor()
	if err != nil {
		fmt.Println("ERROR: Could not find an editor to execute")
		os.Exit(1)
	}
	filename = fmt.Sprintf("content/%s", filename)
	env := os.Environ()

	err = syscall.Exec(editor, []string{filename}, env)
	if err != nil {
		fmt.Printf("ERROR: failed to open editor on %s\n", filename)
		fmt.Println(err)
		os.Exit(1)
	}
}
