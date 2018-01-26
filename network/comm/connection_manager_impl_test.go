package comm

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	pb "it-chain/network/protos"
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
)

func ListenMockServerWithIP(ip string) (*grpc.Server,net.Listener){
	lis, err := net.Listen("tcp", ip)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &Mockserver{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func(){
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			s.Stop()
			lis.Close()
		}
	}()

	return s,lis
}

func TestCommImpl_CreateStreamConn(t *testing.T) {

	server1, listner1 := ListenMockServerWithIP("127.0.0.1:5555")
	server2, listner2 := ListenMockServerWithIP("127.0.0.1:6666")


	comm := NewConnectionManagerImpl()
	comm.CreateStreamConn("1","127.0.0.1:5555",nil)
	comm.CreateStreamConn("2","127.0.0.1:6666",nil)

	defer func(){
		server1.Stop()
		listner1.Close()
		server2.Stop()
		listner2.Close()
	}()

	assert.NotNil(t,comm)
	assert.Equal(t,2,comm.Size())
}

func TestCommImpl_Send(t *testing.T) {
	counter = 0

	server1, listner1 := ListenMockServerWithIP("127.0.0.1:5555")
	server2, listner2 := ListenMockServerWithIP("127.0.0.1:6666")


	comm := NewConnectionManagerImpl()
	comm.CreateStreamConn("1","127.0.0.1:5555",nil)
	comm.CreateStreamConn("2","127.0.0.1:6666",nil)

	envelope := &pb.Envelope{Signature:[]byte("123")}

	comm.SendStream(*envelope,nil, "2")
	comm.SendStream(*envelope, nil, "2")

	defer func(){
		server1.Stop()
		listner1.Close()
		server2.Stop()
		listner2.Close()
		comm.Stop()
	}()

	time.Sleep(3*time.Second)

	assert.Equal(t,2,counter)
}

func TestCommImpl_Stop(t *testing.T) {
	counter = 0
	server1, listner1 := ListenMockServerWithIP("127.0.0.1:5555")
	server2, listner2 := ListenMockServerWithIP("127.0.0.1:6666")


	comm := NewConnectionManagerImpl()
	comm.CreateStreamConn("1","127.0.0.1:5555",nil)
	comm.CreateStreamConn("2","127.0.0.1:6666",nil)

	defer func(){
		server1.Stop()
		listner1.Close()
		server2.Stop()
		listner2.Close()
	}()

	comm.Stop()

	assert.Equal(t,0,comm.Size())
}

func TestCommImpl_Close(t *testing.T) {

	server1, listner1 := ListenMockServerWithIP("127.0.0.1:5555")
	server2, listner2 := ListenMockServerWithIP("127.0.0.1:6666")


	comm := NewConnectionManagerImpl()
	comm.CreateStreamConn("1","127.0.0.1:5555",nil)
	comm.CreateStreamConn("2","127.0.0.1:6666",nil)

	defer func(){
		server1.Stop()
		listner1.Close()
		server2.Stop()
		listner2.Close()
	}()

	comm.Close("1")

	assert.Equal(t,1,comm.Size())

	comm.Stop()
}
