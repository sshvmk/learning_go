## multiple imports can be done in 2 ways :
1. import (
    "fmt"
    "math"
   )
2. import "fmt"
   import "math"

## exported names
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.
When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

## Functions
a function can take zero or more arguments
eg :-
```
func add(x int, y int) int {
	return x + y
}
```

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
eg :-
```
func add(x, y int) int {
	return x + y
}
```

named return values
check split function

## variables
`var` statement declares a list of variables. this var statement can be at package or function level.

var declaration can include initialisers (one per variable)
if you initialise a variable, giving it a type can be omitted. the variable takes the type of initialiser

short variable declaration
:= short assignment statement can be used in place of var declaration with implicit type inside a function
outside a function, every statement begins with a keyword (var, func and so on)

## Basic types (basic-types.go)
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

Variables declared without an explicit initial value are given their zero value.

The zero value is:
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

Type conversion

Type inference
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:
```
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

## Constants (constants.go)
Constants are declared like variables, but with the const keyword. Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

NUMERIC CONSTANTS
Numeric constants are high-precision values.
An untyped constant takes the type needed by its context.
