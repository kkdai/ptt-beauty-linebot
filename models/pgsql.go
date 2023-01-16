package models

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PGSql struct {
	Db   *pg.DB
	data UserFavorite
}

func NewPGSql(url string) *PGSql {
	options, _ := pg.ParseURL(url)
	db := pg.Connect(options)

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return &PGSql{
		Db: db,
	}
}

func (u *PGSql) Add(meta *Model) {
	_, err := u.Db.Model(u).Insert()
	if err != nil {
		log.Println(err)
	}
}

func (u *PGSql) Get(meta *Model) (result *UserFavorite, err error) {
	log.Println("***Get Fav uUID=", u.data.UserId)
	userFav := UserFavorite{}
	err = u.Db.Model(&userFav).
		Where("user_id = ?", u.data.UserId).
		Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	u.data = userFav
	log.Println("UserFavorite DB result= ", u.data)
	return &userFav, nil
}

func (u *PGSql) GetAll(meta *Model) (result []UserFavorite, err error) {
	log.Println("***Get All DB")
	users := []UserFavorite{}
	err = u.Db.Model(&users).Select()
	if err != nil {
		log.Println(err)
	}
	log.Println("***Start server all users =", users)
	if err != nil {
		log.Println("open file error !")
	}

	return users, nil
}

func (u *PGSql) Update(meta *Model) (err error) {
	log.Println("***Update Fav User=", u)

	_, err = u.Db.Model(u.data).
		Set("favorites = ?", u.data.Favorites).
		Where("user_id = ?", u.data.UserId).
		Update()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (u *PGSql) ShowAll(meta *Model) (err error) {
	log.Println("***ShowAll  User -->")
	err = u.Db.Model(&u.data).Select()
	if err != nil {
		log.Println(err)
	}
	log.Println("***Show all users =", u.data)
	return nil
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*UserFavorite)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true})
		if err != nil {
			return err
		}
	}
	return nil
}
