```bash
go build -o go_program.so -buildmode=c-shared go_program.go
g++ -o main main.cpp -ldl
./main
```

```bash
go build -o go_program.so -buildmode=c-shared examples/simple/go_program.go
g++ -o main examples/simple/main.cpp -ldl
./main
```

```bash
go build -o go_program.so -buildmode=c-shared examples/go_struct_with_counter/go_program.go
g++ -o main examples/go_struct_with_counter/main.cpp -ldl
./main
```

```bash
go build -o go_program.so -buildmode=c-shared examples/go_run_goroutine/go_program.go
g++ -o main examples/go_run_goroutine/main.cpp -ldl
./main
```

```bash
go build -o go_program.so -buildmode=c-shared examples/go_struct_in_goroutine/go_program.go
g++ -o main examples/go_struct_in_goroutine/main.cpp -ldl
./main
```
