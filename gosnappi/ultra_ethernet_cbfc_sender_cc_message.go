package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcSenderCcMessage *****
type ultraEthernetCbfcSenderCcMessage struct {
	validation
	obj          *otg.UltraEthernetCbfcSenderCcMessage
	marshaller   marshalUltraEthernetCbfcSenderCcMessage
	unMarshaller unMarshalUltraEthernetCbfcSenderCcMessage
}

func NewUltraEthernetCbfcSenderCcMessage() UltraEthernetCbfcSenderCcMessage {
	obj := ultraEthernetCbfcSenderCcMessage{obj: &otg.UltraEthernetCbfcSenderCcMessage{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcSenderCcMessage) msg() *otg.UltraEthernetCbfcSenderCcMessage {
	return obj.obj
}

func (obj *ultraEthernetCbfcSenderCcMessage) setMsg(msg *otg.UltraEthernetCbfcSenderCcMessage) UltraEthernetCbfcSenderCcMessage {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcSenderCcMessage struct {
	obj *ultraEthernetCbfcSenderCcMessage
}

type marshalUltraEthernetCbfcSenderCcMessage interface {
	// ToProto marshals UltraEthernetCbfcSenderCcMessage to protobuf object *otg.UltraEthernetCbfcSenderCcMessage
	ToProto() (*otg.UltraEthernetCbfcSenderCcMessage, error)
	// ToPbText marshals UltraEthernetCbfcSenderCcMessage to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcSenderCcMessage to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcSenderCcMessage to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcSenderCcMessage struct {
	obj *ultraEthernetCbfcSenderCcMessage
}

type unMarshalUltraEthernetCbfcSenderCcMessage interface {
	// FromProto unmarshals UltraEthernetCbfcSenderCcMessage from protobuf object *otg.UltraEthernetCbfcSenderCcMessage
	FromProto(msg *otg.UltraEthernetCbfcSenderCcMessage) (UltraEthernetCbfcSenderCcMessage, error)
	// FromPbText unmarshals UltraEthernetCbfcSenderCcMessage from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcSenderCcMessage from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcSenderCcMessage from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcSenderCcMessage) Marshal() marshalUltraEthernetCbfcSenderCcMessage {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcSenderCcMessage{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcSenderCcMessage) Unmarshal() unMarshalUltraEthernetCbfcSenderCcMessage {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcSenderCcMessage{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcSenderCcMessage) ToProto() (*otg.UltraEthernetCbfcSenderCcMessage, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcSenderCcMessage) FromProto(msg *otg.UltraEthernetCbfcSenderCcMessage) (UltraEthernetCbfcSenderCcMessage, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcSenderCcMessage) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshalultraEthernetCbfcSenderCcMessage) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalultraEthernetCbfcSenderCcMessage) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalultraEthernetCbfcSenderCcMessage) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalultraEthernetCbfcSenderCcMessage) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalultraEthernetCbfcSenderCcMessage) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ultraEthernetCbfcSenderCcMessage) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcSenderCcMessage) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcSenderCcMessage) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcSenderCcMessage) Clone() (UltraEthernetCbfcSenderCcMessage, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcSenderCcMessage()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// UltraEthernetCbfcSenderCcMessage is cBFC CC_Update (credit consumed) message configuration. CC_Update messages are
// periodically generated by the sender to recover credits that may have leaked
// due to packet drops from link errors.
type UltraEthernetCbfcSenderCcMessage interface {
	Validation
	// msg marshals UltraEthernetCbfcSenderCcMessage to protobuf object *otg.UltraEthernetCbfcSenderCcMessage
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcSenderCcMessage
	// setMsg unmarshals UltraEthernetCbfcSenderCcMessage from protobuf object *otg.UltraEthernetCbfcSenderCcMessage
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcSenderCcMessage) UltraEthernetCbfcSenderCcMessage
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcSenderCcMessage
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcSenderCcMessage
	// validate validates UltraEthernetCbfcSenderCcMessage
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcSenderCcMessage, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnableType1Message returns bool, set in UltraEthernetCbfcSenderCcMessage.
	EnableType1Message() bool
	// SetEnableType1Message assigns bool provided by user to UltraEthernetCbfcSenderCcMessage
	SetEnableType1Message(value bool) UltraEthernetCbfcSenderCcMessage
	// HasEnableType1Message checks if EnableType1Message has been set in UltraEthernetCbfcSenderCcMessage
	HasEnableType1Message() bool
	// EnableType2Message returns bool, set in UltraEthernetCbfcSenderCcMessage.
	EnableType2Message() bool
	// SetEnableType2Message assigns bool provided by user to UltraEthernetCbfcSenderCcMessage
	SetEnableType2Message(value bool) UltraEthernetCbfcSenderCcMessage
	// HasEnableType2Message checks if EnableType2Message has been set in UltraEthernetCbfcSenderCcMessage
	HasEnableType2Message() bool
	// Timer returns uint32, set in UltraEthernetCbfcSenderCcMessage.
	Timer() uint32
	// SetTimer assigns uint32 provided by user to UltraEthernetCbfcSenderCcMessage
	SetTimer(value uint32) UltraEthernetCbfcSenderCcMessage
	// HasTimer checks if Timer has been set in UltraEthernetCbfcSenderCcMessage
	HasTimer() bool
	// CompanyId returns string, set in UltraEthernetCbfcSenderCcMessage.
	CompanyId() string
	// SetCompanyId assigns string provided by user to UltraEthernetCbfcSenderCcMessage
	SetCompanyId(value string) UltraEthernetCbfcSenderCcMessage
	// HasCompanyId checks if CompanyId has been set in UltraEthernetCbfcSenderCcMessage
	HasCompanyId() bool
	// SourceAddress returns string, set in UltraEthernetCbfcSenderCcMessage.
	SourceAddress() string
	// SetSourceAddress assigns string provided by user to UltraEthernetCbfcSenderCcMessage
	SetSourceAddress(value string) UltraEthernetCbfcSenderCcMessage
	// HasSourceAddress checks if SourceAddress has been set in UltraEthernetCbfcSenderCcMessage
	HasSourceAddress() bool
	// DestinationAddress returns string, set in UltraEthernetCbfcSenderCcMessage.
	DestinationAddress() string
	// SetDestinationAddress assigns string provided by user to UltraEthernetCbfcSenderCcMessage
	SetDestinationAddress(value string) UltraEthernetCbfcSenderCcMessage
	// HasDestinationAddress checks if DestinationAddress has been set in UltraEthernetCbfcSenderCcMessage
	HasDestinationAddress() bool
}

// Enable generation of the type 1 CC_Update message which carries the credits
// consumed counters for VC 0 through VC 15.
// EnableType1Message returns a bool
func (obj *ultraEthernetCbfcSenderCcMessage) EnableType1Message() bool {

	return *obj.obj.EnableType1Message

}

// Enable generation of the type 1 CC_Update message which carries the credits
// consumed counters for VC 0 through VC 15.
// EnableType1Message returns a bool
func (obj *ultraEthernetCbfcSenderCcMessage) HasEnableType1Message() bool {
	return obj.obj.EnableType1Message != nil
}

// Enable generation of the type 1 CC_Update message which carries the credits
// consumed counters for VC 0 through VC 15.
// SetEnableType1Message sets the bool value in the UltraEthernetCbfcSenderCcMessage object
func (obj *ultraEthernetCbfcSenderCcMessage) SetEnableType1Message(value bool) UltraEthernetCbfcSenderCcMessage {

	obj.obj.EnableType1Message = &value
	return obj
}

// Enable generation of the type 2 CC_Update message which carries the credits
// consumed counters for VC 16 through VC 31.
// EnableType2Message returns a bool
func (obj *ultraEthernetCbfcSenderCcMessage) EnableType2Message() bool {

	return *obj.obj.EnableType2Message

}

// Enable generation of the type 2 CC_Update message which carries the credits
// consumed counters for VC 16 through VC 31.
// EnableType2Message returns a bool
func (obj *ultraEthernetCbfcSenderCcMessage) HasEnableType2Message() bool {
	return obj.obj.EnableType2Message != nil
}

// Enable generation of the type 2 CC_Update message which carries the credits
// consumed counters for VC 16 through VC 31.
// SetEnableType2Message sets the bool value in the UltraEthernetCbfcSenderCcMessage object
func (obj *ultraEthernetCbfcSenderCcMessage) SetEnableType2Message(value bool) UltraEthernetCbfcSenderCcMessage {

	obj.obj.EnableType2Message = &value
	return obj
}

// The time interval, in microseconds, between generation of successive
// CC_Update messages.
// Timer returns a uint32
func (obj *ultraEthernetCbfcSenderCcMessage) Timer() uint32 {

	return *obj.obj.Timer

}

// The time interval, in microseconds, between generation of successive
// CC_Update messages.
// Timer returns a uint32
func (obj *ultraEthernetCbfcSenderCcMessage) HasTimer() bool {
	return obj.obj.Timer != nil
}

// The time interval, in microseconds, between generation of successive
// CC_Update messages.
// SetTimer sets the uint32 value in the UltraEthernetCbfcSenderCcMessage object
func (obj *ultraEthernetCbfcSenderCcMessage) SetTimer(value uint32) UltraEthernetCbfcSenderCcMessage {

	obj.obj.Timer = &value
	return obj
}

// The 24-bit UEC Company ID (CID) placed in the CC_Update message, as a
// 6 character hexadecimal string. The UEC CID is FA7ACB.
// CompanyId returns a string
func (obj *ultraEthernetCbfcSenderCcMessage) CompanyId() string {

	return *obj.obj.CompanyId

}

// The 24-bit UEC Company ID (CID) placed in the CC_Update message, as a
// 6 character hexadecimal string. The UEC CID is FA7ACB.
// CompanyId returns a string
func (obj *ultraEthernetCbfcSenderCcMessage) HasCompanyId() bool {
	return obj.obj.CompanyId != nil
}

// The 24-bit UEC Company ID (CID) placed in the CC_Update message, as a
// 6 character hexadecimal string. The UEC CID is FA7ACB.
// SetCompanyId sets the string value in the UltraEthernetCbfcSenderCcMessage object
func (obj *ultraEthernetCbfcSenderCcMessage) SetCompanyId(value string) UltraEthernetCbfcSenderCcMessage {

	obj.obj.CompanyId = &value
	return obj
}

// The MAC source address of the CC_Update message.
// SourceAddress returns a string
func (obj *ultraEthernetCbfcSenderCcMessage) SourceAddress() string {

	return *obj.obj.SourceAddress

}

// The MAC source address of the CC_Update message.
// SourceAddress returns a string
func (obj *ultraEthernetCbfcSenderCcMessage) HasSourceAddress() bool {
	return obj.obj.SourceAddress != nil
}

// The MAC source address of the CC_Update message.
// SetSourceAddress sets the string value in the UltraEthernetCbfcSenderCcMessage object
func (obj *ultraEthernetCbfcSenderCcMessage) SetSourceAddress(value string) UltraEthernetCbfcSenderCcMessage {

	obj.obj.SourceAddress = &value
	return obj
}

// The MAC destination address of the CC_Update message. The specification
// uses the reserved address 01:80:C2:00:00:01 or the individual address of
// the destination station.
// DestinationAddress returns a string
func (obj *ultraEthernetCbfcSenderCcMessage) DestinationAddress() string {

	return *obj.obj.DestinationAddress

}

// The MAC destination address of the CC_Update message. The specification
// uses the reserved address 01:80:C2:00:00:01 or the individual address of
// the destination station.
// DestinationAddress returns a string
func (obj *ultraEthernetCbfcSenderCcMessage) HasDestinationAddress() bool {
	return obj.obj.DestinationAddress != nil
}

// The MAC destination address of the CC_Update message. The specification
// uses the reserved address 01:80:C2:00:00:01 or the individual address of
// the destination station.
// SetDestinationAddress sets the string value in the UltraEthernetCbfcSenderCcMessage object
func (obj *ultraEthernetCbfcSenderCcMessage) SetDestinationAddress(value string) UltraEthernetCbfcSenderCcMessage {

	obj.obj.DestinationAddress = &value
	return obj
}

func (obj *ultraEthernetCbfcSenderCcMessage) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Timer != nil {

		if *obj.obj.Timer < 1 || *obj.obj.Timer > 25000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= UltraEthernetCbfcSenderCcMessage.Timer <= 25000 but Got %d", *obj.obj.Timer))
		}

	}

	if obj.obj.CompanyId != nil {

		if !regexp.MustCompile(`^[0-9a-fA-F]{6}$`).MatchString(*obj.obj.CompanyId) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"UltraEthernetCbfcSenderCcMessage.CompanyId should adhere to this regex pattern '%s', but Got %s", `^[0-9a-fA-F]{6}$`, *obj.obj.CompanyId))
		}

	}

	if obj.obj.SourceAddress != nil {

		err := obj.validateMac(obj.SourceAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on UltraEthernetCbfcSenderCcMessage.SourceAddress"))
		}

	}

	if obj.obj.DestinationAddress != nil {

		err := obj.validateMac(obj.DestinationAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on UltraEthernetCbfcSenderCcMessage.DestinationAddress"))
		}

	}

}

func (obj *ultraEthernetCbfcSenderCcMessage) setDefault() {
	if obj.obj.EnableType1Message == nil {
		obj.SetEnableType1Message(false)
	}
	if obj.obj.EnableType2Message == nil {
		obj.SetEnableType2Message(false)
	}
	if obj.obj.Timer == nil {
		obj.SetTimer(10)
	}
	if obj.obj.CompanyId == nil {
		obj.SetCompanyId("FA7ACB")
	}
	if obj.obj.SourceAddress == nil {
		obj.SetSourceAddress("00:00:00:00:00:01")
	}
	if obj.obj.DestinationAddress == nil {
		obj.SetDestinationAddress("01:80:C2:00:00:01")
	}

}
