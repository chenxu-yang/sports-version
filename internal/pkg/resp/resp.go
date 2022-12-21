package resp

type Resp struct {
	Code int32
	Msg  string
	Data interface{}
}

func ToStruct(data interface{}, err error) *Resp {
	if err == nil {
		return &Resp{
			Code: 0,
			Msg:  "success",
			Data: data,
		}
	}
	return &Resp{
		Code: -1,
		Msg:  err.Error(),
		Data: nil,
	}
}
