//Package day1 contains code to implement solutions to the AOC 2020 Day 1 problem
package day1


//Part1 implements a solution to AOC 2020 Day 1 Part One requirement:
//when passed a list of numbers, find the two that sum to the magic number
//2020 and return the product of those numbers.
//Takes a (slice) list of integers as input, returns a 3-part slice of ints
//containing either the factors and their products or all zeros if nothing found.
func Part1(numlist []int) []int {
  for i := 0; i < len(numlist); i++ {
    x := numlist[i]
    for j := 0; j < len(numlist); j++ {
      if i == j {
        continue /*don't self-evaluate*/
      }
      if x + numlist[j] == 2020 {
          return []int{x, numlist[j], x*numlist[j]}
      }
    }
  }
  return []int{0,0,0}
}
