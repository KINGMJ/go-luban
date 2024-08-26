package luban

// MapFn 是一个通用的映射函数类型，它接受一个类型为 E 的元素，并返回一个类型为 R 的元素
type mapFn[E any, R any] func(E) R

// 支持array、slice
func MapSlice[S ~[]E, E any, R any](s S, f mapFn[E, R]) []R {
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
