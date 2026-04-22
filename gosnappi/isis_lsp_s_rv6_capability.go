package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRv6Capability *****
type isisLspSRv6Capability struct {
	validation
	obj            *otg.IsisLspSRv6Capability
	marshaller     marshalIsisLspSRv6Capability
	unMarshaller   unMarshalIsisLspSRv6Capability
	nodeMsdsHolder IsisLspSRv6NodeMsd
}

func NewIsisLspSRv6Capability() IsisLspSRv6Capability {
	obj := isisLspSRv6Capability{obj: &otg.IsisLspSRv6Capability{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRv6Capability) msg() *otg.IsisLspSRv6Capability {
	return obj.obj
}

func (obj *isisLspSRv6Capability) setMsg(msg *otg.IsisLspSRv6Capability) IsisLspSRv6Capability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRv6Capability struct {
	obj *isisLspSRv6Capability
}

type marshalIsisLspSRv6Capability interface {
	// ToProto marshals IsisLspSRv6Capability to protobuf object *otg.IsisLspSRv6Capability
	ToProto() (*otg.IsisLspSRv6Capability, error)
	// ToPbText marshals IsisLspSRv6Capability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRv6Capability to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRv6Capability to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspSRv6Capability struct {
	obj *isisLspSRv6Capability
}

type unMarshalIsisLspSRv6Capability interface {
	// FromProto unmarshals IsisLspSRv6Capability from protobuf object *otg.IsisLspSRv6Capability
	FromProto(msg *otg.IsisLspSRv6Capability) (IsisLspSRv6Capability, error)
	// FromPbText unmarshals IsisLspSRv6Capability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRv6Capability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRv6Capability from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRv6Capability) Marshal() marshalIsisLspSRv6Capability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRv6Capability{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRv6Capability) Unmarshal() unMarshalIsisLspSRv6Capability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRv6Capability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRv6Capability) ToProto() (*otg.IsisLspSRv6Capability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRv6Capability) FromProto(msg *otg.IsisLspSRv6Capability) (IsisLspSRv6Capability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRv6Capability) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRv6Capability) FromPbText(value string) error {
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

func (m *marshalisisLspSRv6Capability) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRv6Capability) FromYaml(value string) error {
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

func (m *marshalisisLspSRv6Capability) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRv6Capability) FromJson(value string) error {
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

func (obj *isisLspSRv6Capability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRv6Capability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRv6Capability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRv6Capability) Clone() (IsisLspSRv6Capability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRv6Capability()
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

func (obj *isisLspSRv6Capability) setNil() {
	obj.nodeMsdsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspSRv6Capability is sRv6 Capabilities Sub-TLV (sub-TLV type 25) learned from the IS-IS Router CAPABILITY TLV (TLV 242) in a received LSP. Indicates that the originating router is an SRv6 Segment Endpoint Node and carries optional OAM support (O-flag) and node-level SRv6 MSD values. Reference: RFC 9352 Section 2, RFC 8491.
type IsisLspSRv6Capability interface {
	Validation
	// msg marshals IsisLspSRv6Capability to protobuf object *otg.IsisLspSRv6Capability
	// and doesn't set defaults
	msg() *otg.IsisLspSRv6Capability
	// setMsg unmarshals IsisLspSRv6Capability from protobuf object *otg.IsisLspSRv6Capability
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRv6Capability) IsisLspSRv6Capability
	// provides marshal interface
	Marshal() marshalIsisLspSRv6Capability
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRv6Capability
	// validate validates IsisLspSRv6Capability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRv6Capability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// OFlag returns bool, set in IsisLspSRv6Capability.
	OFlag() bool
	// SetOFlag assigns bool provided by user to IsisLspSRv6Capability
	SetOFlag(value bool) IsisLspSRv6Capability
	// HasOFlag checks if OFlag has been set in IsisLspSRv6Capability
	HasOFlag() bool
	// NodeMsds returns IsisLspSRv6NodeMsd, set in IsisLspSRv6Capability.
	// IsisLspSRv6NodeMsd is node-level SRv6 Maximum SID Depth (MSD) values learned from a received LSP. Each field represents the advertised depth limit for a specific SRv6 forwarding behavior. MSD Type 41 = Max SL, Type 42 = Max End Pop, Type 44 = Max H.Encaps, Type 45 = Max End D. Reference: RFC 8491, RFC 9352 Section 6.
	NodeMsds() IsisLspSRv6NodeMsd
	// SetNodeMsds assigns IsisLspSRv6NodeMsd provided by user to IsisLspSRv6Capability.
	// IsisLspSRv6NodeMsd is node-level SRv6 Maximum SID Depth (MSD) values learned from a received LSP. Each field represents the advertised depth limit for a specific SRv6 forwarding behavior. MSD Type 41 = Max SL, Type 42 = Max End Pop, Type 44 = Max H.Encaps, Type 45 = Max End D. Reference: RFC 8491, RFC 9352 Section 6.
	SetNodeMsds(value IsisLspSRv6NodeMsd) IsisLspSRv6Capability
	// HasNodeMsds checks if NodeMsds has been set in IsisLspSRv6Capability
	HasNodeMsds() bool
	// CFlag returns bool, set in IsisLspSRv6Capability.
	CFlag() bool
	// SetCFlag assigns bool provided by user to IsisLspSRv6Capability
	SetCFlag(value bool) IsisLspSRv6Capability
	// HasCFlag checks if CFlag has been set in IsisLspSRv6Capability
	HasCFlag() bool
	setNil()
}

// OAM flag (bit 1). When set, the originating router supports the use of the O-bit in the Segment Routing Header (SRH) for OAM operations as defined in RFC 9259.
// OFlag returns a bool
func (obj *isisLspSRv6Capability) OFlag() bool {

	return *obj.obj.OFlag

}

// OAM flag (bit 1). When set, the originating router supports the use of the O-bit in the Segment Routing Header (SRH) for OAM operations as defined in RFC 9259.
// OFlag returns a bool
func (obj *isisLspSRv6Capability) HasOFlag() bool {
	return obj.obj.OFlag != nil
}

// OAM flag (bit 1). When set, the originating router supports the use of the O-bit in the Segment Routing Header (SRH) for OAM operations as defined in RFC 9259.
// SetOFlag sets the bool value in the IsisLspSRv6Capability object
func (obj *isisLspSRv6Capability) SetOFlag(value bool) IsisLspSRv6Capability {

	obj.obj.OFlag = &value
	return obj
}

// Node-level SRv6 Maximum SID Depth (MSD) values learned from MSD sub-TLVs within the Router Capability TLV.
// NodeMsds returns a IsisLspSRv6NodeMsd
func (obj *isisLspSRv6Capability) NodeMsds() IsisLspSRv6NodeMsd {
	if obj.obj.NodeMsds == nil {
		obj.obj.NodeMsds = NewIsisLspSRv6NodeMsd().msg()
	}
	if obj.nodeMsdsHolder == nil {
		obj.nodeMsdsHolder = &isisLspSRv6NodeMsd{obj: obj.obj.NodeMsds}
	}
	return obj.nodeMsdsHolder
}

// Node-level SRv6 Maximum SID Depth (MSD) values learned from MSD sub-TLVs within the Router Capability TLV.
// NodeMsds returns a IsisLspSRv6NodeMsd
func (obj *isisLspSRv6Capability) HasNodeMsds() bool {
	return obj.obj.NodeMsds != nil
}

// Node-level SRv6 Maximum SID Depth (MSD) values learned from MSD sub-TLVs within the Router Capability TLV.
// SetNodeMsds sets the IsisLspSRv6NodeMsd value in the IsisLspSRv6Capability object
func (obj *isisLspSRv6Capability) SetNodeMsds(value IsisLspSRv6NodeMsd) IsisLspSRv6Capability {

	obj.nodeMsdsHolder = nil
	obj.obj.NodeMsds = value.msg()

	return obj
}

// Compression (uSID) flag. When set, the originating router supports Micro-SID (uSID) compressed encoding for SRv6 SIDs. Learned from the Flags field of the SRv6 Capabilities Sub-TLV (sub-TLV type 25). Reference: RFC 9352 Section 2, RFC 9800.
// CFlag returns a bool
func (obj *isisLspSRv6Capability) CFlag() bool {

	return *obj.obj.CFlag

}

// Compression (uSID) flag. When set, the originating router supports Micro-SID (uSID) compressed encoding for SRv6 SIDs. Learned from the Flags field of the SRv6 Capabilities Sub-TLV (sub-TLV type 25). Reference: RFC 9352 Section 2, RFC 9800.
// CFlag returns a bool
func (obj *isisLspSRv6Capability) HasCFlag() bool {
	return obj.obj.CFlag != nil
}

// Compression (uSID) flag. When set, the originating router supports Micro-SID (uSID) compressed encoding for SRv6 SIDs. Learned from the Flags field of the SRv6 Capabilities Sub-TLV (sub-TLV type 25). Reference: RFC 9352 Section 2, RFC 9800.
// SetCFlag sets the bool value in the IsisLspSRv6Capability object
func (obj *isisLspSRv6Capability) SetCFlag(value bool) IsisLspSRv6Capability {

	obj.obj.CFlag = &value
	return obj
}

func (obj *isisLspSRv6Capability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.NodeMsds != nil {

		obj.NodeMsds().validateObj(vObj, set_default)
	}

}

func (obj *isisLspSRv6Capability) setDefault() {

}
