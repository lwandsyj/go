1. fmt 主要包含print 和 scan 

+ print包括print/printf/println： 输出到终端

   + print: 输入不换行
   + printf: 带有格式化的输出  %d-> 数字,%s->字符串 %T->变量类型
   + println: 换行输出

+ Fprint/Fprintf/Fprintln 函数接受两个参数，第一个参数是一个io.writer 

+ Sprint/Sprintf/Sprintln: 只是格式化内容，返回字符串，不输出到某处

2. 格式化

+ %d-> 数字
+ %o-> 八进制
+ %x-> 十六进制，字母小写 a-f
+ %X-> 十六进制，字母大写 A-F
+ %f-> 浮点数

    %9.2f    width 9, precision 2 

+ %c-> 相应Unicode码所表示的字符
+ %U-> unicode

+ %b-> 二进制
+ %s-> 字符串
+ %v-> 默认
+ %p-> 指针

+ 参数占位符

         fmt.Printf("%[2]d, %[1]d\n", 11, 22)  //22, 11，先输出第二个值，再输出第一个值


3. 复合对象

        struct:            {field0 field1 ...} 
        array, slice:      [elem0 elem1 ...] 
        maps:              map[key1:value1 key2:value2] 
        pointer to above:  &{}, &[], &map[]