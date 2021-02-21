- go is a compiled language
- you can do string assigment like `var cards string = "Ace of Spades`
- Java, Go, C++ are static typed langauges. We do care about values assigned to variables.

- you can also just write `card := "Ace of spades"`. Equivalent to above syntax. specifically for strings. Go 
  compiled will figure it out. `:=` tells compiled to figure out the type.
- When reassigning to a string, you don't need to do `:=` again.
- we use `:=` when initializing, declaring, assigning in one step.
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
the function name.
j
and you can call it on cards if it's of the typE:

```go
cards := deck{newCard(), newCard(), "King of Spades"}
```

inside the receiver, `d` is very similar to `this` or `self` in python or es6.
- the receiver in a function allows you to call `receiver.method` in code.
  
- if `fruits[start: notIncludingIndex]`, then `fruits[0:2]` will give you 2 elements. This is the range syntax of 
  slices. It can get you a subset of a slice.
- you can return two values separate with a comma from a function.

- type conversion in go is like this: `[]byte("Hi there")`. Useful for when we need to convert data to another type 
  for a function that only accepts a certain type.
  
tthis is how to convert back to slice of strings:
```go
func (d deck) toString() string {
    []string(d)
}
```
in the above code, deck is a `[]string]` type anyway so this makes sense.

- `error` is an actual type that can bee returnd.

```go
func newDeckFromFile(filename string) deck {
// no receiver since we don't have a deck. The reeturn type is a deck
    bs, err := ioutil.ReadFile(filename)
}
```
in above code, if everything goes correctly, `err` is nil. if there is an error while reading, `err` will be populated.
- nil is a value in Go.

- How to write Go tests? file neeeds to end in `_test` and the test itself needs to begin with `TestXXX` where the 
  first X is capitalzed.
- you can writ ea test like this using a formatted error string:

```go
func TestNewDeck(t *testing.T) {
    testDeck := newDeck()
    deckLength := 16
    if len(testDeck) != deckLength {
        t.Errorf("expected deck legnth of 16 but got %v", len(testDeck))
    }
}
```

## module 2
- what is a struct? a data structure with related properties. Similar to a hash in Ruby or a dictionary in ES6. Go 
  can use order for Struct when using a Struct or you can use keywords. You can do both!
- this is a struct and how to use it:

```go

type contactInfo struct {
    email string
    zipCode int
}

type person struct {
    firstName string
    lastName string
    contactInfo
}
```
we can have custom types and use it to define another custom type like above. Notice the shorthand for contactInfo Notice the shorthand for contactInfo
you can create a person like this:

```go
jim := person{
        firstName: "Jim",
        lastName: "Wan",
        contactInfo: contactInfo{
        email: "someEmailForJim@gmail.com",
        zipCode: 90210,
        },
    }
```

- **receivers can be structs too**:

```go
type person struct {
    firstName string
    lastName string
    contactInfo
}

func main() {
// second way to deal with structs
  jim := person{
    firstName: "Jim",
    lastName: "Wan",
    contactInfo: contactInfo{
    email: "someEmailForJim@gmail.com",
    zipCode: 90210,
    },

  jim.print()
}

func (p person) print() {
    fmt.Printf("%+v", p)
}
```
in the above, jim a type person and so you can call these functions with a receiver of person.

_pass by value_
- go is a pass by value language meaning when we pass a struct to a funciton, a copy is made avaialble inside the 
  function. 
- go is a pass by value language. so when we pass a struct, go will copy all data inside the struct and place it at 
  a new address in memory. it does not pass the same struct in memory. the argument in a function is a new struct, 
  same vlaues, different memory location. so if you change a property in a struct in a function, it does not change 
  the original struct. we're only updating the copy.
```go
    jim.updateName("jimmy")
    jim.print()
}

func (p person) print() {
    fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
  p.firstName = newFirstName
}
```

this fixes the problem:

```go
  jimPointer := &jim
  jimPointer.updateName("jimmy")
  jim.print()
}

func (p person) print() {
  fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}
```

so what do `&` and `*` operator do? `&` gives us the memory (RAM) address of the variable. so `jimPointer := &jim` 
so jimPointer is no longer pointing to the struct, instead it's a reference to the memory address. `jimPointer` when 
printed would be a memory address location.
- the `*` operator gives us the value at that memory address. so in the function we see `*pointerToPerson`, that 
  will return us the value at that `pointerToPerson` memory address. So that will give us the original struct and 
  not a copy.
- BUT a `*` in front of a type is different from a `*` in front of a pointer. when in front of a type like in `func 
  (pointerToPerson *person)`, it means it can only be called with a pointer to a `type person`. But in the body, the 
  `*pointerToPerson` is a value at that memory address.
- So we turn address into value with `*address`.
- We turn value into address with `&value`
- But * in front of a type (like in the argument)  that means we're looking or a type that is a pointer to the 
  specified type (in this case a person). It's not a value when you see `*type`, it just means a pointer to a value 
  of that type. So if you see `func (pointerToPerson *person)` read this as a pointer that points to a person. it 
  does not turn person into a value. It's called a `type description`, means we're working with a pointer to a type.
  
_shortcuts with pointers:_
- if our function is defined like this:
```go
  func (pointerToPerson *person) updateName(newFirstName string) {
```
then that means we can call `updateName` with either a person or a pointer to person. So we can write:
```go
  jimPointer := &jim
  jimPointer.updateName("jimmy")
```
or
```go
  jim.updateName("jimmy")
```
so if you have a variable `jim` that is just type person. but the receiver is of type pointer to person, you can still use the person type when calling the receiver function. Go will automatically convert variable of type person to a  pointer to person for you.
- if we define a receiver of type pointer to a type (say person), go will allow us to call the function either with a variable of type person or a pointer to a person.
- this pass by value is true for structs, ints, and strings. for example, the below code prints "Bill":
  
```go
import "fmt"

func main() {
  name := "Bill"
  updateValue(name)
  fmt.Println(name)
}

func updateValue(n string) {
  n = "Alex"
}
```
and does not print Alex because n in the function is a copied value.

- big HOWEVER in go with pointers... if you pass a slice to af unction and modify it, it DOES modify the original 
  slice even without the use of pointers. that is not how structs work. IF you pass a struct and the function 
  modifies the struct, the original struct is untouched.j
  
big catch... when you pass a slice to a function and mutate the slice... the original slice's values are changed. Even though with structs they do not. Why is this?
- Let's first understand arrays and slices a bit more. remember that slices and grow and shrink unlike arrays which cannot be resized. When you create a slice in go, go is creating two separate data structures; the first is a slice with 3 elements (pointer to head, capacity, and length). capacity is how much the slice can hold. the length is the number of elements in the slice. And, the pointer points to underlying array with list of items. The array is at a separate address in memory. So when you pass a slice into a function, go still copies by value (original at 0001)... it does copy the slice and puts it in another memory location (say 0003), BUT it is still pointing to the underlying array (at another location say 0002). There are now two copies of the slice pointing to the same location of the underlying array (0002).
- What else behaves this way. Slcies are a refrence type because they refer to a data structure (array) in another location in memory. maps and channels are the same way. points and functions too. These are reference types. What are value types (when you have to use pointers)? They are ints, floats, structs, bools, and strings.


## Chapter 3: Maps
- map keys must be of same time and map values must be of same time. The key adn maps don't have to be of the same type to each othet. need to delcare both the key and value type separately.j
this is a map:
  ```go
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#745someHex",
  }
  ```
- other ways to declare map:
```go
  var colors map[string]string
```

```go
  colors := make(map[string]string)
```

both are ways to declare maps to be manipulated and populated later. You cannot use dot notation with maps to add value. Have to use square braces. That's because map keys are typed and inside the braces we need to provide the appropriate typed value. You can use `delete` to delete keys and values off of a map. 

- how to iterate through a map?

```go

for color, hex := range c {
    fmt.Println("color:", color, "hex:", hex)
}
```

- difference between map and struct? 
- structs: you cannot iterate, it's a value type (passed by value and copied so mutations don't affect iteration in function calls), you need to know all different fields at compile time (need to list all the property names at compile), values can be of different types, keys do not support indexing
- maps: you can iterate, passed by reference meaning a copy of the reference is passed to the function, you don't need to know all keys at compile time... you can add to it over the execution of the program.

- quick thing with receivers, if you don't use it, you can omit the variable like this:

```go
func (spanishBot) printGreeting() string {
    return "hola!"
}
```

only the receiver type is present since the variable isn't used in the function.
- you can't overload in Go, even if different argument types. This is not allowed:

```go
func printGreeting(eb englishBot) {
  fmt.Println(eb.getGreeting())
}

func printGreeting(sb englishBot) {
  fmt.Println(sb.getGreeting())
}
```

- can we reduce two diff function with two different argument types into one? use interface
```go

type bot interface {
  getGreeting() string
}

func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}
```

we still have two bots, englishbot and spanish bot each with own digferent implementation of getGreeting function:

```go
func (englishBot) getGreeting() string {
  //very custom logic for english greeting, imagine
  return "hi there"
}

func (spanishBot) getGreeting() string {
  return "hola!"
}
```

but there is an interface of `bot` which says a bot has a method of `getGreeting()` So if you have a method named `getGreeting()` and you return `string`, you can be a `bot` type. You can now call `printGreeting` with that bot. If a member has a method named `getGreeting()` that retunr string, it is promoted to type `bot`.
- with interfaces I want to define methods and attributes that something as. As long as it satisfies that, it is of that interface type.
an interface can list out both the argument adn return types:
  
```go
type <name> interface {
	<method>(int) string
	<otherMethod>(int, string) (float,error)
	<yetAnotherMethod>(user) string
}
```

- int is the argument type and string is the return type above
- you cannot create an interface type but you can create a concrete type which is like a map or struct or int or englishBot which is of type struct, it extends a type struct.
- interfaces are implicit in go, we do not have to explicitly say that our custom type satisfies an interface... it just has to have the methods implemented. we did not have to say englishBot is of type `bot`. Go takes care of the magic for us.
- interfaces are not tests, they do not gaurantee a good implementation of the type. Go will not say that an error is made as long as the implementation is made. The return type can be wrong but as long as the implementation is made, Go will say it's satisfied. It's just a contract but the implementation can be wrong.

- http interfaces:
this code is not enough to get function body:

```go
func main() {
  resp, err := http.Get("http://google.com")
  if err != nil {
    fmt.Println("Error", err)
    os.Exit(1)
  }

  fmt.Println(resp)
}
```