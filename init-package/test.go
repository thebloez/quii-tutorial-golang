package init_package

var someString string

func init() {
	someString = "Hello World"
}

func GetInit() string {
	return someString
}
