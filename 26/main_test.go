package main

import "testing"

func Test1(t *testing.T) {

	want := 2
	list := []int{1, 1, 2}
	answer := removeDuplicates(list)

	if answer != want {
		t.Fatalf("removeDuplicates() = %v, wanted %v", answer, want)
	}

}

func Test2(t *testing.T) {
	want := 5
	list := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	answer := removeDuplicates(list)

	if answer != want {
		t.Fatalf("removeDuplicates() = %v, wanted %v", answer, want)
	}
}

func Test3(t *testing.T) {
	want := 1
	list := []int{1, 1, 1}
	answer := removeDuplicates(list)

	if answer != want {
		t.Fatalf("removeDuplicates() = %v, wanted %v", answer, want)
	}
}
