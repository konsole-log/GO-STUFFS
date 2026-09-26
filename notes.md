### for testing we need to use:
1. file name: xxx_test.go
2. test function starts with word `Test`
3. test function only takes one argument only `t *testing.T`
4. we need to `import "testing"` to use `*testing.T`

### Cycle for Testing:
1. Write a test
2. Make the compiler pass
3. Run the test, see that it fails and check the error message is meaningful
4. Write enough code to make the test pass
5. Refactor

### Learned til now:
> learned to use `pkgsite -open .` which shows the packages installed alongisde the third party sources
* More practice of the TDD workflow
* Integers, addition
* Writing better documentation so users of our code can understand its usage quickly
* Examples of how to use our code, which are checked as part of our tests

> if you want to use `function` outside of the package file you need to start the `function` name with the capitalized and if you want to just use in the package it shoule start with small letter


For Example:

```go
package integers

func Add(x,y int) int{
    return x+y
}
```
> above `Add(x,y int) int` can be accessed to the other file not in same package as `integers.Add(5,3)`


### creating Example for documentation
We can create documentation for the example in same`xxx_test.go`
1. Example starts with `Example`
###### Example:
```go
package integers

import "fmt"

func ExampleAdd(){
    sum := Add(5,2)
    fmt.Println(sum)
    \\ Output: 7
}
```

### BenchMark:

To run benchmark we write this in `xxx_test.go` file which is same as writing test however we use `BenchmarkXxx` for the function name<br>
simple syntax will be:
```go
func BenchMark<Func_name>(b *testing.B){
    //... setup ...
    for b.Loop(){
        //... code to measure ...
    }
    //... cleanup ...
}
```

> Strings in Go are immutable, meaning every concatenation which involves copying memory to accomodate the new string. This impacts performance, particularly during heavy string concatenation.<br>
> The standard library provides the `strings.Builder` type which minimizes memory copying. It implements `WriteString` method which we can use to concatenate strings 

Example for string concatenation:
```go
package iterations

import "strings"

func Repeat(character string, total int) string {
	var repeated strings.Builder
	for i := 0; i < total; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
```
<b>Note</b>: We have to call the `String` method to retrieve the final result

#### For benchmark we use following command
* `go test -bench=.`
* `go test -bench=. -benchmem`

The `-benchmem` flag reports information about memory allocations:
- `B/op`: the number of bytes allocated per iteration
- `allocs/op`: the number of memory allocations per iteration


### Array
Arrays have a fixed capacity which you define when you declare the variable. We can initialize an array in two ways:

- [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
- [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

###### Example:
```go
package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

```
It is sometimes useful to also print the inputs to the function in the error message. Here, we are using the `%v` placeholder to print the "default" format, which works well for arrays.

### Slices
```go
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
```
> We can assign a function to the variable as:
```go
checkSums := func(t *testing.T, got, want []int) {
    t.Helper()
    if !slices.Equal(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

```
