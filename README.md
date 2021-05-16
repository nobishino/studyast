
```go
type MyInt int
type Constraint interface {
    int, int64
}
// int, int64, MyIntなどはConstraintを実装する
```

```go
type Constraint interface {
    ~int | ~int64
}
// int, int64, MyIntなどはConstraintを実装する
```

## 具体例

- `interface{ SomeMethod() }` の型セットは `SomeMethod()`を実装する全ての型からなる集合
- `int` の型セットは`int`のみからなる集合
- `~int` の型セットは、underlying typeが`int`である全ての型からなる集合
- `~MyInt` の型セットは、underlying typeが`MyInt`である全ての型からなる集合