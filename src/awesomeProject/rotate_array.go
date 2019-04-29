package main

import "fmt"
func main(){

   nums := []int{1,2,3,4,5,6,7}
   k := 3
   fmt.Println(rotate(nums, k))
   
}


func rotate(nums []int, k int) []int{

   for i := 0; i < k ; i ++ {
       nums[i] = nums [i + 1]
   }

return nums

}

