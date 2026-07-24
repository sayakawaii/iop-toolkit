package omciDeshape

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"omciAnalyzer/utils"
	"regexp"
	"strconv"
)

type EVTOCD_Deshape struct{}

type ExtVlanTagOpTableEntryData struct {
	filterOuterPriority uint8
	filterOuterVid      uint16
	filterOuterTpidDe   uint8

	filterInnerPriority uint8
	filterInnerVid      uint16
	filterInnerTpidDe   uint8

	filterEtherType uint8

	treatmentTagsToRemove uint8

	treatmentOuterPriority uint8
	treatmentOuterVid      uint16
	treatmentOuterTpidDe   uint8

	treatmentInnerPriority uint8
	treatmentInnerVid      uint16
	treatmentInnerTpidDe   uint8

	bridgePortNum uint8
	/* 0 for create and 1 for delete */
	deleteTableEntryFlag uint8
}

func (e *ExtVlanTagOpTableEntryData) fillExtVlanTagOpTableEntryData(num []uint16) {
	var i int = 1
	e.deleteTableEntryFlag = uint8(num[i])
	i++
	e.filterOuterPriority = uint8(num[i])
	i++
	e.filterOuterVid = uint16(num[i])
	i++
	e.filterOuterTpidDe = uint8(num[i])
	i++
	e.filterInnerPriority = uint8(num[i])
	i++
	e.filterInnerVid = uint16(num[i])
	i++
	e.filterInnerTpidDe = uint8(num[i])
	i++
	e.filterEtherType = uint8(num[i])
	i++
	e.treatmentTagsToRemove = uint8(num[i])
	i++
	e.treatmentOuterPriority = uint8(num[i])
	i++
	e.treatmentOuterVid = uint16(num[i])
	i++
	e.treatmentOuterTpidDe = uint8(num[i])
	i++

	e.treatmentInnerPriority = uint8(num[i])
	i++
	e.treatmentInnerVid = uint16(num[i])
	i++
	e.treatmentInnerTpidDe = uint8(num[i])
	i++

	e.bridgePortNum = uint8(num[i])
}

func (e *ExtVlanTagOpTableEntryData) formatSetVlanTaggingOpTableEntry(tableEntry []uint8) {
	// Table format:
	//----------------------------------------------------------------
	//  Filter Outer Priority (4 bits):	tableEntry [0] = vvvv 0000
	//
	//  Filter Outer VID (13 bits)    :	tableEntry [0] = 0000 vvvv
	//					tableEntry [1] = vvvv vvvv
	//					tableEntry [2] = v000 0000
	//
	//  Filter Outer TPID/DE (3 bits) :	tableEntry [2] = 0vvv 0000
	//
	//  Padding (12 bits)		 :	tableEntry [2] = 0000 vvvv
	//					tableEntry [3] = vvvv vvvv
	//
	//  Filter Inner Priority (4 bits):	tableEntry [4] = vvvv 0000
	//
	//  Filter Inner VID (13 bits)    :	tableEntry [4] = 0000 vvvv
	//					tableEntry [5] = vvvv vvvv
	//					tableEntry [6] = v000 0000
	//
	//  Filter Inner TPID/DE (3 bits) :	tableEntry [6] = 0vvv 0000
	//
	//  Padding (8 bits)      :  tableEntry [6] = 0000 vvvv
	//
	//  Filter EtherType (4 bits)	 :	tableEntry [7] = 0000 vvvv
	//
	//  Treatment (tags to Remove) (2 bits):	tableEntry [8] = vv00 0000
	//
	//  Padding (10 bits)		 :	tableEntry [8] = 00vv vvvv
	//					tableEntry [9] = vvvv 0000
	//
	//  Treatment (outer Priority) (4 bits):	tableEntry [9] = 0000 vvvv
	//
	//  Treatment (outer VID) (13 bits):	tableEntry [10]= vvvv vvvv
	//					tableEntry [11]= vvvv v000
	//
	//  Treatment (outer TPID/DE) (3 bits):	tableEntry [11]= 0000 0vvv
	//
	//  Padding (12 bits)		 :	tableEntry [12]= vvvv vvvv
	//					tableEntry [13]= vvvv 0000
	//
	//  Treatment (inner Priority) (4 bits):	tableEntry [13]= 0000 vvvv
	//
	//  Treatment (inner VID) (13 bits):	tableEntry [14]= vvvv vvvv
	//					tableEntry [15]= vvvv v000
	//
	//  Treatment (outer TPID/DE) (3 bits):	tableEntry [15]= 0000 0vvv
	//----------------------------------------------------------------

	// initialize all to zero
	// unsigned char tableEntry[16] = {0} // 16 bytes

	// Filter Outer Priority
	tableEntry[0] = uint8((e.filterOuterPriority & 0x000F) << 4)

	// Filter Outer Vid
	tableEntry[0] |= uint8((e.filterOuterVid >> 9) & 0x000F)
	tableEntry[1] = uint8((e.filterOuterVid >> 1) & 0x00FF)
	tableEntry[2] = uint8((e.filterOuterVid & 0x0001) << 7)

	// Filter Outer TPID/DE
	tableEntry[2] |= uint8((e.filterOuterTpidDe & 0x0007) << 4)

	// Filter Inner Priority
	tableEntry[4] = uint8((e.filterInnerPriority & 0x000F) << 4)

	// Filter Inner Vid
	tableEntry[4] |= uint8((e.filterInnerVid >> 9) & 0x000F)
	tableEntry[5] = uint8((e.filterInnerVid >> 1) & 0x00FF)
	tableEntry[6] = uint8((e.filterInnerVid & 0x0001) << 7)

	// Filter Inner TPID/DE
	tableEntry[6] |= uint8((e.filterInnerTpidDe & 0x0007) << 4)

	// Filter EtherType
	tableEntry[7] = uint8(e.filterEtherType & 0x000F)
	/* delete entry */
	if e.deleteTableEntryFlag == 1 {
		for i := 8; i < 16; i++ {
			tableEntry[i] = 0xFF
		}
		/*create entry*/
	} else {
		// Treatment Tags To Remove & bridgePort reference
		tableEntry[8] = uint8((e.treatmentTagsToRemove & 0x0003) << 6)

		tableEntry[8] |= uint8((e.bridgePortNum >> 4) & 0x000F)

		// Treatment Outer Priority & bridgePort reference
		tableEntry[9] = uint8(e.treatmentOuterPriority & 0x000F)

		tableEntry[9] |= uint8((e.bridgePortNum << 4) & 0x00F0)

		// Treatment Outer Vid
		tableEntry[10] = uint8((e.treatmentOuterVid >> 5) & 0x00FF)
		tableEntry[11] = uint8((e.treatmentOuterVid & 0x001F) << 3)

		// Treatment Outer TPID/DE
		tableEntry[11] |= uint8(e.treatmentOuterTpidDe & 0x0007)

		// Treatment Inner Priority
		tableEntry[13] = uint8(e.treatmentInnerPriority & 0x000F)

		// Treatment Inner Vid
		tableEntry[14] = uint8((e.treatmentInnerVid >> 5) & 0x00FF)
		tableEntry[15] = uint8((e.treatmentInnerVid & 0x001F) << 3)

		// Treatment Inner TPID/DE
		tableEntry[15] |= uint8(e.treatmentInnerTpidDe & 0x0007)
	}
}

//get numbers from str and return the counter retreived number
func getNum(in string) []uint16 {
	re := regexp.MustCompile("[0-9]+")
	str := re.FindAllString(in, -1)
	num := make([]uint16, len(str))
	cnt := 0
	for _, s := range str {
		i, _ := strconv.Atoi(s)
		// fmt.Printf("found number[%v]: %v\n", cnt, i)

		// num = append(num, uint16(i))
		num[cnt] = uint16(i)
		cnt++
	}
	return num
}

func (e *ExtVlanTagOpTableEntryData) printExtVlanTagOpTableEntryData() {
	fmt.Printf("output EVTOCD rule:\r\n")
	fmt.Printf("(%d,%d,%d),(%d,%d,%d),%d\r\n",
		e.filterOuterPriority,
		e.filterOuterVid,
		e.filterOuterTpidDe,

		e.filterInnerPriority,
		e.filterInnerVid,
		e.filterInnerTpidDe,
		e.filterEtherType)

	fmt.Printf("%d,(%d,%d,%d),(%d,%d,%d)\r\n",
		e.treatmentTagsToRemove,

		e.treatmentOuterPriority,
		e.treatmentOuterVid,
		e.treatmentOuterTpidDe,

		e.treatmentInnerPriority,
		e.treatmentInnerVid,
		e.treatmentInnerTpidDe)
	fmt.Printf("##############################\r\n")
}

var transactionId uint16 = 0

func printHead(wr io.Writer, instId uint16) {
	transactionId++

	var head = [8]uint8{}

	head[0] = uint8((transactionId >> 8) & 0x00ff)
	head[1] = uint8(transactionId & 0x00ff)
	head[2] = 0x48
	head[3] = 0x0a
	head[4] = 0x00
	head[5] = 0xab
	head[6] = uint8((instId >> 8) & 0x00ff)
	head[7] = uint8(instId & 0x00ff)

	for i := 0; i < 8; i++ {
		fmt.Fprintf(wr, "%02x", head[i])
	}

	var mask = [2]uint8{0x04, 0x0}
	for i := 0; i < 2; i++ {
		fmt.Fprintf(wr, "%02x", mask[i])
	}
}

func printEvtocdRawData(wr io.Writer, outRuleRawDat []uint8) {
	for i := 0; i < 16; i++ {
		fmt.Fprintf(wr, "%02x", outRuleRawDat[i])
	}
	// fmt.Printf("\r\n");
}

func printPad(wr io.Writer) {
	var pad = [14]uint8{0}
	for i := 0; i < 14; i++ {
		fmt.Fprintf(wr, "%02x", pad[i])
	}
	fmt.Fprintf(wr, "\r\n")
}

func printEvtocdOmciMsg(wr io.Writer, instId uint16, outRuleRawDat []uint8) {
	printHead(wr, instId)
	printEvtocdRawData(wr, outRuleRawDat)
	printPad(wr)
}

const (
	expected_num_perline = 17
)

func (e *EVTOCD_Deshape) Translate(in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			utils.Log("end of read")
			break
		}
		if len(line) != 0 {
			str := string(line)
			num := getNum(str)
			if len(num) != expected_num_perline {
				utils.Log("failed!!!,file format not correct !!!")
				return errors.New("file format not correct")
			}
			instId := num[0]
			var evtocdRule ExtVlanTagOpTableEntryData
			evtocdRule.fillExtVlanTagOpTableEntryData(num)
			evtocdRule.printExtVlanTagOpTableEntryData()
			outRuleRawData := make([]uint8, 16) // 16 bytes
			evtocdRule.formatSetVlanTaggingOpTableEntry(outRuleRawData)
			printEvtocdOmciMsg(out, instId, outRuleRawData)
		}
	}
	return nil
}

const (
	EVTOCD = "evtocd"
)

func init() {
	newDeshape(EVTOCD, &EVTOCD_Deshape{})
}
