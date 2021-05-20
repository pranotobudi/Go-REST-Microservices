package utils

import "sort"

func BubbleSort(elements []int) []int {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}

	}
	return elements
}

func Sort(elmt []int) []int {
	if len(elmt) < 1000 {
		return BubbleSort(elmt)
	}
	sort.Ints(elmt)
	return elmt
}
