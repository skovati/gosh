package commands

import (
    "fmt"
    "os/exec"
    "os"
)

var commands []string = []string{"cd", "exit"}

func Execute(args []string) error {
    // check if arg[0] is a native gsh command
    for _, v := range commands {
        if v == args[0] {
            // if so, run and return native error
            return execNative(args)
        }
    }

    // otherwise, run as system command
    return execSystem(args)
}

func execSystem(args []string) error {
    // make new command with optional args
    cmd := exec.Command(args[0], args[1:]...)
    // set stderr and stdout
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    // exec and return error
    _, err := exec.LookPath(cmd.Path)
    if err != nil {
        return fmt.Errorf("Error, command not found in PATH")
    }
    return cmd.Run()
}

func execNative(args []string) error {
    switch args[0] {
    case "cd":
        return cd(args)
    case "exit":
        exit()
    }
    return fmt.Errorf("Error, native function could not be run")
}
