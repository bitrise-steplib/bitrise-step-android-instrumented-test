package step

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bitrise-io/go-android/v2/adbmanager"
	"github.com/bitrise-io/go-android/v2/sdk"
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-android-instrumented-test/apk_info"
	"github.com/kballard/go-shellquote"
)

type Input struct {
	MainAPKPath              string `env:"main_apk_path,required"`
	TestAPKPath              string `env:"test_apk_path,required"`
	TestRunnerClass          string `env:"test_runner_class,required"`
	AdditionalTestingOptions string `env:"additional_testing_options"`
}

type Config struct {
	MainAPKPath              string
	TestAPKPath              string
	TestRunnerClass          string
	AdditionalTestingOptions []string
}

type InstrumentedTestRunner struct {
	logger         log.Logger
	inputParser    stepconf.InputParser
	commandFactory command.Factory
}

func New(
	logger log.Logger,
	inputParser stepconf.InputParser,
	commandFactory command.Factory,
) InstrumentedTestRunner {
	return InstrumentedTestRunner{
		logger:         logger,
		inputParser:    inputParser,
		commandFactory: commandFactory,
	}
}

func (testRunner InstrumentedTestRunner) ProcessConfig() (*Config, error) {
	var input Input
	if err := testRunner.inputParser.Parse(&input); err != nil {
		return nil, err
	}
	stepconf.Print(input)

	additionalTestingOptions, err := shellquote.Split(input.AdditionalTestingOptions)
	if err != nil {
		return nil, fmt.Errorf(
			"provided additional testing options (%s) are not valid CLI parameters: %w",
			input.AdditionalTestingOptions, err,
		)
	}

	return &Config{
		MainAPKPath:              input.MainAPKPath,
		TestAPKPath:              input.TestAPKPath,
		TestRunnerClass:          input.TestRunnerClass,
		AdditionalTestingOptions: additionalTestingOptions,
	}, nil
}

func (testRunner InstrumentedTestRunner) Run(config Config) error {
	androidSDK, err := sdk.NewDefaultModel(sdk.Environment{
		AndroidHome:    sdk.NewEnvironment().AndroidHome,
		AndroidSDKRoot: sdk.NewEnvironment().AndroidSDKRoot,
	})
	if err != nil {
		return err
	}

	adb, err := adbmanager.New(androidSDK, testRunner.commandFactory)
	if err != nil {
		return err
	}
	commandOptions := &command.Opts{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	testRunner.logger.Println()
	testRunner.logger.Infof("Installing main APK")

	installMainAPKCommand := adb.InstallAPKCmd(config.MainAPKPath, commandOptions)
	testRunner.logger.Donef("$ %s", installMainAPKCommand.PrintableCommandArgs())
	if err := installMainAPKCommand.Run(); err != nil {
		return err
	}

	testRunner.logger.Println()
	testRunner.logger.Infof("Installing test APK")

	installTestAPKCommand := adb.InstallAPKCmd(config.TestAPKPath, commandOptions)
	testRunner.logger.Donef("$ %s", installTestAPKCommand.PrintableCommandArgs())
	if err := installTestAPKCommand.Run(); err != nil {
		return err
	}

	packageName, err := apk_info.GetAPKPackageName(config.TestAPKPath)
	if err != nil {
		return err
	}

	testRunner.logger.Println()
	testRunner.logger.Infof("Running tests")

	outputBuffer := &bytes.Buffer{}
	testOutputWriter := io.MultiWriter(os.Stdout, outputBuffer)

	runTestsCommand := adb.RunInstrumentedTestsCmd(
		packageName,
		config.TestRunnerClass,
		config.AdditionalTestingOptions,
		&command.Opts{
			Stdout: testOutputWriter,
			Stderr: testOutputWriter,
		},
	)

	testRunner.logger.Donef("$ %s", runTestsCommand.PrintableCommandArgs())
	if err := runTestsCommand.Run(); err != nil {
		return err
	}

	// `adb` does not return an error exit code when tests fail, so we parse the test log instead
	patterns := []string {
		"FAILURES!!!",
		"shortMsg=Process crashed.",
	}
	for _, pattern := range patterns {
		if strings.Contains(outputBuffer.String(), pattern) {
			return fmt.Errorf("test run contained at least one test failure")
		}
	}

	return nil
}
