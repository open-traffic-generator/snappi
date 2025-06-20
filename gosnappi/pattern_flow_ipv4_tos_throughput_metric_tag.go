package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosThroughputMetricTag *****
type patternFlowIpv4TosThroughputMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4TosThroughputMetricTag
	marshaller   marshalPatternFlowIpv4TosThroughputMetricTag
	unMarshaller unMarshalPatternFlowIpv4TosThroughputMetricTag
}

func NewPatternFlowIpv4TosThroughputMetricTag() PatternFlowIpv4TosThroughputMetricTag {
	obj := patternFlowIpv4TosThroughputMetricTag{obj: &otg.PatternFlowIpv4TosThroughputMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosThroughputMetricTag) msg() *otg.PatternFlowIpv4TosThroughputMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4TosThroughputMetricTag) setMsg(msg *otg.PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosThroughputMetricTag struct {
	obj *patternFlowIpv4TosThroughputMetricTag
}

type marshalPatternFlowIpv4TosThroughputMetricTag interface {
	// ToProto marshals PatternFlowIpv4TosThroughputMetricTag to protobuf object *otg.PatternFlowIpv4TosThroughputMetricTag
	ToProto() (*otg.PatternFlowIpv4TosThroughputMetricTag, error)
	// ToPbText marshals PatternFlowIpv4TosThroughputMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosThroughputMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosThroughputMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosThroughputMetricTag struct {
	obj *patternFlowIpv4TosThroughputMetricTag
}

type unMarshalPatternFlowIpv4TosThroughputMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4TosThroughputMetricTag from protobuf object *otg.PatternFlowIpv4TosThroughputMetricTag
	FromProto(msg *otg.PatternFlowIpv4TosThroughputMetricTag) (PatternFlowIpv4TosThroughputMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4TosThroughputMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosThroughputMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosThroughputMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosThroughputMetricTag) Marshal() marshalPatternFlowIpv4TosThroughputMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosThroughputMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosThroughputMetricTag) Unmarshal() unMarshalPatternFlowIpv4TosThroughputMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosThroughputMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosThroughputMetricTag) ToProto() (*otg.PatternFlowIpv4TosThroughputMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosThroughputMetricTag) FromProto(msg *otg.PatternFlowIpv4TosThroughputMetricTag) (PatternFlowIpv4TosThroughputMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosThroughputMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughputMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosThroughputMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughputMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosThroughputMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughputMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosThroughputMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosThroughputMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosThroughputMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosThroughputMetricTag) Clone() (PatternFlowIpv4TosThroughputMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosThroughputMetricTag()
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

// PatternFlowIpv4TosThroughputMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4TosThroughputMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4TosThroughputMetricTag to protobuf object *otg.PatternFlowIpv4TosThroughputMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosThroughputMetricTag
	// setMsg unmarshals PatternFlowIpv4TosThroughputMetricTag from protobuf object *otg.PatternFlowIpv4TosThroughputMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosThroughputMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosThroughputMetricTag
	// validate validates PatternFlowIpv4TosThroughputMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosThroughputMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4TosThroughputMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4TosThroughputMetricTag
	SetName(value string) PatternFlowIpv4TosThroughputMetricTag
	// Offset returns uint32, set in PatternFlowIpv4TosThroughputMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4TosThroughputMetricTag
	SetOffset(value uint32) PatternFlowIpv4TosThroughputMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4TosThroughputMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4TosThroughputMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4TosThroughputMetricTag
	SetLength(value uint32) PatternFlowIpv4TosThroughputMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4TosThroughputMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4TosThroughputMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4TosThroughputMetricTag object
func (obj *patternFlowIpv4TosThroughputMetricTag) SetName(value string) PatternFlowIpv4TosThroughputMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosThroughputMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosThroughputMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4TosThroughputMetricTag object
func (obj *patternFlowIpv4TosThroughputMetricTag) SetOffset(value uint32) PatternFlowIpv4TosThroughputMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosThroughputMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosThroughputMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4TosThroughputMetricTag object
func (obj *patternFlowIpv4TosThroughputMetricTag) SetLength(value uint32) PatternFlowIpv4TosThroughputMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4TosThroughputMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4TosThroughputMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosThroughputMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4TosThroughputMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4TosThroughputMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
