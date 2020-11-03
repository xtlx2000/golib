package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	str1, err := ioutil.ReadFile("pt.txt")
	if err != nil {
		fmt.Printf("read err", err)
		return
	}
	str := string(str1)

	match := regexp.MustCompile("download.php.https=1&id=[0-9]+")

	result := match.FindAllStringSubmatch(str, -1)

	output := ""

	for _, v := range result {
		output += ("https://leaguehd.com/" + v[0] + "\n")
	}

	fmt.Printf("%v", output)
	return
}
