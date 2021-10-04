// package main
//
// import "fmt"
// func main() {
// 	nums := []int{2,5,5,11}
// 	var target int = 10
//   fmt.Println(twoSum(nums, target))
// }
//
// func twoSum(nums []int, target int) []int {
//   var answer []int
// 	for first_count := 0; first_count < len(nums); first_count++ {
// 		for second_count := 0; second_count < len(nums); second_count++ {
// 			var test_number int
//       if first_count != second_count {
//         test_number = nums[first_count] + nums[second_count]
//         if test_number == target {
//           answer = append(answer, first_count)
//           answer = append(answer, second_count)
//           break
//           }
//        }
// 		}
// 	}
//   return answer[0:2]
// }
