# Instructions

* To run tests:
    * copy repo locally 
    *`cd` into package of interest (example: `cd bucketSort`)
    * run `go test -v` for verbose running of tests with any system out print lines, etc.

# LeetCode - Golang Edition

A repo for my Go approaches to Leetcode and other problems. Each problem is listed here with a link to original 
problem detail (leetcode or otherwise) , problem domain detail and links to both the solution and tests.

**NOTE:** please run in Terminal from within package directory with command`go test` to quickly run tests or `go 
    test -v` to see more detailed tests and results or `go test -cover` to see % coverage from tests of method

### 1.  **Two Sum** ###
   * [Problem](https://leetcode.com/problems/two-sum/)
   * Level: Easy
   * *Problem summary:* Given an array of integers, return indices of the two numbers such that they add up to a 
    specific target. You may assume that each input would have exactly one solution, and you may not use the same element twice.
   * [Solution](./src/awesomeProject/twoSum/twosum.go)
   * [Tests](./src/awesomeProject/twoSum/twosum_test.go) 
   * Leetcode Solution Results:
        * _Runtime:_ 36 ms, faster than 32.80% of Go online submissions for Two Sum.
        * _Memory Usage:_ 3 MB, less than 100.00% of Go online 
    
### 2. **Add Two Numbers** ###
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
        
### 287. **Find the Duplicate Number** ###
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

      
### 217. **Contains Duplicate** ###
   * [Problem](https://leetcode.com/problems/contains-duplicate/)
   * Level: Easy
   * *Problem summary:* Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.
   * [Solution](./src/awesomeProject/dupes/dupes.go)
   * [Tests](./src/awesomeProject/dupes/dupes_test.go)
   * LeetCode Solution Results:
      * _Runtime:_ 24 ms, faster than 49.80% of Go online submissions for Contains Duplicate.
      * _Memory Usage:_ 9.2 MB, less than 25.00% of Go online submissions for Contains Duplicate.
      
### 26. **Remove Duplicates from Sorted Array** ###
   * [Problem](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)
   * Level: Easy
   * *Problem summary:* Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
                         
      Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
   * [Solution](./src/awesomeProject/removeDupesSorted/removedupes_sorted.go)
   * [Tests](./src/awesomeProject/removeDupesSorted/removedupes_sorted_test.go)
   * LeetCode Solution Results:
      * _Runtime:_ 56 ms, faster than 31.40% of Go online submissions for Remove Duplicates from Sorted Array.
      * _Memory Usage:_ 7.6 MB, less than 38.46% of Go online submissions for Remove Duplicates from Sorted Array.
      

### 7. **Reverse Integer** ###
   * [Problem](https://leetcode.com/problems/reverse-integer/)
   * Level: Easy
   * *Problem summary:* Given a 32-bit signed integer, reverse digits of an integer.
   * [Solution](./src/awesomeProject/reverseInt/reverse_int.go)
   * [Tests](./src/awesomeProject/reverseInt/reverse_int_test.go)
   * LeetCode Solution Results:
     * _Runtime:_ 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
      * _Memory Usage:_ 2.2 MB, less than 20.00% of Go online submissions for Reverse Integer.
      
### 344. **Reverse String** ###
   * [Problem](https://leetcode.com/problems/reverse-string/)
   * Level: Easy
   * *Problem summary:* Write a function that reverses a string. The input string is given as an array of characters char[].
     Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory
     You may assume all the characters consist of printable ascii characters.
   * [Solution](./src/awesomeProject/reverseString/reverse_strings.go)
   * [Tests](./src/awesomeProject/reverseString/reverse_strings_test.go)
   * LeetCode Solution Results:
     * _Runtime:_ 656 ms, faster than 52.02% of Go online submissions for Reverse String.
     * _Memory Usage:_ 8.7 MB, less than 75.00% of Go online submissions for Reverse String.
     
### 242. **Valid Anagram** ###
   * [Problem](https://leetcode.com/problems/valid-anagram/)
   * Level: Easy
   * *Problem summary:* Given two strings s and t , write a function to determine if t is an anagram of s. What if 
   the inputs contain unicode characters? How would you adapt your solution to such case?
    
     **NOTE** Solution is for unicode characters
   * [Solution](./src/awesomeProject/validAnagram/valid_anagram.go)
   * [Tests](./src/awesomeProject/validAnagram/valid_anagram_test.go)
   * LeetCode Solution Results - nothing to crow about

### Sort List ###
   * [Problem](https://leetcode.com/problems/sort-list/)
   * Level: Medium

# Other Data Structure & Algorithm Problem Solving

### Binary Search
* Chose to create various functions related to binary search including:
    * Verifying a slice/array is a binary sorted 
structure (boolean)
    * Finding the index of an element in a slice of a binary sorted slice (int)
    * Finding the index value of where a new value would belong in a binary sorted slice (int)
    * Inserting a value into a sorted slice in the correct location
    
     * _To Be Completed_ sort a slice

* [Solution](./src/awesomeProject/binarySearch/binary_search.go)
* [Tests](./src/awesomeProject/binarySearch/binary_search_test.go)

### Bucket Sort
* Receiving a slice with only positive integers from 0 - 99, return a sorted slice. Duplicates are allowed.
* [Solution](./src/awesomeProject/bucketSort/bucket_sort.go)
* [Tests](./src/awesomeProject/bucketSort/bucket_sort_test.go)
   
   
### Product Array
* Given an array of integers return an array of products of all elements in that array, excluding the value at that 
index.
* [Solution](./src/awesomeProject/arrayProduct/array_product.go) // not all test cases passing yet
* [Tests](./src/awesomeProject/arrayProduct/array_product_test.go)// not all test cases passing yet
* _To Be Completed_