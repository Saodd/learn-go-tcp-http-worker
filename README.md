# learn-go-tcp-http-worker

鸣谢 & 灵感来源： https://www.v2ex.com/t/727922#

一次HTTP并发编程比赛。

赛题主要内容是： 写一个服务A，接收客户端请求后向指定的服务B发起N次调用，并将N次调用的结果汇总返回。

并发压力：调用端10k并发。
