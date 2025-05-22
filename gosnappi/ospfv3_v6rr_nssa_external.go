package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3V6RRNssaExternal *****
type ospfv3V6RRNssaExternal struct {
	validation
	obj                *otg.Ospfv3V6RRNssaExternal
	marshaller         marshalOspfv3V6RRNssaExternal
	unMarshaller       unMarshalOspfv3V6RRNssaExternal
	capabilitiesHolder Ospfv3V6RRCapabilities
}

func NewOspfv3V6RRNssaExternal() Ospfv3V6RRNssaExternal {
	obj := ospfv3V6RRNssaExternal{obj: &otg.Ospfv3V6RRNssaExternal{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3V6RRNssaExternal) msg() *otg.Ospfv3V6RRNssaExternal {
	return obj.obj
}

func (obj *ospfv3V6RRNssaExternal) setMsg(msg *otg.Ospfv3V6RRNssaExternal) Ospfv3V6RRNssaExternal {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3V6RRNssaExternal struct {
	obj *ospfv3V6RRNssaExternal
}

type marshalOspfv3V6RRNssaExternal interface {
	// ToProto marshals Ospfv3V6RRNssaExternal to protobuf object *otg.Ospfv3V6RRNssaExternal
	ToProto() (*otg.Ospfv3V6RRNssaExternal, error)
	// ToPbText marshals Ospfv3V6RRNssaExternal to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3V6RRNssaExternal to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3V6RRNssaExternal to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3V6RRNssaExternal to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3V6RRNssaExternal struct {
	obj *ospfv3V6RRNssaExternal
}

type unMarshalOspfv3V6RRNssaExternal interface {
	// FromProto unmarshals Ospfv3V6RRNssaExternal from protobuf object *otg.Ospfv3V6RRNssaExternal
	FromProto(msg *otg.Ospfv3V6RRNssaExternal) (Ospfv3V6RRNssaExternal, error)
	// FromPbText unmarshals Ospfv3V6RRNssaExternal from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3V6RRNssaExternal from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3V6RRNssaExternal from JSON text
	FromJson(value string) error
}

func (obj *ospfv3V6RRNssaExternal) Marshal() marshalOspfv3V6RRNssaExternal {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3V6RRNssaExternal{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3V6RRNssaExternal) Unmarshal() unMarshalOspfv3V6RRNssaExternal {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3V6RRNssaExternal{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3V6RRNssaExternal) ToProto() (*otg.Ospfv3V6RRNssaExternal, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3V6RRNssaExternal) FromProto(msg *otg.Ospfv3V6RRNssaExternal) (Ospfv3V6RRNssaExternal, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3V6RRNssaExternal) ToPbText() (string, error) {
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

func (m *unMarshalospfv3V6RRNssaExternal) FromPbText(value string) error {
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

func (m *marshalospfv3V6RRNssaExternal) ToYaml() (string, error) {
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

func (m *unMarshalospfv3V6RRNssaExternal) FromYaml(value string) error {
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

func (m *marshalospfv3V6RRNssaExternal) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3V6RRNssaExternal) ToJson() (string, error) {
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

func (m *unMarshalospfv3V6RRNssaExternal) FromJson(value string) error {
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

func (obj *ospfv3V6RRNssaExternal) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3V6RRNssaExternal) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3V6RRNssaExternal) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3V6RRNssaExternal) Clone() (Ospfv3V6RRNssaExternal, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3V6RRNssaExternal()
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

func (obj *ospfv3V6RRNssaExternal) setNil() {
	obj.capabilitiesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3V6RRNssaExternal is container for the forwarding address of NSSA External route origin configuration.
type Ospfv3V6RRNssaExternal interface {
	Validation
	// msg marshals Ospfv3V6RRNssaExternal to protobuf object *otg.Ospfv3V6RRNssaExternal
	// and doesn't set defaults
	msg() *otg.Ospfv3V6RRNssaExternal
	// setMsg unmarshals Ospfv3V6RRNssaExternal from protobuf object *otg.Ospfv3V6RRNssaExternal
	// and doesn't set defaults
	setMsg(*otg.Ospfv3V6RRNssaExternal) Ospfv3V6RRNssaExternal
	// provides marshal interface
	Marshal() marshalOspfv3V6RRNssaExternal
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3V6RRNssaExternal
	// validate validates Ospfv3V6RRNssaExternal
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3V6RRNssaExternal, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Capabilities returns Ospfv3V6RRCapabilities, set in Ospfv3V6RRNssaExternal.
	// Ospfv3V6RRCapabilities is container for the capabilities associated with route origin.
	Capabilities() Ospfv3V6RRCapabilities
	// SetCapabilities assigns Ospfv3V6RRCapabilities provided by user to Ospfv3V6RRNssaExternal.
	// Ospfv3V6RRCapabilities is container for the capabilities associated with route origin.
	SetCapabilities(value Ospfv3V6RRCapabilities) Ospfv3V6RRNssaExternal
	// HasCapabilities checks if Capabilities has been set in Ospfv3V6RRNssaExternal
	HasCapabilities() bool
	setNil()
}

// Configuration for capabilities associated with route origin.
// Capabilities returns a Ospfv3V6RRCapabilities
func (obj *ospfv3V6RRNssaExternal) Capabilities() Ospfv3V6RRCapabilities {
	if obj.obj.Capabilities == nil {
		obj.obj.Capabilities = NewOspfv3V6RRCapabilities().msg()
	}
	if obj.capabilitiesHolder == nil {
		obj.capabilitiesHolder = &ospfv3V6RRCapabilities{obj: obj.obj.Capabilities}
	}
	return obj.capabilitiesHolder
}

// Configuration for capabilities associated with route origin.
// Capabilities returns a Ospfv3V6RRCapabilities
func (obj *ospfv3V6RRNssaExternal) HasCapabilities() bool {
	return obj.obj.Capabilities != nil
}

// Configuration for capabilities associated with route origin.
// SetCapabilities sets the Ospfv3V6RRCapabilities value in the Ospfv3V6RRNssaExternal object
func (obj *ospfv3V6RRNssaExternal) SetCapabilities(value Ospfv3V6RRCapabilities) Ospfv3V6RRNssaExternal {

	obj.capabilitiesHolder = nil
	obj.obj.Capabilities = value.msg()

	return obj
}

func (obj *ospfv3V6RRNssaExternal) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Capabilities != nil {

		obj.Capabilities().validateObj(vObj, set_default)
	}

}

func (obj *ospfv3V6RRNssaExternal) setDefault() {

}
