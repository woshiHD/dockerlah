// package main
//
// import (
// 	"fmt"
// 	"strconv"
// )
//
// func main() {
// 	fmt.Println(isPalindrome(-121))
// }
//
// func isPalindrome(x int) bool {
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
// 	if x == i {
// 		return true
// 	} else {
// 		return false
// 	}
// }
