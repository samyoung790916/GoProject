package main	// go는 반드시 package로 시작해야 한다. 이 코드가 속한 패키지를 나타낸다.
				// package는 코드를 묶는 단위이다. 
				// main ==> Entry Point 시작점을 포함하는 패키지 
				// go는 main 패키지와 다른 패키지로 구성되어 있다. 

import "fmt"	// fmt를 가져오겠다... fmt라는 패키지를 사용하겠다. 
				// 패키지를 왜 가져오냐? 패키지는 기능을 가지고 있다. 

func main() {
	fmt.Println("Hello Go World")
}


/*

func main(){
}

main은 키워드이다. main은 약속되어 있다. (프로그램 시작점)
fmt.Println() : fmt 패키지 ---> fmt 패키지 안에 있는 Println을 호출한다. 

*/


