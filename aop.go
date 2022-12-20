package aop

import "context"

// AOP generic type for aspect-Oriented Programming.
type AOP[T Pointcut] struct {
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

	if err := proxy.Pointcut(ctx); err != nil {
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
