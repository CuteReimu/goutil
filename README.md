# goutil

![](https://img.shields.io/github/go-mod/go-version/CuteReimu/goutil "语言")
[![](https://img.shields.io/github/actions/workflow/status/CuteReimu/goutil/golangci-lint.yml?branch=master)](https://github.com/CuteReimu/goutil/actions/workflows/golangci-lint.yml "代码分析")
[![](https://img.shields.io/github/contributors/CuteReimu/goutil)](https://github.com/CuteReimu/goutil/graphs/contributors "贡献者")
[![](https://img.shields.io/github/license/CuteReimu/goutil)](https://github.com/CuteReimu/goutil/blob/master/LICENSE "许可协议")

参考其它语言联想出来的一些实用工具类

## math

| 函数                           | 说明         |
|------------------------------|------------|
| `math.MinInt32(1,2,3)` 等     | 获取多个数中的最小值 |
| `math.MaxFloat32(1.2,2.3)` 等 | 获取多个数中的最大值 |

## slices

| 函数                               | 说明                                                           |
|----------------------------------|--------------------------------------------------------------|
| `slices.Contains(arr, e)`        | 判断一个 slice 中是否包含某个元素                                         |
| `slices.Equals(arr1, arr2)`      | 当且仅当两个 slice 长度相同且包含的元素完全相同时返回 ture ，否则返回 false              |
| `slices.CopyOf(arr, 5)`          | 复制指定的 slice ，根据 newLength ，如果有必要，则截取或者在后面填默认的零值              |
| `slices.CopyOfRange(arr, 5, 10)` | 将指定 slice 从 from 到 to 下标复制出来，如果 to 比原 slice 的长度更长，则在后面补默认的零值 |
| `slices.ShuffleN(rand, arr, n)`  | 打乱一个 slice ，但后续只会用前n个值，因此做了一些优化                              |
| `slices.Any(arr, f)`             | 只要 slice 中的任一元素满足给定的 func ，就返回 true ，否则返回 false              |
| `slices.All(arr, f)`             | 仅当 slice 中的所有元素都满足给定的 func ， 才返回 true ，否则返回 false            |
| `slices.Sort(arr, lessFunc)`     | 排序，改变原 slice                                                 |
| `slices.Usort(arr, lessFunc)`    | 排序并去重，返回新的 slice                                             |
| `slices.Uniq(arr)`               | 去重，返回新的 slice                                                |

支持了部分Go1.22的range-over-function试验性特性，请配合`GOEXPERIMENT=rangefunc`使用

```go
import "github.com/CuteReimu/goutil/slices"

func main() {
    for i := range slices.Range(1, 5) {
        fmt.Println(i)
    }
}
```

## strings

| 函数                   | 说明          |
|----------------------|-------------|
| `strings.IsEmpty(s)` | 判断字符串是否为空   |
| `strings.IsBlank(s)` | 判断字符串是否全为空白 |

## PriorityQueue

| 函数                                            | 说明                   |
|-----------------------------------------------|----------------------|
| `goutil.NewDefaultPriorityQueue[int](values)` | 用给定的初始值新建优先队列        |
| `goutil.NewPriorityQueue[int](values, cmp)`   | 用给定的初始值和比较函数新建优先队列   |
| `q.Add(value)`                                | 向优先队列里添加一个元素         |
| `q.Peek()`                                    | 获取（但不移除）优先队列的第一个元素   |
| `q.Remove(value)`                             | 从队列中移除指定的元素          |
| `q.Contains(value)`                           | 判断优先队列里是否包含指定的元素     |
| `q.ToSlice(nil)`                              | 返回队列中的所有元素（不一定按大小顺序） |
| `q.Foreach(f)`                                | 遍历队列中的所有元素（不一定按大小顺序） |
| `q.Len()`                                     | 返回优先队列中元素的个数         |
| `q.Clear()`                                   | 清空优先队列               |
| `q.Poll()`                                    | 移除并返回优先队列的第一个元素      |

## BlockingQueue

线程安全的队列，采用链表的方式实现，支持无限容量（不能超过INT_MAX）。

在 `blocking_queue_test.go` 中有与数组方式实现的性能对比。因为数组方式需要频繁拷贝数组（循环数组的实现方式没法做到无限容量），因此链表方式性能略好约三成左右。

| 函数                               | 说明                  |
|----------------------------------|---------------------|
| `goutil.NewBlockingQueue[int]()` | 创建一个线程安全的队列         |
| `q.Len()`                        | 返回队列中的元素个数          |
| `q.Put(e)`                       | 向队尾插入一个元素，不会阻塞      |
| `q.Take()`                       | 从队首获取一个元素，如果队列为空则阻塞 |
| `q.Poll()`                       | 从队首获取一个元素，不阻塞       |
| `q.Peek()`                       | 返回队首的元素，不取出         |
