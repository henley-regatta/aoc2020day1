package day1

import (
  "testing"
  "fmt"
)

//Test that the function actually returns a valid pair.
func TestPart2(t *testing.T) {
  //NB: test data needs a triplet of numbers that sum to the target value 2020
  //    and we pick a result that has modulus 4*7*2009 = 56252
  testData := []int{1,2,3,4,5,7,2009,4367}
  var r []int
  r = Part2(testData)
  if(len(r) != 4 || r[0] != 4 || r[1] != 7 || r[2] != 2009 || r[3] != 56252) {
    t.Log("Part2 did not find factors correctly: ", r)
    t.Fail()
  }
}

func ExamplePart2() {
  fmt.Println(Part2([]int{1,2,3,4,5,7,2009,4367}))
  // Output: [4 7 2009 56252]
}


//This test is SUPPOSED to fail with no factors
func TestNoFactorPart2(t *testing.T) {
  testData := []int{1,2,3,5,7,11,2021,4096,16384}
  var r []int
  r = Part2(testData)
  if(len(r) != 4 || r[0] != 0 || r[1] != 0 || r[2] != 0 || r[3] != 0) {
    t.Log("Part2 found erroneous factors when none were passed: ", r)
    t.Fail()
  }
}
