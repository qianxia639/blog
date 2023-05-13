package oss

type UploadImageStrategy interface {
	UploadImage(localfile string) (string, error)
}

type Upload struct {
	Strategy  UploadImageStrategy
	LocalFile string
}

func (u *Upload) UploadFile() (string, error) {
	return u.Strategy.UploadImage(u.LocalFile)
}
