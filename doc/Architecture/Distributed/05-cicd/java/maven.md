# Maven

[maven](https://www.runoob.com/maven/maven-tutorial.html)

## Conventional Configuration

|Directory|Function|
|-|-|
|${basedir}|存放 pom.xml 和所有的子目录|
|${basedir}/src/main/java| 项目的 java 源代码|
|${basedir}/src/main/resource| 项目的资源，如 property, springmvc.xml|
|${basedir}/src/test/java|项目的测试类，如 Junit 代码|
|${basedir}/src/test/resources|测试用的资源|
|${basedir}/src/main/webapp/WEB-INF|web 应用文件目录|
|${basedir}/target|打包输出目录|
|${basedir}/target/classes|编译输出目录|
|${basedir}/target/test-classes|测试编译输出目录|
|Test.java|Maven 只会自动运行负荷该命名规则的测试类|
|~/.m2/repository|Maven 默认的本地仓库目录位置|

## Environment Configuration

https://www.runoob.com/maven/maven-setup.html

```bash
sudo vim /etc/profile

export MAVEN_HOME=/usr/local/apache-maven-3.3.9
export PATH=${PATH}:${MAVEN_HOME}/bin

source /etc/profile

mvn -v
```