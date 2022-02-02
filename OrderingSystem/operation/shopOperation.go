package operation

import (
	food2 "OrderingSystem/food"
	menu2 "OrderingSystem/menu"
	shop2 "OrderingSystem/shop"
	"OrderingSystem/user"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ShopOperation(user user.User){
	back := 0
	cmd := exec.Command("bash", "-c", "clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("--------------------------------------")
		fmt.Println("\t欢迎来到商铺管理界面")
		fmt.Println("--------------------------------------")
		fmt.Println("\t请选择操作:")
		fmt.Println("\t1：添加菜品\n\t2：删除菜品\n\t3：修改菜品\n\t4：查看菜单\n\t5：添加商铺\n\tclear：清屏\n\tback：返回上一级\n\tq：退出")
		var opType string
		fmt.Print("请输入：")
		for opType != "1" && opType != "2" && opType != "3" && opType != "4" && opType != "5" && opType != "clear" && opType != "back" {
			fmt.Scanf("%s", &opType)
			switch opType {
			case "1":
				shop := user.GetShop(user.Id)
				var food food2.Food
				fmt.Println("请输入菜名:")
				fmt.Scanf("%s", &food.Name)
				fmt.Println("请输入价格:")
				fmt.Scanf("%f", &food.Price)
				shop.SaveFood(food)
				fmt.Println("添加品成功")
				time.Sleep(2 * time.Second)
			case "2":
				shop := user.GetShop(user.Id)
				var foodName string
				fmt.Println("请输入菜名:")
				fmt.Scanf("%s", &foodName)
				shop.DeleteFood(foodName)
				fmt.Println("删除成功")
				time.Sleep(2 * time.Second)
			case "3":
				shop := user.GetShop(user.Id)
				var foodName string
				var foodPrice float32
				fmt.Println("请输入菜名:")
				fmt.Scanf("%s", &foodName)
				fmt.Println("请输入价格:")
				fmt.Scanf("%f", &foodPrice)
				shop.UpdateFood(foodName, foodPrice)
				fmt.Println("更新成功")
				time.Sleep(2 * time.Second)
			case "4":
				shop := user.GetShop(user.Id)
				fmt.Println("菜单如下")
				foods := shop.SelectAllFoods()
				for _, v := range foods{
					fmt.Println(v.Name, v.Price)
				}
				var quit string
				for quit != "quit"{
					fmt.Println("请输入quit退出！")
					fmt.Scanf("%s", &quit)
				}
				time.Sleep(2 * time.Second)
			case "5":
				var shop shop2.Shop
				shop.BossId = user.Id
				fmt.Println("请输入商铺名:")
				fmt.Scanf("%s", &shop.Name)
				fmt.Println("请输入菜单名:")
				fmt.Scanf("%s", &shop.MenuName)
				shop.MenuName = shop.Name + "_" + shop.MenuName
				user.SetShop(shop)
				fmt.Println("添加商铺成功")
				shop1 := user.GetShop(user.Id)
				var menu menu2.Menu
				menu.ShopId = shop1.Id
				menu.Name = shop1.MenuName
				shop1.SetMenu(menu)
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
			if opType != "1" && opType != "2" && opType != "3" && opType != "4" && opType != "5" && opType != "clear" && opType != "back"{
				fmt.Println("输入不正确，请重新输入：")
			}
		}
		if back == 1{
			break
		}
	}
}
