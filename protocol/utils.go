package protocol

func Options[O any](options ...O) O {
	var opts O
	if len(options) > 0 {
		opts = options[0]
	}
	return opts
}
