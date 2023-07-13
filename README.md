# innocuity_phishing_test 

## 概述

  当你的甲方爸爸需要做钓鱼测试，而你又不想污染cs池的时候，就需要上一个无毒的钓鱼软件来给甲方爸爸做钓鱼测试

---

## 编译方法

有go环境后

~~~
运行build.bat
~~~

---

## 使用方法

1. 在你的服务器上运行程序（linux运行server，windows运行server.exe活server32.exe）
2. 修改client.go中的ip为运行server的IP，并再次编译客户端
3. 服务端的10004为web服务，可以查看到点击客户端的主机名和访问8.8.8.8得到的ip地址以及点击次数
4. 点击按钮可以保存记录并且清除数据重新开始记录