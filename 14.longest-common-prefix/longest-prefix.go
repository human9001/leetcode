package main

func longestCommonPrefix(strs []string) string {
	var res string
	var endLoop bool
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 1; i <= len(strs[0]); i++ {
		prefix := strs[0][:i]
		for _, v := range strs {
			if len(v) >= len(prefix) && v[0:len(prefix)] == prefix {
				res = prefix
			} else {
				if len(prefix) > 0 {
					res = prefix[:len(prefix)-1]
				}
				endLoop = true
				break
			}
		}
		if endLoop {
			break
		}
	}
	return res
}
