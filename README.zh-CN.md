<p align="center">
<img width="250" src="https://user-images.githubusercontent.com/35942268/177622393-c67a9433-eb2b-4de3-a8e9-262d4db48565.png">
</p>

<p align="center">
    <img src="https://img.shields.io/badge/go-1.13+-blue.svg">
    <img src="https://github.com/wgrape/golib/actions/workflows/build.yml/badge.svg">
    <img src="https://img.shields.io/badge/Document-中文/English-orange.svg">
    <a href="https://godoc.org/github.com/WGrape/golib"><img src="https://godoc.org/github.com/WGrape/golib?status.svg" ></a>
    <img src="https://img.shields.io/badge/License-MIT-green.svg">   
</p>

<div align="center">    
    <p>一个简单易用的Go语言封装库</p>
    <p>文档 ：<a href="/README.zh-CN.md">中文</a> / <a href="/README.md">English</a></p>
</div>

## 背景
在使用Go语言时，你可能会觉得时间计算，文件处理，数组操作等经常需要用较多的时间和代码才能完成，非常的不方便。

为了解决这个问题，可以把所有常见的复杂底层操作和业务操作都封装为一个个简单易用的API，而且最重要的是这些API统一且标准，符合Go底层接口规范。

因此golib项目应运而生。

## 安装
使用go get或mod的方式下载依赖

```bash
go get github.com/WGrape/golib
```

## 使用
下载安装成功后，正常调用```golib API```即可

```go
import github.com/WGrape/golib

golib.API()
```

## 包列表

| package | description               | api                                                   |
|---------|---------------------------|-------------------------------------------------------|
| time    | time包提供了更高效的时间处理操作 | [文档](https://pkg.go.dev/github.com/WGrape/golib/time) |
