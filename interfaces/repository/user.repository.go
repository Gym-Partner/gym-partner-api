package repository

import (
    "context"
    "errors"
    "github.com/google/uuid"
    "gitlab.com/Titouan-Esc/api_common/logger"
    "gitlab.com/Titouan-Esc/api_common/middlewares"
    mongodb "gitlab.com/Titouan-Esc/api_common/mongo"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Sql *mongodb.Mongo
	Logger *logger.Log
}

func (u UserRepository) IsExist(data, OPT string) bool {
    var user model.User
    var queryColumn string

    switch OPT {
    case "ID":
        queryColumn = "id"
    case "EMAIL":
        queryColumn = "email"
    }

    filter := bson.D{{queryColumn, data}}

    if err := u.Sql.Database.Collection("users").FindOne(context.TODO(), filter).Decode(&user); err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return false
        }
    }

    if user.Id == "" {
        u.Logger.Warning("User not found in database")
        return false
    } else {
        return true
    }
}

func (u UserRepository) GetAll() (model.Users, error) {
    var users model.Users

    cursor, err := u.Sql.Database.Collection("users").Find(context.TODO(), bson.D{{}}, options.Find().SetProjection(model.UserProjection))
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return model.Users{}, err
        }
    }

    if err = cursor.All(context.TODO(), &users); err != nil {
        u.Logger.Error(err.Error())
        return model.Users{}, err
    }

    return users, nil
}

func (u UserRepository) GetOneById(uid string) (model.User, error) {
    var user model.User

    filter := bson.D{{"id", uid}}

    if err := u.Sql.Database.Collection("users").FindOne(context.TODO(), filter, options.FindOne().SetProjection(model.UserProjection)).Decode(&user); err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            u.Logger.Error(err.Error())
            return model.User{}, err
        }
    }

    return user, nil
}

func (u UserRepository) Create(data model.User) (model.User, error) {
    data.Id = uuid.New().String()
    data.Password = middlewares.HashPassword(data.Password)

    _, err := u.Sql.Database.Collection("users").InsertOne(context.TODO(), data)
    if err != nil {
        u.Logger.Error(err.Error())
        return model.User{}, err
    }
    
    return data, nil
}

func (u UserRepository) Update(data model.User) error {
    filter := bson.D{{"id", data.Id}}
    update := bson.D{{"$set", data}}

    _, err := u.Sql.Database.Collection("users").UpdateOne(context.TODO(), filter, update)
    if err != nil {
        u.Logger.Error(err.Error())
        return err
    }

    return nil
}

func (u UserRepository) Delete(uid string) error {
    //TODO implement me
    panic("implement me")
}
