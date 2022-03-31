# 執行指令

```
go test -v -bench=. -run=none -benchmem .
```

## pprof

```
go test -bench=. -benchmem -cpuprofile profile.out

go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
```

### 使用pprof

```
go tool pprof profile.out

go tool pprof memprofile.out
```

#### 檢查函數需要時間

```
list GenerateBwSnowflake

list GenerateFtSnowflake
```