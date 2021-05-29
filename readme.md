# go mod 包管理,制作自己的包？

#channel 中的buffer 满了之后，写入时会阻塞么？

#多个case 多个buffer 如何处理

#select 和switch 的区别
select只能应用于channel的操作，既可以用于channel的数据接收，也可以用于channel的数据发送。  
如果select的多个分支都满足条件，则会随机的选取其中一个满足条件的分支， 如语言规范中所说  
switch可以为各种类型进行分支操作， 设置可以为接口类型进行分支判断(通过i.(type))

#* 和& 的区别，指针和值
&符号的意思是对变量取地址
*符号的意思是对指针取值
#benchmark 测试没有输出时间
go test -bench=.
#goconvey

#json

#http
内置的http server,路由规则，最长匹配原则
1. 
