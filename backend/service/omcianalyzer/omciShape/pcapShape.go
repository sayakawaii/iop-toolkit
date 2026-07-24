/*
# ------------------------------------------------------------
# -- pcapShape.go
# --
# -- OMCI PCAP file processor
# --
# -- 2026-3-5
# ------------------------------------------------------------
*/

package omciShape

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"omciAnalyzer/utils"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcapgo"
)

type pcapShaper struct {
	outputPath      string
	defaultOnuID    string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

// EtherType for OMCI: 0xa8c8
const OmciEtherType = 0xa8c8

// PrivateHeaderSize is the length of the vendor private header that precedes the
// OMCI message inside the Ethernet payload of the standard capture format.
const PrivateHeaderSize = 9

// EthernetHeaderSize is the length of a standard Ethernet header
// (6-byte destination MAC + 6-byte source MAC + 2-byte EtherType).
const EthernetHeaderSize = 14

// OMCI device identifier values, located in the 4th byte of an OMCI message.
const (
	omciBaselineDeviceID = 0x0a
	omciExtendedDeviceID = 0x0b
)

// omciMinLen is the minimum number of bytes required to inspect the OMCI header.
const omciMinLen = 8

func (shaper *pcapShaper) extractOnuID(payload []byte) string {
	if len(payload) < PrivateHeaderSize {
		return ""
	}
	header := payload[:PrivateHeaderSize]
	onuIDBytes := make([]byte, 0, PrivateHeaderSize)
	for _, b := range header {
		if b >= 0x20 && b <= 0x7E {
			onuIDBytes = append(onuIDBytes, b)
		}
	}
	if len(onuIDBytes) > 0 {
		return string(onuIDBytes)
	}
	return hex.EncodeToString(header)
}

// isRawOmciMessage reports whether the raw record looks like a bare OMCI message.
// In the private capture format each record stores the OMCI message directly,
// without any Ethernet or private header. OMCI messages carry the device
// identifier (baseline 0x0A or extended 0x0B) in the 4th byte, right after the
// 2-byte transaction correlation identifier and the 1-byte message type.
func isRawOmciMessage(data []byte) bool {
	if len(data) < omciMinLen {
		return false
	}
	deviceID := data[3]
	return deviceID == omciBaselineDeviceID || deviceID == omciExtendedDeviceID
}

func (shaper *pcapShaper) appendOmci(onuID string, timestamp time.Time, omciData []byte) {
	if _, exists := shaper.onus[onuID]; !exists {
		shaper.onus[onuID] = shaper.outputPath + onuID
		shaper.omciMetaDataMap[onuID] = []OmciMetaData{}
	}
	shaper.omciMetaDataMap[onuID] = append(shaper.omciMetaDataMap[onuID], OmciMetaData{
		Timestamp: timestamp,
		RawData:   hex.EncodeToString(omciData),
	})
}

func (shaper *pcapShaper) processPacket(packet gopacket.Packet) {
	raw := packet.Data()
	timestamp := packet.Metadata().Timestamp

	// Standard format: Ethernet frame (EtherType 0xa8c8) carrying a vendor
	// private header followed by the OMCI message. The ONU ID is encoded in the
	// private header.
	if len(raw) >= EthernetHeaderSize &&
		binary.BigEndian.Uint16(raw[12:EthernetHeaderSize]) == OmciEtherType {
		payload := raw[EthernetHeaderSize:]
		if len(payload) <= PrivateHeaderSize {
			return
		}
		onuID := shaper.extractOnuID(payload)
		if onuID == "" {
			utils.Log("failed to extract ONU ID from packet")
			return
		}
		shaper.appendOmci(onuID, timestamp, payload[PrivateHeaderSize:])
		return
	}

	// Private capture format: each record stores a raw OMCI message without any
	// Ethernet or private header. There is no per-packet ONU identifier, so all
	// messages are grouped under the capture file name.
	if isRawOmciMessage(raw) {
		shaper.appendOmci(shaper.defaultOnuID, timestamp, raw)
	}
}

func (shaper *pcapShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("PCAP OMCI shape processing started")

	shaper.outputPath = outputPath
	shaper.defaultOnuID = pcapFileBaseName(logPath)
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	pcapFile, err := os.Open(logPath)
	if err != nil {
		utils.Log("open pcap file failed: " + fmt.Sprintf("%s", err))
		return nil
	}
	defer pcapFile.Close()

	reader, err := pcapgo.NewReader(pcapFile)
	if err != nil {
		utils.Log("create pcap reader failed: " + fmt.Sprintf("%s", err))
		return nil
	}
	packetSource := gopacket.NewPacketSource(reader, reader.LinkType())
	for packet := range packetSource.Packets() {
		shaper.processPacket(packet)
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)
	utils.Log("PCAP OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

// pcapFileBaseName returns the capture file name without its directory and
// extension, used as the ONU identifier for the headerless capture format.
func pcapFileBaseName(logPath string) string {
	base := filepath.Base(logPath)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func init() {
	shaper := new(pcapShaper)
	shaperRegist(Pcap, shaper)
}
