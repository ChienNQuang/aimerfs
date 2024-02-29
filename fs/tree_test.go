package fs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTree(t *testing.T) {
	tree := NewTree()
	require.NotNil(t, tree)
}

func TestMkdir(t *testing.T) {
	tree := NewTree()
	err := tree.Mkdir("dir", "/")
	require.Nil(t, err)
	require.Len(t, tree.nodes, 1)
	require.Equal(t, "dir", tree.nodes[0].Name)
	require.Equal(t, "/", tree.nodes[0].Path)
	require.True(t, tree.nodes[0].IsDir)
}

func TestMkdirDuplicate(t *testing.T) {
	tree := NewTree()
	err := tree.Mkdir("dir", "/")
	require.Nil(t, err)
	err = tree.Mkdir("dir", "/")
	require.NotNil(t, err)
	require.EqualError(t, err, ErrConflictDir("dir").Error())
}

func TestMkdirNested(t *testing.T) {
	tree := NewTree()
	err := tree.Mkdir("dir", "/")
	require.Nil(t, err)
	err = tree.Mkdir("nested", "/dir")
	require.Nil(t, err)
	require.Len(t, tree.nodes, 2)
	require.Equal(t, "dir", tree.nodes[0].Name)
	require.Equal(t, "/", tree.nodes[0].Path)
	require.True(t, tree.nodes[0].IsDir)
	require.Equal(t, "nested", tree.nodes[1].Name)
	require.Equal(t, "/dir", tree.nodes[1].Path)
	require.True(t, tree.nodes[1].IsDir)
}

func TestMkdirNoParentPath(t *testing.T) {
	tree := NewTree()
	err := tree.Mkdir("nested", "/dir")
	require.NotNil(t, err)
	require.EqualError(t, err, ErrNotFound("dir").Error())
}
