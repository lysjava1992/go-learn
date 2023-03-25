package main

import (
	"context"
	"flag"
	"fmt"
	"gin-web/grpc/proto/grpc/service"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	service.UnimplementedHelloServer
}

func (s *server) SayHello(ctx context.Context, in *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{Message: "Hello !" + in.GetName() + " This is go !!!"}, nil
}
func (s *server) SayHello2(client service.Hello_SayHello2Server) error {
	fmt.Println("SayHello2被调用")
	count := 0
	for {
		a, err := client.Recv()

		if err == io.EOF {
			return client.SendAndClose(&service.HelloResponse{
				Message: fmt.Sprintf("累计接收消息:%d条", count)})

		}
		if err != nil {

			return err
		}
		fmt.Println("SayHello2被调用: ", a)
		count++
	}
	return nil
}
func (s *server) SayHello3(in *service.HelloRequest, res service.Hello_SayHello3Server) error {
	fmt.Println("SayHello3 被调用")
	for i := 0; i < 10; i++ {
		err := res.Send(&service.HelloResponse{Message: fmt.Sprintf("%s+%d\n", in.Name, i)})
		if err != nil {
			fmt.Println("SayHello3 返回异常")
		}
	}
	return nil
}
func (s *server) SayHello4(in service.Hello_SayHello4Server) error {
	fmt.Println("SayHello4 被调用")
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		for {
			a, err := in.Recv()
			if err != nil {
				fmt.Println("SayHello4 调用异常：", err)
			}
			fmt.Println("SayHello4收到客户端消息: ", a)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			err := in.Send(&service.HelloResponse{Message: fmt.Sprintf("我是服务器 %d\n", i)})
			if err != nil {
				fmt.Println("SayHello4 返回异常")
			}
		}
	}()
	return nil
}

//	protoc --go_out=.  --go_opt=paths=source_relative \
//	   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
//	   helloworld/helloworld.proto

// go_out=. 指定生成的pb.go文件所在目录(如果没有该目录，需要手动提前创建)，
//
//			 . 代表当前protoc执行目录，结合.proto文件中的option go_package，
//			 其最终的生成文件目录为go_out  指定目录/go_package指定目录
//	 --go_opt=paths=source_relative 生成的文件目录不依赖于option
//		  go-grpc_out针对_grpc.pb.go文件，作用同上
func main() {
	flag.Parse()
	s := grpc.NewServer()
	//注册
	service.RegisterHelloServer(s, &server{})
	// 监听端口

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
