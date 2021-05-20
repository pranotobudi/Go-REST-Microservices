package utils

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	elmt := []int{9, 8, 7, 6, 5}
	elmt = BubbleSort(elmt)
	fmt.Println("elmt: ", elmt)
	assert.NotNil(t, elmt)
	assert.EqualValues(t, 5, len(elmt))
	assert.EqualValues(t, 5, elmt[0])
	assert.EqualValues(t, 6, elmt[1])
	assert.EqualValues(t, 7, elmt[2])
	assert.EqualValues(t, 8, elmt[3])
	assert.EqualValues(t, 9, elmt[4])

}

func TestBubbleSortBestCase(t *testing.T) {
	elmt := []int{5, 6, 7, 8, 9}
	elmt = BubbleSort(elmt)
	fmt.Println("elmt: ", elmt)
	assert.NotNil(t, elmt)
	assert.EqualValues(t, 5, len(elmt))
	assert.EqualValues(t, 5, elmt[0])
	assert.EqualValues(t, 6, elmt[1])
	assert.EqualValues(t, 7, elmt[2])
	assert.EqualValues(t, 8, elmt[3])
	assert.EqualValues(t, 9, elmt[4])

}

func GetElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	elmt := GetElements(5)
	assert.NotNil(t, elmt)
	assert.EqualValues(t, 5, len(elmt))
	assert.EqualValues(t, 4, elmt[0])
	assert.EqualValues(t, 3, elmt[1])
	assert.EqualValues(t, 2, elmt[2])
	assert.EqualValues(t, 1, elmt[3])
	assert.EqualValues(t, 0, elmt[4])

}

func BenchmarkBubbleSort10(b *testing.B) {
	elmt := GetElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(elmt)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	elmt := GetElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elmt)
	}
}

func BenchmarkSort1000(b *testing.B) {
	elmt := GetElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elmt)
	}
}
func BenchmarkBubbleSort100000(b *testing.B) {
	elmt := GetElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elmt)
	}
}

func BenchmarkSort100000(b *testing.B) {
	elmt := GetElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elmt)
	}
}
