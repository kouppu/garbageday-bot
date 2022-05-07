# garbageday-bot

ごみ収集日をLINEに通知するBOTです。  
`schedule.json`に、該当するごみ収集日があれば通知します。

## Installation

```bash:
$ git clone https://github.com/suhrr/garbageday-bot.git
```

## Configuration

### LINE Messaging API のアクセス情報を設定

`.env.example`をコピーし、`.env`を作成。  
作成した`.env`に、アクセス情報を記述してください。  
_Heroku などのサービスで、定義されている環境変数があれば、サービスの値が優先されて使用されます。_

### ごみ収集日の設定

`schedule.json`に、ごみ収集日を次のように記述してください。

| Name           | Type   | Detail                                                    |
| -------------- | ------ | --------------------------------------------------------- |
| name           | String | ごみ収集日名                                              |
| weekday        | Int    | 曜日番号 （日 0 月 1 火 2 水 3 木 4 金 5 土 6 ) |
| weekdayInMonth | Int    | 何周目の曜日かを設定 （0なら毎週）                           |

```json:schedule.json
{
  "garbageDays": [
    {
      "name": "普通ごみ",
      "weekday": 1,
      "weekdayInMonth": 0
    },
    {
      "name": "ペットボトル回収",
      "weekday": 2,
      "weekdayInMonth": 1
    },
  ]
}
```

## Usage
```bash:
# 実行
$ bin/garbageday-bot
```

```bash:
# ビルドする場合
$ go build -o bin/garbageday-bot -v .
```

## Test
```bash:
$ go test -v ./weekday/...
```
