// package main
//
// import (
// 	"fmt"
// 	"strconv"
// )
//
// func main() {
// 	fmt.Println(reverse(21474836471))
// }
//
// func reverse(x int) int {
// 	var count int
// 	if x < 0 {
// 		x = Abs(x)
// 		count++
// 	}
// 	y := strconv.Itoa(x)
// 	z := []rune(y)
// 	for i, j := 0, len(z)-1; i < j; i, j = i+1, j-1 {
// 		z[i], z[j] = z[j], z[i]
// 	}
// 	a := string(z)
// 	i, err := strconv.Atoi(a)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
//   if i > 2147483647 || i < -2147483648 {
// 		return 0
// 	} else {
// 		if count == 1 {
// 			return -i
// 		} else {
// 			return i
// 		}
// 	}
// }
//
// func Abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }
