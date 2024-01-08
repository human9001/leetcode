package main

func lengthOfLongestSubstring(s string) int {

	letters := []byte{}
	tmp_letters := []byte{}
	i := 0
	if len(s) == 1 {
		return 1
	}
	for i < len(s) {
		letters = append(letters, s[i])
		ii := i + 1
		for ii < len(s) {

			exists := inSlice(letters, s[ii])

			if !exists {
				letters = append(letters, s[ii])
			}

			if exists || ii == len(s)-1 {

				if len(letters) > len(tmp_letters) {
					tmp_letters = append([]byte{}, letters...)

				}
				break
			}

			ii++
		}
		letters = nil
		i++
	}
	return len(tmp_letters)

}

func inSlice(s []byte, f byte) bool {
	for _, v := range s {
		if v == f {
			return true
		}
	}
	return false
}
