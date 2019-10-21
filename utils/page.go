package utils

// 分页
func Paging(page int, limit int, srcData []string) (data []string) {
	if len(srcData) < 1 {
		return srcData
	}
	startIndex := (page - 1) * limit
	endIndex := limit * page
	if endIndex > len(srcData) {
		endIndex = len(srcData)
	}
	return srcData[startIndex:endIndex]
}
