package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6 *****
type flowIpv6 struct {
	validation
	obj                 *otg.FlowIpv6
	marshaller          marshalFlowIpv6
	unMarshaller        unMarshalFlowIpv6
	versionHolder       PatternFlowIpv6Version
	trafficClassHolder  PatternFlowIpv6TrafficClass
	flowLabelHolder     PatternFlowIpv6FlowLabel
	payloadLengthHolder PatternFlowIpv6PayloadLength
	nextHeaderHolder    PatternFlowIpv6NextHeader
	hopLimitHolder      PatternFlowIpv6HopLimit
	srcHolder           PatternFlowIpv6Src
	dstHolder           PatternFlowIpv6Dst
}

func NewFlowIpv6() FlowIpv6 {
	obj := flowIpv6{obj: &otg.FlowIpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6) msg() *otg.FlowIpv6 {
	return obj.obj
}

func (obj *flowIpv6) setMsg(msg *otg.FlowIpv6) FlowIpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6 struct {
	obj *flowIpv6
}

type marshalFlowIpv6 interface {
	// ToProto marshals FlowIpv6 to protobuf object *otg.FlowIpv6
	ToProto() (*otg.FlowIpv6, error)
	// ToPbText marshals FlowIpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowIpv6 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowIpv6 struct {
	obj *flowIpv6
}

type unMarshalFlowIpv6 interface {
	// FromProto unmarshals FlowIpv6 from protobuf object *otg.FlowIpv6
	FromProto(msg *otg.FlowIpv6) (FlowIpv6, error)
	// FromPbText unmarshals FlowIpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6 from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6) Marshal() marshalFlowIpv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6) Unmarshal() unMarshalFlowIpv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6) ToProto() (*otg.FlowIpv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6) FromProto(msg *otg.FlowIpv6) (FlowIpv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6) FromPbText(value string) error {
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

func (m *marshalflowIpv6) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6) FromYaml(value string) error {
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

func (m *marshalflowIpv6) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowIpv6) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6) FromJson(value string) error {
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

func (obj *flowIpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6) Clone() (FlowIpv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6()
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

func (obj *flowIpv6) setNil() {
	obj.versionHolder = nil
	obj.trafficClassHolder = nil
	obj.flowLabelHolder = nil
	obj.payloadLengthHolder = nil
	obj.nextHeaderHolder = nil
	obj.hopLimitHolder = nil
	obj.srcHolder = nil
	obj.dstHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6 is iPv6 packet header
type FlowIpv6 interface {
	Validation
	// msg marshals FlowIpv6 to protobuf object *otg.FlowIpv6
	// and doesn't set defaults
	msg() *otg.FlowIpv6
	// setMsg unmarshals FlowIpv6 from protobuf object *otg.FlowIpv6
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6) FlowIpv6
	// provides marshal interface
	Marshal() marshalFlowIpv6
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6
	// validate validates FlowIpv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns PatternFlowIpv6Version, set in FlowIpv6.
	// PatternFlowIpv6Version is version number
	Version() PatternFlowIpv6Version
	// SetVersion assigns PatternFlowIpv6Version provided by user to FlowIpv6.
	// PatternFlowIpv6Version is version number
	SetVersion(value PatternFlowIpv6Version) FlowIpv6
	// HasVersion checks if Version has been set in FlowIpv6
	HasVersion() bool
	// TrafficClass returns PatternFlowIpv6TrafficClass, set in FlowIpv6.
	// PatternFlowIpv6TrafficClass is traffic class
	TrafficClass() PatternFlowIpv6TrafficClass
	// SetTrafficClass assigns PatternFlowIpv6TrafficClass provided by user to FlowIpv6.
	// PatternFlowIpv6TrafficClass is traffic class
	SetTrafficClass(value PatternFlowIpv6TrafficClass) FlowIpv6
	// HasTrafficClass checks if TrafficClass has been set in FlowIpv6
	HasTrafficClass() bool
	// FlowLabel returns PatternFlowIpv6FlowLabel, set in FlowIpv6.
	// PatternFlowIpv6FlowLabel is flow label
	FlowLabel() PatternFlowIpv6FlowLabel
	// SetFlowLabel assigns PatternFlowIpv6FlowLabel provided by user to FlowIpv6.
	// PatternFlowIpv6FlowLabel is flow label
	SetFlowLabel(value PatternFlowIpv6FlowLabel) FlowIpv6
	// HasFlowLabel checks if FlowLabel has been set in FlowIpv6
	HasFlowLabel() bool
	// PayloadLength returns PatternFlowIpv6PayloadLength, set in FlowIpv6.
	// PatternFlowIpv6PayloadLength is payload length
	PayloadLength() PatternFlowIpv6PayloadLength
	// SetPayloadLength assigns PatternFlowIpv6PayloadLength provided by user to FlowIpv6.
	// PatternFlowIpv6PayloadLength is payload length
	SetPayloadLength(value PatternFlowIpv6PayloadLength) FlowIpv6
	// HasPayloadLength checks if PayloadLength has been set in FlowIpv6
	HasPayloadLength() bool
	// NextHeader returns PatternFlowIpv6NextHeader, set in FlowIpv6.
	// PatternFlowIpv6NextHeader is next header
	NextHeader() PatternFlowIpv6NextHeader
	// SetNextHeader assigns PatternFlowIpv6NextHeader provided by user to FlowIpv6.
	// PatternFlowIpv6NextHeader is next header
	SetNextHeader(value PatternFlowIpv6NextHeader) FlowIpv6
	// HasNextHeader checks if NextHeader has been set in FlowIpv6
	HasNextHeader() bool
	// HopLimit returns PatternFlowIpv6HopLimit, set in FlowIpv6.
	// PatternFlowIpv6HopLimit is hop limit
	HopLimit() PatternFlowIpv6HopLimit
	// SetHopLimit assigns PatternFlowIpv6HopLimit provided by user to FlowIpv6.
	// PatternFlowIpv6HopLimit is hop limit
	SetHopLimit(value PatternFlowIpv6HopLimit) FlowIpv6
	// HasHopLimit checks if HopLimit has been set in FlowIpv6
	HasHopLimit() bool
	// Src returns PatternFlowIpv6Src, set in FlowIpv6.
	// PatternFlowIpv6Src is source address
	Src() PatternFlowIpv6Src
	// SetSrc assigns PatternFlowIpv6Src provided by user to FlowIpv6.
	// PatternFlowIpv6Src is source address
	SetSrc(value PatternFlowIpv6Src) FlowIpv6
	// HasSrc checks if Src has been set in FlowIpv6
	HasSrc() bool
	// Dst returns PatternFlowIpv6Dst, set in FlowIpv6.
	// PatternFlowIpv6Dst is destination address
	Dst() PatternFlowIpv6Dst
	// SetDst assigns PatternFlowIpv6Dst provided by user to FlowIpv6.
	// PatternFlowIpv6Dst is destination address
	SetDst(value PatternFlowIpv6Dst) FlowIpv6
	// HasDst checks if Dst has been set in FlowIpv6
	HasDst() bool
	setNil()
}

// description is TBD
// Version returns a PatternFlowIpv6Version
func (obj *flowIpv6) Version() PatternFlowIpv6Version {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowIpv6Version().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowIpv6Version{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowIpv6Version
func (obj *flowIpv6) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowIpv6Version value in the FlowIpv6 object
func (obj *flowIpv6) SetVersion(value PatternFlowIpv6Version) FlowIpv6 {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// TrafficClass returns a PatternFlowIpv6TrafficClass
func (obj *flowIpv6) TrafficClass() PatternFlowIpv6TrafficClass {
	if obj.obj.TrafficClass == nil {
		obj.obj.TrafficClass = NewPatternFlowIpv6TrafficClass().msg()
	}
	if obj.trafficClassHolder == nil {
		obj.trafficClassHolder = &patternFlowIpv6TrafficClass{obj: obj.obj.TrafficClass}
	}
	return obj.trafficClassHolder
}

// description is TBD
// TrafficClass returns a PatternFlowIpv6TrafficClass
func (obj *flowIpv6) HasTrafficClass() bool {
	return obj.obj.TrafficClass != nil
}

// description is TBD
// SetTrafficClass sets the PatternFlowIpv6TrafficClass value in the FlowIpv6 object
func (obj *flowIpv6) SetTrafficClass(value PatternFlowIpv6TrafficClass) FlowIpv6 {

	obj.trafficClassHolder = nil
	obj.obj.TrafficClass = value.msg()

	return obj
}

// description is TBD
// FlowLabel returns a PatternFlowIpv6FlowLabel
func (obj *flowIpv6) FlowLabel() PatternFlowIpv6FlowLabel {
	if obj.obj.FlowLabel == nil {
		obj.obj.FlowLabel = NewPatternFlowIpv6FlowLabel().msg()
	}
	if obj.flowLabelHolder == nil {
		obj.flowLabelHolder = &patternFlowIpv6FlowLabel{obj: obj.obj.FlowLabel}
	}
	return obj.flowLabelHolder
}

// description is TBD
// FlowLabel returns a PatternFlowIpv6FlowLabel
func (obj *flowIpv6) HasFlowLabel() bool {
	return obj.obj.FlowLabel != nil
}

// description is TBD
// SetFlowLabel sets the PatternFlowIpv6FlowLabel value in the FlowIpv6 object
func (obj *flowIpv6) SetFlowLabel(value PatternFlowIpv6FlowLabel) FlowIpv6 {

	obj.flowLabelHolder = nil
	obj.obj.FlowLabel = value.msg()

	return obj
}

// description is TBD
// PayloadLength returns a PatternFlowIpv6PayloadLength
func (obj *flowIpv6) PayloadLength() PatternFlowIpv6PayloadLength {
	if obj.obj.PayloadLength == nil {
		obj.obj.PayloadLength = NewPatternFlowIpv6PayloadLength().msg()
	}
	if obj.payloadLengthHolder == nil {
		obj.payloadLengthHolder = &patternFlowIpv6PayloadLength{obj: obj.obj.PayloadLength}
	}
	return obj.payloadLengthHolder
}

// description is TBD
// PayloadLength returns a PatternFlowIpv6PayloadLength
func (obj *flowIpv6) HasPayloadLength() bool {
	return obj.obj.PayloadLength != nil
}

// description is TBD
// SetPayloadLength sets the PatternFlowIpv6PayloadLength value in the FlowIpv6 object
func (obj *flowIpv6) SetPayloadLength(value PatternFlowIpv6PayloadLength) FlowIpv6 {

	obj.payloadLengthHolder = nil
	obj.obj.PayloadLength = value.msg()

	return obj
}

// description is TBD
// NextHeader returns a PatternFlowIpv6NextHeader
func (obj *flowIpv6) NextHeader() PatternFlowIpv6NextHeader {
	if obj.obj.NextHeader == nil {
		obj.obj.NextHeader = NewPatternFlowIpv6NextHeader().msg()
	}
	if obj.nextHeaderHolder == nil {
		obj.nextHeaderHolder = &patternFlowIpv6NextHeader{obj: obj.obj.NextHeader}
	}
	return obj.nextHeaderHolder
}

// description is TBD
// NextHeader returns a PatternFlowIpv6NextHeader
func (obj *flowIpv6) HasNextHeader() bool {
	return obj.obj.NextHeader != nil
}

// description is TBD
// SetNextHeader sets the PatternFlowIpv6NextHeader value in the FlowIpv6 object
func (obj *flowIpv6) SetNextHeader(value PatternFlowIpv6NextHeader) FlowIpv6 {

	obj.nextHeaderHolder = nil
	obj.obj.NextHeader = value.msg()

	return obj
}

// description is TBD
// HopLimit returns a PatternFlowIpv6HopLimit
func (obj *flowIpv6) HopLimit() PatternFlowIpv6HopLimit {
	if obj.obj.HopLimit == nil {
		obj.obj.HopLimit = NewPatternFlowIpv6HopLimit().msg()
	}
	if obj.hopLimitHolder == nil {
		obj.hopLimitHolder = &patternFlowIpv6HopLimit{obj: obj.obj.HopLimit}
	}
	return obj.hopLimitHolder
}

// description is TBD
// HopLimit returns a PatternFlowIpv6HopLimit
func (obj *flowIpv6) HasHopLimit() bool {
	return obj.obj.HopLimit != nil
}

// description is TBD
// SetHopLimit sets the PatternFlowIpv6HopLimit value in the FlowIpv6 object
func (obj *flowIpv6) SetHopLimit(value PatternFlowIpv6HopLimit) FlowIpv6 {

	obj.hopLimitHolder = nil
	obj.obj.HopLimit = value.msg()

	return obj
}

// description is TBD
// Src returns a PatternFlowIpv6Src
func (obj *flowIpv6) Src() PatternFlowIpv6Src {
	if obj.obj.Src == nil {
		obj.obj.Src = NewPatternFlowIpv6Src().msg()
	}
	if obj.srcHolder == nil {
		obj.srcHolder = &patternFlowIpv6Src{obj: obj.obj.Src}
	}
	return obj.srcHolder
}

// description is TBD
// Src returns a PatternFlowIpv6Src
func (obj *flowIpv6) HasSrc() bool {
	return obj.obj.Src != nil
}

// description is TBD
// SetSrc sets the PatternFlowIpv6Src value in the FlowIpv6 object
func (obj *flowIpv6) SetSrc(value PatternFlowIpv6Src) FlowIpv6 {

	obj.srcHolder = nil
	obj.obj.Src = value.msg()

	return obj
}

// description is TBD
// Dst returns a PatternFlowIpv6Dst
func (obj *flowIpv6) Dst() PatternFlowIpv6Dst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = NewPatternFlowIpv6Dst().msg()
	}
	if obj.dstHolder == nil {
		obj.dstHolder = &patternFlowIpv6Dst{obj: obj.obj.Dst}
	}
	return obj.dstHolder
}

// description is TBD
// Dst returns a PatternFlowIpv6Dst
func (obj *flowIpv6) HasDst() bool {
	return obj.obj.Dst != nil
}

// description is TBD
// SetDst sets the PatternFlowIpv6Dst value in the FlowIpv6 object
func (obj *flowIpv6) SetDst(value PatternFlowIpv6Dst) FlowIpv6 {

	obj.dstHolder = nil
	obj.obj.Dst = value.msg()

	return obj
}

func (obj *flowIpv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.TrafficClass != nil {

		obj.TrafficClass().validateObj(vObj, set_default)
	}

	if obj.obj.FlowLabel != nil {

		obj.FlowLabel().validateObj(vObj, set_default)
	}

	if obj.obj.PayloadLength != nil {

		obj.PayloadLength().validateObj(vObj, set_default)
	}

	if obj.obj.NextHeader != nil {

		obj.NextHeader().validateObj(vObj, set_default)
	}

	if obj.obj.HopLimit != nil {

		obj.HopLimit().validateObj(vObj, set_default)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(vObj, set_default)
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6) setDefault() {

}
