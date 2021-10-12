package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"log"
	"sync"
	"wxApp/pkg/util"
)

type WsMsg struct {
	MsType  string      `json:"ms_type"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
}

//连接列表
//map[*melody.Session]*models.SysUser
var Clients = sync.Map{}

func Websocket(mr *melody.Melody) gin.HandlerFunc {
	mr.HandleDisconnect(func(session *melody.Session) {
		if u, ok := Clients.Load(session); ok {
			log.Printf("###[%v] client ws disconnect", u)
			Clients.Delete(session)
			BroadCastOnline()
		}
	})
	mr.HandleError(func(session *melody.Session, e error) {
		if u, ok := Clients.Load(session); ok {
			log.Printf("###[%v] client ws error:  err: %v", u, e)
			Clients.Delete(session)
			BroadCastOnline()
		}
	})
	mr.HandleMessage(func(session *melody.Session, msg []byte) {
		var wsMsg = WsMsg{}
		log.Printf("rec: %v", string(msg))
		util.ShowError("", json.Unmarshal(msg, &wsMsg))

	})
	return func(c *gin.Context) {
		util.ShowError("websocket handle err", mr.HandleRequest(c.Writer, c.Request))
		c.Next()
	}
}

func Send(session *melody.Session, m []byte) {
	err := session.Write(m)
	if err != nil {
		log.Println(err)
	}
}

// 新连接或断开后,向所有人广播在线人数
func BroadCastOnline() {
	var onlineNum int
	Clients.Range(func(s, u interface{}) bool {
		onlineNum++
		return true
	})
	msg := WsMsg{
		MsType: "online",
		Data:   onlineNum,
	}
	msgJsons, err := json.Marshal(msg)
	if err != nil {
		log.Printf("broadcast online fail at marshal json.\n ERR:%v", err)
		return
	}
	Clients.Range(func(s, u interface{}) bool {
		err := s.(*melody.Session).Write(msgJsons)
		if err != nil {
			log.Printf("send online msg err\n ERR:%v", err)
		}
		return true
	})
}

