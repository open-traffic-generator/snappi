package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaTxSc *****
type mkaTxSc struct {
	validation
	obj          *otg.MkaTxSc
	marshaller   marshalMkaTxSc
	unMarshaller unMarshalMkaTxSc
}

func NewMkaTxSc() MkaTxSc {
	obj := mkaTxSc{obj: &otg.MkaTxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaTxSc) msg() *otg.MkaTxSc {
	return obj.obj
}

func (obj *mkaTxSc) setMsg(msg *otg.MkaTxSc) MkaTxSc {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaTxSc struct {
	obj *mkaTxSc
}

type marshalMkaTxSc interface {
	// ToProto marshals MkaTxSc to protobuf object *otg.MkaTxSc
	ToProto() (*otg.MkaTxSc, error)
	// ToPbText marshals MkaTxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaTxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaTxSc to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals MkaTxSc to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalmkaTxSc struct {
	obj *mkaTxSc
}

type unMarshalMkaTxSc interface {
	// FromProto unmarshals MkaTxSc from protobuf object *otg.MkaTxSc
	FromProto(msg *otg.MkaTxSc) (MkaTxSc, error)
	// FromPbText unmarshals MkaTxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaTxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaTxSc from JSON text
	FromJson(value string) error
}

func (obj *mkaTxSc) Marshal() marshalMkaTxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaTxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaTxSc) Unmarshal() unMarshalMkaTxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaTxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaTxSc) ToProto() (*otg.MkaTxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaTxSc) FromProto(msg *otg.MkaTxSc) (MkaTxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaTxSc) ToPbText() (string, error) {
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

func (m *unMarshalmkaTxSc) FromPbText(value string) error {
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

func (m *marshalmkaTxSc) ToYaml() (string, error) {
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

func (m *unMarshalmkaTxSc) FromYaml(value string) error {
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

func (m *marshalmkaTxSc) ToJsonRaw() (string, error) {
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

func (m *marshalmkaTxSc) ToJson() (string, error) {
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

func (m *unMarshalmkaTxSc) FromJson(value string) error {
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

func (obj *mkaTxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaTxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaTxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaTxSc) Clone() (MkaTxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaTxSc()
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

// MkaTxSc is tx secure channel(SC) properties.
type MkaTxSc interface {
	Validation
	// msg marshals MkaTxSc to protobuf object *otg.MkaTxSc
	// and doesn't set defaults
	msg() *otg.MkaTxSc
	// setMsg unmarshals MkaTxSc from protobuf object *otg.MkaTxSc
	// and doesn't set defaults
	setMsg(*otg.MkaTxSc) MkaTxSc
	// provides marshal interface
	Marshal() marshalMkaTxSc
	// provides unmarshal interface
	Unmarshal() unMarshalMkaTxSc
	// validate validates MkaTxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaTxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MkaTxSc.
	Name() string
	// SetName assigns string provided by user to MkaTxSc
	SetName(value string) MkaTxSc
	// SystemId returns string, set in MkaTxSc.
	SystemId() string
	// SetSystemId assigns string provided by user to MkaTxSc
	SetSystemId(value string) MkaTxSc
	// PortId returns uint32, set in MkaTxSc.
	PortId() uint32
	// SetPortId assigns uint32 provided by user to MkaTxSc
	SetPortId(value uint32) MkaTxSc
	// HasPortId checks if PortId has been set in MkaTxSc
	HasPortId() bool
	// StartingMessageNumber returns uint64, set in MkaTxSc.
	StartingMessageNumber() uint64
	// SetStartingMessageNumber assigns uint64 provided by user to MkaTxSc
	SetStartingMessageNumber(value uint64) MkaTxSc
	// HasStartingMessageNumber checks if StartingMessageNumber has been set in MkaTxSc
	HasStartingMessageNumber() bool
}

// Tx SC name.
// Name returns a string
func (obj *mkaTxSc) Name() string {

	return *obj.obj.Name

}

// Tx SC name.
// SetName sets the string value in the MkaTxSc object
func (obj *mkaTxSc) SetName(value string) MkaTxSc {

	obj.obj.Name = &value
	return obj
}

// System ID.
// SystemId returns a string
func (obj *mkaTxSc) SystemId() string {

	return *obj.obj.SystemId

}

// System ID.
// SetSystemId sets the string value in the MkaTxSc object
func (obj *mkaTxSc) SetSystemId(value string) MkaTxSc {

	obj.obj.SystemId = &value
	return obj
}

// Port ID.
// PortId returns a uint32
func (obj *mkaTxSc) PortId() uint32 {

	return *obj.obj.PortId

}

// Port ID.
// PortId returns a uint32
func (obj *mkaTxSc) HasPortId() bool {
	return obj.obj.PortId != nil
}

// Port ID.
// SetPortId sets the uint32 value in the MkaTxSc object
func (obj *mkaTxSc) SetPortId(value uint32) MkaTxSc {

	obj.obj.PortId = &value
	return obj
}

// Starting Message Number.
// StartingMessageNumber returns a uint64
func (obj *mkaTxSc) StartingMessageNumber() uint64 {

	return *obj.obj.StartingMessageNumber

}

// Starting Message Number.
// StartingMessageNumber returns a uint64
func (obj *mkaTxSc) HasStartingMessageNumber() bool {
	return obj.obj.StartingMessageNumber != nil
}

// Starting Message Number.
// SetStartingMessageNumber sets the uint64 value in the MkaTxSc object
func (obj *mkaTxSc) SetStartingMessageNumber(value uint64) MkaTxSc {

	obj.obj.StartingMessageNumber = &value
	return obj
}

func (obj *mkaTxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface MkaTxSc")
	}

	// SystemId is required
	if obj.obj.SystemId == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SystemId is required field on interface MkaTxSc")
	}
	if obj.obj.SystemId != nil {

		err := obj.validateMac(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MkaTxSc.SystemId"))
		}

	}

	if obj.obj.PortId != nil {

		if *obj.obj.PortId < 1 || *obj.obj.PortId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MkaTxSc.PortId <= 65535 but Got %d", *obj.obj.PortId))
		}

	}

	if obj.obj.StartingMessageNumber != nil {

		if *obj.obj.StartingMessageNumber < 1 || *obj.obj.StartingMessageNumber > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MkaTxSc.StartingMessageNumber <= 4294967295 but Got %d", *obj.obj.StartingMessageNumber))
		}

	}

}

func (obj *mkaTxSc) setDefault() {
	if obj.obj.PortId == nil {
		obj.SetPortId(1)
	}
	if obj.obj.StartingMessageNumber == nil {
		obj.SetStartingMessageNumber(1)
	}

}
