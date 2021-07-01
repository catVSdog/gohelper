package slice_test

import (
	"helper/slice"
	"testing"
)

type Person struct {
	children []string
}

func TestIsContainedWhenCompareSlice(t *testing.T) {
	res, err := slice.IsContained([]Person{Person{children: []string{"1"}}}, Person{children: []string{"1"}})
	if err == nil || res != false {
		t.Error("IsContainedWhenCompareSlice failed")
	}
}

func TestIsContainedWhenPass(t *testing.T) {
	res, err := slice.IsContained([]int{1, 2, 3}, 1)
	if err != nil || res != true {
		t.Error("IsContainedWhenPass failed")
	}
}

func TestIsContainedWhenTypeNotEqual(t *testing.T) {
	res, err := slice.IsContained([]int{1, 2, 3}, "1")
	if err != nil || res != false {
		t.Error("IsContainedWhenTypeNotEqual failed")
	}
}
