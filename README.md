# learn-go-sock5-proxy
参考 https://www.v2ex.com/t/727922#

```shell-session
% curl localhost:1080
curl: (52) Empty reply from server

```

```shell-session
2020/11/23 18:12:50 /learn-go-sock5-proxy/main.go:33: [::1]:53198
GET / HTTP/1.1
Host: localhost:1080
User-Agent: curl/7.64.1
Accept: */*


2020/11/23 18:12:50 /learn-go-sock5-proxy/main.go:31: [::1]:53198 closed!
```
