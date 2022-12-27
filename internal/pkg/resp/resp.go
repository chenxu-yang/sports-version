package resp

type Resp struct {
	code int32
	msg  string
	data interface{}
}

func ToStruct(data interface{}, err error) *Resp {
	if err == nil {
		return &Resp{
			code: 0,
			msg:  "success",
			data: data,
		}
	}
	return &Resp{
		code: -1,
		msg:  err.Error(),
		data: nil,
	}
}
