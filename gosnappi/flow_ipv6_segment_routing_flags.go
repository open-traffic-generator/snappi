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
	oamHolder       PatternFlowIpv6SegmentRoutingFlagsOam
	alertHolder     PatternFlowIpv6SegmentRoutingFlagsAlert
	hmacHolder      PatternFlowIpv6SegmentRoutingFlagsHmac
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
	obj.oamHolder = nil
	obj.alertHolder = nil
	obj.hmacHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingFlags is an 8-bit field containing flags. While RFC 8754 reserves all bits as unused, earlier drafts defined specific flags for behavior such as OAM, HMAC, and FRR protection.
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
	// PatternFlowIpv6SegmentRoutingFlagsProtected is protected Flag. Indicates if the segment is protected by a Fast Re-Route (FRR) mechanism.
	Protected() PatternFlowIpv6SegmentRoutingFlagsProtected
	// SetProtected assigns PatternFlowIpv6SegmentRoutingFlagsProtected provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsProtected is protected Flag. Indicates if the segment is protected by a Fast Re-Route (FRR) mechanism.
	SetProtected(value PatternFlowIpv6SegmentRoutingFlagsProtected) FlowIpv6SegmentRoutingFlags
	// HasProtected checks if Protected has been set in FlowIpv6SegmentRoutingFlags
	HasProtected() bool
	// Oam returns PatternFlowIpv6SegmentRoutingFlagsOam, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsOam is oAM Flag. Indicates if the packet is an Operations, Administration, and Maintenance (OAM) packet.
	Oam() PatternFlowIpv6SegmentRoutingFlagsOam
	// SetOam assigns PatternFlowIpv6SegmentRoutingFlagsOam provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsOam is oAM Flag. Indicates if the packet is an Operations, Administration, and Maintenance (OAM) packet.
	SetOam(value PatternFlowIpv6SegmentRoutingFlagsOam) FlowIpv6SegmentRoutingFlags
	// HasOam checks if Oam has been set in FlowIpv6SegmentRoutingFlags
	HasOam() bool
	// Alert returns PatternFlowIpv6SegmentRoutingFlagsAlert, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsAlert is alert Flag. Indicates the presence of important TLVs that must be inspected.
	Alert() PatternFlowIpv6SegmentRoutingFlagsAlert
	// SetAlert assigns PatternFlowIpv6SegmentRoutingFlagsAlert provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsAlert is alert Flag. Indicates the presence of important TLVs that must be inspected.
	SetAlert(value PatternFlowIpv6SegmentRoutingFlagsAlert) FlowIpv6SegmentRoutingFlags
	// HasAlert checks if Alert has been set in FlowIpv6SegmentRoutingFlags
	HasAlert() bool
	// Hmac returns PatternFlowIpv6SegmentRoutingFlagsHmac, set in FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsHmac is hMAC Flag. Indicates the presence of the HMAC TLV for verification.
	Hmac() PatternFlowIpv6SegmentRoutingFlagsHmac
	// SetHmac assigns PatternFlowIpv6SegmentRoutingFlagsHmac provided by user to FlowIpv6SegmentRoutingFlags.
	// PatternFlowIpv6SegmentRoutingFlagsHmac is hMAC Flag. Indicates the presence of the HMAC TLV for verification.
	SetHmac(value PatternFlowIpv6SegmentRoutingFlagsHmac) FlowIpv6SegmentRoutingFlags
	// HasHmac checks if Hmac has been set in FlowIpv6SegmentRoutingFlags
	HasHmac() bool
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
// Oam returns a PatternFlowIpv6SegmentRoutingFlagsOam
func (obj *flowIpv6SegmentRoutingFlags) Oam() PatternFlowIpv6SegmentRoutingFlagsOam {
	if obj.obj.Oam == nil {
		obj.obj.Oam = NewPatternFlowIpv6SegmentRoutingFlagsOam().msg()
	}
	if obj.oamHolder == nil {
		obj.oamHolder = &patternFlowIpv6SegmentRoutingFlagsOam{obj: obj.obj.Oam}
	}
	return obj.oamHolder
}

// description is TBD
// Oam returns a PatternFlowIpv6SegmentRoutingFlagsOam
func (obj *flowIpv6SegmentRoutingFlags) HasOam() bool {
	return obj.obj.Oam != nil
}

// description is TBD
// SetOam sets the PatternFlowIpv6SegmentRoutingFlagsOam value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetOam(value PatternFlowIpv6SegmentRoutingFlagsOam) FlowIpv6SegmentRoutingFlags {

	obj.oamHolder = nil
	obj.obj.Oam = value.msg()

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
// Hmac returns a PatternFlowIpv6SegmentRoutingFlagsHmac
func (obj *flowIpv6SegmentRoutingFlags) Hmac() PatternFlowIpv6SegmentRoutingFlagsHmac {
	if obj.obj.Hmac == nil {
		obj.obj.Hmac = NewPatternFlowIpv6SegmentRoutingFlagsHmac().msg()
	}
	if obj.hmacHolder == nil {
		obj.hmacHolder = &patternFlowIpv6SegmentRoutingFlagsHmac{obj: obj.obj.Hmac}
	}
	return obj.hmacHolder
}

// description is TBD
// Hmac returns a PatternFlowIpv6SegmentRoutingFlagsHmac
func (obj *flowIpv6SegmentRoutingFlags) HasHmac() bool {
	return obj.obj.Hmac != nil
}

// description is TBD
// SetHmac sets the PatternFlowIpv6SegmentRoutingFlagsHmac value in the FlowIpv6SegmentRoutingFlags object
func (obj *flowIpv6SegmentRoutingFlags) SetHmac(value PatternFlowIpv6SegmentRoutingFlagsHmac) FlowIpv6SegmentRoutingFlags {

	obj.hmacHolder = nil
	obj.obj.Hmac = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutingFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Protected != nil {

		obj.Protected().validateObj(vObj, set_default)
	}

	if obj.obj.Oam != nil {

		obj.Oam().validateObj(vObj, set_default)
	}

	if obj.obj.Alert != nil {

		obj.Alert().validateObj(vObj, set_default)
	}

	if obj.obj.Hmac != nil {

		obj.Hmac().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRoutingFlags) setDefault() {

}
