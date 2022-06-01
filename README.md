# simple-go-http-server

Pull Image:

```bash
docker pull yks0000/echoserver
```

Start Server:

```bash
docker run -p 8080:8080 yks0000/echoserver
```

Test:

```bash
$ curl http://127.0.0.1:8080/ -iv
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.79.1
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
HTTP/1.1 200 OK
< Date: Wed, 01 Jun 2022 06:56:18 GMT
Date: Wed, 01 Jun 2022 06:56:18 GMT
< Content-Length: 163
Content-Length: 163
< Content-Type: text/plain; charset=utf-8
Content-Type: text/plain; charset=utf-8

< 
GET / HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: curl/7.79.1
Accept: */*
Hostname: f2e14242bf95
Time: 2022-06-01 06:56:18.5405001 +0000 UTC m=+70.370248601
```
