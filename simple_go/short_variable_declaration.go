package simplego

import "fmt"

func main() {
	f, i := 1, 2
	fmt.Println(f, i)
	f, i := 3, 4 // Error
}
