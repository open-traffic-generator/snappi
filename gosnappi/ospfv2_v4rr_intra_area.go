package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2V4RRIntraArea *****
type ospfv2V4RRIntraArea struct {
	validation
	obj          *otg.Ospfv2V4RRIntraArea
	marshaller   marshalOspfv2V4RRIntraArea
	unMarshaller unMarshalOspfv2V4RRIntraArea
	flagsHolder  Ospfv2V4RRExtdPrefixFlags
}

func NewOspfv2V4RRIntraArea() Ospfv2V4RRIntraArea {
	obj := ospfv2V4RRIntraArea{obj: &otg.Ospfv2V4RRIntraArea{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2V4RRIntraArea) msg() *otg.Ospfv2V4RRIntraArea {
	return obj.obj
}

func (obj *ospfv2V4RRIntraArea) setMsg(msg *otg.Ospfv2V4RRIntraArea) Ospfv2V4RRIntraArea {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2V4RRIntraArea struct {
	obj *ospfv2V4RRIntraArea
}

type marshalOspfv2V4RRIntraArea interface {
	// ToProto marshals Ospfv2V4RRIntraArea to protobuf object *otg.Ospfv2V4RRIntraArea
	ToProto() (*otg.Ospfv2V4RRIntraArea, error)
	// ToPbText marshals Ospfv2V4RRIntraArea to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2V4RRIntraArea to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2V4RRIntraArea to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2V4RRIntraArea to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2V4RRIntraArea struct {
	obj *ospfv2V4RRIntraArea
}

type unMarshalOspfv2V4RRIntraArea interface {
	// FromProto unmarshals Ospfv2V4RRIntraArea from protobuf object *otg.Ospfv2V4RRIntraArea
	FromProto(msg *otg.Ospfv2V4RRIntraArea) (Ospfv2V4RRIntraArea, error)
	// FromPbText unmarshals Ospfv2V4RRIntraArea from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2V4RRIntraArea from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2V4RRIntraArea from JSON text
	FromJson(value string) error
}

func (obj *ospfv2V4RRIntraArea) Marshal() marshalOspfv2V4RRIntraArea {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2V4RRIntraArea{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2V4RRIntraArea) Unmarshal() unMarshalOspfv2V4RRIntraArea {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2V4RRIntraArea{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2V4RRIntraArea) ToProto() (*otg.Ospfv2V4RRIntraArea, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2V4RRIntraArea) FromProto(msg *otg.Ospfv2V4RRIntraArea) (Ospfv2V4RRIntraArea, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2V4RRIntraArea) ToPbText() (string, error) {
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

func (m *unMarshalospfv2V4RRIntraArea) FromPbText(value string) error {
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

func (m *marshalospfv2V4RRIntraArea) ToYaml() (string, error) {
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

func (m *unMarshalospfv2V4RRIntraArea) FromYaml(value string) error {
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

func (m *marshalospfv2V4RRIntraArea) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2V4RRIntraArea) ToJson() (string, error) {
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

func (m *unMarshalospfv2V4RRIntraArea) FromJson(value string) error {
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

func (obj *ospfv2V4RRIntraArea) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2V4RRIntraArea) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2V4RRIntraArea) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2V4RRIntraArea) Clone() (Ospfv2V4RRIntraArea, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2V4RRIntraArea()
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

func (obj *ospfv2V4RRIntraArea) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2V4RRIntraArea is container for Intra-Area.
type Ospfv2V4RRIntraArea interface {
	Validation
	// msg marshals Ospfv2V4RRIntraArea to protobuf object *otg.Ospfv2V4RRIntraArea
	// and doesn't set defaults
	msg() *otg.Ospfv2V4RRIntraArea
	// setMsg unmarshals Ospfv2V4RRIntraArea from protobuf object *otg.Ospfv2V4RRIntraArea
	// and doesn't set defaults
	setMsg(*otg.Ospfv2V4RRIntraArea) Ospfv2V4RRIntraArea
	// provides marshal interface
	Marshal() marshalOspfv2V4RRIntraArea
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2V4RRIntraArea
	// validate validates Ospfv2V4RRIntraArea
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2V4RRIntraArea, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns Ospfv2V4RRExtdPrefixFlags, set in Ospfv2V4RRIntraArea.
	// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
	Flags() Ospfv2V4RRExtdPrefixFlags
	// SetFlags assigns Ospfv2V4RRExtdPrefixFlags provided by user to Ospfv2V4RRIntraArea.
	// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
	SetFlags(value Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRIntraArea
	// HasFlags checks if Flags has been set in Ospfv2V4RRIntraArea
	HasFlags() bool
	setNil()
}

// One-octet field contains flags applicable to the prefix.
// Flags returns a Ospfv2V4RRExtdPrefixFlags
func (obj *ospfv2V4RRIntraArea) Flags() Ospfv2V4RRExtdPrefixFlags {
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
func (obj *ospfv2V4RRIntraArea) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One-octet field contains flags applicable to the prefix.
// SetFlags sets the Ospfv2V4RRExtdPrefixFlags value in the Ospfv2V4RRIntraArea object
func (obj *ospfv2V4RRIntraArea) SetFlags(value Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRIntraArea {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

func (obj *ospfv2V4RRIntraArea) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2V4RRIntraArea) setDefault() {

}
