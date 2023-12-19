# 命令行工具(mongo-helper)使用文档

`mongo-helper`是一款用于操作MongoDB数据库或集合的工具。目前支持使用该工具进行创建索引的操作。

## 安装

```shell
go install github.com/here-Leslie-Lau/mongo-plus/cmd@latest
# 生成的二进制文件默认叫cmd, 需要手动改名为mongo-helper, 或者你可以按照你的喜好来改名
mv $GOPATH/bin/cmd $GOPATH/bin/mongo-helper
```


## 使用

```shell
Usage:
  mongo-helper [command]

Available Commands:
  create      该命令用于创建索引
  help        获取帮助信息
  init        该命令用于初始化连接mongodb的配置信息, 应该优先运行
```

### init

```
Usage:
  mongo-helper init [flags]

Flags:
      --addr string   mongodb连接的ip和端口 (默认 "localhost:27017")
      --db string     要连接的数据库 (默认 "test")
  -h, --help          help for init
      --p string      mongodb的密码 (默认 "root")
      --u string      mongodb的用户名 (默认 "root")
```

使用后会在家目录下生成一个`mongo-plus.json`文件，该文件用于存储连接mongodb的配置信息。

```shell
$ mongo-helper init                                                 
load mongodb success...
init ~/mongo-plus.json success...
```
