## g/>pl

```
g/>pl 1
press Return on empty line to Eval
may the Source be with You

  func fib (n Int) (Int) 
    if < n 2 n + fib - n 1 fib - n 2

[]
  fib 10

[55]
```

### intro
g/>pl is a scripting language/toolkit designed to complement Go.

### syntax
The provided syntax is relatively simple and trivial to customize/replace.

- Forms are separated by whitespace and read left to right.
- The input stream is consumed until all arguments have been collected (recursively) or `EOF`, as a result there is no support for optional/variable arguments.
- Forms may be grouped for macro processing and/or readability using parens.
- All calls including operators are prefix by default.
- The stack is exposed to user code like Forth, `_` may be used to indicate the top value.