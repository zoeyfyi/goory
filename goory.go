package goory

func assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}
