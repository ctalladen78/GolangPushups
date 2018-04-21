// helloworld.go
package main
import (
	"fmt"
	"hello2"
)

func main() {
  fmt.Printf("Hello, world!")
  mystring := "test"
  result := test(mystring)
  fmt.Printf("\n privyit %s",result)
  list := []int{1,2,3}
  sum := sum(list)
  fmt.Printf("\n sum is: %d",sum)
  hello2.RunHello2()
}

// https://stackoverflow.com/questions/11123865/golang-format-a-string-without-printing
// https://www.tutorialspoint.com/c_standard_library/c_function_sprintf.htm
func test( arg string ) string {
  fmt.Printf("\n %s", arg) // same as C lang string interpolation
  return " mir"
}

// passing slice or array into a func
// as opposed to passing a pointer
//https://stackoverflow.com/questions/21719769/passing-an-array-as-an-argument-in-golang 
//https://stackoverflow.com/questions/2439453/using-a-pointer-to-array 
// https://stackoverflow.com/questions/25332725/use-a-literal-for-type-in-assignment
func sum(arr []int) int {
  sum := 0
  for i := 0; i<len(arr); i++ {
    sum += arr[i]
  }
  return sum //https://gistpages.com/posts/go-lang-too-many-arguments-to-return
}
