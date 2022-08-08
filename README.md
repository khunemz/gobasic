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
go test gobasic/./...
```

Test sub test
```go
go test gobasic/services -run="TestCheckGrade/A" -v
```
Coverage
```go
go test gobasic/services -cover
```

Virutalizes Coverage 