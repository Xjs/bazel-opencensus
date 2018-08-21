package debugopencensus

import "go.opencensus.io/plugin/ocgrpc"

func init() {
	var _ ocgrpc.ClientHandler
}
