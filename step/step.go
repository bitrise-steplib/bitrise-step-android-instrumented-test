package step

import (
	"fmt"
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/kballard/go-shellquote"
	"os"
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

func (i InstrumentedTestRunner) ProcessConfig() (*Config, error) {
	var input Input
	err := i.inputParser.Parse(&input)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input: %w", err)
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

func (i InstrumentedTestRunner) Run(config Config) error {
	i.logger.Println()
	i.logger.Infof("Installing main APK:")
	mainAPKInstallError := installAPK(i.commandFactory, config.MainAPKPath)
	if mainAPKInstallError != nil {
		return mainAPKInstallError
	}

	i.logger.Println()
	i.logger.Infof("Installing test APK:")
	testAPKInstallError := installAPK(i.commandFactory, config.TestAPKPath)
	if testAPKInstallError != nil {
		return testAPKInstallError
	}

	packageName, getAPKPackageNameError := getAPKPackageName(config.TestAPKPath)
	if getAPKPackageNameError != nil {
		return getAPKPackageNameError
	}

	i.logger.Println()
	i.logger.Infof("Running tests:")
	instrumentTestErr := runTests(
		i.commandFactory,
		packageName,
		config.TestRunnerClass,
		config.AdditionalTestingOptions,
	)
	if instrumentTestErr != nil {
		return instrumentTestErr
	}

	return nil
}

func installAPK(commandFactory command.Factory, apkPath string) error {
	args := []string{"install", apkPath}
	return runADBCommand(commandFactory, args)
}

func runTests(
	commandFactory command.Factory,
	packageName string,
	testRunnerClass string,
	additionalTestingOptions []string,
) error {
	args := []string{
		"shell",
		"am", "instrument",
		"-w", packageName + "/" + testRunnerClass,
	}
	if len(additionalTestingOptions) > 0 {
		args = append(args, "-e")
		args = append(args, additionalTestingOptions...)
	}
	return runADBCommand(commandFactory, args)
}

func runADBCommand(commandFactory command.Factory, args []string) error {
	cmd := commandFactory.Create("adb", args, &command.Opts{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	cmdError := cmd.Run()
	if cmdError != nil {
		return fmt.Errorf(
			"command: (%s) failed, error: %w", cmd.PrintableCommandArgs(), cmdError,
		)
	} else {
		return nil
	}
}
