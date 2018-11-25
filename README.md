# goplay
PP Playing with Go

~~~
cd C:\Users\ppprakas.ORADEV\go\src\hello>
go get -u github.com/golang/protobuf/proto
~~~
~~~
.\protoc.exe -I=C:\Users\ppprakas.ORADEV\go\src\hello\ --go_out=C:\Users\ppprakas.ORADEV\go\src\hello\ C:\Users\ppprakas.ORADEV\go\src\hello\person.proto
~~~
~~~
go run hello.go person.pb.go
~~~
