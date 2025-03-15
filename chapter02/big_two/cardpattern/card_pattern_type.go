package cardpattern

type CardPatternType string

const (
	CardPatternTypeSingle    CardPatternType = "Single"
	CardPatternTypePair      CardPatternType = "Pair"
	CardPatternTypeStraight  CardPatternType = "Straight"
	CardPatternTypeFullHouse CardPatternType = "FullHouse"
)
