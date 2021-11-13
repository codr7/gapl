## g/>pl

```
g/>pl 1
press Return on empty line to Eval
may the Source be with You

  func fib (n Int) (Int) 
    if < n 2 n + fib - n 1 fib - n 2

[]
  bench 100 fib 20

[55]
```

### intro
g/>pl is a scripting language/toolkit designed to complement Go.

### syntax
The provided syntax is relatively simple and trivial to customize/replace.

- Forms are separated by whitespace and read left to right.
- The input stream is consumed until all arguments have been collected (recursively) or `EOF`.
- Forms may be grouped using parens.
- All calls including operators are prefix.
- The stack is exposed to user code like in Forth, `_` may be used to indicate the top value.

### performance

g/>pl currently runs around 6 times as slow as Python3.

`bench` runs the specified body `n` times and returns elapsed time in milliseconds.

```
  func fib (n Int) (Int) 
    if < n 2 n + fib - n 1 fib - n 2

[]
  bench 100 fib 20

[1406]
```