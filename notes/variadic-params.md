# Variadic Params

When defining our `ParseFS` function, we only allow a single pattern to be passed into the function.

```go
// views/template.go
func ParseFS(fs fs.FS, pattern string) (Template, error) {
  // ...
}
```

If we were to look at the `ParseFS` function inside of the `html/template` package, we would see that it accepts something a bit different.

```go
// inside the html/template package
func ParseFS(fs fs.FS, patterns ...string) (*Template, error)
```

Using three dots (`...`) before a type turns it into a variadic parameter. What this does is allows developers to pass in zero or more arguments of that type. In other words, the following are all valid ways to call the `template.ParseFS` function.

```go
// Zero patterns. Not very useful in this case, but the code will compile
template.ParseFS(embed.FS)

// One pattern. What we do now for our static pages.
template.ParseFS(embed.FS, "home.gohtml")

// Several patterns. Especially useful if you need to include additional
// files that have reusable HTML components, like an alert.
template.ParseFS(embed.FS, "home.gohtml", "footer.gohtml", "alert.gohtml")
```

Variadic parameters can accept a variable number of arguments, so they must be the last parameter defined in a function. The following is NOT valid.

```go
// NOT valid!
func ParseFS(fs fs.FS, patterns ...string, n int)
```

Let's write our own function with a variadic parameter to explore how they work a bit further. Open `cmd/exp/exp.go` and add the following code.

```go
package main

import (
	"fmt"
)

func main() {
	Demo()
	Demo(1)
	Demo(1, 2, 3)
}

func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}
```

When we call a function that uses a variadic parameter, pass individual values into the function call and just include as many as we want. Here we can see that we call `Demo()` with no numbers, just the single `1` value, and a third time with `1`, `2`, and `3` as if the function took three arguments.

```go
Demo()
Demo(1)
Demo(1, 2, 3)
```

Inside the function is a bit different because we don't know for sure how many arguments we are going to receive. To account for this, our code always needs to treat the variadic parameter as a slice. That means it can be ranged over like any other slice.

```go
func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}
```

We can also access the length of the slice and lookup values at a specific index, just like any other slice. Add the following code to your `exp.go` to see this in action:

```go
func main() {
	fmt.Println(Sum())
	fmt.Println(Sum(4))
	fmt.Println(Sum(4, 5, 6))
}

func Sum(numbers ...int) int {
	sum := 0
  for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
  return sum
}
```

To summarize, a function that accepts a variadic parameter will accept individual values and converts them into a slice for your function to use.

What happens if we already have a slice and need to pass that into a function with a variadic parameter? We can also use the triple dots (`...`) to "unravel" a slice into individual arguments. Add the following code to `exp.go` to see this in action.

```go
func main() {
	fib := []int{1, 1, 2, 3, 5, 8}
	Demo(fib...)
	fmt.Println(Sum(fib...))
}
```

In these example we have been looking at `int` variadic parameters, but any data type can be used. `template.ParseFS` uses strings, and we can create a demo function for those as well:

```go
func main() {
	strings := []string{"the", "quick", "brown", "fox"}
	fmt.Println(Join(strings...))
}

func Join(vals ...string) string {
	var sb strings.Builder
	for i, s := range vals {
		sb.WriteString(s)
		if i < len(vals)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
```

All of this begs the question - why don't we just accept a slice as our last argument? We could use functions like:

```go
func Demo(numbers []int)
func Sum(numbers []int) int
func Join(vals []string) string
```

The primary motivation for using variadic parameters is to make a developers life easier, and they will most frequently be used in cases where developers are commonly hand coding the function call. Take our `ParseFS` function - this is very frequently used with hard-coded strings defining what the template file is on disk. As a result, it makes sense to accept a variadic parameter to make the developers life easier.

We are going to update our `ParseFS` function to accept a variadic parameter. Open up `views/template.go` and make the following changes:

```go
func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}
```

The first change is the function definition. We changed the last argument - `patterns` to be a variadic parameter. After that we need to update the following line, as our `patterns` variable is now going to work like a slice inside of our function. To pass a slice into another function that accepts a variadic parameter, we need to use the `...` suffix as we saw earlier. That leads to the line:

```go
htmlTpl, err := template.ParseFS(fs, pattern...)
```

With that our code is now ready to accept multiple patterns in the `ParseFS` function call. In the future we can include additional template files if we need to use them. For instance, we might create reusable components and include them in our final template.


