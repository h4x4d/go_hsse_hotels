package main

import (
	"github.com/h4x4d/go_hsse_hotels/hotel/cmd/grpc"
	"github.com/h4x4d/go_hsse_hotels/hotel/cmd/rest"
	"sync"
)

func main() {
	group := sync.WaitGroup{}
	group.Add(2)
	go rest.StartServer(&group)
	go grpc.StartServer(&group)

	group.Wait()
}
