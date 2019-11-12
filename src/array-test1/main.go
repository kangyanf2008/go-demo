package main

import "fmt"

func main() {
/*	var result []string
	result = append(result,"11")
	result = append(result,"22")
	result = append(result,"33")
	fmt.Println(len(result))
	fmt.Println(cap(result))
	fmt.Println(result)

	var result2 []string
	fmt.Println(len(result2))
	fmt.Println(result2)
	fmt.Println(result2==nil)
	fmt.Println("截取长度")
	fmt.Println(result[0:3])


	vals := make([]interface{}, 3)
	vals[0]=1
	vals[1]=2
	vals[2]=3
	fmt.Printf("vals    %d \n",vals)

	vals2 :=[]interface{}{}
	vals2 = append(vals2,"111")
	vals2 = append(vals2, vals...)
	fmt.Printf("vals1    %s \n",vals2)

	vals3 :=[]interface{}{}
	vals3 = append(append(vals3,"111"), vals...)
	fmt.Printf("vals3    %s \n",vals3)
	fmt.Println(runtime.NumCPU())
*/
	//fmt.Println(removeElement([]int{1,2,3,4,5}, 1))

aa := []interface{}{11,"a",33,"bb"}
var cc []interface{}
	cc = append(cc, aa...)
	cc = append(cc, "tt",22)
	fmt.Println(cc)

}


func removeElement(nums []int, val int) []int {
	//如果是空切片，那就返回0
	if len(nums) == 0 {
		return nums
	}
	//用一个索引
	//循环去比较
	//当一样的时候就删除对应下标的值
	//当不一样的时候，索引加1
	index := 0
	for ; index < len(nums); {
		if nums[index] == val {
			nums = append(nums[:index], nums[index+1:]...)
			continue
		}
		index++
	}
	return nums
}