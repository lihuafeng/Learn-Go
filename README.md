# Learn-Go
Go语言笔记







for range用法
---
for-range支持以下数据类型：
    - array
    - slice
    - string
    - map
    - channel（注意，只写的channel不允许range，例chan<- int）
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
```
