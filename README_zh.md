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


## 安装
使用 `go get` 安装 `go-luban`：
```
go get github.com/KINGMJ/go-luban
```

## 注意：

1. 对于空切片的行为，通常情况下，`Every` 函数对于空切片或空 map 返回 true，因为没有元素可以违反条件。像 js 的 `every` 函数、lodash 的 `every` 都是返回 true。这一点与 `go-funk` 不同，`go-funk` 对于空切片或空 map 返回 false。