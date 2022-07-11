package internal

import (
	"fmt"
	"os/exec"
	"os"
)

/*
These functions all assume that the current working directory is the hugo
root directory.

Call CD() prior.
*/

func Start() {
	hugo, err := exec.LookPath("hugo")
	if err != nil {
		fmt.Println("ERROR: hugo command not found in PATH")
		os.Exit(1)
	}

	procAttr := os.ProcAttr{}

	args := []string{hugo, "serve", "--logFile", "/dev/null"}
	proc, err := os.StartProcess("hugo", args, &procAttr)
	if err != nil {
		fmt.Println("ERROR: could not start hugo server")
		// this is an os.PathError aka io.fs.PathError
		fmt.Println(err.Error())
		// I'd like to be returning the exit code of hugo, but don't know how
		os.Exit(1)
	}
	SetPid(proc.Pid)
	proc.Release()
}

func Stop() {
	// TODO: make this get the process by pid and kill it
	// Getting the process by pid will always succeed, need to check it's real
}

func SetPid(pid int) error {
	pidFile, err := os.Create(".hugo.pid")
	if err != nil {
		return err
	}
	defer pidFile.Close()
	_, err = fmt.Fprintln(pidFile, pid)
	if err != nil {
		return err
	}
	return nil
}

func GetPid() (int, error) {
	var pid int

	//TODO: make this get the pid from the file

	return pid, nil
}
