package utility

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"time"
)

func ImageUpload(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(configs.Instance.EnvCloudName, configs.Instance.EnvCloudAPIKey, configs.Instance.EnvCloudAPISecret)
	if err != nil {

		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: configs.Instance.EnvCloudUploadFolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
