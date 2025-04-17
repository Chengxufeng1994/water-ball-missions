package knowledgeking

type Question struct {
	Problem string
	Answer  string
	Options map[string]string
}

func NewQuestion(problem string, answer string, options map[string]string) *Question {
	return &Question{
		Problem: problem,
		Answer:  answer,
		Options: options,
	}
}
