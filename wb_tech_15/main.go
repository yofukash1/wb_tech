
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	copy(justString, v[:100])
}

func main() {
	someFunc()
}
