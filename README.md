# Muxi promotion service(MPS)
[木犀](http://www.muxixyz.com)产品推广服务,**通用的“新带老”模式的互联网产品推广解决方案**。利用现有产品的用户群来推广你想要推广的任何页面。轻量级、微服务、可重用、与现有产品之间完全解耦。

### 一、Powered BY
- go/iris framework
- redis(数据持久化)
- go-redis(go语言的redis驱动)
- nginx(反向代理)
- let's encrypt(https证书颁发)
- 阿里云提供计算服务

### 二、推广模式简介
本推广模式适用于任何具有用户系统的互联网产品。利用它,我们可以策划一个推广活动,来推广指定的页面。

下面我们以学而的推广为例。我们现在想要策划一场推广活动,思路是让老用户来帮助我们宣传学而。我们事先在学而主页上放出本次推广活动的通知以及具体的活动规则。通用的活动规则大概是这样的:**首先学而注册用户获取本人专属推广链接，然后尽自己最大的可能让自己的专属链接得到更多的点击(用户可以发朋友圈，在各种群里面转发自己的链接或者自己刷点击等。该链接的点击次数会作为该用户为我们产品推广力度的唯一量化指标)。链接最后都被重定向到在获取链接时所指定的被推广页面(比如用户注册页面或者其他的我们想要让更多人看到的页面)**,这样就达到了推广产品的目的。

要想获得一个好的推广效果，必须**事先做好宣传,准备丰厚的奖品**。可以在学而中留出一个推广页面，实时的展示当前参与活动的用户的推广力度排行榜，还可以提供用户的当前排名查询等。我们可以承诺在某一截止时间时，给予榜单前几名的用户丰厚奖励。

理论上来讲，只要宣传做到位、奖品足够吸引人，这个推广活动是会有很好的效果的。

### 三、配置与部署
#### 环境配置
- REDIS_PASSWORD:redis密码，若不设置此变量表明无密码
- REDIS_ADDR:redis地址，默认为`localhost:6379`
- BASIC_AUTH_INFO:Basic Auth账户,默认为`andrewpqc:andrewpqc`
- SECRETKEY:生成和解析token的前面秘钥字符串，默认`fDEtrkpbQbocVxYRLZrnkrXDWJzRZMfO`,该字符串必须是32字节

#### 部署
[docker部署]()
[kubernetes部署]()
[二进制包部署]()
### 四、前端开发者指南
```
http://127.0.0.1:8080/api/v1.0/private-promotion-link?id=1&url=www.baidu.com&ex=1000
http://127.0.0.1:8080/promotion/?t=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEifQ.O4FApAv52Ue-HwMS5mHBaeOnhp_nMcTlANfCfCkOivM&landing=aHR0cHM6Ly93d3cuYmFpZHUuY29t"
```

### 四、API文档地址
https://app.swaggerhub.com/apis/andrewpqc/xueer-promotion/1.0.0

### 五、TODO
