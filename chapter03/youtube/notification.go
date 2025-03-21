package main

type VideoNotification struct {
	Channel *Channel
	Video   Video
}

func NewVideoNotification(channel *Channel, video Video) VideoNotification {
	return VideoNotification{
		Channel: channel,
		Video:   video,
	}
}
