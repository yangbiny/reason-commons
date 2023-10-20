package slice

func MapTo[REQ, RESP any](req []*REQ, fn func(*REQ) RESP) []*RESP {
	if len(req) == 0 {
		return []*RESP{}
	}

	resp := make([]*RESP, len(req))
	for _, r := range req {
		r2 := fn(r)
		resp = append(resp, &r2)
	}
	return resp
}
