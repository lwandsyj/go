1. defer 语句会将函数推迟到外层函数返回之后执行。

        package main

        import "fmt"

        func main() {
            defer fmt.Println("world")

            fmt.Println("hello")
        }

        hello
        world

2. 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用

   在下面示例中，当推迟Println调用时，将评估表达式“ i”。函数返回后，延迟调用将显示“ 0”。

        func a() {
            i := 0
            defer fmt.Println(i)
            i++
            return
        }

        // 返回0


3. ***推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。***

        func b() {
            for i := 0; i < 4; i++ {
                defer fmt.Print(i)
            }
        }
        // 返回 3210