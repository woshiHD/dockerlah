// package main
//
// import "fmt"
//
// func main() {
// 	fmt.Println(isValid("{()}"))
// }
//
// func isValid(s string) bool {
// 	var a []string
// 	var storearray []string
// 	// var x string
// 	var return_var bool
//
// 	for _, r := range s {
// 		c := string(r)
// 		a = append(a, c)
// 	}
//
// 	for index := 0; index < len(a); index++ {
// 		switch a[index] {
// 		case "(":
// 			RemoveIndex(a, index)
// 			for index := 0; index < len(a); index++ {
// 				if a[index] == ")" {
//
// 				}
// 			}
// 		case "[":
//
//
// 		case "{":
//
// 		}
// 	}
//
// 	// switch a[index] {
// 	// case "(":
// 	// 	storearray = append(storearray, "(")
// 	// 	// storearray = append(storearray, ")")
// 	// case "{":
// 	// 	storearray = append(storearray, "{")
// 	// 	// storearray = append(storearray, "}")
// 	// case "[":
// 	// 	storearray = append(storearray, "[")
// 	// 	// storearray = append(storearray, "]")
// 	// case ")":
// 	// 	for index := 0; index < len(storearray); index++ {
// 	// 		if storearray[index] == "(" {
// 	// 			RemoveIndex(storearray, index)
// 	// 			fmt.Println(storearray)
// 	// 		} else {
// 	// 			continue
// 	// 		}
// 	// 	}
// 	// case "}":
// 	// 	for index := 0; index < len(storearray); index++ {
// 	// 		if storearray[index] == "{" {
// 	// 			RemoveIndex(storearray, index)
// 	// 			fmt.Println(storearray)
// 	// 		} else {
// 	// 			continue
// 	// 		}
// 	// 	}
// 	// case "]":
// 	// 	for index := 0; index < len(storearray); index++ {
// 	// 		if storearray[index] == "[" {
// 	// 			RemoveIndex(storearray, index)
// 	// 		} else {
// 	// 			continue
// 	// 		}
// 	// 	}
// 	// }
//
// func RemoveIndex(s []string, index int) []string {
// 	copy(s[index:], s[index+1:]) // Shift a[i+1:] left one index.
// 	s[len(s)-1] = ""             // Erase last element (write zero value).
// 	s = s[:len(s)-1]             // Truncate slice.
// 	return s
// }
