package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"os"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community"
)

func main() {
	waterBallCommunity := community.NewCommunity("Water Ball")
	lines := LoadTestCaseFile("testcases/unit-tests/Forum_BotComment_10Users.in")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "]", 2)
		eventName := strings.TrimPrefix(parts[0], "[")
		var raw string
		if len(parts) == 1 {
			raw = ""
		} else {
			raw = strings.TrimSpace(parts[1])
		}
		communityEvent, err := ParseCommunityEvent(eventName, []byte(raw))
		if err != nil {
			log.Fatalf("failed to parse event: %v", err)
		}

		waterBallCommunity.HandleCommunityEvent(communityEvent)
	}
}

func LoadTestCaseFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func ParseCommunityEvent(event string, data []byte) (community.CommunityEvent, error) {
	switch event {
	case "started":
		var payload community.StartedEventPayload
		err := json.Unmarshal(data, &payload)
		if err != nil {
			return nil, err
		}
		return community.NewStartedEvent(payload), nil
	case "login":
		var payload community.LoginEventPayload
		err := json.Unmarshal(data, &payload)
		if err != nil {
			return nil, err
		}
		return community.NewLoginEvent(payload), nil
	case "logout":
		var payload community.LogoutEventPayload
		err := json.Unmarshal(data, &payload)
		if err != nil {
			return nil, err
		}
		return community.NewLogoutEvent(payload), nil
	case "new message":
		var payload community.NewMessageEventPayload
		err := json.Unmarshal(data, &payload)
		if err != nil {
			return nil, err
		}
		return community.NewNewMessageEvent(payload), nil
	case "new post":
		var payload community.NewPostEventPayload
		err := json.Unmarshal(data, &payload)
		if err != nil {
			if err != nil {
				return nil, err
			}
		}
		return community.NewNewPostEvent(payload), nil
	case "end":
		return community.NewEndEvent(), nil
	}
	return nil, fmt.Errorf("unknown event: %s", event)
}
