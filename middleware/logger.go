package  middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"math"

	"os"
	"time"

)

func Log()gin.HandlerFunc{
	filepath:= "log/ginblog"

	src,err := os.OpenFile(filepath,os.O_RDWR|os.O_CREATE,0755)
	if err!=nil{
		log.Fatal(err)
	}

	logger :=logrus.New()
	logger.Out=src
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		endTime := time.Since(startTime)
        spendTime := fmt.Sprintf("%d ms",int(math.Ceil(float64(endTime.Nanoseconds())/1000000.0)))
        hostName ,_ :=os.Hostname()
        statusCode := c.Writer.Status()
        clientIP := c.ClientIP()
        userAgent :=c.Request.UserAgent()
        dataSize := c.Writer.Size()
        if dataSize< 0{
        	dataSize=0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":hostName,
			"Status":statusCode,
			"SpendTime":spendTime,
			"IP":clientIP,
			"Method":method,
			"Agent":userAgent,
			"Path":path,
		})

		if len(c.Errors)>0{
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}



}