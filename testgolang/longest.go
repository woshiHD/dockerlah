// package main
//
// import "fmt"
//
// func main() {
// 	s := []string{"flower", "flow", "flight"}
// 	longestCommonPrefix(s)
// }
//
// func longestCommonPrefix(strs []string) string {
// 	var s1 string
// 	var s2 string
// 	var a []string
// 	var b []int
// 	var prefix []string
//
// 	// for index := 0; index < len(strs); index++ {
// 	// 	x = strs[index]
// 	// 	y = append(y, x)
// 	// }
//
// 	// s = strs[0]
// 	//
// 	// for _, r := range s {
// 	// 	c := string(r)
// 	// 	a = append(a, c)
// 	// }
//
// 	for index := 0; index < len(strs); index++ {
// 		s1 = strs[index]
// 		for _, r := range s1 {
// 			c := string(r)
// 			a = append(a, c)
// 		}
// 	}
//
// 	for index := 0; index < len(strs); index++ {
// 		s2 = strs[index]
// 		b = append(b, len(s2))
// 	}
//
// 	fmt.Println(a)
// 	fmt.Println(b)
//
// 	fmt.Println(a[b[0]])
//
// 	for index := 0; index < b[0]; index++ {
// 		for index := 0; index < len(b); index++ {
// 			prefix = append(prefix, a[b[index]])
// 		}
// 	}
// 	// for {
// 	//
// 	// }
// 	fmt.Println(prefix)
// 	return "test"
// }
