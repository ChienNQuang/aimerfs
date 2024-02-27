package main

import (
	"fmt"

	"github.com/ChienNQuang/aimerfs/fs"
)

func main() {
	tree := fs.NewTree()

	tree.Mkdir("abc", "/")
	tree.Mkdir("x", "/abc")
	tree.Touch("g", "/abc")
	tree.Touch("e", "/abc/g")
	// l, _ := tree.Ls("/abc")
	// time.Sleep(time.Second)

	// list(l)
	tree.ShowAll()
}

func list(nodes []*fs.Node) {
	for _, n := range nodes {
		if n.IsDir {
			fmt.Printf(fs.GetAbsPath(n.Name, n.Path) + "/\n")
		} else {
			fmt.Printf(fs.GetAbsPath(n.Name, n.Path) + "\n")
		}
	}
}
