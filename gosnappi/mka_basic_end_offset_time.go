package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicEndOffsetTime *****
type mkaBasicEndOffsetTime struct {
	validation
	obj          *otg.MkaBasicEndOffsetTime
	marshaller   marshalMkaBasicEndOffsetTime
	unMarshaller unMarshalMkaBasicEndOffsetTime
}

func NewMkaBasicEndOffsetTime() MkaBasicEndOffsetTime {
	obj := mkaBasicEndOffsetTime{obj: &otg.MkaBasicEndOffsetTime{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicEndOffsetTime) msg() *otg.MkaBasicEndOffsetTime {
	return obj.obj
}

func (obj *mkaBasicEndOffsetTime) setMsg(msg *otg.MkaBasicEndOffsetTime) MkaBasicEndOffsetTime {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicEndOffsetTime struct {
	obj *mkaBasicEndOffsetTime
}

type marshalMkaBasicEndOffsetTime interface {
	// ToProto marshals MkaBasicEndOffsetTime to protobuf object *otg.MkaBasicEndOffsetTime
	ToProto() (*otg.MkaBasicEndOffsetTime, error)
	// ToPbText marshals MkaBasicEndOffsetTime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicEndOffsetTime to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicEndOffsetTime to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicEndOffsetTime struct {
	obj *mkaBasicEndOffsetTime
}

type unMarshalMkaBasicEndOffsetTime interface {
	// FromProto unmarshals MkaBasicEndOffsetTime from protobuf object *otg.MkaBasicEndOffsetTime
	FromProto(msg *otg.MkaBasicEndOffsetTime) (MkaBasicEndOffsetTime, error)
	// FromPbText unmarshals MkaBasicEndOffsetTime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicEndOffsetTime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicEndOffsetTime from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicEndOffsetTime) Marshal() marshalMkaBasicEndOffsetTime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicEndOffsetTime{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicEndOffsetTime) Unmarshal() unMarshalMkaBasicEndOffsetTime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicEndOffsetTime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicEndOffsetTime) ToProto() (*otg.MkaBasicEndOffsetTime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicEndOffsetTime) FromProto(msg *otg.MkaBasicEndOffsetTime) (MkaBasicEndOffsetTime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicEndOffsetTime) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicEndOffsetTime) FromPbText(value string) error {
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

func (m *marshalmkaBasicEndOffsetTime) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicEndOffsetTime) FromYaml(value string) error {
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

func (m *marshalmkaBasicEndOffsetTime) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicEndOffsetTime) FromJson(value string) error {
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

func (obj *mkaBasicEndOffsetTime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicEndOffsetTime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicEndOffsetTime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicEndOffsetTime) Clone() (MkaBasicEndOffsetTime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicEndOffsetTime()
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

// MkaBasicEndOffsetTime is key end offset time in HH:MM. This is relative to key chain start time. A value of "00:00" makes the key valid for lifetime.
type MkaBasicEndOffsetTime interface {
	Validation
	// msg marshals MkaBasicEndOffsetTime to protobuf object *otg.MkaBasicEndOffsetTime
	// and doesn't set defaults
	msg() *otg.MkaBasicEndOffsetTime
	// setMsg unmarshals MkaBasicEndOffsetTime from protobuf object *otg.MkaBasicEndOffsetTime
	// and doesn't set defaults
	setMsg(*otg.MkaBasicEndOffsetTime) MkaBasicEndOffsetTime
	// provides marshal interface
	Marshal() marshalMkaBasicEndOffsetTime
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicEndOffsetTime
	// validate validates MkaBasicEndOffsetTime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicEndOffsetTime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Hh returns uint32, set in MkaBasicEndOffsetTime.
	Hh() uint32
	// SetHh assigns uint32 provided by user to MkaBasicEndOffsetTime
	SetHh(value uint32) MkaBasicEndOffsetTime
	// HasHh checks if Hh has been set in MkaBasicEndOffsetTime
	HasHh() bool
	// Mm returns uint32, set in MkaBasicEndOffsetTime.
	Mm() uint32
	// SetMm assigns uint32 provided by user to MkaBasicEndOffsetTime
	SetMm(value uint32) MkaBasicEndOffsetTime
	// HasMm checks if Mm has been set in MkaBasicEndOffsetTime
	HasMm() bool
}

// Hours in HH format.
// Hh returns a uint32
func (obj *mkaBasicEndOffsetTime) Hh() uint32 {

	return *obj.obj.Hh

}

// Hours in HH format.
// Hh returns a uint32
func (obj *mkaBasicEndOffsetTime) HasHh() bool {
	return obj.obj.Hh != nil
}

// Hours in HH format.
// SetHh sets the uint32 value in the MkaBasicEndOffsetTime object
func (obj *mkaBasicEndOffsetTime) SetHh(value uint32) MkaBasicEndOffsetTime {

	obj.obj.Hh = &value
	return obj
}

// Minutes in MM format.
// Mm returns a uint32
func (obj *mkaBasicEndOffsetTime) Mm() uint32 {

	return *obj.obj.Mm

}

// Minutes in MM format.
// Mm returns a uint32
func (obj *mkaBasicEndOffsetTime) HasMm() bool {
	return obj.obj.Mm != nil
}

// Minutes in MM format.
// SetMm sets the uint32 value in the MkaBasicEndOffsetTime object
func (obj *mkaBasicEndOffsetTime) SetMm(value uint32) MkaBasicEndOffsetTime {

	obj.obj.Mm = &value
	return obj
}

func (obj *mkaBasicEndOffsetTime) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Hh != nil {

		if *obj.obj.Hh > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicEndOffsetTime.Hh <= 4294967295 but Got %d", *obj.obj.Hh))
		}

	}

	if obj.obj.Mm != nil {

		if *obj.obj.Mm > 59 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicEndOffsetTime.Mm <= 59 but Got %d", *obj.obj.Mm))
		}

	}

}

func (obj *mkaBasicEndOffsetTime) setDefault() {
	if obj.obj.Hh == nil {
		obj.SetHh(0)
	}
	if obj.obj.Mm == nil {
		obj.SetMm(0)
	}

}
