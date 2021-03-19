package main

import(
  "fmt"
  "sort"
  "example.com/user/aoc2020day1/day1"
)

var datafile = "testdata/day1_input.txt"

// main is a wrapper to execute the individual problems for AOC2020 Day 1
// it also executes the shared utility functions to prepare the environment
func main() {
  dataset := day1.ReadDatafile(datafile)
  var partanswer []int
  partanswer = day1.Part1(dataset)
  sort.Ints(partanswer)
  fmt.Println("Part 1:", partanswer)
  partanswer = day1.Part2(dataset)
  sort.Ints(partanswer)
  fmt.Println("Part 2:", partanswer)
}
