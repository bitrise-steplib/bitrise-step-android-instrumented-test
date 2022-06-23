package apk_info

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/avast/apkparser"
)

func GetAPKPackageName(apkPath string) (string, error) {
	output := &bytes.Buffer{}
	encoder := xml.NewEncoder(output)
	encoder.Indent("", "\t")

	zipError, resourcesError, manifestError := apkparser.ParseApk(apkPath, encoder)
	if zipError != nil {
		return "", fmt.Errorf("failed to unzip the APK: %w", zipError)
	}
	if resourcesError != nil {
		return "", fmt.Errorf("failed to parse resources: %w", resourcesError)
	}
	if manifestError != nil {
		return "", fmt.Errorf("failed to parse AndroidManifest.xml: %w", manifestError)
	}

	var manifest struct {
		PackageName string `xml:"package,attr"`
	}
	if err := xml.Unmarshal(output.Bytes(), &manifest); err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	return manifest.PackageName, nil
}
