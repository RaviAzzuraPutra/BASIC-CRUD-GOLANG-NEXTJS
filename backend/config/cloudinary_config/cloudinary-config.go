package cloudinary_config

import "os"

type CloudinaryAPI struct {
	CloudName string
	ApiKey    string
	ApiSecret string
}

func ConfigCloudinary() *CloudinaryAPI {
	return &CloudinaryAPI{
		CloudName: os.Getenv("CLOUDINARY_NAME"),
		ApiKey:    os.Getenv("CLOUDINARY_API_KEY"),
		ApiSecret: os.Getenv("CLOUDINARY_API_SECRET"),
	}
}
