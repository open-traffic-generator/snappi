package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpduCollector *****
type flowLacpduCollector struct {
	validation
	obj            *otg.FlowLacpduCollector
	marshaller     marshalFlowLacpduCollector
	unMarshaller   unMarshalFlowLacpduCollector
	typeHolder     PatternFlowLacpduCollectorType
	lengthHolder   PatternFlowLacpduCollectorLength
	maxDelayHolder PatternFlowLacpduCollectorMaxDelay
}

func NewFlowLacpduCollector() FlowLacpduCollector {
	obj := flowLacpduCollector{obj: &otg.FlowLacpduCollector{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpduCollector) msg() *otg.FlowLacpduCollector {
	return obj.obj
}

func (obj *flowLacpduCollector) setMsg(msg *otg.FlowLacpduCollector) FlowLacpduCollector {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpduCollector struct {
	obj *flowLacpduCollector
}

type marshalFlowLacpduCollector interface {
	// ToProto marshals FlowLacpduCollector to protobuf object *otg.FlowLacpduCollector
	ToProto() (*otg.FlowLacpduCollector, error)
	// ToPbText marshals FlowLacpduCollector to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpduCollector to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpduCollector to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpduCollector struct {
	obj *flowLacpduCollector
}

type unMarshalFlowLacpduCollector interface {
	// FromProto unmarshals FlowLacpduCollector from protobuf object *otg.FlowLacpduCollector
	FromProto(msg *otg.FlowLacpduCollector) (FlowLacpduCollector, error)
	// FromPbText unmarshals FlowLacpduCollector from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpduCollector from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpduCollector from JSON text
	FromJson(value string) error
}

func (obj *flowLacpduCollector) Marshal() marshalFlowLacpduCollector {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpduCollector{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpduCollector) Unmarshal() unMarshalFlowLacpduCollector {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpduCollector{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpduCollector) ToProto() (*otg.FlowLacpduCollector, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpduCollector) FromProto(msg *otg.FlowLacpduCollector) (FlowLacpduCollector, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpduCollector) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpduCollector) FromPbText(value string) error {
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

func (m *marshalflowLacpduCollector) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpduCollector) FromYaml(value string) error {
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

func (m *marshalflowLacpduCollector) ToJson() (string, error) {
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

func (m *unMarshalflowLacpduCollector) FromJson(value string) error {
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

func (obj *flowLacpduCollector) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpduCollector) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpduCollector) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpduCollector) Clone() (FlowLacpduCollector, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpduCollector()
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

func (obj *flowLacpduCollector) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.maxDelayHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacpduCollector is information about frame collection parameters.
type FlowLacpduCollector interface {
	Validation
	// msg marshals FlowLacpduCollector to protobuf object *otg.FlowLacpduCollector
	// and doesn't set defaults
	msg() *otg.FlowLacpduCollector
	// setMsg unmarshals FlowLacpduCollector from protobuf object *otg.FlowLacpduCollector
	// and doesn't set defaults
	setMsg(*otg.FlowLacpduCollector) FlowLacpduCollector
	// provides marshal interface
	Marshal() marshalFlowLacpduCollector
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpduCollector
	// validate validates FlowLacpduCollector
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpduCollector, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowLacpduCollectorType, set in FlowLacpduCollector.
	// PatternFlowLacpduCollectorType is tLV Type for Collector Information. The value 0x03 identifies this TLV.
	Type() PatternFlowLacpduCollectorType
	// SetType assigns PatternFlowLacpduCollectorType provided by user to FlowLacpduCollector.
	// PatternFlowLacpduCollectorType is tLV Type for Collector Information. The value 0x03 identifies this TLV.
	SetType(value PatternFlowLacpduCollectorType) FlowLacpduCollector
	// HasType checks if Type has been set in FlowLacpduCollector
	HasType() bool
	// Length returns PatternFlowLacpduCollectorLength, set in FlowLacpduCollector.
	// PatternFlowLacpduCollectorLength is the length of the Collector Information TLV payload in bytes. The value must be 16 (0x10).
	Length() PatternFlowLacpduCollectorLength
	// SetLength assigns PatternFlowLacpduCollectorLength provided by user to FlowLacpduCollector.
	// PatternFlowLacpduCollectorLength is the length of the Collector Information TLV payload in bytes. The value must be 16 (0x10).
	SetLength(value PatternFlowLacpduCollectorLength) FlowLacpduCollector
	// HasLength checks if Length has been set in FlowLacpduCollector
	HasLength() bool
	// MaxDelay returns PatternFlowLacpduCollectorMaxDelay, set in FlowLacpduCollector.
	// PatternFlowLacpduCollectorMaxDelay is indicates the maximum delay, in units of 10 microseconds, that the transmitting system's aggregator will take to buffer frames from its collector before emitting a frame.
	MaxDelay() PatternFlowLacpduCollectorMaxDelay
	// SetMaxDelay assigns PatternFlowLacpduCollectorMaxDelay provided by user to FlowLacpduCollector.
	// PatternFlowLacpduCollectorMaxDelay is indicates the maximum delay, in units of 10 microseconds, that the transmitting system's aggregator will take to buffer frames from its collector before emitting a frame.
	SetMaxDelay(value PatternFlowLacpduCollectorMaxDelay) FlowLacpduCollector
	// HasMaxDelay checks if MaxDelay has been set in FlowLacpduCollector
	HasMaxDelay() bool
	// Reserved returns string, set in FlowLacpduCollector.
	Reserved() string
	// SetReserved assigns string provided by user to FlowLacpduCollector
	SetReserved(value string) FlowLacpduCollector
	// HasReserved checks if Reserved has been set in FlowLacpduCollector
	HasReserved() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowLacpduCollectorType
func (obj *flowLacpduCollector) Type() PatternFlowLacpduCollectorType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowLacpduCollectorType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowLacpduCollectorType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowLacpduCollectorType
func (obj *flowLacpduCollector) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowLacpduCollectorType value in the FlowLacpduCollector object
func (obj *flowLacpduCollector) SetType(value PatternFlowLacpduCollectorType) FlowLacpduCollector {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowLacpduCollectorLength
func (obj *flowLacpduCollector) Length() PatternFlowLacpduCollectorLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowLacpduCollectorLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowLacpduCollectorLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowLacpduCollectorLength
func (obj *flowLacpduCollector) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowLacpduCollectorLength value in the FlowLacpduCollector object
func (obj *flowLacpduCollector) SetLength(value PatternFlowLacpduCollectorLength) FlowLacpduCollector {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// MaxDelay returns a PatternFlowLacpduCollectorMaxDelay
func (obj *flowLacpduCollector) MaxDelay() PatternFlowLacpduCollectorMaxDelay {
	if obj.obj.MaxDelay == nil {
		obj.obj.MaxDelay = NewPatternFlowLacpduCollectorMaxDelay().msg()
	}
	if obj.maxDelayHolder == nil {
		obj.maxDelayHolder = &patternFlowLacpduCollectorMaxDelay{obj: obj.obj.MaxDelay}
	}
	return obj.maxDelayHolder
}

// description is TBD
// MaxDelay returns a PatternFlowLacpduCollectorMaxDelay
func (obj *flowLacpduCollector) HasMaxDelay() bool {
	return obj.obj.MaxDelay != nil
}

// description is TBD
// SetMaxDelay sets the PatternFlowLacpduCollectorMaxDelay value in the FlowLacpduCollector object
func (obj *flowLacpduCollector) SetMaxDelay(value PatternFlowLacpduCollectorMaxDelay) FlowLacpduCollector {

	obj.maxDelayHolder = nil
	obj.obj.MaxDelay = value.msg()

	return obj
}

// Reserved field for future use in the Collector TLV. 12 bytes long and should be set to all zeros.
// Reserved returns a string
func (obj *flowLacpduCollector) Reserved() string {

	return *obj.obj.Reserved

}

// Reserved field for future use in the Collector TLV. 12 bytes long and should be set to all zeros.
// Reserved returns a string
func (obj *flowLacpduCollector) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// Reserved field for future use in the Collector TLV. 12 bytes long and should be set to all zeros.
// SetReserved sets the string value in the FlowLacpduCollector object
func (obj *flowLacpduCollector) SetReserved(value string) FlowLacpduCollector {

	obj.obj.Reserved = &value
	return obj
}

func (obj *flowLacpduCollector) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.MaxDelay != nil {

		obj.MaxDelay().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		err := obj.validateHex(obj.Reserved())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowLacpduCollector.Reserved"))
		}

	}

}

func (obj *flowLacpduCollector) setDefault() {
	if obj.obj.Reserved == nil {
		obj.SetReserved("000000000000000000000000")
	}

}
