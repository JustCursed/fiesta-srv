package services

import (
	"context"
	"github.com/ClickHouse/ch-go"
	"google.golang.org/protobuf/types/known/emptypb"
	"void-studio.net/fiesta/database"
	"void-studio.net/fiesta/pb"
)

type Collector struct {
	pb.UnimplementedCollectorServer
}

func (s *Collector) SaveChatLog(_ context.Context, data *pb.ChatData) (*emptypb.Empty, error) {
	input := database.ChatValues(data)

	database.ExecuteQuery(ch.Query{
		Body:  input.Into("chat"),
		Input: input,
	})

	return &emptypb.Empty{}, nil
}

func (s *Collector) SaveItemLog(_ context.Context, data *pb.ItemData) (*emptypb.Empty, error) {
	input := database.ItemValues(data)

	database.ExecuteQuery(ch.Query{
		Body:  input.Into("items"),
		Input: input,
	})

	return &emptypb.Empty{}, nil
}

func (s *Collector) SaveMovementLog(_ context.Context, data *pb.MovementData) (*emptypb.Empty, error) {
	input := database.MovementValues(data)

	database.ExecuteQuery(ch.Query{
		Body:  input.Into("movement"),
		Input: input,
	})

	return &emptypb.Empty{}, nil
}

func (s *Collector) SaveLoggedLog(_ context.Context, data *pb.LoggedData) (*emptypb.Empty, error) {
	input := database.LoggedValues(data)

	database.ExecuteQuery(ch.Query{
		Body:  input.Into("logged"),
		Input: input,
	})

	return &emptypb.Empty{}, nil
}
