## qbtNotification
一个qBittorrent向qq通知下载开始与完成的简单脚本（基于OneBot协议的[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)qq机器人)   
想法来自[视频链接](https://www.bilibili.com/video/BV1qP411m7zX/)  
(虽然确实已经有其他通知方式——windows和邮箱，但本就图一乐~)

### Usage

```
go run main.go -help
```
```
-check  bool
    是否检查文件名为hash值 (default false)
-group int
    QQ群号
-host string
    go-cqhttp的http地址 (default "127.0.0.1:5700")
-id int
    QQ号
-n string
    下载完成的内容
-s string
    状态 (default "end")
-t string
    access_token
```
#### 根据以上参数提示flag指定相应的参数
如下图
![img.png](https://github.com/Yoak3n/qbtNotification/blob/main/docs/usage.png)

    1. 最好使用程序所在的绝对路径    
    2. 如果为了go-cqhttp部署在公网上的安全而设置了access_token，使用-t指定access_token的值
    3. 如果参数的内容中存在一些符号（如-host、-n)，就用英文双引号""包裹
    4. 由于查询文件名使用了第三方接口，默认不开启（一般文件名为hash值的情况都出现在手动下载时，需求不大，也尽量少用）




