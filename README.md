# **Invoker**
A library to call (invoke) functions by taking names and sample inputs of those functions as parameters. And returns the types and values of output.

## Purpose
Testing the result of multiple functions and business logic.

## Usage
```
go get github.com/odese/invoker
```

### Call Example for Single Function: 
```
package main

import (
	"fmt"

	"github.com/odese/invoker"
)

func main() {
	outputTypes, outputValues, err := invoker.Invoke(Sum, []interface{}{1, 2})
	fmt.Println("outputTypes: ", outputTypes)
	fmt.Println("outputValues: ", outputValues)
	fmt.Println("err: ", err)
}


func Sum(a, b int) (r Response, sum int, err error) {
	r.Field1 = a + b
	r.Field2 = "Sum2 Response"
	sum = a + b
	return r, sum, err
}

type Response struct {
	Field1 int
	Field2 string
}

// Output:
// outputTypes:  [<main.Response Value> <int Value> <error Value>]
// outputValues:  [{3 Sum2 Response} 3 <nil>]
// err:  <nil>
```
### Call Example for Consecutive Functions:
```
package main

import (
	"fmt"

	"github.com/odese/invoker"
)

func main() {
	args := make([]invoker.FuncArguments, 0)
	var arg invoker.FuncArguments
	arg.Function = Sum
	arg.Inputs = []interface{}{1, 2}
	args = append(args, arg)

	arg.Function = Sum
	arg.Inputs = []interface{}{2, 3}
	args = append(args, arg)

	for i := 0; i < len(args); i++ {
		outputTypes, outputValues, err := invoker.Invoke(args[i].Function, args[i].Inputs)
		fmt.Println("outputTypes: ", outputTypes)
		fmt.Println("outputValues: ", outputValues)
		fmt.Println("err: ", err)
	}
}

func Sum(a, b int) (r Response, sum int, err error) {
	r.Field1 = a + b
	r.Field2 = "Sum2 Response"
	sum = a + b
	return r, sum, err
}

type Response struct {
	Field1 int
	Field2 string
}

// Output:
// outputTypes:  [<main.Response Value> <int Value> <error Value>]
// outputValues:  [{3 Sum2 Response} 3 <nil>]
// err:  <nil>
// 
// outputTypes:  [<main.Response Value> <int Value> <error Value>]
// outputValues:  [{5 Sum2 Response} 5 <nil>]
// err:  <nil>
```
### Call Example for Arbitrary Ordered Functions:
```
package main

import (
	"fmt"

	"github.com/odese/invoker"
)

func main() {
	var arg1 invoker.FuncArguments
	arg1.Function = Sum
	arg1.Inputs = []interface{}{1, 2}
	
	var arg2 invoker.FuncArguments
	arg2.Function = Sum
	arg2.Inputs = []interface{}{2, 3}
	
	funcMap := map[string]invoker.FuncArguments{
		"b" : arg2,
		"a" : arg1,
	}

		for k, v := range funcMap {
		fmt.Println(k)
		outputTypes, outputValues, err := invoker.Invoke(funcMap[k].Function, v.Inputs)
		
		fmt.Println("outputTypes: ", outputTypes)
		fmt.Println("outputValues: ", outputValues)
		fmt.Println("err: ", err)
	}
}

func Sum(a, b int) (r Response, sum int, err error) {
	r.Field1 = a + b
	r.Field2 = "Sum2 Response"
	sum = a + b
	return r, sum, err
}

type Response struct {
	Field1 int
	Field2 string
}

// Output:
// b
// outputTypes:  [<main.Response Value> <int Value> <error Value>]
// outputValues:  [{5 Sum2 Response} 5 <nil>]
// err:  <nil>
// a
// outputTypes:  [<main.Response Value> <int Value> <error Value>]
// outputValues:  [{3 Sum2 Response} 3 <nil>]
// err:  <nil>
```
