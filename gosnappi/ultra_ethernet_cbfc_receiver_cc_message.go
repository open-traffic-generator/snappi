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

// ***** UltraEthernetCbfcReceiverCcMessage *****
type ultraEthernetCbfcReceiverCcMessage struct {
	validation
	obj          *otg.UltraEthernetCbfcReceiverCcMessage
	marshaller   marshalUltraEthernetCbfcReceiverCcMessage
	unMarshaller unMarshalUltraEthernetCbfcReceiverCcMessage
}

func NewUltraEthernetCbfcReceiverCcMessage() UltraEthernetCbfcReceiverCcMessage {
	obj := ultraEthernetCbfcReceiverCcMessage{obj: &otg.UltraEthernetCbfcReceiverCcMessage{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcReceiverCcMessage) msg() *otg.UltraEthernetCbfcReceiverCcMessage {
	return obj.obj
}

func (obj *ultraEthernetCbfcReceiverCcMessage) setMsg(msg *otg.UltraEthernetCbfcReceiverCcMessage) UltraEthernetCbfcReceiverCcMessage {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcReceiverCcMessage struct {
	obj *ultraEthernetCbfcReceiverCcMessage
}

type marshalUltraEthernetCbfcReceiverCcMessage interface {
	// ToProto marshals UltraEthernetCbfcReceiverCcMessage to protobuf object *otg.UltraEthernetCbfcReceiverCcMessage
	ToProto() (*otg.UltraEthernetCbfcReceiverCcMessage, error)
	// ToPbText marshals UltraEthernetCbfcReceiverCcMessage to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcReceiverCcMessage to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcReceiverCcMessage to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcReceiverCcMessage struct {
	obj *ultraEthernetCbfcReceiverCcMessage
}

type unMarshalUltraEthernetCbfcReceiverCcMessage interface {
	// FromProto unmarshals UltraEthernetCbfcReceiverCcMessage from protobuf object *otg.UltraEthernetCbfcReceiverCcMessage
	FromProto(msg *otg.UltraEthernetCbfcReceiverCcMessage) (UltraEthernetCbfcReceiverCcMessage, error)
	// FromPbText unmarshals UltraEthernetCbfcReceiverCcMessage from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcReceiverCcMessage from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcReceiverCcMessage from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcReceiverCcMessage) Marshal() marshalUltraEthernetCbfcReceiverCcMessage {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcReceiverCcMessage{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcReceiverCcMessage) Unmarshal() unMarshalUltraEthernetCbfcReceiverCcMessage {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcReceiverCcMessage{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcReceiverCcMessage) ToProto() (*otg.UltraEthernetCbfcReceiverCcMessage, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcReceiverCcMessage) FromProto(msg *otg.UltraEthernetCbfcReceiverCcMessage) (UltraEthernetCbfcReceiverCcMessage, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcReceiverCcMessage) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcReceiverCcMessage) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcReceiverCcMessage) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcReceiverCcMessage) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcReceiverCcMessage) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcReceiverCcMessage) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcReceiverCcMessage) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcReceiverCcMessage) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcReceiverCcMessage) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcReceiverCcMessage) Clone() (UltraEthernetCbfcReceiverCcMessage, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcReceiverCcMessage()
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

// UltraEthernetCbfcReceiverCcMessage is cBFC CC_Update (credit consumed) message reception configuration.
type UltraEthernetCbfcReceiverCcMessage interface {
	Validation
	// msg marshals UltraEthernetCbfcReceiverCcMessage to protobuf object *otg.UltraEthernetCbfcReceiverCcMessage
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcReceiverCcMessage
	// setMsg unmarshals UltraEthernetCbfcReceiverCcMessage from protobuf object *otg.UltraEthernetCbfcReceiverCcMessage
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcReceiverCcMessage) UltraEthernetCbfcReceiverCcMessage
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcReceiverCcMessage
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcReceiverCcMessage
	// validate validates UltraEthernetCbfcReceiverCcMessage
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcReceiverCcMessage, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ExpectedCompanyId returns string, set in UltraEthernetCbfcReceiverCcMessage.
	ExpectedCompanyId() string
	// SetExpectedCompanyId assigns string provided by user to UltraEthernetCbfcReceiverCcMessage
	SetExpectedCompanyId(value string) UltraEthernetCbfcReceiverCcMessage
	// HasExpectedCompanyId checks if ExpectedCompanyId has been set in UltraEthernetCbfcReceiverCcMessage
	HasExpectedCompanyId() bool
	// ValidateCcMinSpacing returns bool, set in UltraEthernetCbfcReceiverCcMessage.
	ValidateCcMinSpacing() bool
	// SetValidateCcMinSpacing assigns bool provided by user to UltraEthernetCbfcReceiverCcMessage
	SetValidateCcMinSpacing(value bool) UltraEthernetCbfcReceiverCcMessage
	// HasValidateCcMinSpacing checks if ValidateCcMinSpacing has been set in UltraEthernetCbfcReceiverCcMessage
	HasValidateCcMinSpacing() bool
	// CcMinSpacing returns uint32, set in UltraEthernetCbfcReceiverCcMessage.
	CcMinSpacing() uint32
	// SetCcMinSpacing assigns uint32 provided by user to UltraEthernetCbfcReceiverCcMessage
	SetCcMinSpacing(value uint32) UltraEthernetCbfcReceiverCcMessage
	// HasCcMinSpacing checks if CcMinSpacing has been set in UltraEthernetCbfcReceiverCcMessage
	HasCcMinSpacing() bool
	// ValidateCcMaxSpacing returns bool, set in UltraEthernetCbfcReceiverCcMessage.
	ValidateCcMaxSpacing() bool
	// SetValidateCcMaxSpacing assigns bool provided by user to UltraEthernetCbfcReceiverCcMessage
	SetValidateCcMaxSpacing(value bool) UltraEthernetCbfcReceiverCcMessage
	// HasValidateCcMaxSpacing checks if ValidateCcMaxSpacing has been set in UltraEthernetCbfcReceiverCcMessage
	HasValidateCcMaxSpacing() bool
	// CcMaxSpacing returns uint32, set in UltraEthernetCbfcReceiverCcMessage.
	CcMaxSpacing() uint32
	// SetCcMaxSpacing assigns uint32 provided by user to UltraEthernetCbfcReceiverCcMessage
	SetCcMaxSpacing(value uint32) UltraEthernetCbfcReceiverCcMessage
	// HasCcMaxSpacing checks if CcMaxSpacing has been set in UltraEthernetCbfcReceiverCcMessage
	HasCcMaxSpacing() bool
}

// The 24-bit UEC Company ID (CID) expected in received CC_Update messages, as
// a 6 character hexadecimal string. The UEC CID is FA7ACB.
// ExpectedCompanyId returns a string
func (obj *ultraEthernetCbfcReceiverCcMessage) ExpectedCompanyId() string {

	return *obj.obj.ExpectedCompanyId

}

// The 24-bit UEC Company ID (CID) expected in received CC_Update messages, as
// a 6 character hexadecimal string. The UEC CID is FA7ACB.
// ExpectedCompanyId returns a string
func (obj *ultraEthernetCbfcReceiverCcMessage) HasExpectedCompanyId() bool {
	return obj.obj.ExpectedCompanyId != nil
}

// The 24-bit UEC Company ID (CID) expected in received CC_Update messages, as
// a 6 character hexadecimal string. The UEC CID is FA7ACB.
// SetExpectedCompanyId sets the string value in the UltraEthernetCbfcReceiverCcMessage object
func (obj *ultraEthernetCbfcReceiverCcMessage) SetExpectedCompanyId(value string) UltraEthernetCbfcReceiverCcMessage {

	obj.obj.ExpectedCompanyId = &value
	return obj
}

// Enable validation of the minimum spacing between received CC_Update
// messages.
// ValidateCcMinSpacing returns a bool
func (obj *ultraEthernetCbfcReceiverCcMessage) ValidateCcMinSpacing() bool {

	return *obj.obj.ValidateCcMinSpacing

}

// Enable validation of the minimum spacing between received CC_Update
// messages.
// ValidateCcMinSpacing returns a bool
func (obj *ultraEthernetCbfcReceiverCcMessage) HasValidateCcMinSpacing() bool {
	return obj.obj.ValidateCcMinSpacing != nil
}

// Enable validation of the minimum spacing between received CC_Update
// messages.
// SetValidateCcMinSpacing sets the bool value in the UltraEthernetCbfcReceiverCcMessage object
func (obj *ultraEthernetCbfcReceiverCcMessage) SetValidateCcMinSpacing(value bool) UltraEthernetCbfcReceiverCcMessage {

	obj.obj.ValidateCcMinSpacing = &value
	return obj
}

// The minimum spacing, in microseconds, expected between received CC_Update
// messages.
// CcMinSpacing returns a uint32
func (obj *ultraEthernetCbfcReceiverCcMessage) CcMinSpacing() uint32 {

	return *obj.obj.CcMinSpacing

}

// The minimum spacing, in microseconds, expected between received CC_Update
// messages.
// CcMinSpacing returns a uint32
func (obj *ultraEthernetCbfcReceiverCcMessage) HasCcMinSpacing() bool {
	return obj.obj.CcMinSpacing != nil
}

// The minimum spacing, in microseconds, expected between received CC_Update
// messages.
// SetCcMinSpacing sets the uint32 value in the UltraEthernetCbfcReceiverCcMessage object
func (obj *ultraEthernetCbfcReceiverCcMessage) SetCcMinSpacing(value uint32) UltraEthernetCbfcReceiverCcMessage {

	obj.obj.CcMinSpacing = &value
	return obj
}

// Enable validation of the maximum spacing between received CC_Update
// messages.
// ValidateCcMaxSpacing returns a bool
func (obj *ultraEthernetCbfcReceiverCcMessage) ValidateCcMaxSpacing() bool {

	return *obj.obj.ValidateCcMaxSpacing

}

// Enable validation of the maximum spacing between received CC_Update
// messages.
// ValidateCcMaxSpacing returns a bool
func (obj *ultraEthernetCbfcReceiverCcMessage) HasValidateCcMaxSpacing() bool {
	return obj.obj.ValidateCcMaxSpacing != nil
}

// Enable validation of the maximum spacing between received CC_Update
// messages.
// SetValidateCcMaxSpacing sets the bool value in the UltraEthernetCbfcReceiverCcMessage object
func (obj *ultraEthernetCbfcReceiverCcMessage) SetValidateCcMaxSpacing(value bool) UltraEthernetCbfcReceiverCcMessage {

	obj.obj.ValidateCcMaxSpacing = &value
	return obj
}

// The maximum spacing, in microseconds, expected between received CC_Update
// messages.
// CcMaxSpacing returns a uint32
func (obj *ultraEthernetCbfcReceiverCcMessage) CcMaxSpacing() uint32 {

	return *obj.obj.CcMaxSpacing

}

// The maximum spacing, in microseconds, expected between received CC_Update
// messages.
// CcMaxSpacing returns a uint32
func (obj *ultraEthernetCbfcReceiverCcMessage) HasCcMaxSpacing() bool {
	return obj.obj.CcMaxSpacing != nil
}

// The maximum spacing, in microseconds, expected between received CC_Update
// messages.
// SetCcMaxSpacing sets the uint32 value in the UltraEthernetCbfcReceiverCcMessage object
func (obj *ultraEthernetCbfcReceiverCcMessage) SetCcMaxSpacing(value uint32) UltraEthernetCbfcReceiverCcMessage {

	obj.obj.CcMaxSpacing = &value
	return obj
}

func (obj *ultraEthernetCbfcReceiverCcMessage) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ExpectedCompanyId != nil {

		if !regexp.MustCompile(`^[0-9a-fA-F]{6}$`).MatchString(*obj.obj.ExpectedCompanyId) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"UltraEthernetCbfcReceiverCcMessage.ExpectedCompanyId should adhere to this regex pattern '%s', but Got %s", `^[0-9a-fA-F]{6}$`, *obj.obj.ExpectedCompanyId))
		}

	}

	if obj.obj.CcMinSpacing != nil {

		if *obj.obj.CcMinSpacing < 1 || *obj.obj.CcMinSpacing > 25000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= UltraEthernetCbfcReceiverCcMessage.CcMinSpacing <= 25000 but Got %d", *obj.obj.CcMinSpacing))
		}

	}

	if obj.obj.CcMaxSpacing != nil {

		if *obj.obj.CcMaxSpacing < 1 || *obj.obj.CcMaxSpacing > 25000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= UltraEthernetCbfcReceiverCcMessage.CcMaxSpacing <= 25000 but Got %d", *obj.obj.CcMaxSpacing))
		}

	}

}

func (obj *ultraEthernetCbfcReceiverCcMessage) setDefault() {
	if obj.obj.ExpectedCompanyId == nil {
		obj.SetExpectedCompanyId("FA7ACB")
	}
	if obj.obj.ValidateCcMinSpacing == nil {
		obj.SetValidateCcMinSpacing(false)
	}
	if obj.obj.CcMinSpacing == nil {
		obj.SetCcMinSpacing(1)
	}
	if obj.obj.ValidateCcMaxSpacing == nil {
		obj.SetValidateCcMaxSpacing(false)
	}
	if obj.obj.CcMaxSpacing == nil {
		obj.SetCcMaxSpacing(25000)
	}

}
