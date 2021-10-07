# Mysql

[docker mysql](https://www.cnblogs.com/sablier/p/11605606.html)

[golang mysql](https://www.jianshu.com/p/9b5cd762e256)


## alter table 

alter table test1
add (name varchar2(30) default ‘无名氏’ not null,

age integer default 22 not null,

has_money number(9,2)

);

## docker

```
docker run -p 3306:3306 --name mysql \
-v /usr/local/docker/mysql/conf:/etc/mysql \
-v /usr/local/docker/mysql/logs:/var/log/mysql \
-v /usr/local/docker/mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
-d mysql:5.7
```

### Import sql

```
# mysql commandline
mysql -uroot -p123456 < xxx.sql

# or source
mysql> create database abc;      # 创建数据库
mysql> use abc;                  # 使用已创建的数据库 
mysql> set names utf8;           # 设置编码
mysql> source /home/abc/abc.sql  # 导入备份数据库
```