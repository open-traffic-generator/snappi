package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2V4RRExternalType2 *****
type ospfv2V4RRExternalType2 struct {
	validation
	obj          *otg.Ospfv2V4RRExternalType2
	marshaller   marshalOspfv2V4RRExternalType2
	unMarshaller unMarshalOspfv2V4RRExternalType2
	flagsHolder  Ospfv2V4RRExtdPrefixFlags
}

func NewOspfv2V4RRExternalType2() Ospfv2V4RRExternalType2 {
	obj := ospfv2V4RRExternalType2{obj: &otg.Ospfv2V4RRExternalType2{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2V4RRExternalType2) msg() *otg.Ospfv2V4RRExternalType2 {
	return obj.obj
}

func (obj *ospfv2V4RRExternalType2) setMsg(msg *otg.Ospfv2V4RRExternalType2) Ospfv2V4RRExternalType2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2V4RRExternalType2 struct {
	obj *ospfv2V4RRExternalType2
}

type marshalOspfv2V4RRExternalType2 interface {
	// ToProto marshals Ospfv2V4RRExternalType2 to protobuf object *otg.Ospfv2V4RRExternalType2
	ToProto() (*otg.Ospfv2V4RRExternalType2, error)
	// ToPbText marshals Ospfv2V4RRExternalType2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2V4RRExternalType2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2V4RRExternalType2 to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2V4RRExternalType2 struct {
	obj *ospfv2V4RRExternalType2
}

type unMarshalOspfv2V4RRExternalType2 interface {
	// FromProto unmarshals Ospfv2V4RRExternalType2 from protobuf object *otg.Ospfv2V4RRExternalType2
	FromProto(msg *otg.Ospfv2V4RRExternalType2) (Ospfv2V4RRExternalType2, error)
	// FromPbText unmarshals Ospfv2V4RRExternalType2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2V4RRExternalType2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2V4RRExternalType2 from JSON text
	FromJson(value string) error
}

func (obj *ospfv2V4RRExternalType2) Marshal() marshalOspfv2V4RRExternalType2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2V4RRExternalType2{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2V4RRExternalType2) Unmarshal() unMarshalOspfv2V4RRExternalType2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2V4RRExternalType2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2V4RRExternalType2) ToProto() (*otg.Ospfv2V4RRExternalType2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2V4RRExternalType2) FromProto(msg *otg.Ospfv2V4RRExternalType2) (Ospfv2V4RRExternalType2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2V4RRExternalType2) ToPbText() (string, error) {
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

func (m *unMarshalospfv2V4RRExternalType2) FromPbText(value string) error {
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

func (m *marshalospfv2V4RRExternalType2) ToYaml() (string, error) {
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

func (m *unMarshalospfv2V4RRExternalType2) FromYaml(value string) error {
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

func (m *marshalospfv2V4RRExternalType2) ToJson() (string, error) {
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

func (m *unMarshalospfv2V4RRExternalType2) FromJson(value string) error {
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

func (obj *ospfv2V4RRExternalType2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2V4RRExternalType2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2V4RRExternalType2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2V4RRExternalType2) Clone() (Ospfv2V4RRExternalType2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2V4RRExternalType2()
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

func (obj *ospfv2V4RRExternalType2) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2V4RRExternalType2 is container for Intra-Area.
type Ospfv2V4RRExternalType2 interface {
	Validation
	// msg marshals Ospfv2V4RRExternalType2 to protobuf object *otg.Ospfv2V4RRExternalType2
	// and doesn't set defaults
	msg() *otg.Ospfv2V4RRExternalType2
	// setMsg unmarshals Ospfv2V4RRExternalType2 from protobuf object *otg.Ospfv2V4RRExternalType2
	// and doesn't set defaults
	setMsg(*otg.Ospfv2V4RRExternalType2) Ospfv2V4RRExternalType2
	// provides marshal interface
	Marshal() marshalOspfv2V4RRExternalType2
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2V4RRExternalType2
	// validate validates Ospfv2V4RRExternalType2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2V4RRExternalType2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns Ospfv2V4RRExtdPrefixFlags, set in Ospfv2V4RRExternalType2.
	// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
	Flags() Ospfv2V4RRExtdPrefixFlags
	// SetFlags assigns Ospfv2V4RRExtdPrefixFlags provided by user to Ospfv2V4RRExternalType2.
	// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
	SetFlags(value Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRExternalType2
	// HasFlags checks if Flags has been set in Ospfv2V4RRExternalType2
	HasFlags() bool
	setNil()
}

// One-octet field contains flags applicable to the prefix.
// Flags returns a Ospfv2V4RRExtdPrefixFlags
func (obj *ospfv2V4RRExternalType2) Flags() Ospfv2V4RRExtdPrefixFlags {
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
func (obj *ospfv2V4RRExternalType2) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One-octet field contains flags applicable to the prefix.
// SetFlags sets the Ospfv2V4RRExtdPrefixFlags value in the Ospfv2V4RRExternalType2 object
func (obj *ospfv2V4RRExternalType2) SetFlags(value Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRExternalType2 {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

func (obj *ospfv2V4RRExternalType2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2V4RRExternalType2) setDefault() {

}
