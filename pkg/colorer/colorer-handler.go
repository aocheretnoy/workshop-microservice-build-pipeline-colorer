package colorer

import (
	"log"
	"time"

	"golang.org/x/net/context"
)

type colorerServer struct {
}

// GetEcho returns the feature at the given point.
func (s *colorerServer) GetColor(ctx context.Context, msg *GetColorRequest) (*GetColorResponse, error) {
	log.Printf("Server colorer called with message (%v)", msg)
	time.Sleep(time.Duration(500) * time.Millisecond)
	return &GetColorResponse{Cold: 0, Hot: 100}, nil
}

func NewServer() ColorerServer {
	s := &colorerServer{}
	return s
}
