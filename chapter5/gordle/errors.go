package gordle

type corpusError string

func (e corpusError) Error() string {
	return string(e)
}
