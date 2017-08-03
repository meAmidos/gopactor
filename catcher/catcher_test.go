package catcher_test

import (
	"testing"

	"github.com/AsynkronIT/protoactor-go/actor"
	. "github.com/smartystreets/goconvey/convey"
	"gopactor/catcher"
	"gopactor/options"
)

type Child struct{}

func TestCatcher_ContextSpawnDummy(t *testing.T) {
	Convey("Subject: Spawn dummy actors", t, func() {
		Convey("Given a catcher and parent/child definitions", func() {
			catch := catcher.New()

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

			Convey("And a parent is created without the dummy spawning", func() {
				parent, err := catch.Spawn(parentProps, options.OptNoInterception)
				So(err, ShouldBeNil)

				Convey("When spawning a child", func() {
					res, err := parent.RequestFuture("spawn", options.DEFAULT_TIMEOUT).Result()
					So(err, ShouldBeNil)
					child, ok := res.(*actor.PID)
					So(ok, ShouldBeTrue)

					Convey("And sending a ping to it", func() {
						res, err = child.RequestFuture("ping", options.DEFAULT_TIMEOUT).Result()

						Convey("Then receive a pong in return", func() {
							So(err, ShouldBeNil)
							resp, ok := res.(string)
							So(ok, ShouldBeTrue)
							So(resp, ShouldEqual, "pong")
						})
					})
				})
			})

			Convey("And a parent is created with the dummy spawning enabled", func() {
				parent, err := catch.Spawn(parentProps, options.OptNoInterception.WithDummySpawning())
				So(err, ShouldBeNil)

				Convey("When spawning a child", func() {
					res, err := parent.RequestFuture("spawn", options.DEFAULT_TIMEOUT).Result()
					So(err, ShouldBeNil)
					child, ok := res.(*actor.PID)
					So(ok, ShouldBeTrue)

					Convey("And sending a ping to it", func() {
						res, err = child.RequestFuture("ping", options.DEFAULT_TIMEOUT).Result()

						// Becuase the child is a dummy actor
						Convey("Then receive nothing in return", func() {
							So(err, ShouldNotBeNil)
							So(err.Error(), ShouldContainSubstring, "timeout")
						})
					})
				})
			})
		})
	})
}

func TestCatcher_ContextSpawnInterception(t *testing.T) {
	Convey("Subject: Spawn interception in catcher", t, func() {
		Convey("Given a catcher and parent/child definitions", func() {
			catch := catcher.New()

			// Define a dummy child
			childProps := actor.FromFunc(func(ctx actor.Context) {})

			// Define a parent that can spawn a child when asked to.
			// If a child is spawned, it's PID is sent back to the requestor.
			parentProps := actor.FromFunc(func(ctx actor.Context) {
				switch m := ctx.Message().(type) {
				case string:
					if m == "spawn" && ctx.Sender() != nil {
						child := ctx.SpawnPrefix(childProps, "my-dear-dummy")
						ctx.Respond(child)
					}
				}
			})

			Convey("And using a parent without spawn interception", func() {
				parent, err := catch.Spawn(parentProps, options.OptNoInterception.WithPrefix("parent1"))
				So(err, ShouldBeNil)

				Convey("When sending a request to spawn a child", func() {
					res, err := parent.RequestFuture("spawn", options.DEFAULT_TIMEOUT).Result()

					Convey("Get a response with the child's PID", func() {
						So(err, ShouldBeNil)
						_, ok := res.(*actor.PID)
						So(ok, ShouldBeTrue)
					})
				})
			})

			Convey("And using a parent with spawn interception enabled", func() {
				parent, err := catch.Spawn(parentProps, options.OptNoInterception.WithSpawnInterception().WithPrefix("parent1"))
				So(err, ShouldBeNil)

				Convey("When sending a request to spawn a child", func() {
					f := actor.NewFuture(options.DEFAULT_TIMEOUT)
					parent.Request("spawn", f.PID())

					Convey("And trying to get a response right away", func() {
						_, err := f.Result()

						// We end up with an error because the parent has intercepted the spawning
						// and now is stuck waiting for us to see that.
						Convey("Then get a timeout error", func() {
							So(err, ShouldNotBeNil)
							So(err.Error(), ShouldContainSubstring, "timeout")
						})
					})

					Convey("And when asserting that the child is spawned (wrong name)", func() {
						res := catch.ShouldSpawn("foobar")

						Convey("Then get an assertion error", func() {
							So(res, ShouldContainSubstring, "PID does not match")
						})
					})

					// First, assert spawning
					// This will unblock the catcher (and thus the "parent")
					Convey("And when asserting that the child is spawned", func() {
						So(catch.ShouldSpawn("my-dear-dummy"), ShouldBeEmpty)

						// Second, get the result
						Convey("Then get a response with the child's PID", func() {
							res, err := f.Result()
							So(err, ShouldBeNil)

							child, ok := res.(*actor.PID)
							So(ok, ShouldBeTrue)
							So(child.String(), ShouldContainSubstring, "my-dear-dummy")
						})
					})
				})
			})
		})
	})
}
