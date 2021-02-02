1. 在 go 语言中，循环只有for 语句，没有while ,loop 等, 但是for 有很多变体

   在go语言中，for 语句不用必须使用() 包裹

+ 正常for 循环

        // 正常的三个语句
        for i:=0;i<10;i++ {
            fmt.Println(i)
        }

+ 省略第一个语句

        i:=0

        for ;i<10;i++ {

        }

+ 省略第一个和第三个语句

        // 实现while 的效果
        sum := 1
        for sum < 1000 {
            sum += sum
        }

+ for省略三个语句，死循环

        for{

        }

2. for range 用于循环字符串，数组，切片，map 等

+ 字符串，切片，数组返回index,value

        for index,value :=range arr{

        }

+ 循环map 返回key,value

        for key,value :=range map{
            
        }