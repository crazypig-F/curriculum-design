package decoding

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func NewCodebook(passwordFile string) map[string]string {
	codeBook := make(map[string]string)
	file1, err := ioutil.ReadFile(passwordFile)
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range strings.Split(string(file1), "\n") {
		line := strings.Split(string(line), " ")
		if len(line) == 2 {
			codeBook[line[1]] = line[0]
		}
	}
	return codeBook
}
