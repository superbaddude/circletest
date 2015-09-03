package circletest

import (
	"runtime"
	"testing"

	"github.com/ezoic/sol/ezutil"
)

func Test_GolangVersion(t *testing.T) {
	//when golang changes versions, it sometimes breaks code.  circleci does not allow
	//to specify the version of golang.  therefore, make a test that fails when the version
	//does not match the expected version
	expectedVersions := []string{"go1.5"}
	if ezutil.StrInSlice(runtime.Version(), expectedVersions) == false {
		t.Fatalf("Expected golang version: %+v, actual: %s\n", expectedVersions, runtime.Version())
	}
}
