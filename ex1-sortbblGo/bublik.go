package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func bubble_sort(nums []int64) {
	x := len(nums)
	for m := 0; m < x; m++ {
		swapped := false
		for n := 0; n < x-1; n++ {
			if nums[n] > nums[n+1] {
				nums[n], nums[n+1] = nums[n+1], nums[n]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
func main() {
	input, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("No input")
		os.Exit(1)
	}

	input_data := make([]byte, 64)
	for {
		_, err := input.Read(input_data)
		if err == io.EOF {
			break
		}
	}

	var nums []int64
	for _, m := range strings.Split(string(input_data), " ") {
		for _, n := range strings.Split(string(m), "\r") {
			for _, num := range strings.Split(string(n), "\x00") {
				tmp, _ := strconv.ParseInt(num, 10, 64)
				if tmp != 0 {
					nums = append(nums, tmp)
				}
			}
		}
	}
	input.Close()
	output, err := os.Create("output.txt")
	defer output.Close()

	bubble_sort(nums)

	output_str := ""
	for _, m := range nums {
		output_str += fmt.Sprintf("%d ", m)
	}

	output.WriteString(output_str)

	fmt.Println(nums)

}
