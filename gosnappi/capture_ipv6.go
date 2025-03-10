package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureIpv6 *****
type captureIpv6 struct {
	validation
	obj                 *otg.CaptureIpv6
	marshaller          marshalCaptureIpv6
	unMarshaller        unMarshalCaptureIpv6
	versionHolder       CaptureField
	trafficClassHolder  CaptureField
	flowLabelHolder     CaptureField
	payloadLengthHolder CaptureField
	nextHeaderHolder    CaptureField
	hopLimitHolder      CaptureField
	srcHolder           CaptureField
	dstHolder           CaptureField
}

func NewCaptureIpv6() CaptureIpv6 {
	obj := captureIpv6{obj: &otg.CaptureIpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *captureIpv6) msg() *otg.CaptureIpv6 {
	return obj.obj
}

func (obj *captureIpv6) setMsg(msg *otg.CaptureIpv6) CaptureIpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureIpv6 struct {
	obj *captureIpv6
}

type marshalCaptureIpv6 interface {
	// ToProto marshals CaptureIpv6 to protobuf object *otg.CaptureIpv6
	ToProto() (*otg.CaptureIpv6, error)
	// ToPbText marshals CaptureIpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureIpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureIpv6 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals CaptureIpv6 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalcaptureIpv6 struct {
	obj *captureIpv6
}

type unMarshalCaptureIpv6 interface {
	// FromProto unmarshals CaptureIpv6 from protobuf object *otg.CaptureIpv6
	FromProto(msg *otg.CaptureIpv6) (CaptureIpv6, error)
	// FromPbText unmarshals CaptureIpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureIpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureIpv6 from JSON text
	FromJson(value string) error
}

func (obj *captureIpv6) Marshal() marshalCaptureIpv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureIpv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureIpv6) Unmarshal() unMarshalCaptureIpv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureIpv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureIpv6) ToProto() (*otg.CaptureIpv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureIpv6) FromProto(msg *otg.CaptureIpv6) (CaptureIpv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureIpv6) ToPbText() (string, error) {
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

func (m *unMarshalcaptureIpv6) FromPbText(value string) error {
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

func (m *marshalcaptureIpv6) ToYaml() (string, error) {
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

func (m *unMarshalcaptureIpv6) FromYaml(value string) error {
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

func (m *marshalcaptureIpv6) ToJsonRaw() (string, error) {
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

func (m *marshalcaptureIpv6) ToJson() (string, error) {
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

func (m *unMarshalcaptureIpv6) FromJson(value string) error {
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

func (obj *captureIpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureIpv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureIpv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureIpv6) Clone() (CaptureIpv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureIpv6()
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

func (obj *captureIpv6) setNil() {
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

// CaptureIpv6 is description is TBD
type CaptureIpv6 interface {
	Validation
	// msg marshals CaptureIpv6 to protobuf object *otg.CaptureIpv6
	// and doesn't set defaults
	msg() *otg.CaptureIpv6
	// setMsg unmarshals CaptureIpv6 from protobuf object *otg.CaptureIpv6
	// and doesn't set defaults
	setMsg(*otg.CaptureIpv6) CaptureIpv6
	// provides marshal interface
	Marshal() marshalCaptureIpv6
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureIpv6
	// validate validates CaptureIpv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureIpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns CaptureField, set in CaptureIpv6.
	// CaptureField is description is TBD
	Version() CaptureField
	// SetVersion assigns CaptureField provided by user to CaptureIpv6.
	// CaptureField is description is TBD
	SetVersion(value CaptureField) CaptureIpv6
	// HasVersion checks if Version has been set in CaptureIpv6
	HasVersion() bool
	// TrafficClass returns CaptureField, set in CaptureIpv6.
	// CaptureField is description is TBD
	TrafficClass() CaptureField
	// SetTrafficClass assigns CaptureField provided by user to CaptureIpv6.
	// CaptureField is description is TBD
	SetTrafficClass(value CaptureField) CaptureIpv6
	// HasTrafficClass checks if TrafficClass has been set in CaptureIpv6
	HasTrafficClass() bool
	// FlowLabel returns CaptureField, set in CaptureIpv6.
	// CaptureField is description is TBD
	FlowLabel() CaptureField
	// SetFlowLabel assigns CaptureField provided by user to CaptureIpv6.
	// CaptureField is description is TBD
	SetFlowLabel(value CaptureField) CaptureIpv6
	// HasFlowLabel checks if FlowLabel has been set in CaptureIpv6
	HasFlowLabel() bool
	// PayloadLength returns CaptureField, set in CaptureIpv6.
	// CaptureField is description is TBD
	PayloadLength() CaptureField
	// SetPayloadLength assigns CaptureField provided by user to CaptureIpv6.
	// CaptureField is description is TBD
	SetPayloadLength(value CaptureField) CaptureIpv6
	// HasPayloadLength checks if PayloadLength has been set in CaptureIpv6
	HasPayloadLength() bool
	// NextHeader returns CaptureField, set in CaptureIpv6.
	// CaptureField is description is TBD
	NextHeader() CaptureField
	// SetNextHeader assigns CaptureField provided by user to CaptureIpv6.
	// CaptureField is description is TBD
	SetNextHeader(value CaptureField) CaptureIpv6
	// HasNextHeader checks if NextHeader has been set in CaptureIpv6
	HasNextHeader() bool
	// HopLimit returns CaptureField, set in CaptureIpv6.
	// CaptureField is description is TBD
	HopLimit() CaptureField
	// SetHopLimit assigns CaptureField provided by user to CaptureIpv6.
	// CaptureField is description is TBD
	SetHopLimit(value CaptureField) CaptureIpv6
	// HasHopLimit checks if HopLimit has been set in CaptureIpv6
	HasHopLimit() bool
	// Src returns CaptureField, set in CaptureIpv6.
	// CaptureField is description is TBD
	Src() CaptureField
	// SetSrc assigns CaptureField provided by user to CaptureIpv6.
	// CaptureField is description is TBD
	SetSrc(value CaptureField) CaptureIpv6
	// HasSrc checks if Src has been set in CaptureIpv6
	HasSrc() bool
	// Dst returns CaptureField, set in CaptureIpv6.
	// CaptureField is description is TBD
	Dst() CaptureField
	// SetDst assigns CaptureField provided by user to CaptureIpv6.
	// CaptureField is description is TBD
	SetDst(value CaptureField) CaptureIpv6
	// HasDst checks if Dst has been set in CaptureIpv6
	HasDst() bool
	setNil()
}

// description is TBD
// Version returns a CaptureField
func (obj *captureIpv6) Version() CaptureField {
	if obj.obj.Version == nil {
		obj.obj.Version = NewCaptureField().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &captureField{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a CaptureField
func (obj *captureIpv6) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the CaptureField value in the CaptureIpv6 object
func (obj *captureIpv6) SetVersion(value CaptureField) CaptureIpv6 {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// TrafficClass returns a CaptureField
func (obj *captureIpv6) TrafficClass() CaptureField {
	if obj.obj.TrafficClass == nil {
		obj.obj.TrafficClass = NewCaptureField().msg()
	}
	if obj.trafficClassHolder == nil {
		obj.trafficClassHolder = &captureField{obj: obj.obj.TrafficClass}
	}
	return obj.trafficClassHolder
}

// description is TBD
// TrafficClass returns a CaptureField
func (obj *captureIpv6) HasTrafficClass() bool {
	return obj.obj.TrafficClass != nil
}

// description is TBD
// SetTrafficClass sets the CaptureField value in the CaptureIpv6 object
func (obj *captureIpv6) SetTrafficClass(value CaptureField) CaptureIpv6 {

	obj.trafficClassHolder = nil
	obj.obj.TrafficClass = value.msg()

	return obj
}

// description is TBD
// FlowLabel returns a CaptureField
func (obj *captureIpv6) FlowLabel() CaptureField {
	if obj.obj.FlowLabel == nil {
		obj.obj.FlowLabel = NewCaptureField().msg()
	}
	if obj.flowLabelHolder == nil {
		obj.flowLabelHolder = &captureField{obj: obj.obj.FlowLabel}
	}
	return obj.flowLabelHolder
}

// description is TBD
// FlowLabel returns a CaptureField
func (obj *captureIpv6) HasFlowLabel() bool {
	return obj.obj.FlowLabel != nil
}

// description is TBD
// SetFlowLabel sets the CaptureField value in the CaptureIpv6 object
func (obj *captureIpv6) SetFlowLabel(value CaptureField) CaptureIpv6 {

	obj.flowLabelHolder = nil
	obj.obj.FlowLabel = value.msg()

	return obj
}

// description is TBD
// PayloadLength returns a CaptureField
func (obj *captureIpv6) PayloadLength() CaptureField {
	if obj.obj.PayloadLength == nil {
		obj.obj.PayloadLength = NewCaptureField().msg()
	}
	if obj.payloadLengthHolder == nil {
		obj.payloadLengthHolder = &captureField{obj: obj.obj.PayloadLength}
	}
	return obj.payloadLengthHolder
}

// description is TBD
// PayloadLength returns a CaptureField
func (obj *captureIpv6) HasPayloadLength() bool {
	return obj.obj.PayloadLength != nil
}

// description is TBD
// SetPayloadLength sets the CaptureField value in the CaptureIpv6 object
func (obj *captureIpv6) SetPayloadLength(value CaptureField) CaptureIpv6 {

	obj.payloadLengthHolder = nil
	obj.obj.PayloadLength = value.msg()

	return obj
}

// description is TBD
// NextHeader returns a CaptureField
func (obj *captureIpv6) NextHeader() CaptureField {
	if obj.obj.NextHeader == nil {
		obj.obj.NextHeader = NewCaptureField().msg()
	}
	if obj.nextHeaderHolder == nil {
		obj.nextHeaderHolder = &captureField{obj: obj.obj.NextHeader}
	}
	return obj.nextHeaderHolder
}

// description is TBD
// NextHeader returns a CaptureField
func (obj *captureIpv6) HasNextHeader() bool {
	return obj.obj.NextHeader != nil
}

// description is TBD
// SetNextHeader sets the CaptureField value in the CaptureIpv6 object
func (obj *captureIpv6) SetNextHeader(value CaptureField) CaptureIpv6 {

	obj.nextHeaderHolder = nil
	obj.obj.NextHeader = value.msg()

	return obj
}

// description is TBD
// HopLimit returns a CaptureField
func (obj *captureIpv6) HopLimit() CaptureField {
	if obj.obj.HopLimit == nil {
		obj.obj.HopLimit = NewCaptureField().msg()
	}
	if obj.hopLimitHolder == nil {
		obj.hopLimitHolder = &captureField{obj: obj.obj.HopLimit}
	}
	return obj.hopLimitHolder
}

// description is TBD
// HopLimit returns a CaptureField
func (obj *captureIpv6) HasHopLimit() bool {
	return obj.obj.HopLimit != nil
}

// description is TBD
// SetHopLimit sets the CaptureField value in the CaptureIpv6 object
func (obj *captureIpv6) SetHopLimit(value CaptureField) CaptureIpv6 {

	obj.hopLimitHolder = nil
	obj.obj.HopLimit = value.msg()

	return obj
}

// description is TBD
// Src returns a CaptureField
func (obj *captureIpv6) Src() CaptureField {
	if obj.obj.Src == nil {
		obj.obj.Src = NewCaptureField().msg()
	}
	if obj.srcHolder == nil {
		obj.srcHolder = &captureField{obj: obj.obj.Src}
	}
	return obj.srcHolder
}

// description is TBD
// Src returns a CaptureField
func (obj *captureIpv6) HasSrc() bool {
	return obj.obj.Src != nil
}

// description is TBD
// SetSrc sets the CaptureField value in the CaptureIpv6 object
func (obj *captureIpv6) SetSrc(value CaptureField) CaptureIpv6 {

	obj.srcHolder = nil
	obj.obj.Src = value.msg()

	return obj
}

// description is TBD
// Dst returns a CaptureField
func (obj *captureIpv6) Dst() CaptureField {
	if obj.obj.Dst == nil {
		obj.obj.Dst = NewCaptureField().msg()
	}
	if obj.dstHolder == nil {
		obj.dstHolder = &captureField{obj: obj.obj.Dst}
	}
	return obj.dstHolder
}

// description is TBD
// Dst returns a CaptureField
func (obj *captureIpv6) HasDst() bool {
	return obj.obj.Dst != nil
}

// description is TBD
// SetDst sets the CaptureField value in the CaptureIpv6 object
func (obj *captureIpv6) SetDst(value CaptureField) CaptureIpv6 {

	obj.dstHolder = nil
	obj.obj.Dst = value.msg()

	return obj
}

func (obj *captureIpv6) validateObj(vObj *validation, set_default bool) {
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

func (obj *captureIpv6) setDefault() {

}
