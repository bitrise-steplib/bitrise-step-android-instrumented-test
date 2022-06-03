# Android Instrument Test

[![Step changelog](https://shields.io/github/v/release/hisaac/bitrise-step-android-instrument-test?include_prereleases&label=changelog&color=blueviolet)](https://github.com/hisaac/bitrise-step-android-instrument-test/releases)

Runs Instrument tests on an existing APK

<details>
<summary>Description</summary>

Runs Instrument tests on an existing APK
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `main_apk_path` | The path to the app's main APK | required | `$BITRISE_APK_PATH` |
| `test_apk_path` | The path to the app's test APK | required | `$BITRISE_TEST_APK_PATH` |
| `test_runner_class` | The name of the test runner |  | `androidx.test.runner.AndroidJUnitRunner` |
| `additional_testing_options` | Additional options to pass to the `adb` command  Options will be passed to the command like so:  ```shell adb shell am instrument [...] -e $additional_testing_options ``` |  |  |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/hisaac/bitrise-step-android-instrument-test/pulls) and [issues](https://github.com/hisaac/bitrise-step-android-instrument-test/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
