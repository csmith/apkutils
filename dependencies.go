package apkutils

import (
	"fmt"
	"strings"
)

// FlattenDependencies walks and flattens the package dependency tree, returning
// the named packages and all their transitive dependencies.
//
// Ignores version constraints and conflicts.
func FlattenDependencies(packages map[string]*PackageInfo, names ...string) (map[string]*PackageInfo, error) {
	res := make(map[string]*PackageInfo)
	queue := append([]string{}, names...)

	for len(queue) > 0 {
		if _, ok := res[queue[0]]; ok {
			// We've already got a resolution for this package, skip it.
			queue = queue[1:]
			continue
		}

		if strings.HasPrefix(queue[0], "!") {
			//Package conflict, skip it
			queue = queue[1:]
			continue
		}

		p, ok := packages[queue[0]]
		if !ok {
			return nil, fmt.Errorf("package required but not found: %s", queue[0])
		}

		queue = append(queue[1:], p.Dependencies...)
		res[p.Name] = p
	}

	return res, nil
}
