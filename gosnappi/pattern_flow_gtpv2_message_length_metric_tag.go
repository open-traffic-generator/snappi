package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2MessageLengthMetricTag *****
type patternFlowGtpv2MessageLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv2MessageLengthMetricTag
	marshaller   marshalPatternFlowGtpv2MessageLengthMetricTag
	unMarshaller unMarshalPatternFlowGtpv2MessageLengthMetricTag
}

func NewPatternFlowGtpv2MessageLengthMetricTag() PatternFlowGtpv2MessageLengthMetricTag {
	obj := patternFlowGtpv2MessageLengthMetricTag{obj: &otg.PatternFlowGtpv2MessageLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2MessageLengthMetricTag) msg() *otg.PatternFlowGtpv2MessageLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageLengthMetricTag) setMsg(msg *otg.PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2MessageLengthMetricTag struct {
	obj *patternFlowGtpv2MessageLengthMetricTag
}

type marshalPatternFlowGtpv2MessageLengthMetricTag interface {
	// ToProto marshals PatternFlowGtpv2MessageLengthMetricTag to protobuf object *otg.PatternFlowGtpv2MessageLengthMetricTag
	ToProto() (*otg.PatternFlowGtpv2MessageLengthMetricTag, error)
	// ToPbText marshals PatternFlowGtpv2MessageLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2MessageLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2MessageLengthMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2MessageLengthMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2MessageLengthMetricTag struct {
	obj *patternFlowGtpv2MessageLengthMetricTag
}

type unMarshalPatternFlowGtpv2MessageLengthMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv2MessageLengthMetricTag from protobuf object *otg.PatternFlowGtpv2MessageLengthMetricTag
	FromProto(msg *otg.PatternFlowGtpv2MessageLengthMetricTag) (PatternFlowGtpv2MessageLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv2MessageLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2MessageLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2MessageLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2MessageLengthMetricTag) Marshal() marshalPatternFlowGtpv2MessageLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2MessageLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2MessageLengthMetricTag) Unmarshal() unMarshalPatternFlowGtpv2MessageLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2MessageLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2MessageLengthMetricTag) ToProto() (*otg.PatternFlowGtpv2MessageLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2MessageLengthMetricTag) FromProto(msg *otg.PatternFlowGtpv2MessageLengthMetricTag) (PatternFlowGtpv2MessageLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2MessageLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageLengthMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2MessageLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2MessageLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2MessageLengthMetricTag) Clone() (PatternFlowGtpv2MessageLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2MessageLengthMetricTag()
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

// PatternFlowGtpv2MessageLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv2MessageLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv2MessageLengthMetricTag to protobuf object *otg.PatternFlowGtpv2MessageLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2MessageLengthMetricTag
	// setMsg unmarshals PatternFlowGtpv2MessageLengthMetricTag from protobuf object *otg.PatternFlowGtpv2MessageLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2MessageLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2MessageLengthMetricTag
	// validate validates PatternFlowGtpv2MessageLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2MessageLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv2MessageLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv2MessageLengthMetricTag
	SetName(value string) PatternFlowGtpv2MessageLengthMetricTag
	// Offset returns uint32, set in PatternFlowGtpv2MessageLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv2MessageLengthMetricTag
	SetOffset(value uint32) PatternFlowGtpv2MessageLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv2MessageLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv2MessageLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv2MessageLengthMetricTag
	SetLength(value uint32) PatternFlowGtpv2MessageLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv2MessageLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv2MessageLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv2MessageLengthMetricTag object
func (obj *patternFlowGtpv2MessageLengthMetricTag) SetName(value string) PatternFlowGtpv2MessageLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2MessageLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2MessageLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv2MessageLengthMetricTag object
func (obj *patternFlowGtpv2MessageLengthMetricTag) SetOffset(value uint32) PatternFlowGtpv2MessageLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2MessageLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2MessageLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv2MessageLengthMetricTag object
func (obj *patternFlowGtpv2MessageLengthMetricTag) SetLength(value uint32) PatternFlowGtpv2MessageLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv2MessageLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv2MessageLengthMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageLengthMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv2MessageLengthMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv2MessageLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
