package version

import (
	"runtime"
	"fmt"
	"html/template"
	"io"
)

var (
	version   string
	buildDate string
	gitCommit string
)

type Version struct {
	Version   string `json:"version"`
	GitCommit string `json:"commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

func Get() Version {
	return Version{
		Version:   version,
		GitCommit: gitCommit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

var versionTemplate = ` Version:		{{.Version}}
	Git commit:		{{.GitCommit}}
	Go version:		{{.GoVersion}}
	Built:			{{.BuildDate}}
	Platform:		{{.Platform}}
`

func FormatVersion(w io.Writer) error {
	tmpl, _ := template.New("version").Parse(versionTemplate)
	return tmpl.Execute(w, Get())
}
