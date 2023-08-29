package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Packages   Packages
	PackageMap map[string]*Package
	Path       string
}

func NewRepository(path string) (*Repository, error) {
	r := &Repository{
		Path:       path,
		PackageMap: map[string]*Package{},
		Packages:   Packages{},
	}

	logrus.Debugf("Looking for packages in %s", r.Path)
	packageFiles, err := os.ReadDir(r.Path)
	if err != nil {
		return nil, err
	}

	for _, packageFile := range packageFiles {
		if !packageFile.IsDir() {
			continue
		}

		p := &Package{
			Name:       packageFile.Name(),
			Path:       filepath.Join(r.Path, packageFile.Name()),
			VersionMap: map[string]*Version{},
			Versions:   Versions{},
		}

		logrus.Debugf("Looking for versions in %s", p.Path)
		versionFiles, err := os.ReadDir(p.Path)
		if err != nil {
			return nil, err
		}

		for _, versionFile := range versionFiles {
			if !versionFile.IsDir() {
				continue
			}
			version := versionFile.Name()

			parsedVersion, err := semver.NewVersion(version)
			if err != nil {
				logrus.Warnf("Skipping unparsable version of %s: %s", p.Name, version)
				continue
			}

			v := &Version{
				Version:       version,
				ParsedVersion: parsedVersion,
				Path:          filepath.Join(p.Path, version),
				Package:       p,
			}
			p.VersionMap[version] = v
			p.Versions = append(p.Versions, v)
		}

		sort.Sort(&p.Versions)

		if len(p.VersionMap) > 0 {
			r.PackageMap[p.Name] = p
			r.Packages = append(r.Packages, p)
		}
	}
	return r, nil
}

func (r *Repository) FindPackageVersion(query string) (*Version, error) {
	var packageName string
	var packageVersion string

	if strings.Contains(query, ":") {
		queryParts := strings.SplitN(query, ":", 2)
		packageName = queryParts[0]
		packageVersion = queryParts[1]
	} else {
		packageName = query
	}

	p, ok := r.PackageMap[packageName]
	if !ok {
		return nil, fmt.Errorf("Package not found: %s", packageName)
	}

	if packageVersion != "" {
		v, ok := p.VersionMap[packageVersion]
		if !ok {
			return nil, fmt.Errorf("Package %s has no version: %s", packageName, packageVersion)
		}
		return v, nil
	}

	return p.Versions.Latest(), nil
}
