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
+ Operator overloading
+ More to come!

The intention of the language is to allow you to write strongly typed functional algorithms while still having the vast (and growing) Golang libraries.  Ideally, a system's program could use both GoF and Golang interchangeably.

===

Lambdas
=======

Lambdas are the only way to define a function or method:

```
fibonacci -> n int -> int
  subFib -> a int -> b int -> n int-> int
    match n
      0,1,2 -> a
      n     -> subFib a+b a n-1
  subFib n
```  

  This example actually demonstrates a few things:
  + Lambdas (`->` syntax)
  + Pattern Matching (`match` keyword)
  + Tabs over brackets
  + Lack of `return` keyword
  + Tail Recursion
   
  
Generics
========

Generics are not defined explicitly, but by leaving type information out for the compiler to determine:

```
double -> a -> a'
  a + a
  
double 2 // Results in 4
```

The only requirement is that whatever type `a` turns out to be, has the `+` operator. Also notice that the return type: `a'` (pronouced A-Prime). This is how a generic return type is defined.  It means that whatever type `a` is, that's the type the function will return.  So say there are two parameters:

```
encode -> a -> b -> b`
  a.encode b
```  

This function will take two parameters (`a` and `b`) where the type `a` has a method `encode` that takes type `b`.  The method `encode` will have to return the type `b`

Pattern Matching
================

Pattern matching is nothing more than a fancy `switch` statement, however most would argue that the syntax is much cleaner.  The way it works is by defining several `lambdas` with the desired value that result in that lambda being called:

```
randomName -> n int -> string
  0 -> "Bill" // Called if n is 0
  9 -> "Sandy" // Called if n is 9
  n -> "John" // default
```

Immutability
============

GoF doesn't allow a variable to be altered past initialization.  This is similar to most functional languages.  The thought is that if one has several variables that are being altered, then the code is more difficult to troubleshoot, and far easier to break.  Instead, think recursion.  If you need a value to be altered, then you need recursion.

This has some large implications, such as no loops! Again, think recursion (ideally tail recursion).


Higher order functions
======================

Operator overloading
====================
