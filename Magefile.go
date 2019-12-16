// +build mage

package main

import (
	"errors"
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
	me := getFunctionName()

	methods, err := getMethods(Test{}, []string{me})
	if err != nil {
		return err
	}

	errChans := make([]chan error, len(methods))
	for i := range errChans {
		errChans[i] = make(chan error, 0)
	}

	for i := range methods {
		go func(method call, err chan error) {
			for _, r := range method() {
				switch v := r.(type) {
				case nil:
				case error:
					err <- v
				default:
					err <- fmt.Errorf("unexpected return value: %#v", r)
				}
			}

			close(err)

		}(methods[i], errChans[i])
	}

	errs := 0
	for _, errChan := range errChans {
		for range errChan {
			errs++
		}
	}

	if errs != 0 {
		return errors.New("some tests failed")
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

// Test the "arrays-and-slices" package
func (Test) ArraysAndSlices() error {
	return run("arrays-and-slices")
}

// Test the "structs-methods-and-interfaces" package
func (Test) StructsMethodsAndInterfaces() error {
	return run("structs-methods-and-interfaces")
}

// Test the "pointers-and-errors" package
func (Test) PointersAndErrors() error {
	return run("pointers-and-errors")
}

// Test the "maps" package
func (Test) Maps() error {
	return run("maps")
}

// Test the "dependency-injection" package
func (Test) DependencyInjection() error {
	return run("dependency-injection")
}

// Test the "mocking" package
func (Test) Mocking() error {
	return run("mocking")
}

// Test the "concurrency" package
func (Test) Concurrency() error {
	return run("concurrency")
}

// Test the "select" package
func (Test) Select() error {
	return run("select")
}

// Test the "reflection" package
func (Test) Reflection() error {
	return run("reflection")
}

// Test the "sync" package
func (Test) Sync() error {
	return run("sync")
}

// Test the "context" package
func (Test) Context() error {
	return run("context")
}

// Test the "property-based-tests" package
func (Test) PropertyBasedTests() error {
	return run("property-based-tests")
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

// A bunch of functions that wrap the complexities of the reflect package.

// getFunctionName returns the name of the function calling this function.
func getFunctionName() string {
	// Return the pointer to the previous call in the call stack,
	// so the function that called getFunctionName().
	p, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("could not get function name: " +
			"failed to retrieve pointer to calling function")
	}

	// Return information about the function behind the pointer.
	frames := runtime.CallersFrames([]uintptr{p})
	frame, _ := frames.Next()

	// frame.Function looks like "main.func1.func2",
	// so split on the dots and return the last part.
	parts := strings.Split(frame.Function, ".")
	return parts[len(parts)-1]
}

type call func(args ...interface{}) []interface{}

// getMethods returns all methods for the given struct, except the ones in the list of exceptions.
func getMethods(s interface{}, except []string) ([]call, error) {
	structType := reflect.TypeOf(s)
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("give me a struct, dimwit")
	}

	structValue := reflect.ValueOf(s)

	calls := make([]call, 0)

	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)

		if contains(except, method.Name) {
			continue
		}

		calls = append(calls,
			func(args ...interface{}) []interface{} {
				in := make([]reflect.Value, len(args)+1)

				in[0] = structValue
				for i := range args {
					in[i+1] = reflect.ValueOf(args[i])
				}

				out := method.Func.Call(in)

				ret := make([]interface{}, len(out))
				for i := range ret {
					ret[i] = out[i].Interface()
				}

				return ret
			})
	}

	return calls, nil
}

func contains(haystack []string, needle string) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}
