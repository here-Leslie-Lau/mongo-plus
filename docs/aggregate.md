# Aggregate Usage Documentation

English | [中文](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/aggregate_cn.md)

* [Call the Aggregate method to perform aggregation operations](#Call-the-Aggregate-method-to-perform-aggregation-operations)
* [Get the Match Stage](#Get-the-Match-Stage)
* [Get the Sort Stage](#Get-the-Sort-Stage)
* [Get the Limit Stage](#Get-the-Limit-Stage)
* [Get the Skip Stage](#Get-the-Skip-Stage)
* [Get the Unset Stage](#Get-the-Unset-Stage)
* [Get the Group Stage](#Get-the-Group-Stage)
  * [Get the Avg Stage](#Get-the-Avg-Stage)
  * [Get the Sum Stage](#Get-the-Sum-Stage)

## Call the Aggregate method to perform aggregation operations

```go
func (ch *Chain) Aggregate(des interface{}, stages ...bson.D) error
```

Request Parameter Explanation: des - Result set after method execution, stages - Stages within the aggregation operation.

This stage can directly invoke methods provided by mongo-plus to obtain, as will be explained in the following text. If not satisfied, custom input is also possible.

## Get the Match Stage

```go
func (ch *Chain) GetMatchStage(filed, val string) bson.D
```

Request Parameter Explanation: field - Field to match, val - Specific value to match.

example:

```go
// Query documents with the name "leslie" and save the result set into the "documents" container
matchStage := ch.GetMatchStage("name", "leslie")
ch.Aggregate(&documents, matchStage)
```

## Get the Sort Stage

```go
func (ch *Chain) GetSortStage(rules ...SortRule) bson.D
```

Request Parameter Explanation: rules - Specific set of sorting rules, refer to `mongo.SortRule` for details.

example:

```go
// First, query documents with the name "leslie", then sort the result set in descending order based on the "age" field, and finally save the result set into the "documents" container
matchStage := ch.GetMatchStage("name", "leslie")
sortStage := ch.GetSortStage(mongo.SortRule{Typ: mongo.SortTypeDESC, Field: "age"})
ch.Aggregate(&documents, matchStage, sortStage)
```

## Get the Limit Stage

```go
func (ch *Chain) GetLimitStage(num int64) bson.D
```

Request Parameter Explanation: num - Number of documents to limit.

example:

```go
// Query two documents
limitStage := ch.GetLimitStage(2)
ch.Aggregate(&documents, limitStage)
```

## Get the Skip Stage

```go
func (ch *Chain) GetSkipStage(num int64) bson.D
```

Request Parameter Explanation: num - Number of documents to skip.

example:

```go
// Skip the first two documents, query the following two documents, and save the result set into the "documents" container
skipStage := ch.GetSkipStage(2)
limitStage := ch.GetLimitStage(2)
ch.Aggregate(&documents, skipStage, limitStage)
```

## Get the Unset Stage

```go
func (ch *Chain) GetUnsetStage(fileds ...string) bson.D
```

Request Parameter Explanation: fields - Names of fields to ignore

example:

```go
unsetStage := ch.GetUnsetStage("name", "age")
ch.Aggregate(&documents, unsetStage)
```

## Get the Group Stage

```go
func (ch *Chain) GetGroupStage(groupFiled string, subStages ...bson.D) bson.D
```

Request Parameter Explanation: groupField - Name of the field to group by, subStages - Sub-stages, if needed, pass them in
tips: Currently, this library provides only a limited set of supported stages. You can also customize BSON input as needed

### Get the Avg Stage

```go
func (ch *Chain) GetAvgStage(calledFiled, filed string) bson.D
```

Request Parameter Explanation: calledField - Name of the field after calculating the average, field - Field to calculate the average of
This stage is a sub-stage of the "group" stage, so it's usually used in combination with the GetGroupStage and passed as subStages

example:

```go
// Group by "class," calculate the average age for each class, and save the final results into the "documents" container
avgStage := ch.GetAvgStage("age", "age")
groupStage := ch.GetGroupStage("class", avgStage)
ch.Aggregate(&documents, groupStage)
```

### Get the Sum Stage

```go
func (ch *Chain) GetSumStage(calledFiled, filed string) bson.D
```

Request Parameter Explanation: calledField - Name of the field after calculating the sum, field - Field to calculate the sum of
This stage is a sub-stage of the "group" stage, so it's usually used in combination with the GetGroupStage and passed as subStages

example:

```go
// Group by "class," calculate the total sum of ages for each class, name the total sum as "total_age," and save the final results into the "documents" container
sumStage := ch.GetSumStage("total_age", "age")
groupStage := ch.GetGroupStage("class", sumStage)
ch.Aggregate(&documents, groupStage)
```
