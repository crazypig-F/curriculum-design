package operation

import (
	shop2 "OrderingSystem/shop"
	"OrderingSystem/user"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func orderFood(customer user.User, shop shop2.Shop) {
	cmd := exec.Command("bash", "-c", "clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	back := 0
	for {
		fmt.Println("--------------------------------------")
		fmt.Println("\t欢迎来到点餐界面")
		fmt.Println("--------------------------------------")
		fmt.Println("\t1：查看所有菜品\n\t2：添加菜品\n\t3：删除菜品\n\t4：查看订单\n\t5：下单\n\tclear：清屏\n\tback：返回上一级\n\tq：退出")
		var opType string
		for opType != "1" && opType != "2" && opType != "3" && opType != "4" && opType != "5" && opType != "clear" && opType != "back" {
			fmt.Print("请输入：")
			fmt.Scanf("%s", &opType)
			switch opType {
			case "1":
				fmt.Println("--------------------------------------")
				fmt.Println(shop.GetMenu().Name, "\t菜单如下")
				fmt.Println("编号", "名称", "价格")
				fmt.Println("--------------------------------------")
				foods := shop.SelectAllFoods()
				for index, food := range foods {
					fmt.Println(index+1, "\t", food.Name, food.Price)
				}
				var quit string
				fmt.Println("输入quit退出")
				for quit != "quit" {
					fmt.Scanf("%s", &quit)
					if quit != "quit" {
						fmt.Println("输入不正确，请重新输入！")
					}
				}
			case "2":
				var foodId int
				var foodNum int
				fmt.Println("请输入菜品ID:")
				fmt.Scanf("%d", &foodId)
				for foodNum == 0 {
					fmt.Println("请输入数量:")
					fmt.Scanf("%d", &foodNum)
					if foodNum == 0 {
						fmt.Println("数量不正确，请重新输入")
					}
				}
				food := shop.GetFood(foodId)
				if food.Price != 0 {
					customer.AddFood(foodId, shop.Id, foodNum)
				}
				fmt.Println("添加成功")
				time.Sleep(2 * time.Second)
			case "3":
				var foodId int
				fmt.Println("请输入菜名ID:")
				fmt.Scanf("%d", &foodId)
				customer.DeleteFood(foodId, shop.Id)
				fmt.Println("删除成功")
				time.Sleep(2 * time.Second)
			case "4":
				fmt.Println("--------------------------------------")
				fmt.Println(customer.Name, "(先生/女士)，您的订单如下")
				fmt.Println("编号", "名称", "价格", "数量")
				fmt.Println("--------------------------------------")
				foods, numbers := customer.SelectAllFoods(shop.Id)
				var totalPrice float32
				for index, f := range foods {
					fmt.Println(index+1, f.Name, f.Price, numbers[index])
					totalPrice += f.Price * float32(numbers[index])
				}
				fmt.Println("总价为：", totalPrice )
				var quit string
				fmt.Println("输入quit退出")
				for quit != "quit" {
					fmt.Scanf("%s", &quit)
					if quit != "quit" {
						fmt.Println("输入不正确，请重新输入！")
					}
				}
			case "5":
				customer.Order(shop.Id)
				fmt.Println("下单成功")
				back = 1
				time.Sleep(2 * time.Second)
			case "clear":
				cmd := exec.Command("bash", "-c", "clear")
				cmd.Stdout = os.Stdout
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
			case "back":
				back = 1
			case "q":
				os.Exit(0)
			}
			if opType != "1" && opType != "2" && opType != "3" && opType != "4" && opType != "5" && opType != "clear" && opType != "back" {
				fmt.Println("输入不正确，请重新输入：")
			}
		}
		if back == 1 {
			break
		}
	}
}

func CustomerOperation(customer user.User) {
	back := 0
	cmd := exec.Command("bash", "-c", "clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("--------------------------------------")
		fmt.Println("\t欢迎来到点餐界面")
		fmt.Println("--------------------------------------")
		fmt.Println("\t请选择操作:")
		fmt.Println("\t1：查看所有商铺\n\tclear：清屏\n\tback：返回上一级\n\tq：退出")
		var opType string
		fmt.Print("请输入：")
		for opType != "1" && opType != "clear" && opType != "back" {
			fmt.Scanf("%s", &opType)
			switch opType {
			case "1":
				shops := customer.GetAllShop()
				for index, s := range shops {
					fmt.Println(index+1, "\t", s.Name)
				}
				fmt.Println("请选择商铺：")
				var shopIndex int
				for shopIndex <= 0 || shopIndex > len(shops) {
					fmt.Scanf("%d", &shopIndex)
					if shopIndex <= 0 || shopIndex > len(shops) {
						fmt.Println("输入错误，请重新输入")
					} else {
						orderFood(customer, shops[shopIndex-1])
					}
				}
			case "clear":
				cmd := exec.Command("bash", "-c", "clear")
				cmd.Stdout = os.Stdout
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
			case "back":
				back = 1
			case "q":
				os.Exit(0)
			}
			if opType != "1" && opType != "clear" && opType != "back" {
				fmt.Println("输入不正确，请重新输入：")
			}
		}
		if back == 1 {
			break
		}
	}
}
