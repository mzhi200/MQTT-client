package main
import(
	"fmt"
)

var (
	help = [...]string{
		"r c:   Connect mqtt server",
		"r d:   Disconnect mqtt server",
		"r s:   Subscribe event",
		"r u:   Unsubscribe event",
		"r l:   Publish Login",
		"r o:   Publish Logout",
		"r p:   Publish Data",
	}
)
func triggerCommandFun() {
	var err error
	cmd := [5]string{}
	log := newLog(nil)
	log.Info("Start CMD Test...")

	for {
		fmt.Printf("diag> ")
		//clear cmd
		for k:=0; k<5; k++ {cmd[k] = ""}
		num, _ := fmt.Scanln(&cmd[0], &cmd[1], &cmd[2], &cmd[3], &cmd[4])
		num = num
		//fmt.Printf("cmd num: %d cmd: %+v\n", num, cmd)
		if cmd[0] == "r" {
			switch cmd[1] {
			case "c":
				//Connect mqtt server
				ue.client, err = oneNetConnect(config)
				if err != nil {
					log.Error("Failed to Connect; Err: %v", err)
					return
				}

			case "d":
				//Disconnect mqtt server
				ue.client.Disconnect(250)

			case "s":
				//Subscribe event
				err = ue.subscribeEven(TopicSubAllEv)
				if err != nil {
					log.Error("Failed to Subscribe; Err: %v", err)
					return
				}

			case "u":
				//Unsubscribe event
				err = ue.unSubscribeEven(TopicSubAllEv)
				if err != nil {
					log.Error("Failed to UnSubscribe; Err: %v", err)
					return
				}

			case "l":
				//Publish Login
				login := date{
					Id: 123,
					Dp: DpType{
						MsgType: []MessageType{
							{
								V: LOGIN,
							},
						},
					},
				}
				ue.publishData(login)
				if err != nil {
					log.Error("Failed to Publish; Err: %v", err)
					return
				}

			case "o":
				//Publish Login
				logout := date{
					Id: 123,
					Dp: DpType{
						MsgType: []MessageType{
							{
								V: LOGOUT,
							},
						},
					},
				}
				ue.publishData(logout)
				if err != nil {
					log.Error("Failed to Publish; Err: %v", err)
					return
				}

			case "p":
				//Publish Data
				text := date{
					Id: 123,
					Dp: DpType{
						Tp: []TpData{
							{
								V: 30,
							},
						},
						Power: []PowerData{
							{
								V: 4.5,
							},
						},
						MsgType: []MessageType{
							{
								V: DP,
							},
						},
					},
				}
				ue.publishData(text)
				if err != nil {
					log.Error("Failed to Publish; Err: %v", err)
					return
				}

			default:
				fmt.Println("CMD List:")
				for _,cmdItem := range help {
					fmt.Println("  ", cmdItem)
				}
			}
		} else if cmd[0] == "?" || cmd[0] == "help"{
			fmt.Println("CMD List:")
			for _,cmdItem := range help {
				fmt.Println("  ", cmdItem)
			}
		}
	}
}