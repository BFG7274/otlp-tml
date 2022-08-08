package logs

import (
	context "context"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LogClient struct {
	logClient LogsDataClient
}

func (lc *LogClient) WriteLog(msg, level string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lc.logClient.WriteLog(ctx, &LogMsg{
		Timestamp: time.Now().Unix(),
		Msg:       msg,
		Level:     level,
	})
	return nil
}

func NewLogClient() *LogClient {
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := NewLogsDataClient(conn)
	return &LogClient{
		logClient: c,
	}
}
