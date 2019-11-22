# 博客网站API

## 用户注册

```
POST /api/users
```

### Input

| 项目名称    | 类型   | 描述         |
| ----------- | ------ | ------------ |
| username    | string | 用户名       |
| password    | string | 密码         |
| captchaId   | string | 图形验证码id |
| captchaCode | string | 用户识别结果 |

### Example

```
{
	"username": "Alice",
	"password": "123456",
	"captchaId": "UUID of a captcha image",
	"captchaCode": "123abc",
	"createdTime": "2019-11-20T00:00:00Z"
}
```

### Response

>Status: 201 Created
>
>Location: /api/users

```
{
	"id": 1,
	"username": "Alice",
	"createdTime": "2019-11-20T00:00:00Z"
}
```



## 获得图形验证码

```
GET /api/users/captcha
```

### Response

>Status: 200 OK
>
>Location: api/users/captcha

```
{
   	"captchaId": "UUID of a captcha image",
    "base64": "Base64 encoding of a captcha image"
}
```



## 用户登录

```
POST /api/users/login
```

### Input

| 项目名称    | 类型   | 描述         |
| ----------- | ------ | ------------ |
| username    | string | 用户名       |
| password    | string | 密码         |
| captchaId   | string | 图形验证码id |
| captchaCode | string | 用户识别结果 |

### Example

```
{
	"username": "Alice",
	"password": "123456",
	"captchaId": "UUID of a captcha image",
	"captchaCode": "123abc"
}
```

### Response

> Status: 200 OK
>
> Location: /api/users/login

```
{
	"accessToken": "string",
	"expiresTime": "date"
}
```



## 创建分类列表

```
POST /api/categories
```

### Parameters

| 项目名称     | 类型   | 描述         |
| ------------ | ------ | ------------ |
| categoryName | string | 分类列表名称 |

### Response

>Status: 201 Created
>
>Location: /api/categories

```
{
    "categoryId": 0,
    "categoryName": "Programming"
}
```



## 获取分类列表

```
GET /api/categories
```

### Parameters

| 项目名称 | 类型   | 描述     |
| -------- | ------ | -------- |
| page     | number | 页码     |
| size     | number | 页内项数 |

### Response

>Status: 200 OK
>
>Location: /api/categories

```
{
    "categories": [
        {
        "categoryId": 0,
        "categoryName": "Programming",
        "count": 10
    	},
    	{
        "categoryId": 1,
        "categoryName": "DataStructure",
        "count": 20
    	}
    ]
}
```



## 更新分类列表

```
PUT /api/categories/:categoryId
```

### Parameters

| 项目名称     | 类型   | 描述         |
| ------------ | ------ | ------------ |
| categoryName | string | 分类列表名称 |

### Response

>Status: 200 OK
>
>Location: /api/categories/0



## 删除分类列表

```
DELETE /api/categories/:categoryId
```

### Parameters

| 项目名称     | 类型   | 描述         |
| ------------ | ------ | ------------ |
| categoryName | string | 分类列表名称 |

### Response

>Status: 204 No Content
>
>Location: /api/categories/0



## 创建文章

```
POST /api/articles
```

### Parameters

| 项目名称       | 类型   | 描述        |
| -------------- | ------ | ----------- |
| categoryId     | number | 分类列表 id |
| articleTitle   | string | 文章标题    |
| articleContent | string | 文章内容    |

### Response

>Status: 201 Created
>
>Location: /api/articles/0

```
{
    "articleId": 0,
    "categoryId": 0,
    "articleTitle": "Title",
    "articleContent": "xxxxxx",
    "createdTime": "2019-11-21T00:00:00Z"
}
```



## 获取文章

```
GET /api/articles/:articleId
```

### Parameters

| 项目名称     | 类型   | 描述     |
| ------------ | ------ | -------- |
| articleTitle | string | 文章标题 |

### Response

>Status: 200 OK
>
>Location: /api/articles/0

```
{
    "articleId": 0,
    "articleTitle": "Title",
    "articleContent": "xxxxxx"
}
```



## 更新文章

```
PUT /api/articles/:articleId
```

### Parameters

| 项目名称       | 类型   | 描述        |
| -------------- | ------ | ----------- |
| categoryId     | number | 分类列表 id |
| articleTitle   | string | 文章标题    |
| articleContent | string | 文章内容    |

### Response

>Status: 200 OK
>
>Location: /api/articles/0

```
{
    "articleId": 0,
    "categoryId": 0,
    "articleTitle": "Title",
    "articleContent": "xxxxxx"
}
```



## 删除文章

```
DELETE /api/articles/:articleId
```

### Parameters

| 项目名称     | 类型   | 描述     |
| ------------ | ------ | -------- |
| articleTitle | string | 文章标题 |

### Response

>Status: 204 No Content
>
>Location: /api/articles/0