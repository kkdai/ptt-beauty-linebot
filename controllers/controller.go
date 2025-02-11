package controllers

import (
	"errors"
	"fmt"
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

// Helper to create an article document given an index.
// Now checks error from GetAllFromURL (returns error string) and uses fallback values.
func createArticle(ptt *PTT, index int) favdb.ArticleDocument {
	title := ptt.GetPostTitleByIndex(index)
	url := ptt.GetPostUrlByIndex(index)
	post := favdb.ArticleDocument{
		ArticleTitle: title,
		URL:          url,
		ArticleID:    utils.GetPttIDFromURL(url),
	}
	// Check for error from GetAllFromURL using empty-string check.
	if errStr, images, push, boo := ptt.GetAllFromURL(url); errStr != "" {
		// Fallback values if error (such as too many redirects)
		post.ImageLinks = []string{}
		post.MessageCount.Push = 0
		post.MessageCount.Boo = 0
		post.MessageCount.All = 0
	} else {
		post.ImageLinks = images
		post.MessageCount.Push = push
		post.MessageCount.Boo = boo
		post.MessageCount.All = push + boo
	}
	return post
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
	fmt.Println("Total count:", count)
	for i := 0; i < count && i < perPage; i++ {
		title := ptt.GetPostTitleByIndex(i)
		if utils.CheckTitleWithBeauty(title) {
			ret = append(ret, createArticle(ptt, i))
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
		index := rands[i]  // use the random index
		post := createArticle(ptt, index)
		post.MessageCount.Count = ptt.GetPostStarByIndex(index)
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
		ret = append(ret, createArticle(ptt, i))
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
		url := ptt.GetPostUrlByIndex(i)
		post := favdb.ArticleDocument{
			ArticleTitle: title,
			URL:          url,
			ArticleID:    utils.GetPttIDFromURL(url),
			MessageCount: favdb.MessageCount{
				Count: ptt.GetPostStarByIndex(i),
			},
		}
		ret = append(ret, post)
	}
	// Sort articles by like count in descending order.
	sort.Sort(favdb.AllArticles(ret))
	// Get the top limit entries.
	ret = ret[0:limit]
	// Enrich each article with images and like details with error handling.
	for k := range ret {
		url := ret[k].URL
		if errStr, images, push, boo := ptt.GetAllFromURL(url); errStr != "" {
			ret[k].ImageLinks = []string{}
			ret[k].MessageCount.Push = 0
			ret[k].MessageCount.Boo = 0
			ret[k].MessageCount.All = 0
		} else {
			ret[k].ImageLinks = images
			ret[k].MessageCount.Push = push
			ret[k].MessageCount.Boo = boo
			ret[k].MessageCount.All = push + boo
		}
	}
	return ret, nil
}
