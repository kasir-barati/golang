# Install Go

```cmd
# Manjaro linux
sudo pacman -Syu go
```

# Structure of Go Code

- Packages:
  - A bunch of go files :smile:.
- Modules:
  - A bunch of packages :smile:.
  - So basically when we initiate a new project we are creating a new module.

```cmd
go mod init github.com/kasir-barati/golang
```

> [!TIP]
>
> The name of module usually is the URL of a GitHub repository. But we can use any name we want, i.e. `git mod init test`.

![go mod init result is a go.mod file](./assets/go-mod-res.png)

Now you can start creating new packages:

```cmd
code src/utils/main.go
```

- This is a special package, it is the entry point of the program.
- It is where the compiler starts executing the code from.
- You must define a function called `main` in this package. **This is mandatory**.

Now it's time to build and run the program:

```cmd
go build src/utils/main.go
```

- This will create a binary file called `main` at the root of the project.
- I gitignore the output of this command.

Now let's run the program:

```cmd
./main
```

> [!TIP]
>
> You can also run the program with the `go run src/utils/main.go` command.

# Variables & Constants

- To declare a variable we use the `var` keyword.
- The type of the variable can be inferred from the value.
- We can also specify the type of the variable explicitly.

```go
var name string = "John"
```

> [!TIP]
>
> We must use every variable we define, otherwise the compiler will throw an error.

<table>
  <thead>
    <tr>
      <th colspan=5>Variable Types</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>int</code>, <code>uint</code></td>
      <td><code>int8</code>, <code>uint8</code></td>
      <td><code>int16</code>, <code>uint16</code></td>
      <td><code>int32</code>, <code>uint32</code></td>
      <td><code>int64</code>, <code>uint64</code></td>
    </tr>
    <tr>
      <td>-</td>
      <td>-</td>
      <td>-</td>
      <td><code>float32</code></td>
      <td><code>float64</code></td>
    </tr>
    <tr>
      <td><code>string</code></td>
      <td>-</td>
      <td>-</td>
      <td>-</td>
      <td>-</td>
    </tr>
    <tr>
      <td><code>rune</code></td>
      <td>-</td>
      <td>-</td>
      <td>-</td>
      <td>-</td>
    </tr>
    <tr>
      <td><code>bool</code></td>
      <td>-</td>
      <td>-</td>
      <td>-</td>
      <td>-</td>
    </tr>
  </tbody>
</table>

> [!TIP]
>
> **Overflow**:
>
> The size of the variable matters and in runtime you might run into the issue of overflow.
>
> ```go
> var score int16 = 32767
> score++
> println(score) // -32768
> var avg float32 = 999999.99
> fmt.Println(avg) // 1e+06
> ```

> [!NOTE]
>
> **`len` VS `utf8.RuneCountInString`**:
>
> `len` of a string is the number of bytes, **NOT** the number of characters!
>
> ```go
> var name string = "Öl"
> println(len(name)) // 3
> ```
>
> To get the number of characters we can use the `utf8` package.
>
> ```go
> import "unicode/utf8"
> var name string = "Öl"
> println(utf8.RuneCountInString(name)) // 2
> ```

- We can use the `int` type which picks `int23` or `int64` based on the system architecture.
- Int division is always an integer.
- `rune`s are equivalent of `char` types in c++.
- The default value of a variable depends on its type, e.g. `int` is `0`, `string` is `""`, etc.
- We can use the `:=` operator to declare and initialize a variable in one line, e.g. `name := "Mahdi"`.
  - Sometimes it is better to be as explicit as possible, `res := someFunc()`.

> [!NOTE]
>
> `const` is immutable & they must be initialized at declaration.

# Functions

- We can define functions in Go by using the `func` keyword.
- We need to let the compiler know the return type of the function + the types of the parameters.

```go
package main

import (
	"fmt"
)

func main() {
	var num1 int = 11
	var num2 int = 20
	var result int = add(num1, num2)
	fmt.Println(result)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
```

- We can return multiple values from a function.
  - **Note**, this ain't the same as `tuple`s in Python.

```go
func divide(numerator int, denominator int) (int, int) {
  var result int = numerator / denominator
  var remainder int = numerator % denominator
  return result, remainder
}
```

> [!TIP]
>
> A common pattern in Go is to return an error as the last value:
>
> ```go
> import "errors"
> func divide(numerator int, denominator int) (int, int, error) {
>   if denominator == 0 {
>      return 0, 0, errors.New("cannot divide by zero")
>   }
>   var result int = numerator / denominator
>   var remainder int = numerator % denominator
>   return result, remainder, nil
> }
>
> var divisionResult, divisionRemainder, error = divide(num1, num2)
> if error != nil {
>   fmt.Println(error.Error())
>   return
> }
> ```

# Logical Operators & Comparison Operators

<table>
  <tbody>
    <tr>
      <td><code>==</code></td>
      <td><code>!=</code></td>
      <td><code>&&</code></td>
      <td><code>||</code></td>
      <td><code>></code></td>
      <td><code><</code></td>
    </tr>
    <tr>
      <td><code>if res == 10 {</code></td>
      <td><code>if err != nil {</code></td>
      <td><code>if num2 > 0 && num1 > 0 {</code></td>
      <td><code>if res == 2 || res == 4 {</code></td>
      <td><code>if num1 > 0 {</code></td>
      <td><code>if num2 < 0 {</code></td>
    </tr>
  </tbody>
</table>

> [!NOTE]
>
> - If the first statement in an `if` statement is true, the second statement is **NOT** evaluated for the `||` operator.
> - If the first statement in an `if` statement is false, the second statement is **NOT** evaluated for the `&&` operator.

# `switch` Statement

- No need to use `break` in Go.
  - It is implied.
- You can have a `switch` statement without a condition.

```go
var num int = 10
switch num {
  case 1, 2:
    println("num is 1 or 2")
  case 10:
    println("num is 10")
  case 20:
    println("num is 20")
  default:
    println("num is not 10 or 20")
)
switch {
  case num > 0:
    println("num is positive")
  // ...
}
```

# Arrays

- Fixed-size.
- Homogeneous.
- Indexed.
  - Zero-based.

```go
import "fmt"

var scores [5]int = [5]int{1, 2, 3, 4, 5}
fmt.Println(scores[:])   // [0 0 0 0 0]
fmt.Println(scores[1:3]) // [2 3]
```

> [!TIP]
>
> ```go
> scores := [...]int{1, 2, 3, 4, 5}
> var names [3]string
> ```

# Slices

- Wrappers around arrays.
- Omit the size of the array, and you'll have your slice.
- It will automatically resizes the array for you.
  - This has performance implications.
  - So better to think carefully about the size/capacity of the slice.

```go
var scores []int
scores = append(scores, 1)
```

> [!TIP]
>
> Use speared operator to append a list of values to a slice:
>
> ```go
> numbers := []int{1, 2, 3, 4, 5}
> scores = append(scores, numbers...)
> ```

## `make` Function

- Creates a slice with a specified length and capacity.
- The capacity is the size of the underlying array.
  - This means that you can create a bigger slice than what you currently need in order to avoid recreating the underlying array when you need to add more elements.

```go
var scores []int = make([]int, 5, 10)
```
