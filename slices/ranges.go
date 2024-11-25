package slices

// Range 生成一个左开右闭区间的整数序列，包含 start ，不含 end
//
// 使用方法：
//
//	for i := range Range(1, 5) {
//	    fmt.Println(i)
//	}
func Range[E integer](start, end E) func(func(E) bool) {
	return Progression(start, end, 1)
}

// Progression 生成一个左开右闭区间的整数序列，包含 start ，不含 end ，步长为 step
//
// 使用方法：
//
//	for i := range Progression(1, 5, 2) {
//	    fmt.Println(i)
//	}
func Progression[E integer](start, end, step E) func(func(E) bool) {
	return func(f func(E) bool) {
		for i := start; i < end; i += step {
			if !f(i) {
				break
			}
		}
	}
}
