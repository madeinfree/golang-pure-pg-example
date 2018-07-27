# Golang api server

練習 Golang API Server，讓 Backend Server 能脫離 Nodejs，
搭配 docker-compose 做練習，利用 docker volume 來儲存永久資料，
為了保持內網一致，必須先設定 network subnet，才能 container 保持在同一個動態網段。

```
[ docker 內網設定 ]
docker network create -d bridge --subnet 192.168.0.0/24 --gateway 192.168.0.1 dockernet

[ pgsql ]
自動建立 id 遞增 SERIAL

[ Golang package ]
https://godoc.org/github.com/lib/pq#pkg-examples

[ Golang 文章 ]
<JSON>
https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
https://stackoverflow.com/questions/15672556/handling-json-post-request-in-go

[ Golang warning 解法 ]
// Person is a representation of a person
type Person struct {
    Name string
    Phone string
}

[ Golang 回應表頭設定 ]
write.Header().Set
```
