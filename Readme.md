github项目地址:https://github.com/Ledgerbiggg/m3u8_download

## 什么是m3u8?

- M3U8是一种用于存储多媒体播放列表的文件格式。它通常用于指导流媒体播放器播放音频或视频流。M3U8文件本质上是文本文件，其中包含一系列URL，这些URL指向媒体文件或流的各个片段。这种文件格式通常用于网络直播、在线视频和音频流服务，因为它允许媒体内容以分段的方式传输，从而更有效地适应网络带宽和客户端设备的特性。M3U8文件也支持指定播放顺序、分辨率和编码等参数，以提供更好的用户体验。

**简单看一下**

- URL 秘钥文件的地址(一般是和m3u8文件放一起)
- IV（Initialization Vector，初始化向量）是一种随机生成的初始输入值。IV在对称加密和一些加密模式中起着重要作用
- 之后有各种的ts分片,这些是视频的一部分文件

![img](https://img2.imgtp.com/2024/03/17/vTtn3JuL.png)

## 项目介绍

### 简介

- 这个项目是一个m3u8的多线程下载器,只需配置一个m3u8的url地址,会自动多线程下载并解密,最后拼接成一个完成的视频
- 这里是几个已知的m3u8的地址

```yaml
mu8:
  #  https://cdn1.xlzys.com/play/9b6BWALb/index.m3u8 新神榜杨戬
  #  https://cdn.wlcdn99.com:777/69fcc1e2/index.m3u8 功夫熊猫
  #  https://cdn.wlcdn99.com:777/4b868107/index.m3u8 周处除三害
  #  https://m3u.haiwaikan.com/xm3u8/286a8da263780575721b36a490557330fe5fcbc155e95a999f864a2f50c98f919921f11e97d0da21.m3u8 热辣滚烫
  #  新神榜杨戬
  url: https://cdn1.xlzys.com/play/9b6BWALb/index.m3u8
  #下载ts文件的线程数
  thread: 30
```

### **就是比别人要快!!!!!**

### 使用方法:

1. 配置好url地址和下载的线程(推荐10-50)
2. 双击Down_m3u8.exe启动程序
3. 等待下载和拼接完成

![](https://img2.imgtp.com/2024/03/17/dcXJm84F.png)

1. 直到出现打印:完成下载请打开target.ts观看

![img](https://img2.imgtp.com/2024/03/17/wPc7Hhhw.png)

