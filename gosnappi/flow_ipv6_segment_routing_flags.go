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
	obj             *otg.FlowIpv6SegmentRoutingFlags
	marshaller      marshalFlowIpv6SegmentRoutingFlags
	unMarshaller    unMarshalFlowIpv6SegmentRoutingFlags
	protectedHolder PatternFlowIpv6SegmentRoutingFlagsProtected
	alertHolder     PatternFlowIpv6SegmentRoutingFlagsAlert
	oFlagHolder     PatternFlowIpv6SegmentRoutingFlagsOFlag
	hFlagHolder     PatternFlowIpv6SegmentRoutingFlagsHFlag
	u1FlagHolder    PatternFlowIpv6SegmentRoutingFlagsU1Flag
	u2FlagHolder    PatternFlowIpv6SegmentRoutingFlagsU2Flag
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
	obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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

func (obj *flowIpv6SegmentRoutingFlags) setNil() {
	obj.protectedHolder = nil
	obj.alertHolder = nil
	obj.oFlagHolder = nil
	obj.hFlagHolder = nil
	obj.u1FlagHolder = nil
	obj.u2FlagHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
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
	// Protected returns PatternFlowIpv6SegmentRoutingFlagsProtected, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsProtected is protected (P) flag. Indicates that the segment is protected by a Fast Re-Route (FRR) mechanism (RFC 8754 Section 2.1).
	Protected() PatternFlowIpv6SegmentRoutingFlagsProtected
	// SetProtected assigns PatternFlowIpv6SegmentRoutingFlagsProtected provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsProtected is protected (P) flag. Indicates that the segment is protected by a Fast Re-Route (FRR) mechanism (RFC 8754 Section 2.1).
	SetProtected(value PatternFlowIpv6SegmentRoutingFlagsProtected) FlowIpv6SegmentRoutingFlags
	// HasProtected checks if Protected has been set in FlowIpv6SegmentRoutingFlags
	HasProtected() bool
	// Alert returns PatternFlowIpv6SegmentRoutingFlagsAlert, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsAlert is alert (A) flag. Indicates the presence of important TLVs in the SRH that must be examined by the receiving endpoint (RFC 8754 Section 2.1).
	Alert() PatternFlowIpv6SegmentRoutingFlagsAlert
	// SetAlert assigns PatternFlowIpv6SegmentRoutingFlagsAlert provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsAlert is alert (A) flag. Indicates the presence of important TLVs in the SRH that must be examined by the receiving endpoint (RFC 8754 Section 2.1).
	SetAlert(value PatternFlowIpv6SegmentRoutingFlagsAlert) FlowIpv6SegmentRoutingFlags
	// HasAlert checks if Alert has been set in FlowIpv6SegmentRoutingFlags
	HasAlert() bool
	// OFlag returns PatternFlowIpv6SegmentRoutingFlagsOFlag, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsOFlag is oAM (O) flag (RFC 9259 Section 2). When set, marks this packet as an active OAM probe. SRv6 endpoint nodes that process OAM packets inspect SRH TLVs such as the Ingress Node TLV (type 1) and Egress Node TLV (type 2) to perform path verification. Reference: RFC 8754 Section 2.1, RFC 9259 Section 2.
	OFlag() PatternFlowIpv6SegmentRoutingFlagsOFlag
	// SetOFlag assigns PatternFlowIpv6SegmentRoutingFlagsOFlag provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsOFlag is oAM (O) flag (RFC 9259 Section 2). When set, marks this packet as an active OAM probe. SRv6 endpoint nodes that process OAM packets inspect SRH TLVs such as the Ingress Node TLV (type 1) and Egress Node TLV (type 2) to perform path verification. Reference: RFC 8754 Section 2.1, RFC 9259 Section 2.
	SetOFlag(value PatternFlowIpv6SegmentRoutingFlagsOFlag) FlowIpv6SegmentRoutingFlags
	// HasOFlag checks if OFlag has been set in FlowIpv6SegmentRoutingFlags
	HasOFlag() bool
	// HFlag returns PatternFlowIpv6SegmentRoutingFlagsHFlag, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsHFlag is hMAC (H) flag (RFC 8754 Section 2.1). When set, indicates that this SRH includes an HMAC TLV (type 5) providing data integrity and source authentication for the segment routing path. Reference: RFC 8754 Section 2.1.
	HFlag() PatternFlowIpv6SegmentRoutingFlagsHFlag
	// SetHFlag assigns PatternFlowIpv6SegmentRoutingFlagsHFlag provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsHFlag is hMAC (H) flag (RFC 8754 Section 2.1). When set, indicates that this SRH includes an HMAC TLV (type 5) providing data integrity and source authentication for the segment routing path. Reference: RFC 8754 Section 2.1.
	SetHFlag(value PatternFlowIpv6SegmentRoutingFlagsHFlag) FlowIpv6SegmentRoutingFlags
	// HasHFlag checks if HFlag has been set in FlowIpv6SegmentRoutingFlags
	HasHFlag() bool
	// U1Flag returns PatternFlowIpv6SegmentRoutingFlagsU1Flag, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsU1Flag is reserved bit U1 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
	U1Flag() PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// SetU1Flag assigns PatternFlowIpv6SegmentRoutingFlagsU1Flag provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsU1Flag is reserved bit U1 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
	SetU1Flag(value PatternFlowIpv6SegmentRoutingFlagsU1Flag) FlowIpv6SegmentRoutingFlags
	// HasU1Flag checks if U1Flag has been set in FlowIpv6SegmentRoutingFlags
	HasU1Flag() bool
	// U2Flag returns PatternFlowIpv6SegmentRoutingFlagsU2Flag, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsU2Flag is reserved bit U2 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
	U2Flag() PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// SetU2Flag assigns PatternFlowIpv6SegmentRoutingFlagsU2Flag provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsU2Flag is reserved bit U2 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
	SetU2Flag(value PatternFlowIpv6SegmentRoutingFlagsU2Flag) FlowIpv6SegmentRoutingFlags
	// HasU2Flag checks if U2Flag has been set in FlowIpv6SegmentRoutingFlags
	HasU2Flag() bool
	setNil()
}

// description is TBD
// Protected returns a PatternFlowIpv6SegmentRoutingFlagsProtected
func (obj *flowIpv6SegmentRoutingFlags) Protected() PatternFlowIpv6SegmentRoutingFlagsProtected {
	if obj.obj.Protected == nil {
		obj.obj.Protected = NewPatternFlowIpv6SegmentRoutingFlagsProtected().msg()
	}
	if obj.protectedHolder == nil {
		obj.protectedHolder = &patternFlowIpv6SegmentRoutingFlagsProtected{obj: obj.obj.Protected}
	}
	return obj.protectedHolder
}

// description is TBD
// Protected returns a PatternFlowIpv6SegmentRoutingFlagsProtected
func (obj *flowIpv6SegmentRoutingFlags) HasProtected() bool {
	return obj.obj.Protected != nil
}

// description is TBD
// SetProtected sets the PatternFlowIpv6SegmentRoutingFlagsProtected value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetProtected(value PatternFlowIpv6SegmentRoutingFlagsProtected) FlowIpv6SegmentRoutingFlags {

	obj.protectedHolder = nil
	obj.obj.Protected = value.msg()

	return obj
}

// description is TBD
// Alert returns a PatternFlowIpv6SegmentRoutingFlagsAlert
func (obj *flowIpv6SegmentRoutingFlags) Alert() PatternFlowIpv6SegmentRoutingFlagsAlert {
	if obj.obj.Alert == nil {
		obj.obj.Alert = NewPatternFlowIpv6SegmentRoutingFlagsAlert().msg()
	}
	if obj.alertHolder == nil {
		obj.alertHolder = &patternFlowIpv6SegmentRoutingFlagsAlert{obj: obj.obj.Alert}
	}
	return obj.alertHolder
}

// description is TBD
// Alert returns a PatternFlowIpv6SegmentRoutingFlagsAlert
func (obj *flowIpv6SegmentRoutingFlags) HasAlert() bool {
	return obj.obj.Alert != nil
}

// description is TBD
// SetAlert sets the PatternFlowIpv6SegmentRoutingFlagsAlert value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetAlert(value PatternFlowIpv6SegmentRoutingFlagsAlert) FlowIpv6SegmentRoutingFlags {

	obj.alertHolder = nil
	obj.obj.Alert = value.msg()

	return obj
}

// description is TBD
// OFlag returns a PatternFlowIpv6SegmentRoutingFlagsOFlag
func (obj *flowIpv6SegmentRoutingFlags) OFlag() PatternFlowIpv6SegmentRoutingFlagsOFlag {
	if obj.obj.OFlag == nil {
		obj.obj.OFlag = NewPatternFlowIpv6SegmentRoutingFlagsOFlag().msg()
	}
	if obj.oFlagHolder == nil {
		obj.oFlagHolder = &patternFlowIpv6SegmentRoutingFlagsOFlag{obj: obj.obj.OFlag}
	}
	return obj.oFlagHolder
}

// description is TBD
// OFlag returns a PatternFlowIpv6SegmentRoutingFlagsOFlag
func (obj *flowIpv6SegmentRoutingFlags) HasOFlag() bool {
	return obj.obj.OFlag != nil
}

// description is TBD
// SetOFlag sets the PatternFlowIpv6SegmentRoutingFlagsOFlag value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetOFlag(value PatternFlowIpv6SegmentRoutingFlagsOFlag) FlowIpv6SegmentRoutingFlags {

	obj.oFlagHolder = nil
	obj.obj.OFlag = value.msg()

	return obj
}

// description is TBD
// HFlag returns a PatternFlowIpv6SegmentRoutingFlagsHFlag
func (obj *flowIpv6SegmentRoutingFlags) HFlag() PatternFlowIpv6SegmentRoutingFlagsHFlag {
	if obj.obj.HFlag == nil {
		obj.obj.HFlag = NewPatternFlowIpv6SegmentRoutingFlagsHFlag().msg()
	}
	if obj.hFlagHolder == nil {
		obj.hFlagHolder = &patternFlowIpv6SegmentRoutingFlagsHFlag{obj: obj.obj.HFlag}
	}
	return obj.hFlagHolder
}

// description is TBD
// HFlag returns a PatternFlowIpv6SegmentRoutingFlagsHFlag
func (obj *flowIpv6SegmentRoutingFlags) HasHFlag() bool {
	return obj.obj.HFlag != nil
}

// description is TBD
// SetHFlag sets the PatternFlowIpv6SegmentRoutingFlagsHFlag value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetHFlag(value PatternFlowIpv6SegmentRoutingFlagsHFlag) FlowIpv6SegmentRoutingFlags {

	obj.hFlagHolder = nil
	obj.obj.HFlag = value.msg()

	return obj
}

// description is TBD
// U1Flag returns a PatternFlowIpv6SegmentRoutingFlagsU1Flag
func (obj *flowIpv6SegmentRoutingFlags) U1Flag() PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	if obj.obj.U1Flag == nil {
		obj.obj.U1Flag = NewPatternFlowIpv6SegmentRoutingFlagsU1Flag().msg()
	}
	if obj.u1FlagHolder == nil {
		obj.u1FlagHolder = &patternFlowIpv6SegmentRoutingFlagsU1Flag{obj: obj.obj.U1Flag}
	}
	return obj.u1FlagHolder
}

// description is TBD
// U1Flag returns a PatternFlowIpv6SegmentRoutingFlagsU1Flag
func (obj *flowIpv6SegmentRoutingFlags) HasU1Flag() bool {
	return obj.obj.U1Flag != nil
}

// description is TBD
// SetU1Flag sets the PatternFlowIpv6SegmentRoutingFlagsU1Flag value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetU1Flag(value PatternFlowIpv6SegmentRoutingFlagsU1Flag) FlowIpv6SegmentRoutingFlags {

	obj.u1FlagHolder = nil
	obj.obj.U1Flag = value.msg()

	return obj
}

// description is TBD
// U2Flag returns a PatternFlowIpv6SegmentRoutingFlagsU2Flag
func (obj *flowIpv6SegmentRoutingFlags) U2Flag() PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	if obj.obj.U2Flag == nil {
		obj.obj.U2Flag = NewPatternFlowIpv6SegmentRoutingFlagsU2Flag().msg()
	}
	if obj.u2FlagHolder == nil {
		obj.u2FlagHolder = &patternFlowIpv6SegmentRoutingFlagsU2Flag{obj: obj.obj.U2Flag}
	}
	return obj.u2FlagHolder
}

// description is TBD
// U2Flag returns a PatternFlowIpv6SegmentRoutingFlagsU2Flag
func (obj *flowIpv6SegmentRoutingFlags) HasU2Flag() bool {
	return obj.obj.U2Flag != nil
}

// description is TBD
// SetU2Flag sets the PatternFlowIpv6SegmentRoutingFlagsU2Flag value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetU2Flag(value PatternFlowIpv6SegmentRoutingFlagsU2Flag) FlowIpv6SegmentRoutingFlags {

	obj.u2FlagHolder = nil
	obj.obj.U2Flag = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutingFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Protected != nil {

		obj.Protected().validateObj(vObj, set_default)
	}

	if obj.obj.Alert != nil {

		obj.Alert().validateObj(vObj, set_default)
	}

	if obj.obj.OFlag != nil {

		obj.OFlag().validateObj(vObj, set_default)
	}

	if obj.obj.HFlag != nil {

		obj.HFlag().validateObj(vObj, set_default)
	}

	if obj.obj.U1Flag != nil {

		obj.U1Flag().validateObj(vObj, set_default)
	}

	if obj.obj.U2Flag != nil {

		obj.U2Flag().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRoutingFlags) setDefault() {

}
