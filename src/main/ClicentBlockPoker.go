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
var SngType int
var TableName string
var MainAllianceID int32
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

/* 客户端封装 */
type Client struct {
	i		int
	login	bool
	chairID	uint16
	tableID	uint16
	tableCode int32
	conn	net.Conn
	logonChan   chan bool
	stopChan	chan bool
	sendChan	chan []byte
	gameType	int
}

func init() {
	conf := new(Config)
	conf.InitConfig("./ClicentBlockPoker.ini")

	RoutineNum, _ = strconv.Atoi(conf.Read("Base", "cliNum"))
	beginUserID, _ = strconv.Atoi(conf.Read("Base", "beginUID"))
	GameType, _ = strconv.Atoi(conf.Read("Base", "gameType"))
	TableName = conf.Read("Base", "tableName")
	SngType,_ = strconv.Atoi(conf.Read("Base","sngType"))
	RobotTableID=0
}

/*常用方法 */

func putOutNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05  ")
}

/* 创建牌桌 */
func (c *Client) createTable() {
	fmt.Println(c.i+beginUserID,"创建牌桌")
	createTable := bytes.NewBuffer([]byte{})

	/* 创建源头, 0个人 1圈子 2俱乐部 3联盟 */
	binary.Write(createTable, binary.LittleEndian, uint16(3))
	/* 创建人ID */
	binary.Write(createTable, binary.LittleEndian, int32(MainAllianceID))
	/*牌桌名*/
	userName := make([]byte, 64)
	name := []byte(strconv.Itoa(c.i+beginUserID))
	copy(userName, name)
	createTable.Write(userName)
	//wKindID
	binary.Write(createTable, binary.LittleEndian, uint16(7))
	/* 玩法类型 0普通局；1MTT9；3.猎人赛 4.奥马哈 6.SNG */
	switch GameType {
	case 0:
		binary.Write(createTable, binary.LittleEndian, uint16(GameType))
	case 1:
		binary.Write(createTable, binary.LittleEndian, uint16(GameType))
	case 3:
		binary.Write(createTable, binary.LittleEndian, uint16(GameType))
	case 4:
		binary.Write(createTable, binary.LittleEndian, uint16(GameType))
	case 6:
		binary.Write(createTable, binary.LittleEndian, uint16(GameType))
	default:
		fmt.Printf("GameType %d is not ligel,Create table fail!\n", GameType)
		return
	}

	/* 牌桌名称 */
	tableName := make([]byte, 64)
	name = []byte(TableName + "牌桌")
	copy(tableName, name)
	createTable.Write(tableName)
	//ip,gps
	binary.Write(createTable, binary.LittleEndian, false)
	binary.Write(createTable, binary.LittleEndian, false)

	/* 椅子数 */
	if GameType==6{
		binary.Write(createTable, binary.LittleEndian, int32(SngType))
	}else {
		binary.Write(createTable, binary.LittleEndian, int32(9))
	}

	/* 游戏属性 */
	switch GameType {
	case 6:
		binary.Write(createTable, binary.LittleEndian, int32(1000))//nStartScore
		binary.Write(createTable, binary.LittleEndian, int32(1500))//nRequestScore
		binary.Write(createTable, binary.LittleEndian, int32(150))//nServiceScore
		binary.Write(createTable, binary.LittleEndian, int32(60))//nBlindRoseTime
		binary.Write(createTable, binary.LittleEndian, int32(0))//nStartTime
		binary.Write(createTable, binary.LittleEndian, false)//bAuthorizedEntry
		binary.Write(createTable, binary.LittleEndian, int32(0))//nAddOnCount
		binary.Write(createTable, binary.LittleEndian, int32(0))//nRebuyCount
		binary.Write(createTable, binary.LittleEndian, int32(0))//nBlindRoseType
	default:
		fmt.Printf("GameType %d is not ligel,Create table fail!\n", GameType)
		return
	}

	/* 发送 */
	c.sendData(199, 1, createTable.Bytes())
}

/* 进入牌桌 */
func (c *Client) enterTable(nTableID uint16, nTableCode int32) {
	c.tableID = nTableID
	c.tableCode = nTableCode
	enterTable := bytes.NewBuffer([]byte{})
	binary.Write(enterTable, binary.LittleEndian, nTableID)
	binary.Write(enterTable, binary.LittleEndian, uint16(1))
	binary.Write(enterTable, binary.LittleEndian, nTableCode)
	binary.Write(enterTable, binary.LittleEndian, int32(c.i+beginUserID))
	/* 发送 */
	c.sendData(199, 4, enterTable.Bytes())
}
func (c *Client) backToGame(nTableID uint16, nTableCode int32, nSoure uint16) {
	c.tableID = nTableID
	//c.tableCode = nTableCode
	enterTable := bytes.NewBuffer([]byte{})
	binary.Write(enterTable, binary.LittleEndian, nTableID)
	binary.Write(enterTable, binary.LittleEndian, nSoure)
	binary.Write(enterTable, binary.LittleEndian, nTableCode)
	binary.Write(enterTable, binary.LittleEndian, int32(c.i+beginUserID))
	/* 发送 */
	c.sendData(199, 4, enterTable.Bytes())
}
/* 坐下 */
func (c *Client) sitDown() {
	sitDown := bytes.NewBuffer([]byte{})

	binary.Write(sitDown, binary.LittleEndian, c.tableCode)
	binary.Write(sitDown, binary.LittleEndian, int32(c.i+beginUserID))
	binary.Write(sitDown, binary.LittleEndian, c.tableID)
	binary.Write(sitDown, binary.LittleEndian, c.chairID)

	/* 发送 */
	c.sendData(199, 22, sitDown.Bytes())
}

/* 报名 */
func (c *Client) matchSighUp() {
	fmt.Println(c.i+beginUserID,"报名")
	enterTable := bytes.NewBuffer([]byte{})
	binary.Write(enterTable, binary.LittleEndian, c.tableCode)
	binary.Write(enterTable, binary.LittleEndian, int32(c.i+beginUserID))
	/* 发送 */
	c.sendData(199, 40, enterTable.Bytes())
}

/* 买分 */
func (c *Client) buyScore() {
	buyScore := bytes.NewBuffer([]byte{})

	binary.Write(buyScore, binary.LittleEndian, int32(c.i+beginUserID))
	binary.Write(buyScore, binary.LittleEndian, c.tableCode)
	binary.Write(buyScore, binary.LittleEndian, int32(100))

	/* 发送 */
	c.sendData(199, 24, buyScore.Bytes())
}

/* 开始游戏 */
func (c *Client) startGame() {
	startGame := bytes.NewBuffer([]byte{})

	binary.Write(startGame, binary.LittleEndian, c.tableCode)

	/* 发送 */
	c.sendData(199, 29, startGame.Bytes())
}

/* 抢庄 */
func (c *Client) beginSnach(ableSnatch bool) {
	/* 随机数秒延迟操作,防止机器人动作太快 */
	time.Sleep(time.Duration(rand.Intn(8) + 1) * time.Second)
	beginSnach := bytes.NewBuffer([]byte{})

	binary.Write(beginSnach, binary.LittleEndian, c.tableCode)
	binary.Write(beginSnach, binary.LittleEndian, int32(c.i+beginUserID))
	binary.Write(beginSnach, binary.LittleEndian, int32(8))
	binary.Write(beginSnach, binary.LittleEndian, int32(8))

	binary.Write(beginSnach, binary.LittleEndian, ableSnatch)
	c.sendData(200, 1, beginSnach.Bytes())
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
	//time.Sleep(time.Duration(1) * time.Second)
	addScore := bytes.NewBuffer([]byte{})

	binary.Write(addScore, binary.LittleEndian, int32(c.tableCode))
	binary.Write(addScore, binary.LittleEndian, int32(c.i+beginUserID))
	binary.Write(addScore, binary.LittleEndian, int32(7))
	binary.Write(addScore, binary.LittleEndian, int32(GameType))

	binary.Write(addScore, binary.LittleEndian, int32(0))
	binary.Write(addScore, binary.LittleEndian, byte(2))

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

				/* 掉线后5秒后充连 */
				time.Sleep(time.Duration(5) * time.Second)
				client, err := connect(c.i)
				if err != nil {
					fmt.Println("client", c.i, "connect to server fail!")
					continue
				}

				go client.dealSend()
				go client.dealRecv()

				mapClis[c.i] = client
				client.logon()

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
	var szLen = make([]byte, 2)
	/* 定义一个定时器，用于发送心跳 */
	timer := time.NewTimer(time.Second * 60)
	for {
		select {
		case <- c.stopChan:
			goto RecvEnd
		/* 一分钟定时器，发送心跳 */
		case <-timer.C:
			c.sendData(0, 1, []byte{})
		default:
			break
		}

		msgLen, err := c.conn.Read(recvData[:4])
		if err != nil || msgLen <= 0 {
			fmt.Println(c.i+beginUserID,"Read from Socket Fail!", err)

			/* 掉线后5秒后充连 */
			time.Sleep(time.Duration(5) * time.Second)
			client, err := connect(c.i)
			if err != nil {
				fmt.Println("client", c.i, "connect to server fail!")
				continue
			}

			go client.dealSend()
			go client.dealRecv()

			mapClis[c.i] = client
			client.logon()

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
			/* 登陆成功 */
			case 101:
				if nSubID == 1 {
					c.login = true
					var nUserID int32
					userID := make([]byte, 4)
					copy(userID, data[8:])

					codeBuf := bytes.NewBuffer(userID)
					binary.Read(codeBuf, binary.LittleEndian, &nUserID)

					fmt.Println(nUserID,"登陆成功")
					if c.i == 0 {
						c.createTable()
					}
				}
			/*牌桌消息*/
			case 199:
				switch nSubID {
				//创建牌桌成功
				case 50:
					var nTableCode int32
					tableCode := make([]byte, 4)
					copy(tableCode, data[8:])
					codeBuf := bytes.NewBuffer(tableCode)
					binary.Read(codeBuf, binary.LittleEndian, &nTableCode)

					fmt.Println(putOutNowTime(), c.i+beginUserID, "创建牌桌成功:", nTableCode)
					RobotTableID=int(nTableCode)
					//c.tableCode = nTableCode
					/* 睡眠2秒，预防后端牌桌还没准备好 */
					time.Sleep(time.Duration(2) * time.Second)
					for _, client := range mapClis {
						client.enterTable(0, nTableCode)
					}
				/*用户状态*/
				case 102:
					//var status byte = data[36]
					var nUserID int32
					var nTableID int32
					var nChairID uint16
					handCardData :=make([]byte,2)

					userID := make([]byte, 4)
					copy(userID, data[8:])

					tableID := make([]byte,4)
					copy(tableID,data[12:])

					chairID := make([]byte, 2)
					copy(chairID, data[16:])

					copy(handCardData,data[18:])

					userIDBuf := bytes.NewBuffer(userID)
					binary.Read(userIDBuf, binary.LittleEndian, &nUserID)

					tableIDBuf := bytes.NewBuffer(tableID)
					binary.Read(tableIDBuf, binary.LittleEndian, &nTableID)

					chairIDBuf := bytes.NewBuffer(chairID)
					binary.Read(chairIDBuf, binary.LittleEndian, &nChairID)

					fmt.Println(c.i+beginUserID,nTableID,nChairID,commonMethod.GetCardStr(handCardData[0]),commonMethod.GetCardStr(handCardData[1]))

					//fmt.Println(putOutNowTime(),c.i+beginUserID,nUserID," chairID=",nChairID)
					c.chairID = nChairID
				default:
					break
				}
			/*游戏消息*/
			case 200:
				//fmt.Println(200,"SubID:",nSubID)
				switch nSubID {
				case 151:
					var nTableCode int32
					tableCode := make([]byte, 4)
					copy(tableCode, data[8:])
					codeBuf := bytes.NewBuffer(tableCode)
					binary.Read(codeBuf, binary.LittleEndian, &nTableCode)

					var nChairID uint16
					chairID := make([]byte, 2)
					copy(chairID, data[37:])
					chairIDBuf := bytes.NewBuffer(chairID)
					binary.Read(chairIDBuf, binary.LittleEndian, &nChairID)
					fmt.Println(nTableCode,nChairID)
					if nChairID == c.chairID && nTableCode==c.tableCode{
						fmt.Println(putOutNowTime(),"c.chairID=",c.chairID," nChairID=",nChairID)
						c.TexasAddScore()
					}

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

var mapClis = make(map[int]*Client)

func main() {
	for i:=0; i<RoutineNum; i++ {
		client, err := connect(i)
		if err != nil {
			fmt.Println("client", i, "connect to server fail!")
			continue
		}
		mapClis[i] = client
		go client.dealSend()
		go client.dealRecv()
		client.logon()
	}
	for {
		for _, client := range mapClis {
			client.sendData(0, 1, []byte{})
		}
		fmt.Println(putOutNowTime(), len(mapClis), "clients are running...")
		time.Sleep(time.Duration(300) * time.Second)
	}
}

/* 连接服务器 */
func connect(index int) (*Client, error){
	conn, err := net.Dial("tcp", ":29220")
	if err != nil {
		return nil, err
	}

	client := &Client{
		i:		index,
		chairID: 0xFFFF,
		tableID: 0,
		tableCode: 0,
		login:	false,
		conn:	conn,
		stopChan: make(chan bool),
		sendChan: make(chan []byte),
		gameType: 8,
	}

	return client, nil
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

	binary.Write(szLen, binary.LittleEndian, int16(nDataLen))

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

