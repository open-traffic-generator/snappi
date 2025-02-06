package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicKeySourcePsk *****
type mkaBasicKeySourcePsk struct {
	validation
	obj          *otg.MkaBasicKeySourcePsk
	marshaller   marshalMkaBasicKeySourcePsk
	unMarshaller unMarshalMkaBasicKeySourcePsk
}

func NewMkaBasicKeySourcePsk() MkaBasicKeySourcePsk {
	obj := mkaBasicKeySourcePsk{obj: &otg.MkaBasicKeySourcePsk{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicKeySourcePsk) msg() *otg.MkaBasicKeySourcePsk {
	return obj.obj
}

func (obj *mkaBasicKeySourcePsk) setMsg(msg *otg.MkaBasicKeySourcePsk) MkaBasicKeySourcePsk {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicKeySourcePsk struct {
	obj *mkaBasicKeySourcePsk
}

type marshalMkaBasicKeySourcePsk interface {
	// ToProto marshals MkaBasicKeySourcePsk to protobuf object *otg.MkaBasicKeySourcePsk
	ToProto() (*otg.MkaBasicKeySourcePsk, error)
	// ToPbText marshals MkaBasicKeySourcePsk to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicKeySourcePsk to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicKeySourcePsk to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicKeySourcePsk struct {
	obj *mkaBasicKeySourcePsk
}

type unMarshalMkaBasicKeySourcePsk interface {
	// FromProto unmarshals MkaBasicKeySourcePsk from protobuf object *otg.MkaBasicKeySourcePsk
	FromProto(msg *otg.MkaBasicKeySourcePsk) (MkaBasicKeySourcePsk, error)
	// FromPbText unmarshals MkaBasicKeySourcePsk from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicKeySourcePsk from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicKeySourcePsk from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicKeySourcePsk) Marshal() marshalMkaBasicKeySourcePsk {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicKeySourcePsk{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicKeySourcePsk) Unmarshal() unMarshalMkaBasicKeySourcePsk {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicKeySourcePsk{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicKeySourcePsk) ToProto() (*otg.MkaBasicKeySourcePsk, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicKeySourcePsk) FromProto(msg *otg.MkaBasicKeySourcePsk) (MkaBasicKeySourcePsk, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicKeySourcePsk) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicKeySourcePsk) FromPbText(value string) error {
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

func (m *marshalmkaBasicKeySourcePsk) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicKeySourcePsk) FromYaml(value string) error {
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

func (m *marshalmkaBasicKeySourcePsk) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicKeySourcePsk) FromJson(value string) error {
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

func (obj *mkaBasicKeySourcePsk) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicKeySourcePsk) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicKeySourcePsk) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicKeySourcePsk) Clone() (MkaBasicKeySourcePsk, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicKeySourcePsk()
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

// MkaBasicKeySourcePsk is the container for Pre-shared key(PSK).
type MkaBasicKeySourcePsk interface {
	Validation
	// msg marshals MkaBasicKeySourcePsk to protobuf object *otg.MkaBasicKeySourcePsk
	// and doesn't set defaults
	msg() *otg.MkaBasicKeySourcePsk
	// setMsg unmarshals MkaBasicKeySourcePsk from protobuf object *otg.MkaBasicKeySourcePsk
	// and doesn't set defaults
	setMsg(*otg.MkaBasicKeySourcePsk) MkaBasicKeySourcePsk
	// provides marshal interface
	Marshal() marshalMkaBasicKeySourcePsk
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicKeySourcePsk
	// validate validates MkaBasicKeySourcePsk
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicKeySourcePsk, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CakValue returns string, set in MkaBasicKeySourcePsk.
	CakValue() string
	// SetCakValue assigns string provided by user to MkaBasicKeySourcePsk
	SetCakValue(value string) MkaBasicKeySourcePsk
	// HasCakValue checks if CakValue has been set in MkaBasicKeySourcePsk
	HasCakValue() bool
	// CakName returns string, set in MkaBasicKeySourcePsk.
	CakName() string
	// SetCakName assigns string provided by user to MkaBasicKeySourcePsk
	SetCakName(value string) MkaBasicKeySourcePsk
	// HasCakName checks if CakName has been set in MkaBasicKeySourcePsk
	HasCakName() bool
	// StartTime returns string, set in MkaBasicKeySourcePsk.
	StartTime() string
	// SetStartTime assigns string provided by user to MkaBasicKeySourcePsk
	SetStartTime(value string) MkaBasicKeySourcePsk
	// HasStartTime checks if StartTime has been set in MkaBasicKeySourcePsk
	HasStartTime() bool
	// EndTime returns string, set in MkaBasicKeySourcePsk.
	EndTime() string
	// SetEndTime assigns string provided by user to MkaBasicKeySourcePsk
	SetEndTime(value string) MkaBasicKeySourcePsk
	// HasEndTime checks if EndTime has been set in MkaBasicKeySourcePsk
	HasEndTime() bool
}

// CAK value.
// CakValue returns a string
func (obj *mkaBasicKeySourcePsk) CakValue() string {

	return *obj.obj.CakValue

}

// CAK value.
// CakValue returns a string
func (obj *mkaBasicKeySourcePsk) HasCakValue() bool {
	return obj.obj.CakValue != nil
}

// CAK value.
// SetCakValue sets the string value in the MkaBasicKeySourcePsk object
func (obj *mkaBasicKeySourcePsk) SetCakValue(value string) MkaBasicKeySourcePsk {

	obj.obj.CakValue = &value
	return obj
}

// CAK name.
// CakName returns a string
func (obj *mkaBasicKeySourcePsk) CakName() string {

	return *obj.obj.CakName

}

// CAK name.
// CakName returns a string
func (obj *mkaBasicKeySourcePsk) HasCakName() bool {
	return obj.obj.CakName != nil
}

// CAK name.
// SetCakName sets the string value in the MkaBasicKeySourcePsk object
func (obj *mkaBasicKeySourcePsk) SetCakName(value string) MkaBasicKeySourcePsk {

	obj.obj.CakName = &value
	return obj
}

// Key start time in HH:MM.
// StartTime returns a string
func (obj *mkaBasicKeySourcePsk) StartTime() string {

	return *obj.obj.StartTime

}

// Key start time in HH:MM.
// StartTime returns a string
func (obj *mkaBasicKeySourcePsk) HasStartTime() bool {
	return obj.obj.StartTime != nil
}

// Key start time in HH:MM.
// SetStartTime sets the string value in the MkaBasicKeySourcePsk object
func (obj *mkaBasicKeySourcePsk) SetStartTime(value string) MkaBasicKeySourcePsk {

	obj.obj.StartTime = &value
	return obj
}

// Key end time in HH:MM.
// EndTime returns a string
func (obj *mkaBasicKeySourcePsk) EndTime() string {

	return *obj.obj.EndTime

}

// Key end time in HH:MM.
// EndTime returns a string
func (obj *mkaBasicKeySourcePsk) HasEndTime() bool {
	return obj.obj.EndTime != nil
}

// Key end time in HH:MM.
// SetEndTime sets the string value in the MkaBasicKeySourcePsk object
func (obj *mkaBasicKeySourcePsk) SetEndTime(value string) MkaBasicKeySourcePsk {

	obj.obj.EndTime = &value
	return obj
}

func (obj *mkaBasicKeySourcePsk) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CakValue != nil {

		if len(*obj.obj.CakValue) < 16 || len(*obj.obj.CakValue) > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"16 <= length of MkaBasicKeySourcePsk.CakValue <= 32 but Got %d",
					len(*obj.obj.CakValue)))
		}

	}

	if obj.obj.CakName != nil {

		if len(*obj.obj.CakName) < 1 || len(*obj.obj.CakName) > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of MkaBasicKeySourcePsk.CakName <= 32 but Got %d",
					len(*obj.obj.CakName)))
		}

	}

}

func (obj *mkaBasicKeySourcePsk) setDefault() {
	if obj.obj.CakValue == nil {
		obj.SetCakValue("F123456789ABCDEF0123456789ABCDEF")
	}
	if obj.obj.CakName == nil {
		obj.SetCakName("F123456789ABCDEF0123456789ABCDEFF123456789ABCDEF0123456789ABCDEF")
	}
	if obj.obj.StartTime == nil {
		obj.SetStartTime("00:00")
	}
	if obj.obj.EndTime == nil {
		obj.SetEndTime("00:00")
	}

}
