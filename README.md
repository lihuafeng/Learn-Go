# Learn-Go
Go语言笔记

#### 依赖管理工具--[govendor](https://github.com/kardianos/govendor)

似PHP的composer,Python的pip


go内置的数据结构
---
|类型	|底层类型|
| ------------- |:-------------:| 
|array	|普通数组|
|string	|底层为struct，包含长度len和一个指向底层数组的指针|
|slice	|底层为struct，包含长度len，容量cap和一个指向底层数组的指针|
|map	|指向一个struct|
|channel	|指向一个struct|


Slice用法
---

slice的底层：
```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

slice的底层结构由一个指向数组的指针ptr和长度len，容量cap构成，也就是说slice的数据存在数组当中。

slice的重要知识点
1. slice的底层是数组指针。
2. 当append后，slice长度不超过容量cap，新增的元素将直接加在数组中。
3. 当append后，slice长度超过容量cap，将会返回一个新的slice。

实例：
```
func main() {
    s := []int{1, 2, 3} // len=3, cap=3
    a := s
    s[0] = 888
    s = append(s, 4)

    fmt.Println(a, len(a), cap(a)) // 输出：[888 2 3] 3 3
    fmt.Println(s, len(s), cap(s)) // 输出：[888 2 3 4] 4 6
}
因为slice的底层是数组指针，所以slice a和s指向的是同一个底层数组，所以当修改s[0]时，a也会被修改。
当s进行append时，因为长度len和容量cap是int值类型，所以不会影响到a。
```
总结
1. 谨记slice的底层结构是指针数组，并且len和cap是值类型。
2. 使用cap观察append后是否分配了新的数组。
3. Go的函数传参都是值拷贝传递。


defer用法
---

* defer是什么？

	延迟执行，但是只作用于函数。

* defer函数的执行顺序

	被defer的函数会被放入一个栈中，所以是先进后出的执行顺序。而被defer的函数在函数reture之后执行。
	
* 执行recover

	被 defer 的函数在 return 之后执行，这个时机点正好可以捕获函数抛出的 panic，因而defer 的另一个重要用途就是执行 recover ，而 recover 也只有在 defer 中才会起作用。
	
	```
	//recover 要放在 panic 点的前面，一般放在函数的起始的位置就可以了
	func test() {
	    defer func() {
		if ok := recover(); ok != nil {
		    fmt.Println("recover")
		}
	    }()
	    panic("error")
	}
	```
* defer与return的关系

	也就是说 return 语句不是原子操作，它被拆成了两步
	rval = xxx // 返回值赋值给rval
	ret // 函数返回
	而 defer 语句就是在这两条语句之间执行，也就是
	rval = xxx // 返回值赋值给rval
	defer_func  // 执行defer函数
	ret // 函数返回
	
实例：
```
func f1() {
    for i := 0; i < 5; i++ {
        defer fmt.Println(i)
    }
}
// 因为defer的调用是先进后出的顺序
// 所以输出：5, 4, 3, 2, 1


func f2() {
    for i := 0; i < 5; i++ {
    	defer func() {
    	    fmt.Println(i)
    	}()
    }
}
// 上面说到，i是一个闭包引用
// 所以当执行defer时，i已经是5了
// 所以输出：5，5，5，5，5


func f3() {
    for i := 0; i < 5; i++ {
    	defer func(n int) {
    	    fmt.Println(n)
    	}(i)
    }
}
// Go的函数参数是值拷贝，所以这是普通的函数传值
// 所以输出：5，4，3，2，1


func f4() int {
    t := 5
    defer func() {
    	t++
    }()
    return t
}
// 注意：f4函数的返回值是没有声明变量的
// 所以t虽然是闭包引用，但返回值rval不是闭包引用
// 可以拆解为
// rval = t
// t++
// return rval
// 所以输出是5


func f5() (r int) {
    defer func() {
        r++
    }()
    return 0
}
// 注意：f5函数的返回值是有声明变量的
// 所以返回值r是闭包引用
// 可以拆解为
// r = 0
// rval = r
// r++
// return rval
// 所以输出：1


func f6() (r int) {
    t := 5
    defer func() {
    	t = t + 5
    }()
    return t
}
// 这里t虽然是闭包引用，但返回值r不是闭包引用
// 可以拆解为
// r = t
// rval = r
// t = t + 5
// return rval
// 所以输出：5


func f7() (r int) {
    defer func(r int) {
	r = r + 5
    }(r)
    return 1
}
// 因为匿名函数的参数也是r，所以相当于是
// 匿名函数的参数r = r + 5，不影响外部
// 所以输出：1
```
	
小结：

1. 谨记defer和return执行的顺序
2. 注意返回值是否为闭包引用
	

for range用法
---
<pre>
for-range支持以下数据类型：
    - array
    - slice
    - string
    - map
    - channel（注意，只写的channel不允许range，例chan<- int）
</pre>
```
for key := range v {
	fmt.Println(key)
}
for key,val := range v{
	fmt.Println(key,val)
}
for _,val :=range v{
	fmt.Println(val)
}

dict := make(map[string]string,4)
dict["one"] = "1"
dict["two"] = "2"
dict["three"] = "3"
for k, v := range dict {
	fmt.Println(k, v)
	dict["four"] = "4"
}
```
<pre>
在Go里：
    （1）在遍历map的时候，添加或删除元素是安全的
    （2）但添加的元素，在往后的遍历中不一定会出现

    我们知道，Go的map底层是用hashmap实现的，因此map的遍历使用迭代器实现的，而不是像array这种形式的遍历。
   而且在 Go maps in action 中提到，map的遍历顺序被随机化了，也就是说遍历的顺序将会不固定，那么在for循环里给map新添加的元素有可能在迭代器之前，这时不会出现在往后的遍历中，反之，则会出现在往后的遍历中。
</pre>

interface
---

interface 是一组 method 的集合，是 duck-type programming 的一种体现。不关心属性（数据），只关心行为（方法）。具体使用中你可以自定义自己的 struct，并提供特定的 interface 里面的 method 就可以把它当成 interface 来使用。下面是一种 interface 的典型用法，定义函数的时候参数定义成 interface，调用函数的时候就可以做到非常的灵活。


```
type Myinterface interface {
	Print()
}

func Fn(myinter Myinterface)  {
	
}

type Mystr struct {
	
}

func (m Mystr) Print()  {

}
func main() {
	var me Mystr
	Fn(me)
}
```



设计模式
---

* 装箱者模式


* 单例模式

Go可以使用sync.Once包很优雅的实现单例

```
// The zero value for Once is ready to use
var oSingle sync.Once
var single *myType

func getSingle() *myType {
	oSingle.Do(func(){ single = newMyType() })
	return single
}
```

上面的实现方式有以下优点：
1. 无论调用多少次Do，都只会调用一次newMyType()
2. getSingle调用非常高效

    使用装饰模式和单例模式，再配合sync.Once你可以巧妙且方便地将一个不线程安全的API转换成安全的API。
