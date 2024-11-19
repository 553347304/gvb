package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/INIT/global"
	"gvb_server/models"
	"io"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
)

type imageList struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Size      string `json:"size"`
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

func save(c *gin.Context, file *multipart.FileHeader, filePath string) (string, bool) {
	// 判断文件后缀名
	suffixName := strings.TrimPrefix(filepath.Ext(file.Filename), ".") // 去除.---文件后缀名
	suffix := strings.ToLower(suffixName)
	// 文件小写
	if !BY.Map.IsList(suffix, WhiteImageList) {
		return "非法扩展名", false
	}

	// 查询图片是否已存在
	fileObj, _ := file.Open()
	byteData, _ := io.ReadAll(fileObj)
	imageHash := BY.Data.Md5(byteData)

	var bannerModel models.BannerModel
	is := FK_SQL.Is(&bannerModel, "hash = ?", imageHash)

	if is {
		return "图片已存在", true
	}

	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
		return "保存文件失败", false
	}

	// 图片入库
	global.DB.Create(&models.BannerModel{
		Path: "/" + filePath,
		Hash: imageHash,
		Name: file.Filename,
	})
	return "上传成功", true
}

// ImageUploadView 图片上传
func (ImagesApi) ImageUploadView(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		BY.JSONList(c, 0, imageList{Message: "绑定错误"})
		return
	}

	var total int
	var list = make([]imageList, 0)

	fileList, _ := form.File["image"]
	for i, file := range fileList {
		filePath := path.Join(global.Config.Upload.Path, file.Filename) // 文件路径
		size := float64(file.Size) / (1024 * 1024)                      // 文件大小
		message, isSuccess := save(c, file, filePath)

		list = append(list, imageList{
			Name:      file.Filename,
			Path:      "/" + filePath,
			Size:      fmt.Sprintf("%.2fMB", size),
			Message:   message,
			IsSuccess: isSuccess,
		})
		total = i
	}

	BY.JSONList(c, total, list)
}
