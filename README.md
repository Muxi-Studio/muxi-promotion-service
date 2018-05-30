# Muxi promotion service(MPS)
[木犀](http://www.muxixyz.com)产品推广服务,**通用的“新带老”模式的互联网产品推广解决方案**。利用现有产品的用户群来推广你想要推广的任何页面。轻量级、微服务、可重用、与现有产品之间完全解耦。

### 一、Powered BY
- go/[iris](https://github.com/kataras/iris) framework
- [redis](https://redis.io/)(数据持久化)
- [go-redis](https://github.com/go-redis/redis)(go语言的redis驱动)
- [nginx](http://nginx.org/)(反向代理)
- [let's encrypt](https://letsencrypt.org/)(https证书颁发)
- [阿里云](https://cn.aliyun.com/?utm_content=se_980105&gclid=Cj0KCQjw9LPYBRDSARIsAHL7J5l2CnX6oYbFSvzhFnnsZOrEoaPWnfB8Nc1m_hH7y35-NUypq847NxAaArl8EALw_wcB)提供计算服务

### 二、推广模式简介
本推广模式适用于任何具有用户系统的互联网产品。利用它,我们可以策划一个推广活动,来推广指定的页面。

下面我们以学而的推广为例。我们现在想要策划一场推广活动,思路是让老用户来帮助我们宣传学而。我们事先在学而主页上放出本次推广活动的通知以及具体的活动规则。通用的活动规则大概是这样的:**首先学而注册用户获取本人专属推广链接，然后尽自己最大的可能让自己的专属链接得到更多的点击(用户可以发朋友圈，在各种群里面转发自己的链接或者自己刷点击等。该链接的点击次数会作为该用户为我们产品推广力度的唯一量化指标)。链接最后都被重定向到在获取链接时所指定的被推广页面(比如用户注册页面或者其他的我们想要让更多人看到的页面)**,这样就达到了推广产品的目的。

要想获得一个好的推广效果，必须**事先做好宣传,准备丰厚的奖品**。可以在学而中留出一个推广页面，实时的展示当前参与活动的用户的推广力度排行榜，还可以提供用户的当前排名查询等。我们可以承诺在某一截止时间时，给予榜单前几名的用户丰厚奖励。

理论上来讲，只要宣传做到位、奖品足够吸引人，这个推广活动是会有很好的效果的。

### 三、配置与部署
#### 环境变量配置
- REDIS_PASSWORD:redis密码，若不设置此变量表明无密码
- REDIS_ADDR:redis地址，默认为`localhost:6379`
- BASIC_AUTH_INFO:Basic Auth账户,默认为`andrewpqc:andrewpqc`
- SECRETKEY:生成和解析token的前面秘钥字符串，默认`fDEtrkpbQbocVxYRLZrnkrXDWJzRZMfO`,该字符串必须是32字节

#### 部署
[kubernetes部署](https://github.com/Andrewpqc/xueer-promotion-service/blob/develop/deploy/k8s/README.md)

[docker部署](https://github.com/Andrewpqc/xueer-promotion-service/blob/develop/deploy/docker/README.md)

[二进制包部署](https://github.com/Andrewpqc/xueer-promotion-service/blob/develop/deploy/binary/README.md)

### 四、前端开发者指南
首先发送GET请求到下面的URL
```
https://promotion.andrewpqc.xyz/private-promotion-link/
```
该请求有下面的查询参数:`id`,`url`,`ex`。这三个参数分别的含义是用户标识，需要推广的页面url和链接有效期。其中前两个参数为必须参数，后面的`ex`为可选参数，ex的单位为秒。如不传`ex`表明不给生成的链接设置有效期，返回的链接永久有效。
下面是两个示例:
```
https://promotion.andrewpqc.xyz/api/v1.0/private-promotion-link/?id=1&url=www.baidu.com&ex=1000
```
上面的链接就是获取用户标识为1,被推广URL为百度首页，过期时间为1000秒的推广链接。
```
https://promotion.andrewpqc.xyz/api/v1.0/private-promotion-link/?id=1&url=www.baidu.com
```
这个链接就是获取用户标识为1,被推广URL为百度首页，永久有效的推广链接。

下面是请求返回的示例推广链接
```
http://promotion.andrewpqc.xyz/promotion/?t=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEifQ.O4FApAv52Ue-HwMS5mHBaeOnhp_nMcTlANfCfCkOivM&landing=aHR0cHM6Ly93d3cuYmFpZHUuY29t"
```
随后用户可以分发此链接，此链接会将请求发送至本推广服务，我们就会将数据库对应数据更新，并且将请求重定向至被推广页面。如果该链接设置了过期，那么这个handler还会检查该链接是否过期，如果过期则数据库不会更改，请求也不会被重定向，程序返回相应状态码提示链接过期。

**给推广链接设置过期时间有啥好处?**
设置过期时间，用户必须在一段时间后重新获取专属推广链接并且重新分发。这对于增加推广效果有益。(前提是奖品非常诱人，用户非常有刷榜的欲望，不然的话，链接过期他就真的就不玩儿了:)

更多其他的API请参考下面的API文档。

### 四、API文档
https://app.swaggerhub.com/apis/andrewpqc/xueer-promotion/1.0.0

### 五、TODO

