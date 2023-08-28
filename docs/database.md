# Documentation for Database-Level Usage

- [Get Database Instance](#Get-Database-Instance)
- [Execute MongoDB database management commands](#Execute-MongoDB-database-management-commands)
- [Check the status of the database service](#Check-the-status-of-the-database-service)
- [Check if the server is available](#Check-if-the-server-is-available)
- [Get database statistics](#Get-database-statistics)
- [Get information about the server's status](#Get-information-about-the-server's-status)

## Get Database Instance

```go
func (c *Conn) GetDB() *mongo.Database
```

Calling this method allows you to obtain the `Database` instance of the official MongoDB Go driver

## Execute MongoDB database management commands

```go
func (c *Conn) RunCommand(ctx context.Context, cmd bson.D, des interface{}) error
```

Request parameter description: ctx - standard library context, cmd - command collection of type bson.D, des - result of command execution, pointer to a struct

## Check the status of the database service

```go
func (c *Conn) IsMaster(ctx context.Context, des interface{}) error
```

Request parameter description: ctx - standard library context, des - command execution result, pointer to a structure

This command is used to check the status of the server, especially to determine whether the server is a primary node, a replica set, and other basic information about the server. This command can help you understand the current status and topology of the MongoDB deployment

## Check if the server is available

```go
func (c *Conn) Ping(ctx context.Context) error
```

Request parameter description: ctx - standard library context

Used to check if the server is available

## Get database statistics

```go
func (c *Conn) DbStats(ctx context.Context, des interface{}) error
```

Request parameter description: ctx - standard library context, des - command execution result, pointer to a structure

Used to retrieve statistics about the database. The information returned by this command includes the size of the database, the number of objects, the number of indexes, and other statistical data

## Get information about the server's status

```go
func (c *Conn) ServerStatus(ctx context.Context, des interface{}) error
```

Request parameter description: ctx - standard library context, des - command execution result, pointer to a structure

Used to retrieve information about the server's status. This command returns information such as the server's version, operating system, memory usage, connection count, and other statistical data
