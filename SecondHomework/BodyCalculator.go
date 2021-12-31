package main

import "fmt"

func main() {
	var count int
	var sum float32
	for {
		var opt int
		var name string
		var sex string
		var age int
		var height float32
		var weight float32
		var gender bool
		_ = gender
		var bmi float32
		_ = bmi
		var rateOfBody float32
		var average float32
		_ = average
		fmt.Println("成年人体脂计算器\n请输入：\n1.计算体脂\n2.退出")
		fmt.Scanln(&opt)
		if opt == 1 {
			count++
			fmt.Println("请输入您的姓名")
			fmt.Scanln(&name)
			fmt.Println("请输入您的性别")
			fmt.Scanln(&sex)
			if sex == "男" {
				gender = true
			} else if sex == "女" {
				gender = false
			} else {
				fmt.Println("请输入一个正确的性别")
				continue
			}
			fmt.Println("请输入您的年龄")
			fmt.Scanln(&age)
			if age < 18 {
				fmt.Println("对不起，未授权未成年人使用")
				continue
			}
			if age < 0 && age > 150 {
				fmt.Println("请输入一个正确的年龄")
				continue
			}
			fmt.Println("请输入您的身高")
			fmt.Scanln(&height)
			if height <= 0 {
				fmt.Println("输入错误")
				continue
			}
			height = height / 100
			fmt.Println("请输入您的体重")
			fmt.Scanln(&weight)
			if weight <= 0 {
				fmt.Println("输入错误")
				continue
			}
			bmi = weight / (height * height)
			if gender {
				rateOfBody = 1.2*bmi + 0.23*float32(age) - 5.4 - 10.8
			} else if !gender {
				rateOfBody = 1.2*bmi + 0.23*float32(age) - 5.4
			}
			fmt.Println("姓名:")
			fmt.Println(name)
			fmt.Println("BMI:")
			fmt.Printf("%.2f\n", bmi)
			fmt.Println("体脂率：")
			fmt.Printf("%.2f\n", rateOfBody)
			sum = sum + rateOfBody
			average = sum / float32(count)
			if gender {
				if age >= 18 && age <= 39 {
					if rateOfBody >= 5 && rateOfBody < 10 {
						fmt.Println("偏瘦")
					} else if rateOfBody >= 10 && rateOfBody < 21 {
						fmt.Println("标准")
					} else if rateOfBody >= 21 && rateOfBody < 26 {
						fmt.Println("偏胖")
					} else if rateOfBody >= 26 && rateOfBody < 45 {
						fmt.Println("过胖")
					} else {
						fmt.Println("数据异常，如有需要请立即去医院就诊")
					}
				}
				if age >= 40 && age <= 59 {
					if rateOfBody >= 5 && rateOfBody < 11 {
						fmt.Println("偏瘦")
					} else if rateOfBody >= 11 && rateOfBody < 22 {
						fmt.Println("标准")
					} else if rateOfBody >= 22 && rateOfBody < 27 {
						fmt.Println("偏胖")
					} else if rateOfBody >= 27 && rateOfBody < 45 {
						fmt.Println("过胖")
					} else {
						fmt.Println("数据异常，如有需要请立即去医院就诊")
					}
				}
				if age >= 60 {
					if rateOfBody >= 5 && rateOfBody < 13 {
						fmt.Println("偏瘦")
					} else if rateOfBody >= 13 && rateOfBody < 24 {
						fmt.Println("标准")
					} else if rateOfBody >= 24 && rateOfBody < 29 {
						fmt.Println("偏胖")
					} else if rateOfBody >= 29 && rateOfBody < 45 {
						fmt.Println("过胖")
					} else {
						fmt.Println("数据异常，如有需要请立即去医院就诊")
					}
				}
			}
			if !gender {
				if age >= 18 && age <= 39 {
					if rateOfBody >= 5 && rateOfBody < 20 {
						fmt.Println("偏瘦")
					} else if rateOfBody >= 20 && rateOfBody < 34 {
						fmt.Println("标准")
					} else if rateOfBody >= 34 && rateOfBody < 39 {
						fmt.Println("偏胖")
					} else if rateOfBody >= 39 && rateOfBody < 45 {
						fmt.Println("过胖")
					} else {
						fmt.Println("数据异常，如有需要请立即去医院就诊")
					}
				}
				if age >= 40 && age <= 59 {
					if rateOfBody >= 5 && rateOfBody < 21 {
						fmt.Println("偏瘦")
					} else if rateOfBody >= 21 && rateOfBody < 35 {
						fmt.Println("标准")
					} else if rateOfBody >= 35 && rateOfBody < 40 {
						fmt.Println("偏胖")
					} else if rateOfBody >= 40 && rateOfBody < 45 {
						fmt.Println("过胖")
					} else {
						fmt.Println("数据异常，如有需要请立即去医院就诊")
					}
				}
				if age >= 60 {
					if rateOfBody >= 5 && rateOfBody < 22 {
						fmt.Println("偏瘦")
					} else if rateOfBody >= 22 && rateOfBody < 36 {
						fmt.Println("标准")
					} else if rateOfBody >= 36 && rateOfBody < 41 {
						fmt.Println("偏胖")
					} else if rateOfBody >= 41 && rateOfBody < 45 {
						fmt.Println("过胖")
					} else {
						fmt.Println("数据异常，如有需要请立即去医院就诊")
					}
				}
			}
			fmt.Println("人数")
			fmt.Println(count)
			fmt.Println("平均体脂率")
			fmt.Printf("%.2f\n", average)
			fmt.Println("结束,重新选择")
		} else if opt == 2 {
			count = 0
			sum = 0
			fmt.Println("退出")
			break
		} else {
			fmt.Println("请重新输入正确的数字")
		}
	}
}
