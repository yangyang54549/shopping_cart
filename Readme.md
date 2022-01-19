---
### 简介
> 这是一个商城购物车模块

### 用法

1. 使用您喜欢的存储器（演示使用redis） 定义用户key和连接
> 	storage := new(shopping_cart.RedisDb).New("shopping_card:" + strconv.Itoa(int(uid)), redisDb)

2. 引入存储器
> 	cart := shopping_cart.New(storage)

3. 使用

保存
> cart.Save(data)

删除
> cart.Clean()

删除一行
> cart.CleanOne(rowId)

修改
> cart.Edit(rowId, num, isSelect)

获取
> cart.GetAll()

获取一行
> cart.Get(rowId)

### 扩展

实现storage存储器接口，并引入即可