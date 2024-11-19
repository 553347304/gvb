package images_api

type ImagesApi struct {
}

const msg = "图片"

var (
	// WhiteImageList 图片上传白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"bmp",
		"ico",
		"gif",
	}
)
