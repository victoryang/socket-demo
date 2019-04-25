# syslog日志系统
## Linux系统提供了接口来访问系统日志
- openlog()
- syslog() or vsyslog()
- closelog()

>其中openlog()和closelog()可选;
>
>openlog可以在每条log记录上追加类似进程名的信息;
>
>closelog用来关闭与syslog守护进程通信的socket的通信描述符。

## syslog配置
配置信息存储在/etc/syslog.conf中

> 类型.级别 [；类型.级别] `TAB` 动作
>
> 类型：信息产生的源头，如auth、cron等
>
> 级别：log级别
>
> 动作: 消息目的地，如主机、用户等
