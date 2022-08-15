package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

/**
each source file can define its own niladic init function to set up whatever state is required.
(Actually each file can have multiple init functions.)
And finally means finally: init is called after all the variable declarations in the package have evaluated their initializers,
and those are evaluated only after all the imported packages have been initialized.
**/
func init() {
	fmt.Println("Start init")

	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	// gopath may be overridden by --gopath flag on command line.
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func main() {
	fmt.Println(home)
	fmt.Println(user)
	fmt.Println(gopath)
}
