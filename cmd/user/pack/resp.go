package pack

//
//// // BuildBaseResp build baseResp from error
//func BuildBaseResp(err error) user.DouyinUserRegisterResponse {
//	if err == nil {
//		return user.DouyinUserResponse
//	}
//	e := errno.ErrNo{}
//	if errors.As(err, &e) {
//		return
//	}
//
//	s := errno.ServiceErr.WithMessage(err.Error())
//	return baseResp(s)
//}
