package models

type UserFavorite struct {
	Id        int64    `bson:"_id"`
	UserId    string   `json:"user_id" bson:"user_id"`
	Favorites []string `json:"favorites" bson:"favorites"`
}

type UserFavData interface {
	Add(meta *Model)
	Get(meta *Model) (result *UserFavorite, err error)
	Update(meta *Model) (err error)
	GetAll(meta *Model) (result []UserFavorite, err error)
}
