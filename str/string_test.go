package str_test

import (
	"helper/str"
	"testing"
)

func TestIsAlphanumericWhenPass(t *testing.T) {
	a := "abc123"
	if !str.IsAlphanumeric(a) {
		t.Error("IsAlphanumericWhenPass failed")
	}
}

func TestIsAlphanumericWhenChinese(t *testing.T) {
	a := "你好"
	if !str.IsAlphanumeric(a) {
		t.Error("IsAlphanumericWhenPass failed")
	}
}

func TestIsAlphanumericFailed(t *testing.T) {
	a := "]d3!"
	if str.IsAlphanumeric(a) {
		t.Error("IsAlphanumericFailed failed")
	}
}

func TestIsUpperWhenPass(t *testing.T) {
	a := "A"
	if !str.IsUpper(a) {
		t.Error("IsUpperWhenPass failed")
	}
}

func TestIsUpperWhenFailed(t *testing.T) {
	a := "a"
	if str.IsUpper(a) {
		t.Error("IsUpperWhenFailed failed")
	}
}

func TestIsLowerWhenPass(t *testing.T) {
	a := "a"
	if !str.IsLower(a) {
		t.Error("IsLowerWhenPass failed")
	}
}

func TestIsLowerWhenFailed(t *testing.T) {
	a := "A"
	if str.IsLower(a) {
		t.Error("IsLowerWhenFailed failed")
	}
}
