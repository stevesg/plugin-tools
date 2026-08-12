//go:build mage
// +build mage

package main

// This file adds grafana-app-sdk code generation to the standard plugin mage targets. It is a
// separate file from Magefile.go on purpose: mage collects targets from every mage-tagged file in
// the package, so your Magefile.go stays untouched and upgradeable.

import (
	"fmt"
	"strings"

	"github.com/magefile/mage/sh"
)

// appSDKModule is the Go module providing both the app-sdk library and its code generator.
const appSDKModule = "github.com/grafana/grafana-app-sdk"

// appSDKVersion resolves the app-sdk version from go.mod, so the generator always matches the
// library your generated code is compiled against. Bump it with:
//
//	go get github.com/grafana/grafana-app-sdk@latest && mage generate && go mod tidy
//
// Note the ordering: run `mage generate` before `go mod tidy`. Until generated code exists, nothing
// imports the app-sdk, so `tidy` would drop the requirement and this lookup would fail.
func appSDKVersion() (string, error) {
	version, err := sh.Output("go", "list", "-m", "-f", "\{{.Version}}", appSDKModule)
	if err != nil {
		return "", fmt.Errorf("could not resolve %s version from go.mod (is it a dependency?): %w", appSDKModule, err)
	}
	return strings.TrimSpace(version), nil
}

// Generate runs grafana-app-sdk code generation from the CUE kinds in ./kinds.
//
// Output paths are configured in kinds/config.cue: Go types to pkg/generated/, TypeScript types to
// src/generated/, and the CRD + app manifest JSON to definitions/. The generated manifest is copied
// into the plugin bundle as app-sdk-manifest.json by the frontend build.
//
// Re-run this whenever you change anything under ./kinds or upgrade the app-sdk. Generated code is
// intended to be committed.
func Generate() error {
	version, err := appSDKVersion()
	if err != nil {
		return err
	}

	return sh.RunV("go", "run", appSDKModule+"/cmd/grafana-app-sdk@"+version, "generate", "--source", "kinds")
}
