# Maven

[Maven教程](https://www.runoob.com/maven/maven-tutorial.html)

## 基础

### Maven 功能

- 构建
- 文档生成
- 报告
- 依赖
- SCMs
- 发布
- 分发
- 邮件列表

### 约定配置

|目录|目的|
|:-|:-|
|${base}|存放pom.xml和所有的子目录|
|${basedir}/src/main/java|项目的源代码|
|${basedir}/src/main/resources|项目的资源，比如说property文件，springmvc.xml|
|${basedir}/src/test/java|项目的测试类，比如说Junit代码|
|${basedir}/src/test/resources|测试用的资源|
|${basedir}/src/main/webapp/WEB-INF|web应用文件目录，web项目的信息，比如存放web.xml、本地图片、jsp视图页面|
|${basedir}/target|打包输出目录|
|${basedir}/target/classes|编译输出目录|
|${basedir}/target/test-classes|测试编译输出目录|
|Test.java|Maven只会自动运行符合该命名规则的测试类|
|~/.m2/repository|Maven默认的本地仓库目录位置|

### Maven的特点

- 项目构建遵循统一的规则
- 任意工程中共享
- 依赖管理包括自动更新
- 一个庞大且不断增长的库
- 可扩展，能够轻松编写JAVA或脚本语言的插件
- 只需要很少或不需要额外配置即可即时访问新功能
- **基于模型的构建** - Maven能够将任意数量的项目构建到预定义的输出类型中，如 JAR，WAR或基于项目元数据的分发，而不需要在大多数情况下执行任何脚本
- **项目信息的一致性站点** - 使用与构建过程相同的元数据，Maven都够生成一个网站或PDF，包括您要添加的任何文档，并添加到关于项目开发状态的标准报告中
- **发布管理和发布单独的输出** - Maven将不需要额外的配置，就可以与源代码管理系统集成，并可以基于某个标签管理项目的发布。它也可以将其发布到分发位置供其它项目使用。Maven 能够发布单独的输出，如 JAR，包含其它依赖和文档的归档，或者作为源代码发布。
- **向后兼容性** - 你可以很轻松的从旧版本 Maven的多个模块移植到 Maven 3
- 子项目使用父项目依赖时，正常情况子项目应该继承父项目依赖，无需使用版本号
- 并行构建 - 编译速度能够普遍提高 20-50%
- 更好的错误报告 - Maven 改进了错误报告，它为你提供了 Maven wiki页面的链接，你可以点击链接查看错误的完整描述