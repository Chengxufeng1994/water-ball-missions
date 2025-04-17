package recordreplayformatstrategy

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot"

type DefaultRecordReplayFormat struct {
}

var _ bot.RecordReplayFormatStrategy = &DefaultRecordReplayFormat{}

func NewDefaultRecordReplayFormat() *DefaultRecordReplayFormat {
	return &DefaultRecordReplayFormat{}
}

func (d *DefaultRecordReplayFormat) Format() string {
	return "default"
}
