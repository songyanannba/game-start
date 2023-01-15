# protoc 
    protoc --go_out=. *.proto

# 分布式架构 适用于 web(兼容gin) ； 或者长链接的服务（websocket）

    client 客户端调用程序
    gateway 网管层
    player and room 服务层

# 架构描述 分布式
    客户端（gateway） -> nats(发布订阅) -> 服务端(player)

    

  