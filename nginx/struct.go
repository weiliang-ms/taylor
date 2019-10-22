package nginx

// 前台操作配置的相关信息（自定义配置、上传配置接口）
type NginxConfigInfo struct {
	Path       string `json:"path"`       // 配置文件路径
	Content    string `json:"content"`    // 配置文件内容
	Protocol   string `json:"protocol"`   // nginx代理的协议
	Naxsi      string `json:"naxsi"`      // 是否开启naxsi模块
	ExposePort string `json:"exposePort"` // 防火墙是否开放端口
	Port       int `json:"port"`       // 配置监听端口
	Name       string `json:"name"`       // 配置文件名称
}

type UploadFile struct {
	Filename string `form:"filename"`
}

type NginxInstanceInfo struct {
	IP             string `json:"ip"` // 宿主机IP
	WorkerNum      string `json:"worker_num"`
	WafStatus      string `json:"waf_status"`
	InstanceStatus string `json:"instance_status"`
}
