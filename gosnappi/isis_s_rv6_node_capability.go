package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6NodeCapability *****
type isisSRv6NodeCapability struct {
	validation
	obj            *otg.IsisSRv6NodeCapability
	marshaller     marshalIsisSRv6NodeCapability
	unMarshaller   unMarshalIsisSRv6NodeCapability
	nodeMsdsHolder IsisSRv6NodeMsd
}

func NewIsisSRv6NodeCapability() IsisSRv6NodeCapability {
	obj := isisSRv6NodeCapability{obj: &otg.IsisSRv6NodeCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6NodeCapability) msg() *otg.IsisSRv6NodeCapability {
	return obj.obj
}

func (obj *isisSRv6NodeCapability) setMsg(msg *otg.IsisSRv6NodeCapability) IsisSRv6NodeCapability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6NodeCapability struct {
	obj *isisSRv6NodeCapability
}

type marshalIsisSRv6NodeCapability interface {
	// ToProto marshals IsisSRv6NodeCapability to protobuf object *otg.IsisSRv6NodeCapability
	ToProto() (*otg.IsisSRv6NodeCapability, error)
	// ToPbText marshals IsisSRv6NodeCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6NodeCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6NodeCapability to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6NodeCapability struct {
	obj *isisSRv6NodeCapability
}

type unMarshalIsisSRv6NodeCapability interface {
	// FromProto unmarshals IsisSRv6NodeCapability from protobuf object *otg.IsisSRv6NodeCapability
	FromProto(msg *otg.IsisSRv6NodeCapability) (IsisSRv6NodeCapability, error)
	// FromPbText unmarshals IsisSRv6NodeCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6NodeCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6NodeCapability from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6NodeCapability) Marshal() marshalIsisSRv6NodeCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6NodeCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6NodeCapability) Unmarshal() unMarshalIsisSRv6NodeCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6NodeCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6NodeCapability) ToProto() (*otg.IsisSRv6NodeCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6NodeCapability) FromProto(msg *otg.IsisSRv6NodeCapability) (IsisSRv6NodeCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6NodeCapability) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6NodeCapability) FromPbText(value string) error {
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

func (m *marshalisisSRv6NodeCapability) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6NodeCapability) FromYaml(value string) error {
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

func (m *marshalisisSRv6NodeCapability) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6NodeCapability) FromJson(value string) error {
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

func (obj *isisSRv6NodeCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6NodeCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6NodeCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6NodeCapability) Clone() (IsisSRv6NodeCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6NodeCapability()
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

func (obj *isisSRv6NodeCapability) setNil() {
	obj.nodeMsdsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRv6NodeCapability is sRv6 Capabilities Sub-TLV (sub-TLV type 25) carried within the IS-IS Router CAPABILITY TLV (TLV 242, RFC 7981). Announces that this IS-IS router is an SRv6 Segment Endpoint Node and optionally supports the O-bit for OAM operations. Node-level SRv6 Maximum SID Depth (MSD) values are also advertised here per RFC 8491. Reference: RFC 9352 Section 2.
type IsisSRv6NodeCapability interface {
	Validation
	// msg marshals IsisSRv6NodeCapability to protobuf object *otg.IsisSRv6NodeCapability
	// and doesn't set defaults
	msg() *otg.IsisSRv6NodeCapability
	// setMsg unmarshals IsisSRv6NodeCapability from protobuf object *otg.IsisSRv6NodeCapability
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6NodeCapability) IsisSRv6NodeCapability
	// provides marshal interface
	Marshal() marshalIsisSRv6NodeCapability
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6NodeCapability
	// validate validates IsisSRv6NodeCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6NodeCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// OFlag returns bool, set in IsisSRv6NodeCapability.
	OFlag() bool
	// SetOFlag assigns bool provided by user to IsisSRv6NodeCapability
	SetOFlag(value bool) IsisSRv6NodeCapability
	// HasOFlag checks if OFlag has been set in IsisSRv6NodeCapability
	HasOFlag() bool
	// NodeMsds returns IsisSRv6NodeMsd, set in IsisSRv6NodeCapability.
	// IsisSRv6NodeMsd is node-level SRv6 Maximum SID Depth (MSD) capabilities. A non-zero value causes the corresponding MSD sub-TLV to be included in the advertisement; a value of 0 (default) suppresses it. This follows RFC 8491 where MSD is a list of (MSD-Type, MSD-Value) tuples. Reference: RFC 8491, RFC 9352 Section 6.
	NodeMsds() IsisSRv6NodeMsd
	// SetNodeMsds assigns IsisSRv6NodeMsd provided by user to IsisSRv6NodeCapability.
	// IsisSRv6NodeMsd is node-level SRv6 Maximum SID Depth (MSD) capabilities. A non-zero value causes the corresponding MSD sub-TLV to be included in the advertisement; a value of 0 (default) suppresses it. This follows RFC 8491 where MSD is a list of (MSD-Type, MSD-Value) tuples. Reference: RFC 8491, RFC 9352 Section 6.
	SetNodeMsds(value IsisSRv6NodeMsd) IsisSRv6NodeCapability
	// HasNodeMsds checks if NodeMsds has been set in IsisSRv6NodeCapability
	HasNodeMsds() bool
	// CFlag returns bool, set in IsisSRv6NodeCapability.
	CFlag() bool
	// SetCFlag assigns bool provided by user to IsisSRv6NodeCapability
	SetCFlag(value bool) IsisSRv6NodeCapability
	// HasCFlag checks if CFlag has been set in IsisSRv6NodeCapability
	HasCFlag() bool
	setNil()
}

// OAM flag (bit 0 of the Flags field in the SRv6 Capabilities Sub-TLV). When set, indicates that this router supports the use of the O-bit in the Segment Routing Header (SRH) for Operations, Administration and Maintenance (OAM) operations as defined in RFC 9259.
// OFlag returns a bool
func (obj *isisSRv6NodeCapability) OFlag() bool {

	return *obj.obj.OFlag

}

// OAM flag (bit 0 of the Flags field in the SRv6 Capabilities Sub-TLV). When set, indicates that this router supports the use of the O-bit in the Segment Routing Header (SRH) for Operations, Administration and Maintenance (OAM) operations as defined in RFC 9259.
// OFlag returns a bool
func (obj *isisSRv6NodeCapability) HasOFlag() bool {
	return obj.obj.OFlag != nil
}

// OAM flag (bit 0 of the Flags field in the SRv6 Capabilities Sub-TLV). When set, indicates that this router supports the use of the O-bit in the Segment Routing Header (SRH) for Operations, Administration and Maintenance (OAM) operations as defined in RFC 9259.
// SetOFlag sets the bool value in the IsisSRv6NodeCapability object
func (obj *isisSRv6NodeCapability) SetOFlag(value bool) IsisSRv6NodeCapability {

	obj.obj.OFlag = &value
	return obj
}

// Node-level SRv6 Maximum SID Depth (MSD) capabilities advertised as MSD sub-TLVs within the Router Capability TLV per RFC 8491. These signal the maximum SRH stack depth this router supports for specific SRv6 forwarding behaviors (RFC 9352 Section 6).
// NodeMsds returns a IsisSRv6NodeMsd
func (obj *isisSRv6NodeCapability) NodeMsds() IsisSRv6NodeMsd {
	if obj.obj.NodeMsds == nil {
		obj.obj.NodeMsds = NewIsisSRv6NodeMsd().msg()
	}
	if obj.nodeMsdsHolder == nil {
		obj.nodeMsdsHolder = &isisSRv6NodeMsd{obj: obj.obj.NodeMsds}
	}
	return obj.nodeMsdsHolder
}

// Node-level SRv6 Maximum SID Depth (MSD) capabilities advertised as MSD sub-TLVs within the Router Capability TLV per RFC 8491. These signal the maximum SRH stack depth this router supports for specific SRv6 forwarding behaviors (RFC 9352 Section 6).
// NodeMsds returns a IsisSRv6NodeMsd
func (obj *isisSRv6NodeCapability) HasNodeMsds() bool {
	return obj.obj.NodeMsds != nil
}

// Node-level SRv6 Maximum SID Depth (MSD) capabilities advertised as MSD sub-TLVs within the Router Capability TLV per RFC 8491. These signal the maximum SRH stack depth this router supports for specific SRv6 forwarding behaviors (RFC 9352 Section 6).
// SetNodeMsds sets the IsisSRv6NodeMsd value in the IsisSRv6NodeCapability object
func (obj *isisSRv6NodeCapability) SetNodeMsds(value IsisSRv6NodeMsd) IsisSRv6NodeCapability {

	obj.nodeMsdsHolder = nil
	obj.obj.NodeMsds = value.msg()

	return obj
}

// Compression (uSID) flag at the node level (RFC 9800). When set, announces that this IS-IS router supports Micro-SID (uSID) compressed SRv6 encoding. This flag is the prerequisite: it must be set before individual End SIDs or Adj SIDs are marked as uSIDs via their own c_flag fields. It is carried in the SRv6 Capabilities Sub-TLV (sub-TLV type 25) Flags field. To advertise a uSID-capable node using the F3216 format (uSID block fc00::/32, node locator fc00:0:1:: /48):
// Step 1 - Set this c_flag = true.
// Step 2 - Configure IsisSRv6.Locator with:
// locator "fc00:0:1::", prefix_length 48,
// sid_structure: lb=32, ln=16, fn=16, arg=0.
// Step 3 - Add End SIDs under the locator (c_flag=true per SID):
// function "0001" => uSID fc00:0:1:1::  (End)
// function "0002" => uSID fc00:0:1:2::  (End.DT46)
// Step 4 - Add Adj SIDs on the interface (c_flag=true per SID):
// function "00c8" => uSID fc00:0:1:c8::  (End.X)
// A headend can then pack uSIDs from the same /32 block into a single 128-bit container - no SRH required. With three nodes (fc00:0:1::, fc00:0:2::, fc00:0:3::) each advertising End function "0001", the headend constructs destination address fc00:0:1:1:2:1:3:1 to steer traffic over the 3-hop path in one 128-bit address. Reference: RFC 9352 Section 2, RFC 9800.
// CFlag returns a bool
func (obj *isisSRv6NodeCapability) CFlag() bool {

	return *obj.obj.CFlag

}

// Compression (uSID) flag at the node level (RFC 9800). When set, announces that this IS-IS router supports Micro-SID (uSID) compressed SRv6 encoding. This flag is the prerequisite: it must be set before individual End SIDs or Adj SIDs are marked as uSIDs via their own c_flag fields. It is carried in the SRv6 Capabilities Sub-TLV (sub-TLV type 25) Flags field. To advertise a uSID-capable node using the F3216 format (uSID block fc00::/32, node locator fc00:0:1:: /48):
// Step 1 - Set this c_flag = true.
// Step 2 - Configure IsisSRv6.Locator with:
// locator "fc00:0:1::", prefix_length 48,
// sid_structure: lb=32, ln=16, fn=16, arg=0.
// Step 3 - Add End SIDs under the locator (c_flag=true per SID):
// function "0001" => uSID fc00:0:1:1::  (End)
// function "0002" => uSID fc00:0:1:2::  (End.DT46)
// Step 4 - Add Adj SIDs on the interface (c_flag=true per SID):
// function "00c8" => uSID fc00:0:1:c8::  (End.X)
// A headend can then pack uSIDs from the same /32 block into a single 128-bit container - no SRH required. With three nodes (fc00:0:1::, fc00:0:2::, fc00:0:3::) each advertising End function "0001", the headend constructs destination address fc00:0:1:1:2:1:3:1 to steer traffic over the 3-hop path in one 128-bit address. Reference: RFC 9352 Section 2, RFC 9800.
// CFlag returns a bool
func (obj *isisSRv6NodeCapability) HasCFlag() bool {
	return obj.obj.CFlag != nil
}

// Compression (uSID) flag at the node level (RFC 9800). When set, announces that this IS-IS router supports Micro-SID (uSID) compressed SRv6 encoding. This flag is the prerequisite: it must be set before individual End SIDs or Adj SIDs are marked as uSIDs via their own c_flag fields. It is carried in the SRv6 Capabilities Sub-TLV (sub-TLV type 25) Flags field. To advertise a uSID-capable node using the F3216 format (uSID block fc00::/32, node locator fc00:0:1:: /48):
// Step 1 - Set this c_flag = true.
// Step 2 - Configure IsisSRv6.Locator with:
// locator "fc00:0:1::", prefix_length 48,
// sid_structure: lb=32, ln=16, fn=16, arg=0.
// Step 3 - Add End SIDs under the locator (c_flag=true per SID):
// function "0001" => uSID fc00:0:1:1::  (End)
// function "0002" => uSID fc00:0:1:2::  (End.DT46)
// Step 4 - Add Adj SIDs on the interface (c_flag=true per SID):
// function "00c8" => uSID fc00:0:1:c8::  (End.X)
// A headend can then pack uSIDs from the same /32 block into a single 128-bit container - no SRH required. With three nodes (fc00:0:1::, fc00:0:2::, fc00:0:3::) each advertising End function "0001", the headend constructs destination address fc00:0:1:1:2:1:3:1 to steer traffic over the 3-hop path in one 128-bit address. Reference: RFC 9352 Section 2, RFC 9800.
// SetCFlag sets the bool value in the IsisSRv6NodeCapability object
func (obj *isisSRv6NodeCapability) SetCFlag(value bool) IsisSRv6NodeCapability {

	obj.obj.CFlag = &value
	return obj
}

func (obj *isisSRv6NodeCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.NodeMsds != nil {

		obj.NodeMsds().validateObj(vObj, set_default)
	}

}

func (obj *isisSRv6NodeCapability) setDefault() {
	if obj.obj.OFlag == nil {
		obj.SetOFlag(false)
	}
	if obj.obj.CFlag == nil {
		obj.SetCFlag(false)
	}

}
