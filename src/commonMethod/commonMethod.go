package commonMethod

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
	"bytes"
	"encoding/binary"
)




type Student struct {
	Name string
	Id   int
}

func Ssss() {
	s := make(map[string]*Student)
	s["chenchao"] = &Student{
		Name: "chenchao",
		Id:   111,
	}
	s["chenchao"].Id = 222
	for _, ss := range s {
		ss.Id = 333
	}
	fmt.Println(s["chenchao"].Id)
}

func SayHello() {
	fmt.Println("SayHello()-->Hello")
}

func PutOutNowTime() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05  "))
}

func CreateRandNumber() int32 {
	//site :=fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//i,_ :=strconv.ParseInt(site,0,32)
	//rand.Intn(8)
	//return int32(i)
	sum:=(rand.Intn(9)+1)*100000 + rand.Intn(10)*10000 + rand.Intn(10)*1000 + rand.Intn(10)*100 + rand.Intn(10)*10+rand.Intn(10)
	return int32(sum)
}



func ReadInt32FromData(data []byte,beginPath int) int32 {
	var nTableCode int32
	tableCode := make([]byte, 4)
	copy(tableCode, data[beginPath:])
	codeBuf := bytes.NewBuffer(tableCode)
	binary.Read(codeBuf, binary.LittleEndian, &nTableCode)
	return nTableCode
}

func ReadByteFromData(data []byte,beginPath int) byte  {
	var nByte byte
	byteData:=make([]byte,1)
	copy(byteData,data[beginPath:])
	byteBuf := bytes.NewBuffer(byteData)
	binary.Read(byteBuf, binary.LittleEndian, &nByte)
	return nByte
}


func GetCardStr(card byte) string {
	var str1  string
	var str2  string
	cardType:= int(card/0x10)
	cardValue:= int(card%0x10)
	switch cardType {
	case 0:
		str1="方块"
	case 1:
		str1="梅花"
	case 2:
		str1="红桃"
	case 3:
		str1="黑桃"
	}
	switch cardValue {
	case 1:
		str2="A"
	case 11:
		str2="J"
	case 12:
		str2="Q"
	case 13:
		str2="K"
	default:
		str2=strconv.Itoa(cardValue)
	}
	return  str1+str2
}
func GetAllCardStr(card []byte) string {
	str1:=GetCardStr(card[0])
	for i := 1; i< len(card);i++  {
		str1=str1+" "+GetCardStr(card[i])
	}
	return str1
}

func SortCard(cardData []byte)  {
	var valueData []byte
	valueData=make([]byte, len(cardData))
	for i:=0;i< len(cardData);i++{
		valueData[i]=GetCardLogicValue(cardData[i])
	}
	beSorted:=false
	bLast:= len(cardData)-1
	var bTempData byte
	for true {
		beSorted=true
		for i := 0; i < bLast; i++ {
			if valueData[i]<valueData[i+1] || (valueData[i]==valueData[i+1] && cardData[i]<cardData[i+1]){
				bTempData=valueData[i]
				valueData[i]=valueData[i+1]
				valueData[i+1]=bTempData
				bTempData=cardData[i]
				cardData[i]=cardData[i+1]
				cardData[i+1]=bTempData
				beSorted=false
			}
		}
		bLast--
		if beSorted {
			break
		}
	}
}
func GetCardValue(bCardDate byte) byte {
	return bCardDate&0x0F
}
func GetCardColor(bCardDate byte) byte {
	return bCardDate&0xF0
}
func GetCardLogicValue(bCardDate byte) byte {
	bValue:=GetCardValue(bCardDate)
	if bValue==1 {
		return 0x0e
	}else{
		return bValue
	}
}

type tagAnalyseResult struct {
	fourCount                      byte						//四张数目
	threeCount                     byte						//三张数目
	twoCount                       byte						//两张数目
	signedCount                    byte						//单张数目
}
func AnalyseCardData(cardData []byte) tagAnalyseResult {
	var AnalyseResult =tagAnalyseResult{
		fourCount:0,
		threeCount:0,
		twoCount:0,
		signedCount:0,
	}
	for i := 0; i< len(cardData); i++  {
		sameCount:=1
		logicValue:=GetCardLogicValue(cardData[i])
		for j := i + 1; j< len(cardData); j++  {
			if cardData[j]==0{continue}
			if logicValue!=GetCardLogicValue(cardData[j]){break}
			sameCount++
		}
		switch sameCount {
		case 1:
			AnalyseResult.signedCount++
		case 2:
			AnalyseResult.twoCount++
		case 3:
			AnalyseResult.threeCount++
		case 4:
			AnalyseResult.fourCount++
		default:
			break
		}
		i+=sameCount-1
	}
	return AnalyseResult
}
func GetCardType(bCardDate []byte) byte {
	if len(bCardDate)!=5 {
		return 0
	}
	isSameColor := true
	isLineCard := true
	bFirstValue := GetCardLogicValue(bCardDate[0])
	bFirstColor := GetCardColor(bCardDate[0])
	//i:=0
	for i := 0; i< len(bCardDate); i++  {
		if bFirstColor!=GetCardColor(bCardDate[i]) {isSameColor=false}
		if bFirstValue!=GetCardLogicValue(bCardDate[i])+byte(i) {isLineCard=false}
		if !isSameColor && !isLineCard{
			break
		}
	}
	if !isLineCard && bFirstValue==0x0E{
		i:=1
		for i = 1; i< len(bCardDate); i++ {
			if bFirstValue!=(GetCardLogicValue(bCardDate[i])+byte(i)+8) {
				break
			}
		}
		if i == len(bCardDate) {
			isLineCard=true
		}
	}
	if isLineCard && isSameColor && GetCardLogicValue(bCardDate[1])==13 {
		return 10      //皇家同花顺
	}
	if isSameColor && isLineCard {
		return 9      //同花顺型
	}
	if isSameColor && !isLineCard{
		return 6      //同花类型
	}
	if !isSameColor && isLineCard{
		return 5      //顺子类型
	}


	AnalyseResult:=AnalyseCardData(bCardDate)
	if AnalyseResult.fourCount==1{
		return 8        //铁支类型
	}
	if AnalyseResult.threeCount==1 && AnalyseResult.twoCount==1{
		return 7        //葫芦类型
	}
	if AnalyseResult.threeCount==1 && AnalyseResult.twoCount==0 {
		return 4         //三条类型
	}
	if AnalyseResult.twoCount==2 {
		return 3         //两对类型
	}
	if AnalyseResult.twoCount==1 && AnalyseResult.signedCount==3 {
		return 2          //对子类型
	}
	return 0x01
}

func DeleteFromByteSlice(sliceData []byte,x int)[]byte{
	if x> len(sliceData)-1{
		fmt.Println("DeleteFromSlice err")
		return nil
	}
	bTempData:=make([]byte, len(sliceData)-1)
	copy(bTempData[:x],sliceData[:x])
	copy(bTempData[x:],sliceData[x+1:])
	return bTempData
}

func CompareCard(firstData []byte,nextData []byte) byte  {
	if len(firstData)!=5 || len(nextData)!=5{
		fmt.Println("CompareCard err")
		return 1
	}
	firstType:=GetCardType(firstData)
	nextType:=GetCardType(nextData)
	if firstType>nextType {
		return 1
	}
	if firstType<nextType {
		return 2
	}

	if firstType!=1{
		i:=0
		for i= 0; i< len(firstData);  i++{
			firstValue:=GetCardLogicValue(firstData[i])
			nextValue:=GetCardLogicValue(nextData[i])
			if firstValue>nextValue { return 1}
			if firstValue<nextValue { return 2}
			if firstValue==nextValue { continue}
		}
		if i == len(firstData) {return 0}
	}

	switch firstType {
	case 1:
		i:=0
		for i= 0; i< len(firstData);  i++{
			firstValue:=GetCardLogicValue(firstData[i])
			nextValue:=GetCardLogicValue(nextData[i])
			if firstValue>nextValue { return 1}
			if firstValue<nextValue { return 2}
			if firstValue==nextValue { continue}
		}
		if i == len(firstData) {return 0}
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
		return 0
	default:
		break
	}

	return 1
}

func GetGreatestCardType(bHandCard []byte,bCenterCard []byte) (byte,[]byte) {
	var bTempCard []byte
	bTempCard=append(bCenterCard,bHandCard...)
	SortCard(bTempCard)
	//fmt.Println("SortCard",GetAllCardStr(bTempCard),bTempCard)
	//if len(bTempCard)==5 {
	//	return GetCardType(bTempCard),bTempCard
	//}

	//var lastCard []byte
	//var greatestCard = DeleteFromByteSlice(bTempCard,0)
	//cardKind:=GetCardType(greatestCard)
	//if len(bTempCard)==6 {
	//	for i := 0; i< len(bTempCard);i++  {
	//		lastCard=DeleteFromByteSlice(bTempCard,i)
	//		if cardKind<=GetCardType(lastCard){
	//			cardKind=GetCardType(lastCard)
	//			copy(greatestCard,lastCard)
	//		}
	//	}
	//}
	//if cardKind==1{
	//	greatestCard=make([]byte,5)
	//	copy(greatestCard,bTempCard)
	//}
	//
	//return cardKind,greatestCard
	return GetFinalCardType(bTempCard)
}

func GetFinalCardType(cardData []byte)(byte,[]byte){
	if len(cardData)<=5 {
		return GetCardType(cardData),cardData
	}else {
		var lastCard []byte
		var greatestCard = DeleteFromByteSlice(cardData,0)
		cardKind:=GetCardType(greatestCard)
		for i := 0; i< len(cardData);i++  {
			lastCard=DeleteFromByteSlice(cardData,i)
			if cardKind<=GetCardType(lastCard){
				cardKind=GetCardType(lastCard)
				copy(greatestCard,lastCard)
			}
		}
		return GetFinalCardType(greatestCard)
	}

}

