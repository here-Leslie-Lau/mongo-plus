package test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
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
	conn, f := newConnWithMonitor(monitor)
	defer f()
	res := new(result)
	err := conn.Collection(&demo{"demo"}).Where("name", "leslie").FindOne(res)
	require.Nil(t, err)

	require.Equal(t, "leslie", res.Name)
}

func TestCommandMonitor(t *testing.T) {
	m := sync.Map{}
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			m.Store(evt.RequestID, evt.Command)
		},
		Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
			var commands bson.Raw
			v, ok := m.Load(evt.RequestID)
			if ok {
				commands = v.(bson.Raw)
			}
			fmt.Printf("\n [%.3fms] %s, %v\n", float64(evt.Duration)/1e6, commands.String(), evt.Reply)
		},
		Failed: func(_ context.Context, evt *event.CommandFailedEvent) {
			var commands bson.Raw
			v, ok := m.Load(evt.RequestID)
			if ok {
				commands = v.(bson.Raw)
			}
			fmt.Printf("\n [%.3fms] %s, %v\n", float64(evt.Duration)/1e6, commands.String(), evt.Failure)
		},
	}
	conn, f := newConnWithMonitor(monitor)
	defer f()

	res := new(result)
	err := conn.Collection(&demo{"demo"}).Where("name", "leslie").FindOne(res)
	require.Nil(t, err)

	require.Equal(t, "leslie", res.Name)
}

func BenchmarkCommandMonitor(b *testing.B) {
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			fmt.Println("Started")
		},
		Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
			fmt.Println("Succeeded")
		},
		Failed: func(_ context.Context, evt *event.CommandFailedEvent) {
			fmt.Println("Failed")
		},
	}
	conn, f := newConnWithMonitor(monitor)
	defer f()

	for i := 0; i < b.N; i++ {
		res := new(result)
		err := conn.Collection(&demo{"demo"}).Where("name", "leslie").FindOne(res)
		require.Nil(b, err)

		require.Equal(b, "leslie", res.Name)
	}
}

func TestServerMonitor(t *testing.T) {
	monitor := &event.ServerMonitor{
		TopologyOpening: func(_ *event.TopologyOpeningEvent) {
			fmt.Println("TopologyOpening")
		},
		ServerClosed: func(_ *event.ServerClosedEvent) {
			fmt.Println("ServerClosed")
		},
	}
	conn, f := newConnWithMonitor(monitor)
	defer f()

	res := new(result)
	err := conn.Collection(&demo{"demo"}).Where("name", "leslie").FindOne(res)
	require.Nil(t, err)

	require.Equal(t, "leslie", res.Name)
}
