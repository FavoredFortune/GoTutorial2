# LeetCode - Golang Edition

A repo for my Go approaches to Leetcode and other problems. Each problem is listed here with a ink to original 
problem detail (leetcode or otherwise) , problem domain detail and links to both the solution and tests.

**NOTE:** please run in Terminal from within package directory with command`go test` to quikly run tests or `go 
    test -v` to see more detailed tests and results or `go test -cover` to see % coverage from tests of method

###1.  **Two Sum**
   * [Problem](https://leetcode.com/problems/two-sum/)
   * Level: Easy
   * *Problem summary:* Given an array of integers, return indices of the two numbers such that they add up to a 
    specific target. You may assume that each input would have exactly one solution, and you may not use the same element twice.
   * [Solution](./src/awesomeProject/twoSum/twosum.go)
   * [Tests](./src/awesomeProject/twoSum/twosum_test.go) 
   * Leetcode Solution Results:
        * _Runtime:_ 36 ms, faster than 32.80% of Go online submissions for Two Sum.
        * _Memory Usage:_ 3 MB, less than 100.00% of Go online 
    
###2. **Add Two Numbers**
   * [Problem](https://leetcode.com/problems/add-two-numbers/)
   * Level: Medium
   * *Problem summary:* You are given two non-empty linked lists representing two non-negative integers. The digits 
    are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as 
    a linked list. You may assume the two numbers do not contain any leading zero, except the number 0 itself.
   * [Solution](./src/awesomeProject/addTwoNumbers/addtwonum.go)
   * [Tests](./src/awesomeProject/addTwoNumbers/addtwonum_test.go) 
   * Leetcode Solution Results:
      * _Runtime:_ 16 ms, faster than 21.14% of Go online submissions for Add Two Numbers.
      * _Memory Usage:_ 5 MB, less than 51.22% of Go online submissions for Add Two Numbers.
        
###287. **Find the Duplicate Number**
   * [Problem](https://leetcode.com/problems/find-the-duplicate-number/)
   * Level: Medium
   * *Problem summary:* Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.
        **Note:**
        
       - You must not modify the array (assume the array is read only).
       - You must use only constant, O(1) extra space.
       - Your runtime complexity should be less than O(n2).
       - There is only one duplicate number in the array, but it could be repeated more than once.
   * [Solution](./src/awesomeProject/findDuplicate/findDuplicate.go)
   * [Tests](./src/awesomeProject/findDuplicate/findDuplicate_test.go)
   * LeetCode Solution Results:
      * _Runtime_: 16 ms, faster than 13.92% of Go online submissions for Find the Duplicate Number.
      * _Memory Usage:_ 4.8 MB, less than 100.00% of Go online submissions for Find the Duplicate Number.
      
###217. **Contains Duplicate**
   * [Problem](https://leetcode.com/problems/contains-duplicate/)
   * Level: Easy
   * *Problem summary:* Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.
   * [Solution](./src/awesomeProject/dupes/dupes.go)
   * [Tests](./src/awesomeProject/dupes/dupes_test.go)
   * LeetCode Solution Results:
      * _Runtime:_ 24 ms, faster than 49.80% of Go online submissions for Contains Duplicate.
      * _Memory Usage:_ 9.2 MB, less than 25.00% of Go online submissions for Contains Duplicate.
      
###26. **Remove Duplicates from Sorted Array**
   * [Problem](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)
   * Level: Easy
   * *Problem summary:* Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
                         
      Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
   * [Solution](./src/awesomeProject/removeDupesSorted/removedupes_sorted.go)
   * [Tests](./src/awesomeProject/removeDupesSorted/removedupes_sorted_test.go)
   * LeetCode Solution Results:
      * _Runtime:_ 56 ms, faster than 31.40% of Go online submissions for Remove Duplicates from Sorted Array.
      * _Memory Usage:_ 7.6 MB, less than 38.46% of Go online submissions for Remove Duplicates from Sorted Array.
      
###7. **Reverse Integer**
   * [Problem](https://leetcode.com/problems/reverse-integer/)
   * Level: Easy
   * *Problem summary:* Given a 32-bit signed integer, reverse digits of an integer.
   * [Solution](./src/awesomeProject/reverseInt/reverse_int.go)
   * [Tests](./src/awesomeProject/reverseInt/reverse_int_test.go)
   * LeetCode Solution Results:
     * _Runtime:_ 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
      * _Memory Usage:_ 2.2 MB, less than 20.00% of Go online submissions for Reverse Integer.
      
###344. **Reverse String**
   * [Problem](https://leetcode.com/problems/reverse-string/)
   * Level: Easy
   * *Problem summary:* Write a function that reverses a string. The input string is given as an array of characters char[].
                         
     Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
                         
     You may assume all the characters consist of printable ascii characters.
   * [Solution](./src/awesomeProject/reverseString/reverse_strings.go)
   * [Tests](./src/awesomeProject/reverseString/reverse_strings_test.go)
   * LeetCode Solution Results:
     * _Runtime:_ 656 ms, faster than 52.02% of Go online submissions for Reverse String.
     * _Memory Usage:_ 8.7 MB, less than 75.00% of Go online submissions for Reverse String.
