# go-luban

`go-luban` 是一个基于泛型的现代 Go 工具库，参考了 [go-funk](https://github.com/thoas/go-funk)、[Lodash](https://www.lodashjs.com/) 等现代工具库。

## TODO

### 集合操作

- [x] Map
- [x] Filter
- [x] Reduce
- [x] Each
- [x] EachRight
- [x] Every


- [ ] Find
- [ ] FindIndex

### 集合操作（Map）

- [x] MapMap
- [x] FilterMap
- [x] ReduceMap
- [x] EachMap
- [x] EveryMap


### 字符串处理



### 函数组合

- [ ] 柯里化
- [ ] 组合



### 深拷贝

### 其他常用工具






## 注：

1. 对于空切片的行为，通常情况下，`Every` 函数对于空切片或空 map 返回 true，因为没有元素可以违反条件。像 js 的 `every` 函数、lodash 的 `every` 都是返回 true。这一点与 `go-funk` 不同，`go-funk` 对于空切片或空 map 返回 false。