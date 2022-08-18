# Golang项目函数调用链跟踪工具。
> 自动注入跟踪代码并输出有层次感的函数调用链跟踪命令行工具。

# 用法:

1.下载项目自己编译使用

```bash

git clone https://github.com/hanzhongzi/golang_function_tracer.git

cd golang_function_tracer ; go build github.com/hanzhongzi/golang_function_tracer

chmod a+x golang_function_tracer

```

 比如你有一个函数特别多的文件，函数调用复杂，日志输出不完善，不能快速定位问题。
` cat explame/demo/demo.go `

```golang

package main

func foo() {

        bar()
}

func bar() {

}

func main() {
        
        foo()
}

```


 使用本项目编译出的 golang_function_tracer 二进制命令行工具注入函数调用逻辑展示

```
./golang_function_tracer -w explame/demo/demo.go 
```

 当再次运行此函数时会打印出详细函数调用逻辑并且会有层次感的输出，配合grep 达到快速理解文件内函数调用逻辑。
```golang
go run  explame/demo/demo.go                                                                           

g[00001]:    ->main.main
g[00001]:        ->main.foo
g[00001]:            ->main.bar
g[00001]:            <-main.bar
g[00001]:        <-main.foo
g[00001]:    <-main.main

```

