package main

//func main() {
//	str := ""
//	fmt.Println(positions(str, len(str)))
//}

const NUM = 31

func positions(str string, n int) (result uint8) {
	for i := 0; i < n; i++ {
		result += str[i] & NUM
	}
	return
}

func count(a int, b int) int {
	return 0
}
