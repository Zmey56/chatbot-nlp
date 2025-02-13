package client

type NLPClient interface {
	SendRequest(text string) (string, error)
}
