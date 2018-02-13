package main

import (
    "reservation/src/common"
    "reservation/src/handler"
    "reservation/src/infrastructure/db"
    "flag"
    "os"
    "syscall"

    "github.com/fvbock/endless"
    "github.com/gin-gonic/gin"
)

func main() {

    fg := flag.Int("fg", 0, "run as foreground process")
    addr := flag.String("addr", ":80", "run by port and limited IP address")
    flag.Parse()

    // Init gin
    r := gin.Default()

    // Set routing
    router.POST("/login", func(c *gin.Context) {
        id := c.PostForm("message")
        phone := c.PostForm("message")
        session := c.PostForm("message")

        c.JSON(200, gin.H{
            "id":  id,
            "phone": phone,
            "session":    session,
        })
    })

    r.GET("/student/:sid", handler.Stundet)

    if *fg == 0 {
        // Init endless for graceful restart
        srv := endless.NewServer(*addr, r)
        // Hook of SIGHUP and SIGTERM
        srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGTERM] = append(
            srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGTERM],
            postSigTerm)

        // Execute as a background process
        err := srv.ListenAndServe()
        if err != nil {
            common.Log.Error(err)
        }
    } else {
        r.Run(*addr)
        postSigTerm()
    }

    os.Exit(0)
}

// Call before main()
func init() {
    if err := db.ConnectMySQL("student"); err != nil {
        panic(err.Error())
    }
    if err := common.LogOpen(); err != nil {
        panic(err.Error())
    }
}

// Call when catch signal of SIGTERM
func postSigTerm() {
    common.LogClose()
}