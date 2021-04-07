package shell

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "os/exec"
    "path"
)

var commands []string = []string{"cd", "exit"}

const reset = "\033[0m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const purple = "\033[35m"
const cyan = "\033[36m"
const white = "\033[37m"

func Repl() {
    for true {
        // print gsh prompt
        prompt()
        // read stdin
        input, err := readCommandLine()
        must(err)
        args := parseCommand(input)
        must(execute(args))
    }
}

func must(e error) {
    if e != nil {
        printErr(e)
    }
}

// gsh prompt
func prompt() {
    wd, err := os.Getwd()
    must(err)
    curDir := path.Base(wd)

    fmt.Printf("%sgsh %s%s %s", yellow, green, curDir, reset)
}

// reads in os.Stdin and passes
func readCommandLine() (string, error) {
    r := bufio.NewReader(os.Stdin)
    input, err := r.ReadString('\n')
    // error throw if string doesnt end in \n
    if err != nil {
        printErr(err)
        // return, since we can't parse input
        return "", err
    }
    return input, nil
}

func parseCommand(raw string) []string {
    // trim \n, and parse command
    command := strings.TrimSuffix(raw, "\n")
    // split by spaces
    return strings.Split(command, " ")
}

func execute(args []string) error {
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

func printErr(e error) {
    fmt.Fprintf(os.Stderr, e.Error()+"\n")
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
        return os.Chdir(args[1])
    case "exit":
        os.Exit(0)
        return nil
    default:
        os.Exit(0)
        return nil
    }

}
