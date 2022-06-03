package main

import (
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-android-instrumented-test/step"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.NewLogger()
	step := createStep(logger)
	exitCode := 0

	config, err := step.ProcessConfig()
	if err != nil {
		logger.Errorf(err.Error())
		exitCode = 1
	}

	err = step.Run(*config)
	if err != nil {
		logger.Errorf(err.Error())
		exitCode = 1
	}

	return exitCode
}

func createStep(logger log.Logger) step.AndroidInstrumentedTestStep {
	osEnvs := env.NewRepository()
	inputParser := stepconf.NewInputParser(osEnvs)
	return step.New(logger, inputParser)
}
