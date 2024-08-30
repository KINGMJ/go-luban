# go-luban
一个现代化、安全且类型安全的 Go 工具库，基于 Go 泛型提供了一系列常用的集合操作助手 (Map, Filter, Reduce, Each, Every, ...)。

## 描述

`go-luban` 是一个实用工具库，旨在简化对数组、切片和 map 等集合的常见操作。该库基于 Go 强大的泛型特性构建，使其既现代又类型安全。`go-luban` 借鉴了 JavaScript 的 Lodash 和 Go-funk 库的设计思想，但在安全性和易用性上有了显著提升。

## 为什么选择 go-luban？
- **基于泛型**：与传统依赖反射的工具库不同，`go-luban` 利用 Go 的泛型提供了类型安全和更高的性能。
- **增强的安全性**：`go-luban` 解决了类似 Go-funk 库中的不足之处，后者使用 `any` 类型常常导致 nil 指针异常。通过确保类型安全，`go-luban` 减少了运行时错误的风险。
- **现代化 API**：`go-luban` 专注于简洁和易用性，为开发者提供了一个干净且直观的 API。诸如 Map、Filter、Reduce、Each 和 Every 等函数能够无缝地处理你的数据结构，提供了你需要的灵活性，同时不牺牲安全性。

## 功能

### 集合操作

- [x] Map
- [x] Filter
- [x] Reduce
- [x] Each
- [x] EachRight
- [x] Every
- [x] Some
- [x] Find
- [x] Chunk
- [x] Compact



### 集合操作（Map）

- [x] MapMap
- [x] FilterMap
- [x] ReduceMap
- [x] EachMap
- [x] EveryMap
- [x] SomeMap
- [x] FindMap

### 字符串处理

### 函数组合

- [ ] 柯里化
- [ ] 组合


### 深拷贝

### 其他常用工具


## 安装
使用 `go get` 安装 `go-luban`：
```
go get github.com/KINGMJ/go-luban
```

## 设计理念：

### 1. 关于`Every`与`Some`对于空切片与空 map 的行为

通常情况下，`Every` 函数对于空切片或空 map 返回 true，因为没有元素可以违反条件。像 js 的 `every` 函数、lodash 的 `every` 都是返回 true。这一点与 `go-funk` 不同，`go-funk` 对于空切片或空 map 返回 false。

`Some` 检查的是**是否有任何元素满足条件**，对于空切片或空 map 来说，因为没有元素可供检查，所以它返回 false。

这种逻辑在许多编程语言（包括 JavaScript、Python 等）中都被广泛采用，确保了一致性和可预测性。

### 2. `Find` 函数

由于golang支持多返回值，所以`Find`函数可以直接返回找到元素的`key`和`value`，无需再使用`FindIndex`函数。

### 3. `Compact` 函数

golang 的 slices 包中有 `Compact` 函数，但它实现的效果是用单个副本替换连续出现的相同元素，它并不等同于去重的功能。实际的使用场景中，对相同元素去重更常见。
所以像`go-funk`库或者`lodash`的`Compact`函数，都是过滤零值的函数。去重操作我们额外提供一个`Uniq`函数来实现。

### 4. 错误处理
`go-luban` 对于错误处理的设计原则是：当出现错误时，应该返回错误，而不是抛出异常。这样可以避免不必要的错误处理逻辑，并使代码更加清晰和可读。比如：`Chunk` 函数，如果传入一个小于1的`size`，应该直接返回错误。在`go1.23.0`版本中，它的`Chunk`函数，如果传入了一个小于1的`size`，会直接抛出异常，导致程序崩溃。
