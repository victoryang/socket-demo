# Assemble

[从汇编角度理解 golang 多值返回和闭包](http://www.360doc.com/content/16/1226/16/19227797_617826588.shtml)

[逃逸分析](https://blog.csdn.net/chenchongg/article/details/88113778)

[汇编](https://zhuanlan.zhihu.com/p/264871176)

## go tool

```bash
# 汇编
go tool compile -S test.go > test.s

# 逃逸分析
go build -gcflags '-m'
```