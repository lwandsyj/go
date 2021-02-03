1. 在 go 语言中使用map 表示对象或者字典

        map[keyType]valueType

        a:=map[string]int{"name":1,"age":18}

        表示key 是string类型的，value 是int 型的

2. map 中key 如果是字符串要用双引号括起来，不然回去查找变量，避免造成错误

        d := map[string]string{name: "1", age: "2"}

        ./fn.go:40:25: undefined: name
        ./fn.go:40:36: undefined: age

3. 使用len 返回map 的key 的个数

         d := map[string]string{"name": "1", "age": "2"}
        fmt.Println(len(d))

4. 删除某个属性

        // delete 从map 中删除key 对象的value
        // 定义： func delete(m map[kType]vType,key Type)
        func learnDelete() {
            c := make(map[string]string)
            c = map[string]string{"name": "张三", "age": "14", "pwd": "pwd"}
            fmt.Println(c) //map[age:14 name:张三 pwd:pwd]
            // 删除pwd key
            delete(c, "pwd") //map[age:14 name:张三]
            fmt.Println(c)
        }