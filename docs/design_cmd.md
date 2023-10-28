该文档主要用于设计cmd相关，相当于草稿

## 前言

打算编写一个名为`mongo-cli`或者名为`mongoplus-plus`的工具。

特性:

- 简单易用，大量语法糖(仓库初衷)
- 支持创建数据库、集合、索引
- 支持查询结果集并将结果输出至终端
- 运行原生命令(不确定mongodb的go官方驱动是否支持)
- 方便安装

## 技术栈

- 命令行工具使用[cobra](https://github.com/spf13/cobra)
- 逻辑实现采用[本仓库](https://github.com/here-Leslie-Lau/mongo-plus)

其他待补充

## 命令

这里只是设计，后续可能会改

### 大命令

```shell
mongo-cli help

create 创建数据库、集合、索引等
query 查询结果集, 条件方面的拼装目前还需要想想如何设计
TODO...
```

### 子命令

TODO

## 示例
