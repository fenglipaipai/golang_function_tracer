package trace

import (
	trace "github.com/hanzhongzi/golang_function_tracer"
)

func a() {
	defer trace.Trace()()
	b()
}

func b() {
	defer trace.Trace()()
	c()
}

func c() {
	defer trace.Trace()()
	d()
}

func d() {
	defer trace.Trace()()
}

func ExampleTrace() {
	a()
	// Output:
	// g[00001]:    ->github.com/hanzhongzi/golang_function_tracer.a
	// g[00001]:        ->github.com/hanzhongzi/golang_function_tracer.b
	// g[00001]:            ->github.com/hanzhongzi/golang_function_tracer.c
	// g[00001]:                ->github.com/hanzhongzi/golang_function_tracer.d
	// g[00001]:                <-github.com/hanzhongzi/golang_function_tracer.d
	// g[00001]:            <-github.com/hanzhongzi/golang_function_tracer.c
	// g[00001]:        <-github.com/hanzhongzi/golang_function_tracer.b
	// g[00001]:    <-github.com/hanzhongzi/golang_function_tracer.a
}
