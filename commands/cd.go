package commands

import (
    "os"
)

func cd(args []string) error {
    return os.Chdir(args[1])
}
