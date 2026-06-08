package selectstmt

func FirstResult(ch1, ch2 <-chan string) string {
	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	}
}
