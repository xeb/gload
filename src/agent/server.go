package agent

import (
	"fmt"
	stime "github.com/xeb/gload/src/time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

type server struct{}

func (s *server) GetTime(ctx context.Context, in *stime.TimeRequest) (*stime.TimeReply, error) {
	return &stime.TimeReply{TimeValue: time.Now().UnixNano()}, nil
}

func Bind(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		etxt := err.Error()
		if strings.Contains(etxt, "address already in use") || strings.Contains(etxt, "permission denied") {
			fmt.Printf("Attempting to rebind on %d\n", port+1)
			Bind(port + 1)
		} else {
			log.Fatalf("failed to listen: %v", err)
		}
	}
	s := grpc.NewServer()
	stime.RegisterTimerServer(s, &server{})
	fmt.Printf("FOUND PORT - Listening on %d\n", port)
	s.Serve(lis)

	// Register the Agent with Discovery

	// Bind Other commands

}
