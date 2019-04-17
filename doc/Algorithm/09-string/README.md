# String

## 普通字符串的查找
    - 指针是否为空，否则返回
    - 判断str是否为'\0'，判断剩下来的字符串长度是否>=模板字符串的长度，只要一个不符合，函数结束运行
    - 依次比较字符串和模板字符串的内容，如果全部符合，返回；只要有一个不符合，break跳出，str加1，转上一步

    char* strstr() {
        if len(target) == 0 || len(sub) == 0 {
            fmt.Println(false)
            return
        }

        var length = len(sub)
        for len(target)>=length {
            var i int
            for i=0; i<length; i++ {
                if target[i]!=sub[i]{
                    break
                }
            }

            if i == length {
                fmt.Println(true)
                return
            }

            target = target[i+1:]
        }

        fmt.Println(false)
    }

## KMP
    