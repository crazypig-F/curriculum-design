package encode

import (
	"HuffmanCoding/password"
	"HuffmanCoding/tree"
)

func TreeEncode(root *tree.HuffmanTreeNode, flag string, word password.PassWord) {
	if root == nil{
		return
	}
	if root.Left == nil && root.Right == nil{
		word.Set(root.Val, flag)
	}
	TreeEncode(root.Left, flag + "0", word)
	TreeEncode(root.Right, flag + "1", word)
}

func Encode(root *tree.HuffmanTreeNode, inputFile string) string{
	word := password.NewPassword()
	TreeEncode(root, "", word)
	word.Save("data/passwordFile.txt")
	code := ""
	for _, s := range inputFile{
		if string(s) == " "{
			code += word.Get("\\s")
		}else if string(s) == "\n"{
			code += word.Get("\\n")
		}else{
			code += word.Get(string(s))
		}
	}
	return code
}