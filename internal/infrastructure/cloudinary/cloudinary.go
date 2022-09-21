package cloudinary

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"time"
)

func ImageUpload(input string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(utils.AddToStr(configs.Instance.CloudinaryName), utils.AddToStr(configs.Instance.CloudinaryAPIKey), utils.AddToStr(configs.Instance.CloudinarySecret))
	if err != nil {

		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{
		Folder: string(constant.Directory),
	})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
