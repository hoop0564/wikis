package ch40_pipe_filter

type Request interface {
}

type Response interface {
}

type Filter interface {
	Process(data Request) (Response, error)
}
