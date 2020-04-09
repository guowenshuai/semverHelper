package semverHelper

import (
	"testing"
)

func Test_semver(t *testing.T) {
	v := NewSemVersion()
	v = "v1.10.11+alpha"
	t.Log(MajorPlus(v))
	t.Log(MinorPlus(v))
	t.Log(PatchPlus(v))
}