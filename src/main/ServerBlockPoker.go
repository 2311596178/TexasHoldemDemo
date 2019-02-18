package main

import (
	"bufio"
	"bytes"
	"commonMethod"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var RoutineNum int
var beginUserID int
var GameType int
var RobotTableID int


var SendByteMap = [...]byte {
	0x0D,0x36,0x9B,0x0B,0xD4,0xC4,0x39,0x74,0x45,0x23,0x16,0x14,0x06,0xEB,0x04,0x3E,
	0x12,0x5C,0x8B,0xBC,0x61,0x63,0xF6,0xA5,0xE1,0x65,0xD8,0xF5,0x5A,0x07,0xF0,0x13,
	0xD7,0x2F,0x40,0x5F,0x44,0x8E,0x6E,0xBF,0x7E,0xAB,0x2C,0x1F,0xB4,0xAC,0x9D,0x91,
	0xF2,0x20,0x6B,0x4A,0x24,0x59,0x89,0x64,0x70,0xE0,0x6A,0x5E,0x3D,0x0A,0x77,0x42,
	0x80,0x27,0xB8,0xC5,0x8C,0x0E,0xFA,0x8A,0xD5,0x29,0x56,0x57,0x6C,0x53,0x67,0x41,
	0xE8,0x00,0x1A,0xE4,0x86,0x83,0xB0,0x22,0x28,0x4D,0x3F,0x26,0x46,0x4F,0x6F,0x2B,
	0x72,0x3A,0xF1,0x8D,0x97,0x95,0x49,0x84,0xE5,0xE3,0x79,0x8F,0x51,0x17,0xA8,0x82,
	0x93,0xAF,0x69,0x0C,0x71,0x31,0xDE,0x21,0x75,0xA0,0xAA,0xBA,0x7C,0x38,0x02,0xB7,
	0x87,0xF8,0x15,0x05,0x3C,0xD3,0xA4,0x85,0x2E,0xFB,0xEE,0x47,0x3B,0xEF,0x37,0x7F,
	0xC7,0xAE,0x96,0x35,0xD0,0xBB,0xD2,0xC8,0xA2,0x08,0xF3,0xD1,0x73,0xF4,0x48,0x2D,
	0x81,0x01,0xFD,0xE7,0x1D,0xCC,0xCD,0xBD,0x1B,0x7A,0x2A,0xAD,0x66,0xBE,0x55,0x33,
	0xC6,0xDD,0xFF,0xFC,0xCE,0xCF,0xB3,0x09,0x5D,0xEA,0x9C,0x34,0xF9,0x10,0x9F,0xDA,
	0x03,0xDB,0x88,0xB2,0x1E,0x4E,0xB9,0xE6,0xC2,0xF7,0xCB,0x7D,0xC9,0x62,0xC3,0xA6,
	0xDC,0xA7,0x50,0xB5,0x4B,0x94,0xC0,0xED,0x4C,0x11,0x5B,0x78,0xD9,0xB1,0x92,0x19,
	0xE9,0xA1,0x1C,0xB6,0x32,0x99,0xA3,0x76,0x9E,0x7B,0x6D,0x9A,0x30,0xD6,0xA9,0x25,
	0x90,0xCA,0xE2,0x58,0xC1,0x18,0x52,0xFE,0xDF,0x68,0x98,0x54,0xEC,0x60,0x43,0x0F}

var RecvByteMap = [...]byte {
	0x51,0xa1,0x7e,0xc0,0x0e,0x83,0x0c,0x1d,0x99,0xb7,0x3d,0x03,0x73,0x00,0x45,0xff,
	0xbd,0xd9,0x10,0x1f,0x0b,0x82,0x0a,0x6d,0xf5,0xdf,0x52,0xa8,0xe2,0xa4,0xc4,0x2b,
	0x31,0x77,0x57,0x09,0x34,0xef,0x5b,0x41,0x58,0x49,0xaa,0x5f,0x2a,0x9f,0x88,0x21,
	0xec,0x75,0xe4,0xaf,0xbb,0x93,0x01,0x8e,0x7d,0x06,0x61,0x8c,0x84,0x3c,0x0f,0x5a,
	0x22,0x4f,0x3f,0xfe,0x24,0x08,0x5c,0x8b,0x9e,0x66,0x33,0xd4,0xd8,0x59,0xc5,0x5d,
	0xd2,0x6c,0xf6,0x4d,0xfb,0xae,0x4a,0x4b,0xf3,0x35,0x1c,0xda,0x11,0xb8,0x3b,0x23,
	0xfd,0x14,0xcd,0x15,0x37,0x19,0xac,0x4e,0xf9,0x72,0x3a,0x32,0x4c,0xea,0x26,0x5e,
	0x38,0x74,0x60,0x9c,0x07,0x78,0xe7,0x3e,0xdb,0x6a,0xa9,0xe9,0x7c,0xcb,0x28,0x8f,
	0x40,0xa0,0x6f,0x55,0x67,0x87,0x54,0x80,0xc2,0x36,0x47,0x12,0x44,0x63,0x25,0x6b,
	0xf0,0x2f,0xde,0x70,0xd5,0x65,0x92,0x64,0xfa,0xe5,0xeb,0x02,0xba,0x2e,0xe8,0xbe,
	0x79,0xe1,0x98,0xe6,0x86,0x17,0xcf,0xd1,0x6e,0xee,0x7a,0x29,0x2d,0xab,0x91,0x71,
	0x56,0xdd,0xc3,0xb6,0x2c,0xd3,0xe3,0x7f,0x42,0xc6,0x7b,0x95,0x13,0xa7,0xad,0x27,
	0xd6,0xf4,0xc8,0xce,0x05,0x43,0xb0,0x90,0x97,0xcc,0xf1,0xca,0xa5,0xa6,0xb4,0xb5,
	0x94,0x9b,0x96,0x85,0x04,0x48,0xed,0x20,0x1a,0xdc,0xbf,0xc1,0xd0,0xb1,0x76,0xf8,
	0x39,0x18,0xf2,0x69,0x53,0x68,0xc7,0xa3,0x50,0xe0,0xb9,0x0d,0xfc,0xd7,0x8a,0x8d,
	0x1e,0x62,0x30,0x9a,0x9d,0x1b,0x16,0xc9,0x81,0xbc,0x46,0x89,0xb3,0xa2,0xf7,0xb2}

var cardMap=map[int]byte{
	0:0x1,1:0x2,2:0x3,3:0x4,4:0x5,5:0x6,6:0x7,7:0x8,8:0x9,9:0xa,10:0xb,11:0xc,12:0xd,                 //方块 A - K
	13:0x11,14:0x12,15:0x13,16:0x14,17:0x15,18:0x16,19:0x17,20:0x18,21:0x19,22:0x1a,23:0x1b,24:0x1c,25:0x1d, //梅花 A - K
	26:0x21,27:0x22,28:0x23,29:0x24,30:0x25,31:0x26,32:0x27,33:0x28,34:0x29,35:0x2a,36:0x2b,37:0x2c,38:0x2d, //红桃 A - K
	39:0x31,40:0x32,41:0x33,42:0x34,43:0x35,44:0x36,45:0x37,46:0x38,47:0x39,48:0x3a,49:0x3b,50:0x3c,51:0x3d} //黑桃 A - K

const  (
	enAddScoreNULL=0				//空值状态
	enAddScoreGIVEUP=1			//放弃状态
	enAddScoreCHECK=2			//让牌状态
	enAddScoreFOLLOW=3			//跟注状态
	enAddScoreADD=4				//加注状态
	enAddScoreSHOWHAND=5			//梭哈状态
)

/* 客户端封装 */
type Client struct {
	i		int
	userID  int32
	login	bool
	chairID	uint16
	tableID	uint16
	tableCode int32
	score     int
	conn	net.Conn
	logonChan   chan bool
	stopChan	chan bool
	sendChan	chan []byte
	gameType	int
}

type Table struct {
	tableCode    int32
	chairSlice   []int32 //chairSlice[nChairID]=nUserID
	addUserChan  chan AddUser
	maxUser      int
	btCenterCard []byte                  //公牌
	userInfo     []UserInfo
	tableInfo    TableInfo
	closeEnter   bool
}
type AddUser struct {
	userID      int32
	userScore   int
}
type UserInfo struct {
	score        int
	handCard     []byte           //手牌
	maxScore     int             //用户最大筹码
	betScore     int             //累计下注筹码
	tableScore       int             //桌上下注筹码
	addScoreStatus       byte            //0x0正常 0x01梭哈 0x02弃牌
	tableStatus      bool             //true 正常 false淘汰
}
type TableInfo struct {
	currentUser      int               //当前用户
	dealer		     int               //D庄家
	chopCount        int               //摇摆次数
	balanceCount     int               //平衡次数
	lMinRaiseScore   int               //加最小注
	balanceScore     int               //下注平衡值
	lPotScore        int               //底池
}

func init() {
	conf := new(Config)
	conf.InitConfig("./ClicentBlockPoker.ini")
	RoutineNum, _ = strconv.Atoi(conf.Read("Base", "cliNum"))
	beginUserID, _ = strconv.Atoi(conf.Read("Base", "beginUID"))
	GameType, _ = strconv.Atoi(conf.Read("Base", "gameType"))
	RobotTableID=0
}

/*常用方法 */
func putOutNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05  ")
}

/* 创建牌桌 */
func (c *Client) createTable() {
	nTableCode := commonMethod.CreateRandNumber()
	tableItem :=&Table{
		tableCode:    nTableCode,
		addUserChan:  make(chan AddUser),
		maxUser:      RoutineNum,
		userInfo:     make([]UserInfo,RoutineNum),
		closeEnter:   false,
	}

	for i := 0; i< len(tableItem.userInfo);i++  {
		p:=tableItem.userInfo[i]
		p.handCard=make([]byte,2)
		p.score=0
		p.maxScore=0
		p.betScore=0
		p.tableScore=0
		p.addScoreStatus=0
		p.tableStatus=false
	}
	tableItem.tableInfo=TableInfo{
		currentUser:0,
		dealer:0,
		chopCount:0,
		balanceScore:0,
		lMinRaiseScore:50,
		balanceCount:0,
		lPotScore:0,
	}
	go tableItem.enterTable()
	mapTable[nTableCode] = tableItem
	createTable := bytes.NewBuffer([]byte{})
	binary.Write(createTable, binary.LittleEndian, int32(nTableCode))
	c.sendData(199, 50, createTable.Bytes())
	fmt.Println("创建牌桌成功",nTableCode)
}

/* 进入牌桌 */
func (c *Client) enterTable(nTableCode int32) {
	tableItem := mapTable[nTableCode]
	if  tableItem!= nil{
		if tableItem.closeEnter{
			return
		}
		addUser:=AddUser{
			userID:     c.userID,
			userScore:  c.score,
		}
		tableItem.addUserChan <- addUser
		c.score=0
	}
}

/* 下注 */
func (c *Client) addScore() {
	/* 随机数秒延迟操作,防止机器人动作太快 */
	time.Sleep(time.Duration(rand.Intn(8) + 1) * time.Second)
	addScore := bytes.NewBuffer([]byte{})

	binary.Write(addScore, binary.LittleEndian, c.tableCode)
	binary.Write(addScore, binary.LittleEndian, int32(c.i+beginUserID))
	binary.Write(addScore, binary.LittleEndian, int32(8))
	binary.Write(addScore, binary.LittleEndian, int32(8))

	if GameType == 2 {
		nScroe := rand.Intn(20) + 1
		binary.Write(addScore, binary.LittleEndian, int64(nScroe))
	} else {
		binary.Write(addScore, binary.LittleEndian, int64(1))
	}
	binary.Write(addScore, binary.LittleEndian, false)
	c.sendData(200, 2, addScore.Bytes())

}
func (c *Client) TexasAddScore()  {
	time.Sleep(time.Duration(1) * time.Second)
	addScore := bytes.NewBuffer([]byte{})

	binary.Write(addScore, binary.LittleEndian, int32(RobotTableID))
	binary.Write(addScore, binary.LittleEndian, int32(c.i+beginUserID))
	binary.Write(addScore, binary.LittleEndian, int32(7))
	binary.Write(addScore, binary.LittleEndian, int32(GameType))

	binary.Write(addScore, binary.LittleEndian, int32(0))

	opetateType := make([]byte, 1)
	cbOpetateType := []byte(strconv.Itoa(2))
	copy(opetateType, cbOpetateType)
	addScore.Write(opetateType)
	fmt.Println(putOutNowTime(),c.i+beginUserID," 下注 ","tableCode:",c.tableCode)
	c.sendData(200, 1, addScore.Bytes())
}

/* 开牌 */
func (c *Client) openCard() {
	/* 随机数秒延迟操作,防止机器人动作太快 */
	time.Sleep(time.Duration(rand.Intn(6) + 1) * time.Second)
	openCard := bytes.NewBuffer([]byte{})

	binary.Write(openCard, binary.LittleEndian, c.tableCode)
	binary.Write(openCard, binary.LittleEndian, int32(c.i+beginUserID))
	binary.Write(openCard, binary.LittleEndian, int32(8))
	binary.Write(openCard, binary.LittleEndian, int32(8))

	binary.Write(openCard, binary.LittleEndian, true)
	c.sendData(200, 3, openCard.Bytes())
}

/* 登陆 */
func (c *Client) logon() {
	var logon []byte

	userID := bytes.NewBuffer([]byte{})
	binary.Write(userID, binary.LittleEndian, int32(c.i+beginUserID))

	logon = append(logon, userID.Bytes()...)
	logon = append(logon, 1)

	faceID := bytes.NewBuffer([]byte{})
	binary.Write(faceID, binary.LittleEndian, int32(c.i))
	logon = append(logon, faceID.Bytes()...)

	userName := make([]byte, 64)
	logon = append(logon, userName...)

	passWord := make([]byte, 33)
	for i:=0; i<6; i++ {
		passWord[i] = '1'
	}
	logon = append(logon, passWord...)

	userGuid := make([]byte, 64)
	logon = append(logon, userGuid...)

	cliAddr := bytes.NewBuffer([]byte{})
	binary.Write(cliAddr, binary.LittleEndian, int32(c.i))
	logon = append(logon, cliAddr.Bytes()...)

	c.sendData(101, 1, logon)
}



func (c *Client) sendData(nMainID int16, nSubID int16, szData []byte) {
	var sendData = make([]byte, 8)

	mainID := bytes.NewBuffer([]byte{})
	subID := bytes.NewBuffer([]byte{})

	binary.Write(mainID, binary.LittleEndian, nMainID)
	binary.Write(subID, binary.LittleEndian, nSubID)

	copy(sendData[4:], mainID.Bytes())
	copy(sendData[6:], subID.Bytes())

	sendData = append(sendData, szData...)

	/* 加密 */
	encryptData(sendData, 8 + len(szData))

	/* 放入管道发送 */
	c.sendChan <- sendData
}

func (c *Client) dealSend() {
	defer c.conn.Close()

	for {
		select {
		case data := <- c.sendChan:
			_, err := c.conn.Write(data)
			if err != nil {
				fmt.Printf("Client %d Write data to Server Fail!\n", c.i)

				delete(mapClis, c.i)

				c.stopChan <- true
				c.stopChan <- true
			}
		case <- c.stopChan:
			goto SendEnd
		}
	}
SendEnd:
	return
}

func (c *Client) dealRecv() {
	defer c.conn.Close()

	recvData := make([]byte, 1024)
	var nLen uint16
	var szLen  = make([]byte, 2)

	for {
		select {
		case <- c.stopChan:
			goto RecvEnd
		default:
			break
		}

		msgLen, err := c.conn.Read(recvData[:4])
		if err != nil || msgLen <= 0 {
			fmt.Println(c.i+beginUserID,"Read from Socket Fail!", err)
			delete(mapClis, c.i)
			c.stopChan <- true
			c.stopChan <- true
			continue
		}

		copy(szLen, recvData[2:])
		lenBuf := bytes.NewBuffer(szLen)

		binary.Read(lenBuf, binary.LittleEndian, &nLen)

		if nLen >= 1024 {
			fmt.Println("recv illeg MsgLen=", nLen)
			continue
		}

		msgLen, err = c.conn.Read(recvData[4:nLen])
		if err != nil || msgLen <= 0 {
			fmt.Println(c.i+beginUserID,"Read from Socket Fail!", err)
			continue
		}

		conChan := make(chan bool)
		/* 收到数据后，启动一个匿名协程处理 */
		go func() {
			data := make([]byte, 1024)
			sendMsg := bytes.NewBuffer([]byte{})
			copy(data, recvData)
			n := msgLen + 4
			conChan <- true

			/* 定义主消息切片 */
			var msgID = make([]byte, 2)
			var nMainID, nSubID uint16
			/* 解密 */
			crevasseData(data, n)
			/* 主消息 */
			copy(msgID, data[4:])
			pTemp := bytes.NewBuffer(msgID)
			binary.Read(pTemp, binary.LittleEndian, &nMainID)
			/* 从消息 */
			copy(msgID, data[6:])
			pTemp = bytes.NewBuffer(msgID)
			binary.Read(pTemp, binary.LittleEndian, &nSubID)
			switch nMainID {
			/* 登陆 */
			case 101:
				if nSubID == 1 {
					c.login = true
					var nUserID int32
					userID := make([]byte, 4)
					copy(userID, data[8:])
					codeBuf := bytes.NewBuffer(userID)
					binary.Read(codeBuf, binary.LittleEndian, &nUserID)
					fmt.Println(nUserID,"logon in")
					c.userID = nUserID
					mapClis[int(nUserID)]=c
					binary.Write(sendMsg, binary.LittleEndian, int32(nUserID))
					c.sendData(101,1,sendMsg.Bytes())
				}
			/*牌桌消息*/
			case 199:
				switch nSubID {
				case 1:
					fmt.Println(c.userID,"创建牌桌")
					c.createTable()
				case 4:
					/* 加入牌桌 */
					//var nTableCode int32
					//tableCode := make([]byte, 4)
					//copy(tableCode, data[8:])
					//codeBuf := bytes.NewBuffer(tableCode)
					//binary.Read(codeBuf, binary.LittleEndian, &nTableCode)
					nTableCode := commonMethod.ReadInt32FromData(data,12)
					fmt.Println(c.userID,"加入牌桌",nTableCode)
					c.enterTable(nTableCode)


				default:
					break
				}
			/*游戏消息*/
			case 200:
				//fmt.Println(200,"SubID:",nSubID)
				switch nSubID {
				case 1:
					nTableCode:=commonMethod.ReadInt32FromData(data,8)
					nUserID:=commonMethod.ReadInt32FromData(data,12)
					t:=mapTable[nTableCode]
					if c.userID != nUserID || t==nil || t.chairSlice[t.tableInfo.currentUser]!=nUserID{
						break
					}
					addScore:=commonMethod.ReadInt32FromData(data,16)
					addScoreStatus:=commonMethod.ReadByteFromData(data,20)
					if t.userInfo[t.tableInfo.currentUser].score<=0 && t.userInfo[t.tableInfo.currentUser].addScoreStatus!=enAddScoreSHOWHAND{
						t.userInfo[t.tableInfo.currentUser].tableStatus=false
					}
					t.userInfo[t.tableInfo.currentUser].addScoreStatus=addScoreStatus
					fmt.Println(200,1)
					t.addUserScore(t.tableInfo.currentUser,int(addScore),addScoreStatus)


				default:
					break
				}

			case 0:
				if nSubID == 1 {
					c.sendData(0, 3, []byte{})
				}
			default:
				break
			}
		}()

		<- conChan
	}

RecvEnd:
	return
}

func (t *Table)enterTable()  {
	for{
		addUser:=<-t.addUserChan
		t.chairSlice=append(t.chairSlice,addUser.userID)
		t.userInfo[len(t.chairSlice)-1].score=addUser.userScore
		t.userInfo[len(t.chairSlice)-1].maxScore=addUser.userScore
		t.userInfo[len(t.chairSlice)-1].tableStatus=true
		if len(t.chairSlice)>=9 {
			t.closeEnter=true
			t.gameStart()
			return
		}
	}
}

func (t *Table) gameStart()  {
	fmt.Println(t.chairSlice)
	t.getCardArr()
	fmt.Println("游戏开始 发牌")
	t.sendHandCard()

	t.tableInfo.lPotScore = 0
	for i := 0; i< len(t.userInfo);i++  {
		//p:=&t.userInfo[i]
		if t.userInfo[i].tableStatus == true && t.userInfo[i].score>0{
			if t.userInfo[i].score<t.tableInfo.lMinRaiseScore{
				score:=t.userInfo[i].score
				t.userInfo[i].score=0
				t.userInfo[i].betScore += score
				t.userInfo[i].tableScore = score
				t.userInfo[i].addScoreStatus = enAddScoreCHECK
				t.tableInfo.lPotScore += score
				continue
			}
			t.userInfo[i].score -= t.tableInfo.lMinRaiseScore
			t.userInfo[i].betScore += t.tableInfo.lMinRaiseScore
			t.userInfo[i].tableScore = t.tableInfo.lMinRaiseScore
			t.userInfo[i].addScoreStatus = enAddScoreCHECK
			t.tableInfo.lPotScore += t.tableInfo.lMinRaiseScore
			//p.tableStatus=true
		}
	}
	t.tableInfo.balanceScore = t.tableInfo.lMinRaiseScore

	nextUser := t.tableInfo.dealer
	isEnd:=true
	for i := 0; i < t.maxUser; i++ {
		nextUser=(nextUser+i+1)%(t.maxUser-1)
		fmt.Println("nextUser",nextUser,t.chairSlice[nextUser],t.userInfo[nextUser].tableStatus)
		if t.userInfo[nextUser].tableStatus == true {
			isEnd=false
			break
		}
	}
	if !isEnd {
		t.tableInfo.dealer=nextUser
		t.tableInfo.currentUser=nextUser
		fmt.Println("游戏开始 玩家列表:",t.chairSlice)
		fmt.Println("桌上筹码",t.tableInfo.lPotScore)
		fmt.Println(t.userInfo)
		t.addUserScore(t.tableInfo.currentUser,t.tableInfo.lMinRaiseScore,enAddScoreCHECK)
	}else{
		fmt.Println("game end error")
	}

}
func (t *Table)gameEnd() {
	//commonMethod.SortCard(t.)
	fmt.Println("游戏结束",t.userInfo)
	fmt.Println("tableInfo",t.tableInfo)
	fmt.Println("btCenterCard",commonMethod.GetAllCardStr(t.btCenterCard),t.btCenterCard)
	cardKind:=make([]byte,RoutineNum)
	//tableScore:=0
	bEndCard :=make(map[int][]byte)
	for i := 0; i < RoutineNum; i++ {
		if t.userInfo[i].tableStatus==false || t.userInfo[i].addScoreStatus==enAddScoreGIVEUP {
			cardKind[i]=0
			bTempCard:=append(t.userInfo[i].handCard,t.btCenterCard...)
			commonMethod.SortCard(bTempCard)
			bEndCard[i]=make([]byte,5)
			copy(bEndCard[i],bTempCard)
		}else{
			cardKind[i],bEndCard[i]=commonMethod.GetGreatestCardType(t.userInfo[i].handCard,t.btCenterCard)
		}
		fmt.Println(commonMethod.GetAllCardStr(t.userInfo[i].handCard),t.userInfo[i].handCard)
		fmt.Println(t.chairSlice[i],t.userInfo[i].score,commonMethod.GetAllCardStr(bEndCard[i]),bEndCard[i],cardKind[i])
		//tableScore+=t.userInfo[i].tableScore
		t.userInfo[i].tableScore=0
	}
	//fmt.Println(tableScore,t.tableInfo.lPotScore)
	fmt.Println(bEndCard)

	var winKind byte=0
	var winUser []int

	for i:=0; i< len(cardKind); i++{
		if cardKind[i]>winKind {
			winKind=cardKind[i]
		}
	}
	for i:=0; i< len(cardKind); i++{
		if cardKind[i]==winKind {
			winUser=append(winUser,i)
		}
	}
	fmt.Println("winKind",winKind,"winUser",winUser)
	var lastWinnerList []int
	lastWinner:=winUser[0]
	for i:= 0; i < len(winUser); i++ {
		if commonMethod.CompareCard(bEndCard[lastWinner],bEndCard[winUser[i]])==2 {lastWinner=i}
		fmt.Println(222)
	}
	for i:= 0; i < len(winUser); i++ {
		x:= commonMethod.CompareCard(bEndCard[lastWinner],bEndCard[winUser[i]])
		if x==0 {
			lastWinnerList=append(lastWinnerList,winUser[i])
		}
	}
	if lastWinnerList==nil {
		lastWinnerList=append(lastWinnerList,lastWinner)
		fmt.Println("lastWinnerList add")
	}
	fmt.Println("lastWinner",lastWinner,"lastWinnerList",lastWinnerList)
	winScore:=t.tableInfo.lPotScore/ len(lastWinnerList)
	//fmt.Println("winUser",winUser,"底池",winScore)
	for i:= 0; i < len(lastWinnerList); i++ {
		t.userInfo[lastWinnerList[i]].score+=winScore
	}
	lPotScore:=t.tableInfo.lPotScore
	t.tableInfo.lPotScore=0
	t.tableInfo.balanceCount=0
	t.tableInfo.balanceScore=0
	userCount:=0
	for i:= 0; i < len(t.userInfo); i++ {
		if t.userInfo[i].score>0 {
			t.userInfo[i].addScoreStatus=enAddScoreGIVEUP
			t.userInfo[i].tableStatus=true
			userCount++
		}else {
			t.userInfo[i].tableStatus=false
		}
	}
	if userCount>1{
		//time.Sleep(time.Duration(3) * time.Second)
		//fmt.Println(t.userInfo)
		num :=0
		isErr := false
		for i := 0; i< len(t.userInfo);i++  {
			fmt.Println(t.chairSlice[i],t.userInfo[i].score,t.userInfo[i].tableStatus)
			num+=t.userInfo[i].score
			if t.userInfo[i].score < -18{
				isErr=true
			}
		}
		fmt.Println("allScore",num,userCount)
		if isErr{
			fmt.Println("num err",t.tableInfo.dealer,lPotScore,len(lastWinnerList))
			return
		}
		if num==9000{
			t.gameStart()
		}else {
			fmt.Println("num err",t.tableInfo.dealer,lPotScore,len(lastWinnerList))
			t.gameStart()
		}

	}else {
		num :=0
		for i := 0; i< len(t.userInfo);i++  {
			fmt.Println(t.chairSlice[i],t.userInfo[i].score,t.userInfo[i].tableStatus)
			num+=t.userInfo[i].score
		}
		fmt.Println("allScore",num,userCount)
		fmt.Println(t.tableInfo,"牌桌结束")
	}

}

func (t *Table)getCardArr()  {
	i:=0
	CardData:=make([]byte,52)
	for _,v:=range cardMap{
		if i> len(CardData) {
			break
		}
		CardData[i]=v
		i++
	}
	t.btCenterCard=make([]byte,5)
	copy(t.btCenterCard,CardData)
	fmt.Println(CardData[len(t.btCenterCard):])

	fmt.Println("11111",t.chairSlice,len(t.btCenterCard))
	i= len(t.btCenterCard)
	for  x:=0;x< len(t.chairSlice); x++ {
		if i+1>= len(CardData){
			fmt.Println(i, len(CardData))
			break
		}
		t.userInfo[x].handCard=make([]byte,2)
		copy(t.userInfo[x].handCard,CardData[i:])
		//fmt.Println("33333",t.chairSlice[x],t.userInfo[x].handCard)
		i=i+2
	}

}

func (t *Table)sendHandCard()  {
	handCardData:=make([]byte,2)
	for chairID,userID :=range t.chairSlice {
		c:=mapClis[int(userID)]
		if c !=nil{
			sendCardData := bytes.NewBuffer([]byte{})
			_ = binary.Write(sendCardData, binary.LittleEndian, c.userID)
			_ = binary.Write(sendCardData, binary.LittleEndian, t.tableCode)
			_ = binary.Write(sendCardData, binary.LittleEndian, uint16(chairID))
			copy(handCardData,t.userInfo[chairID].handCard[0:])
			sendCardData.Write(handCardData)
			fmt.Println(c.i+beginUserID,commonMethod.GetCardStr(handCardData[0]),commonMethod.GetCardStr(handCardData[1]),t.userInfo[chairID].score)
			/* 发送 */
			c.sendData(199, 102, sendCardData.Bytes())
		}
	}
}

func (t *Table)sendCenterCard(balanceCount int)  {
	//第1次下注平衡后就开始发给三张公牌
	//第2次下注平衡后就开始发第四张公牌
	//第3次下注平衡后就开始发第五张公牌
	//第4次下注平衡后就结束游戏
	fmt.Println("平衡:",balanceCount)
	sendCard := bytes.NewBuffer([]byte{})
	switch balanceCount {
	case 1:
		_ = binary.Write(sendCard, binary.LittleEndian, int32(3))
		cardData:=make([]byte,5)
		copy(cardData,t.btCenterCard[0:2])
		sendCard.Write(cardData)
		t.sendTableData(200,153,sendCard.Bytes())
	case 2:
		_ = binary.Write(sendCard, binary.LittleEndian, int32(4))
		cardData:=make([]byte,5)
		copy(cardData,t.btCenterCard[0:3])
		sendCard.Write(cardData)
		t.sendTableData(200,153,sendCard.Bytes())
	case 3:
		_ = binary.Write(sendCard, binary.LittleEndian, int32(5))
		cardData:=make([]byte,5)
		copy(cardData,t.btCenterCard[0:4])
		sendCard.Write(cardData)
		t.sendTableData(200,153,sendCard.Bytes())
	case 4:
		t.gameEnd()
	default:
		break
	}
}

func (t *Table)addUserScore(chairID int,score int,enAddScoreUserState byte)  {

	//score
	nDifference := t.tableInfo.balanceScore-t.userInfo[chairID].tableScore	//差额
	remainScore := t.userInfo[chairID].score  //剩余
	if remainScore<0{remainScore=0}

	if t.userInfo[chairID].tableStatus==false{
		fmt.Println("111")
		goto addScoreEnd
	}

	//加注不够的情况下
	if score < nDifference {
		if remainScore>=nDifference {
			score = nDifference	    //身上剩余钱够付差额，加注值修正为差额值
		}else {
			score = remainScore	//allin情况下，加注值为剩余钱
			t.userInfo[chairID].addScoreStatus=enAddScoreSHOWHAND
		}
	}else if score > nDifference{      //加注
		if score > remainScore {      //超出判断
			score = remainScore
		}
	}
	if t.userInfo[chairID].tableStatus==false{
		fmt.Println("111")
		goto addScoreEnd
	}
	if score<0 || enAddScoreUserState==enAddScoreGIVEUP {
		fmt.Println("addScore err",score)
		score=0
	}
	if t.userInfo[chairID].tableStatus==false{
		fmt.Println("111")
		goto addScoreEnd
	}
	fmt.Println("玩家",t.chairSlice[chairID],"加注前",t.userInfo[chairID].score,"最小注",score)
	t.userInfo[chairID].score-=score
	t.userInfo[chairID].tableScore+=score
	t.userInfo[chairID].betScore+=score
	t.tableInfo.lPotScore+=score
	if t.userInfo[chairID].score<=0 && t.userInfo[chairID].addScoreStatus!=enAddScoreSHOWHAND{
		fmt.Println(t.chairSlice[chairID],"is out")
		t.userInfo[chairID].tableStatus=false
	}
	for i := 0; i< len(t.userInfo); i++ {
		if t.userInfo[i].tableScore>t.tableInfo.balanceScore{
			t.tableInfo.balanceScore=t.userInfo[i].tableScore
		}
	}
	if t.userInfo[chairID].tableScore>t.tableInfo.balanceScore{
		t.tableInfo.balanceScore=t.userInfo[chairID].tableScore
	}
	goto addScoreEnd
addScoreEnd:
	nextUser:=0
	isEnd:=true
	for i := 0; i < t.maxUser; i++ {
		nextUser=(t.tableInfo.currentUser+i+1)%t.maxUser
		t.tableInfo.chopCount++
		if t.userInfo[nextUser].tableStatus == true &&
			t.userInfo[nextUser].addScoreStatus!=enAddScoreGIVEUP &&
			t.userInfo[nextUser].addScoreStatus!=enAddScoreSHOWHAND {
			isEnd=false
			break
		}
	}
	if !isEnd{
		t.tableInfo.currentUser=nextUser
		if t.tableInfo.chopCount >= t.maxUser {
			t.tableInfo.balanceCount++
			t.tableInfo.chopCount=0

			if t.tableInfo.balanceCount >= 4 {
				t.sendCenterCard(t.tableInfo.balanceCount)
				fmt.Println(4444)
				return
			}else {
				t.sendCenterCard(t.tableInfo.balanceCount)
			}
		}
	}else{
		t.sendCenterCard(4)
		fmt.Println("isEnd")
		return
	}
	if t.tableInfo.balanceCount >= 4 {
		fmt.Println(44442)
		return
	}
	//发送数据
	addScore := bytes.NewBuffer([]byte{})
	//8-21
	_ = binary.Write(addScore, binary.LittleEndian, int32(score))
	_ = binary.Write(addScore, binary.LittleEndian, int32(t.userInfo[chairID].score))
	_ = binary.Write(addScore, binary.LittleEndian, int32(chairID))
	_ = binary.Write(addScore, binary.LittleEndian, enAddScoreUserState)
	//21-33
	_ = binary.Write(addScore, binary.LittleEndian, int32(t.tableInfo.currentUser))
	_ = binary.Write(addScore, binary.LittleEndian, int32(t.userInfo[t.tableInfo.currentUser].score))
	_ = binary.Write(addScore, binary.LittleEndian, int32(t.userInfo[t.tableInfo.currentUser].tableScore))
	//33-41
	_ = binary.Write(addScore, binary.LittleEndian, int32(t.tableInfo.balanceScore))
	_ = binary.Write(addScore, binary.LittleEndian, int32(t.tableInfo.lPotScore))
	fmt.Println(nextUser,t.chairSlice[nextUser],"发送加注")
	t.sendTableData(200,151,addScore.Bytes())

}

func (t *Table)sendTableData(nMainID int16, nSubID int16, szData []byte)  {
	for _,userID:=range t.chairSlice {
		c:=mapClis[int(userID)]
		if c != nil {
			addScore := bytes.NewBuffer([]byte{})
			_ = binary.Write(addScore, binary.LittleEndian, int32(t.tableCode))
			_ = binary.Write(addScore, binary.LittleEndian, int32(c.userID))
			_ = binary.Write(addScore, binary.LittleEndian, int32(7))
			_ = binary.Write(addScore, binary.LittleEndian, int32(GameType))
			sendData:=append(addScore.Bytes(),szData...)
			c.sendData(nMainID, nSubID, sendData)
		}
	}
}


var mapClis = make(map[int]*Client)
var mapTable = make(map[int32]*Table)

func main() {
	l, err := net.Listen("tcp", ":29220")
	if err != nil {
		fmt.Println("error listen:", err)
		return
	}
	defer l.Close()
	fmt.Println("listen ok")
	var index int
	go func() {
		for {
			conn, err := l.Accept()
			if  err != nil {
				fmt.Println(index," accept error:", err)
				continue
			}
			client := &Client{
				i:		index,
				userID:  0,
				chairID: 0xFFFF,
				tableID: 0,
				tableCode: 0,
				score:  1000,
				login:	false,
				conn:	conn,
				stopChan: make(chan bool),
				sendChan: make(chan []byte),
				gameType: 8,
			}
			//mapClis[index] = client
			go client.dealSend()
			go client.dealRecv()
			index++
			//fmt.Printf("%d: accept a new connection\n", index)
		}
	}()

	for {
		for _, client := range mapClis {
			client.sendData(0, 1, []byte{})
		}

		fmt.Println(putOutNowTime(), len(mapClis), "clients are running...")
		time.Sleep(time.Duration(300) * time.Second)
	}
}


func encryptData(szData []byte, nDataLen int) {
	var nCheckCode byte = 0

	for i:=4; i<nDataLen; i++ {
		nCheckCode += szData[i]
		szData[i] = SendByteMap[szData[i]]
	}

	szData[0] = 0x01
	szData[1] = ^nCheckCode + 1

	szLen := bytes.NewBuffer([]byte{})

	_ = binary.Write(szLen, binary.LittleEndian, int16(nDataLen))

	copy(szData[2:], szLen.Bytes())
}

func crevasseData(szData []byte, nDataLen int) {
	var nCheckCode byte

	nCheckCode = szData[1]

	for i:=4; i<nDataLen; i++ {
		szData[i] = RecvByteMap[szData[i]]
		nCheckCode += szData[i]
	}

}

////////////////////////////////////////////////////////////////
/* 读配置文件相关 */
type Config struct {
	Mymap  map[string]string
	strcet string
}

func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + "." + frist
		c.Mymap[key] = strings.TrimSpace(second)
	}
}

func (c Config) Read(node, key string) string {
	key = node + "." + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}

