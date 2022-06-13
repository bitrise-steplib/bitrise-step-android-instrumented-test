package apk_info

import (
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetAPKPackageName(t *testing.T) {
	// Given
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok)
	mockAPKPath := path.Dir(filename) + "/testdata/app-debug-androidTest.apk"

	expectedPackageName := "io.bitrise.sample.android.test"

	// When
	actualPackageName, err := GetAPKPackageName(mockAPKPath)

	// Then
	require.NoError(t, err)
	require.Equal(t, expectedPackageName, actualPackageName)
}
