package recursion

func reverseString(s []byte) {
	length := len(s)
	recurse(0, s, int(length))
}

func recurse(index int, s []byte, length int) {
	if nil == s || 0 == length {
		return
	}

	if length%2 == 0 {
		// even number
		if index == length/2 {
			return
		}
	} else {
		if index == length/2+1 {
			return
		}
	}

	recurse(index+1, s, length)

	tmp := s[length-index-1]
	s[length-index-1] = s[index]
	s[index] = tmp
}

func main() {
	var str = []byte("Gopher!")
	reverseString(str)
	print(string(str))
}
