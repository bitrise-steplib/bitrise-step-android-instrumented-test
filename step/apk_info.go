package step

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/avast/apkparser"
)

type AndroidManifest struct {
	XMLName     xml.Name `xml:"manifest"`
	PackageName string   `xml:"package,attr"`
}

func getAPKPackageName(apkPath string) (string, error) {
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

	var manifest AndroidManifest
	err := xml.Unmarshal(output.Bytes(), &manifest)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %w", err)
	} else {
		return manifest.PackageName, nil
	}
}
