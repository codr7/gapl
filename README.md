## g/>pl

```
g/>pl 3
press Return on empty line to Eval
may the Source be with You

  func fib (n Int) (Int) 
    if < n 2 n + fib - n 1 fib - n 2
  fib 10

[55]
```

### intro
g/>pl is a scripting language/toolkit designed to complement Go.

### syntax
The provided syntax is relatively simple and trivial to customize/replace.

- Forms are separated by whitespace and read left to right.
- The input stream is consumed until all arguments have been collected (recursively) or `EOF`.
- All calls including operators are prefix.
- Forms may be grouped using parens.

### the stack
`d` may be used to drop the top `n` values.

```
  1 2 3 4 5

[1 2 3 4 5]
  dd

[1 2 3]
```

### functions
New functions may be defined using `func`.

```
  func foo () (Int) 42

[]
  foo

[42]
```

Anonymous functions may be created by simply omitting the name.

```
  func () (Int) 42

[Func(() (Int))]
  call _

[42]
```

Functions are lexically scoped,

```
  func foo () () (
    func bar () (Int) 42
    bar
  )

[]
  foo

[42]
  bar

Error in repl at line 0, column 0: Unknown id: bar
```

#### call flags
Call flags may be specified by prefixing with `|`.

##### |d(rop)
Drops returned values.

```
  func foo () (Int) 42
  foo|d
  
[]
```

##### |t(co)
Performs tail call optimization, may be used outside of tail position which causes an immediate return.

```
  func foo (n Int) (Int)
    if = n 0 n foo|t - n 1
  foo 42
  
[0]
```

### performance
g/>pl currently runs around 6 times as slow as Python3.

`bench` runs the specified body `n` times and returns elapsed time in milliseconds.

```
  func fib (n Int) (Int) 
    if < n 2 n + fib - n 1 fib - n 2
  bench 100 fib|d 20

[1344]
```