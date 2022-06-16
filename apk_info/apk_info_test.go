package apk_info

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetAPKPackageName(t *testing.T) {
	// Given
	mockAPKPath := "./testdata/app-debug-androidTest.apk"
	expectedPackageName := "io.bitrise.sample.android.test"

	// When
	actualPackageName, err := GetAPKPackageName(mockAPKPath)

	// Then
	require.NoError(t, err)
	require.Equal(t, expectedPackageName, actualPackageName)
}
