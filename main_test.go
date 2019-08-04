package bst

import (
	"testing"
)

func createBST(root int, values ...int) *binarySearchTree {
	tree := &binarySearchTree{value: root}
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

// https://en.wikipedia.org/wiki/File:Binary_search_tree.svg
var test = createBST(8, 3, 1, 6, 10, 4, 7, 14, 13)

func TestRootValue(t *testing.T) {
	output := test.value
	expected := 8
	if output != expected {
		t.Fail()
	}
}

func TestLeftValue(t *testing.T) {
	output := test.left.value
	expected := 3
	if output != expected {
		t.Fail()
	}
}

func TestLeftLeftValue(t *testing.T) {
	output := test.left.left.value
	expected := 1
	if output != expected {
		t.Fail()
	}
}

func TestLeftLeftLeftNode(t *testing.T) {
	output := test.left.left.left
	if output != nil {
		t.Fail()
	}
}

func TestLeftLeftRightNode(t *testing.T) {
	output := test.left.left.right
	if output != nil {
		t.Fail()
	}
}

func TestLeftRightValue(t *testing.T) {
	output := test.left.right.value
	expected := 6
	if output != expected {
		t.Fail()
	}
}

func TestLeftRightLeftValue(t *testing.T) {
	output := test.left.right.left.value
	expected := 4
	if output != expected {
		t.Fail()
	}
}

func TestLeftRightLeftLeft(t *testing.T) {
	output := test.left.right.left.left
	if output != nil {
		t.Fail()
	}
}

func TestLeftRightLeftRight(t *testing.T) {
	output := test.left.right.left.right
	if output != nil {
		t.Fail()
	}
}

func TestLeftRightRightValue(t *testing.T) {
	output := test.left.right.right.value
	expected := 7
	if output != expected {
		t.Fail()
	}
}

func TestLeftRightRightLeft(t *testing.T) {
	output := test.left.right.right.left
	if output != nil {
		t.Fail()
	}
}

func TestLeftRightRightRight(t *testing.T) {
	output := test.left.right.right.right
	if output != nil {
		t.Fail()
	}
}

func TestRightValue(t *testing.T) {
	output := test.right.value
	expected := 10
	if output != expected {
		t.Fail()
	}
}

func TestRightLeft(t *testing.T) {
	output := test.right.left
	if output != nil {
		t.Fail()
	}
}

func TestRightRightValue(t *testing.T) {
	output := test.right.right.value
	expected := 14
	if output != expected {
		t.Fail()
	}
}

func TestRightRightRight(t *testing.T) {
	output := test.right.right.right
	if output != nil {
		t.Fail()
	}
}

func TestRightRightLeftValue(t *testing.T) {
	output := test.right.right.left.value
	expected := 13
	if output != expected {
		t.Fail()
	}
}

func TestRightRightLeftLeft(t *testing.T) {
	output := test.right.right.left.left
	if output != nil {
		t.Fail()
	}
}

func TestRightRightLeftRight(t *testing.T) {
	output := test.right.right.left.right
	if output != nil {
		t.Fail()
	}
}

func TestContainsTrue(t *testing.T) {
	testData := []int{8, 3, 1, 6, 10, 4, 7, 14, 13}
	for _, val := range testData {
		output := test.Contains(val)
		expected := true
		if output != expected {
			t.Fail()
		}
	}
}

func TestContainsFalse(t *testing.T) {
	testData := []int{9, 18, 12, 36, 100, 84, 0, 2, 11}
	for _, val := range testData {
		output := test.Contains(val)
		expected := false
		if output != expected {
			t.Fail()
		}
	}
}

func TestRemoveValue(t *testing.T) {
	testRemoved := test
	testRemoved.Remove(8)
	if testRemoved.value != 10 {
		t.Fail()
	}
	if testRemoved.left.value != 3 {
		t.Fail()
	}
	if testRemoved.right.value != 14 {
		t.Fail()
	}
}

func TestRemoveLeft(t *testing.T) {
	testRemoved := test
	testRemoved.Remove(3)
	if testRemoved.left.value != 4 {
		t.Fail()
	}
	if testRemoved.left.left.value != 1 {
		t.Fail()
	}
	if testRemoved.left.right.value != 6 {
		t.Fail()
	}
}

func TestRemoveRight(t *testing.T) {
	testRemoved := test
	testRemoved.Remove(10)
	if testRemoved.right.value != 14 {
		t.Fail()
	}
}
