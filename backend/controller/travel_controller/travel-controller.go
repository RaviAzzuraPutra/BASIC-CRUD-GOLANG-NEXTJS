package travel_controller

import (
	"backend/config/cloudinary_config"
	"backend/database"
	"backend/helper"
	"backend/model"
	"backend/request"
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddTravel(ctx *gin.Context) {
	travelRequest := new(request.TravelRequest)

	if errReq := ctx.ShouldBind(&travelRequest); errReq != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errReq.Error(),
		})
		return
	}

	// menangkap file yang di upload
	file, errFile := ctx.FormFile("photo")

	if errFile != nil {
		ctx.JSON(400, gin.H{
			"Message": "Terjadi Kesalahan Saat Mengambil File",
			"Error":   errFile.Error(),
		})
		return
	}

	// membuka file yang di upload
	src, errSrc := file.Open()

	if errSrc != nil {
		ctx.JSON(400, gin.H{
			"Message": "Terjadi Kesalahan Saat Membuka File",
			"Error":   errSrc.Error(),
		})
		return
	}

	defer src.Close()

	// mengambil konfigurasi cloudinary dari config
	cloudinaryCFG := cloudinary_config.ConfigCloudinary()

	// inisialisasi cloudinary
	cloudinaryInit, errINIT := cloudinary.NewFromParams(cloudinaryCFG.CloudName, cloudinaryCFG.ApiKey, cloudinaryCFG.ApiSecret)
	if errINIT != nil {
		log.Printf("Terjadi Kesalahan Saat Inisialisasi Cloudinary: %v", errINIT)
		return
	}
	// pastikan menggunakan secure url (https)
	cloudinaryInit.Config.URL.Secure = true

	customFileName := "travel" + "-" + *travelRequest.Name + "-" + uuid.New().String()

	// mengupload file ke cloudinary
	uploadFile, errUploadFile := cloudinaryInit.Upload.Upload(context.Background(), src, uploader.UploadParams{
		Folder:       "BASIC-CRUD-GOLANG-NEXTJS",
		PublicID:     customFileName,
		ResourceType: "image",
	})

	if errUploadFile != nil {
		ctx.JSON(400, gin.H{
			"Message": "Terjadi Kesalahan Saat Mengupload File",
			"Error":   errUploadFile.Error(),
		})
		return
	}

	travel := new(model.Travel)
	travel.Name = travelRequest.Name
	travel.Description = travelRequest.Description
	photoURL := &uploadFile.SecureURL
	travel.Photo = photoURL
	travel.Price = travelRequest.Price

	errDB := database.DB.Table("travels").Create(&travel).Error

	if errDB != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Menyimpan Data",
			"Error":   errDB.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Menambahkan Data Travel",
		"Data":    travel,
	})

}

func GetAllTravel(ctx *gin.Context) {
	travels := new([]model.Travel)

	errDB := database.DB.Table("travels").Where("deleted_at IS NULL").Order("created_at asc").Find(&travels).Error

	if errDB != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Mengambil Data",
			"Error":   errDB.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Mendapatkan Data Travel",
		"Data":    travels,
	})
}

func GetTravelByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	travel := new(model.Travel)

	errDB := database.DB.Table("travels").Where("id = ? AND deleted_at IS NULL", ID).Find(&travel).Error

	if errDB != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Mengambil Data",
			"Error":   errDB.Error(),
		})
		return
	}

	if travel.Id == nil {
		ctx.JSON(404, gin.H{
			"Message": "Data Tidak Ditemukan",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Mendapatkan Data Travel Berdasarkan ID",
		"Data":    travel,
	})
}

func DeleteTravel(ctx *gin.Context) {
	ID := ctx.Param("id")

	travel := new(model.Travel)

	errFind := database.DB.Table("travels").Where("id = ? AND deleted_at IS NULL", ID).Find(&travel).Error

	if errFind != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Mengambil Data",
			"Error":   errFind.Error(),
		})
		return
	}

	if travel.Id == nil {
		ctx.JSON(404, gin.H{
			"Message": "Data Tidak Ditemukan",
		})
		return
	}

	errDB := database.DB.Table("travels").Where("id = ?", ID).Delete(&travel).Error

	if errDB != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Menghapus Data",
			"Error":   errDB.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Menghapus Data Travel",
	})

}

func UpdateTravel(ctx *gin.Context) {
	ID := ctx.Param("id")

	travelRequest := new(request.TravelRequest)

	if errReq := ctx.ShouldBind(&travelRequest); errReq != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errReq.Error(),
		})
		return
	}

	travel := new(model.Travel)

	errFind := database.DB.Table("travels").Where("id = ? AND deleted_at IS NULL", ID).Find(&travel).Error

	if errFind != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Mengambil Data",
			"Error":   errFind.Error(),
		})
		return
	}

	if travel.Id == nil {
		ctx.JSON(404, gin.H{
			"Message": "Data Tidak Ditemukan",
		})
		return
	}

	file, errFile := ctx.FormFile("photo")

	if errFile == nil {
		src, errSrc := file.Open()

		if errSrc != nil {
			ctx.JSON(400, gin.H{
				"Message": "Terjadi Kesalahan Saat Membuka File",
				"Error":   errSrc.Error(),
			})
			return
		}

		defer src.Close()

		cloudinaryCFG := cloudinary_config.ConfigCloudinary()

		cloudinaryInit, errINIT := cloudinary.NewFromParams(cloudinaryCFG.CloudName, cloudinaryCFG.ApiKey, cloudinaryCFG.ApiSecret)
		if errINIT != nil {
			log.Printf("Terjadi Kesalahan Saat Inisialisasi Cloudinary: %v", errINIT)
			return
		}

		cloudinaryInit.Config.URL.Secure = true

		if travel.Photo != nil {
			publicID := helper.ExtractPublicID(*travel.Photo)
			_, errDel := cloudinaryInit.Upload.Destroy(context.Background(), uploader.DestroyParams{
				PublicID:     publicID,
				ResourceType: "image",
			})

			if errDel != nil {
				log.Print("Terjadi Kesalahan Saat Menghapus File: ", errDel.Error())
			}
		}

		customFileName := "travel" + "-" + *travelRequest.Name + "-" + uuid.New().String()

		uploadFile, errUploadFile := cloudinaryInit.Upload.Upload(context.Background(), src, uploader.UploadParams{
			Folder:       "BASIC-CRUD-GOLANG-NEXTJS",
			PublicID:     customFileName,
			ResourceType: "image",
		})

		if errUploadFile != nil {
			ctx.JSON(400, gin.H{
				"Message": "Terjadi Kesalahan Saat Mengupload File",
				"Error":   errUploadFile.Error(),
			})
			return
		}

		travel.Photo = &uploadFile.SecureURL
	}

	travel.Name = travelRequest.Name
	travel.Description = travelRequest.Description
	travel.Price = travelRequest.Price

	errDB := database.DB.Table("travels").Where("id = ?", ID).Updates(&travel).Error

	if errDB != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Mengupdate Data",
			"Error":   errDB.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Mengupdate Data Travel",
		"Data":    travel,
	})

}
