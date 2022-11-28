参考资料：
 - https://github.com/zxh0/jvmgo-book
 - 《自己动手写Java虚拟机》
 - [Java语言规范、Java虚拟机规范](https://docs.oracle.com/javase/specs/index.html)

```
preferences -> go -> GOPATH -> Project GOPATH
/Users/nibnait/project-go/jvm-go
``` 

## ch01 命令行工具
```
> cd ~/project-go/jvm-go/src/main
> go build .
> ./main -version
version 0.0.1
> ./main -help
Usage: ./main [-option] class [args...]
> ./main -cp foo/bar MyApp arg1 arg2
classpath: foo/bar, class: MyApp, args: [arg1 arg2] 

```

## ch02 搜索 class 文件
```
> cd ~/project-go/jvm-go/src/main
> go build .
> ./main -XJre "/Library/Java/JavaVirtualMachines/jdk1.8.0_231.jdk/Contents/Home/jre" java.lang.Object
```
