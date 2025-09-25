package main 

import "fmt"

func add(num1 int, num2 int)int{
	return num1 + num2;
}

func divide(a, b int)(int,error){
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}

func addAndMulti(a , b int)(sum int, product int){
	 sum = a + b
	 product = a * b
	 return
}
func main(){
	

 sum := 1

 for sum < 20{

	fmt.Println(sum)
  	sum += sum
 }
	
}





