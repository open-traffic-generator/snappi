package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingFlags *****
type flowIpv6SegmentRoutingFlags struct {
	validation
	obj          *otg.FlowIpv6SegmentRoutingFlags
	marshaller   marshalFlowIpv6SegmentRoutingFlags
	unMarshaller unMarshalFlowIpv6SegmentRoutingFlags
}

func NewFlowIpv6SegmentRoutingFlags() FlowIpv6SegmentRoutingFlags {
	obj := flowIpv6SegmentRoutingFlags{obj: &otg.FlowIpv6SegmentRoutingFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingFlags) msg() *otg.FlowIpv6SegmentRoutingFlags {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingFlags) setMsg(msg *otg.FlowIpv6SegmentRoutingFlags) FlowIpv6SegmentRoutingFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingFlags struct {
	obj *flowIpv6SegmentRoutingFlags
}

type marshalFlowIpv6SegmentRoutingFlags interface {
	// ToProto marshals FlowIpv6SegmentRoutingFlags to protobuf object *otg.FlowIpv6SegmentRoutingFlags
	ToProto() (*otg.FlowIpv6SegmentRoutingFlags, error)
	// ToPbText marshals FlowIpv6SegmentRoutingFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingFlags to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingFlags struct {
	obj *flowIpv6SegmentRoutingFlags
}

type unMarshalFlowIpv6SegmentRoutingFlags interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingFlags from protobuf object *otg.FlowIpv6SegmentRoutingFlags
	FromProto(msg *otg.FlowIpv6SegmentRoutingFlags) (FlowIpv6SegmentRoutingFlags, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingFlags from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingFlags) Marshal() marshalFlowIpv6SegmentRoutingFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingFlags) Unmarshal() unMarshalFlowIpv6SegmentRoutingFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingFlags) ToProto() (*otg.FlowIpv6SegmentRoutingFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingFlags) FromProto(msg *otg.FlowIpv6SegmentRoutingFlags) (FlowIpv6SegmentRoutingFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingFlags) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingFlags) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingFlags) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingFlags) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingFlags) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingFlags) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingFlags) Clone() (FlowIpv6SegmentRoutingFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingFlags()
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

// FlowIpv6SegmentRoutingFlags is sRH Flags field (RFC 8754 Section 2.1). An 8-bit field; RFC 8754 marks all bits as reserved. IxNetwork exposes the following bits for OAM, HMAC, FRR, and protocol testing: Protected (FRR), Alert, O (OAM, RFC 9259), H (HMAC), and two reserved bits U1 and U2.
type FlowIpv6SegmentRoutingFlags interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingFlags to protobuf object *otg.FlowIpv6SegmentRoutingFlags
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingFlags
	// setMsg unmarshals FlowIpv6SegmentRoutingFlags from protobuf object *otg.FlowIpv6SegmentRoutingFlags
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingFlags) FlowIpv6SegmentRoutingFlags
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingFlags
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingFlags
	// validate validates FlowIpv6SegmentRoutingFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Protected returns bool, set in FlowIpv6SegmentRoutingFlags.
	Protected() bool
	// SetProtected assigns bool provided by user to FlowIpv6SegmentRoutingFlags
	SetProtected(value bool) FlowIpv6SegmentRoutingFlags
	// HasProtected checks if Protected has been set in FlowIpv6SegmentRoutingFlags
	HasProtected() bool
	// Alert returns bool, set in FlowIpv6SegmentRoutingFlags.
	Alert() bool
	// SetAlert assigns bool provided by user to FlowIpv6SegmentRoutingFlags
	SetAlert(value bool) FlowIpv6SegmentRoutingFlags
	// HasAlert checks if Alert has been set in FlowIpv6SegmentRoutingFlags
	HasAlert() bool
	// OFlag returns bool, set in FlowIpv6SegmentRoutingFlags.
	OFlag() bool
	// SetOFlag assigns bool provided by user to FlowIpv6SegmentRoutingFlags
	SetOFlag(value bool) FlowIpv6SegmentRoutingFlags
	// HasOFlag checks if OFlag has been set in FlowIpv6SegmentRoutingFlags
	HasOFlag() bool
	// HFlag returns bool, set in FlowIpv6SegmentRoutingFlags.
	HFlag() bool
	// SetHFlag assigns bool provided by user to FlowIpv6SegmentRoutingFlags
	SetHFlag(value bool) FlowIpv6SegmentRoutingFlags
	// HasHFlag checks if HFlag has been set in FlowIpv6SegmentRoutingFlags
	HasHFlag() bool
	// U1Flag returns bool, set in FlowIpv6SegmentRoutingFlags.
	U1Flag() bool
	// SetU1Flag assigns bool provided by user to FlowIpv6SegmentRoutingFlags
	SetU1Flag(value bool) FlowIpv6SegmentRoutingFlags
	// HasU1Flag checks if U1Flag has been set in FlowIpv6SegmentRoutingFlags
	HasU1Flag() bool
	// U2Flag returns bool, set in FlowIpv6SegmentRoutingFlags.
	U2Flag() bool
	// SetU2Flag assigns bool provided by user to FlowIpv6SegmentRoutingFlags
	SetU2Flag(value bool) FlowIpv6SegmentRoutingFlags
	// HasU2Flag checks if U2Flag has been set in FlowIpv6SegmentRoutingFlags
	HasU2Flag() bool
}

// Protected (P) flag. Indicates that the segment is protected by a Fast Re-Route (FRR) mechanism (RFC 8754 Section 2.1).
// Protected returns a bool
func (obj *flowIpv6SegmentRoutingFlags) Protected() bool {

	return *obj.obj.Protected

}

// Protected (P) flag. Indicates that the segment is protected by a Fast Re-Route (FRR) mechanism (RFC 8754 Section 2.1).
// Protected returns a bool
func (obj *flowIpv6SegmentRoutingFlags) HasProtected() bool {
	return obj.obj.Protected != nil
}

// Protected (P) flag. Indicates that the segment is protected by a Fast Re-Route (FRR) mechanism (RFC 8754 Section 2.1).
// SetProtected sets the bool value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetProtected(value bool) FlowIpv6SegmentRoutingFlags {

	obj.obj.Protected = &value
	return obj
}

// Alert (A) flag. Indicates the presence of important TLVs in the SRH that must be examined by the receiving endpoint (RFC 8754 Section 2.1).
// Alert returns a bool
func (obj *flowIpv6SegmentRoutingFlags) Alert() bool {

	return *obj.obj.Alert

}

// Alert (A) flag. Indicates the presence of important TLVs in the SRH that must be examined by the receiving endpoint (RFC 8754 Section 2.1).
// Alert returns a bool
func (obj *flowIpv6SegmentRoutingFlags) HasAlert() bool {
	return obj.obj.Alert != nil
}

// Alert (A) flag. Indicates the presence of important TLVs in the SRH that must be examined by the receiving endpoint (RFC 8754 Section 2.1).
// SetAlert sets the bool value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetAlert(value bool) FlowIpv6SegmentRoutingFlags {

	obj.obj.Alert = &value
	return obj
}

// OAM (O) flag (RFC 9259 Section 2). When set, marks this packet as an active OAM probe. SRv6 endpoint nodes that process OAM packets inspect SRH TLVs such as the Ingress Node TLV (type 1) and Egress Node TLV (type 2) to perform path verification. Reference: RFC 8754 Section 2.1, RFC 9259 Section 2.
// OFlag returns a bool
func (obj *flowIpv6SegmentRoutingFlags) OFlag() bool {

	return *obj.obj.OFlag

}

// OAM (O) flag (RFC 9259 Section 2). When set, marks this packet as an active OAM probe. SRv6 endpoint nodes that process OAM packets inspect SRH TLVs such as the Ingress Node TLV (type 1) and Egress Node TLV (type 2) to perform path verification. Reference: RFC 8754 Section 2.1, RFC 9259 Section 2.
// OFlag returns a bool
func (obj *flowIpv6SegmentRoutingFlags) HasOFlag() bool {
	return obj.obj.OFlag != nil
}

// OAM (O) flag (RFC 9259 Section 2). When set, marks this packet as an active OAM probe. SRv6 endpoint nodes that process OAM packets inspect SRH TLVs such as the Ingress Node TLV (type 1) and Egress Node TLV (type 2) to perform path verification. Reference: RFC 8754 Section 2.1, RFC 9259 Section 2.
// SetOFlag sets the bool value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetOFlag(value bool) FlowIpv6SegmentRoutingFlags {

	obj.obj.OFlag = &value
	return obj
}

// HMAC (H) flag (RFC 8754 Section 2.1). When set, indicates that this SRH includes an HMAC TLV (type 5) providing data integrity and source authentication for the segment routing path. Reference: RFC 8754 Section 2.1.
// HFlag returns a bool
func (obj *flowIpv6SegmentRoutingFlags) HFlag() bool {

	return *obj.obj.HFlag

}

// HMAC (H) flag (RFC 8754 Section 2.1). When set, indicates that this SRH includes an HMAC TLV (type 5) providing data integrity and source authentication for the segment routing path. Reference: RFC 8754 Section 2.1.
// HFlag returns a bool
func (obj *flowIpv6SegmentRoutingFlags) HasHFlag() bool {
	return obj.obj.HFlag != nil
}

// HMAC (H) flag (RFC 8754 Section 2.1). When set, indicates that this SRH includes an HMAC TLV (type 5) providing data integrity and source authentication for the segment routing path. Reference: RFC 8754 Section 2.1.
// SetHFlag sets the bool value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetHFlag(value bool) FlowIpv6SegmentRoutingFlags {

	obj.obj.HFlag = &value
	return obj
}

// Reserved bit U1 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
// U1Flag returns a bool
func (obj *flowIpv6SegmentRoutingFlags) U1Flag() bool {

	return *obj.obj.U1Flag

}

// Reserved bit U1 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
// U1Flag returns a bool
func (obj *flowIpv6SegmentRoutingFlags) HasU1Flag() bool {
	return obj.obj.U1Flag != nil
}

// Reserved bit U1 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
// SetU1Flag sets the bool value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetU1Flag(value bool) FlowIpv6SegmentRoutingFlags {

	obj.obj.U1Flag = &value
	return obj
}

// Reserved bit U2 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
// U2Flag returns a bool
func (obj *flowIpv6SegmentRoutingFlags) U2Flag() bool {

	return *obj.obj.U2Flag

}

// Reserved bit U2 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
// U2Flag returns a bool
func (obj *flowIpv6SegmentRoutingFlags) HasU2Flag() bool {
	return obj.obj.U2Flag != nil
}

// Reserved bit U2 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
// SetU2Flag sets the bool value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetU2Flag(value bool) FlowIpv6SegmentRoutingFlags {

	obj.obj.U2Flag = &value
	return obj
}

func (obj *flowIpv6SegmentRoutingFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowIpv6SegmentRoutingFlags) setDefault() {
	if obj.obj.Protected == nil {
		obj.SetProtected(false)
	}
	if obj.obj.Alert == nil {
		obj.SetAlert(false)
	}
	if obj.obj.OFlag == nil {
		obj.SetOFlag(false)
	}
	if obj.obj.HFlag == nil {
		obj.SetHFlag(false)
	}
	if obj.obj.U1Flag == nil {
		obj.SetU1Flag(false)
	}
	if obj.obj.U2Flag == nil {
		obj.SetU2Flag(false)
	}

}
