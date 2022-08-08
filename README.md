# Testing


```go
go test gobasic/services -v 
```

Result 
```go
=== RUN   TestCheckGrade
--- PASS: TestCheckGrade (0.00s)
=== RUN   TestHello
--- PASS: TestHello (0.00s)
PASS
ok      gobasic/services        0.420s
```

Test Everything
```go
go test ./...
```

Test sub test
```go
go test gobasic/services -run="TestCheckGrade/A" -v
```
Coverage
```go
go test gobasic/services -cover
```

Benchnark
```go
go test gobasic/services -bench=.

// benchmark memory
go test gobasic/services -bench=. -benchmem
```

Go doc
```go
# terminal
go get golang.org/x/tools/cmd/godoc
godoc -http=8000
```

mysql
```
CREATE USER 'user1'@'localhost' IDENTIFIED WITH mysql_native_password BY 'P#ssw0rd'
# do not use this in production
GRANT ALL PRIVILEGES ON *.* TO 'user1'@'localhost' WITH GRANT OPTION;
```