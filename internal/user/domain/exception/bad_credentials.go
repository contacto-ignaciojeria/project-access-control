package exception

type BadCredentials struct{}

func Error() string {
	return "bad credentials"
}