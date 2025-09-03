package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpCollector *****
type flowLacpCollector struct {
	validation
	obj            *otg.FlowLacpCollector
	marshaller     marshalFlowLacpCollector
	unMarshaller   unMarshalFlowLacpCollector
	typeHolder     PatternFlowLacpCollectorType
	lengthHolder   PatternFlowLacpCollectorLength
	maxDelayHolder PatternFlowLacpCollectorMaxDelay
}

func NewFlowLacpCollector() FlowLacpCollector {
	obj := flowLacpCollector{obj: &otg.FlowLacpCollector{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpCollector) msg() *otg.FlowLacpCollector {
	return obj.obj
}

func (obj *flowLacpCollector) setMsg(msg *otg.FlowLacpCollector) FlowLacpCollector {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpCollector struct {
	obj *flowLacpCollector
}

type marshalFlowLacpCollector interface {
	// ToProto marshals FlowLacpCollector to protobuf object *otg.FlowLacpCollector
	ToProto() (*otg.FlowLacpCollector, error)
	// ToPbText marshals FlowLacpCollector to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpCollector to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpCollector to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpCollector struct {
	obj *flowLacpCollector
}

type unMarshalFlowLacpCollector interface {
	// FromProto unmarshals FlowLacpCollector from protobuf object *otg.FlowLacpCollector
	FromProto(msg *otg.FlowLacpCollector) (FlowLacpCollector, error)
	// FromPbText unmarshals FlowLacpCollector from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpCollector from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpCollector from JSON text
	FromJson(value string) error
}

func (obj *flowLacpCollector) Marshal() marshalFlowLacpCollector {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpCollector{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpCollector) Unmarshal() unMarshalFlowLacpCollector {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpCollector{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpCollector) ToProto() (*otg.FlowLacpCollector, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpCollector) FromProto(msg *otg.FlowLacpCollector) (FlowLacpCollector, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpCollector) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpCollector) FromPbText(value string) error {
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

func (m *marshalflowLacpCollector) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpCollector) FromYaml(value string) error {
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

func (m *marshalflowLacpCollector) ToJson() (string, error) {
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

func (m *unMarshalflowLacpCollector) FromJson(value string) error {
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

func (obj *flowLacpCollector) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpCollector) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpCollector) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpCollector) Clone() (FlowLacpCollector, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpCollector()
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

func (obj *flowLacpCollector) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.maxDelayHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacpCollector is information about frame collection parameters.
type FlowLacpCollector interface {
	Validation
	// msg marshals FlowLacpCollector to protobuf object *otg.FlowLacpCollector
	// and doesn't set defaults
	msg() *otg.FlowLacpCollector
	// setMsg unmarshals FlowLacpCollector from protobuf object *otg.FlowLacpCollector
	// and doesn't set defaults
	setMsg(*otg.FlowLacpCollector) FlowLacpCollector
	// provides marshal interface
	Marshal() marshalFlowLacpCollector
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpCollector
	// validate validates FlowLacpCollector
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpCollector, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowLacpCollectorType, set in FlowLacpCollector.
	// PatternFlowLacpCollectorType is tLV Type for Collector Information. The value 0x03 identifies this TLV.
	Type() PatternFlowLacpCollectorType
	// SetType assigns PatternFlowLacpCollectorType provided by user to FlowLacpCollector.
	// PatternFlowLacpCollectorType is tLV Type for Collector Information. The value 0x03 identifies this TLV.
	SetType(value PatternFlowLacpCollectorType) FlowLacpCollector
	// HasType checks if Type has been set in FlowLacpCollector
	HasType() bool
	// Length returns PatternFlowLacpCollectorLength, set in FlowLacpCollector.
	// PatternFlowLacpCollectorLength is the length of the Collector Information TLV payload in bytes. The value must be 16 (0x10).
	Length() PatternFlowLacpCollectorLength
	// SetLength assigns PatternFlowLacpCollectorLength provided by user to FlowLacpCollector.
	// PatternFlowLacpCollectorLength is the length of the Collector Information TLV payload in bytes. The value must be 16 (0x10).
	SetLength(value PatternFlowLacpCollectorLength) FlowLacpCollector
	// HasLength checks if Length has been set in FlowLacpCollector
	HasLength() bool
	// MaxDelay returns PatternFlowLacpCollectorMaxDelay, set in FlowLacpCollector.
	// PatternFlowLacpCollectorMaxDelay is indicates the maximum delay, in units of 10 microseconds, that the transmitting system's aggregator will take to buffer frames from its collector before emitting a frame.
	MaxDelay() PatternFlowLacpCollectorMaxDelay
	// SetMaxDelay assigns PatternFlowLacpCollectorMaxDelay provided by user to FlowLacpCollector.
	// PatternFlowLacpCollectorMaxDelay is indicates the maximum delay, in units of 10 microseconds, that the transmitting system's aggregator will take to buffer frames from its collector before emitting a frame.
	SetMaxDelay(value PatternFlowLacpCollectorMaxDelay) FlowLacpCollector
	// HasMaxDelay checks if MaxDelay has been set in FlowLacpCollector
	HasMaxDelay() bool
	// Reserved returns string, set in FlowLacpCollector.
	Reserved() string
	// SetReserved assigns string provided by user to FlowLacpCollector
	SetReserved(value string) FlowLacpCollector
	// HasReserved checks if Reserved has been set in FlowLacpCollector
	HasReserved() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowLacpCollectorType
func (obj *flowLacpCollector) Type() PatternFlowLacpCollectorType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowLacpCollectorType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowLacpCollectorType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowLacpCollectorType
func (obj *flowLacpCollector) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowLacpCollectorType value in the FlowLacpCollector object
func (obj *flowLacpCollector) SetType(value PatternFlowLacpCollectorType) FlowLacpCollector {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowLacpCollectorLength
func (obj *flowLacpCollector) Length() PatternFlowLacpCollectorLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowLacpCollectorLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowLacpCollectorLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowLacpCollectorLength
func (obj *flowLacpCollector) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowLacpCollectorLength value in the FlowLacpCollector object
func (obj *flowLacpCollector) SetLength(value PatternFlowLacpCollectorLength) FlowLacpCollector {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// MaxDelay returns a PatternFlowLacpCollectorMaxDelay
func (obj *flowLacpCollector) MaxDelay() PatternFlowLacpCollectorMaxDelay {
	if obj.obj.MaxDelay == nil {
		obj.obj.MaxDelay = NewPatternFlowLacpCollectorMaxDelay().msg()
	}
	if obj.maxDelayHolder == nil {
		obj.maxDelayHolder = &patternFlowLacpCollectorMaxDelay{obj: obj.obj.MaxDelay}
	}
	return obj.maxDelayHolder
}

// description is TBD
// MaxDelay returns a PatternFlowLacpCollectorMaxDelay
func (obj *flowLacpCollector) HasMaxDelay() bool {
	return obj.obj.MaxDelay != nil
}

// description is TBD
// SetMaxDelay sets the PatternFlowLacpCollectorMaxDelay value in the FlowLacpCollector object
func (obj *flowLacpCollector) SetMaxDelay(value PatternFlowLacpCollectorMaxDelay) FlowLacpCollector {

	obj.maxDelayHolder = nil
	obj.obj.MaxDelay = value.msg()

	return obj
}

// Reserved field for future use in the Collector TLV. 12 bytes long and should be set to all zeros.
// Reserved returns a string
func (obj *flowLacpCollector) Reserved() string {

	return *obj.obj.Reserved

}

// Reserved field for future use in the Collector TLV. 12 bytes long and should be set to all zeros.
// Reserved returns a string
func (obj *flowLacpCollector) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// Reserved field for future use in the Collector TLV. 12 bytes long and should be set to all zeros.
// SetReserved sets the string value in the FlowLacpCollector object
func (obj *flowLacpCollector) SetReserved(value string) FlowLacpCollector {

	obj.obj.Reserved = &value
	return obj
}

func (obj *flowLacpCollector) validateObj(vObj *validation, set_default bool) {
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
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowLacpCollector.Reserved"))
		}

	}

}

func (obj *flowLacpCollector) setDefault() {
	if obj.obj.Reserved == nil {
		obj.SetReserved("000000000000000000000000")
	}

}
