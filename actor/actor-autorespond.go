package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"time"
)

// 返回值Actor
// 请求
type myAutoResponder struct {
	name string
}

// GetAutoResponse 请求回调函数
func (m *myAutoResponder) GetAutoResponse(context actor.Context) interface{} {
	return myAutoResponse{
		name: m.name + " " + context.Self().Id,
	}

}

// 返回
type myAutoResponse struct {
	name string
}

func main() {
	system := actor.NewActorSystem()

	//构建有返回值的props 定义目标actor的行为动作
	props := actor.PropsFromFunc(func(ctx actor.Context) {
		switch msg := ctx.Message().(type) {
		//match message based on type
		case *myAutoResponder:
			msg.name = msg.name + " ---"
		}
	})

	// 创建actor 由root创建 返回pid地址
	pid := system.Root.Spawn(props)
	res, _ := system.Root.RequestFuture(pid, &myAutoResponder{name: "Tom"}, 10*time.Second).Result()

	fmt.Println(res)
}
