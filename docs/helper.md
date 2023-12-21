# Command-line tool (mongo-helper) usage documentation.

`mongo-helper` is a tool for operating MongoDB databases or collections. Currently, it supports operations for creating indexes using this tool.

## Installation

```shell
go install github.com/here-Leslie-Lau/mongo-plus/cmd@latest
# The generated binary file is named "cmd" by default. You need to manually rename it to "mongo-helper", or you can rename it according to your preference.
mv $GOPATH/bin/cmd $GOPATH/bin/mongo-helper
```

## Usage

```shell
Usage:
  mongo-helper [command]

Available Commands:
  create      This command is used to create indexes, etc.
  help        Help about any command
  init        The command is used to initialize the configuration file, which is used for connecting 
to MongoDB. Should be used before other commands.
```

### init

```
Usage:
  mongo-helper init [flags]

Flags:
      --addr string   The address of MongoDB (default "localhost:27017")
      --db string     The database of MongoDB (default "test")
  -h, --help          help for init
      --p string      The password of MongoDB (default "root")
      --u string      The username of MongoDB (default "root")
```

After usage, a `mongo-plus.json` file will be generated in the home directory. This file is used to store the configuration information for connecting to MongoDB.

```shell
$ mongo-helper init                                                 
load mongodb success...
init ~/mongo-plus.json success...
```

### create

```
$ mongo-helper create --help                                        
This command is used to create indexes, etc.

Usage:
  mongo-helper create [flags]

Flags:
      --coll string      The name of the collection to be operated
  -h, --help             help for create
      --indexs strings   The format for the fields to create an index is 'field_name_sorting_rule(1 o
r -1)', such as: 'name_1', 'age_-1', 'name_1_age_-1'
      --ope string       The name of the operation to be performed, now only index
```

#### index

Create a compound index within the demo collection with index fields "name" and "age," where "name" is in descending order, and "age" is in ascending order.

```shell
$ mongo-helper create --coll=demo --ope=index --indexs=name_-1,age_1
load mongodb success...
Index created successfully, index name: name_-1_age_1
```

The connection information read by this command is the configuration stored in the `~/mongo-plus.json` file. If you want to use different configuration information, you can specify it using the `--addr`, `--db`, `--u`, `--p` parameters.

## Troubleshooting

Feel free to raise an issue, and I will respond as soon as possible
