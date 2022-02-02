package main

import (
	"OrderingSystem/operation"
	user2 "OrderingSystem/user"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	for {
		cmd := exec.Command("bash", "-c", "clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("--------------------------------------")
		fmt.Println("\t欢迎来到点餐系统")
		fmt.Println("--------------------------------------")
		fmt.Println("\t请选择进入方式:")
		fmt.Println("\t1：登录\n\t2：注册\n\tclear：清屏\n\tq：退出")
		var enterType string
		fmt.Print("请输入：")
		for enterType != "1" && enterType != "2" && enterType != "clear" {
			fmt.Scanf("%s", &enterType)
			switch enterType {
			case "1":
				var user user2.User
				fmt.Println("请输入用户名：")
				fmt.Scanf("%s", &user.Name)
				fmt.Println("请输入密码：")
				fmt.Scanf("%s", &user.Password)
				if loginUser, flag := user.Login(); flag{
					userType := strconv.Itoa(loginUser.Type)
					switch userType {
					case "1":
						operation.ShopOperation(loginUser)
						break
					case "2":
						operation.CustomerOperation(loginUser)
					}
				}else{
					fmt.Println("用户名密码不正确！")
					time.Sleep(2 * time.Second)
				}
			case "2":
				var user user2.User
				fmt.Println("请输入用户名：")
				fmt.Scanf("%s", &user.Name)
				fmt.Println("请输入密码：")
				fmt.Scanf("%s", &user.Password)
				fmt.Println("请选择用户类型(1：商家，2：顾客)：")
				fmt.Scanf("%d", &user.Type)
				fmt.Println("请输入手机号：")
				fmt.Scanf("%s", &user.Phone)
				fmt.Println("请输入地址：")
				fmt.Scanf("%s", &user.Address)
				user.Register()
				fmt.Println("注册成功")
				time.Sleep(2 * time.Second)
			case "clear":
				cmd := exec.Command("bash", "-c", "clear")
				cmd.Stdout = os.Stdout
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
			case "q":
				os.Exit(0)
			}
			if enterType != "1" && enterType != "2" && enterType != "clear" {
				fmt.Println("输入不正确，请重新输入：")
			}
		}
	}
}
