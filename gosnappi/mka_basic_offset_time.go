package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicOffsetTime *****
type mkaBasicOffsetTime struct {
	validation
	obj          *otg.MkaBasicOffsetTime
	marshaller   marshalMkaBasicOffsetTime
	unMarshaller unMarshalMkaBasicOffsetTime
}

func NewMkaBasicOffsetTime() MkaBasicOffsetTime {
	obj := mkaBasicOffsetTime{obj: &otg.MkaBasicOffsetTime{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicOffsetTime) msg() *otg.MkaBasicOffsetTime {
	return obj.obj
}

func (obj *mkaBasicOffsetTime) setMsg(msg *otg.MkaBasicOffsetTime) MkaBasicOffsetTime {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicOffsetTime struct {
	obj *mkaBasicOffsetTime
}

type marshalMkaBasicOffsetTime interface {
	// ToProto marshals MkaBasicOffsetTime to protobuf object *otg.MkaBasicOffsetTime
	ToProto() (*otg.MkaBasicOffsetTime, error)
	// ToPbText marshals MkaBasicOffsetTime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicOffsetTime to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicOffsetTime to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicOffsetTime struct {
	obj *mkaBasicOffsetTime
}

type unMarshalMkaBasicOffsetTime interface {
	// FromProto unmarshals MkaBasicOffsetTime from protobuf object *otg.MkaBasicOffsetTime
	FromProto(msg *otg.MkaBasicOffsetTime) (MkaBasicOffsetTime, error)
	// FromPbText unmarshals MkaBasicOffsetTime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicOffsetTime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicOffsetTime from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicOffsetTime) Marshal() marshalMkaBasicOffsetTime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicOffsetTime{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicOffsetTime) Unmarshal() unMarshalMkaBasicOffsetTime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicOffsetTime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicOffsetTime) ToProto() (*otg.MkaBasicOffsetTime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicOffsetTime) FromProto(msg *otg.MkaBasicOffsetTime) (MkaBasicOffsetTime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicOffsetTime) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicOffsetTime) FromPbText(value string) error {
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

func (m *marshalmkaBasicOffsetTime) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicOffsetTime) FromYaml(value string) error {
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

func (m *marshalmkaBasicOffsetTime) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicOffsetTime) FromJson(value string) error {
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

func (obj *mkaBasicOffsetTime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicOffsetTime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicOffsetTime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicOffsetTime) Clone() (MkaBasicOffsetTime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicOffsetTime()
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

// MkaBasicOffsetTime is offset time.
type MkaBasicOffsetTime interface {
	Validation
	// msg marshals MkaBasicOffsetTime to protobuf object *otg.MkaBasicOffsetTime
	// and doesn't set defaults
	msg() *otg.MkaBasicOffsetTime
	// setMsg unmarshals MkaBasicOffsetTime from protobuf object *otg.MkaBasicOffsetTime
	// and doesn't set defaults
	setMsg(*otg.MkaBasicOffsetTime) MkaBasicOffsetTime
	// provides marshal interface
	Marshal() marshalMkaBasicOffsetTime
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicOffsetTime
	// validate validates MkaBasicOffsetTime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicOffsetTime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Hh returns uint32, set in MkaBasicOffsetTime.
	Hh() uint32
	// SetHh assigns uint32 provided by user to MkaBasicOffsetTime
	SetHh(value uint32) MkaBasicOffsetTime
	// HasHh checks if Hh has been set in MkaBasicOffsetTime
	HasHh() bool
	// Mm returns uint32, set in MkaBasicOffsetTime.
	Mm() uint32
	// SetMm assigns uint32 provided by user to MkaBasicOffsetTime
	SetMm(value uint32) MkaBasicOffsetTime
	// HasMm checks if Mm has been set in MkaBasicOffsetTime
	HasMm() bool
}

// Hours in HH format.
// Hh returns a uint32
func (obj *mkaBasicOffsetTime) Hh() uint32 {

	return *obj.obj.Hh

}

// Hours in HH format.
// Hh returns a uint32
func (obj *mkaBasicOffsetTime) HasHh() bool {
	return obj.obj.Hh != nil
}

// Hours in HH format.
// SetHh sets the uint32 value in the MkaBasicOffsetTime object
func (obj *mkaBasicOffsetTime) SetHh(value uint32) MkaBasicOffsetTime {

	obj.obj.Hh = &value
	return obj
}

// Minutes in MM format.
// Mm returns a uint32
func (obj *mkaBasicOffsetTime) Mm() uint32 {

	return *obj.obj.Mm

}

// Minutes in MM format.
// Mm returns a uint32
func (obj *mkaBasicOffsetTime) HasMm() bool {
	return obj.obj.Mm != nil
}

// Minutes in MM format.
// SetMm sets the uint32 value in the MkaBasicOffsetTime object
func (obj *mkaBasicOffsetTime) SetMm(value uint32) MkaBasicOffsetTime {

	obj.obj.Mm = &value
	return obj
}

func (obj *mkaBasicOffsetTime) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Hh != nil {

		if *obj.obj.Hh > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicOffsetTime.Hh <= 4294967295 but Got %d", *obj.obj.Hh))
		}

	}

	if obj.obj.Mm != nil {

		if *obj.obj.Mm > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicOffsetTime.Mm <= 4294967295 but Got %d", *obj.obj.Mm))
		}

	}

}

func (obj *mkaBasicOffsetTime) setDefault() {
	if obj.obj.Hh == nil {
		obj.SetHh(0)
	}
	if obj.obj.Mm == nil {
		obj.SetMm(0)
	}

}
