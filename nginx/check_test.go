package nginx

import (
	"testing"
)

// 检测端口合法性单元测试
func TestCheckPort(t *testing.T) {
	cases := []struct{
		port int
		expected bool
	}{
		{6643266432,false},
		{4432,true},
		{-188,false},
	}
	for _, v := range cases {
		_, re := checkPort(v.port)
		if re != v.expected {
			t.Fatalf("check port function failed, port: %d, execpted:%t, result:%t", v.port, v.expected, re)
		}
	}
}

// 检测配置文件名称
func TestCheckConfigName(t *testing.T) {
	cases := []struct{
		name string
		path string
		expected bool
	}{
		{"*dsdd-!~", HttpConfigDir() + "*dsdd-!~",false},
		{"12nadsd",HttpConfigDir() + "12nadsd",true},
		{"*dsdd-!~", StreamConfigDir() + "*dsdd-!~",false},
		{"12nadsd",StreamConfigDir() + "12nadsd",true},
	}

	for _, v := range cases {
		msg, re := checkConfigName(v.name,v.path)
		if re != v.expected {
			t.Fatalf("check config name function failed, name: %s, path: %s, execpted:%t, result:%t, msg: %s", v.name, v.path, v.expected, re, msg)
		}
	}
}
