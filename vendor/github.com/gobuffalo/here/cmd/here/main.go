package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gobuffalo/here"
)

func main() {
	defer func() {
		c := exec.Command("go", "mod", "tidy")
		c.Run()
	}()
	pwd, _ := os.Getwd()

	args := os.Args[1:]
	if len(args) == 0 {
		args = append(args, pwd)
	}

	h := here.New()

	fn := h.Dir
	switch args[0] {
	case "pkg":
		fn = h.Package
		args = args[1:]
		if len(args) == 0 {
			log.Fatalf("you must pass at least one package name")
		}
	}

	for _, a := range args {
		i, err := fn(a)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stdout, i.String())
	}
}
