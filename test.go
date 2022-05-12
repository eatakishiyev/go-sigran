package main

import (
	"encoding/hex"
	"fmt"
	"go-sigtran/m3ua"
	"go-sigtran/sccp"
	"go-sigtran/tcap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Queue struct {
	q []int
}

func (q *Queue) Dequeue() int {
	var e = q.q[0]
	q.q = q.q[1:]
	return e
}

func (q *Queue) Enqueue(e int) {
	q.q = append(q.q, e)
}

func main() {

	queue := Queue{
		q: []int{1, 2, 3},
	}
	dequeue := queue.Dequeue()
	fmt.Printf("dequeued %d\n", dequeue)

	dequeue = queue.Dequeue()
	fmt.Printf("dequeued %d\n", dequeue)

	var tcapMessageData = "671a49049b1391926b122810060700118605010101a0056403800100"

	data, err := hex.DecodeString(tcapMessageData)
	if err != nil {
		fmt.Printf("error occurred while decode hex-string %s", err)
	}

	tcapStack := tcap.NewTcapStack("TestTcapStack", 1, 10, time.Duration(10)*time.Second)

	tcapStack.OnMessage(nil, nil, data, true, 1, sccp.ReturnMessageOnError)

	asp := &m3ua.Asp{
		BindPort: 2905,
		BindIp:   "127.0.0.1",
		IsServer: true,
	}

	asp.Start()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	stack := m3ua.GetM3UAStackInstance()
	fmt.Printf("Stack = %s\n", stack)
	stack = m3ua.GetM3UAStackInstance()
	fmt.Printf("Stack = %s\n", stack)
}
