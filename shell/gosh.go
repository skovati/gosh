package shell

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "path"

    "github.com/skovati/gosh/commands"
)

const reset = "\033[0m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const purple = "\033[35m"
const cyan = "\033[36m"
const white = "\033[37m"

// main repl loop
func Repl() {
    for true {
        // print gsh prompt
        prompt()
        // read stdin
        input, err := readCommandLine()
        must(err)
        args := parseCommand(input)
        must(commands.Execute(args))
    }
}

func must(e error) {
    if e != nil {
        printErr(e)
    }
}

// gsh prompt
func prompt() {
    // get working dir
    wd, err := os.Getwd()
    must(err)
    // get just last dir
    curDir := path.Base(wd)
    // print with colors, no new line
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

func printErr(e error) {
    fmt.Fprintf(os.Stderr, e.Error()+"\n")
}
