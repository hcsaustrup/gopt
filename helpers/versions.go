package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/coreos/go-semver/semver"
)

type Versions []*Version

func (v *Versions) Len() int {
	return len(*v)
}

func (v *Versions) Less(i int, j int) bool {
	return (*v)[i].ParsedVersion.LessThan((*(*v)[j].ParsedVersion))
}

func (v Versions) Swap(i int, j int) {
	x := v[i]
	v[i] = v[j]
	v[j] = x
}

func (v *Versions) Latest() *Version {
	return (*v)[len(*v)-1]
}

type Version struct {
	Version       string
	ParsedVersion *semver.Version
	Path          string
	Package       *Package
}

func (v *Version) AddToPath(path string, prepend bool) string {
	pathElements := []string{}

	for _, pathElement := range strings.Split(path, ":") {
		prefix := fmt.Sprintf("%s%c", v.Package.Path, os.PathSeparator)
		if strings.HasPrefix(pathElement, prefix) {
			continue
		}
		pathElements = append(pathElements, pathElement)
	}

	if prepend {
		pathElements = append([]string{v.Path}, pathElements...)
	} else {
		pathElements = append(pathElements, v.Path)
	}

	return strings.Join(pathElements, ":")
}
