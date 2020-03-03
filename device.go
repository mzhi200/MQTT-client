package main

type eventType uint8
const (
	LoginEv  eventType = 0
	LogoutEv eventType = 1
	DataEv   eventType = 2
)

type event struct {
	msg interface{}
    eventID eventType
}

func (ev *event) eventHandle(dv *deviceData) (err error) {
	log := newLog(nil)
	procFlag := false
	defer func() {
		dv.chProcDone <- procComplete{procCpltFlag: procFlag}
	}()

	oneNetMsg := date{Id: dv.deviceId}
	switch ev.eventID {
	case LoginEv:
		oneNetMsg.Dp.MsgType = []MessageType{{V: LOGIN,}}
		err = ue.db.Set(GetDeviceDbKey(dv.deviceId), "login", 0).Err()
		if err != nil {
			log.Error("Failed to save device_%d to db; Err: %v", dv.deviceId, err)
			return
		}

	case LogoutEv:
		oneNetMsg.Dp.MsgType = []MessageType{{V: LOGOUT,}}
		err = ue.db.Del(GetDeviceDbKey(dv.deviceId)).Err()
		if err != nil {
			log.Error("Failed to delele device_%d from db; Err: %v", dv.deviceId, err)
			return
		}

	case DataEv:
		dvData := ev.msg.(DeviceData)
		oneNetMsg.Dp.MsgType = []MessageType{{V: DP,}}
		oneNetMsg.Dp.Tp = []TpData{{V: dvData.Tp}}
		oneNetMsg.Dp.Power = []PowerData{{V: dvData.Power}}
		_, err1 := ue.db.Get(GetDeviceDbKey(dv.deviceId)).Result()
		if err1 != nil {
			log.Error("Failed to get device_%d from db; Err: %v", dv.deviceId, err1)
			err = err1
			return
		}
	}
	err = ue.publishData(oneNetMsg)
	if err != nil {
		log.Error("Failed to Publish; Err: %v", err)
		return
	}
	procFlag = true
	return
}

type procComplete struct {
	procCpltFlag bool
	//msg interface{}
	//eventID eventType
}

type deviceData struct {
	deviceId    uint32
	chEvent     chan event
	chProcDone  chan procComplete
}


func (ue *deviceData) deviceEventHandleLoop() {
	log := newLog(nil)

	for {
		select {
		case ev := <- ue.chEvent:
			err := ev.eventHandle(ue)
			if err != nil {
				log.Error("Failed to eventHandle; Err: %v", err)
			}

		case pr := <- ue.chProcDone:
			if pr.procCpltFlag {
				//save to db
			} else {
				//error
			}
			log.Debug("deviceEventHandleLoop Done.")
			return
		}
	}
}

func newDeviceCb() *deviceData {
	dv := &deviceData{
		deviceId: 0,
		chEvent: make(chan event, 1),
		chProcDone: make(chan procComplete, 1),
	}

	go dv.deviceEventHandleLoop()
	return dv
}