package oss

type UploadImageStrategy interface {
	UploadImage(localfile string) (string, error)
}

type Upload struct {
	ImageStrategy UploadImageStrategy
	LocalFile     string
}

func (u *Upload) UploadImage() (string, error) {
	return u.ImageStrategy.UploadImage(u.LocalFile)
}
