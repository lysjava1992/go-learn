package main

import (
	"fmt"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"reflect"
)

// HelloMessage  消息
type HelloMessage struct {
	Who string
}

type HelloActor struct {
}

func (state *HelloActor) Receive(context actor.Context) {
	fmt.Println("type:", reflect.TypeOf(context.Message()))
	switch msg := context.Message().(type) {
	//match message based on type
	case *HelloMessage:
		fmt.Printf("Hello %v\n", msg.Who)
	}

}
func main() {

	// 创建actor
	system := actor.NewActorSystem()
	//构建props 定义目标actor的行为动作
	props := actor.PropsFromProducer(func() actor.Actor { return &HelloActor{} })
	// 由 rootActor创建actor
	pid := system.Root.Spawn(props)

	//组装消息
	h := HelloMessage{Who: "Tom"}
	hMessage := &h

	//发送消息
	system.Root.Send(pid, hMessage)

	_, _ = console.ReadLine()
}
