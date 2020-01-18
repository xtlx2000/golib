package main

func main() {
	mm := make(map[int]string)
	mm[1] = "hehe"

	m, isExist := mm[1]
	m = ""
}
