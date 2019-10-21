package utils

import "strings"

func DeleteSubString(s string,args ...string) string{
	return func(args []string)string{
		for _, v := range args {
			s = strings.Replace(s,v,"",-1)
		}
		return s
	}(args)
}
