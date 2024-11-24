package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2V4RRNssaExternal *****
type ospfv2V4RRNssaExternal struct {
	validation
	obj          *otg.Ospfv2V4RRNssaExternal
	marshaller   marshalOspfv2V4RRNssaExternal
	unMarshaller unMarshalOspfv2V4RRNssaExternal
	flagsHolder  Ospfv2V4RRExtdPrefixFlags
}

func NewOspfv2V4RRNssaExternal() Ospfv2V4RRNssaExternal {
	obj := ospfv2V4RRNssaExternal{obj: &otg.Ospfv2V4RRNssaExternal{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2V4RRNssaExternal) msg() *otg.Ospfv2V4RRNssaExternal {
	return obj.obj
}

func (obj *ospfv2V4RRNssaExternal) setMsg(msg *otg.Ospfv2V4RRNssaExternal) Ospfv2V4RRNssaExternal {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2V4RRNssaExternal struct {
	obj *ospfv2V4RRNssaExternal
}

type marshalOspfv2V4RRNssaExternal interface {
	// ToProto marshals Ospfv2V4RRNssaExternal to protobuf object *otg.Ospfv2V4RRNssaExternal
	ToProto() (*otg.Ospfv2V4RRNssaExternal, error)
	// ToPbText marshals Ospfv2V4RRNssaExternal to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2V4RRNssaExternal to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2V4RRNssaExternal to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2V4RRNssaExternal to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2V4RRNssaExternal struct {
	obj *ospfv2V4RRNssaExternal
}

type unMarshalOspfv2V4RRNssaExternal interface {
	// FromProto unmarshals Ospfv2V4RRNssaExternal from protobuf object *otg.Ospfv2V4RRNssaExternal
	FromProto(msg *otg.Ospfv2V4RRNssaExternal) (Ospfv2V4RRNssaExternal, error)
	// FromPbText unmarshals Ospfv2V4RRNssaExternal from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2V4RRNssaExternal from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2V4RRNssaExternal from JSON text
	FromJson(value string) error
}

func (obj *ospfv2V4RRNssaExternal) Marshal() marshalOspfv2V4RRNssaExternal {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2V4RRNssaExternal{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2V4RRNssaExternal) Unmarshal() unMarshalOspfv2V4RRNssaExternal {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2V4RRNssaExternal{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2V4RRNssaExternal) ToProto() (*otg.Ospfv2V4RRNssaExternal, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2V4RRNssaExternal) FromProto(msg *otg.Ospfv2V4RRNssaExternal) (Ospfv2V4RRNssaExternal, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2V4RRNssaExternal) ToPbText() (string, error) {
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

func (m *unMarshalospfv2V4RRNssaExternal) FromPbText(value string) error {
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

func (m *marshalospfv2V4RRNssaExternal) ToYaml() (string, error) {
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

func (m *unMarshalospfv2V4RRNssaExternal) FromYaml(value string) error {
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

func (m *marshalospfv2V4RRNssaExternal) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2V4RRNssaExternal) ToJson() (string, error) {
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

func (m *unMarshalospfv2V4RRNssaExternal) FromJson(value string) error {
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

func (obj *ospfv2V4RRNssaExternal) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2V4RRNssaExternal) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2V4RRNssaExternal) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2V4RRNssaExternal) Clone() (Ospfv2V4RRNssaExternal, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2V4RRNssaExternal()
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

func (obj *ospfv2V4RRNssaExternal) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2V4RRNssaExternal is container for Intra-Area.
type Ospfv2V4RRNssaExternal interface {
	Validation
	// msg marshals Ospfv2V4RRNssaExternal to protobuf object *otg.Ospfv2V4RRNssaExternal
	// and doesn't set defaults
	msg() *otg.Ospfv2V4RRNssaExternal
	// setMsg unmarshals Ospfv2V4RRNssaExternal from protobuf object *otg.Ospfv2V4RRNssaExternal
	// and doesn't set defaults
	setMsg(*otg.Ospfv2V4RRNssaExternal) Ospfv2V4RRNssaExternal
	// provides marshal interface
	Marshal() marshalOspfv2V4RRNssaExternal
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2V4RRNssaExternal
	// validate validates Ospfv2V4RRNssaExternal
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2V4RRNssaExternal, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns Ospfv2V4RRExtdPrefixFlags, set in Ospfv2V4RRNssaExternal.
	// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
	Flags() Ospfv2V4RRExtdPrefixFlags
	// SetFlags assigns Ospfv2V4RRExtdPrefixFlags provided by user to Ospfv2V4RRNssaExternal.
	// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
	SetFlags(value Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRNssaExternal
	// HasFlags checks if Flags has been set in Ospfv2V4RRNssaExternal
	HasFlags() bool
	// Propagation returns bool, set in Ospfv2V4RRNssaExternal.
	Propagation() bool
	// SetPropagation assigns bool provided by user to Ospfv2V4RRNssaExternal
	SetPropagation(value bool) Ospfv2V4RRNssaExternal
	// HasPropagation checks if Propagation has been set in Ospfv2V4RRNssaExternal
	HasPropagation() bool
	setNil()
}

// One-octet field contains flags applicable to the prefix.
// Flags returns a Ospfv2V4RRExtdPrefixFlags
func (obj *ospfv2V4RRNssaExternal) Flags() Ospfv2V4RRExtdPrefixFlags {
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
func (obj *ospfv2V4RRNssaExternal) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One-octet field contains flags applicable to the prefix.
// SetFlags sets the Ospfv2V4RRExtdPrefixFlags value in the Ospfv2V4RRNssaExternal object
func (obj *ospfv2V4RRNssaExternal) SetFlags(value Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRNssaExternal {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The flag is set True if LSA will be propagated between Areas.
// Propagation returns a bool
func (obj *ospfv2V4RRNssaExternal) Propagation() bool {

	return *obj.obj.Propagation

}

// The flag is set True if LSA will be propagated between Areas.
// Propagation returns a bool
func (obj *ospfv2V4RRNssaExternal) HasPropagation() bool {
	return obj.obj.Propagation != nil
}

// The flag is set True if LSA will be propagated between Areas.
// SetPropagation sets the bool value in the Ospfv2V4RRNssaExternal object
func (obj *ospfv2V4RRNssaExternal) SetPropagation(value bool) Ospfv2V4RRNssaExternal {

	obj.obj.Propagation = &value
	return obj
}

func (obj *ospfv2V4RRNssaExternal) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2V4RRNssaExternal) setDefault() {
	if obj.obj.Propagation == nil {
		obj.SetPropagation(false)
	}

}
