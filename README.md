# AOP-golang: A AOP library based on generic implementation.

🇬🇧 English | 🇨🇳 [中文](./README_ZH.md)

## Overview

_AOP_ (Aspect Oriented Programming) is a kind of programming design idea, is the continuation of OOP (Object Oriented Programming), is a hot spot in software development. AOP can be used to isolate each part of business logic, so as to reduce the coupling degree between the parts of business logic, improve the reusability of the program, and improve the efficiency of development.

## Installation

`go get github.com/MicroWiller/AOP-golang`

## Usage

1) First, define a structure type that implement the Aspect interface.
```go
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

2) Instantiate the generic architecture AOP.
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

3) Instantiate a pre/post pointcut.
```go
// Pre-cut point Option
beforePoint := RegisterBefore[BusProxy]
// Post cut point Option
afterPoint := RegisterAfter[BusProxy]
```

4) Pointcut structures implement generic interfaces Before[T Aspect] / After[T Aspect].
```go
type Police struct {
}

func (p Police) Before(ctx context.Context, bp *BusProxy) error {
    return nil
}

func (p Police) After(ctx context.Context, bp *BusProxy) {
}
```

5) Generate Option load points.
```go
p := Police{}
pBeforeOpt := beforePoint(p)
pAfterOpt := afterPoint(p)
```

6) Execute AOP proxy methods.
```go
busAop.Proxy(ctx, pBeforeOpt, pAfterOpt)
```

## Finally
更多使用示例请查看`test文件`，欢迎提`issue`。