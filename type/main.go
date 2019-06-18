package main

import "fmt"

func main() {

	// A slice B数组


	var A [4]int =[4]int{1,2,3,4}
	fmt.Println(A)
	B :=A[0:3]
	fmt.Println(B)
	fmt.Println(&A[0]) // 0xc0000a0000
	fmt.Println(&B[0]) // 0xc0000a0000 肯定一样啊 因为B[0]这个操作就是他里面strcut指向的
	fmt.Println(&B) // &[1 2 3]
	fmt.Printf("A指针是 %p\n",&A) // 0xc00009c000
	fmt.Printf("B[0]指针是 %p\n",&B[0]) //  0xc00009c000
	fmt.Printf("B指针是 %p",&B) // 0xc0000a6000
	// 指向这个slice的内存指针跟A确实不一样 那肯定啊 内存得有帮slice开辟内存的吧
}

// go参数的传递就只有一种 值传递 可通过指针拷贝（实际就是值传递）到达引用传递的效果

// 如果重复调用某函数 而且传参数有一个是固定的 那么就用闭bao会优雅 f2('jpg',22),f2('jpg',11)

// go数组是直接指向内存空间 而不是指向地址所以数组是值类型
// slice是引用类型 竟然是引用那么就得有对应得引用对象 数组
// slice定义完后还不能使用需要让其引用到一个数组或者make一个空间
// 指针指向内存

// 指针放在堆里面可以共享
// 什么是引用类型 -- 可以理解成他们执行的是变量内存 比如说slice指向的是数组内存
// 在golang中只有三种引用类型它们分别是切片slice、字典map、管道channel。其它的全部是值类型，引用类型可以简单的理解为指针类型，它们都是通过make完成初始化
// map 是 key-value 数据结构，又称为字段或者关联数组。类似其它编程语言的集合，
// 在编程中是经常使用到
// map跟我们javascript语言的json对象很相似 但是map存放的数据类型不能不同一定要一样 所以通常用map+struct->js对象


// 引用类型跟指针做好区别 虽然实际是一样的 都是指向那块变量内存
// 但是语言区分类型是因为你赋值或者参数传递是以拷贝还是指向内存空间的 注意 指针也是指向内存空间的
// 注意在golang中只有三种引用类型它们分别是切片slice、字典map、管道channel。其它的全部是值类型，引用类型可以简单的理解为指针类型，它们都是通过make完成初始化
// 值类型说明传递或者赋值就是拷贝
// 你可能会很奇怪为什么我打印的map类型不是一个指针（地址），因为引用类型又不是指针，map里面有做了对应的底层处理让他可以指向变量内存
// 注意引用类型指向内存变量 值类型也是指向内存变量 （不指向他怎么拿到）？对
// 不同的是引用类型指向公共的 值都是自己开辟自己的 -- 好像就是放在栈还是放在堆








// --------------------------------------------- 看这里
/*
再总结一遍 golang的参数传递只有值传递也就是拷贝 即使你传个指针 也是拷贝了指针的

值类型：int、float、bool和string这些类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
(对啊这个变量指向内存 不然怎么拿到值，而我们所说的值类型只是说在一系列赋值或者参数传递是拷贝啊还是指向同一个内存这才是他们的区别 变量永远都是指向内存的 这对于你画内存析图很重要)
(很重要的一点 以后记得写在笔记里)
值类型的变量的值存储在栈中。当使用等号=将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝。可以通过 &i 获取变量 i 的内存地址。  值拷贝

引用类型：特指slice、map、channel这三种预定义类型。引用类型拥有更复杂的存储结构:(1)分配内存
(2)初始化一系列属性等一个引用类型的变量r1存储的是r1的值所在的内存地址（数字）
或内存地址中第一个字所在的位置，这个内存地址被称之为指针，这个指针实际上也被存在另外的某一个字中。  

两者的主要区别：拷贝操作和函数传参。

来看下slice的数据结构
slice 从底层来说，其实就是一个数据结构(struct 结构体) type slice struct {
ptr *[2]int
len int cap int

}

slice底层确实存了一个指针所以指向内存 但是他本身自己也是个内存也就是这样的工作形式的 A slice 这个变量你得有内存去存它吧 而它这个变量里面又存了个指针 指向B数组
我们要操作的就是这个B数组 其实本身A内存我们不会过多的关注

*/


// 指针也是种类型 指针类型吗 不然还能怎么样

// 引用类型 A->栈(地址)->堆   过程应该是这样的

// slice不管以什么方式声明 底层都是有先声明个数组供它引用的


// 务必记得 make 仅适用于 map，slice 和 channel，并且返回的不是指针。应当用 new 获得特定的指针。

// map跟slice 的数据结构还是类似 https://i6448038.github.io/2018/08/26/map-secret/

// 底层都是指针指向

