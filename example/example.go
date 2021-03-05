package example

import (
	"../model"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {

	header := model.PDUHeader{
		ProtocolVersion:      1,
		CompatibilityVersion: 1,
		ExerciseID:           1,
		PUDType:              13,
		PDUStatus:            1,
		PDUHeaderLength:      31,
		Timestamp:            uint64(time.Now().Unix()),
	}
	body := model.PDUBody{
		EntityIDRecord: model.EntityIDRecord{
			SiteIdentifierField:        1,
			ApplicationIdentifierField: 1,
			EntityIDField:              1,
		},
		SequenceNumber: 1,
		EntityType: model.EntityType{
			KindField:        1,
			DomainField:      1,
			CountryField:     1,
			CategoryField:    1,
			SubcategoryField: 1,
			SpecificField:    1,
			ExtraField:       1,
		},
		EntityAppearance:   1,
		EntityCapabilities: 1,
		EntityLocation: model.EntityLocation{
			XCoordinateField: 1,
			YCoordinateField: 1,
			ZCoordinateField: 1,
		},
		EntityOrientation: model.EntityOrientation{
			PSIField:   1,
			THETAField: 1,
			PHIField:   1,
		},
		EntityOrientationFixed: 1,
		EntityMarking:          1,
		EntityMarkingStringRecord: model.EntityMarkingStringRecord{
			FirstCharacterField:    1,
			SecondCharacterField:   1,
			ThirdCharacterField:    1,
			FourthCharacterField:   1,
			FifthCharacterField:    1,
			SixthCharacterField:    1,
			SeventhCharacterField:  1,
			EighthCharacterField:   1,
			NinthCharacterField:    1,
			TenthCharacterField:    1,
			EleventhCharacterField: 1,
		},
		EntityMarkingFixed:               1,
		EntityMarkingForceID:             1,
		EntityMarkingDRA:                 1,
		EntityMarkingNumExtensionRecords: 1,
	}

	fmt.Println("Hex dump of go packet serialization output:")
	fmt.Println(hex.Dump(append(model.PDUHeader.TrimHeader(header),
		model.PDUBody.TrimBody(body)...)))
}
