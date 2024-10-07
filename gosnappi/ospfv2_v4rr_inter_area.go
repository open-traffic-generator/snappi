package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2V4RRInterArea *****
type ospfv2V4RRInterArea struct {
	validation
	obj          *otg.Ospfv2V4RRInterArea
	marshaller   marshalOspfv2V4RRInterArea
	unMarshaller unMarshalOspfv2V4RRInterArea
	flagsHolder  Ospfv2V4RRExtdPrefixFlags
}

func NewOspfv2V4RRInterArea() Ospfv2V4RRInterArea {
	obj := ospfv2V4RRInterArea{obj: &otg.Ospfv2V4RRInterArea{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2V4RRInterArea) msg() *otg.Ospfv2V4RRInterArea {
	return obj.obj
}

func (obj *ospfv2V4RRInterArea) setMsg(msg *otg.Ospfv2V4RRInterArea) Ospfv2V4RRInterArea {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2V4RRInterArea struct {
	obj *ospfv2V4RRInterArea
}

type marshalOspfv2V4RRInterArea interface {
	// ToProto marshals Ospfv2V4RRInterArea to protobuf object *otg.Ospfv2V4RRInterArea
	ToProto() (*otg.Ospfv2V4RRInterArea, error)
	// ToPbText marshals Ospfv2V4RRInterArea to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2V4RRInterArea to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2V4RRInterArea to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2V4RRInterArea struct {
	obj *ospfv2V4RRInterArea
}

type unMarshalOspfv2V4RRInterArea interface {
	// FromProto unmarshals Ospfv2V4RRInterArea from protobuf object *otg.Ospfv2V4RRInterArea
	FromProto(msg *otg.Ospfv2V4RRInterArea) (Ospfv2V4RRInterArea, error)
	// FromPbText unmarshals Ospfv2V4RRInterArea from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2V4RRInterArea from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2V4RRInterArea from JSON text
	FromJson(value string) error
}

func (obj *ospfv2V4RRInterArea) Marshal() marshalOspfv2V4RRInterArea {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2V4RRInterArea{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2V4RRInterArea) Unmarshal() unMarshalOspfv2V4RRInterArea {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2V4RRInterArea{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2V4RRInterArea) ToProto() (*otg.Ospfv2V4RRInterArea, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2V4RRInterArea) FromProto(msg *otg.Ospfv2V4RRInterArea) (Ospfv2V4RRInterArea, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2V4RRInterArea) ToPbText() (string, error) {
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

func (m *unMarshalospfv2V4RRInterArea) FromPbText(value string) error {
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

func (m *marshalospfv2V4RRInterArea) ToYaml() (string, error) {
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

func (m *unMarshalospfv2V4RRInterArea) FromYaml(value string) error {
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

func (m *marshalospfv2V4RRInterArea) ToJson() (string, error) {
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

func (m *unMarshalospfv2V4RRInterArea) FromJson(value string) error {
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

func (obj *ospfv2V4RRInterArea) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2V4RRInterArea) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2V4RRInterArea) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2V4RRInterArea) Clone() (Ospfv2V4RRInterArea, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2V4RRInterArea()
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

func (obj *ospfv2V4RRInterArea) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2V4RRInterArea is container for Intra-Area.
type Ospfv2V4RRInterArea interface {
	Validation
	// msg marshals Ospfv2V4RRInterArea to protobuf object *otg.Ospfv2V4RRInterArea
	// and doesn't set defaults
	msg() *otg.Ospfv2V4RRInterArea
	// setMsg unmarshals Ospfv2V4RRInterArea from protobuf object *otg.Ospfv2V4RRInterArea
	// and doesn't set defaults
	setMsg(*otg.Ospfv2V4RRInterArea) Ospfv2V4RRInterArea
	// provides marshal interface
	Marshal() marshalOspfv2V4RRInterArea
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2V4RRInterArea
	// validate validates Ospfv2V4RRInterArea
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2V4RRInterArea, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns Ospfv2V4RRExtdPrefixFlags, set in Ospfv2V4RRInterArea.
	// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
	Flags() Ospfv2V4RRExtdPrefixFlags
	// SetFlags assigns Ospfv2V4RRExtdPrefixFlags provided by user to Ospfv2V4RRInterArea.
	// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
	SetFlags(value Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRInterArea
	// HasFlags checks if Flags has been set in Ospfv2V4RRInterArea
	HasFlags() bool
	setNil()
}

// One-octet field contains flags applicable to the prefix.
// Flags returns a Ospfv2V4RRExtdPrefixFlags
func (obj *ospfv2V4RRInterArea) Flags() Ospfv2V4RRExtdPrefixFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewOspfv2V4RRExtdPrefixFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &ospfv2V4RRExtdPrefixFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// One-octet field contains flags applicable to the prefix.
// Flags returns a Ospfv2V4RRExtdPrefixFlags
func (obj *ospfv2V4RRInterArea) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One-octet field contains flags applicable to the prefix.
// SetFlags sets the Ospfv2V4RRExtdPrefixFlags value in the Ospfv2V4RRInterArea object
func (obj *ospfv2V4RRInterArea) SetFlags(value Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRInterArea {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

func (obj *ospfv2V4RRInterArea) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2V4RRInterArea) setDefault() {

}
