package signalr

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSignalR(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SignalR Suite")
}

func connect(hubProto HubInterface) *testingConnection {
	server, err := NewServer(SimpleHubFactory(hubProto),
		Logger(log.NewLogfmtLogger(os.Stderr), false),
		HubChanReceiveTimeout(200*time.Millisecond),
		StreamBufferCapacity(5))
	if err != nil {
		Fail(err.Error())
		return nil
	}
	conn := newTestingConnection()
	go server.Run(context.TODO(), conn)
	return conn
}
