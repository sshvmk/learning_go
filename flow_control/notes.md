## Loops
there is only for loop in Go
there are no parentheses surrounding the three components of the for statement and the braces {} are always required.

The init and post statements are optional.

FOREVER - If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
```
func main() {
	for {
	}
}
```

## if and else
Variables declared inside an if short statement are also available inside any of the else blocks.

## Switch
A switch statement is a shorter way to write a sequence of if-else statements.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
