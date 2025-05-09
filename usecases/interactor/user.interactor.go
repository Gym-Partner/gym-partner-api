package interactor

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gitlab.com/gym-partner1/api/gym-partner-api/core/awsService"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type IUserInteractor interface {
	Create(ctx *gin.Context) (model.User, *core.Error)
	GetAll() (model.Users, *core.Error)
	GetOne(c *gin.Context) (model.User, *core.Error)
	GetOneByEmail(ctx *gin.Context) (model.User, *core.Error)
	Search(query string, limit, offset int) (model.Users, *core.Error)
	UploadImage(ctx *gin.Context) (model.UserImage, *core.Error)
	Update(ctx *gin.Context) *core.Error
	Delete(ctx *gin.Context) *core.Error
}

type UserInteractor struct {
	IUserRepository   repository.IUserRepository
	IFollowRepository repository.IFollowRepository
	IUtils            utils.IUtils[model.User]
}

func MockUserInteractor(userMock *mock.UserInteractorMock, utilsMock *mock.UtilsMock[model.User]) *UserInteractor {
	return &UserInteractor{
		IUserRepository: userMock,
		IUtils:          utilsMock,
	}
}

// -------------------------- CRUD ------------------------------

func (ui *UserInteractor) Create(ctx *gin.Context) (model.User, *core.Error) {
	data, err := ui.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.User{}, err
	}

	exist := ui.IUserRepository.IsExist(data.Email, "EMAIL")
	if exist {
		return model.User{}, core.NewError(
			http.StatusBadRequest,
			fmt.Sprintf(core.ErrAppINTUserExist, data.Email))
	}

	data.Id = ui.IUtils.GenerateUUID()

	data.Password, _ = ui.IUtils.HashPassword(data.Password)
	user, err := ui.IUserRepository.Create(data)
	return user, err
}

func (ui *UserInteractor) GetAll() (model.Users, *core.Error) {
	users, err := ui.IUserRepository.GetAll()
	if err != nil {
		return model.Users{}, err
	}

	// Followers part and user's part
	for key, user := range users {
		followers, err := ui.IFollowRepository.GetAllByUserId(user.Id)
		if err != nil {
			return model.Users{}, err
		}

		userImage, err := ui.IUserRepository.GetImageByUserId(user.Id)
		if err != nil {
			return model.Users{}, err
		}

		users[key].Followers = followers.Followers
		users[key].Following = followers.Followings
		users[key].UserImage = userImage.ImageURL
	}

	return users, nil
}

func (ui *UserInteractor) GetOne(c *gin.Context) (model.User, *core.Error) {
	uid, _ := c.Get("uid")

	user, err := ui.IUserRepository.GetOneById(uid.(string))
	if err != nil {
		return model.User{}, err
	}

	// Followers part
	followers, err := ui.IFollowRepository.GetAllByUserId(uid.(string))
	if err != nil {
		return model.User{}, err
	}

	// User's image part
	userImage, err := ui.IUserRepository.GetImageByUserId(uid.(string))
	if err != nil {
		return model.User{}, err
	}

	user.Followers = followers.Followers
	user.Following = followers.Followings
	user.UserImage = userImage.ImageURL
	return user, nil
}

func (ui *UserInteractor) GetOneByEmail(ctx *gin.Context) (model.User, *core.Error) {
	data, err := ui.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.User{}, err
	}

	user, err := ui.IUserRepository.GetOneByEmail(data.Email)
	if err != nil {
		return model.User{}, err
	}

	// Followers part
	followers, err := ui.IFollowRepository.GetAllByUserId(data.Id)
	if err != nil {
		return model.User{}, err
	}

	// User's image part
	userImage, err := ui.IUserRepository.GetImageByUserId(data.Id)
	if err != nil {
		return model.User{}, err
	}

	user.Password = data.Password
	user.Followers = followers.Followers
	user.Following = followers.Followings
	user.UserImage = userImage.ImageURL
	return user, err
}

func (ui *UserInteractor) Search(query string, limit, offset int) (model.Users, *core.Error) {
	users, err := ui.IUserRepository.Search(query, limit, offset)
	if err != nil {
		return model.Users{}, err
	}

	// Followers part
	for key, user := range users {
		followers, err := ui.IFollowRepository.GetAllByUserId(user.Id)
		if err != nil {
			return model.Users{}, err
		}

		users[key].Followers = followers.Followers
		users[key].Following = followers.Followings
	}

	return users, nil
}

func (ui *UserInteractor) UploadImage(ctx *gin.Context) (model.UserImage, *core.Error) {
	uid, _ := ctx.Get("uid")
	file, err := ctx.FormFile("image")
	if err != nil {
		return model.UserImage{}, core.NewError(
			http.StatusNotAcceptable,
			fmt.Sprintf(core.ErrAppINTUserImageNotFound, uid),
			err)
	}

	filename := uid.(string) + filepath.Ext(file.Filename)

	src, err := file.Open()
	if err != nil {
		return model.UserImage{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppINTUserImageNotOpen, uid),
			err)
	}
	defer src.Close()

	// Initialization AWS services
	awsSess := awsService.NewAWSService()
	s3Client := awsService.NewAWSS3(awsSess)
	bucketName := viper.GetString("AWS_S3_BUCKET_NAME")

	exist := ui.IUserRepository.UserImageIsExist(uid.(string))
	if exist {
		// Remove image in database and S3 service
		_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(filename),
		})
		if err != nil {
			return model.UserImage{}, core.NewError(
				http.StatusInternalServerError,
				fmt.Sprintf(core.ErrAppINTUserImageDeleteS3, uid),
				err)
		}

		if err = ui.IUserRepository.DeleteUserImage(uid.(string)); err != nil {
			return model.UserImage{}, core.NewError(
				http.StatusInternalServerError,
				fmt.Sprintf(core.ErrAppINTUserImageDeletePsql, uid),
				err)
		}
	}

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   src,
		//ContentType: aws.String(file.Header.Get("Content-Type")),
	})
	if err != nil {
		log.Println(err.Error())
		return model.UserImage{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppINTUserImageUpload, uid),
			err)
	}

	imageURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s",
		bucketName,
		viper.GetString("AWS_REGION"),
		filename)

	userImage := model.UserImage{
		Id:        uuid.New().String(),
		UserId:    uid.(string),
		ImageURL:  imageURL,
		CreatedAt: time.Now(),
	}

	if err := ui.IUserRepository.UploadImage(userImage); err != nil {
		return model.UserImage{}, err
	}

	return userImage, nil
}

func (ui *UserInteractor) Update(ctx *gin.Context) *core.Error {
	uid, _ := ctx.Get("uid")
	patch, err := ui.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return err
	}
	patch.Id = *uid.(*string)

	exist := ui.IUserRepository.IsExist(patch.Id, "ID")
	if !exist {
		return core.NewError(
			http.StatusBadRequest,
			fmt.Sprintf(core.ErrAppINTUserNotExist, patch.Id))
	}

	target, err := ui.IUserRepository.GetOneById(patch.Id)
	if err != nil {
		return err
	}

	if err = ui.IUtils.Bind(&target, patch); err != nil {
		return err
	}

	err = ui.IUserRepository.Update(target)
	return err
}

func (ui *UserInteractor) Delete(ctx *gin.Context) *core.Error {
	uid, _ := ctx.Get("uid")

	exist := ui.IUserRepository.IsExist(*uid.(*string), "ID")
	if !exist {
		return core.NewError(
			http.StatusBadRequest,
			fmt.Sprintf(core.ErrAppINTUserNotExist, *uid.(*string)))
	}

	err := ui.IUserRepository.Delete(*uid.(*string))
	return err
}
