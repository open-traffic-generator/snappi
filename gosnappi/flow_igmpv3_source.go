package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIgmpv3Source *****
type flowIgmpv3Source struct {
	validation
	obj           *otg.FlowIgmpv3Source
	marshaller    marshalFlowIgmpv3Source
	unMarshaller  unMarshalFlowIgmpv3Source
	addressHolder PatternFlowIgmpv3SourceAddress
}

func NewFlowIgmpv3Source() FlowIgmpv3Source {
	obj := flowIgmpv3Source{obj: &otg.FlowIgmpv3Source{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIgmpv3Source) msg() *otg.FlowIgmpv3Source {
	return obj.obj
}

func (obj *flowIgmpv3Source) setMsg(msg *otg.FlowIgmpv3Source) FlowIgmpv3Source {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIgmpv3Source struct {
	obj *flowIgmpv3Source
}

type marshalFlowIgmpv3Source interface {
	// ToProto marshals FlowIgmpv3Source to protobuf object *otg.FlowIgmpv3Source
	ToProto() (*otg.FlowIgmpv3Source, error)
	// ToPbText marshals FlowIgmpv3Source to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIgmpv3Source to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIgmpv3Source to JSON text
	ToJson() (string, error)
}

type unMarshalflowIgmpv3Source struct {
	obj *flowIgmpv3Source
}

type unMarshalFlowIgmpv3Source interface {
	// FromProto unmarshals FlowIgmpv3Source from protobuf object *otg.FlowIgmpv3Source
	FromProto(msg *otg.FlowIgmpv3Source) (FlowIgmpv3Source, error)
	// FromPbText unmarshals FlowIgmpv3Source from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIgmpv3Source from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIgmpv3Source from JSON text
	FromJson(value string) error
}

func (obj *flowIgmpv3Source) Marshal() marshalFlowIgmpv3Source {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIgmpv3Source{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIgmpv3Source) Unmarshal() unMarshalFlowIgmpv3Source {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIgmpv3Source{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIgmpv3Source) ToProto() (*otg.FlowIgmpv3Source, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIgmpv3Source) FromProto(msg *otg.FlowIgmpv3Source) (FlowIgmpv3Source, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIgmpv3Source) ToPbText() (string, error) {
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

func (m *unMarshalflowIgmpv3Source) FromPbText(value string) error {
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

func (m *marshalflowIgmpv3Source) ToYaml() (string, error) {
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

func (m *unMarshalflowIgmpv3Source) FromYaml(value string) error {
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

func (m *marshalflowIgmpv3Source) ToJson() (string, error) {
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

func (m *unMarshalflowIgmpv3Source) FromJson(value string) error {
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

func (obj *flowIgmpv3Source) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIgmpv3Source) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIgmpv3Source) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIgmpv3Source) Clone() (FlowIgmpv3Source, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIgmpv3Source()
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

func (obj *flowIgmpv3Source) setNil() {
	obj.addressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIgmpv3Source is a source address entry. The entries form a vector of IPv4 unicast addresses, one per source, and their number is given by number_of_sources.
type FlowIgmpv3Source interface {
	Validation
	// msg marshals FlowIgmpv3Source to protobuf object *otg.FlowIgmpv3Source
	// and doesn't set defaults
	msg() *otg.FlowIgmpv3Source
	// setMsg unmarshals FlowIgmpv3Source from protobuf object *otg.FlowIgmpv3Source
	// and doesn't set defaults
	setMsg(*otg.FlowIgmpv3Source) FlowIgmpv3Source
	// provides marshal interface
	Marshal() marshalFlowIgmpv3Source
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIgmpv3Source
	// validate validates FlowIgmpv3Source
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIgmpv3Source, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns PatternFlowIgmpv3SourceAddress, set in FlowIgmpv3Source.
	// PatternFlowIgmpv3SourceAddress is unicast IPv4 source address, use a counter to send a range of sources.
	Address() PatternFlowIgmpv3SourceAddress
	// SetAddress assigns PatternFlowIgmpv3SourceAddress provided by user to FlowIgmpv3Source.
	// PatternFlowIgmpv3SourceAddress is unicast IPv4 source address, use a counter to send a range of sources.
	SetAddress(value PatternFlowIgmpv3SourceAddress) FlowIgmpv3Source
	// HasAddress checks if Address has been set in FlowIgmpv3Source
	HasAddress() bool
	setNil()
}

// description is TBD
// Address returns a PatternFlowIgmpv3SourceAddress
func (obj *flowIgmpv3Source) Address() PatternFlowIgmpv3SourceAddress {
	if obj.obj.Address == nil {
		obj.obj.Address = NewPatternFlowIgmpv3SourceAddress().msg()
	}
	if obj.addressHolder == nil {
		obj.addressHolder = &patternFlowIgmpv3SourceAddress{obj: obj.obj.Address}
	}
	return obj.addressHolder
}

// description is TBD
// Address returns a PatternFlowIgmpv3SourceAddress
func (obj *flowIgmpv3Source) HasAddress() bool {
	return obj.obj.Address != nil
}

// description is TBD
// SetAddress sets the PatternFlowIgmpv3SourceAddress value in the FlowIgmpv3Source object
func (obj *flowIgmpv3Source) SetAddress(value PatternFlowIgmpv3SourceAddress) FlowIgmpv3Source {

	obj.addressHolder = nil
	obj.obj.Address = value.msg()

	return obj
}

func (obj *flowIgmpv3Source) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		obj.Address().validateObj(vObj, set_default)
	}

}

func (obj *flowIgmpv3Source) setDefault() {

}
