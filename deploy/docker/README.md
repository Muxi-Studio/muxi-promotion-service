### docker部署
构建程序，生成二进制文件
``` bash
$ go build
```

将二进制文件放置与本目录，配置本目录的`Dockerfile`并构建进镜像:
``` bash
$ docker build -t pqcsdockerhub/muxi-promotion-service-image .
```

将镜像传至服务器，然后启动:
``` bash
$ docker run -d --name muxi-promotion-service-ct -p 8080:8080 pqcsdockerhub/muxi-promotion-service-image
```

然后配置nginx反向代理，配置域名，配置HTTPS。部署即可完成