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

### switch evaluation order
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
```
switch i {
case 0:
case f():
}
```
does not call f if i==0

> The Go programming language's playground, an online tool for running Go code, sets its time to always start at 2009-11-10 23:00:00 UTC. This specific timestamp is a nod to the history of the Go language itself.
Explanation:
Significance of the Date: November 10, 2009, is the date when the Go programming language was officially announced to the public by Google. This was when Go was first released as an open-source project, marking a significant milestone in its development.
> Why a Fixed Time? The Go playground uses a deterministic environment to ensure consistent and reproducible results for code execution. By fixing the time to 2009-11-10 23:00:00 UTC, the playground ensures that any time-dependent code (e.g., functions like time.Now()) yields predictable outputs, which is crucial for testing and sharing code snippets reliably.
> Cultural Reference: The phrase "left as an exercise for the reader" is a playful hint often used in technical documentation, suggesting that the significance of the date is something enthusiasts or curious developers can uncover themselves. In this case, it points to the Go announcement date, rewarding those familiar with the language's history.

### switch with no condition
Switch without a condition is the same as switch true

## Defer
A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

stacking defer -> https://go.dev/blog/defer-panic-and-recover
