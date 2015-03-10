# go-dynamic

Dynamic variable binding for Go

This package implements dynamic variable binding (a Lispy take on
thread-local storage) for a modified version of Go (specifically, [my
fork](https://github.com/dkolbly/go), which adds reflection on
goroutine ids and a concept of goroutine group ids).

The [gls](https://github.com/jtolds/gls) package is also very
interesting (and employs a truly creative systems hack!) and allows
you to do something very similar, but seems like it might be slow
(I haven't verified this) and doesn't support inheritance across
new goroutines which is important for my usecase (logging request
context across a family of sub-goroutines).  On the plus side,
`gls` doesn't require a patched Go :-)


