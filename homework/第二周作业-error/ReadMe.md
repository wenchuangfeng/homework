我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？



```
应该wrap这个error，因为像sql.ErrNoRows这种标准库只会返回根因错误，而dao层是数据持久层，一般来说不做具体的业务处理，那么就需要wrap一下error，加上一些扩展信息，例如查询失败"query failed"。然后往上抛这个错误，交给业务层来对错误进行处理。demo代码见errordemo.go
```

