# goutil

![](https://img.shields.io/github/languages/top/CuteReimu/goutil "语言")
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

## strings

| 函数                   | 说明          |
|----------------------|-------------|
| `strings.IsEmpty(s)` | 判断字符串是否为空   |
| `strings.IsBlank(s)` | 判断字符串是否全为空白 |
