# golang plugin 机制
Go 1.8为我们提供了一个创建共享库的新工具，称为Plugins！让我们来创建和使用一个插件。 目前的插件只能在Linux和Darwin上工作。
## 应用场景

## 步骤
1. 声明插件文件
2. 构建插件
3. 在其他go程序动态加载调用插件

## 2.构建插件
```shell
#  构建当前目录下的多有go文件
go build -buildmode=plugin

# 指定构建go文件 -o 指定文件名
go build -buildmode=plugin -o aplugin.so aplugin.go
```

## 3.在其他go程序调用插件

1. plugin.Open 打开插件
2. Lookup() 查找插件的方法


## 注意

使用 go run 执行demo程序

1. 在idea中运行不了,原因是idea是先build当前的程序到另外一个目录, 但是插件so文件并没有拷贝过去


