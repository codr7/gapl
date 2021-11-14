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

### status
The functionality described in this document is implemented and verified to work but not much else at the moment.
Any ideas for improvements are most welcome.

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

### bindings
Identifiers may be bound once per scope using `let`.

```
  let foo 42

[]
  let foo 42

Error in repl at line 0, column 0: Duplicate binding: foo 42
```

`_` may be used as a placeholder to pop the stack.

```
  42

[42]
  let foo _

[]
  foo

[42]
```

### functions
New functions may be defined using `func`.

```
  func foo () (Int) 42

[]
  foo

[42]
```

Anonymous functions may be created by omitting the name.

```
  func () (Int Int) (35 7)

[Func(() (Int Int))]
  call _

[35 7]
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

and capture their defining environment.

```
  func foo () (Func) (
    let bar 42
    func () (Int) bar
  )

[]
  call foo

[42]
```

#### call flags
Call flags may be specified by prefixing with `|`.

##### |d(rop)
Drops returned values.

##### |t(co)
Performs tail call optimization.

##### |u(unsafe)
Disables all type checks for the duration of the call.

### performance
g/>pl currently runs around twice as slow as Python3, any ideas on how to make it run faster are most welcome.

```
$ cd bench
$ python3 fibrec.py
233
```

```
  func fibrec (n Int) (Int) 
    if < n 2 n + fibrec - n 1 fibrec - n 2
  bench 100 fibrec|d|u 20

[458]
```

```
$ python3 fibtail.py
105
```

```
  func fibtail (n Int a Int b Int) (Int)
    if = n 0 a if = n 1 b fibtail|t - n 1 b + a b
  bench 10000 fibtail|d|u 70 0 1

[112]
```