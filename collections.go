package luban

// MapFn 是一个通用的映射函数类型，它接受一个类型为 E 的元素，并返回一个类型为 R 的元素
type mapFn[E any, R any] func(E) R

// 支持array、slice
func Map[S ~[]E, E any, R any](s S, f mapFn[E, R]) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// MapMapFn 是一个通用的映射函数类型，它接受 map 的键和值，并返回一个新的值类型 R
type mapMapFn[K comparable, V any, R any] func(K, V) R

func MapMap[K comparable, V any, R any](m map[K]V, f mapMapFn[K, V, R]) map[K]R {
	if m == nil {
		return map[K]R{}
	}
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = f(k, v)
	}
	return result
}

// filterFn 是一个通用的过滤函数类型，它接受一个类型为 E 的元素，并返回一个布尔值
type filterFn[E any] func(E) bool

func Filter[S ~[]E, E any](s S, f filterFn[E]) S {
	var result S
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// filterMapFn 是一个通用的过滤函数类型，它接受 map 的键和值，并返回一个布尔值
type filterMapFn[K comparable, V any] func(K, V) bool

func FilterMap[M ~map[K]V, K comparable, V any](m M, f filterMapFn[K, V]) M {
	result := make(M)
	for k, v := range m {
		if f(k, v) {
			result[k] = v
		}
	}
	return result
}

// reduceFn 是一个通用的归约函数类型，它接受一个累加器类型为 R 和一个当前元素类型为 E，并返回一个新的累加器类型为 R
type reduceFn[E any, R any] func(R, E) R

func Reduce[S ~[]E, E any, R any](s S, f reduceFn[E, R], initial R) R {
	result := initial
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

type reduceMapFn[K comparable, V any, R any] func(R, K, V) R

func ReduceMap[M ~map[K]V, K comparable, V any, R any](m M, f reduceMapFn[K, V, R], initial R) R {
	result := initial
	for k, v := range m {
		result = f(result, k, v)
	}
	return result
}

// eachFn 是一个接受一个元素并对其执行操作的函数类型
type eachFn[E any] func(E)

func Each[S ~[]E, E any](s S, f eachFn[E]) {
	for _, v := range s {
		f(v)
	}
}

type eachMapFn[K comparable, V any] func(K, V)

func EachMap[M ~map[K]V, K comparable, V any](m M, f eachMapFn[K, V]) {
	for k, v := range m {
		f(k, v)
	}
}
