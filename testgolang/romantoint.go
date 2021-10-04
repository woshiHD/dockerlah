// package main
//
// import "fmt"
//
// func main() {
// 	fmt.Println(romanToInt("MCMXCIV"))
// }
//
// func romanToInt(s string) int {
// 	var a []string
// 	var count int
//
// 	for _, r := range s {
// 		c := string(r)
// 		a = append(a, c)
// 	}
// 	fmt.Println(a)
//
// 	for index := 1; index < len(a); index++ {
// 		switch a[index] {
// 		case "M":
// 			if a[index-1] == "C" {
// 				count += 800
// 			} else {
// 				count += 1000
// 			}
// 		case "D":
// 			if a[index-1] == "C" {
// 				count += 300
// 			} else {
// 				count += 500
// 			}
// 		case "C":
// 			if a[index-1] == "X" {
// 				count += 80
// 			} else {
// 				count += 100
// 			}
// 		case "L":
// 			if a[index-1] == "X" {
// 				count += 30
// 			} else {
// 				count += 50
// 			}
// 		case "X":
// 			if a[index-1] == "I" {
// 				count += 8
// 			} else {
// 				count += 10
// 			}
// 		case "V":
// 			if a[index-1] == "I" {
// 				count += 3
// 			} else {
// 				count += 5
// 			}
// 		case "I":
// 			count += 1
// 		}
// 	}
//
// 	switch a[0] {
// 	case "M":
// 		count += 1000
// 	case "D":
// 		count += 500
// 	case "C":
// 		count += 100
// 	case "L":
// 		count += 50
// 	case "X":
// 		count += 10
// 	case "V":
// 		count += 5
// 	case "I":
// 		count += 1
// 	}
//
// 	return count
// }
