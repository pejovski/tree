package tree

import (
	"testing"
)

func TestGetChildren(t *testing.T) {
	tree := generateTree()

	children, err := tree.getChildren("name1")
	if err != nil {
		t.Errorf("Error getChildren: %s", err)
	}

	if len(children) != 2 {
		t.Error("Wrong number of children")
	}

	if children[0] != "name2" {
		t.Error("Wrong name of the first child")
	}
}

func TestAddCategory(t *testing.T) {
	tree := generateTree()

	err := tree.addCategory("name1", "name1")
	if err == nil {
		t.Errorf("addCategory error for category that already exist was expected")
	}

	err = tree.addCategory("name7", "name8")
	if err == nil {
		t.Errorf("addCategory error for parent that not exist was expected")
	}

	err = tree.addCategory("name5", "name1")
	if err != nil {
		t.Errorf("addCategory error was not expected: %s", err)
	}

	err = tree.addCategory("name6", "")
	if err != nil {
		t.Errorf("addCategory error was not expected: %s", err)
	}
}

func TestAddThanGetCategory(t *testing.T) {
	tree := &Tree{}

	err := tree.addCategory("name1", "")
	if err != nil {
		t.Errorf("addCategory error was not expected: %s", err)
	}

	err = tree.addCategory("name2", "name1")
	if err != nil {
		t.Errorf("addCategory error was not expected: %s", err)
	}

	err = tree.addCategory("name3", "name1")
	if err != nil {
		t.Errorf("addCategory error was not expected: %s", err)
	}

	children, err := tree.getChildren("name1")
	if err != nil {
		t.Errorf("Error getChildren: %s", err)
	}

	if len(children) != 2 {
		t.Error("Wrong number of children")
	}

	if children[0] != "name2" {
		t.Error("Wrong name of the first child")
	}

	if children[1] != "name3" {
		t.Error("Wrong name of the first child")
	}
}

func generateTree() *Tree {
	return &Tree{
		Category{
			Name:   "name1",
			Parent: "",
		},
		Category{
			Name:   "name2",
			Parent: "name1",
		},
		Category{
			Name:   "name3",
			Parent: "name1",
		},
		Category{
			Name:   "name4",
			Parent: "name2",
		},
	}
}
