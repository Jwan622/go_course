- go is a compiled language
- you can do string assigment like `var cards string = "Ace of Spades`
- Java, Go, C++ are static typed langauges. We do care about values assigned to variables.

- you can also just write `card := "Ace of spades"`. Equivalent to above syntax. specifically for strings. Go 
  compiled will figure it out. `:=` tells compiled to figure out the type.
- When reassigning to a string, you don't need to do `:=` again.

```go
func newCard() string {
    return "Five of Diamonds"
}
```

in above, you tell function what newCard returns a value of type string. This is needed if you do something like: 
`card := newCard()`. The compiler sees that you're doing assignment so newCard has to return something, not nothing 
sees that you're doing assignment so newCard has to return something, not nothing.

- array is fixed length in go. but a slice is an array that can grow or shrink.
- slices and arrays must be of identical type because they are defined with a data type.
- an array or slice cannot be ['five', 102, 'six]
- use `range` keyword to iterate over inside of a slice.

Go is not an object oriented langauge. no classes.

- you can create a custom type in Go like this:
```go
type deck []string
```
the above type is the same a slice of strings.

What are receivers? Receivers are functions that can be called on any variable of the receiver's type. This is a 
receiver:

```go
func (d deck) print() {
    for i, card := range d {
    fmt.Println(i, card)
    }
}
```
in the above receiver, the `deck` variable is available in the function as a variable called `d`. Now every variable 
of type `deck` can call the `print` function on itself. usually by convention the variable is 1 or 2 letters 
starting with the same letter as the type. Thatt's why it's simply just `d`. notee the receiver is the part beefore 
the function name.}
j
and you can call it on cards if it's of the typE:

```go
cards := deck{newCard(), newCard(), "King of Spades"}
```

inside the receiver, `d` is very similar to `this` or `self` in python or es6.