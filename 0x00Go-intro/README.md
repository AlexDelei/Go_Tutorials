```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

#### Functions
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

```
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

#### ....Functions
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

```
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

```

#### ...Functions - (Multiple results)
```
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

#### Names return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

**A return statement without arguments returns the named return values. This is known as a "naked" return.**

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(19))
}

```

#### Variables
The var statement **declares a list of variables** ; as in function argument lists, the type is last.

A var statement can be at package or function level.

```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

#### Variables with initilializers
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

#### Short variable declarations
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

**Outside a function**, every statement begins with a keyword (var, func, and so on) and so the **:= construct is NOT available**.

```
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "yeah!"
	fmt.Println(i, j, k, c, python, java)
}
```

#### Basic Types
Go's basic types are:

<ol type="number">
	<li>bool</li>
	<li>string</li>
	<li>int, int8, int16, int32, int64</li>
	<li>uint, uint8, uint16, uint32, uint64, uintptr</li>
	<li>byte <b>Alias for uint8</b></li>
	<li>rune <b>Alias for int32 -- Represents a unicode code point</b></li>
	<li>float32, float64</li>
	<li>complex64, complex128</li>
</ol>


example:
```
package main

import (
	"fmt"
	"math/cpmlx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

output:
```
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
```

#### Zero Values
Variables declared without an explicit initial value are given their __zero value__.

The zero value is:
	0 for numeric types
	false for the boolean type
	"" for empty strings


```
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

#### Type Conversions
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)
Unlike in C, in Go assignment between items of different type **requires an explicit conversion**.


#### Type inference
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant.

#### Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

**Constants cannot be declared using the __:=__ syntax.**

```
package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const truth = true
	fmt.Println("Go rules ?", truth)
}
```

#### Numeric Constants
Numeric constants are high-precision values.

An untyped constant takes the type needed by its context.


(An int can store at maximum a 64-bit integer, and sometimes less.)

```
import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 94
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

output:
```
21
0.2
1.2676506002282295e+29
```

----------------------------------------------------------
----------------------------------------------------------

### For Statements
```
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello Alex!")
	}
	fmt.Println("Hello guys!")
}
```

### Go's While statements
```
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

output:
```
1024
```

----------------------------------------------------------
----------------------------------------------------------

#### If statements
Go's if statements are like its for loops; the expression **need not** be surrounded by parentheses ( ) but the **braces { } are required**.

```
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```
