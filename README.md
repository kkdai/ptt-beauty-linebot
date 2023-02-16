# 表特看看 - LINE 聊天機器人 for PTT Beauty

 [![GoDoc](https://godoc.org/github.com/kkdai/ptt-beauty-linebot.svg?status.svg)](https://godoc.org/github.com/kkdai/ptt-beauty-linebot) [![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/ptt-beauty-linebot)](https://goreportcard.com/report/github.com/kkdai/ptt-beauty-linebot) [![Go](https://github.com/kkdai/ptt-beauty-linebot/actions/workflows/go.yml/badge.svg)](https://github.com/kkdai/ptt-beauty-linebot/actions/workflows/go.yml)

# Supported Features

- Easy use LINE Bot to check PTT Beauty board.
- Save your favorite ptt posts.
- For developer, you could easily to switch pgsql or memory DB by add environment "DATABASE_URL".
- For developer, you could use Github Issue as DB to storage your favorite links by add environment "GITHUB_URL".

# Support DB

- [x] Memory DB
- [x] PostgresSQL  
- [x] Github issue.
- [ ] Firestore.  [#1](https://github.com/kkdai/ptt-beauty-linebot/issues/1)

# Working in Progress Features

- [x] Search keyword on ptt. [#17](https://github.com/kkdai/ptt-beauty-linebot/issues/17)

# How to use it

## For User

### 掃描 QR Code 或點選連結

[<img src="resource/qr_code.png">](https://line.me/R/ti/p/SFXWQpzdaY)

## For Developer

### Installation and Usage

#### Deploy on Web Platform

- Deploy on [Heroku](https://heroku.com)

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

- Deploy on [Reder.com](https://render.com)

[![Deploy to Render](http://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)

More detail, please check my [LINE Bot Template project](https://github.com/kkdai/LineBotTemplate).

### 截圖

- 功能選單

<img src="resource/screen1.jpg" height="480">

- 熱門照片

<img src="resource/screen2.jpg" height="480">

- 對話直接搜尋

<img src="resource/screen3.jpg" height="480">
