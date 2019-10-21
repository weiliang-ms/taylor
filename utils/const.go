package utils

const (
	CustomConfigSuccessful    = "自定义配置文件已生效！"
	ModifySuccessful          = "修改成功！"
	UploadSuccessful          = "上传成功！"
	SuccessfulCode            = 0
	FailureCode               = 1
	CoreType                  = "core"
	HttpType                  = "http"
	TcpType                   = "tcp"
	WafType                   = "waf"
	OtherType                 = "other"
	DisableType               = "disable"
	HttpConfigSuffix          = ".conf"
	TcpConfigSuffix           = ".tcp"
	BackupSuffix              = ".bak"
	PortOutOfRange            = "端口超过取值范围"
	PortIllegal               = "端口非法，请检测输入的端口，必须为纯数字"
	PortInUse                 = "端口已经被占用"
	PortCheckSuccessful       = "端口检测成功"
	ConfigNameCheckSuccessful = "配置文件名称检测成功"
	ConfigNameIllegal         = "配置文件名称非法，只能以数字、英文字母、下划线、'-'命名"
	ConfigNameExist           = "配置文件名称重复，请更改名称！"
	CheckPortDirect           = "netstat -aln|grep "
	CoreConfigNotAllowModify  = "核心配置文件禁止修改！"
	ActiveStatus              = "active"
	DisableStatus             = "disable"
)
