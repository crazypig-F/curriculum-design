package password

import (
	"fmt"
	"io/ioutil"
)

type PassWord map[string]string
func NewPassword() PassWord{
	return make(PassWord)
}

func (p PassWord)Set(key, val string){
	p[key] = val
}

func (p PassWord)Get(key string)string{
	return p[key]
}

func (p PassWord)Save(passwordFile string){
	data := ""
	for key, val := range p{
		data += key + " " + val + "\n"
	}
	err := ioutil.WriteFile(passwordFile, []byte(data), 0777)
	if err != nil {
		fmt.Println(err)
	}
}



