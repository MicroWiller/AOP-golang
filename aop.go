package aop

import "context"

// Before pre-cut point
type Before[T Aspect] interface {
	Before(ctx context.Context, t *T) error
}

// After post cut point
type After[T Aspect] interface {
	After(ctx context.Context, t *T)
}

// AOP generic type for aspect-Oriented Programming.
type AOP[T Aspect] struct {
	befores []Before[T]
	afters  []After[T]
	proxy   T
}

func (aop *AOP[T]) GetBefores() []Before[T] {
	if aop != nil {
		return aop.befores
	}
	return nil
}

func (aop *AOP[T]) GetAfters() []After[T] {
	if aop != nil {
		return aop.afters
	}
	return nil
}

func (aop *AOP[T]) GetProxy() T {
	var t T
	if aop != nil {
		t = aop.proxy
	}
	return t
}

func (aop *AOP[T]) SetProxy(t T) {
	aop.proxy = t
}

// Proxy load the pointcut, and execute the broker.
func (aop *AOP[T]) Proxy(ctx context.Context, opts ...Option[T]) error {

	for _, opt := range opts {
		opt(aop)
	}

	proxy := aop.GetProxy()

	if err := aop.Before(ctx, &proxy); err != nil {
		return err
	}

	if err := proxy.Aspect(ctx); err != nil {
		return err
	}

	aop.After(ctx, &proxy)

	return nil
}

func (aop *AOP[T]) Before(ctx context.Context, t *T) error {
	for _, before := range aop.GetBefores() {
		if err := before.Before(ctx, t); err != nil {
			return err
		}
	}
	return nil
}

func (aop *AOP[T]) After(ctx context.Context, t *T) {
	for _, after := range aop.GetAfters() {
		after.After(ctx, t)
	}
}

type Option[T Aspect] func(aop *AOP[T])

func RegisterBefore[T Aspect](befores ...Before[T]) Option[T] {
	return func(aop *AOP[T]) {
		aop.befores = append(aop.befores, befores...)
	}
}

func RegisterAfter[T Aspect](afters ...After[T]) Option[T] {
	return func(aop *AOP[T]) {
		aop.afters = append(aop.afters, afters...)
	}
}

// Aspect actual business interface.
type Aspect interface {
	Aspect(ctx context.Context) error
}
