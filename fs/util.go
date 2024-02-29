package fs

import "fmt"

func GetAbsPath(name, path string) string {
	if path == "/" {
		return fmt.Sprintf("%s%s", path, name)
	} else {
		return fmt.Sprintf("%s/%s", path, name)
	}
}
