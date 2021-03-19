package day1

//Part2 implements a solution to AOC2020 Day 1 Part Two.
//From a list of integers find the THREE that sum to the magic number 2020 and
//return their product.
//Input is a (slice) list of integers, output is a 4-part int slice containing
//the three factors and their product, or all zeros if no answer found.
func Part2(numlist []int) []int {
  for i := 0; i < len(numlist); i++ {
    x := numlist[i]
    for j := 0; j < len(numlist); j++ {
      if i == j {
        continue //don't self-evaluate
      }
      y := numlist[j]
      xy := x + y
      for k := 0; k < len(numlist); k++ {
        if (k == i) || (k == j) {
          continue // no self-evaluation.
        }
        if xy + numlist[k] == 2020 {
          return []int{x,y,numlist[k], x*y*numlist[k]}
        }
      }
    }
  }
  return []int{0,0,0,0}
}
