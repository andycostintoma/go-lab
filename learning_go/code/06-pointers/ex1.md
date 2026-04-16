Compiling the code with `-gcflags="-m"` shows when values escape to the heap (it also shows when functions are inlined, which is a compiler optimization that's not covered by this book):

````shell
$ go build -gcflags="-m" ex1.go
./main.go:11:6: can inline MakePerson
./main.go:19:6: can inline MakePersonPointer
./main.go:28:17: inlining call to MakePerson
./main.go:29:13: inlining call to fmt.Println
./main.go:30:25: inlining call to MakePersonPointer
./main.go:31:13: inlining call to fmt.Println
./main.go:11:17: leaking param: firstName to result ~r0 level=0
./main.go:11:28: leaking param: lastName to result ~r0 level=0
./main.go:19:24: leaking param: firstName
./main.go:19:35: leaking param: lastName
./main.go:20:9: &Person{...} escapes to heap
./main.go:29:13: ... argument does not escape
./main.go:29:14: p escapes to heap
./main.go:30:25: &Person{...} escapes to heap
./main.go:31:13: ... argument does not escape
````

The `&Person{}` returned from `MakePersonPointer` escapes to the heap. Any time a pointer is returned from a function, the pointer is returned on the stack, but the value it points to must be stored on the heap.

Surprisingly, it also says that `p` escapes to the heap. The reason for this is that it is passed into `fmt.Println`. This is because the parameter to `fmt.Println` are `...any`. The current Go compiler moves to the heap any value that is passed in to a function via a parameter that is of an interface type.