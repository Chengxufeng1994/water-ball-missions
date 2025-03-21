package main

import "fmt"

type Namer interface {
	Name() string
}

type ChannelSubscriber interface {
	Namer
	HandleNotification(notification VideoNotification)
}

type BasedSubscriber struct {
	name string
}

func NewBasedSubscriber(name string) *BasedSubscriber {
	return &BasedSubscriber{
		name: name,
	}
}

func (b *BasedSubscriber) Name() string {
	return b.name
}

// WaterBallSubscriber (水球訂閱者)
type WaterBallSubscriber struct {
	*BasedSubscriber
}

var _ ChannelSubscriber = (*WaterBallSubscriber)(nil)

func NewWaterBallSubscriber() *WaterBallSubscriber {
	return &WaterBallSubscriber{
		BasedSubscriber: NewBasedSubscriber("水球"),
	}
}

func (w *WaterBallSubscriber) HandleNotification(notification VideoNotification) {
	if notification.Video.Length >= 180 {
		notification.Video.Like()
		fmt.Printf("%s 對影片 \"%s\" 按讚\n", w.Name(), notification.Video.Title)
	}
}

// FireBallSubscriber (火球訂閱者)
type FireBallSubscriber struct {
	*BasedSubscriber
}

var _ ChannelSubscriber = (*FireBallSubscriber)(nil)

func NewFireBallSubscriber() *FireBallSubscriber {
	return &FireBallSubscriber{
		BasedSubscriber: NewBasedSubscriber("火球"),
	}
}

func (f *FireBallSubscriber) HandleNotification(notification VideoNotification) {
	if notification.Video.Length <= 60 {
		notification.Channel.RemoveSubscriber(f)
	}
}
