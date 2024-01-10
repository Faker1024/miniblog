package version

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"runtime"
)

var (
	// GitVersion 是语义化的版本号
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间， $(date -u + '%Y-%m-%dT%H:%M:%SZ') 命令输出
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit 是Git的SHA1值，$(git rev-parse HEAD) 命令的输出
	GitCommit = "$Format:%H$"
	// GitTreeState 代表构建时Git仓库的状态，可能的值：clean，dirty
	GitTreeState = ""
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	s, err := info.Text()
	if err == nil {
		return string(s)
	}
	return info.GoVersion
}

func (info Info) ToJson() string {
	marshal, _ := json.Marshal(info)
	return string(marshal)
}

func (info Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", info.GitVersion)
	table.AddRow("gitCommit:", info.GitCommit)
	table.AddRow("gitTreeState:", info.GitTreeState)
	table.AddRow("buildDate:", info.BuildDate)
	table.AddRow("goVersion:", info.GoVersion)
	table.AddRow("compiler:", info.Compiler)
	table.AddRow("platform:", info.Platform)
	return table.Bytes(), nil
}

func Get() Info {
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
