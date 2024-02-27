package fs

import (
	"errors"
	"fmt"
	"strings"
)

type Tree struct {
	nodes []*Node
}

func NewTree() *Tree {
	return &Tree{
		nodes: []*Node{},
	}
}

func (t *Tree) Touch(name, path string) error {
	if !strings.HasPrefix(path, "/") || strings.HasSuffix(name, "/") {
		return errors.New("invalid path")
	}

	if strings.HasSuffix(name, "/") {
		return errors.New("invalid name")
	}

	if name == "" || path == "" {
		return errors.New("stupid")
	}
	return t.addNode(name, path)
}

func (t *Tree) Mkdir(name, path string) error {
	if !strings.HasPrefix(path, "/") || strings.HasSuffix(name, "/") {
		return errors.New("invalid path")
	}

	if !strings.HasSuffix(name, "/") {
		name += "/"
	}
	if name == "" || path == "" {
		return errors.New("stupid")
	}
	return t.addNode(name, path)
}

// addNode adds a node to the tree
func (t *Tree) addNode(name, path string) error {
	isDir := strings.HasSuffix(name, "/")
	name = strings.TrimSuffix(name, "/")

	// check if path exists
	if path != "/" {
		slashCount := strings.Count(path, "/")
		var parentPath string

		if slashCount <= 1 { // if the path is at root
			parentPath = "/"
		} else { // if the path is at somewhere else
			parentPath = path[:strings.LastIndex(path, "/")]
		}
		parentName := path[strings.LastIndex(path, "/")+1:]
		var p *Node
		for _, n := range t.nodes {
			if n.Path == parentPath && n.Name == parentName {
				p = n
			}
		}

		if p == nil {
			return fmt.Errorf("%s not found", path)
		}
	}

	// handle duplication
	for _, n := range t.nodes {
		if n.Path == path && n.Name == name && n.IsDir == isDir {
			if isDir {
				return fmt.Errorf("directory %s already exists", name)
			} else {
				return fmt.Errorf("file %s already exists", name)
			}
		}
	}

	n := &Node{
		Name:  name,
		Path:  path,
		IsDir: isDir,
	}
	t.nodes = append(t.nodes, n)
	return nil
}

func (t *Tree) Ls(path string) (items []string, err error) {
	list := []string{}
	for _, n := range t.nodes {
		if n.Path == path {
			if n.IsDir {
				list = append(list, n.Name+"/")
			} else {
				list = append(list, n.Name)
			}
		}
	}

	return list, nil
}
