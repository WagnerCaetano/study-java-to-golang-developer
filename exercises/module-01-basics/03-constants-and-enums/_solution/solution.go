package constants

type Status int

const (
	StatusOK      Status = 200
	StatusNotFound Status = 404
	StatusError   Status = 500
)

func StatusMessage(s Status) string {
	messages := map[Status]string{
		StatusOK:      "OK",
		StatusNotFound: "Not Found",
		StatusError:   "Internal Server Error",
	}
	if msg, ok := messages[s]; ok {
		return msg
	}
	return "Unknown"
}
