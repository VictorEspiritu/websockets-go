// Package main provides ...
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "github.com/gorilla/websocket"
   "net/http"
)

func main() {
   f := "index.html"
   r := gin.Default()
   r.LoadHTMLFiles(f)

   r.GET("/", func (c *gin.Context) {
      c.HTML(200, f, nil)
   })

   r.GET("/ws", func (c *gin.Context) {
      wshandler(c.Writer, c.Request)
   })

   r.Run("localhost:2020")
}

func wshandler(w http.ResponseWriter, r *http.Request) {
   conn, err := wsupgrader.Upgrade(w, r, nil)
   if err != nil{
      fmt.Println("Failed to upgrade websocket: %+v", err)
      return
   }

   for{
      t, msg, err := conn.ReadMessage()
      if err != nil{
         fmt.Println("Failed to Read websocket: %+v", err)
         break
      }
      conn.WriteMessage(t, msg)
   }
}

var wsupgrader = websocket.Upgrader{
   ReadBufferSize : 1024,
   WriteBufferSize : 1024,
}

