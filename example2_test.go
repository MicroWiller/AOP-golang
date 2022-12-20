package aop

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestAOP(t *testing.T) {

	busAop := NewBus("131", "13", 13)

	beforePoint := RegisterBefore[BusProxy]
	afterPoint := RegisterAfter[BusProxy]

	p := Police{}
	pBeforeOpt := beforePoint(p)
	pAfterOpt := afterPoint(p)

	ctx := context.Background()
	err := busAop.Proxy(ctx, pBeforeOpt, pAfterOpt)
	if err != nil {
		panic(err)
	}

}

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

// BusProxy proxy bus
type BusProxy struct {
	bus *Bus
}

func (ap BusProxy) Pointcut(ctx context.Context) error {
	return ap.bus.Drive(ctx)
}

type Bus struct {
	Name       string
	Route      string
	Passengers int64
}

func (b Bus) Drive(ctx context.Context) error {
	fmt.Printf("this %s bus is driving %s , has %d Passengers\n",
		b.Name, b.Route, b.Passengers)
	return nil
}

type Police struct {
}

func (p Police) Before(ctx context.Context, bp *BusProxy) error {
	fmt.Println("police before check")
	if bp.bus.Passengers <= 20 {
		fmt.Println("police check Passengers is OK!")
	}
	if bp.bus.Passengers > 30 {
		return errors.New("police check overload")
	}
	return nil
}

func (p Police) After(ctx context.Context, bp *BusProxy) {
	fmt.Println("police after doing")
	fmt.Printf("police say %s check done where is %s\n", bp.bus.Name, bp.bus.Route)
}
