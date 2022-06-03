package step

import (
	"fmt"
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/log"
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

type AndroidInstrumentedTestStep struct {
	logger      log.Logger
	inputParser stepconf.InputParser
}

func New(logger log.Logger, inputParser stepconf.InputParser) AndroidInstrumentedTestStep {
	return AndroidInstrumentedTestStep{
		logger:      logger,
		inputParser: inputParser,
	}
}

func (a AndroidInstrumentedTestStep) ProcessConfig() (*Config, error) {
	var input Input
	err := a.inputParser.Parse(&input)
	if err != nil {
		return nil, err
	}

	stepconf.Print(input)

	additionalTestingOptions, err := shellquote.Split(input.AdditionalTestingOptions)
	if err != nil {
		return nil, fmt.Errorf("provided additional testing options (%s) are not valid CLI parameters: %w", input.AdditionalTestingOptions, err)
	}

	return &Config{
		MainAPKPath:              input.MainAPKPath,
		TestAPKPath:              input.TestAPKPath,
		TestRunnerClass:          input.TestRunnerClass,
		AdditionalTestingOptions: additionalTestingOptions,
	}, nil
}

func (a AndroidInstrumentedTestStep) Run(config Config) error {
	a.logger.Println()
	a.logger.Infof("Running tests:")

	//adb install "$BITRISE_APK_PATH"
	//adb install "$BITRISE_TEST_APK_PATH"

	//TEST_APP_PACKAGE_NAME=$(apkanalyzer manifest application-id "$BITRISE_TEST_APK_PATH")

	//adb shell am instrument -w "$TEST_APP_PACKAGE_NAME/androidx.test.runner.AndroidJUnitRunner"

	return nil
}
