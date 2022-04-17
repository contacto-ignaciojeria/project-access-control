package exception

type BadCredentials struct{}

func (e BadCredentials) Error() string {
	return "bad credentials"
}