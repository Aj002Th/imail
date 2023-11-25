# iMail

## 项目介绍
iMail是一个基于go实现的信息获取和推送工具，支持邮件推送爬虫的结果。  
可以通过爬虫爬取相应的信息, 例如视频的更新信息等, 并通过邮件推送给用户。

## 项目结构
```text
.
├─cmd     # 命令行工具
├─common  # 工具包
├─data    # 用户数据 
└─server  # 业务模块
    ├─catcher   # 爬虫模块
    ├─manager   # 管理器模块
    └─messager  # 消息推送模块
```

## 项目依赖
- go 1.21+
  1. 参考官方文档 https://golang.google.cn/doc/install
- 无头浏览器
  1. 项目使用 go-rod 库操作无头浏览器实现爬虫, 在绝大多数操作系统下, 项目首次启动时能自动下载和安装合适版本的无头浏览器, 在某些平台上，您可能需要手动安装浏览器，Rod 无法保证自动下载的浏览器始终有效
  2. 若自动下载失败, 可参考 https://go-rod.github.io/#/compatibility?id=compatibility 进行手动下载安装

## 项目启动
1. 下载项目代码
    ```shell
    git clone https://github.com/Aj002Th/imail.git
    ```
2. 安装第三方依赖
    ```bash
    cd imail
    go mod tidy
    ```
3. 编写配置文件  
   - 在项目根目录创建 imail.yaml 文件  
   - 相关配置项可参考项目根目录下的配置文件模板 imail.yaml.example  
   - 后文也有对于配置文件的说明
4. 启动项目
```bash
go build -o imail
./imail server
```

## 配置文件说明
```yaml
global:
  env: "debug"  # debug | release 运行模式

catcher:
  # 各类爬虫的配置
  bilibiliVideo:  #[bilibili视频更新爬取]的相关配置
    source: "bilibili" #来源名称
    target:
      - uid: "uid001"  #up主的uid
        category: "开发"  #信息分类
      - uid: "uid002"
        category: "经济学"

manager:
  crontab: "0 19 * * *"  #cron表达式, 控制爬虫爬取和消息推送的频率
  immediate: true  #是否在项目启动时立马执行一次爬虫和消息推送

messager:
  email:
    sender:
      nickname: imailHelper  #发件人昵称
      host: smtp.qq.com  #smtp服务器地址
      port: 587  #smtp服务器端口
      username: xxxx@xxx.com  #smtp登录账号
      password: password  #smtp登录密码
    receiver:
      users:  #收件人邮箱地址
        - xxxx@xxx.com
```