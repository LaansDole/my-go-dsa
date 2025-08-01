package array

import (
	"slices"
	"strconv"
	"testing"
)

func addSliceOfTwoNumbers(slice1, slice2 []int) (resultSlice []int) {
    convSliceToInt := func(slice []int) int {
        var numStr string
        for _, num := range slice {
            numStr += strconv.Itoa(num)
        }
        
        result, _ := strconv.Atoi(numStr)
        return result
    }
	result := convSliceToInt(slice1) + convSliceToInt(slice2)

    for _, digit := range strconv.Itoa(result) {
        digitInt, _ := strconv.Atoi(string(digit))
        resultSlice = append(resultSlice, digitInt)
    }
    
    return resultSlice
}

/*
TestAddSliceOfTwoNumbers tests solution(s) with the following signature and problem description:

	AddTwoNumbers(num1, num2 []int) []int

A slice representation of a positive integer like 283 looks like {2,8,3}. Given two positive
integers represented in this format return their sum in the same format.

For example given {2,9} and {9,9,9}, return {1,0,2,8}.
Because 29+999=1028.
*/
func TestAddSliceOfTwoNumbers(t *testing.T) {
	tests := []struct {
		num1, num2, sum []int
	}{
		{[]int{1}, []int{}, []int{1}},
		{[]int{1}, []int{0}, []int{1}},
		{[]int{1}, []int{1}, []int{2}},
		{[]int{1}, []int{9}, []int{1, 0}},
		{[]int{2, 5}, []int{3, 5}, []int{6, 0}},
		{[]int{2, 9}, []int{9, 9, 9}, []int{1, 0, 2, 8}},
		{[]int{9, 9, 9}, []int{9, 9, 9}, []int{1, 9, 9, 8}},
	}
	for i, test := range tests {
		if got := addSliceOfTwoNumbers(test.num1, test.num2); !slices.Equal(got, test.sum) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sum, got)
		}
	}
}
