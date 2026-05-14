package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6UsidDstuSid *****
type flowIpv6UsidDstuSid struct {
	validation
	obj          *otg.FlowIpv6UsidDstuSid
	marshaller   marshalFlowIpv6UsidDstuSid
	unMarshaller unMarshalFlowIpv6UsidDstuSid
}

func NewFlowIpv6UsidDstuSid() FlowIpv6UsidDstuSid {
	obj := flowIpv6UsidDstuSid{obj: &otg.FlowIpv6UsidDstuSid{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6UsidDstuSid) msg() *otg.FlowIpv6UsidDstuSid {
	return obj.obj
}

func (obj *flowIpv6UsidDstuSid) setMsg(msg *otg.FlowIpv6UsidDstuSid) FlowIpv6UsidDstuSid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6UsidDstuSid struct {
	obj *flowIpv6UsidDstuSid
}

type marshalFlowIpv6UsidDstuSid interface {
	// ToProto marshals FlowIpv6UsidDstuSid to protobuf object *otg.FlowIpv6UsidDstuSid
	ToProto() (*otg.FlowIpv6UsidDstuSid, error)
	// ToPbText marshals FlowIpv6UsidDstuSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6UsidDstuSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6UsidDstuSid to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6UsidDstuSid struct {
	obj *flowIpv6UsidDstuSid
}

type unMarshalFlowIpv6UsidDstuSid interface {
	// FromProto unmarshals FlowIpv6UsidDstuSid from protobuf object *otg.FlowIpv6UsidDstuSid
	FromProto(msg *otg.FlowIpv6UsidDstuSid) (FlowIpv6UsidDstuSid, error)
	// FromPbText unmarshals FlowIpv6UsidDstuSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6UsidDstuSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6UsidDstuSid from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6UsidDstuSid) Marshal() marshalFlowIpv6UsidDstuSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6UsidDstuSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6UsidDstuSid) Unmarshal() unMarshalFlowIpv6UsidDstuSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6UsidDstuSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6UsidDstuSid) ToProto() (*otg.FlowIpv6UsidDstuSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6UsidDstuSid) FromProto(msg *otg.FlowIpv6UsidDstuSid) (FlowIpv6UsidDstuSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6UsidDstuSid) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6UsidDstuSid) FromPbText(value string) error {
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

func (m *marshalflowIpv6UsidDstuSid) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6UsidDstuSid) FromYaml(value string) error {
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

func (m *marshalflowIpv6UsidDstuSid) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6UsidDstuSid) FromJson(value string) error {
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

func (obj *flowIpv6UsidDstuSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6UsidDstuSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6UsidDstuSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6UsidDstuSid) Clone() (FlowIpv6UsidDstuSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6UsidDstuSid()
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

// FlowIpv6UsidDstuSid is one uSID value within the no-SRH uSID container (RFC 9800 Section 3). For F3216 (16-bit uSID): 4 hex characters. The value 0x0000 is reserved as the End-of-Container marker and must not be used.
type FlowIpv6UsidDstuSid interface {
	Validation
	// msg marshals FlowIpv6UsidDstuSid to protobuf object *otg.FlowIpv6UsidDstuSid
	// and doesn't set defaults
	msg() *otg.FlowIpv6UsidDstuSid
	// setMsg unmarshals FlowIpv6UsidDstuSid from protobuf object *otg.FlowIpv6UsidDstuSid
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6UsidDstuSid) FlowIpv6UsidDstuSid
	// provides marshal interface
	Marshal() marshalFlowIpv6UsidDstuSid
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6UsidDstuSid
	// validate validates FlowIpv6UsidDstuSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6UsidDstuSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Usid returns string, set in FlowIpv6UsidDstuSid.
	Usid() string
	// SetUsid assigns string provided by user to FlowIpv6UsidDstuSid
	SetUsid(value string) FlowIpv6UsidDstuSid
	// HasUsid checks if Usid has been set in FlowIpv6UsidDstuSid
	HasUsid() bool
}

// The uSID value as a hex string. For F3216: 4 hex characters. Any '0' padding before the before uSID value should be explicitly configured. Example: "0001".
// Usid returns a string
func (obj *flowIpv6UsidDstuSid) Usid() string {

	return *obj.obj.Usid

}

// The uSID value as a hex string. For F3216: 4 hex characters. Any '0' padding before the before uSID value should be explicitly configured. Example: "0001".
// Usid returns a string
func (obj *flowIpv6UsidDstuSid) HasUsid() bool {
	return obj.obj.Usid != nil
}

// The uSID value as a hex string. For F3216: 4 hex characters. Any '0' padding before the before uSID value should be explicitly configured. Example: "0001".
// SetUsid sets the string value in the FlowIpv6UsidDstuSid object
func (obj *flowIpv6UsidDstuSid) SetUsid(value string) FlowIpv6UsidDstuSid {

	obj.obj.Usid = &value
	return obj
}

func (obj *flowIpv6UsidDstuSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Usid != nil {

		err := obj.validateHex(obj.Usid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6UsidDstuSid.Usid"))
		}

	}

}

func (obj *flowIpv6UsidDstuSid) setDefault() {

}
