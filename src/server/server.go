package main

import (
    "log"
    "net"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "calc"
    "google.golang.org/grpc/reflection"
    "fmt"
    "errors"
)

const (
    port = ":50051"
)

// server is used to implement calc.
type server struct{}

func (s *server) Add(ctx context.Context, in *calc.AddRequest) (*calc.AddReply, error){
	fmt.Println("%v + %v = %v",in.Num1,in.Num2,in.Num1+in.Num2)
	return &calc.AddReply{Result:(in.Num1+in.Num2)}, nil
}

func (s *server) Sub(ctx context.Context, in *calc.SubRequest) (*calc.SubReply, error){
	fmt.Println("%v - %v = %v", in.Num1, in.Num2, in.Num1 - in.Num2)
	return &calc.SubReply{Result:(in.Num1 - in.Num2)}, nil
}

func (s *server) Mult(ctx context.Context, in *calc.MultRequest) (*calc.MultReply, error){
        fmt.Println("%v * %v = %v",in.Num1,in.Num2,in.Num1*in.Num2)
        return &calc.MultReply{Result:(in.Num1*in.Num2)}, nil
}

func (s *server) Div(ctx context.Context, in *calc.DivRequest) (*calc.DivReply, error){
    	if 0 == in.Num2{
		return nil,errors.New("divisor is 0") 
	} 
        fmt.Println("%v / %v = %v",in.Num1,in.Num2,in.Num1/in.Num2)
        return &calc.DivReply{Result:(in.Num1/in.Num2)}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    calc.RegisterCalcServer(s, &server{})
    // Register reflection service on gRPC server.
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
