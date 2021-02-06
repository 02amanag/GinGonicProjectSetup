package usecase

type UsecaseStruct struct {
}

func NewUsecaseStruct() *UsecaseStruct {
	return &UsecaseStruct{}
}

func (u *UsecaseStruct) DetailsForm(param string) string {
	param += "added the latest"
	return param
}
