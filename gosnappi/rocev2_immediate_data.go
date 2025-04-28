package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2ImmediateData *****
type rocev2ImmediateData struct {
	validation
	obj          *otg.Rocev2ImmediateData
	marshaller   marshalRocev2ImmediateData
	unMarshaller unMarshalRocev2ImmediateData
}

func NewRocev2ImmediateData() Rocev2ImmediateData {
	obj := rocev2ImmediateData{obj: &otg.Rocev2ImmediateData{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2ImmediateData) msg() *otg.Rocev2ImmediateData {
	return obj.obj
}

func (obj *rocev2ImmediateData) setMsg(msg *otg.Rocev2ImmediateData) Rocev2ImmediateData {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2ImmediateData struct {
	obj *rocev2ImmediateData
}

type marshalRocev2ImmediateData interface {
	// ToProto marshals Rocev2ImmediateData to protobuf object *otg.Rocev2ImmediateData
	ToProto() (*otg.Rocev2ImmediateData, error)
	// ToPbText marshals Rocev2ImmediateData to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2ImmediateData to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2ImmediateData to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2ImmediateData struct {
	obj *rocev2ImmediateData
}

type unMarshalRocev2ImmediateData interface {
	// FromProto unmarshals Rocev2ImmediateData from protobuf object *otg.Rocev2ImmediateData
	FromProto(msg *otg.Rocev2ImmediateData) (Rocev2ImmediateData, error)
	// FromPbText unmarshals Rocev2ImmediateData from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2ImmediateData from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2ImmediateData from JSON text
	FromJson(value string) error
}

func (obj *rocev2ImmediateData) Marshal() marshalRocev2ImmediateData {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2ImmediateData{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2ImmediateData) Unmarshal() unMarshalRocev2ImmediateData {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2ImmediateData{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2ImmediateData) ToProto() (*otg.Rocev2ImmediateData, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2ImmediateData) FromProto(msg *otg.Rocev2ImmediateData) (Rocev2ImmediateData, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2ImmediateData) ToPbText() (string, error) {
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

func (m *unMarshalrocev2ImmediateData) FromPbText(value string) error {
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

func (m *marshalrocev2ImmediateData) ToYaml() (string, error) {
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

func (m *unMarshalrocev2ImmediateData) FromYaml(value string) error {
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

func (m *marshalrocev2ImmediateData) ToJson() (string, error) {
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

func (m *unMarshalrocev2ImmediateData) FromJson(value string) error {
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

func (obj *rocev2ImmediateData) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2ImmediateData) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2ImmediateData) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2ImmediateData) Clone() (Rocev2ImmediateData, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2ImmediateData()
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

// Rocev2ImmediateData is four bytes of immediate Data for SEND/WRITE with immediate.
type Rocev2ImmediateData interface {
	Validation
	// msg marshals Rocev2ImmediateData to protobuf object *otg.Rocev2ImmediateData
	// and doesn't set defaults
	msg() *otg.Rocev2ImmediateData
	// setMsg unmarshals Rocev2ImmediateData from protobuf object *otg.Rocev2ImmediateData
	// and doesn't set defaults
	setMsg(*otg.Rocev2ImmediateData) Rocev2ImmediateData
	// provides marshal interface
	Marshal() marshalRocev2ImmediateData
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2ImmediateData
	// validate validates Rocev2ImmediateData
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2ImmediateData, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ImmediateData returns string, set in Rocev2ImmediateData.
	ImmediateData() string
	// SetImmediateData assigns string provided by user to Rocev2ImmediateData
	SetImmediateData(value string) Rocev2ImmediateData
	// HasImmediateData checks if ImmediateData has been set in Rocev2ImmediateData
	HasImmediateData() bool
}

// Four bytes of immediate Data for SEND/WRITE with immediate.
// ImmediateData returns a string
func (obj *rocev2ImmediateData) ImmediateData() string {

	return *obj.obj.ImmediateData

}

// Four bytes of immediate Data for SEND/WRITE with immediate.
// ImmediateData returns a string
func (obj *rocev2ImmediateData) HasImmediateData() bool {
	return obj.obj.ImmediateData != nil
}

// Four bytes of immediate Data for SEND/WRITE with immediate.
// SetImmediateData sets the string value in the Rocev2ImmediateData object
func (obj *rocev2ImmediateData) SetImmediateData(value string) Rocev2ImmediateData {

	obj.obj.ImmediateData = &value
	return obj
}

func (obj *rocev2ImmediateData) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ImmediateData != nil {

		err := obj.validateHex(obj.ImmediateData())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Rocev2ImmediateData.ImmediateData"))
		}

	}

}

func (obj *rocev2ImmediateData) setDefault() {
	if obj.obj.ImmediateData == nil {
		obj.SetImmediateData("00000000")
	}

}
