// instrument_trace/examples/demo/demo.go

package main

import "github.com/hanzhongzi/golang_function_tracer"

func foo() {
	defer trace.Trace()()

	bar()
}

func bar() {
	defer trace.Trace()()

}

func main() {
	defer trace.Trace()()

	foo()
}
