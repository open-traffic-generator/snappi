package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingUsidSegment *****
type flowIpv6SegmentRoutingUsidSegment struct {
	validation
	obj                 *otg.FlowIpv6SegmentRoutingUsidSegment
	marshaller          marshalFlowIpv6SegmentRoutingUsidSegment
	unMarshaller        unMarshalFlowIpv6SegmentRoutingUsidSegment
	locatorHolder       PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	locatorLengthHolder PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
}

func NewFlowIpv6SegmentRoutingUsidSegment() FlowIpv6SegmentRoutingUsidSegment {
	obj := flowIpv6SegmentRoutingUsidSegment{obj: &otg.FlowIpv6SegmentRoutingUsidSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) msg() *otg.FlowIpv6SegmentRoutingUsidSegment {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) setMsg(msg *otg.FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingUsidSegment struct {
	obj *flowIpv6SegmentRoutingUsidSegment
}

type marshalFlowIpv6SegmentRoutingUsidSegment interface {
	// ToProto marshals FlowIpv6SegmentRoutingUsidSegment to protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	ToProto() (*otg.FlowIpv6SegmentRoutingUsidSegment, error)
	// ToPbText marshals FlowIpv6SegmentRoutingUsidSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingUsidSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingUsidSegment to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingUsidSegment struct {
	obj *flowIpv6SegmentRoutingUsidSegment
}

type unMarshalFlowIpv6SegmentRoutingUsidSegment interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	FromProto(msg *otg.FlowIpv6SegmentRoutingUsidSegment) (FlowIpv6SegmentRoutingUsidSegment, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingUsidSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingUsidSegment from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Marshal() marshalFlowIpv6SegmentRoutingUsidSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingUsidSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Unmarshal() unMarshalFlowIpv6SegmentRoutingUsidSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingUsidSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToProto() (*otg.FlowIpv6SegmentRoutingUsidSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromProto(msg *otg.FlowIpv6SegmentRoutingUsidSegment) (FlowIpv6SegmentRoutingUsidSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingUsidSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsidSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsidSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Clone() (FlowIpv6SegmentRoutingUsidSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingUsidSegment()
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

func (obj *flowIpv6SegmentRoutingUsidSegment) setNil() {
	obj.locatorHolder = nil
	obj.locatorLengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingUsidSegment is one CSID container entry in the SRH segment list (RFC 9800, RFC 8754 Section 2.1).
// Supports both NEXT-CSID (locator_length > 0) and REPLACE-CSID packed format
// (locator_length = 0) as described in Flow.Ipv6SegmentRoutingUsid.
//
// NEXT-CSID example (F3216: locator_length=32, 16-bit CSIDs):
// locator fc00::/32, usids ["0001","0002","0003"] -> SRH entry fc00:0:1:2:3::
//
// REPLACE-CSID packed example (LNFL=32, K=4, locator_length=0):
// usids ["00010002","00030004"] -> wire [0][0][00030004][00010002] (MSB->LSB)
// -> SRH entry ::3:4:1:2
type FlowIpv6SegmentRoutingUsidSegment interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingUsidSegment to protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingUsidSegment
	// setMsg unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidSegment
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingUsidSegment
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingUsidSegment
	// validate validates FlowIpv6SegmentRoutingUsidSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingUsidSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Locator returns PatternFlowIpv6SegmentRoutingUsidSegmentLocator, set in FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits of every uSID assembled from this entry. For F3216, this is a /32 prefix (e.g., fc00::). The prefix length is given by locator_length. Ignored when locator_length is 0 (REPLACE-CSID flavor, RFC 9800 Section 5).
	Locator() PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// SetLocator assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocator provided by user to FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits of every uSID assembled from this entry. For F3216, this is a /32 prefix (e.g., fc00::). The prefix length is given by locator_length. Ignored when locator_length is 0 (REPLACE-CSID flavor, RFC 9800 Section 5).
	SetLocator(value PatternFlowIpv6SegmentRoutingUsidSegmentLocator) FlowIpv6SegmentRoutingUsidSegment
	// HasLocator checks if Locator has been set in FlowIpv6SegmentRoutingUsidSegment
	HasLocator() bool
	// LocatorLength returns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength, set in FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Valid range: 0-112. For NEXT-CSID (locator_length > 0): high-order locator_length bits of locator form the LB; usids are packed after it in forward order. For REPLACE-CSID first container (locator_length > 0): same as NEXT-CSID structure; use a single usid = Locator-Node+Function value (LNFL bits). For REPLACE-CSID packed containers (locator_length = 0): the locator field is ignored; the 128-bit SRH entry is K = floor(128/LNFL) slots of LNFL bits each; usids are packed from the LSB (RFC 9800 Section 4.2).
	LocatorLength() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// SetLocatorLength assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength provided by user to FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Valid range: 0-112. For NEXT-CSID (locator_length > 0): high-order locator_length bits of locator form the LB; usids are packed after it in forward order. For REPLACE-CSID first container (locator_length > 0): same as NEXT-CSID structure; use a single usid = Locator-Node+Function value (LNFL bits). For REPLACE-CSID packed containers (locator_length = 0): the locator field is ignored; the 128-bit SRH entry is K = floor(128/LNFL) slots of LNFL bits each; usids are packed from the LSB (RFC 9800 Section 4.2).
	SetLocatorLength(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) FlowIpv6SegmentRoutingUsidSegment
	// HasLocatorLength checks if LocatorLength has been set in FlowIpv6SegmentRoutingUsidSegment
	HasLocatorLength() bool
	// Usids returns []string, set in FlowIpv6SegmentRoutingUsidSegment.
	Usids() []string
	// SetUsids assigns []string provided by user to FlowIpv6SegmentRoutingUsidSegment
	SetUsids(value []string) FlowIpv6SegmentRoutingUsidSegment
	setNil()
}

// description is TBD
// Locator returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocator
func (obj *flowIpv6SegmentRoutingUsidSegment) Locator() PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	if obj.obj.Locator == nil {
		obj.obj.Locator = NewPatternFlowIpv6SegmentRoutingUsidSegmentLocator().msg()
	}
	if obj.locatorHolder == nil {
		obj.locatorHolder = &patternFlowIpv6SegmentRoutingUsidSegmentLocator{obj: obj.obj.Locator}
	}
	return obj.locatorHolder
}

// description is TBD
// Locator returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocator
func (obj *flowIpv6SegmentRoutingUsidSegment) HasLocator() bool {
	return obj.obj.Locator != nil
}

// description is TBD
// SetLocator sets the PatternFlowIpv6SegmentRoutingUsidSegmentLocator value in the FlowIpv6SegmentRoutingUsidSegment object
func (obj *flowIpv6SegmentRoutingUsidSegment) SetLocator(value PatternFlowIpv6SegmentRoutingUsidSegmentLocator) FlowIpv6SegmentRoutingUsidSegment {

	obj.locatorHolder = nil
	obj.obj.Locator = value.msg()

	return obj
}

// description is TBD
// LocatorLength returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
func (obj *flowIpv6SegmentRoutingUsidSegment) LocatorLength() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	if obj.obj.LocatorLength == nil {
		obj.obj.LocatorLength = NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength().msg()
	}
	if obj.locatorLengthHolder == nil {
		obj.locatorLengthHolder = &patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength{obj: obj.obj.LocatorLength}
	}
	return obj.locatorLengthHolder
}

// description is TBD
// LocatorLength returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
func (obj *flowIpv6SegmentRoutingUsidSegment) HasLocatorLength() bool {
	return obj.obj.LocatorLength != nil
}

// description is TBD
// SetLocatorLength sets the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength value in the FlowIpv6SegmentRoutingUsidSegment object
func (obj *flowIpv6SegmentRoutingUsidSegment) SetLocatorLength(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) FlowIpv6SegmentRoutingUsidSegment {

	obj.locatorLengthHolder = nil
	obj.obj.LocatorLength = value.msg()

	return obj
}

// Ordered list of CSID hex strings for this container (RFC 9800 Section 3). For NEXT-CSID (locator_length > 0): CSIDs are packed in forward order after the Locator Block. usids[0] immediately follows the LB, usids[1] follows, etc. The value 0x0 (all zeros) is reserved as the End-of-Container marker. Example F3216: ["0001","0002","0003"] with locator fc00::/32 -> fc00:0:1:2:3:: For REPLACE-CSID packed format (locator_length = 0): CSIDs are packed in reverse order from the LSB (RFC 9800 Section 4.2, Figure 4). K = floor(128 / LNFL) where LNFL is inferred from hex string width (8 chars = 32-bit CSID, K = 4). usids[0] goes to position K-1 (LSB, bits [128-LNFL..127]), usids[1] to position K-2, etc. Unused MSB positions are zero-padded. Example LNFL=32: ["00010002","00030004"] -> wire [0][0][00030004][00010002] (MSB to LSB) -> SRH entry ::3:4:1:2
// Usids returns a []string
func (obj *flowIpv6SegmentRoutingUsidSegment) Usids() []string {
	if obj.obj.Usids == nil {
		obj.obj.Usids = make([]string, 0)
	}
	return obj.obj.Usids
}

// Ordered list of CSID hex strings for this container (RFC 9800 Section 3). For NEXT-CSID (locator_length > 0): CSIDs are packed in forward order after the Locator Block. usids[0] immediately follows the LB, usids[1] follows, etc. The value 0x0 (all zeros) is reserved as the End-of-Container marker. Example F3216: ["0001","0002","0003"] with locator fc00::/32 -> fc00:0:1:2:3:: For REPLACE-CSID packed format (locator_length = 0): CSIDs are packed in reverse order from the LSB (RFC 9800 Section 4.2, Figure 4). K = floor(128 / LNFL) where LNFL is inferred from hex string width (8 chars = 32-bit CSID, K = 4). usids[0] goes to position K-1 (LSB, bits [128-LNFL..127]), usids[1] to position K-2, etc. Unused MSB positions are zero-padded. Example LNFL=32: ["00010002","00030004"] -> wire [0][0][00030004][00010002] (MSB to LSB) -> SRH entry ::3:4:1:2
// SetUsids sets the []string value in the FlowIpv6SegmentRoutingUsidSegment object
func (obj *flowIpv6SegmentRoutingUsidSegment) SetUsids(value []string) FlowIpv6SegmentRoutingUsidSegment {

	if obj.obj.Usids == nil {
		obj.obj.Usids = make([]string, 0)
	}
	obj.obj.Usids = value

	return obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Locator != nil {

		obj.Locator().validateObj(vObj, set_default)
	}

	if obj.obj.LocatorLength != nil {

		obj.LocatorLength().validateObj(vObj, set_default)
	}

	if obj.obj.Usids != nil {

		err := obj.validateHexSlice(obj.Usids())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SegmentRoutingUsidSegment.Usids"))
		}

	}

}

func (obj *flowIpv6SegmentRoutingUsidSegment) setDefault() {

}
