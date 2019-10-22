package nginx

import (
	"io/ioutil"
	"taylor/utils"
	"os"
	"runtime"
)

const (
	CheckConfigDirect  = " -t"
	StopInstanceDirect = " -s stop"
	QuitInstanceDirect = " -s quit"
	ReloadConfigDirect = " -s reload"

	BinaryPathVariable        = "NginxBinaryPath"
	HttpConfigDirVariable     = "NginxHttpDir"
	StreamConfigDirVariable   = "NginxStreamDir"
	CoreConfigDirVariable     = "NginxCoreConfigPath"
	WafConfigDirVariable      = "NginxWafConfigDir"
	WafWhiteConfigDirVariable = "NginxWafWhiteConfigDir"
	WafBlackConfigDirVariable = "NginxWafBlackConfigDir"

	UnixDefaultBaseDir = "/opt/nginx/"
	//WindowsDefaultBaseDir = "D:\\Program Files (x86)\\nginx-1.16.1\\"
	WindowsDefaultBaseDir = "F:\\Program Files (x86)\\nginx-1.12.2\\"

	WindowsDefaultBinaryPath     = WindowsDefaultBaseDir + "sbin\\nginx"
	UnixDefaultBinaryPath        = UnixDefaultBaseDir + "sbin/nginx"
	WindowsDefaultHttpDir        = WindowsDefaultBaseDir + "conf\\conf.d\\"
	UnixDefaultHttpDir           = UnixDefaultBaseDir + "conf/conf.d/"
	WindowsDefaultStreamDir      = WindowsDefaultBaseDir + "conf\\stream.d\\"
	UnixDefaultStreamDir         = UnixDefaultBaseDir + "conf/stream.d/"
	WindowsDefaultCoreConfigPath = WindowsDefaultBaseDir + "conf\\nginx.conf"
	UnixDefaultCoreConfigPath    = UnixDefaultBaseDir + "conf/nginx.conf"
	WindowsNginxVersionPath      = WindowsDefaultBaseDir + "conf\\version"
	UnixNginxtVersionPath        = UnixDefaultBaseDir + "conf/version"
	UnixDefaultWafDir            = UnixDefaultBaseDir + "conf/waf/"
	WindowsDefaultWafDir         = WindowsDefaultBaseDir + "conf\\waf\\"
	UnixDefaultWafWhiteDir       = UnixDefaultBaseDir + "conf/waf/white/"
	WindowsDefaultWafWhiteDir    = WindowsDefaultBaseDir + "conf\\waf\\white\\"
	UnixDefaultWafBlackDir       = UnixDefaultBaseDir + "conf/waf/black/"
	WindowsDefaultWafBlackDir    = WindowsDefaultBaseDir + "conf\\waf\\balck\\"
)

// 获取二进制文件路径
func BinaryPath() string {
	if os.Getenv(BinaryPathVariable) == "" {
		switch runtime.GOOS {
		case "windows":
			return WindowsDefaultBinaryPath
		default:
			return UnixDefaultBinaryPath
		}
	}
	return utils.DirAppendSlash(os.Getenv(BinaryPathVariable))
}

// 获取http server配置路径
func HttpConfigDir() string {
	if os.Getenv(HttpConfigDirVariable) == "" {
		switch runtime.GOOS {
		case "windows":
			return WindowsDefaultHttpDir
		default:
			return UnixDefaultHttpDir
		}
	}
	return utils.DirAppendSlash(os.Getenv(BinaryPathVariable))
}

// 获取tcp server配置路径
func StreamConfigDir() string {
	if os.Getenv(StreamConfigDirVariable) == "" {
		switch runtime.GOOS {
		case "windows":
			return WindowsDefaultStreamDir
		default:
			return UnixDefaultStreamDir
		}
	}
	return utils.DirAppendSlash(os.Getenv(StreamConfigDirVariable))
}

// 获取waf配置文件路径
func WafConfigDir() string {
	if os.Getenv(WafConfigDirVariable) == "" {
		switch runtime.GOOS {
		case "windows":
			return WindowsDefaultWafDir
		default:
			return UnixDefaultWafDir
		}
	}
	return utils.DirAppendSlash(os.Getenv(WafConfigDirVariable))
}

// 获取waf配置文件路径
func WafWhiteConfigDir() string {
	if os.Getenv(WafWhiteConfigDirVariable) == "" {
		switch runtime.GOOS {
		case "windows":
			return WindowsDefaultWafWhiteDir
		default:
			return UnixDefaultWafWhiteDir
		}
	}
	return utils.DirAppendSlash(os.Getenv(WafWhiteConfigDirVariable))
}

// 获取waf配置文件路径
func WafBlackConfigDir() string {
	if os.Getenv(WafBlackConfigDirVariable) == "" {
		switch runtime.GOOS {
		case "windows":
			return WindowsDefaultWafBlackDir
		default:
			return UnixDefaultWafBlackDir
		}
	}
	return utils.DirAppendSlash(os.Getenv(WafBlackConfigDirVariable))
}

// 获取tcp server配置路径
func CoreConfigPath() string {
	if os.Getenv(CoreConfigDirVariable) == "" {
		switch runtime.GOOS {
		case "windows":
			return WindowsDefaultCoreConfigPath
		default:
			return UnixDefaultCoreConfigPath
		}
	}
	return utils.DirAppendSlash(os.Getenv(CoreConfigDirVariable))
}

// 配置列表
func ConfigList() (configs []string) {
	configFiles := append(configs, CoreConfigPath(), HttpConfigDir(), StreamConfigDir())
	for _, f := range configFiles {
		fileInfo, _ := os.Stat(f)
		if fileInfo != nil && !utils.IsDir(f) {
			configs = append(configs, f)
		} else if fileInfo != nil && utils.IsDir(f) {
			configFiles, _ := ioutil.ReadDir(f)
			for _, v := range configFiles {
				if !v.IsDir() {
					configs = append(configs, utils.DirAppendSlash(f)+v.Name())
				}
			}
		}
	}
	return configs
}
