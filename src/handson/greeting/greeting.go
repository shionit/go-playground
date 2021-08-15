package greeting

func Do() string {
	return "こんにちは" + packageValue
}

var packageValue string = "123"

func init() {
	packageValue = "スッキリ"
}
