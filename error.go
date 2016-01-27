package mecab

type mecabError struct {
	message string
}

func newMecabError(message string) mecabError {
	return mecabError{
		message: message,
	}
}

func (err mecabError) Error() string {
	return err.message
}
