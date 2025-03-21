package main

import (
	"fmt"
	"slices"
)

type Channel struct {
	Name        string
	Videos      []Video
	Subscribers []ChannelSubscriber
}

func NewChannel(name string) *Channel {
	subscribers := make([]ChannelSubscriber, 0)
	return &Channel{
		Name:        name,
		Subscribers: subscribers,
	}
}

func (channel *Channel) Upload(video Video) {
	channel.Videos = append(channel.Videos, video)
	for _, sub := range channel.Subscribers {
		sub.HandleNotification(NewVideoNotification(channel, video))
	}
}

func (c *Channel) AddSubscriber(sub ChannelSubscriber) {
	c.Subscribers = append(c.Subscribers, sub)
	fmt.Println(sub.Name(), "訂閱了", c.Name)
}

func (c *Channel) RemoveSubscriber(sub ChannelSubscriber) {
	for i, s := range c.Subscribers {
		if s.Name() == sub.Name() {
			c.Subscribers = slices.Delete(c.Subscribers, i, i+1)
			fmt.Println(sub.Name(), "取消訂閱了", c.Name)
			break
		}
	}
}
