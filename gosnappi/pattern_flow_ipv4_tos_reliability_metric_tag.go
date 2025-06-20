package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosReliabilityMetricTag *****
type patternFlowIpv4TosReliabilityMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4TosReliabilityMetricTag
	marshaller   marshalPatternFlowIpv4TosReliabilityMetricTag
	unMarshaller unMarshalPatternFlowIpv4TosReliabilityMetricTag
}

func NewPatternFlowIpv4TosReliabilityMetricTag() PatternFlowIpv4TosReliabilityMetricTag {
	obj := patternFlowIpv4TosReliabilityMetricTag{obj: &otg.PatternFlowIpv4TosReliabilityMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosReliabilityMetricTag) msg() *otg.PatternFlowIpv4TosReliabilityMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4TosReliabilityMetricTag) setMsg(msg *otg.PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosReliabilityMetricTag struct {
	obj *patternFlowIpv4TosReliabilityMetricTag
}

type marshalPatternFlowIpv4TosReliabilityMetricTag interface {
	// ToProto marshals PatternFlowIpv4TosReliabilityMetricTag to protobuf object *otg.PatternFlowIpv4TosReliabilityMetricTag
	ToProto() (*otg.PatternFlowIpv4TosReliabilityMetricTag, error)
	// ToPbText marshals PatternFlowIpv4TosReliabilityMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosReliabilityMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosReliabilityMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosReliabilityMetricTag struct {
	obj *patternFlowIpv4TosReliabilityMetricTag
}

type unMarshalPatternFlowIpv4TosReliabilityMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4TosReliabilityMetricTag from protobuf object *otg.PatternFlowIpv4TosReliabilityMetricTag
	FromProto(msg *otg.PatternFlowIpv4TosReliabilityMetricTag) (PatternFlowIpv4TosReliabilityMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4TosReliabilityMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosReliabilityMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosReliabilityMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosReliabilityMetricTag) Marshal() marshalPatternFlowIpv4TosReliabilityMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosReliabilityMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosReliabilityMetricTag) Unmarshal() unMarshalPatternFlowIpv4TosReliabilityMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosReliabilityMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosReliabilityMetricTag) ToProto() (*otg.PatternFlowIpv4TosReliabilityMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosReliabilityMetricTag) FromProto(msg *otg.PatternFlowIpv4TosReliabilityMetricTag) (PatternFlowIpv4TosReliabilityMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosReliabilityMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliabilityMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosReliabilityMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliabilityMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosReliabilityMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliabilityMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosReliabilityMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosReliabilityMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosReliabilityMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosReliabilityMetricTag) Clone() (PatternFlowIpv4TosReliabilityMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosReliabilityMetricTag()
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

// PatternFlowIpv4TosReliabilityMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4TosReliabilityMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4TosReliabilityMetricTag to protobuf object *otg.PatternFlowIpv4TosReliabilityMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosReliabilityMetricTag
	// setMsg unmarshals PatternFlowIpv4TosReliabilityMetricTag from protobuf object *otg.PatternFlowIpv4TosReliabilityMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosReliabilityMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosReliabilityMetricTag
	// validate validates PatternFlowIpv4TosReliabilityMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosReliabilityMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4TosReliabilityMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4TosReliabilityMetricTag
	SetName(value string) PatternFlowIpv4TosReliabilityMetricTag
	// Offset returns uint32, set in PatternFlowIpv4TosReliabilityMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4TosReliabilityMetricTag
	SetOffset(value uint32) PatternFlowIpv4TosReliabilityMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4TosReliabilityMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4TosReliabilityMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4TosReliabilityMetricTag
	SetLength(value uint32) PatternFlowIpv4TosReliabilityMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4TosReliabilityMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4TosReliabilityMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4TosReliabilityMetricTag object
func (obj *patternFlowIpv4TosReliabilityMetricTag) SetName(value string) PatternFlowIpv4TosReliabilityMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosReliabilityMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosReliabilityMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4TosReliabilityMetricTag object
func (obj *patternFlowIpv4TosReliabilityMetricTag) SetOffset(value uint32) PatternFlowIpv4TosReliabilityMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosReliabilityMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosReliabilityMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4TosReliabilityMetricTag object
func (obj *patternFlowIpv4TosReliabilityMetricTag) SetLength(value uint32) PatternFlowIpv4TosReliabilityMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4TosReliabilityMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4TosReliabilityMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosReliabilityMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4TosReliabilityMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4TosReliabilityMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
