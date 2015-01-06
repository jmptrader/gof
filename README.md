GoF
===

A Functional programming language that compiles down to Go.  Currently GoF is nothing more than a set of ideas.

If you have some that you would like to share, please feel free!  GoF is intended to be what Go hasn't for coders, so if you seem a feature that would fit, speak up!

===

Go Functional (GoF) has many of the features that Go users crave:
+ Lambdas (the only way to define a function)
+ Generics
+ Pattern Matching
+ Immutability
+ Higher order functions
+ More to come!

The intention of the language is to allow you to write strongly typed functional algorithms while still having the vast (and growing) Golang libraries.  Ideally, a system's program could use both GoF and Golang interchangeably.

===

Language Details
================

####Spaces vs. Tabs

GoF uses tabs instead of spaces to match what `go fmt` enforces.  Even though `go fmt` won't work with GoF code, the convention is still adhered to.

####Numeric Constants

GoF is strict with types, so it picking the correct numeric type is important.  Like many languages, use a suffix to tell the compiler which numeric type was intended.

Type    | Suffix
----    | ------
int8    | b
uint8   | ub
int16   | h
uint16  | uh
int32   | (default)
uint32  | ui
int64   | l
uint64  | ul
float32 | f
float64 | (default with .0)

NOTE: The suffix IS case sensitive. Meaning a `3UI` is invalid.

To denote a hex value, prefix the number with a `0x`. For octal, prefix the value with a `0`.  This is the same as Go.

####Functions
=======
Functions
=======
Language Details
================

####Spaces vs. Tabs

GoF uses tabs instead of spaces to match what `go fmt` enforces.  Even though `go fmt` won't work with GoF code, the convention is still adhered to.

####Numeric Constants

GoF is strict with types, so it picking the correct numeric type is important.  Like many languages, use a suffix to tell the compiler which numeric type was intended.

Type    | Suffix
----    | ------
int8    | b
uint8   | ub
int16   | h
uint16  | uh
int32   | (default)
uint32  | ui
int64   | l
uint64  | ul
float32 | f
float64 | (default with .0)

NOTE: The suffix IS case sensitive. Meaning a `3UI` is invalid.

Lambdas
=======

=======
Functions have a simple syntax:
```
func fibonacci -> n int -> int
  func subFib -> a int -> b int -> n int-> int
    match n
      0,1,2 -> a
      n     -> subFib a+b a n-1
  subFib n
```

  This example actually demonstrates a few things:
  + Function Currying
  + Pattern Matching (`match` keyword)
  + Tabs over brackets
  + Lack of `return` keyword
  + Tail Recursion (not yet implemented)

  Important requirements of a function:
  + First token is the name of the function.
  + Last token is the return type. (Use `unit` if there isn't one)
  + An argument in a function declaration has to have a name assigned to it (e.g. `n`).

####Generics

Generics are not defined explicitly, but by leaving type information out for the compiler to determine:

```
func double -> a' -> a'
  a + a

double 2 // Results in 4
```

The only requirement is that whatever type `a` turns out to be, has the `+` operator. Also notice the types: `a'` (pronouced A-Prime). This is how a generic type is denoted.  So say there are two parameters:

```
func encode -> a' -> b' -> b`
  a.encode b
```

This function will take two parameters (`a` and `b`) where the type `a` has a method `encode` that takes type `b`.  The method `encode` will have to return the type `b`

####Pattern Matching

Pattern matching is nothing more than a fancy `switch` statement, however most would argue that the syntax is much cleaner.  The way it works is by defining several `lambdas` with the desired value that result in that lambda being called:

```
func randomName -> n int -> string
  match randomName
    0 -> "Bill" // Called if n is 0
    9 -> "Sandy" // Called if n is 9
    n -> "John" // default
```

####Immutability

GoF doesn't allow a variable to be altered past initialization.  This is similar to most functional languages.  The thought is that if one has several variables that are being altered, then the code is more difficult to troubleshoot, and far easier to break.  Instead, think recursion.  If you need a value to be altered, then you need recursion.

This has some large implications, such as no loops! Again, think recursion (ideally tail recursion).


####Higher order functions

Functions can take functions as a parameter:

```
func funcA -> a int -> func funcB -> b int -> int -> int
  funcB a
```

This example demonstrates a function (`funcA`) that takes an `int` and a function (`funcB`) and returns an `int`. `funcB` takes an `int` and returns an `int`.

Functions can return a function:

```
func addTo -> a int -> func -> int -> int
  func -> b int -> int
    a + b
```

This example demonstrates anonymous functions and that the last entry is returned from the function.

Functions can be partially called:

```
func add -> a int -> b int-> int
  a + b

x = add 1 // x is now a func -> a int -> int
y = x 2 // y is now 3
```

GoF Tool
========

####Options

The GoF tool (`gof`) is intended to resemble Golang's tool (`go`), however with different options:

######Summary
Option   | Description
------   | -----------
generate | Generate .go files from .gof files

######generate
`gof generate [DIRECTORY]`

Generate takes a directory (or `.` if none is provided) and converts each `.gof` file to a `.go` file.  It will name it using the `.gof`'s name.
```
go run gof.go generate tests
tests/singleFunc.gof -> tests/singleFunc.go
tests/twoFunc.gof -> tests/twoFunc.go
```

Notes
=====

This is a new project and really a proof of concept.  Types (among other features) haven't been nailed down yet, but are intended to be added.
