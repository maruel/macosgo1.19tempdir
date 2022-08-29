# macOS issue

```
go install golang.org/dl/go1.18.5@latest
go1.18.5 download
go install golang.org/dl/go1.19@latest
go1.19 download

go1.18.5 run main.go
go1.19 run main.go
```

The current working directory used to be resolved for symlinks and doesn't anymore.
The path printed out used to be within /private/var/..., it's now in /var/.

## Output with go1.18.5

```
$ go1.18.5 run main.go
Created: /var/folders/xh/c8h6f2jx5b15m1w7whmm1l_8001nj2/T/2642791120
Running from: /private/var/folders/xh/c8h6f2jx5b15m1w7whmm1l_8001nj2/T/2642791120
```

## Output with go1.19

```
$ go1.19 run main.go
Created: /var/folders/xh/c8h6f2jx5b15m1w7whmm1l_8001nj2/T/406562626
Running from: /var/folders/xh/c8h6f2jx5b15m1w7whmm1l_8001nj2/T/406562626
```
