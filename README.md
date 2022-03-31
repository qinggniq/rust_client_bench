# Bench of rust client
## env
```shell
$ rustc --version
rustc 1.61.0-nightly (1eb72580d 2022-03-08)
$ go version
go version go1.15.15 darwin/amd64
```
## usage
### Bench HTTP2
#### run server
```shell
$ cd ./cmd/server
$ HTTP2=true go run main.go
HTTP2 Listening [0.0.0.0:1010]...
```
#### run go client
```shell
$ cd ./cmd/client
$ HTTP2=true go run main.go
HTTP2 bench start
avgQps 34680
avgQps 37906
avgQps 34680
avgQps 36216
avgQps 35821
...
```
#### run rust client
```shell
$ cd ./cmd/rust_client
$ cargo build --release
$ HTTP2=true ./target/release/rust_client
avg qps 11374
avg qps 7021
avg qps 4428
...
```
Calculating qps every 5s.

### Bench HTTP1
#### run server
```shell
$ cd ./cmd/server
$ go run main.go
HTTP1 Listening [0.0.0.0:1010]...
```
#### run go client
```shell
$ cd ./cmd/client
$ go run main.go
HTTP1 bench start
avgQps 83671
avgQps 81433
avgQps 90595
...
```
#### run rust client
```shell
$ cd ./cmd/rust_client
$ cargo build --release
$ /target/release/rust_client
avg qps 66956
avg qps 66924
avg qps 64205
avg qps 66796
avg qps 61512
avg qps 67504
avg qps 61505
...
```
Calculating qps every 5s.
