package database

import (
	"fmt"
	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
	"strings"
	"time"
	"void-studio.net/fiesta/pb"
)

func ChatValues(data *pb.ChatData) proto.Input {
	var (
		player   proto.ColStr
		server   proto.ColStr
		message  proto.ColStr
		location proto.ColBool
		command  proto.ColBool
		private  proto.ColBool
		cords    proto.ColStr
		clock    proto.ColDateTime
	)

	player.Append(data.Player)
	server.Append(data.Server)
	message.Append(data.Message)
	location.Append(strings.HasPrefix(data.Message, "!"))
	command.Append(strings.HasPrefix(data.Message, "/"))
	private.Append(data.Private)
	cords.Append(data.Cords)
	clock.Append(time.Unix(int64(data.Time), 0))

	return proto.Input{
		{Name: "player", Data: player},
		{Name: "message", Data: message},
		{Name: "server", Data: server},
		{Name: "location", Data: location},
		{Name: "command", Data: command},
		{Name: "private", Data: private},
		{Name: "cords", Data: cords},
		{Name: "time", Data: clock},
	}
}

func ItemValues(data *pb.ItemData) proto.Input {
	var (
		player proto.ColStr
		item   proto.ColStr
		amount proto.ColUInt8
		action proto.ColBool
		server proto.ColStr
		cords  proto.ColStr
		clock  proto.ColDateTime
	)

	player.Append(data.Player)
	server.Append(data.Server)
	item.Append(data.Item)
	amount.Append(uint8(data.Amount))
	action.Append(data.Action)
	cords.Append(data.Cords)
	clock.Append(time.Unix(int64(data.Time), 0))

	return proto.Input{
		{Name: "player", Data: player},
		{Name: "item", Data: item},
		{Name: "amount", Data: amount},
		{Name: "server", Data: server},
		{Name: "cords", Data: cords},
		{Name: "time", Data: clock},
	}
}

func MovementValues(data *pb.MovementData) proto.Input {
	var (
		player proto.ColStr
		from   proto.ColStr
		to     proto.ColStr
		clock  proto.ColDateTime
	)

	player.Append(data.Player)
	from.Append(data.From)
	to.Append(data.To)
	clock.Append(time.Unix(int64(data.Time), 0))

	return proto.Input{
		{Name: "player", Data: player},
		{Name: "from", Data: from},
		{Name: "to", Data: to},
		{Name: "time", Data: clock},
	}
}

func LoggedValues(data *pb.LoggedData) proto.Input {
	var (
		player proto.ColStr
		server proto.ColStr
		action proto.ColBool
		cords  proto.ColStr
		clock  proto.ColDateTime
	)

	player.Append(data.Player)
	server.Append(data.Server)
	action.Append(data.Action)
	cords.Append(data.Cords)
	clock.Append(time.Unix(int64(data.Time), 0))

	return proto.Input{
		{Name: "player", Data: player},
		{Name: "server", Data: server},
		{Name: "action", Data: action},
		{Name: "cords", Data: cords},
		{Name: "time", Data: clock},
	}
}

func ExecuteQuery(query ch.Query) bool {
	err := pool.Do(ctx, query)

	if err != nil {
		_ = fmt.Errorf("failed to execute query: %v", err)
		return false
	}

	return true
}
