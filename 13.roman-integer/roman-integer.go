package main

func romanToInt(s string) int {
	var sum int
	mapRomans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	if len(s) == 1 {
		return mapRomans[s]
	}

	current := 0
	next := 0
	for i := 0; i < len(s)-1; i++ {
		current = mapRomans[string(s[i])]
		next = mapRomans[string(s[i+1])]
		if current < next {
			sum += (next - current)
			i++
			if i == len(s)-2 {
				sum += mapRomans[string(s[len(s)-1])]
			}
		} else {
			sum += current
			if i == len(s)-2 {
				sum += next
			}

		}

	}
	return sum
}
