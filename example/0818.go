package example

import (
	"fmt"
	"strings"
)

// & 메모리 값 복사, * 살펴보거나 훝어본다, 메모리 주소를 본다
// b는 a를 살펴보는 포인터라고 생각하면 된다.

func UpperCase(name string) (length int, upperCase string) {
	length = len(name)
	upperCase = strings.ToUpper(name)
	return
}

func DrinkAge(age int) bool {
	if koreanAge := age + 1; koreanAge > 20 {
		return true
	}

	return false
}

func Pointer() {
	a := 10 //메모리 주소를 변수에 할당
	b := a  //메모리 주소에 있는 값을 새로운 메모리 주소에 할당
	a = 20  // 메모리 주소에 있는 값을 변경
	// => 메모리 주소가 다름으로 b는 a의 값을 deep copy(독립적인 메모리 할당)

	c := 100 //메모리 주소를 변수에 할당
	d := &c  //c의 메모리 주소를 d에 할당 shallow copy(동일한 메모리 주소)
	c = 200  //동일한 메모리 주소의 값을 변경
	// => 같은 주소를 가르키고 있는 값을 변경 했음으로 c,d의 값은 같다.

	e := 1000 //메모리 주소를 변수에 할당
	f := &e   // e의 메모리 주소를 f에 할당 shallow copy(동일한 메모리 주소)
	*f = 2000 // 동일한 메모리 주소를 가지고 있는 f의 값을 변경
	// => '*'는 포인터의 개념으로 *사용하지 않으면 주소의 값을 보여준다.
	//'*'를 사용하여 주소안의 값을 변경하여 e,f 둘다 변경한다.

	z := [4]int{1, 2, 3, 4}
	y := z
	y[2] = 100
	// => golang에서는 단순 객체의 복제의 개념이 없다, mutable, immutable과 관계없이 안됨

	fmt.Println(a, b)  // 20, 10
	fmt.Println(c, *d) // 100, 100
	fmt.Println(e, *f) // 1000, 1000
	fmt.Println(z, y)  // [1 2 3 4] [1 2 100 4]

}

func main() {
	//example 1
	fmt.Println(UpperCase("wyatt"))

	//example 2
	fmt.Println(DrinkAge(18))

	//example 3
	Pointer()
}
