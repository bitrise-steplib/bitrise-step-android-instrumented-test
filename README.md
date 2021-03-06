# Android Instrumented Test

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/bitrise-step-android-instrumented-test?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/bitrise-step-android-instrumented-test/releases)

Runs Instrumented tests on an existing APK

<details>
<summary>Description</summary>

Runs Instrumented tests on an existing APK
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
| `test_runner_class` | The name of the test runner | required | `androidx.test.runner.AndroidJUnitRunner` |
| `additional_testing_options` | A space-delimited list of additional options to pass to the test runner  Example:  If a value of `KEY1 true KEY2 false` is passed to this input, then it will be passed to the `adb` command like so:  ```shell adb shell am instrument -e "KEY1" "true" "KEY2" "false" [...] ```  See the [`adb` documentation](https://developer.android.com/studio/command-line/adb#am) for more info. |  |  |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/bitrise-step-android-instrumented-test/pulls) and [issues](https://github.com/bitrise-steplib/bitrise-step-android-instrumented-test/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
