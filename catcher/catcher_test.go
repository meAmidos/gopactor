package catcher_test

import (
	"testing"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/meamidos/gopactor/catcher"
	"github.com/meamidos/gopactor/options"
	"github.com/stretchr/testify/assert"
)

type Child struct{}

func TestCatcher_ContextSpawn(t *testing.T) {
	a := assert.New(t)

	// Define a child that can respond to ping
	childProps := actor.FromFunc(func(ctx actor.Context) {
		switch m := ctx.Message().(type) {
		case string:
			if m == "ping" && ctx.Sender() != nil {
				ctx.Respond("pong")
			}
		}
	})

	// Define a parent that can spawn a child when asked to.
	// If a child is spawned, it's PID is sent back to the requestor.
	parentProps := actor.FromFunc(func(ctx actor.Context) {
		switch m := ctx.Message().(type) {
		case string:
			if m == "spawn" && ctx.Sender() != nil {
				child := ctx.Spawn(childProps)
				ctx.Respond(child)
			}
		}
	})

	catch := catcher.New()

	// Case 1: A parent without the dummy spawning
	parent, err := catch.Spawn(parentProps, options.OptNoInterception)
	a.Nil(err)

	// Spawn a child and get its PID
	res, err := parent.RequestFuture("spawn", options.DEFAULT_TIMEOUT).Result()
	a.Nil(err)
	child, ok := res.(*actor.PID)
	a.True(ok)

	// Send a ping to the child and receive a pong in return.
	res, err = child.RequestFuture("ping", options.DEFAULT_TIMEOUT).Result()
	a.Nil(err)
	resp, ok := res.(string)
	a.True(ok)
	a.Equal("pong", resp)

	// Case 2: A parent with the dummy spawning enabled
	parent, err = catch.Spawn(parentProps, options.OptNoInterception.WithDummySpawning())
	a.Nil(err)

	// Spawn a child and get its PID
	res, err = parent.RequestFuture("spawn", options.DEFAULT_TIMEOUT).Result()
	a.Nil(err)
	child, ok = res.(*actor.PID)
	a.True(ok)

	// Send a ping to the child and receive nothing because the child is a dummy actor
	res, err = child.RequestFuture("ping", options.DEFAULT_TIMEOUT).Result()
	a.Error(err)
	a.Contains(err.Error(), "timeout")
}
