package main
import(
	"encoding/json"
	"fmt"
)

var (
	help = [...]string{
		"r e:   Connect mqtt server",
		"r d:   Disconnect mqtt server",
		"r s:   Subscribe event",
		"r u:   Unsubscribe event",
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
			case "e":
				//Connect mqtt server
				ue.client, err = oneNetConnect(config)
				if err != nil {
					log.Error("Failed to Connect; Err: %v", err)
					return
				}

			case "s":
				//Subscribe event
				err = ue.subscribeEven("#")
				if err != nil {
					log.Error("Failed to Subscribe; Err: %v", err)
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
					},
				}
				textString, err := json.Marshal(text)
				if err != nil {
					log.Error("Failed to Marshal(text); Err: %v",err)
					return
				}
				ue.publishData("dp/post/json", textString)
				if err != nil {
					log.Error("Failed to Publish; Err: %v", err)
					return
				}

			case "u":
				//Unsubscribe event
				err = ue.unSubscribeEven("#")
				if err != nil {
					log.Error("Failed to UnSubscribe; Err: %v", err)
					return
				}

			case "d":
				//Disconnect mqtt server
				ue.client.Disconnect(250)

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