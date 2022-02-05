package main

import (
	"fmt"
	"os"

	"github.com/benji2512/multiple-package/cmd"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to get working directory: %s",
			err.Error())
	}

	res, err := cmd.ExcuteLs(wd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to list files in %s, error: %s",
			wd, err.Error())
	}

	fmt.Printf("%s\n", res)
}
