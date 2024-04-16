package helper

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// SortMapByFunc 自定义map排序
//func SortMapByFunc[K comparable, V any](m map[K]V, sortFunc func(a, b Pair[K, V]) bool) []Pair[K, V] {
//	p := make([]Pair[K, V], len(m))
//	i := 0
//	for k, v := range m {
//		p[i] = Pair[K, V]{Key: k, Value: v}
//		i++
//	}
//	slices.SortFunc(p, sortFunc)
//	return p
//}

func ArrToMap[V comparable](arr []V) map[V]bool {
	var m = make(map[V]bool, len(arr))
	for _, v := range arr {
		m[v] = true
	}
	return m
}

func ArrToMapFunc[M comparable, V any](arr []V, f func(item V) (M, bool)) map[M][]V {
	var ms = map[M][]V{}
	for _, v := range arr {
		m, ok := f(v)
		if ok {
			ms[m] = append(ms[m], v)
		}

	}
	return ms
}

func MapFirst[V any, T comparable](maps map[T]V) V {
	for _, v := range maps {
		return v
	}
	return *new(V)
}

func MapFilter[V comparable, T, R any](maps map[V]T, f func(v V, t T) (R, bool)) []R {
	rs := make([]R, 0)
	for v, t := range maps {
		r, b := f(v, t)
		if b {
			rs = append(rs, r)
		}
	}
	return rs
}

func MapConversion[K comparable, V, A any](vMap map[K]V, f func(K, V) (A, bool)) []A {
	var res []A
	if len(vMap) == 0 {
		return res
	}
	for k, v := range vMap {
		r, b := f(k, v)
		if b {
			res = append(res, r)
		}
	}
	return res
}

func MapGetMaxLen[K comparable, V any](vMap map[K][]V) (maxLen int) {
	for _, v := range vMap {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	return
}

func MapGetKeys[K comparable, V any](vMap map[K]V) []K {
	var keys []K
	for k := range vMap {
		keys = append(keys, k)
	}
	return keys
}

func MapKeysToBoolMap[K comparable](arr []K) map[K]bool {
	var m = make(map[K]bool, len(arr))
	for _, v := range arr {
		m[v] = true
	}
	return m
}
