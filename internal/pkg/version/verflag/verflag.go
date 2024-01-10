package verflag

import (
	"fmt"
	"github.com/marmotedu/miniblog/internal/pkg/version"
	"github.com/spf13/pflag"
	"os"
	"strconv"
)

type versionValue int

const (
	VersionFalse versionValue = 0
	VersionTrue  versionValue = 1
	VersionRaw   versionValue = 2
)

const (
	strRawVersion   = "raw"
	versionFlagName = "version"
)

var versionFlag = Version(versionFlagName, VersionFalse, "Print version information and quit.")

// String 实现了 pflag.Value 接口中的String方法
func (v *versionValue) String() string {
	//TODO implement me
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}

// Set 实现了 pflag.Value 接口中的Set方法
func (v *versionValue) Set(s string) error {
	//TODO implement me
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	parseBool, err := strconv.ParseBool(s)
	if parseBool {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}

func (v *versionValue) Type() string {
	return "version"
}

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func (v *versionValue) Get() interface{} {
	return v
}

func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)
	return p
}

func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	pflag.Var(p, name, usage)
	// `--version` 等价于 `--version=true`
	pflag.Lookup(name).NoOptDefVal = "true"
}

// AddFlags 在任意FlagSet上注册这个包的标志，这样他们指向与全局标志相同的值
func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

// PrintAndExitIfRequest 将检查是否传递了 `--version`标志，如果是，则打印版本并退出
func PrintAndExitIfRequest() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
