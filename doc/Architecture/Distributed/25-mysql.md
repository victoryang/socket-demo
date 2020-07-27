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

docker run -p 3306:3306 --name mysql \
-v /usr/local/docker/mysql/conf:/etc/mysql \
-v /usr/local/docker/mysql/logs:/var/log/mysql \
-v /usr/local/docker/mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
-d mysql:5.7