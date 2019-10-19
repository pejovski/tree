package tree

import (
	"fmt"
)

type Category struct {
	Name   string
	Parent string
}

type Tree []Category

func (t Tree) getChildren(parentName string) ([]string, error) {
	var children []string
	for _, c := range t {
		if c.Parent == parentName {
			children = append(children, c.Name)
		}
	}

	return children, nil
}

func (t *Tree) addCategory(name string, parent string) error {
	if hasCategory, hasParent := t.checkCategoryAndParent(name, parent); hasCategory || !hasParent {
		if hasCategory {
			return fmt.Errorf("invalid argument: Category already exist")
		}
		return fmt.Errorf("invalid argument: Parent does not exist")
	}

	*t = append(*t, Category{
		Name:   name,
		Parent: parent,
	})

	return nil
}

func (t Tree) checkCategoryAndParent(category string, parent string) (hasCategory bool, hasParent bool) {
	if parent == "" {
		hasParent = true
	}

	for _, c := range t {

		if !hasCategory && c.Name == category {
			hasCategory = true
		}

		if !hasParent && c.Name == parent {
			hasParent = true
		}

		// if both are already true there is no need to continue
		if hasCategory && hasParent {
			break
		}
	}

	return hasCategory, hasParent
}
