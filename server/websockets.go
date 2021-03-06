package main

import (
	"code.google.com/p/go.net/websocket"
	//"code.google.com/p/google-api-go-client/mirror/v1"
	"github.com/wearscript/wearscript-go/wearscript"
	"fmt"
	//"strings"
	"time"
	"sync"
)

var Locks = map[string]*sync.Mutex{} // [user]
var UserSockets = map[string][]chan *[]interface{}{} // [user]
var Managers = map[string]*wearscript.ConnectionManager{}

func CurTime() float64 {
	return float64(time.Now().UnixNano()) / 1000000000.
}

func WSHandler(ws *websocket.Conn) {
	defer ws.Close()
	fmt.Println("Connected with glass")
	fmt.Println(ws.Request().URL.Path)
	userId, err := userID(ws.Request())
    //    if err != nil || userId == "" {
    //            path := strings.Split(ws.Request().URL.Path, "/")
	//	fmt.Println(path)
    //            if len(path) != 3 {
    //                    fmt.Println("Bad path")
    //                    return
    //            }
    //            userId, err = getSecretUser("ws", secretHash(path[len(path)-1]))
    //            if err != nil {
    //                    fmt.Println(err)
    //                    return
    //            }
    //    }
	//userId = SanitizeUserId(userId)
	//_, err = mirror.New(authTransport(userId).Client()) // svc
	//if err != nil {
	//	LogPrintf("ws: mirror")
	//	//WSSend(c, &[]interface{}{[]uint8("error"), []uint8("Unable to create mirror transport")})
	//	return
	//}
	// TODO(brandyn): Add lock here when adding users
	if Managers[userId] == nil {
		Managers[userId], err = wearscript.ConnectionManagerFactory("server", "demo")
		if err != nil {
			LogPrintf("ws: cm")
			return
		}
		Managers[userId].SubscribeTestHandler()
		//Managers[userId].Subscribe("gist", func(channel string, dataRaw []byte, data []interface{}) {
		//	GithubGistHandle(Managers[userId], userId, data)
		//})
		//Managers[userId].Subscribe("weariverse", func(channel string, dataRaw []byte, data []interface{}) {
		//	WeariverseGistHandle(Managers[userId], userId, data)
		//})
        //Managers[userId].Subscribe("ping", func (c string, dataBin []byte, data []interface{}) {
        //    resultChan, ok := data[1].(string)
        //    if !ok {return}
        //    Managers[userId].Publish(resultChan, data[2], time.Now().UnixNano() / 1000000000., Managers[userId].GroupDevice())
        //    //fmt.PrintLn(fmt.Sprintf(resultChan));
        //})
        //Managers[userId].Subscribe("sensors", func (c string, dataBin []byte, data []interface{}) {
        //    //names := data[1]
        //    //samples := data[2]
        //    //fmt.Println(fmt.Sprintf("Sensors[%s] Names[%v] Samples[%v]", c, names, samples))
        //})
	}
	cm := Managers[userId]
	conn, err := cm.NewConnection(ws) // con
	if err != nil {
		fmt.Println("Failed to create conn")
		return
	}
	cm.HandlerLoop(conn)
}
