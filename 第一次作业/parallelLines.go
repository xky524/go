package main

import "fmt"

func main() {
	line1 :=[][]int{}
	row1 :=[]int{1,3}
	list1 :=[]int{1,3}
	line1 = append(line1,row1)
	line1 = append(line1,list1)
	line2 :=[][]int{}
	row2 :=[]int{4,6}
	list2 :=[]int{6,8}
	line2=append(line2,row2)
	line2=append(line2,list2)
	var k int
	k=(line1[0][1]-line1[0][0])*(line2[1][1]-line2[1][0])-(line2[0][1]-line2[0][0])*(line1[1][1]-line1[1][0])
	var k1 int
	//两条线各取一点形成新的直线，若新直线与原直线不平行，则原两条直线不重合
	k1=(line2[0][1]-line1[0][0])*(line1[1][1]-line1[1][0])-(line1[0][1]-line1[0][0])*(line2[1][1]-line1[1][0])
	if k==0 && k1!=0{
		fmt.Println("两条直线平行")
	}else {
		fmt.Println("不平行或重合")
	}



}
