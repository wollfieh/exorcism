# discussion from code review on exercism

## The idea from the reviewer on exercism

is to create a func that receives the filter and the slice, and execute the repeating behavior of iterate, filter, append generically, e.g.

``` go
func ApplyFilterTo(filter func(X) bool, slice []X) []X {
  var result X
 for _, v := range slice {
  if filter(v) {
   result = append(result, v)
  }
 }
 return result
}
```

And then you will only do something like this inside the other functions:

```go
func (i Ints) Keep(filter func(int) bool) Ints {
 return ApplyFilterTo(filter, i)
}
```

Note: not sure if everything compiles, can't check it right now, but as soon as I can I will try locally.

## my thoughts on this

ok I see where we're going with this. This all would be inside the package, so you could keep the applyFilterTo local and just expose the types.

I don't know yet how I feel about the mixing of OO style methods and functions like ApplyFilterTo , any idea on that?

I'll try to make your example run, I find this interesting.
