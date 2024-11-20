package grpc

import "sync"

func StartServer(group *sync.WaitGroup) {
	defer group.Done()

}
