package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingTlvPad *****
type flowIpv6SegmentRoutingTlvPad struct {
	validation
	obj          *otg.FlowIpv6SegmentRoutingTlvPad
	marshaller   marshalFlowIpv6SegmentRoutingTlvPad
	unMarshaller unMarshalFlowIpv6SegmentRoutingTlvPad
}

func NewFlowIpv6SegmentRoutingTlvPad() FlowIpv6SegmentRoutingTlvPad {
	obj := flowIpv6SegmentRoutingTlvPad{obj: &otg.FlowIpv6SegmentRoutingTlvPad{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingTlvPad) msg() *otg.FlowIpv6SegmentRoutingTlvPad {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingTlvPad) setMsg(msg *otg.FlowIpv6SegmentRoutingTlvPad) FlowIpv6SegmentRoutingTlvPad {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingTlvPad struct {
	obj *flowIpv6SegmentRoutingTlvPad
}

type marshalFlowIpv6SegmentRoutingTlvPad interface {
	// ToProto marshals FlowIpv6SegmentRoutingTlvPad to protobuf object *otg.FlowIpv6SegmentRoutingTlvPad
	ToProto() (*otg.FlowIpv6SegmentRoutingTlvPad, error)
	// ToPbText marshals FlowIpv6SegmentRoutingTlvPad to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingTlvPad to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingTlvPad to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingTlvPad struct {
	obj *flowIpv6SegmentRoutingTlvPad
}

type unMarshalFlowIpv6SegmentRoutingTlvPad interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingTlvPad from protobuf object *otg.FlowIpv6SegmentRoutingTlvPad
	FromProto(msg *otg.FlowIpv6SegmentRoutingTlvPad) (FlowIpv6SegmentRoutingTlvPad, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingTlvPad from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingTlvPad from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingTlvPad from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingTlvPad) Marshal() marshalFlowIpv6SegmentRoutingTlvPad {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingTlvPad{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingTlvPad) Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvPad {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingTlvPad{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingTlvPad) ToProto() (*otg.FlowIpv6SegmentRoutingTlvPad, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingTlvPad) FromProto(msg *otg.FlowIpv6SegmentRoutingTlvPad) (FlowIpv6SegmentRoutingTlvPad, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingTlvPad) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvPad) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvPad) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvPad) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvPad) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvPad) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingTlvPad) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvPad) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvPad) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingTlvPad) Clone() (FlowIpv6SegmentRoutingTlvPad, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingTlvPad()
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

// FlowIpv6SegmentRoutingTlvPad is padding TLV used when more than one byte of padding is required to align subsequent TLVs or pad the header to an 8-octet boundary.
type FlowIpv6SegmentRoutingTlvPad interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingTlvPad to protobuf object *otg.FlowIpv6SegmentRoutingTlvPad
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingTlvPad
	// setMsg unmarshals FlowIpv6SegmentRoutingTlvPad from protobuf object *otg.FlowIpv6SegmentRoutingTlvPad
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingTlvPad) FlowIpv6SegmentRoutingTlvPad
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingTlvPad
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvPad
	// validate validates FlowIpv6SegmentRoutingTlvPad
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingTlvPad, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Padding returns string, set in FlowIpv6SegmentRoutingTlvPad.
	Padding() string
	// SetPadding assigns string provided by user to FlowIpv6SegmentRoutingTlvPad
	SetPadding(value string) FlowIpv6SegmentRoutingTlvPad
	// HasPadding checks if Padding has been set in FlowIpv6SegmentRoutingTlvPad
	HasPadding() bool
}

// Variable-length padding bytes. Must be set to 0 on transmission.
// Padding returns a string
func (obj *flowIpv6SegmentRoutingTlvPad) Padding() string {

	return *obj.obj.Padding

}

// Variable-length padding bytes. Must be set to 0 on transmission.
// Padding returns a string
func (obj *flowIpv6SegmentRoutingTlvPad) HasPadding() bool {
	return obj.obj.Padding != nil
}

// Variable-length padding bytes. Must be set to 0 on transmission.
// SetPadding sets the string value in the FlowIpv6SegmentRoutingTlvPad object
func (obj *flowIpv6SegmentRoutingTlvPad) SetPadding(value string) FlowIpv6SegmentRoutingTlvPad {

	obj.obj.Padding = &value
	return obj
}

func (obj *flowIpv6SegmentRoutingTlvPad) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Padding != nil {

		err := obj.validateHex(obj.Padding())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SegmentRoutingTlvPad.Padding"))
		}

	}

}

func (obj *flowIpv6SegmentRoutingTlvPad) setDefault() {
	if obj.obj.Padding == nil {
		obj.SetPadding("00")
	}

}
