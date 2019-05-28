# Learn-Go
Go语言笔记



go内置的数据结构
---
|类型	|底层类型|
| ------------- |:-------------:| 
|array	|普通数组|
|string	|底层为struct，包含长度len和一个指向底层数组的指针|
|slice	|底层为struct，包含长度len，容量cap和一个指向底层数组的指针|
|map	|指向一个struct|
|channel	|指向一个struct|




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
