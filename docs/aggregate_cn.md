# aggregate(聚合操作)使用文档

[English](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/aggregate.md) | 中文

* [调用Aggregate方法完成聚合操作](#调用Aggregate方法完成聚合操作)
* [获取match stage](#获取match-stage)
* [获取sort stage](#获取sort-stage)
* [获取limit stage](#获取limit-stage)
* [获取skip stage](#获取skip-stage)
* [获取unset stage](#获取unset-stage)
* [获取group stage](#获取group-stage)
  * [获取avg stage](#获取avg-stage)
  * [获取sum stage](#获取sum-stage)

## 调用Aggregate方法完成聚合操作

```go
func (ch *Chain) Aggregate(des interface{}, stages ...bson.D) error
```

请求参数说明: des-方法执行后的结果集, stages-聚合操作内的stage(阶段)

该stage可以调用mongo-plus提供的方法直接获取, 下文会说明。如果不满足也可以自定义传入

## 获取match stage

```go
func (ch *Chain) GetMatchStage(filed, val string) bson.D
```

请求参数说明 filed: 要匹配(match)的字段, val: 具体要匹配的值

example:

```go
// 查询name为leslie的文档，并将结果集保存至documents内
matchStage := ch.GetMatchStage("name", "leslie")
ch.Aggregate(&documents, matchStage)
```

## 获取sort stage

```go
func (ch *Chain) GetSortStage(rules ...SortRule) bson.D
```

请求参数说明 rules: 具体的排序规则集合, 可参考mongo.SortRule

example:

```go
// 先查询name为leslie的文档, 再按照age字段倒序, 最后将结果集保存至documents内
matchStage := ch.GetMatchStage("name", "leslie")
sortStage := ch.GetSortStage(mongo.SortRule{Typ: mongo.SortTypeDESC, Field: "age"})
ch.Aggregate(&documents, matchStage, sortStage)
```

## 获取limit stage

```go
func (ch *Chain) GetLimitStage(num int64) bson.D
```

请求参数说明 num: 要限制的文档数

example:

```go
// 查询2条文档
limitStage := ch.GetLimitStage(2)
ch.Aggregate(&documents, limitStage)
```

## 获取skip stage

```go
func (ch *Chain) GetSkipStage(num int64) bson.D
```

请求参数说明 num: 要跳过的文档数

example:

```go
// 跳过前两条文档, 查询之后的两条文档, 并将结果集保存至documents内
skipStage := ch.GetSkipStage(2)
limitStage := ch.GetLimitStage(2)
ch.Aggregate(&documents, skipStage, limitStage)
```

## 获取unset stage

```go
func (ch *Chain) GetUnsetStage(fileds ...string) bson.D
```

请求参数说明 fileds: 要忽略的字段名

example:

```go
unsetStage := ch.GetUnsetStage("name", "age")
ch.Aggregate(&documents, unsetStage)
```

## 获取group stage

```go
func (ch *Chain) GetGroupStage(groupFiled string, subStages ...bson.D) bson.D
```

请求参数说明 groupFiled: 要分组的字段名, subStages: 子stage, 如果需要则传入
tips: 目前该库只提供少量stage支持, 也可以自定义bson传入

### 获取avg stage

```go
func (ch *Chain) GetAvgStage(calledFiled, filed string) bson.D
```

请求参数说明 calledFiled: 计算出平均值之后的字段命名, filed: 要计算平均值的字段
该stage为group的子stage, 故一般与GetGroupStage组合使用, 作为subStages传入

example:

```go
// 根据class分组, 计算每个class的age平均值, 最终结果保存至documents内
avgStage := ch.GetAvgStage("age", "age")
groupStage := ch.GetGroupStage("class", avgStage)
ch.Aggregate(&documents, groupStage)
```

### 获取sum stage

```go
func (ch *Chain) GetSumStage(calledFiled, filed string) bson.D
```

请求参数说明 calledFiled: 计算出总和之后的字段命名, filed: 要计算总和的字段
该stage为group的子stage, 故一般与GetGroupStage组合使用, 作为subStages传入

example:

```go
// 根据class分组, 计算每个class的age总和, 总和命名为total_age, 最终结果保存至documents内
sumStage := ch.GetSumStage("total_age", "age")
groupStage := ch.GetGroupStage("class", sumStage)
ch.Aggregate(&documents, groupStage)
```
