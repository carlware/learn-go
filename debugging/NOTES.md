## DELVE

### install
```shell
go install github.com/go-delve/delve/cmd/dlv@latest
```

### useful commands
```text
dlv debug .
continue
restart
list main.go:14
break main.go:14
breakpoints
step
args
locals
next
stack
threads
exit
clearall
// same as continue
c
// show source code
list
```

### remote debug commands
```shell
docker build --tag godebug .
docker run --security-opt="seccomp=unconfined" --cap-add=SYS_PTRACE -p:5001:5000 -p:2345:2345 godebug 
```