# Code Generator

[code-generator](https://blog.csdn.net/zhonglinzhang/article/details/105022125)

## 生成clientset

>../../../k8s.io/code-generator/generate-groups.sh client,informer,lister github.com/REPONAME/PROJECTNAME/pkg/client github.com/REPONAME/PROJECTNAME/pkg/apis GROUP:v1alpha1

命令的第一段是脚本路径, 需要用code-generator相对与operator工程项目的相 对路径, 第二个参数表示要生成什么, 我们这里只需要 client/informer/lister, 至于deepcopy, 由于operator-sdk已经帮我们生成了, 所以没有用code-generator生成. 第三个参数是生成的代码要放到什么路径. 注 意这里不是相对operator工程的相对路径, 而是类似golang里import的路径, 也 就是相对于GOPATH的路径. 第四个参数是apis所在的路径, 同样需要是相对于 GOPATH的路径. 最后一个参数是Group+Version, 用冒号分割. code-generator 就会去第四个参数拼上Group+Version后的路径, 也就是 pkg/apis/GROUP/VERSION这个目录下去找CRD的定义.

