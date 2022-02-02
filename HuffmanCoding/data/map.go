package data

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type Node map[string]int

// Sort 冒泡排序
func Sort(queue []Node) {
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(queue)-i-1; j++ {
			var val1, val2 int
			for _, v := range queue[j] {
				val1 = v
			}
			for _, v := range queue[j+1] {
				val2 = v
			}
			if val1 > val2 {
				queue[j], queue[j+1] = queue[j+1], queue[j]
			}
		}
	}
}

func CountCharacter(inputFile []byte, outputFile string) []Node {
	m := make(Node)
	for _, c := range inputFile {
		if string(c) == "\n" {
			m["\\n"] += 1
		} else if string(c) == " " {
			m["\\s"] += 1
		} else {
			m[string(c)] += 1
		}
	}

	// 通过m的val的大小来排序
	var Queue []Node
	for key, val := range m {
		n := Node{key: val}
		Queue = append(Queue, n)
	}
	Sort(Queue)
	data := ""
	for _, node := range Queue {
		for key, val := range node {
			data += key + " " + strconv.Itoa(val) + "\n"
		}
	}
	err := ioutil.WriteFile(outputFile, []byte(data), 0777)
	if err != nil {
		fmt.Println(err)
	}
	return Queue
}
