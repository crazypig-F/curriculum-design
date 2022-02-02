package main

import (
	"HuffmanCoding/data"
	"HuffmanCoding/decoding"
	"HuffmanCoding/encode"
	"HuffmanCoding/tree"
	"fmt"
	"io/ioutil"
)


func main()  {
	file1, err := ioutil.ReadFile("data/inputFile.txt")
	if err != nil {
		fmt.Print(err)
	}
	queue := tree.ChangeNode(data.CountCharacter(file1, "data/hfmTree.txt"))
	root := tree.CreateTree(queue)
	code := encode.Encode(root, string(file1))
	fmt.Println(code)
	decoding.Decoding(code)
}


