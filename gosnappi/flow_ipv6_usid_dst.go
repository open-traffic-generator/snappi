package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6UsidDst *****
type flowIpv6UsidDst struct {
	validation
	obj                 *otg.FlowIpv6UsidDst
	marshaller          marshalFlowIpv6UsidDst
	unMarshaller        unMarshalFlowIpv6UsidDst
	locatorHolder       PatternFlowIpv6UsidDstLocator
	locatorLengthHolder PatternFlowIpv6UsidDstLocatorLength
}

func NewFlowIpv6UsidDst() FlowIpv6UsidDst {
	obj := flowIpv6UsidDst{obj: &otg.FlowIpv6UsidDst{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6UsidDst) msg() *otg.FlowIpv6UsidDst {
	return obj.obj
}

func (obj *flowIpv6UsidDst) setMsg(msg *otg.FlowIpv6UsidDst) FlowIpv6UsidDst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6UsidDst struct {
	obj *flowIpv6UsidDst
}

type marshalFlowIpv6UsidDst interface {
	// ToProto marshals FlowIpv6UsidDst to protobuf object *otg.FlowIpv6UsidDst
	ToProto() (*otg.FlowIpv6UsidDst, error)
	// ToPbText marshals FlowIpv6UsidDst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6UsidDst to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6UsidDst to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6UsidDst struct {
	obj *flowIpv6UsidDst
}

type unMarshalFlowIpv6UsidDst interface {
	// FromProto unmarshals FlowIpv6UsidDst from protobuf object *otg.FlowIpv6UsidDst
	FromProto(msg *otg.FlowIpv6UsidDst) (FlowIpv6UsidDst, error)
	// FromPbText unmarshals FlowIpv6UsidDst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6UsidDst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6UsidDst from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6UsidDst) Marshal() marshalFlowIpv6UsidDst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6UsidDst{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6UsidDst) Unmarshal() unMarshalFlowIpv6UsidDst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6UsidDst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6UsidDst) ToProto() (*otg.FlowIpv6UsidDst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6UsidDst) FromProto(msg *otg.FlowIpv6UsidDst) (FlowIpv6UsidDst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6UsidDst) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6UsidDst) FromPbText(value string) error {
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

func (m *marshalflowIpv6UsidDst) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6UsidDst) FromYaml(value string) error {
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

func (m *marshalflowIpv6UsidDst) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6UsidDst) FromJson(value string) error {
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

func (obj *flowIpv6UsidDst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6UsidDst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6UsidDst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6UsidDst) Clone() (FlowIpv6UsidDst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6UsidDst()
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

func (obj *flowIpv6UsidDst) setNil() {
	obj.locatorHolder = nil
	obj.locatorLengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6UsidDst is structured input for the SRv6 uSID reduced encapsulation case (RFC 9800 Section 4) where the entire SR path fits in a single 128-bit uSID container placed directly in the outer IPv6 dst, with no Segment Routing Header. The implementation packs the fields as: LB (locator_length high-order bits of locator) || uSID-1 || uSID-2 || ... || EoC (zero-pad to 128 bits).
// For F3216 format (RFC 9800 Section 3): LB = 32 bits, each uSID = 16 bits, up to 6 uSIDs per container. Example - locator fc00::/32, locator_length 32, usids ["0001","0002","0003"]: assembled dst = fc00:0:1:2:3::
// When this field is present, ipv6.dst should be ignored by the implementation.
type FlowIpv6UsidDst interface {
	Validation
	// msg marshals FlowIpv6UsidDst to protobuf object *otg.FlowIpv6UsidDst
	// and doesn't set defaults
	msg() *otg.FlowIpv6UsidDst
	// setMsg unmarshals FlowIpv6UsidDst from protobuf object *otg.FlowIpv6UsidDst
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6UsidDst) FlowIpv6UsidDst
	// provides marshal interface
	Marshal() marshalFlowIpv6UsidDst
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6UsidDst
	// validate validates FlowIpv6UsidDst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6UsidDst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Locator returns PatternFlowIpv6UsidDstLocator, set in FlowIpv6UsidDst.
	// PatternFlowIpv6UsidDstLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits packed into the 128-bit dst. For F3216 this is a /32 prefix (e.g., fc00::). The number of bits used is given by locator_length.
	Locator() PatternFlowIpv6UsidDstLocator
	// SetLocator assigns PatternFlowIpv6UsidDstLocator provided by user to FlowIpv6UsidDst.
	// PatternFlowIpv6UsidDstLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits packed into the 128-bit dst. For F3216 this is a /32 prefix (e.g., fc00::). The number of bits used is given by locator_length.
	SetLocator(value PatternFlowIpv6UsidDstLocator) FlowIpv6UsidDst
	// LocatorLength returns PatternFlowIpv6UsidDstLocatorLength, set in FlowIpv6UsidDst.
	// PatternFlowIpv6UsidDstLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Determines how many high-order bits of locator are used as the LB and how many bits remain for CSID packing. Valid range: 1-112. For NEXT-CSID F3216 (RFC 9800 Section 3): 32. For REPLACE-CSID first container (RFC 9800 Section 4.2): use the LB length (e.g. 32 or 48); the outer IPv6 DA always carries a fully formed SRv6 SID with a LB.
	LocatorLength() PatternFlowIpv6UsidDstLocatorLength
	// SetLocatorLength assigns PatternFlowIpv6UsidDstLocatorLength provided by user to FlowIpv6UsidDst.
	// PatternFlowIpv6UsidDstLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Determines how many high-order bits of locator are used as the LB and how many bits remain for CSID packing. Valid range: 1-112. For NEXT-CSID F3216 (RFC 9800 Section 3): 32. For REPLACE-CSID first container (RFC 9800 Section 4.2): use the LB length (e.g. 32 or 48); the outer IPv6 DA always carries a fully formed SRv6 SID with a LB.
	SetLocatorLength(value PatternFlowIpv6UsidDstLocatorLength) FlowIpv6UsidDst
	// HasLocatorLength checks if LocatorLength has been set in FlowIpv6UsidDst
	HasLocatorLength() bool
	// Usids returns []string, set in FlowIpv6UsidDst.
	Usids() []string
	// SetUsids assigns []string provided by user to FlowIpv6UsidDst
	SetUsids(value []string) FlowIpv6UsidDst
	setNil()
}

// description is TBD
// Locator returns a PatternFlowIpv6UsidDstLocator
func (obj *flowIpv6UsidDst) Locator() PatternFlowIpv6UsidDstLocator {
	if obj.obj.Locator == nil {
		obj.obj.Locator = NewPatternFlowIpv6UsidDstLocator().msg()
	}
	if obj.locatorHolder == nil {
		obj.locatorHolder = &patternFlowIpv6UsidDstLocator{obj: obj.obj.Locator}
	}
	return obj.locatorHolder
}

// description is TBD
// SetLocator sets the PatternFlowIpv6UsidDstLocator value in the FlowIpv6UsidDst object
func (obj *flowIpv6UsidDst) SetLocator(value PatternFlowIpv6UsidDstLocator) FlowIpv6UsidDst {

	obj.locatorHolder = nil
	obj.obj.Locator = value.msg()

	return obj
}

// description is TBD
// LocatorLength returns a PatternFlowIpv6UsidDstLocatorLength
func (obj *flowIpv6UsidDst) LocatorLength() PatternFlowIpv6UsidDstLocatorLength {
	if obj.obj.LocatorLength == nil {
		obj.obj.LocatorLength = NewPatternFlowIpv6UsidDstLocatorLength().msg()
	}
	if obj.locatorLengthHolder == nil {
		obj.locatorLengthHolder = &patternFlowIpv6UsidDstLocatorLength{obj: obj.obj.LocatorLength}
	}
	return obj.locatorLengthHolder
}

// description is TBD
// LocatorLength returns a PatternFlowIpv6UsidDstLocatorLength
func (obj *flowIpv6UsidDst) HasLocatorLength() bool {
	return obj.obj.LocatorLength != nil
}

// description is TBD
// SetLocatorLength sets the PatternFlowIpv6UsidDstLocatorLength value in the FlowIpv6UsidDst object
func (obj *flowIpv6UsidDst) SetLocatorLength(value PatternFlowIpv6UsidDstLocatorLength) FlowIpv6UsidDst {

	obj.locatorLengthHolder = nil
	obj.obj.LocatorLength = value.msg()

	return obj
}

// Ordered list of uSID values (hex strings) to pack after the Locator Block (RFC 9800 Section 3). For F3216 each uSID is 16 bits (4 hex chars); up to 6 uSIDs fit per container. The End-of-Container zero-pad is appended automatically by the implementation. The value 0x0000 is reserved as the End-of-Container marker and must not be used. Example: ["0001","0002","0003"] with locator fc00::/32 assembles to fc00:0:1:2:3::
// Usids returns a []string
func (obj *flowIpv6UsidDst) Usids() []string {
	if obj.obj.Usids == nil {
		obj.obj.Usids = make([]string, 0)
	}
	return obj.obj.Usids
}

// Ordered list of uSID values (hex strings) to pack after the Locator Block (RFC 9800 Section 3). For F3216 each uSID is 16 bits (4 hex chars); up to 6 uSIDs fit per container. The End-of-Container zero-pad is appended automatically by the implementation. The value 0x0000 is reserved as the End-of-Container marker and must not be used. Example: ["0001","0002","0003"] with locator fc00::/32 assembles to fc00:0:1:2:3::
// SetUsids sets the []string value in the FlowIpv6UsidDst object
func (obj *flowIpv6UsidDst) SetUsids(value []string) FlowIpv6UsidDst {

	if obj.obj.Usids == nil {
		obj.obj.Usids = make([]string, 0)
	}
	obj.obj.Usids = value

	return obj
}

func (obj *flowIpv6UsidDst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Locator is required
	if obj.obj.Locator == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Locator is required field on interface FlowIpv6UsidDst")
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
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6UsidDst.Usids"))
		}

	}

}

func (obj *flowIpv6UsidDst) setDefault() {

}
