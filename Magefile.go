// +build mage

package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Format all code
func Fmt() error {
	return sh.RunV("gofmt", "-s", "-w", ".")
}

type Test mg.Namespace

// Run all tasks defined under "test"
func (Test) All() error {
	// Ask the runtime for the name of this function.
	// I just wanted to see what it looks like. Never do this in code that other people need to look at.
	pc, _, _, _ := runtime.Caller(0)
	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next()
	parts := strings.Split(frame.Function, ".")
	this := parts[len(parts)-1]

	// Create a reflect.Type from the Test type.
	typ := reflect.TypeOf(Test{})

	// Iterate over all methods defined on the Test type and execute them.
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		// If the retrieved method is the one currently executing, skip it.
		// Lest we recurse infinitely.
		if method.Name == this {
			continue
		}

		// Slice of arguments that will be passed to the method.
		args := []reflect.Value{
			reflect.ValueOf(Test{}), // The first argument is the method's receiver.
		}

		// Execute the method and catch any number of return values.
		ret := method.Func.Call(args)

		for _, r := range ret {
			// We expect any of the methods defined on type Test to either return nil or an error.
			switch v := r.Interface().(type) {
			case nil:
			case error:
				return v
			default:
				return fmt.Errorf("unexpected return value: %q", r)
			}
		}
	}

	return nil
}

// Test the "hello" package
func (Test) Hello() error {
	return run("hello")
}

// Test the "integers" package
func (Test) Integers() error {
	return run("integers")
}

// Test the "iteration" package
func (Test) Iteration() error {
	return run("iteration")
}

// Test the "arrays_and_slices" package
func (Test) ArraysAndSlices() error {
	return run("arrays_and_slices")
}

// Test the "structs_methods_and_interfaces" package
func (Test) StructsMethodsAndInterfaces() error {
	return run("structs_methods_and_interfaces")
}

// Test the "pointers_and_errors" package
func (Test) PointersAndErrors() error {
	return run("pointers_and_errors")
}

// Test the "maps" package
func (Test) Maps() error {
	return run("maps")
}

// task defines the type for a function that runs a task in the given directory.
type task func(dir string) error

// run runs a collection of tasks; usually all of them.
func run(dir string) error {
	fmt.Printf("Testing %s...\n", dir)

	tasks := []task{
		format,
		vet,
		test,
	}

	for _, t := range tasks {
		if err := t(dir); err != nil {
			return err
		}
	}

	return nil
}

// format runs 'gofmt' in the given directory, and returns an error when one or more files need formatting.
func format(dir string) error {
	output, err := sh.Output("gofmt", "-l", "-s", dir)
	if err != nil {
		return err
	}

	if output != "" {
		return fmt.Errorf("%s\n%s\n\n%s",
			"The following files need formatting:",
			output,
			"You can fix this problem by running 'mage fmt'.")
	}

	return nil
}

// vet runs 'go vet' in the given directory.
func vet(dir string) error {
	return sh.RunV("go", "vet", "./"+dir)
}

// test runs 'go test' in the given directory.
func test(dir string) error {
	return sh.RunV("go", "test",
		"-count=1", // Disable test caching
		"-cover",   // Enable coverage reporting
		"./"+dir)
}
