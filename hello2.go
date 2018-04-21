// helloworld.go
package hello2

import (
	"fmt"
)

type fizzBuzzRet struct {
	msg string
	sum int
	b   int
}

// exportable package
// https://medium.com/golangspec/exported-identifiers-in-go-518e93cc98af
func RunHello2() {
	fmt.Printf("main running")
	mystring := "test 2"
	result := hello(mystring)
	fmt.Printf("\n privyit %s \n", result)
	myslice := []int{1, 2, 3}
	fmt.Printf("\n list %d", myslice)
	message, sum, sum2 := example1(myslice)
	fmt.Printf("is golang function pass by ref by default?")
	fmt.Printf("\n list %d", myslice)
	// fzbz := new(fizzBuzzRet) // declaring this doesnt work
	fzbz := fizzBuzzRet{}
	fzbz = example2(myslice)
	fmt.Printf("\n ex1 result %s %d %d  \n", message, sum, sum2)
	fmt.Printf("\n ex2 result %v", fzbz)
	// fmt.Println(float64(sum*3) / 2)
	myMap := map[string]int{"foo": 1, "baz": 3}
	//https://blog.golang.org/go-maps-in-action
	fmt.Println("\n map", myMap["baz"])
	f := func(n int) int { return n * 2 }
	g := func() int { return 4 }
	fmt.Println("\n named functions", f(g()))
	// playing with slices
	size := 10
	capacity := 11
	myslice2 := make([]int, size, capacity)
	fmt.Println(myslice2)
	fmt.Println("len ", len(myslice2))      // length of dynamic slice
	fmt.Println("capacity ", cap(myslice2)) // capacity is length of underlying array

}

// https://stackoverflow.com/questions/11123865/golang-format-a-string-without-printing
// https://www.tutorialspoint.com/c_standard_library/c_function_sprintf.htm
func hello(arg string) string {
	fmt.Printf("\n %s", arg) // same as C lang string interpolation
	return " mir"
}

// passing slice or array into a func
// as opposed to passing a pointer
//https://stackoverflow.com/questions/21719769/passing-an-array-as-an-argument-in-golang
//https://stackoverflow.com/questions/2439453/using-a-pointer-to-array
// https://stackoverflow.com/questions/25332725/use-a-literal-for-type-in-assignment
//https://blog.golang.org/defer-panic-and-recover
//http://jsplain.com/javascript/index.php/Thread/182-Go-Golang-error-message-multiple-value-in-single-value-context/
func example1(arr []int) (string, int, int) {
	fmt.Printf("\n ex1 running")
	val := 0
	for i := 0; i < len(arr); i++ {
		fmt.Println("\n arr ", arr[i], i)
		val = arr[i]
		arr[i] = 0
		if val%3 == 0 {
			fmt.Println("fizz", val, i)
		} else {
			fmt.Println("buzz", val, i)
		}
	}
	//https://gistpages.com/post//https://gistpages.com/posts/go-lang-too-many-arguments-to-returns/go-lang-too-many-arguments-to-return
	//http://ernestmicklei.com/2013/11/from-multiple-to-single-value-context-in-go/
	return "ex1 sum is", val, 0
}

//https://letslearngo.wordpress.com/2016/01/09/struct-map-array-and-slice/
//https://dzone.com/articles/try-and-catch-in-golang
//https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html
func example2(arr []int) fizzBuzzRet {
	fmt.Printf("\n ex2 running")
	s := fizzBuzzRet{"fb2 sum is", 0, 0}
	for i := 0; i < len(arr); i++ {
		s.sum += arr[i]
		fmt.Println("\n arr ", arr[i], i)
		if s.sum%3 == 0 {
			fmt.Println("fizz", arr[i], i)
			s.b = i
		} else {
			fmt.Println("buzz", arr[i], i)
			s.b = i
		}
	}
	//https://gistpages.com/post//https://gistpages.com/posts/go-lang-too-many-arguments-to-returns/go-lang-too-many-arguments-to-return
	//http://ernestmicklei.com/2013/11/from-multiple-to-single-value-context-in-go/
	return s
}

// https://blog.golang.org/go-slices-usage-and-internals
// https://blog.golang.org/go-maps-in-action
//https://www.youtube.com/watch?v=fhdA-6LcOxk
func example3(arr []int) int {
	// assert that arr != nil
	return 0
}
