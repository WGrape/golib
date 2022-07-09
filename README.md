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
    <p>A simple and easy-to-use library in go</p>
    <p>Document ：<a href="/README.zh-CN.md">中文</a> / <a href="/README.md">English</a></p>
</div>

## Background

When using the Go language, you may feel that time calculation, file processing, array operations, etc. often require a lot of time and code to complete, which is very inconvenient.

In order to solve this problem, all common complex underlying operations and business operations can be encapsulated into simple and easy-to-use APIs, and the most important thing is that these APIs are unified and standard, and conform to the Go interface specification.

So the golib project came into being.

## Download
Using ```go get``` or ```go mod``` the two ways to download this project. It is recommended to download the latest version.

### go get
```bash
go get github.com/WGrape/golibt@latest
```

### go mod
If you put the word latest in place of the tag in the ```go.mod``` file it will get changed to the latest tag the modules.

```mod
module XXX

go 1.16

require (
    github.com/WGrape/golib latest
)

```

## Usage
After the download and import are successful, you can call the ```golib package``` as follows.

> The code is from [matching](https://github.com/WGrape/matching/blob/main/pkg/strategy/strategy.go) project.

```go
import "github.com/WGrape/golib/permutation"

// getCombinationList get the combination list of properties
func (strategy *UseStrategy) getCombinationList(propertyList []string) []string {
    return permutation.GetCombinationsWithImplode(propertyList, ";")
}
```

## Package list

| package | description                                                                             | api                                                     |
|---------|-----------------------------------------------------------------------------------------|---------------------------------------------------------|
| time    | Package time provides functionality for measuring and displaying time more efficiently. | [doc](https://pkg.go.dev/github.com/WGrape/golib/time)  |
 | array | Package array provides functionality for measuring and displaying array more efficiently. | [doc](https://pkg.go.dev/github.com/WGrape/golib/array) |
| permutation | Package permutation provides many algorithms about permutation and combination. | [doc](https://pkg.go.dev/github.com/WGrape/golib/permutation) |
