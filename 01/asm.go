package one

func str2ToInt(str string) int {
	num := 0
	n := len(str)
	for i := n - 1; i >= 0; i-- {
		c := int(byte(str[i] - '0'))
		num += c * (1 << (n - i - 1))
	}
	return num
}

func str16ToInt(str string) int {
	num := 0
	n := len(str)
	for i := n - 1; i >= 0; i-- {
		c := 0
		if str[i] < '9' {
			c = int(byte(str[i] - '0'))
		} else {
			c = int(byte(str[i]-'a')) + 10
		}

		num += c * (1 << ((n - i - 1) * 4))
	}
	return num
}
