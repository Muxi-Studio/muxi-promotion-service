### Kubernetes部署
- 登上集群
- 将本目录的内容`clone`至服务器
- 修改`mps.deployment.yaml`的内容，设置容器的环境变量
- 运行下面的命令:

创建命名`promotion`空间,
``` bash
$ kubectl create ns promotion
```
在`promotion`命名空间中创建`deployment`和service,
``` bash
$ kubectl create -f k8s/ -n promotion
```
查看容器是否正常启动，
``` bash
$ kubectl get pod -n promotion
```
如果未正常启动，则运行`logs`,`describe`子命令进行debug...

--::==