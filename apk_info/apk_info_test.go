package apk_info

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/bitrise-io/go-utils/command/git"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/stretchr/testify/require"
)

func Test_GetAPKPackageName(t *testing.T) {
	// Setup
	logger := log.NewLogger()

	logger.Infof("Creating temporary directory")
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		logger.Errorf("Setup: Failed to create temporary directory, error: %w", err)
	}
	logger.Infof("Temporary directory created at %s", tmpDir)

	defer func(path string, logger log.Logger) {
		if err := os.RemoveAll(path); err != nil {
			logger.Warnf("failed to remove temporary directory, error: %w", err)
		} else {
			logger.Donef("Temporary directory (%s) deleted", tmpDir)
		}
	}(tmpDir, logger)

	gitCommand, err := git.New(tmpDir)
	if err != nil {
		logger.Errorf("Setup: Failed to create directory for test artifact git project, error: %w", err)
	}

	logger.Infof("Cloning sample artifacts git repo")
	if err := gitCommand.Clone("https://github.com/bitrise-io/sample-artifacts.git").Run(); err != nil {
		logger.Errorf("Setup: Failed to clone test artifact repo, error: %w", err)
	}

	logger.Infof("Setup complete, running test")

	// Given
	mockAPKPath := path.Join(tmpDir, "apks/app-debug.apk")
	expectedPackageName := "com.example.birmachera.myapplication"

	// When
	actualPackageName, err := GetAPKPackageName(mockAPKPath)

	// Then
	require.NoError(t, err)
	require.Equal(t, expectedPackageName, actualPackageName)
}
