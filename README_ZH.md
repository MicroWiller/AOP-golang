# AOP: 基于泛型实现的AOP库，简单轻量。

🇬🇧 [English](./README.md) | 🇨🇳 中文

## Overview

_AOP_（面向切面编程）是一种编程设计思想，是OOP（面向对象程序设计）的延续，是软件开发中的一个热点。利用AOP可以对业务逻辑的各个部分进行隔离，从而使得业务逻辑各部分之间的耦合度降低，提高程序的可重用性，同时提升开发的效率。

## Installation

`go get github.com/MicroWiller/AOP-golang`

## Usage

1) 首先，定义一个结构类型，继承 Aspect 接口
```go
package aop

import "context"

// BusProxy proxy bus
type BusProxy struct {
	bus *Bus
}

func (ap BusProxy) Aspect(ctx context.Context) error {
	return nil
}

type Bus struct {
	Name       string
	Route      string
	Passengers int64
}
```

2) 实例化泛型结构AOP
```go
// NewBus instantiate generic AOP for BusProxy.
func NewBus(name, route string, p int64) AOP[BusProxy] {
busProxyAop := AOP[BusProxy]{}
bus := Bus{
Name:       name,
Route:      route,
Passengers: p,
}
proxy := BusProxy{bus: &bus}
busProxyAop.SetProxy(proxy)
return busProxyAop
}
```

3) 实例化前置/后置切入点
```go
// Pre-cut point Option
beforePoint := RegisterBefore[BusProxy]
// Post cut point Option
afterPoint := RegisterAfter[BusProxy]
```

4) 切入点结构继承泛型接口 Before[T Aspect] / After[T Aspect]
```go
type Police struct {
}

func (p Police) Before(ctx context.Context, bp *BusProxy) error {
    return nil
}

func (p Police) After(ctx context.Context, bp *BusProxy) {
}
```

5) 生成Option加载点
```go
p := Police{}
pBeforeOpt := beforePoint(p)
pAfterOpt := afterPoint(p)
```

6) 执行AOP代理方法
```go
busAop.Proxy(ctx, pBeforeOpt, pAfterOpt)
```

## 最后
更多使用示例请查看`test文件`，欢迎提`issue`。