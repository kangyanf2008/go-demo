package main
/**
// C 标志io头文件，你也可以使用里面提供的函数
#include <stdio.h>
void pri(){
	printf("fdsfs");
    printf("hey");
}
int add(int a,int b){
    return a+b;
}
 */
import "C"
import "fmt"

func main() {
	fmt.Println(C.add(1,2))
	C.pri()
}
