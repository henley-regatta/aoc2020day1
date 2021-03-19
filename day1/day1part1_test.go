package day1

import (
  "testing"
  "fmt"
)

//Test that the function actually returns a valid pair.
func TestPart1(t *testing.T) {
  //NB: test data needs a pair of numbers that sum to the target value 2020
  //    and we pick a result that has modulus 11*2009 = 22099
  testData := []int{1,2,3,5,7,11,2009,4367}
  var r []int
  r = Part1(testData)
  if(len(r) != 3 || r[0] != 11 || r[1] != 2009 || r[2] != 22099) {
    t.Log("Part1 did not find factors correctly: ", r)
    t.Fail()
  }
}

func ExamplePart1() {
  fmt.Println(Part1([]int{1,2,3,5,7,11,2009,4367}))
  // Output: [11 2009 22099]
}

//This test is SUPPOSED to fail with no factors
func TestNoFactorPart1(t *testing.T) {
  testData := []int{1,2,3,5,7,11,2021,4096,16384}
  var r []int
  r = Part1(testData)
  if(len(r) != 3 || r[0] != 0 || r[1] != 0 || r[2] != 0) {
    t.Log("Part1 found erroneous factors when none were passed: ", r)
    t.Fail()
  }
}
