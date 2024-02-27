package fs

import (
	"fmt"
)

type Node struct {
	// absxyz
	Name string
	// asdf/asfdasfd
	Path  string
	IsDir bool
}

func (n *Node) FullPath() string {
	return fmt.Sprintf("%s/%s", n.Path, n.Name)
}
