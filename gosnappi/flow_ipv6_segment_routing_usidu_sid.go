package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingUsiduSid *****
type flowIpv6SegmentRoutingUsiduSid struct {
	validation
	obj          *otg.FlowIpv6SegmentRoutingUsiduSid
	marshaller   marshalFlowIpv6SegmentRoutingUsiduSid
	unMarshaller unMarshalFlowIpv6SegmentRoutingUsiduSid
}

func NewFlowIpv6SegmentRoutingUsiduSid() FlowIpv6SegmentRoutingUsiduSid {
	obj := flowIpv6SegmentRoutingUsiduSid{obj: &otg.FlowIpv6SegmentRoutingUsiduSid{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingUsiduSid) msg() *otg.FlowIpv6SegmentRoutingUsiduSid {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingUsiduSid) setMsg(msg *otg.FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsiduSid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingUsiduSid struct {
	obj *flowIpv6SegmentRoutingUsiduSid
}

type marshalFlowIpv6SegmentRoutingUsiduSid interface {
	// ToProto marshals FlowIpv6SegmentRoutingUsiduSid to protobuf object *otg.FlowIpv6SegmentRoutingUsiduSid
	ToProto() (*otg.FlowIpv6SegmentRoutingUsiduSid, error)
	// ToPbText marshals FlowIpv6SegmentRoutingUsiduSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingUsiduSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingUsiduSid to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingUsiduSid struct {
	obj *flowIpv6SegmentRoutingUsiduSid
}

type unMarshalFlowIpv6SegmentRoutingUsiduSid interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingUsiduSid from protobuf object *otg.FlowIpv6SegmentRoutingUsiduSid
	FromProto(msg *otg.FlowIpv6SegmentRoutingUsiduSid) (FlowIpv6SegmentRoutingUsiduSid, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingUsiduSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingUsiduSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingUsiduSid from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingUsiduSid) Marshal() marshalFlowIpv6SegmentRoutingUsiduSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingUsiduSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingUsiduSid) Unmarshal() unMarshalFlowIpv6SegmentRoutingUsiduSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingUsiduSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingUsiduSid) ToProto() (*otg.FlowIpv6SegmentRoutingUsiduSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingUsiduSid) FromProto(msg *otg.FlowIpv6SegmentRoutingUsiduSid) (FlowIpv6SegmentRoutingUsiduSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingUsiduSid) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsiduSid) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsiduSid) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsiduSid) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsiduSid) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsiduSid) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingUsiduSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsiduSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsiduSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingUsiduSid) Clone() (FlowIpv6SegmentRoutingUsiduSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingUsiduSid()
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

// FlowIpv6SegmentRoutingUsiduSid is one uSID value within a compressed uSID container (RFC 9800 Section 3). Represents a single network function (node SID, adjacency SID, or service SID) packed consecutively after the Locator Block inside the 128-bit container. For F3216 (RFC 9800 Section 3), each uSID is 16 bits (4 hex characters).
type FlowIpv6SegmentRoutingUsiduSid interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingUsiduSid to protobuf object *otg.FlowIpv6SegmentRoutingUsiduSid
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingUsiduSid
	// setMsg unmarshals FlowIpv6SegmentRoutingUsiduSid from protobuf object *otg.FlowIpv6SegmentRoutingUsiduSid
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsiduSid
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingUsiduSid
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingUsiduSid
	// validate validates FlowIpv6SegmentRoutingUsiduSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingUsiduSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Usid returns string, set in FlowIpv6SegmentRoutingUsiduSid.
	Usid() string
	// SetUsid assigns string provided by user to FlowIpv6SegmentRoutingUsiduSid
	SetUsid(value string) FlowIpv6SegmentRoutingUsiduSid
	// HasUsid checks if Usid has been set in FlowIpv6SegmentRoutingUsiduSid
	HasUsid() bool
}

// The uSID value expressed as a hex string (RFC 9800 Section 3). The string width must match the uSID bit length divided by 4. For F3216 (16-bit uSID): 4 hex characters. The value 0x0000 is reserved as the End-of-Container (EoC) marker and must not be used as a uSID value (RFC 9800 Section 3). Example: "0001" identifies node 1; with locator fc00::/32 this contributes the 0001 field to produce container fc00:0:1:...::
// Usid returns a string
func (obj *flowIpv6SegmentRoutingUsiduSid) Usid() string {

	return *obj.obj.Usid

}

// The uSID value expressed as a hex string (RFC 9800 Section 3). The string width must match the uSID bit length divided by 4. For F3216 (16-bit uSID): 4 hex characters. The value 0x0000 is reserved as the End-of-Container (EoC) marker and must not be used as a uSID value (RFC 9800 Section 3). Example: "0001" identifies node 1; with locator fc00::/32 this contributes the 0001 field to produce container fc00:0:1:...::
// Usid returns a string
func (obj *flowIpv6SegmentRoutingUsiduSid) HasUsid() bool {
	return obj.obj.Usid != nil
}

// The uSID value expressed as a hex string (RFC 9800 Section 3). The string width must match the uSID bit length divided by 4. For F3216 (16-bit uSID): 4 hex characters. The value 0x0000 is reserved as the End-of-Container (EoC) marker and must not be used as a uSID value (RFC 9800 Section 3). Example: "0001" identifies node 1; with locator fc00::/32 this contributes the 0001 field to produce container fc00:0:1:...::
// SetUsid sets the string value in the FlowIpv6SegmentRoutingUsiduSid object
func (obj *flowIpv6SegmentRoutingUsiduSid) SetUsid(value string) FlowIpv6SegmentRoutingUsiduSid {

	obj.obj.Usid = &value
	return obj
}

func (obj *flowIpv6SegmentRoutingUsiduSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Usid != nil {

		err := obj.validateHex(obj.Usid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SegmentRoutingUsiduSid.Usid"))
		}

	}

}

func (obj *flowIpv6SegmentRoutingUsiduSid) setDefault() {

}
