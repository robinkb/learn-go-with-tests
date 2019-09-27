// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
	// mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Test() error {
	fmt.Println("Testing...")
	return sh.RunV("go", "test", "./01_hello_world/...")
}
