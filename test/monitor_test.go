package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/event"
)

func TestPoolMonitor(t *testing.T) {
	monitor := &event.PoolMonitor{
		Event: func(evt *event.PoolEvent) {
			switch evt.Type {
			case event.GetSucceeded:
				fmt.Println("GetSucceeded")
			case event.ConnectionReturned:
				fmt.Println("ConnectionReturned")
			case event.PoolCreated:
				fmt.Println("PoolCreated")
			}
		},
	}
	conn, f := newConnWithPoolMonitor(monitor)
	defer f()

	res := new(result)
	err := conn.Collection(&demo{"demo"}).Where("name", "leslie").FindOne(res)
	require.Nil(t, err)

	fmt.Printf("res: %+v\n", res)
	require.Equal(t, "leslie", res.Name)
}
