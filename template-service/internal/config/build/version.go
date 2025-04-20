package build

import (
	"fmt"
	"runtime/debug"
)

type Version struct {
	AppName   string
	Version   string
	OS        string
	Arch      string
	BuildTime string
	ShaCommit string
}

func (v *Version) String() string {
	return fmt.Sprintf(
		"AppName: %s, Version: %s, OS: %s, Arch: %s, BuildTime: %s, ShaCommit: %s",
		v.AppName, v.Version, v.OS, v.Arch, v.BuildTime, v.ShaCommit,
	)
}

func GetAppVersion() *Version {
	return &Version{
		AppName:   appName,
		Version:   version,
		OS:        os,
		Arch:      arch,
		BuildTime: buildTime,
		ShaCommit: shaCommit,
	}
}

func init() {
	readDebugInfo()
}

func readDebugInfo() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	for _, setting := range info.Settings {
		switch setting.Key {
		case "GOARCH":
			arch = setting.Value
		case "GOOS":
			os = setting.Value
		}
	}
}
