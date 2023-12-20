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

### create

```
$ mongo-helper create --help                                        
该命令目前仅用于创建索引

Usage:
  mongo-helper create [flags]

Flags:
      --coll string      要操作的集合名字
  -h, --help             帮助信息
      --indexs strings   创建索引的字段格式为 '索引列_排序规则(1或-1)', 例如: 'name_1', 'age_-1', 'name_1_age_-1'
      --ope string       需要执行的操作, 目前仅支持index(索引)
```

#### index

在demo集合内创建一个复合索引, 索引字段为name和age, name为降序, age为升序

```shell
$ mongo-helper create --coll=demo --ope=index --indexs=name_-1,age_1
load mongodb success...
Index created successfully, index name: name_-1_age_1
```

该命令读取的连接信息为`~/mongo-plus.json`文件中的配置信息, 如果你想使用其他配置信息, 可以使用`--addr`, `--db`, `--u`, `--p`参数来指定

## 问题反馈

直接提issue, 我会尽快回复
