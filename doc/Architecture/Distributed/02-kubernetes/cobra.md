# Cobra

[github](github.com/spf13/cobra/cobra)

[example](https://www.cnblogs.com/sparkdev/archive/2019/05/14/10856077.html)

## 主要功能

- 简易的子命令行模式，如 docker run等
- 完全兼容 posix 命令行模式
- 嵌套子命令 subcommand
- 支持全局，局部，串联 flags
- 使用cobra 很容易生成应用程序和命令
- 自动生成详细 help 信息

## 概念

cobra 中有个重要的概念，分别是 commands，arguments 和 flags。其中commands代表行为，arguments就是命令行参数，flags 代表对行为的改变。

**APP COMMAND ARG --FLAG**
**APPNAME VERB NOUN --ADJECTIVE**

### Command

Command is the central point of the application. Each interaction that the application supports will be contained in a Command. A command can have children commands and optionally run an action.

### Flags

A flag is a way to modify the behavior of a command. Cobra supports fully POSIX-compliant flags as well as the Go flag package. A Cobra command can define flags that persist through to children commands and flags that are only available to that command.

Flag functionality is provided by the pflag library, a fork of the flag standard library which maintains the same interface while adding POSIX compliance.