package helpers

type Packages []*Package
type Package struct {
	Name       string
	Versions   Versions
	VersionMap map[string]*Version
	Path       string
}
