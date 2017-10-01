package main

func main() {
	w := make([]*struct{}, 5)
	for _, p := range w {
		p = &struct{}{}
		_ = p
	}
}
