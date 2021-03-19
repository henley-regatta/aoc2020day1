package day1

import (
  "os"
  "log"
  "bufio"
  "strconv"
)

// ReadDatafile is a utility function for AOC2020 Day1 to parse a file into
// a list (slice) of integers ready for use by the part solvers.
func ReadDatafile(datafile string) []int {

  file,err := os.Open(datafile)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  numlist := []int{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    numStr := scanner.Text()
    num, _ := strconv.Atoi(numStr)
    numlist = append(numlist,num)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return numlist
}
