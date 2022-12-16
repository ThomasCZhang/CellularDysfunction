package main

import (
	"fmt"
	//"io/ioutil"
	"math"
	//"os"
	//"strconv"
	//"strings"
	"testing"
)

func TestDistance(t *testing.T) {

  type test struct {
    point1, point2 OrderedPair
    answer float64
  }

  tests := make([]test, 9)
  tests[0].point1.x = 0
  tests[0].point1.y = 0
  tests[0].point2.x = 0
  tests[0].point2.y = 0
  tests[0].answer = 0.0

  tests[1].point1.x = 0
  tests[1].point1.y = 0
  tests[1].point2.x = 3
  tests[1].point2.y = 4
  tests[1].answer = 5.0

  tests[2].point1.x = 0
  tests[2].point1.y = 0
  tests[2].point2.x = -3
  tests[2].point2.y = -4
  tests[2].answer = 5

  tests[3].point1.x = 0
  tests[3].point1.y = 0
  tests[3].point2.x = 30000000
  tests[3].point2.y = 40000000
  tests[3].answer = 50000000

  tests[4].point1.x = 0
  tests[4].point1.y = 0
  tests[4].point2.x = -30000000
  tests[4].point2.y = -40000000
  tests[4].answer = 50000000

  tests[5].point1.x = 50000000
  tests[5].point1.y = 50000000
  tests[5].point2.x = 50000000
  tests[5].point2.y = 50000000
  tests[5].answer = 0.0

  tests[6].point1.x = 6
  tests[6].point1.y = 8
  tests[6].point2.x = -6
  tests[6].point2.y = -8
  tests[6].answer = 20.0

  tests[7].point1.x = 50000000
  tests[7].point1.y = 0
  tests[7].point2.x = -50000000
  tests[7].point2.y = 0
  tests[7].answer = 100000000

  tests[8].point1.x = 0
  tests[8].point1.y = -50000000
  tests[8].point2.x = 0
  tests[8].point2.y = 50000000
  tests[8].answer = 100000000

  for i, test := range tests {
  outcome := ComputeDistance(test.point1, test.point2)

    if outcome != test.answer {
      t.Errorf("Error! For input test dataset %d, your code gives %f, and the correct distance is %f", i, outcome, test.answer)
    } else {
      fmt.Println("Correct! When the points are ", test.point1.x, test.point1.y, "and",  test.point2.x, test.point2.y, "the distance is", test.answer)
    }
  }
}

func TestMagnitude(t *testing.T) {
  type test struct {
    point OrderedPair
    answer float64
  }

  tests := make([]test, 8)

  tests[0].point.x = 0
  tests[0].point.y = 0
  tests[0].answer = 0

  tests[1].point.x = -3
  tests[1].point.y = -4
  tests[1].answer = 5

  tests[2].point.x = 3
  tests[2].point.y = 4
  tests[2].answer = 5

  tests[3].point.x = 0
  tests[3].point.y = 50000000
  tests[3].answer = 50000000

  tests[4].point.x = -50000000
  tests[4].point.y = 0
  tests[4].answer = 50000000

  tests[5].point.x = 0
  tests[5].point.y = 0
  tests[5].answer = 0

  tests[6].point.x = 0
  tests[6].point.y = 0
  tests[6].answer = 0

  tests[7].point.x = 0
  tests[7].point.y = 0
  tests[7].answer = 0

  for i, test := range tests {
  outcome := test.point.Magnitude()

    if outcome != test.answer {
      t.Errorf("Error! For input test dataset %d, your code gives %f, and the correct magnitude is %f", i, outcome, test.answer)
    } else {
      fmt.Println("Correct! When the point is", test.point.x, test.point.y, "the richness is", test.answer)
    }
  }
}

func TestFindLine(t *testing.T) {

  type test struct {
    point1, point2 OrderedPair
    answer1, answer2 float64
  }

  tests := make([]test, 9)
  tests[0].point1.x = 0
  tests[0].point1.y = 0
  tests[0].point2.x = 4
  tests[0].point2.y = 0
  tests[0].answer1 = 0.0
  tests[0].answer2 = 0.0

  tests[1].point1.x = 0
  tests[1].point1.y = 5
  tests[1].point2.x = -4
  tests[1].point2.y = 5
  tests[1].answer1 = 0.0
  tests[1].answer2 = 5.0

  tests[2].point1.x = 2
  tests[2].point1.y = 0
  tests[2].point2.x = 0
  tests[2].point2.y = -4
  tests[2].answer1 = -2.0
  tests[2].answer2 = -4.0

  tests[3].point1.x = 0
  tests[3].point1.y = 0
  tests[3].point2.x = 20000000
  tests[3].point2.y = 40000000
  tests[3].answer1 = 2.0
  tests[3].answer2 = 0.0

  tests[4].point1.x = 0
  tests[4].point1.y = 0
  tests[4].point2.x = -20000000
  tests[4].point2.y = -40000000
  tests[4].answer1 = 20000000
  tests[4].answer2 = 0

  tests[5].point1.x = 50000000
  tests[5].point1.y = 50000001
  tests[5].point2.x = 50000001
  tests[5].point2.y = 50000000
  tests[5].answer1 = -1.0
  tests[5].answer2 = 0.0

  tests[6].point1.x = 6
  tests[6].point1.y = 8
  tests[6].point2.x = -6
  tests[6].point2.y = -7
  tests[6].answer1 = 1.25
  tests[6].answer2 = 0


  tests[7].point1.x = 50000000
  tests[7].point1.y = 0
  tests[7].point2.x = -50000000
  tests[7].point2.y = 0
  tests[7].answer1 = 0
  tests[7].answer2 = 0

  tests[8].point1.x = 0
  tests[8].point1.y = -50000000
  tests[8].point2.x = 0
  tests[8].point2.y = 50000000
  tests[8].answer1 = math.Inf(1)
  tests[8].answer2 = math.NaN()

  for i, test := range tests {
  outcome1, outcome2 := FindLine(test.point1, test.point2)

    if outcome1 != test.answer1 && outcome2 != test.answer2 {
      t.Errorf("Error! For input test dataset %d, your code gives (%f, %f), and the correct m and b is (%f, %f).", i, outcome1, outcome2, test.answer1, test.answer2)
    } else {
      fmt.Println("Correct! When the points are", test.point1.x, test.point1.y, "and",  test.point2.x, test.point2.y, "the line is", test.answer1, test.answer2)
    }
  }
}

func TestFindHomogeneousLine(t *testing.T) {

  type test struct {
    point1, point2 OrderedPair
    answer1, answer2, answer3 float64
  }

  tests := make([]test, 9)
  tests[0].point1.x = 0
  tests[0].point1.y = 0
  tests[0].point2.x = 4
  tests[0].point2.y = 0
  tests[0].answer1 = 0.0
  tests[0].answer2 = 0.0
  tests[0].answer3 = 0.0

  tests[1].point1.x = 0
  tests[1].point1.y = 5
  tests[1].point2.x = -4
  tests[1].point2.y = 5
  tests[1].answer1 = 0.0
  tests[1].answer2 = 5.0
  tests[1].answer3 = 0.0

  tests[2].point1.x = 2
  tests[2].point1.y = 0
  tests[2].point2.x = 0
  tests[2].point2.y = -4
  tests[2].answer1 = -4.0
  tests[2].answer2 = 2.0
  tests[2].answer3 = 8.0

  tests[3].point1.x = 0
  tests[3].point1.y = 0
  tests[3].point2.x = 20000000
  tests[3].point2.y = 40000000
  tests[3].answer1 = 2.0
  tests[3].answer2 = 0.0
  tests[3].answer3 = 0.0

  tests[4].point1.x = 0
  tests[4].point1.y = 0
  tests[4].point2.x = -20000000
  tests[4].point2.y = -40000000
  tests[4].answer1 = 20000000
  tests[4].answer2 = 0
  tests[4].answer3 = 0.0

  tests[5].point1.x = 50000000
  tests[5].point1.y = 50000001
  tests[5].point2.x = 50000001
  tests[5].point2.y = 50000000
  tests[5].answer1 = -1.0
  tests[5].answer2 = 0.0
  tests[5].answer3 = 0.0

  tests[6].point1.x = 6
  tests[6].point1.y = 8
  tests[6].point2.x = -6
  tests[6].point2.y = -7
  tests[6].answer1 = -15
  tests[6].answer2 = 12
  tests[6].answer3 = -6.0

  tests[7].point1.x = 50000000
  tests[7].point1.y = 0
  tests[7].point2.x = -50000000
  tests[7].point2.y = 0
  tests[7].answer1 = 0
  tests[7].answer2 = 0
  tests[7].answer3 = 0.0

  tests[8].point1.x = 0
  tests[8].point1.y = -50000000
  tests[8].point2.x = 0
  tests[8].point2.y = 50000000
  tests[8].answer1 = math.Inf(1)
  tests[8].answer2 = math.NaN()
  tests[8].answer3 = 0.0

  for i, test := range tests {
  outcome1, outcome2, outcome3 := FindHomogenousLine(test.point1, test.point2)

    if outcome1 != test.answer1 && outcome2 != test.answer2 && outcome3 != test.answer3 {
      t.Errorf("Error! For input test dataset %d, your code gives (%f, %f, %f), and the correct A, B, and C is (%f, %f, %f).", i, outcome1, outcome2, outcome3, test.answer1, test.answer2, test.answer3)
    } else {
      fmt.Println("Correct! When the points are", test.point1.x, test.point1.y, "and",  test.point2.x, test.point2.y, "the A, B, C values are", test.answer1, test.answer2, test.answer3)
    }
  }
}

func TestProjectVector(t *testing.T) {

  type test struct {
    point1, point2 OrderedPair
    answer1, answer2 float64
  }

  tests := make([]test, 9)
  tests[0].point1.x = 0
  tests[0].point1.y = 0
  tests[0].point2.x = 4
  tests[0].point2.y = 0
  tests[0].answer1 = 0.0
  tests[0].answer2 = 0.0

  tests[1].point1.x = 0
  tests[1].point1.y = 5
  tests[1].point2.x = -4
  tests[1].point2.y = 5
  tests[1].answer1 = -2.439024
  tests[1].answer2 = 3.048780

  tests[2].point1.x = 2
  tests[2].point1.y = 0
  tests[2].point2.x = 0
  tests[2].point2.y = -4
  tests[2].answer1 = 0
  tests[2].answer2 = 0

  tests[3].point1.x = 0
  tests[3].point1.y = 0
  tests[3].point2.x = 20000000
  tests[3].point2.y = 40000000
  tests[3].answer1 = 2.0
  tests[3].answer2 = 0.0

  tests[4].point1.x = 0
  tests[4].point1.y = 0
  tests[4].point2.x = -20000000
  tests[4].point2.y = -40000000
  tests[4].answer1 = 20000000
  tests[4].answer2 = 0

  tests[5].point1.x = 50000000
  tests[5].point1.y = 50000001
  tests[5].point2.x = 50000001
  tests[5].point2.y = 50000000
  tests[5].answer1 = 50000001
  tests[5].answer2 = 50000000

  tests[6].point1.x = 6
  tests[6].point1.y = 8
  tests[6].point2.x = -6
  tests[6].point2.y = -7
  tests[6].answer1 = 6.494118
  tests[6].answer2 = 7.576471

  tests[7].point1.x = 50000000
  tests[7].point1.y = 0
  tests[7].point2.x = -50000000
  tests[7].point2.y = 0
  tests[7].answer1 = 0
  tests[7].answer2 = 0

  tests[8].point1.x = 0
  tests[8].point1.y = -50000000
  tests[8].point2.x = 0
  tests[8].point2.y = 50000000
  tests[8].answer1 = 0
  tests[8].answer2 = -50000000

  for i, test := range tests {
    outcome := ProjectVector(test.point1, test.point2)
    outcome1 := outcome.x
    outcome2 := outcome.y

    if math.Round(outcome1/6)*6 != math.Round(test.answer1/6)*6 && math.Round(outcome2/6)*6 != math.Round(test.answer2/6)*6 {
      t.Errorf("Error! For input test dataset %d, your code gives (%f, %f), and the correct vector is (%f, %f).", i, outcome1, outcome2, test.answer1, test.answer2)
    } else {
      fmt.Println("Correct! When the points are", test.point1.x, test.point1.y, "and",  test.point2.x, test.point2.y, "the vector is", test.answer1, test.answer2)
    }
  }
}

func TestDotProduct2D(t *testing.T) {

  type test struct {
    point1, point2 OrderedPair
    answer float64
  }

  tests := make([]test, 9)
  tests[0].point1.x = 0
  tests[0].point1.y = 0
  tests[0].point2.x = 4
  tests[0].point2.y = 0
  tests[0].answer = 0.0

  tests[1].point1.x = 0
  tests[1].point1.y = 5
  tests[1].point2.x = -4
  tests[1].point2.y = 5
  tests[1].answer = 25

  tests[2].point1.x = 2
  tests[2].point1.y = 0
  tests[2].point2.x = 0
  tests[2].point2.y = -4
  tests[2].answer = 0

  tests[3].point1.x = 0
  tests[3].point1.y = 0
  tests[3].point2.x = 20000000
  tests[3].point2.y = 40000000
  tests[3].answer = 2.0

  tests[4].point1.x = 0
  tests[4].point1.y = 0
  tests[4].point2.x = -20000000
  tests[4].point2.y = -40000000
  tests[4].answer = 0

  tests[5].point1.x = 50000000
  tests[5].point1.y = 50000001
  tests[5].point2.x = 50000001
  tests[5].point2.y = 50000000
  tests[5].answer = 5000000100000000

  tests[6].point1.x = 6
  tests[6].point1.y = 8
  tests[6].point2.x = -6
  tests[6].point2.y = -7
  tests[6].answer = -92

  tests[7].point1.x = 50000000
  tests[7].point1.y = 0
  tests[7].point2.x = -50000000
  tests[7].point2.y = 0
  tests[7].answer = -2500000000000000

  tests[8].point1.x = 0
  tests[8].point1.y = -50000000
  tests[8].point2.x = 0
  tests[8].point2.y = 50000000
  tests[8].answer = -2500000000000000

  for i, test := range tests {
    outcome := DotProduct2D(test.point1, test.point2)

    if math.Round(outcome/6)*6 != math.Round(test.answer/6)*6 {
      t.Errorf("Error! For input test dataset %d, your code gives %f, and the correct dot product is %f.", i, outcome, test.answer)
    } else {
      fmt.Println("Correct! When the points are", test.point1.x, test.point1.y, "and",  test.point2.x, test.point2.y, "the dot product is", test.answer)
    }
  }
}

func TestMultiplyVectorByConstant2D(t *testing.T) {

  type test struct {
    point OrderedPair
    constant float64
    answer1, answer2 float64
  }

  tests := make([]test, 9)
  tests[0].point.x = 0
  tests[0].point.y = 0
  tests[0].constant = 5
  tests[0].answer1 = 0.0
  tests[0].answer2 = 0.0

  tests[1].point.x = 0
  tests[1].point.y = 5
  tests[1].constant = -5
  tests[1].answer1 = -2.439024
  tests[1].answer2 = 3.048780

  tests[2].point.x = 2
  tests[2].point.y = 0
  tests[2].constant = 3
  tests[2].answer1 = 0
  tests[2].answer2 = 0

  tests[3].point.x = 20000000
  tests[3].point.y = 40000000
  tests[3].constant = 4
  tests[3].answer1 = 80000000
  tests[3].answer2 = 160000000

  tests[4].point.x = -20000000
  tests[4].point.y = -40000000
  tests[4].constant = -5
  tests[4].answer1 = 100000000
  tests[4].answer2 = 200000000

  tests[5].point.x = 50000000
  tests[5].point.y = 50000001
  tests[5].constant = 0
  tests[5].answer1 = 0
  tests[5].answer2 = 0

  tests[6].point.x = 6
  tests[6].point.y = 8
  tests[6].constant = -6
  tests[6].answer1 = -36
  tests[6].answer2 = -48

  tests[7].point.x = -50000000
  tests[7].point.y = 0
  tests[7].constant = -1
  tests[7].answer1 = 50000000
  tests[7].answer2 = 0

  tests[8].point.x = 0
  tests[8].point.y = -50000000
  tests[8].constant = 100
  tests[8].answer1 = 0
  tests[8].answer2 = -5000000000

  for i, test := range tests {
    outcome := MultiplyVectorByConstant2D(test.point, test.constant)
    outcome1 := outcome.x
    outcome2 := outcome.y

    if math.Round(outcome1/6)*6 != math.Round(test.answer1/6)*6 && math.Round(outcome2/6)*6 != math.Round(test.answer2/6)*6 {
      t.Errorf("Error! For input test dataset %d, your code gives (%f, %f), and the correct vector is (%f, %f).", i, outcome1, outcome2, test.answer1, test.answer2)
    } else {
      fmt.Println("Correct! When the points is", test.point.x, test.point.y, "and the constant is", test.constant,  "the vector is", test.answer1, test.answer2)
    }
  }
}

func TestCalculateAngleBetweenVectors2D(t *testing.T) {

    type test struct {
      point1, point2 OrderedPair
      answer float64
    }

    tests := make([]test, 9)
    tests[0].point1.x = 1
    tests[0].point1.y = 1
    tests[0].point2.x = -1
    tests[0].point2.y = 1
    tests[0].answer = 1.570796

    tests[1].point1.x = 0
    tests[1].point1.y = 5
    tests[1].point2.x = -4
    tests[1].point2.y = 5
    tests[1].answer = 0.674741

    tests[2].point1.x = 2
    tests[2].point1.y = 0
    tests[2].point2.x = 0
    tests[2].point2.y = -4
    tests[2].answer = 0

    tests[3].point1.x = 30000000
    tests[3].point1.y = 40000000
    tests[3].point2.x = 30000000
    tests[3].point2.y = 40000000
    tests[3].answer = 2.0

    tests[4].point1.x = -12
    tests[4].point1.y = -16
    tests[4].point2.x = -30
    tests[4].point2.y = -40
    tests[4].answer = 0

    tests[5].point1.x = -5
    tests[5].point1.y = 10
    tests[5].point2.x = -5
    tests[5].point2.y = -10
    tests[5].answer = 2.214297

    tests[6].point1.x = 6
    tests[6].point1.y = 8
    tests[6].point2.x = -6
    tests[6].point2.y = -7
    tests[6].answer = 3.076467

    tests[7].point1.x = 50000000
    tests[7].point1.y = 0
    tests[7].point2.x = -50000000
    tests[7].point2.y = 0
    tests[7].answer = 3.141593

    tests[8].point1.x = 30000000
    tests[8].point1.y = 40000000
    tests[8].point2.x = 160000000
    tests[8].point2.y = 120000000
    tests[8].answer = 0.283794

    for i, test := range tests {
      outcome := CalculateAngleBetweenVectors2D(test.point1, test.point2)

      if math.Round(outcome/6)*6 != math.Round(test.answer/6)*6 {
        t.Errorf("Error! For input test dataset %d, your code gives %f, and the correct angle is %f.", i, outcome, test.answer)
      } else {
        fmt.Println("Correct! When the points are", test.point1.x, test.point1.y, "and",  test.point2.x, test.point2.y, "the angle is", test.answer)
      }
    }
  }
