# macOS issue

```
go install golang.org/dl/go1.18.5@latest
go1.18.5 download
go install golang.org/dl/go1.19@latest
go1.19 download

go1.18.5 run main.go
go1.19 run main.go
```

The path printed out by a stack trace used to be within /private/var/..., it's
now in /var/.

## Output with go1.18.5

```
$ go1.18.5 run main.go
Running within /var/folders/xh/c8h6f2jx5b15m1w7whmm1l_8001nj2/T/3121730881
panic: oops

goroutine 1 [running]:
main.main()
        /private/var/folders/xh/c8h6f2jx5b15m1w7whmm1l_8001nj2/T/3121730881/inner/main.go:4 +0x30
```

## Output with go1.19

```
$ go1.19 run main.go
Running within /var/folders/xh/c8h6f2jx5b15m1w7whmm1l_8001nj2/T/2440236045
panic: oops

goroutine 1 [running]:
main.main()
        /var/folders/xh/c8h6f2jx5b15m1w7whmm1l_8001nj2/T/2440236045/inner/main.go:4 +0x2c
```
