package main

func main() {
	Server().Run()
}

//	r := gin.Default()
//	r.Use(middleware.DummyMiddleWare)
//	r.Use(middleware.ErrorMiddleWare())
//	r.Use(middleware.TokenAuthMiddleWare())
//
//	r.GET("/enigma", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": "pong",
//		})
//	})
//	r.GET("/product", gettingWithQueryParam())
//	r.POST("/product", posting())
//	r.Run("localhost:3000")
//	err := r.Run("localhost:3000")
//	if err != nil {
//		panic(err)
//	}
//	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}
