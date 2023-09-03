package oss

const (
	QI_NIU_BASE_URL = "https://demo.com/"
)

func GetImgUrl(id string) string {
	return QI_NIU_BASE_URL + id
}
