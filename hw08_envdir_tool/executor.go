package main

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	commandStart := cmd[0]
	cmdForRun := exec.Command(commandStart, cmd[1:]...)
	for k, v := range env {
		if v.NeedRemove {
			os.Unsetenv(k)
		} else {
			if _, ok := os.LookupEnv(k); ok {
				os.Unsetenv(k)
			}

			os.Setenv(k, v.Value)
		}
	}

	cmdForRun.Stdin = os.Stdin
	cmdForRun.Stdout = os.Stdout
	cmdForRun.Stderr = os.Stderr

	if err := cmdForRun.Run(); err != nil {
		fmt.Println(err)
		return cmdForRun.ProcessState.ExitCode()
	}

	return 0
}
