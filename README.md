# goplay
PP Playing with Go

* Define person.proto
* Run protoc and compile to go

* [protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf)
* [protobuf/examples](https://github.com/protocolbuffers/protobuf/tree/master/examples)

~~~
cd C:\Users\ppprakas.ORADEV\go\src\hello>
~~~
~~~
go get -u github.com/golang/protobuf/proto
~~~
~~~
.\protoc.exe -I=C:\Users\ppprakas.ORADEV\go\src\hello\ --go_out=C:\Users\ppprakas.ORADEV\go\src\hello\ C:\Users\ppprakas.ORADEV\go\src\hello\person.proto
~~~
~~~
go run hello.go person.pb.go
~~~
