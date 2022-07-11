package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Linux has up to 32768
// MacOS up to 99998
// FreeBSD up to 99999
// So 5 bytes
const MAX_PID_SIZE int = 5
const PID_FILE string = ".hugo.pid"

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
	pid, err := GetPid()
	if err != nil {
		fmt.Println("ERROR: Could not read pid file, is your hugo running?")
		fmt.Println(err)
		os.Exit(0)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("WARNING: %d is not running.  Maybe hugo died?\n", pid)
		os.Remove(PID_FILE)
		return
	}

	err = proc.Kill()
	if err != nil {
		fmt.Println("ERROR: failed to stop hugo")
	}
	os.Remove(PID_FILE)
}

func SetPid(pid int) error {
	pidFile, err := os.Create(PID_FILE)
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
	pidBytes := make([]byte, MAX_PID_SIZE)

	pidFile, err := os.Open(PID_FILE)
	if err != nil {
		return 0, err
	}
	_, err = pidFile.Read(pidBytes)
	if err != nil {
		return 0, err
	}

	pid, err = strconv.Atoi(string(pidBytes))
	if err != nil {
		return 0, err
	}

	return pid, nil
}
