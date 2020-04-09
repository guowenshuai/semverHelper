/*  Package semverHelper
 *
 * Given a version number MAJOR.MINOR.PATCH, increment the:
 *
 *  MAJOR version when you make incompatible API changes,
 *  MINOR version when you add functionality in a backwards compatible manner, and
 *  PATCH version when you make backwards compatible bug fixes.
 */

package semverHelper

import (
	"strconv"
	"strings"

	"golang.org/x/mod/semver"
)

func NewSemVersion() string {
	return semver.Canonical("v1.0.0")
}

func MajorPlus(v string) (string, error) {
	if !semver.IsValid(v) {
		return NewSemVersion(), nil
	}
	return plus(v, x)
}
func MinorPlus(v string) (string, error) {
	if !semver.IsValid(v) {
		return NewSemVersion(), nil
	}
	return plus(v, y)
}
func PatchPlus(v string) (string, error) {
	if !semver.IsValid(v) {
		return NewSemVersion(), nil
	}
	return plus(v, z)
}

const (
	x = iota
	y
	z
)

func plus(v string, n int) (string, error) {
	suffix := semver.Build(v)
	obj := semver.Canonical(v)
	objs := strings.Split(obj[1:], ".")
	i, err := strconv.Atoi(objs[n])
	if err != nil {
		return "", err
	}
	objs[n] = strconv.Itoa(i + 1)
	return strings.Join(objs, ".") + suffix, nil
}
