package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicStartOffsetTime *****
type mkaBasicStartOffsetTime struct {
	validation
	obj          *otg.MkaBasicStartOffsetTime
	marshaller   marshalMkaBasicStartOffsetTime
	unMarshaller unMarshalMkaBasicStartOffsetTime
}

func NewMkaBasicStartOffsetTime() MkaBasicStartOffsetTime {
	obj := mkaBasicStartOffsetTime{obj: &otg.MkaBasicStartOffsetTime{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicStartOffsetTime) msg() *otg.MkaBasicStartOffsetTime {
	return obj.obj
}

func (obj *mkaBasicStartOffsetTime) setMsg(msg *otg.MkaBasicStartOffsetTime) MkaBasicStartOffsetTime {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicStartOffsetTime struct {
	obj *mkaBasicStartOffsetTime
}

type marshalMkaBasicStartOffsetTime interface {
	// ToProto marshals MkaBasicStartOffsetTime to protobuf object *otg.MkaBasicStartOffsetTime
	ToProto() (*otg.MkaBasicStartOffsetTime, error)
	// ToPbText marshals MkaBasicStartOffsetTime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicStartOffsetTime to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicStartOffsetTime to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicStartOffsetTime struct {
	obj *mkaBasicStartOffsetTime
}

type unMarshalMkaBasicStartOffsetTime interface {
	// FromProto unmarshals MkaBasicStartOffsetTime from protobuf object *otg.MkaBasicStartOffsetTime
	FromProto(msg *otg.MkaBasicStartOffsetTime) (MkaBasicStartOffsetTime, error)
	// FromPbText unmarshals MkaBasicStartOffsetTime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicStartOffsetTime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicStartOffsetTime from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicStartOffsetTime) Marshal() marshalMkaBasicStartOffsetTime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicStartOffsetTime{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicStartOffsetTime) Unmarshal() unMarshalMkaBasicStartOffsetTime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicStartOffsetTime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicStartOffsetTime) ToProto() (*otg.MkaBasicStartOffsetTime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicStartOffsetTime) FromProto(msg *otg.MkaBasicStartOffsetTime) (MkaBasicStartOffsetTime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicStartOffsetTime) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicStartOffsetTime) FromPbText(value string) error {
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

func (m *marshalmkaBasicStartOffsetTime) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicStartOffsetTime) FromYaml(value string) error {
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

func (m *marshalmkaBasicStartOffsetTime) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicStartOffsetTime) FromJson(value string) error {
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

func (obj *mkaBasicStartOffsetTime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicStartOffsetTime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicStartOffsetTime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicStartOffsetTime) Clone() (MkaBasicStartOffsetTime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicStartOffsetTime()
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

// MkaBasicStartOffsetTime is key start offset time in HH:MM. This is relative to key chain start time.
type MkaBasicStartOffsetTime interface {
	Validation
	// msg marshals MkaBasicStartOffsetTime to protobuf object *otg.MkaBasicStartOffsetTime
	// and doesn't set defaults
	msg() *otg.MkaBasicStartOffsetTime
	// setMsg unmarshals MkaBasicStartOffsetTime from protobuf object *otg.MkaBasicStartOffsetTime
	// and doesn't set defaults
	setMsg(*otg.MkaBasicStartOffsetTime) MkaBasicStartOffsetTime
	// provides marshal interface
	Marshal() marshalMkaBasicStartOffsetTime
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicStartOffsetTime
	// validate validates MkaBasicStartOffsetTime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicStartOffsetTime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Hh returns uint32, set in MkaBasicStartOffsetTime.
	Hh() uint32
	// SetHh assigns uint32 provided by user to MkaBasicStartOffsetTime
	SetHh(value uint32) MkaBasicStartOffsetTime
	// HasHh checks if Hh has been set in MkaBasicStartOffsetTime
	HasHh() bool
	// Mm returns uint32, set in MkaBasicStartOffsetTime.
	Mm() uint32
	// SetMm assigns uint32 provided by user to MkaBasicStartOffsetTime
	SetMm(value uint32) MkaBasicStartOffsetTime
	// HasMm checks if Mm has been set in MkaBasicStartOffsetTime
	HasMm() bool
}

// Hours in HH format.
// Hh returns a uint32
func (obj *mkaBasicStartOffsetTime) Hh() uint32 {

	return *obj.obj.Hh

}

// Hours in HH format.
// Hh returns a uint32
func (obj *mkaBasicStartOffsetTime) HasHh() bool {
	return obj.obj.Hh != nil
}

// Hours in HH format.
// SetHh sets the uint32 value in the MkaBasicStartOffsetTime object
func (obj *mkaBasicStartOffsetTime) SetHh(value uint32) MkaBasicStartOffsetTime {

	obj.obj.Hh = &value
	return obj
}

// Minutes in MM format.
// Mm returns a uint32
func (obj *mkaBasicStartOffsetTime) Mm() uint32 {

	return *obj.obj.Mm

}

// Minutes in MM format.
// Mm returns a uint32
func (obj *mkaBasicStartOffsetTime) HasMm() bool {
	return obj.obj.Mm != nil
}

// Minutes in MM format.
// SetMm sets the uint32 value in the MkaBasicStartOffsetTime object
func (obj *mkaBasicStartOffsetTime) SetMm(value uint32) MkaBasicStartOffsetTime {

	obj.obj.Mm = &value
	return obj
}

func (obj *mkaBasicStartOffsetTime) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Hh != nil {

		if *obj.obj.Hh > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicStartOffsetTime.Hh <= 4294967295 but Got %d", *obj.obj.Hh))
		}

	}

	if obj.obj.Mm != nil {

		if *obj.obj.Mm > 59 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicStartOffsetTime.Mm <= 59 but Got %d", *obj.obj.Mm))
		}

	}

}

func (obj *mkaBasicStartOffsetTime) setDefault() {
	if obj.obj.Hh == nil {
		obj.SetHh(0)
	}
	if obj.obj.Mm == nil {
		obj.SetMm(0)
	}

}
