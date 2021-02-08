# Go开源框架

## gin

- [gin日志打印到某个目录，并实现拆分](https://blog.csdn.net/u010918487/article/details/86146691)

- 自定义中间件

  ```go
  func Logger() gin.HandlerFunc {
  	return func(c *gin.Context) {
  		t := time.Now()
  
  		// Set example variable
  		c.Set("example", "12345")
  
  		// before request
  
  		c.Next()
  
  		// after request
  		latency := time.Since(t)
  		log.Print(latency)
  
  		// access the status we are sending
  		status := c.Writer.Status()
  		log.Println(status)
  	}
  }
  
  func main() {
  	r := gin.New()
  	r.Use(Logger())
  
  	r.GET("/test", func(c *gin.Context) {
  		example := c.MustGet("example").(string)
  
  		// it would print: "12345"
  		log.Println(example)
  	})
  
  	// Listen and serve on 0.0.0.0:8080
  	r.Run(":8080")
  }
  ```

  

## go-micro

- [Go Micro 中文文档 3.x](https://learnku.com/docs/go-micro/3.x)

- [Go Micro入门](http://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/GoMicro%E5%85%A5%E9%97%A8.html)

