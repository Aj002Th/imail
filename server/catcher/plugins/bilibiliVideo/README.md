## bilibiliVideo

获取 up 主的视频更新情况

### 使用方法

在配置文件的 catcher 部分添加相应配置

```yaml
bilibiliVideo:
  source: "bilibili" #来源名称
  target:
    - uid: "uid001"  #up主的uid
      category: "开发"  #信息分类
    - uid: "uid002"
      category: "经济学"
```

### 注意事项

只爬取了视频列表第一页(包含最多30个视频), 如果一个爬取间隔内更新超过30个视频, 则会有更新信息被遗漏

通常情况下也不会有这么勤奋的up主吧 :)
