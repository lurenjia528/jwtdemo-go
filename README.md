# jwtdemo-go

测试说明:

mongo命令查看

``` bash
# 显示所有数据库
show dbs
# 切换数据库
use Movies
# 删除数据库
db.dropDatabase()
# 显示所有表
show tables
# 全表扫描 db.(show tables 结果).find()
db.UserModel.find()
```

查询时,header为 authorization: token  中间没有Bearer(也可以自行处理)

注册

![1](./testimg/注册.png)

登录

![1](./testimg/登录.png)

上传

![1](./testimg/上传.png)

获取所有

![1](./testimg/获取所有.png)
