package model

import (
	"bytes"
	"encoding/binary"
)

// PDUHeader
//
// A PDU Header Record shall be the first part of each PDU.
//
//
//	Field Size 		Field Name 					Data Type
//	(bytes)
//	1 				Protocol Version  			8-bit enumeration
//	1 				Compatibility Version	 	8-bit enumeration
//	1 				Exercise ID 				8-bit unsigned integer
//	1 				PDU Type 					8-bit enumeration
//	1 				PDU Status 					8-bit record
//	1 				PDU Header Length	 		8-bit unsigned integer
//	2 				PDU Length 					16-bit unsigned integer
//	8 				Timestamp 					64-bit signed integer
//	Total PDU Header record size = 16 bytes
type PDUHeader struct {
	//
	//ProtocolVersion
	//This field shall specify the version of protocol used in a PDU.
	//
	//	Value			Description
	//	0				Other
	//	1				DIS PDU version 1.0 (May 92)
	//	2				IEEE 1278-1993
	//	3				DIS PDU version 2.0 - third draft (May 93)
	//	4				DIS PDU version 2.0 - fourth draft (revised) March 16, 1994
	//	5				IEEE 1278.1-1995
	ProtocolVersion      uint8
	CompatibilityVersion uint8
	//	ExerciseID
	//
	//	Exercise Identification shall be unique to each exercise being conducted simultaneously on the same communications medium.
	//
	ExerciseID uint8
	//	PUDType
	//	Value	Description
	//  0	Other
	//  1	Entity State
	//  10	Repair Response
	//  11	Create Entity
	//  12	Remove Entity
	//  129	Announce Object
	//  13	Start/Resume
	//  130	Delete Object
	//  131	Describe Application
	//  132	Describe Event
	//  133	Describe Object
	//  134	Request Event
	//  135	Request Object
	//  14	Stop/Freeze
	//  140	Time Space Position Indicator - FI
	//  141	Appearance-FI
	//  142	Articulated Parts - FI
	//  143	Fire - FI
	//  144	Detonation - FI
	//  15	Acknowledge
	//  150	Point Object State
	//  151	Linear Object State
	//  152	Areal Object State
	//  153	Environment
	//  155	Transfer Control Request
	//  156	Transfer Control
	//  157	Transfer Control Acknowledge
	//  16	Action Request
	//  160	Intercom Control
	//  161	Intercom Signal
	//  17	Action Response
	//  170	Aggregate
	//  18	Data Query
	//  19	Set Data
	//  2	Fire
	//  20	Data
	//  21	Event Report
	//  22	Comment
	//  23	Electromagnetic Emission
	//  24	Designator
	//  25	Transmitter
	//  26	Signal
	//  27	Receiver
	//  3	Detonation
	//  4	Collision
	//  5	Service Request
	//  6	Resupply Offer
	//  7	Resupply Received
	//  8	Resupply Cancel
	//  9	Repair Complete
	PUDType         uint8
	PDUStatus       uint8
	PDUHeaderLength uint16
	Timestamp       uint64
}
type PDUPacket interface {
	TrimHeader() []byte
	TrimBody() []byte
}

type PDUBody struct {
	EntityIDRecord
	SequenceNumber uint8
	EntityType
	EntityAppearance   uint
	EntityCapabilities uint
	EntityLocation
	EntityOrientation
	EntityOrientationFixed uint
	EntityMarking          uint8
	EntityMarkingStringRecord
	EntityMarkingFixed               uint
	EntityMarkingForceID             uint
	EntityMarkingDRA                 uint
	EntityMarkingNumExtensionRecords uint
}

//
//	EntityIDRecord
//
//	Each Entity in a given exercise executing on a DIS application shall be assigned an Entity Identifier Record Unique to the exercise.
//
//	Item Name						Bit Length		Opt		Opt Ctl		Rpt		Rpt Ctl
//	Simulation Address Record		32
//	Entity Identity Field			16
//
//
//	Simulation Address Record
//
//	Item Name						Bit Length	Opt	Opt 	Ctl	Rpt		Rpt Ctl
//	Site Identifier Field			16
//	Application Identifier Field	16
//
//
type EntityIDRecord struct {

	//
	//	SiteIdentifierField
	//
	//	Each DIS site shall be assigned a unique Site Identifier.
	//	No site shall be assigned an ID contain all zero's, (2e16-1) or (2e16-2).
	//	The mechanism by which Site IDs are assigned is outside the scope of this standard.
	//  The simulation manager shall use the reserved site IDs to identify receivers of Simulation Management PDUs.
	//  A site ID equal to zero shall mean no site; this may be used to annotate a PDU log.
	//  A site ID equal to all ones (2e16-1) shall mean all sites; this may be used to start all sites.
	//  An application ID equal to (2e16-2) shall have no meaning but is reserved for consistency.
	//
	SiteIdentifierField uint16

	//
	//	ApplicationIdentifierField
	//
	//  Each simulation application at a DIS site shall be assigned an application identifier unique within that site.
	//  No simulation application shall be assigned an id contain all zeros, (2e16-1), or (2e16-2).
	//  One or more simulation applications may reside in a single host computer.
	//  The mechanism by which application IDs are assigned is outside the scope of this standard.
	//  The simulation manager shall use the reserved application IDs to identify receivers of Simulation Management PDUs.
	//  An application ID equal to zero shall mean no application.
	//  An application ID equal to all ones (2e16-1) shall mean all applications; this may be used to start all applications within a site.
	//  An application ID equal to (2e16-2) shall have no meaning but is reserved for consistency.
	ApplicationIdentifierField uint16

	//
	// EntityIDField
	//
	// Each entity in a given DIS application shall be given an Entity Identifier (Entity ID) unique to all other entities in that application and in that exercise.
	// This identifier is valid for the duration of the exercise; however, entity IDs shall be reused when all possible entity IDs have been exhausted.
	// No entity shall have an ID of zero, (2e16 -1) or (2e16 -2).
	// This number need not be registered or retained for future exercises.
	EntityIDField uint16
}

type PDUExtensionRecord struct {
	RecordType uint
	//	RecordType == 2001 means Attached Parts extension record
	//	RecordType == 2010 means DRA 10 extension record
	RecordTypeFixed uint
	RecordLength    uint
	RecordNumParts  uint
	DRA10Flags      uint
	AttachParts     uint
}

type EntityType struct {
	KindField        uint8
	DomainField      uint8
	CountryField     uint16
	CategoryField    uint8
	SubcategoryField uint8
	SpecificField    uint8
	ExtraField       uint8
}

type EntityLocation struct {
	XCoordinateField uint64
	YCoordinateField uint64
	ZCoordinateField uint64
}

type EntityOrientation struct {
	PSIField   uint32
	THETAField uint32
	PHIField   uint32
}

type EntityMarkingStringRecord struct {
	FirstCharacterField    uint8
	SecondCharacterField   uint8
	ThirdCharacterField    uint8
	FourthCharacterField   uint8
	FifthCharacterField    uint8
	SixthCharacterField    uint8
	SeventhCharacterField  uint8
	EighthCharacterField   uint8
	NinthCharacterField    uint8
	TenthCharacterField    uint8
	EleventhCharacterField uint8
}

func (header PDUHeader) TrimHeader() []byte {
	byteHeader := new(bytes.Buffer)
	_ = binary.Write(byteHeader, binary.LittleEndian, header)
	return byteHeader.Bytes()
}
func (body PDUBody) TrimBody() []byte {
	byteHeaderEntityIDRecord := new(bytes.Buffer)
	_ = binary.Write(byteHeaderEntityIDRecord, binary.LittleEndian, body.EntityIDRecord)
	PDUBodyEntityIDRecord := byteHeaderEntityIDRecord.Bytes()
	byteHeaderEntityType := new(bytes.Buffer)
	_ = binary.Write(byteHeaderEntityType, binary.LittleEndian, body.EntityType)
	PDUBodyEntityType := byteHeaderEntityIDRecord.Bytes()
	byteHeaderEntityOrientation := new(bytes.Buffer)
	_ = binary.Write(byteHeaderEntityOrientation, binary.LittleEndian, body.EntityOrientation)
	PDUBodyEntityOrientation := byteHeaderEntityIDRecord.Bytes()
	byteHeaderEntityLocation := new(bytes.Buffer)
	_ = binary.Write(byteHeaderEntityLocation, binary.LittleEndian, body.EntityLocation)
	PDUBodyEntityLocation := byteHeaderEntityIDRecord.Bytes()
	byteHeaderEntityMarkingStringRecord := new(bytes.Buffer)
	_ = binary.Write(byteHeaderEntityMarkingStringRecord, binary.LittleEndian, body.EntityMarkingStringRecord)
	PDUBodyEntityMarkingStringRecord := byteHeaderEntityIDRecord.Bytes()
	return join(append(join(append(join(append(PDUBodyEntityIDRecord, body.SequenceNumber), PDUBodyEntityType), PDUBodyEntityLocation...), PDUBodyEntityOrientation), body.EntityMarking), PDUBodyEntityMarkingStringRecord)

}

func join(slice []byte, slices []byte) []byte {
	return append(slice, slices...)
}
