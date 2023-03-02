package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()

	envMap, err := ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	exitCodeRun := RunCmd(args[1:], envMap)

	os.Exit(exitCodeRun)
}
