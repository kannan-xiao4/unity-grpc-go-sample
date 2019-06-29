package controller

import (
	"context"
	"fmt"
	"io"
	"log"
	"sample/pb"
)

type ChatController struct {
	savedGreet [] *sample_caht.Replay
}

func (*ChatController) Hello(ctx context.Context, name *sample_caht.Name) (*sample_caht.Replay, error) {
	var str = fmt.Sprintf("Hello, %s", name.Name)
	var rep = sample_caht.Replay{Message: str}
	return &rep, nil
}

func (controller *ChatController) ListGreet(empty *sample_caht.Empty, stream sample_caht.RouteGuide_ListGreetServer) error {
	for index, word := range controller.savedGreet {
		var replay = sample_caht.Replay{Message:fmt.Sprintf("index: %d, word: %s", index, word)}
		if err := stream.Send(&replay); err != nil {
			return err
		}
	}
	return nil
}

func (controller *ChatController) RecordGreet(stream sample_caht.RouteGuide_RecordGreetServer) error {
	for {
		var name, err = stream.Recv()
		if err == io.EOF {
			var summery = sample_caht.Summary{}
			return stream.SendAndClose(&summery)
		}
		if err != nil {
			return err
		}

		log.Printf("hogehoge:%s", name.Name)
	}
}

func (controller *ChatController) Chatting(sample_caht.RouteGuide_ChattingServer) error {
	panic("implement me")
}


