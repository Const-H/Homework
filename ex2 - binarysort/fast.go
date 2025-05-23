package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error при чтении:", err)
		return
	}

	nums := strings.Fields(string(data))
	if len(nums) < 3 {
		fmt.Println("Error:(a, b & c)")
		return
	}

	a, err1 := strconv.ParseFloat(nums[0], 64)
	b, err2 := strconv.ParseFloat(nums[1], 64)
	c, err3 := strconv.ParseFloat(nums[2], 64)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error при преобразовании:", err1, err2, err3)
		return
	}

	D := b*b - 4*a*c

	var root1, root2 string

	if D >= 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		root1 = fmt.Sprintf("%.2f", x1)
		root2 = fmt.Sprintf("%.2f", x2)
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-D) / (2 * a)
		root1 = fmt.Sprintf("%.2f+%.2fi", realPart, imaginaryPart)
		root2 = fmt.Sprintf("%.2f-%.2fi", realPart, imaginaryPart)
	}

	fmt.Println("Корни уравнения:")
	fmt.Println("x1 =", root1)
	fmt.Println("x2 =", root2)

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s\n%s", root1, root2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}