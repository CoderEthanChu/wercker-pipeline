# wercker-pipeline

### 何为wercker?
wercker是一家被oracle收购的一家初创公司的产品，提供容器化及微服务快速开发，部署的CI平台。

### 为什么用wercker?

* 云端，简单，免费
* 特别适合独立开发者，初创公司

### 20分钟构建自己专属部署流水线基础环境
*1.注册用户推荐直接使用github账户登陆，[wercker](https://app.wercker.com/)*

> wercker注册有一个坑：注册时候需要进行选择图片验证，但是朝内网络无法加载，需要科学上网

*2.在wercher上创建应用*

![创建应用](https://raw.githubusercontent.com/CoderEthanChu/wercker-pipeline/master/illustration/create.png)

*3.绑定github代码库*

![绑定代码库](https://raw.githubusercontent.com/CoderEthanChu/wercker-pipeline/master/illustration/bind.png)

*4.提交代码触发流水线*

![触发流水线](https://raw.githubusercontent.com/CoderEthanChu/wercker-pipeline/master/illustration/run.png)

*5.增加dockerhub推送*

![dockerhub](https://raw.githubusercontent.com/CoderEthanChu/wercker-pipeline/master/illustration/dockerhub.png)

> 推送dockerhub这一步有个坑，官方文档的registry地址需要修改
```yaml
- internal/docker-push:
    # specify the image to be pushed - this is the one we created earlier
    image-name: my-new-image
    username: $USERNAME # Registry username
    password: $PASSWORD # Registry password
    registry: https://hub.docker.com
    repository: $USERNAME/docker-build-golang
```

> 需要将https://hub.docker.com 改成 https://registry.hub.docker.com
```yaml
deploy:
  steps:
    - internal/docker-push:
        username: $USERNAME # Registry username
        password: $PASSWORD # Registry password
        cmd: /wercker-pipline
        tag: latest
        registry: https://registry.hub.docker.com
        repository: $USERNAME/wercker-pipeline
```