<div align="center">
<img width="250" src="https://user-images.githubusercontent.com/35942268/177622393-c67a9433-eb2b-4de3-a8e9-262d4db48565.png">
</div>

<div align="center">
    <!-- oscs: https://www.oscs1024.com/cd/1543980900807675904?sign=a3d02348 -->
    <a href="https://www.oscs1024.com/project/oscs/WGrape/golib?ref=badge_small" alt="OSCS Status"><img src="https://www.oscs1024.com/platform/badge/WGrape/golib.svg?size=small"/></a>
    <img src="https://img.shields.io/badge/go-1.13+-blue.svg">
    <img src="https://github.com/wgrape/golib/actions/workflows/build.yml/badge.svg">
    <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/wgrape/golib">
    <img src="https://img.shields.io/badge/Document-中文/English-orange.svg">
    <a href="https://godoc.org/github.com/WGrape/golib"><img src="https://godoc.org/github.com/WGrape/golib?status.svg" ></a>
    <img src="https://img.shields.io/badge/License-MIT-green.svg">   
</div>

<div align="center">    
    <p>一个简单易用的Go语言封装库</p>
    <p>文档 ：<a href="/README.zh-CN.md">中文</a> / <a href="/README.md">English</a></p>
</div>

## 背景
在使用Go语言时，你可能会觉得时间计算，文件处理，数组操作等经常需要用较多的时间和代码才能完成，非常的不方便。

为了解决这个问题，可以把所有常见的复杂底层操作和业务操作都封装为一个个简单易用的API，而且最重要的是这些API统一且标准，符合Go底层接口规范。

因此golib项目应运而生。

## 下载
使用go get或mod这两种方式下载此项目即可，建议下载最新版本。

### go get
```bash
go get github.com/WGrape/golibt@latest
```

### go mod
如果在```go.mod```文件中使用```latest```作为版本，那么在使用```go tidy```等命令下载时，会自动替换为最新的版本号。

```mod
module XXX

go 1.16

require (
    github.com/WGrape/golib latest
)

```


## 使用
下载导入成功后，像如下例子一样正常调用```golib package```即可

> 代码来自于 [matching](https://github.com/WGrape/matching/blob/main/pkg/strategy/strategy.go) 项目

```go
import "github.com/WGrape/golib/permutation"

// getCombinationList get the combination list of properties
func (strategy *UseStrategy) getCombinationList(propertyList []string) []string {
    return permutation.GetCombinationsWithImplode(propertyList, ";")
}
```

## 贡献
由于本项目由各个独立的```package```组成，所以几乎不会有代码阅读成本。如果你有更好的想法，非常欢迎加入一起贡献。关于如何贡献请参考[贡献文档](.github/CONTRIBUTING.md)。

<img src="https://contrib.rocks/image?repo=wgrape/golib">

## 包列表

| package         | description                           | api                                                              |
|-----------------|---------------------------------------|------------------------------------------------------------------|
| safego          | safego包提供了一系列的安全性行为                   | [文档](https://pkg.go.dev/github.com/WGrape/golib/sagego)          |
| permutation     | permutation包提供了排列组合算法相关的操作            | [文档](https://pkg.go.dev/github.com/WGrape/golib/permutation)     |
| http            | http包提供了高效高性能的网络操作                    | [文档](https://pkg.go.dev/github.com/WGrape/golib/http)            |
| redis           | redis包提供了更高级的操作，如解决缓存穿透、缓存击穿、缓存雪崩等问题. | [文档](https://pkg.go.dev/github.com/WGrape/golib/redis)           |
| desensitization | desensitization包提供数据脱敏支持              | [文档](https://pkg.go.dev/github.com/WGrape/golib/desensitization) |
| time            | time包提供了更高效的时间处理操作                    | [文档](https://pkg.go.dev/github.com/WGrape/golib/time)            |
| array           | array包提供了更高效的数组操作                     | [文档](https://pkg.go.dev/github.com/WGrape/golib/array)           |
| binary          | binary包提供了更高效的十进制数字和二进制相关的操作.         | [文档](https://pkg.go.dev/github.com/WGrape/golib/binary)    |
| convert         | convert包提供了更加简单的整型、浮点型、字符串类型之间的转换     | [文档](https://pkg.go.dev/github.com/WGrape/golib/convert)         |
| set             | set包提供了更加简单的集合运算操作                    | [文档](https://pkg.go.dev/github.com/WGrape/golib/set)             |
| rand            | rand包提供了一些随机性的操作                      | [文档](https://pkg.go.dev/github.com/WGrape/golib/rand)            |
| math            | math包提供了更加简单的关于数学计算相关的操作              | [文档](https://pkg.go.dev/github.com/WGrape/golib/math)            |
| system          | system包提供了与Linux交互的接口                 | [文档](https://pkg.go.dev/github.com/WGrape/golib/system)          |
| string          | string包提供了更简单的String处理接口              | [文档](https://pkg.go.dev/github.com/WGrape/golib/string)          |
| slice           | slice包提供了更简单的Slice处理接口                | [文档](https://pkg.go.dev/github.com/WGrape/golib/slice)           |
| frontend        | frontend包提供了与前端交互工作时更简单的接口            | [文档](https://pkg.go.dev/github.com/WGrape/golib/frontend)        |
