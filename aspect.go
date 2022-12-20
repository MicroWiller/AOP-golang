package aop

import "context"

type Pointcut interface {
	Pointcut(ctx context.Context) error
}

// Before advice
type Before[T Pointcut] interface {
	Before(ctx context.Context, t *T) error
}

// After advice
type After[T Pointcut] interface {
	After(ctx context.Context, t *T)
}

type Option[T Pointcut] func(aop *AOP[T])

func RegisterBefore[T Pointcut](befores ...Before[T]) Option[T] {
	return func(aop *AOP[T]) {
		aop.befores = append(aop.befores, befores...)
	}
}

func RegisterAfter[T Pointcut](afters ...After[T]) Option[T] {
	return func(aop *AOP[T]) {
		aop.afters = append(aop.afters, afters...)
	}
}
