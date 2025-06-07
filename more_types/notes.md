# Pointers
The * operator denotes the pointer's underlying value.

```
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

This is known as "dereferencing" or "indirecting".

*Unlike C, Go has no pointer arithmetic.*

# Structs
A struct is a collection of fields

Struct fields : Struct fields are accessed using a dot.

## Pointer to Structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

## Struct literals
A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.

# Arrays
The type [n]T is an array of n values of type T
```
var a [10]int
```
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized.

SLICING IN ARRAY -
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

Slicing does not really store any data, it only describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

## Slice literals
A slice literal is like an array literal without the length.

This is an array literal:
```
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:
```
[]bool{true, true, false}
```

## Slice defaults
when slicing an array we can use default higher and lower bound insetad of mentioning the values.
lower bound default is 0
higher bound default is length of slices
```
a[0:10]
a[:10]
a[0:]
a[:]
```

## Slice length and capacity
length of slice is number of elements it contains `len(s)`
capacity of slice is number of elements in underlying array `cap(s)`

We can extend a slice's length by re-slicing it, provided it has sufficient capacity

## Nil slices
The zero value of a slice is nil.
A nil slice has a length and capacity of 0 and has no underlying array.
