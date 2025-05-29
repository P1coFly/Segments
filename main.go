package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type segment struct {
	Start int
	End   int
}

func minPoints(segments []segment) int {
	if len(segments) == 0 {
		return 0
	}

	sort.Slice(segments, func(i, j int) bool {
		return segments[i].End < segments[j].End
	})

	count := 1
	currentPoint := segments[0].End

	for i := 1; i < len(segments); i++ {
		// Если текущая точка не лежит в отрезке,
		// берём новую точку в правой границе seg
		if currentPoint < segments[i].Start {
			currentPoint = segments[i].End
			count++
		}
	}

	return count
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	segments := make([]segment, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &segments[i].Start, &segments[i].End)
	}

	result := minPoints(segments)
	fmt.Println(result)
}
