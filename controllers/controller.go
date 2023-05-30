package controllers

import (
	"errors"
	"sort"

	"github.com/kkdai/favdb"

	"github.com/kkdai/linebot-ptt-beauty/utils"
	. "github.com/kkdai/photomgr"
)

type UserFavorite struct {
	Id        int64    `bson:"_id"`
	UserId    string   `json:"user_id" bson:"user_id"`
	Favorites []string `json:"favorites" bson:"favorites"`
}

func GetOne(url string) (result *favdb.ArticleDocument, err error) {
	ptt := NewPTT()
	post := favdb.ArticleDocument{}
	post.URL = url
	post.ArticleID = utils.GetPttIDFromURL(url)
	post.ArticleTitle = ptt.GetUrlTitle(url)
	post.ImageLinks = ptt.GetAllImageAddress(url)
	like, dis := ptt.GetPostLikeDis(url)
	post.MessageCount.Push = like
	post.MessageCount.Boo = dis
	post.MessageCount.All = like + dis
	post.MessageCount.Count = like - dis
	return &post, nil
}

func Get(page int, perPage int) (results []favdb.ArticleDocument, err error) {
	var ret []favdb.ArticleDocument
	ptt := NewPTT()
	count := ptt.ParsePttPageByIndex(page, true)
	for i := 0; i < count && i < perPage; i++ {
		title := ptt.GetPostTitleByIndex(i)
		if utils.CheckTitleWithBeauty(title) {
			post := favdb.ArticleDocument{}
			url := ptt.GetPostUrlByIndex(i)
			post.ArticleTitle = title
			post.URL = url
			post.ArticleID = utils.GetPttIDFromURL(url)
			_, post.ImageLinks, post.MessageCount.Push, post.MessageCount.Boo = ptt.GetAllFromURL(url)
			post.MessageCount.All = post.MessageCount.Push + post.MessageCount.Boo
			post.MessageCount.Count = ptt.GetPostStarByIndex(i)
			ret = append(ret, post)
		}
	}

	return ret, nil
}

func GetRandom(count int) (results []favdb.ArticleDocument, err error) {
	rands := utils.GetRandomIntSet(100, 10)
	ptt := NewPTT()
	pCount := ptt.ParsePttByNumber(101, 0)
	if pCount == 0 {
		return nil, errors.New("NotFound")
	}
	var ret []favdb.ArticleDocument
	for i := 0; i < count; i++ {
		title := ptt.GetPostTitleByIndex(rands[i])
		post := favdb.ArticleDocument{}
		url := ptt.GetPostUrlByIndex(rands[i])
		post.ArticleTitle = title
		post.URL = url
		post.ArticleID = utils.GetPttIDFromURL(url)
		_, post.ImageLinks, post.MessageCount.Push, post.MessageCount.Boo = ptt.GetAllFromURL(url)
		post.MessageCount.All = post.MessageCount.Push + post.MessageCount.Boo
		post.MessageCount.Count = ptt.GetPostStarByIndex(i)
		ret = append(ret, post)
	}
	return ret, nil
}

func GetKeyword(count int, keyword string) (results []favdb.ArticleDocument, err error) {

	ptt := NewPTT()
	pCount := ptt.ParseSearchByKeyword(keyword)
	if pCount == 0 {
		return nil, errors.New("NotFound")
	}
	var ret []favdb.ArticleDocument
	for i := 0; i < count; i++ {
		title := ptt.GetPostTitleByIndex(i)
		post := favdb.ArticleDocument{}
		url := ptt.GetPostUrlByIndex(i)
		post.ArticleTitle = title
		post.URL = url
		post.ArticleID = utils.GetPttIDFromURL(url)
		_, post.ImageLinks, post.MessageCount.Push, post.MessageCount.Boo = ptt.GetAllFromURL(url)
		post.MessageCount.All = post.MessageCount.Push + post.MessageCount.Boo
		post.MessageCount.Count = ptt.GetPostStarByIndex(i)
		ret = append(ret, post)
	}
	return ret, nil
}

func GetMostLike(total int, limit int) (results []favdb.ArticleDocument, err error) {
	ptt := NewPTT()
	pCount := ptt.ParsePttByNumber(total, 0)
	if pCount == 0 {
		return nil, errors.New("NotFound")
	}

	var ret []favdb.ArticleDocument
	for i := 0; i < pCount; i++ {
		title := ptt.GetPostTitleByIndex(i)
		post := favdb.ArticleDocument{}
		url := ptt.GetPostUrlByIndex(i)
		post.ArticleTitle = title
		post.URL = url
		post.ArticleID = utils.GetPttIDFromURL(url)
		post.MessageCount.Count = ptt.GetPostStarByIndex(i)
		ret = append(ret, post)
	}
	//Sort it.
	sort.Sort(favdb.AllArticles(ret))

	//Get the first limit (10)
	ret = ret[0:limit]

	//Add each images and like/dis to top 10.
	for k, _ := range ret {
		url := ret[k].URL
		_, ret[k].ImageLinks, ret[k].MessageCount.Push, ret[k].MessageCount.Boo = ptt.GetAllFromURL(url)
		ret[k].MessageCount.All = ret[k].MessageCount.Push + ret[k].MessageCount.Boo
	}

	return ret, nil
}
