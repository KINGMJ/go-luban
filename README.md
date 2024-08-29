# go-luban
A modern and type-safe Go utility library providing a collection of helpers (Map, Filter, Reduce, Each, Every, ...) based on Go generics.

## Description
`go-luban` is a utility library designed to simplify common operations on collections such as arrays, slices, and maps. It is built using Go's powerful generics, making it both modern and type-safe. This library draws inspiration from JavaScript's Lodash and the Go-funk library, but with significant improvements in safety and usability.

## Why go-luban?
- **Generics-Powered**: Unlike traditional utility libraries that rely on reflection, `go-luban` leverages Go's generics to offer type safety and performance.

- **Enhanced Safety**: go-luban addresses the shortcomings of libraries like Go-funk, where the use of `any` types often led to nil pointer exceptions. By ensuring type safety, `go-luban` reduces the risk of runtime errors.

- **Modern API**: With a focus on simplicity and ease of use, `go-luban` provides a clean and intuitive API for developers. Functions like `Map`, `Filter`, `Reduce`, `Each`, and `Every` work seamlessly with your data structures, offering the flexibility you need without compromising on safety.

## Features

### Collections

- [x] Map
- [x] Filter
- [x] Reduce
- [x] Each
- [x] EachRight
- [x] Every
- [x] Some
- [x] Find

### Collections (Map)

- [x] MapMap
- [x] FilterMap
- [x] ReduceMap
- [x] EachMap
- [x] EveryMap
- [x] SomeMap
- [x] FindMap

## Installation
To install `go-luban`, use `go get`:

```
go get github.com/KINGMJ/go-luban
```

## Note

### 1. Behavior of `Every` and `Some` with Empty Slices and Maps
By design, the `Every` function returns `true` for empty slices or empty maps because there are no elements to violate the condition. This behavior is consistent with functions like JavaScript's `every` and Lodash's `every`, both of which return `true` for empty collections. This is different from `go-funk`, which returns `false` for empty slices or maps.

On the other hand, `Some` checks whether **any element satisfies the condition**. For empty slices or maps, since there are no elements to check, it returns `false`.

This logic is widely adopted in many programming languages, including `JavaScript` and `Python`, ensuring consistency and predictability.

### 2. `Find` Function
Since Golang supports multiple return values, the Find function can directly return the key and value of the found element, making it unnecessary to use a FindIndex function.

## Documentation
For detailed documentation, examples, and API references, please visit the go-luban documentation.

## Contributing
Contributions are welcome! Please submit issues and pull requests on the GitHub repository.

## License
This project is licensed under the MIT License. See the LICENSE file for details.