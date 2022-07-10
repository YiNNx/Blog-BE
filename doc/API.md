# API Doc

## Basic Format:

- Address：` `
- API prefix：`/api/v1`    ( omitted below )

## Data Exchange Format:

- success:

```json
{
    "success": true,
    "message": "",
    "hint": "",
    "data": {
        ...
    }
}
```

- error:

```json
{
    "success": false,
    "message": "error message",
    "hint": "debug info",
    "data": {}
}
```

Error will return an http status code other than `200`.

## Authorization

Authenticate with Json Web Token. Be like:

```
Authorization: Bearer TOKEN
```

APIs is divided into three parts. View and Comment parts are available for everyone, while Post part needs Authorization. Sign-up is not opened publicly, and only the website owner  is able to sign in, post or revise posts and view the back-end data.

## **Paginate**

Most of the `GET` requests support pagination fuction by using query params `from` and `amount`.

For example:

`GET /post/all?from=0&amount=10`

## API - View & Comment

#### Get Posts With Query Params (outline)

`GET /post`

`year` - specified year

`month` - specified month

`keyword` - search

`tag` - get posts by tag

`from` & `amount` - paginate

Response

```
{
	[	{
            "pid": 0,
            "status": 0,	//0-Public | 1-Private | 2-Script
            "title": "",
            "time": "2022-01-01 20:00:00",
            "tag":["",""],
            "excerpt": "",
            "stats": {
                "likes": 0,
                "views": 0,
                "comments": 0
            }
        },
        ...
    ]
}
```

#### Get Post By Pid

`GET /post/<pid>`

Response

```json
{
    "pid": 0,
    "status": 0,	//0-Public | 1-Private | 2-Script
    "title": "",
    "time": "2022-01-01 20:00:00",
    "tag":["",""],
    "type": 0,	//0-PlainText | 1-Markdown | 2-HTML
    "content": "",
    "stats": {
        "likes": 0,
        "views": 0,
        "comments": 0
    }
}
```

#### Get Comments of Post

`GET /post/<pid>/comment`

Response

```
{
	[	{
            "cid": 1,
            "from": "",
            "from_url": "",
            "time": "2022-01-01 20:00:00",
            "content": "",
        },
        ...
    ]
}
```

#### Get All Tags

`GET /tag`

```
{
	["TAG","TAG",...]
}
```

#### Like

`PUT /post/<pid>/like`

Request

```
{
	"status":true
}
```

#### Comment

`POST /post/<pid>/comment`

Request

```
{
	"from":"",
	"email":"",
	"from_url":"",
	"content":""
}
```

Response

```
{
	"cid":1
}
```

#### Sub Comment

`POST /post/<pid>/comment`

Request

```
{
	"from":"",
	"parent_cid":"",
	"email":"",
	"from_url":"",
	"content":""
}
```

Response

```
{
	"cid":1
}
```

#### Stats

`GET /stats`

Response

```
{
	"posts":0,
	"timestamp":0,
    "views": 0,
    "likes": 0,
    "comments": 0
}
```

### API - Post

*Token Needed

#### New Post

`POST /post`

Request

```json
{
	"status":0,	//0-Public | 1-Private | 2-Script
	"title":"title",
	"excerpt":"",
	"type":0, //0-PlainText | 1-Markdown | 2-HTML
	"content":"content",
    "tag":["tag1","tag2"]
}
```

#### Update Post

`PUT /post/<pid>`

Request

```
{
	"status":0,	//0-Public | 1-Private | 2-Script
	"title":"title",
	"excerpt":"",
	"type":0, //0-PlainText | 1-Markdown | 2-HTML
	"content":"content",
    "tag":["tag1","tag2"]
}
```

#### Delete Post

`DELETE /post/<pid>`

Request

```
{
	"status":true
}
```

#### Detele Comment

`DELETE /comment/<cid>`

Request

```
{
	"status":true
}
```

#### Change Profile

`POST /info`

Request

```
{
	
}
```

