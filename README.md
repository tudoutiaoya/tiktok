# tiktok



## 基础接口

### 视频流接口

校验参数

按照时间倒叙查询(滑动分页)

视频信息封装加上作者信息

如果登录加上是否关注，是否点赞



待完善：

推拉模式



### 用户注册

校验参数

判断用户是否存在

Bcriypt加密

保存用户

返回token





### 用户登录

校验参数

查询用户

判断用户是否存在

判断密码是否正确

返回token和用户id



### 用户信息

验证参数

查询用户

判断用户是否存在

判断密码是否正确

返回用户信息





### 投稿接口

登录拦截：根据token

校验参数

获取视频数据

上传到七牛云

解析封面

保存视频到数据库

// 思考？是否文件秒传、断点续传？？？



待完善：

视频流不能有自己的视频

给关注用户发通知？？？





### 发布列表

登录拦截

查询发布列表

视频信息封装加上作者信息

封装是否点赞



## 互动接口

### 赞操作

查看视频是否存在

如果是点赞

​	判断是否已经点过赞

如果是取消点赞



### 喜欢列表

登录拦截

查询视频id

查询视频列表

添加是否关注，是否点赞



待完善：

改成子查询



### 评论操作

登录拦截

校验参数

如果是评论

​	保存

如果是删除

​	校验评论id

​	删除评论



### 评论列表

查询评论列表

添加作者信息





todo:

是否添加关注？







## 社交接口

### 关注操作

登录校验

如果是关注

​	不能关注自己

​	判断关注用户存在不存在

​	判断是否关注过

​	关注

如果是取消关注

​	取消关注





### 关注列表

登录拦截

查看关注用户列表



### 粉丝列表

> 关注你的

登录拦截

查看粉丝列表



### 好友列表

登录拦截

查看好友列表



todo:

添加最新消息



### 发送消息

登录校验

雪花算法生成全局唯一id

发送消息到对方消息邮箱(邮箱是指redis中的list)

mq异步保存到数据库



### 聊天记录

登录校验

去拉取自己的消息邮箱







todo:

待完善很多











































