package tutorial

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

func init() {
	fmt.Printf("\"I am in init function in tutoral package\": %v\n", "I am in init function in tutoral package")
}

const PI float32 = 3.14

func Logical() {
	var (
		b1 = true
		b2 = false
	)

	age := 19
	gender := "M"

	if age > -18 && gender == "M" {
		fmt.Printf("\"你是成年男子\": %v\n", "你是成年男子")
	}

	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
}

func Integer() {
	var a int = 10

	fmt.Printf("a: %b\n", a)
	fmt.Printf("a: %o\n", a)
	fmt.Printf("a: %x\n", a)
	fmt.Printf("a: %d\n", a)
	fmt.Printf("PI: %.2f\n", math.Pi)
}

func String() {
	var s string = "Hello world"
	var html string = `
		<html>
			<head></head>
			<body>hello world</body>
		</html>
	`

	var prev string = "prev "
	var last string = "last"

	fmt.Printf("s: %v\n", s)
	fmt.Printf("html: %v\n", html)
	msg := prev + last
	fmt.Printf("%s\n", msg)

	result := strings.Join([]string{prev, last}, "")
	fmt.Printf("%s\n", result)

	var buffer bytes.Buffer
	buffer.WriteString("Tom")
	buffer.WriteString(", ")
	buffer.WriteString("20")
	fmt.Println(buffer.String())

	str := "hello world"
	fmt.Println(len(str))

	fmt.Printf("strings.Split(str): %v\n", strings.Split(str, " "))
	fmt.Printf("strings.Contains(\"hello\"): %v\n", strings.Contains(str, "hello"))
	fmt.Printf("strings.ToLower(str): %v\n", strings.ToLower(str))
	fmt.Printf("strings.ToUpper(str): %v\n", strings.ToUpper(str))
	fmt.Printf("strings.HasPrefix(str, \"hello\"): %v\n", strings.HasPrefix(str, "hello"))
	fmt.Printf("strings.HasSuffix(str, \"world\"): %v\n", strings.HasSuffix(str, "world"))
	fmt.Printf("strings.Index(str, \"l\"): %v\n", strings.Index(str, "l"))
	fmt.Printf("strings.LastIndex(str, \"l\"): %v\n", strings.LastIndex(str, "l"))

}

func Week() {
	var c string
	fmt.Printf("\"Please enter a char: \": %v\n", "Please enter a char: ")
	fmt.Scan(&c)

	switch strings.ToUpper(c) {
	case "F":
		fmt.Printf("\"Friday\": %v\n", "Friday")
	case "M":
		fmt.Printf("\"Monday\": %v\n", "Monday")
	case "W":
		fmt.Printf("\"Wendesday\": %v\n", "Wendesday")
	case "T":
		var s string
		fmt.Printf("\"Please enter into another char \": %v\n", "Please enter into another char ")
		fmt.Scan(&s)

		switch strings.ToUpper(s) {
		case "H":
			fmt.Printf("\"Thursday\": %v\n", "Thursday")
		case "U":
			fmt.Printf("\"Tuesday\": %v\n", "Tuesday")
		default:
			fmt.Printf("\"Error\": %v\n", "Error")
		}
	case "S":
		var s string
		fmt.Printf("\"Please enter into another char \": %v\n", "Please enter into another char ")
		fmt.Scan(&s)

		switch strings.ToUpper(s) {
		case "U":
			fmt.Printf("\"Sunday\": %v\n", "Sunday")
		case "A":
			fmt.Printf("\"Saturday\": %v\n", "Saturday")
		default:
			fmt.Printf("\"Error\": %v\n", "Error")
		}
	default:
		fmt.Printf("\"Error\": %v\n", "Error")
	}
}

func IF() {
	var a int = 100

	if a > 20 {
		fmt.Printf("\">\": %v\n", ">")
	} else {
		fmt.Printf("\"<\": %v\n", "<")
	}

	flag := true
	if flag {
		fmt.Printf("\"True\": %v\n", "True")
	} else {
		fmt.Printf("\"False\": %v\n", "False")
	}

	if age, name := 20, "Bob"; age > 18 {
		fmt.Printf("\"Adult\": %v\n", "Adult")
		fmt.Printf("age: %v\n", age)
		fmt.Printf("name: %v\n", name)

	} else {
		fmt.Printf("\"Teen\": %v\n", "Teen")
		fmt.Printf("age: %v\n", age)
	}

	// fmt.Printf("\"Please enter values: \": %v\n", "Please enter values: ")
	// var (
	// 	name  string
	// 	age   int
	// 	email string
	// )
	// fmt.Scan(&name, &age, &email)
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)
	// fmt.Printf("email: %v\n", email)

	score := 95
	if score > 90 {
		fmt.Printf("\"great\": %v\n", "great")
	} else if score > 70 {
		fmt.Printf("\"good\": %v\n", "good")
	} else {
		fmt.Printf("\"not good\": %v\n", "not good")
	}

	// Week()

	a, b, c := 1, 2, 3

	if a > b {
		if a > c {
			fmt.Printf("\"max\": %v\n", a)
		} else {
			fmt.Printf("\"max\": %v\n", c)
		}
	} else {
		if b > c {
			fmt.Printf("\"max\": %v\n", b)
		} else {
			fmt.Printf("\"max\": %v\n", c)
		}
	}

	score_01 := 81
	switch {
	case score_01 > 90:
		fmt.Printf("\"great\": %v\n", "great")
	case score_01 > 80:
		fmt.Printf("\"good\": %v\n", "good")
		fallthrough
	case score_01 > 60:
		fmt.Printf("\"not good\": %v\n", "not good")
		// fallthrough
	default:
		fmt.Printf("\"bad\": %v\n", "bad")
	}
}

func For() {

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\t", i)
	}
	fmt.Println("")

	j := 1
	for ; j < 5; j++ {
		fmt.Printf("j: %v\t", j)
	}
	fmt.Println("")

	k := 1
	for k < 5 {
		fmt.Printf("k: %v\t", k)
		k++
	}
	fmt.Println("")

	// for {
	// 	fmt.Printf("\"永远循环\": %v\n", "永远循环")
	// }

	var a = [3]int{1, 2, 3}
	for _, idx := range a {
		fmt.Printf("idx: %v\t", idx)
	}
	fmt.Println("")

	var age = [...]int{1, 1, 2, 3, 5, 4}
	for _, v := range age {
		fmt.Printf("v: %v\t", v)
	}
	fmt.Println("")

	var s = []int{1, 2, 3}
	for _, s_idx := range s {
		fmt.Printf("s_idx: %v\t", s_idx)
	}
	fmt.Println("")

	var m = make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "29"
	m["gender"] = "M"
	for k, v := range m {
		fmt.Printf("k: %v\t", k)
		fmt.Printf("v: %v\t", v)
	}
	fmt.Println("")

	var str = "hello world"
	for _, val := range str {
		fmt.Printf("val: %c\t", val)
	}
	fmt.Println("")

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\t", i)
		if i >= 5 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				goto END
			}
			fmt.Println(i, " ", j)
		}
	}

END:
	fmt.Printf("\"Ending...\": %v\n", "Ending...")
}

func Array() {
	// var a1 [2]int
	// var a2 [3]string
	// var a3 [4]bool
	// var a4 [5]float32

	// fmt.Printf("a1: %T\n", a1)
	// fmt.Printf("a1: %v\n", a1)
	// fmt.Printf("a2: %T\n", a2)
	// fmt.Printf("a2: %v\n", a2)
	// fmt.Printf("a3: %T\n", a3)
	// fmt.Printf("a3: %v\n", a3)
	// fmt.Printf("a4: %T\n", a4)
	// fmt.Printf("a4: %v\n", a4)

	var a1 = [3]int{1, 3, 3}
	fmt.Printf("a1: %v\n", a1)

	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)

	var a3 = [...]int{1, 2, 3, 4}
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("len(a3): %v\n", len(a3))

	var a4 = [...]string{1: "a_dab", 2: "da_sd"}
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a4[0]: %v\n", a4[0])
	fmt.Printf("len(a4[0]): %v\n", len(a4[0]))

	var a5 = [...]int{1, 2, 3, 5, 5, 5, 2}
	for i := 0; i < len(a5); i++ {
		fmt.Printf("a5[%v]: %v\t", i, a5[i])
	}
	fmt.Println("")

	for _, v := range a5 {
		fmt.Printf("v: %v\t", v)
	}
	fmt.Println("")
}

func Slice() {
	// Slice 可以动态扩容的数组
	// var identifier []type
	// 若是不确定长度的数组，括号内为..., 而此处是可变长度的数组
	// var slice_1 []type = make([]type, len)
	// slice_1 := make([]type, len)
	// len为slice的起始长度，capacity为容量长度（capacity >= len)
	// slice_1 := make([]type, len, capacity)
	var s1 []int
	var s2 []string
	fmt.Printf("s1: %T\n", s1)
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("s2: %T\n", s2)
	fmt.Printf("len(s2): %v\n", len(s2))

	var s3 = make([]int, 3)
	fmt.Printf("s3: %T\n", s3)
	fmt.Printf("len(s3): %v\n", len(s3))

	// 初始化
	var s4 = []int{1, 3, 4}
	fmt.Printf("len(s4): %v\n", len(s4))
	fmt.Printf("s4[1]: %v\n", s4[1])

	var a1 = [...]int{1, 3, 4}
	fmt.Printf("a1: %T\n", a1)

	s5 := a1[:]
	fmt.Printf("s5: %T\n", s5)

	s6 := []int{1, 2, 4}
	s6 = append(s6, 4)
	fmt.Printf("s6: %v\n", s6)

	s7 := append(s6[:2], s6[3:]...)
	fmt.Printf("s7: %v\n", s7)

}

func Map() {
	// 变量声明
	// var identifier map[key_type]value_type
	// 使用:=时，若是有直接赋初始值，就不能用make；若是没有初始值，则必须要使用make
	// identifier := make(map[key_type]value_type)
	// var identifier = make(map[key_type]value_type)
	m1 := make(map[string]string)

	// 定义好之后再初始化
	// m1[key] = value
	m1["name"] = "Bob"
	fmt.Printf("m1[\"name\"]: %v\n", m1["name"])
	// 定义时直接赋值，此时不能使用make
	// identifier := map[key_type][value_type]{
	// 	"name": "bob",
	// 	"age": "20"
	// }
	m2 := map[string]string{
		"name":   "Bo",
		"gender": "M",
	}
	fmt.Printf("m2: %v\n", m2)

	m3 := make(map[string]string)
	m3["salary"] = "1000"
	fmt.Printf("m3: %v\n", m3)

	for k, v := range m3 {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

	m4 := map[string]string{
		"gender": "F",
		"job":    "Developer",
	}
	for k := range m4 {
		fmt.Printf("k: %v\t", k)
	}
	fmt.Println("")

	for _, v := range m4 {
		fmt.Printf("v: %v\t", v)
	}
	fmt.Println("")

}

func Sum(a int, b int) (sum int) {
	sum = a + b
	return sum
}

func f1() {
	fmt.Println("I am a function without return value")
}

func f2(a int) (ret int, b int) {
	ret = a + 1
	b = a + 2
	return ret, b
}

func f3(a int, b int) int {
	return a + b
}

func slice(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func changeable(input ...int) {
	for _, v := range input {
		fmt.Printf("v: %v\n", v)
	}
}

func changeable_02(name string, salary float32, other ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("salary: %v\n", salary)
	for _, v := range other {
		fmt.Printf("v: %v\n", v)
	}
}

func sayHello(name string) string {
	return name
}

func adv_f(name string, f func(string) string) string {
	return f(name)
}

func opr(operator string) func(int, int) int {
	switch operator {
	case "+":
		return func(a int, b int) int {
			return a + b
		}
	case "-":
		return func(a, b int) int {
			return a - b
		}
	default:
		return nil
	}
}

func isS(number int) bool {
	a, b, c := number/100, number%10, number/10%10
	if a*a*a+b*b*b+c*c*c == number {
		return true
	}
	return false
}

func allS() (res []int, count int) {
	result := []int{}
	for i := 100; i < 1000; i++ {
		if isS(i) {
			result = append(result, i)
		}
	}
	return result, len(result)
}

//闭包
func outer(y int) func(int) int {
	var x int = 100

	res_f := func(z int) int {
		return z + x + y
	}

	return res_f
}

func accu() func(int) int {
	var x int = 0
	return func(y int) int {
		x += y
		return x
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(a int) int {
		base -= a
		return base
	}

	return add, sub
}

func recuisive(a int) int {
	if a == 1 {
		return a
	}
	return a * recuisive(a-1)
}

func fibo(a int) int {
	if a == 1 || a == 2 {
		return 1
	}
	return fibo(a-1) + fibo(a-2)
}

func Func() {
	// func FuncName(args type, ...) (return_val, type, ...) {
	// }
	// f1()
	// ret, b := f2(3)

	// sum := f3(1, 2)
	// fmt.Printf("sum: %v\n", sum)

	// fmt.Printf("ret: %v\n", ret)
	// fmt.Printf("b: %v\n", b)

	// var input = []int{1, 2, 3, 4, 5, 6}
	// result := slice(input)
	// fmt.Printf("result: %v\n", result)

	// changeable(1, 3, 4, 5, 6)
	// changeable_02("Jack", 134.232, 1, 3, 4, 2, 4)
	// changeable_02("12", 132.23, 121, 12, 12121, 121, 4)

	// 预先定义一个函数类型的变量，此时funcName相当于int/float32/string/bool地位
	// 相当于给函数一个签名
	type funcName func(int, int) int

	// var f funcName = f3
	// fmt.Printf("f(1, 4): %v\n", f(1, 4))

	// f = Sum
	// fmt.Printf("f(3, 4): %v\n", f(3, 4))
	// fmt.Printf("adv_f(\"Tom\", sayHello): %v\n", adv_f("Tom", sayHello))

	// fmt.Printf("opr(\"+\")(1, 2): %v\n", opr("+")(1, 2))
	// fmt.Printf("opr(\"-\")(1, 2): %v\n", opr("-")(1, 2))

	// func(a int, b string) {
	// 	fmt.Printf("a: %v\n", a)
	// 	fmt.Printf("b: %v\n", b)
	// }(2, "ban")

	// fmt.Printf("isS(153): %v\n", isS(153))
	// al, count := allS()
	// fmt.Printf("al: %v\n", al)
	// fmt.Printf("count: %v\n", count)

	// ou := outer(100)
	// fmt.Printf("ou(10): %v\n", ou(10))

	// f := accu()
	// fmt.Printf("f(10): %v\n", f(10))
	// fmt.Printf("f(20): %v\n", f(20))
	// fmt.Printf("f(30): %v\n", f(30))
	// f_a, f_s := calc(10)
	// fmt.Println(f_a(1), f_s(2))
	// fmt.Println(f_a(3), f_s(4))
	// fmt.Println(f_a(5), f_s(6))

	// fmt.Printf("recuisive(6): %v\n", recuisive(6))
	fmt.Printf("fibo(1): %v\n", fibo(1))
	fmt.Printf("fibo(2): %v\n", fibo(2))
	fmt.Printf("fibo(10): %v\n", fibo(10))
}

func Pointer() {
	// var p1 *int
	// fmt.Printf("p1: %v\n", p1)
	// fmt.Printf("p1: %T\n", p1)

	// var i int = 100
	// p1 = &i
	// fmt.Printf("p1: %v\n", p1)
	// fmt.Printf("p1: %T\n", p1)
	// fmt.Printf("p1: %v\n", *p1)

	// var p2 *string
	// var s string = "hello"
	// p2 = &s
	// fmt.Printf("p2: %T\n", p2)
	// fmt.Printf("p2: %v\n", *p2)

	// var p3 *bool
	// var b bool = true
	// p3 = &b
	// fmt.Printf("p3: %T\n", p3)
	// fmt.Printf("p3: %v\n", *p3)

	var arr_p [5]*int
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		v := v
		arr_p[i] = &v
	}
	// for i := 0; i < len(arr); i++ {
	// 	arr_p[i] = &arr[i]
	// }

	for _, v := range arr_p {
		fmt.Printf("v: %T\n", v)
		fmt.Printf("v: %v\n", *v)
	}
}

func Type() {
	// 类型定义
	// 定义一个新的类型变量
	type MyInt int
	var i MyInt = 1
	fmt.Printf("i: %T\n", i)
	fmt.Printf("i: %v\n", i)

	// 类型别名
	// 没有产生新类型，就是加一个小名字
	type MyInt2 = int
	var j MyInt2 = 2
	fmt.Printf("j: %T\n", j)
	fmt.Printf("j: %v\n", j)

	var k int = 3
	fmt.Printf("k: %T\n", k)
	fmt.Printf("k: %v\n", k)

}

func Struct() {
	// 结构体定义，就是对象和类的概念
	// 需要用到上文的类型定义，但是多了一个struct关键字
	// type struct_type struct {
	// member definition;
	// member definition
	// }
	/*
		type: 结构体关键字
		struct_type: 定义的结构体名称，类似于类名
		struct； struct关键字
		member： 成员名
		definition：成员的类型
		若是成员名的类型相同，可以写到一行
			例如： member1， member2 definition
	*/
	// 创建结构体, 此后Person就可以当成一个类型，（和int，float，bool同等地位）
	type Person struct {
		name, email string
		id, age     int
	}

	// 实例化一个结构体
	var person Person
	// 或者
	person2 := Person{}
	fmt.Printf("person: %v\n", person)
	fmt.Printf("person2: %v\n", person2)

	// person.name = "P1"
	// person.age = 32
	// person.email = "p1@gmail.com"
	// person.id = 1

	// 初始化：按成员名字传入，键值对
	person = Person{
		name:  "jack",
		age:   32,
		email: "jack@gmail.com",
		id:    2,
	}

	// 初始化-02: 按定义时的顺序传入
	person = Person{
		"Jack",
		"Jack@gmail.com",
		3,
		23,
	}

	person2.name = "P2"
	person2.age = 31
	person2.email = "p2@gmail.com"
	person2.id = 2

	fmt.Printf("person: %v\n", person)
	fmt.Printf("person2: %v\n", person2)

}
