package main

import (
	"os"

	"github.com/bitrise-steplib/bitrise-step-android-instrumented-test/step"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.NewLogger()
	instrumentedTestRunner := createTestRunner(logger)
	exitCode := 0

	config, err := instrumentedTestRunner.ProcessConfig()
	if err != nil {
		logger.Errorf(err.Error())
		exitCode = 1
		return exitCode
	}

	if err := instrumentedTestRunner.Run(*config); err != nil {
		logger.Errorf(err.Error())
		exitCode = 1
	}

	return exitCode
}

func createTestRunner(logger log.Logger) step.InstrumentedTestRunner {
	osEnvs := env.NewRepository()
	inputParser := stepconf.NewInputParser(osEnvs)
	commandFactory := command.NewFactory(osEnvs)
	return step.New(logger, inputParser, commandFactory)
}
