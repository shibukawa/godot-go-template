//+build mage

// This is the build script for Mage. The install target is all you really need.
// The release target is for generating official releases and is really only
// useful to project admins.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type BuildPlatform struct {
	OS   string
	Arch string
}

var (
	godotBin   string
	ci         bool
	targetOS   string
	targetArch string
)

func init() {
	var (
		ok bool
	)

	if targetOS, ok = os.LookupEnv("TARGET_OS"); !ok {
		targetOS = runtime.GOOS
	}

	if targetArch, ok = os.LookupEnv("TARGET_ARCH"); !ok {
		targetArch = runtime.GOARCH
	}

	envCI, _ := os.LookupEnv("CI")
	ci = envCI == "true"
}

func envWithPlatform(platform BuildPlatform) map[string]string {
	envs := map[string]string{
		"GOOS":        targetOS,
		"GOARCH":      targetArch,
		"CGO_ENABLED": "1",
	}

	// enable for cross-compiling from linux
	// case "windows":
	// 	envs["CC"] = "i686-w64-mingw32-gcc"
	// }

	return envs
}

func Build() error {
	appPath := filepath.Join("entrypoint")
	outputPath := filepath.Join("dist")

	return buildGodotPlugin(
		appPath,
		outputPath,
		BuildPlatform{
			OS:   targetOS,
			Arch: targetArch,
		},
	)
}

func buildGodotPlugin(appPath string, outputPath string, platform BuildPlatform) error {
	return sh.RunWith(envWithPlatform(platform), mg.GoCmd(), "build",
		"-tags", "tools", "-buildmode=c-shared", "-x", "-trimpath",
		"-o", filepath.Join(outputPath, platform.godotPluginCSharedName(appPath)),
		filepath.Join(appPath, "main.go"),
	)
}

func (p BuildPlatform) godotPluginCSharedName(appPath string) string {
	// NOTE: these files needs to line up with CI as well as the naming convention
	//       expected by the test godot project
	switch p.OS {
	case "windows":
		return fmt.Sprintf("libgodotgo-myscript-windows-4.0-%s.dll", p.Arch)
	case "darwin":
		return fmt.Sprintf("libgodotgo-myscript-darwin-10.6-%s.dylib", p.Arch)
	case "linux":
		return fmt.Sprintf("libgodotgo-myscript-linux-%s.so", p.Arch)
	default:
		panic(fmt.Errorf("unsupported build platform: %s", p.OS))
	}
}

var Default = Build
