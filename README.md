# Go_Shop 服务端

## 🐱‍🏍简介

​	该项目为Go_shop项目的服务端，服务被拆分成了5个独立的微服务，分别为：

​		`goods_servcie`：商品服务

​		`inventory_servcie`：库存服务

 	   `order_service`：订单服务(内含购物车服务)

​		`user_service`：用户服务

​		`userop_service`：用户留言收藏服务

​	微服务定位：位于对外暴露接口的Web层和数据库之间，主要负责对数据库中信息的查询以及分布式问题的解决。不涉及具体业务内容。仅对外体统服务支持。同时使用注册中心，完成了负载均衡，服务注册，减轻单个微服务的运行压力。

## 🥽技术选择

​		微服务框架—`GRPC`

​		数据库操作—`GORM`

​		服务注册与负载均衡—`consul`

​		配置中心—`Nacos`

​		数据库—`MySQL`

​		搜索：`ElasticSearch`

## 🚧系统架构总览

![image-20221029174645592](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221029174645592.png)

​		每个微服务拥有一个独立的数据库，服务启动，自动将服务注册到conusl中。

​		微服务之间的调用流程 (服务A 调用 服务B 中的服务)：

​			1.A服务启动注册到consul中

​			2.A服务在consul中检查B服务是否注册到consul中

​			2.1 B服务已注册— A服务获取B服务注册IP与端口 建立连接 调用B服务

​			2.2 B服务未注册— A服务提示报错 B服务未注册 持续监听consul中B服务的注册

​		各个微服务的配置，统一由Nacos管理。

![image-20221029175257002](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221029175257002.png)

![image-20221029175348813](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221029175348813.png)