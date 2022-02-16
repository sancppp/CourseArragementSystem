package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Session(secret string) gin.HandlerFunc {
	//store 存储引擎
	store := cookie.NewStore([]byte(secret))
	//Also set Secure: true if using SSL, you should though

	//store的参数设置
	//store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions("camp-session", store)
}

// var cookieName = "my_gin"
// var lifeTime = 3600
// func Session(c *context.Context) {

// 	cookie, err := c.Cookie(cookieName)

// 	if err == nil {

// 		sessionString, err := redis.Client().Get(context2.TODO(), cookie).Result()

// 		if err == nil {

// 			var session context.Session

// 			json.Unmarshal([]byte(sessionString), &session)

// 			//存储到context中，方便当前请求中的其他函数好操作session
// 			c.Set("_session", session)

// 			return
// 		}

// 	}

// 	sessionKey := uuid.NewV4().String()

// 	c.SetCookie(cookieName, sessionKey, lifeTime, "/", c.Domain(), false, true)

// 	session := context.Session{
// 		Cookie:      sessionKey,
// 		ExpireTime:  time.Now().Unix() + int64(lifeTime),
// 		SessionList: make(map[string]interface{}),
// 	}

// 	//这里也要
// 	c.Set("_session", session)

// 	jsonString, _ := json.Marshal(session)

// 	redis.Client().Set(context2.TODO(), sessionKey, jsonString, time.Second*time.Duration(lifeTime))

// }
