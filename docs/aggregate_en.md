# Aggregate Usage Documentation

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
