package main

type Hoge struct {
	N int
}

// Fuga型にHoge型を埋め込む
type Fuga struct {
	Hoge // 名前のないフィールドになる
}

func main() {
	s := Fuga{}
	s.N = 1000
	println(s.Hoge.N)
}
