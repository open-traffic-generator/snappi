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
	obj                   *otg.MkaBasicKeySourcePsk
	marshaller            marshalMkaBasicKeySourcePsk
	unMarshaller          unMarshalMkaBasicKeySourcePsk
	startOffsetTimeHolder MkaBasicStartOffsetTime
	endOffsetTimeHolder   MkaBasicEndOffsetTime
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
	obj.setNil()
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
	// ToJsonRaw marshals MkaBasicKeySourcePsk to raw JSON text
	ToJsonRaw() (string, error)
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
	m.obj.setNil()
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalmkaBasicKeySourcePsk) ToJsonRaw() (string, error) {
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
	return string(data), nil
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
	m.obj.setNil()
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

func (obj *mkaBasicKeySourcePsk) setNil() {
	obj.startOffsetTimeHolder = nil
	obj.endOffsetTimeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
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
	// StartOffsetTime returns MkaBasicStartOffsetTime, set in MkaBasicKeySourcePsk.
	// MkaBasicStartOffsetTime is key start offset time in HH:MM. This is relative to key chain start time.
	StartOffsetTime() MkaBasicStartOffsetTime
	// SetStartOffsetTime assigns MkaBasicStartOffsetTime provided by user to MkaBasicKeySourcePsk.
	// MkaBasicStartOffsetTime is key start offset time in HH:MM. This is relative to key chain start time.
	SetStartOffsetTime(value MkaBasicStartOffsetTime) MkaBasicKeySourcePsk
	// HasStartOffsetTime checks if StartOffsetTime has been set in MkaBasicKeySourcePsk
	HasStartOffsetTime() bool
	// EndOffsetTime returns MkaBasicEndOffsetTime, set in MkaBasicKeySourcePsk.
	// MkaBasicEndOffsetTime is key end offset time in HH:MM. This is relative to key chain start time. A value of "00:00" makes the key valid for lifetime.
	EndOffsetTime() MkaBasicEndOffsetTime
	// SetEndOffsetTime assigns MkaBasicEndOffsetTime provided by user to MkaBasicKeySourcePsk.
	// MkaBasicEndOffsetTime is key end offset time in HH:MM. This is relative to key chain start time. A value of "00:00" makes the key valid for lifetime.
	SetEndOffsetTime(value MkaBasicEndOffsetTime) MkaBasicKeySourcePsk
	// HasEndOffsetTime checks if EndOffsetTime has been set in MkaBasicKeySourcePsk
	HasEndOffsetTime() bool
	setNil()
}

// Connectivity association key(CAK) value. It can be 128 bits or 256 bits long depending on the chosen MKA key derivation function.
// CakValue returns a string
func (obj *mkaBasicKeySourcePsk) CakValue() string {

	return *obj.obj.CakValue

}

// Connectivity association key(CAK) value. It can be 128 bits or 256 bits long depending on the chosen MKA key derivation function.
// CakValue returns a string
func (obj *mkaBasicKeySourcePsk) HasCakValue() bool {
	return obj.obj.CakValue != nil
}

// Connectivity association key(CAK) value. It can be 128 bits or 256 bits long depending on the chosen MKA key derivation function.
// SetCakValue sets the string value in the MkaBasicKeySourcePsk object
func (obj *mkaBasicKeySourcePsk) SetCakValue(value string) MkaBasicKeySourcePsk {

	obj.obj.CakValue = &value
	return obj
}

// Connectivity association key(CAK) name.
// CakName returns a string
func (obj *mkaBasicKeySourcePsk) CakName() string {

	return *obj.obj.CakName

}

// Connectivity association key(CAK) name.
// CakName returns a string
func (obj *mkaBasicKeySourcePsk) HasCakName() bool {
	return obj.obj.CakName != nil
}

// Connectivity association key(CAK) name.
// SetCakName sets the string value in the MkaBasicKeySourcePsk object
func (obj *mkaBasicKeySourcePsk) SetCakName(value string) MkaBasicKeySourcePsk {

	obj.obj.CakName = &value
	return obj
}

// StartOffsetTime returns a MkaBasicStartOffsetTime
func (obj *mkaBasicKeySourcePsk) StartOffsetTime() MkaBasicStartOffsetTime {
	if obj.obj.StartOffsetTime == nil {
		obj.obj.StartOffsetTime = NewMkaBasicStartOffsetTime().msg()
	}
	if obj.startOffsetTimeHolder == nil {
		obj.startOffsetTimeHolder = &mkaBasicStartOffsetTime{obj: obj.obj.StartOffsetTime}
	}
	return obj.startOffsetTimeHolder
}

// StartOffsetTime returns a MkaBasicStartOffsetTime
func (obj *mkaBasicKeySourcePsk) HasStartOffsetTime() bool {
	return obj.obj.StartOffsetTime != nil
}

// SetStartOffsetTime sets the MkaBasicStartOffsetTime value in the MkaBasicKeySourcePsk object
func (obj *mkaBasicKeySourcePsk) SetStartOffsetTime(value MkaBasicStartOffsetTime) MkaBasicKeySourcePsk {

	obj.startOffsetTimeHolder = nil
	obj.obj.StartOffsetTime = value.msg()

	return obj
}

// EndOffsetTime returns a MkaBasicEndOffsetTime
func (obj *mkaBasicKeySourcePsk) EndOffsetTime() MkaBasicEndOffsetTime {
	if obj.obj.EndOffsetTime == nil {
		obj.obj.EndOffsetTime = NewMkaBasicEndOffsetTime().msg()
	}
	if obj.endOffsetTimeHolder == nil {
		obj.endOffsetTimeHolder = &mkaBasicEndOffsetTime{obj: obj.obj.EndOffsetTime}
	}
	return obj.endOffsetTimeHolder
}

// EndOffsetTime returns a MkaBasicEndOffsetTime
func (obj *mkaBasicKeySourcePsk) HasEndOffsetTime() bool {
	return obj.obj.EndOffsetTime != nil
}

// SetEndOffsetTime sets the MkaBasicEndOffsetTime value in the MkaBasicKeySourcePsk object
func (obj *mkaBasicKeySourcePsk) SetEndOffsetTime(value MkaBasicEndOffsetTime) MkaBasicKeySourcePsk {

	obj.endOffsetTimeHolder = nil
	obj.obj.EndOffsetTime = value.msg()

	return obj
}

func (obj *mkaBasicKeySourcePsk) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CakValue != nil {

		if len(*obj.obj.CakValue) < 1 || len(*obj.obj.CakValue) > 64 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of MkaBasicKeySourcePsk.CakValue <= 64 but Got %d",
					len(*obj.obj.CakValue)))
		}

	}

	if obj.obj.CakName != nil {

		if len(*obj.obj.CakName) < 1 || len(*obj.obj.CakName) > 64 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of MkaBasicKeySourcePsk.CakName <= 64 but Got %d",
					len(*obj.obj.CakName)))
		}

	}

	if obj.obj.StartOffsetTime != nil {

		obj.StartOffsetTime().validateObj(vObj, set_default)
	}

	if obj.obj.EndOffsetTime != nil {

		obj.EndOffsetTime().validateObj(vObj, set_default)
	}

}

func (obj *mkaBasicKeySourcePsk) setDefault() {
	if obj.obj.CakValue == nil {
		obj.SetCakValue("F123456789ABCDEF0123456789ABCDEF")
	}
	if obj.obj.CakName == nil {
		obj.SetCakName("F123456789ABCDEF0123456789ABCDEFF123456789ABCDEF0123456789ABCDEF")
	}

}
