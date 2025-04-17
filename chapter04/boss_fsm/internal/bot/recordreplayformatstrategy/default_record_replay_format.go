package recordreplayformatstrategy

type DefaultRecordReplayFormat struct {
}

var _ RecordReplayFormatStrategy = &DefaultRecordReplayFormat{}

func NewDefaultRecordReplayFormat() *DefaultRecordReplayFormat {
	return &DefaultRecordReplayFormat{}
}

func (d *DefaultRecordReplayFormat) Format() string {
	return "[Record Replay] %s"
}
