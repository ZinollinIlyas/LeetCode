package easy

func LongestCommonPrefix(strs []string) string {
	resArr := []byte(strs[0])

	if len(resArr) == 0 {
		return ""
	}
	for _, str := range strs {
		if len(str) < len(resArr) {
			resArr = resArr[:len(str)]
		}
		for j := 0; j < len(resArr); j++ {
			if str[j] != resArr[j] {
				if j == 0 {
					return ""
				} else {
					resArr = resArr[:j]
					break
				}
			}
		}

	}

	return string(resArr)

}
