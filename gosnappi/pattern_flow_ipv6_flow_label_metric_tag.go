package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6FlowLabelMetricTag *****
type patternFlowIpv6FlowLabelMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv6FlowLabelMetricTag
	marshaller   marshalPatternFlowIpv6FlowLabelMetricTag
	unMarshaller unMarshalPatternFlowIpv6FlowLabelMetricTag
}

func NewPatternFlowIpv6FlowLabelMetricTag() PatternFlowIpv6FlowLabelMetricTag {
	obj := patternFlowIpv6FlowLabelMetricTag{obj: &otg.PatternFlowIpv6FlowLabelMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6FlowLabelMetricTag) msg() *otg.PatternFlowIpv6FlowLabelMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv6FlowLabelMetricTag) setMsg(msg *otg.PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6FlowLabelMetricTag struct {
	obj *patternFlowIpv6FlowLabelMetricTag
}

type marshalPatternFlowIpv6FlowLabelMetricTag interface {
	// ToProto marshals PatternFlowIpv6FlowLabelMetricTag to protobuf object *otg.PatternFlowIpv6FlowLabelMetricTag
	ToProto() (*otg.PatternFlowIpv6FlowLabelMetricTag, error)
	// ToPbText marshals PatternFlowIpv6FlowLabelMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6FlowLabelMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6FlowLabelMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6FlowLabelMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6FlowLabelMetricTag struct {
	obj *patternFlowIpv6FlowLabelMetricTag
}

type unMarshalPatternFlowIpv6FlowLabelMetricTag interface {
	// FromProto unmarshals PatternFlowIpv6FlowLabelMetricTag from protobuf object *otg.PatternFlowIpv6FlowLabelMetricTag
	FromProto(msg *otg.PatternFlowIpv6FlowLabelMetricTag) (PatternFlowIpv6FlowLabelMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv6FlowLabelMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6FlowLabelMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6FlowLabelMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6FlowLabelMetricTag) Marshal() marshalPatternFlowIpv6FlowLabelMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6FlowLabelMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6FlowLabelMetricTag) Unmarshal() unMarshalPatternFlowIpv6FlowLabelMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6FlowLabelMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6FlowLabelMetricTag) ToProto() (*otg.PatternFlowIpv6FlowLabelMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6FlowLabelMetricTag) FromProto(msg *otg.PatternFlowIpv6FlowLabelMetricTag) (PatternFlowIpv6FlowLabelMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6FlowLabelMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6FlowLabelMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6FlowLabelMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6FlowLabelMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6FlowLabelMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6FlowLabelMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6FlowLabelMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6FlowLabelMetricTag) Clone() (PatternFlowIpv6FlowLabelMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6FlowLabelMetricTag()
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

// PatternFlowIpv6FlowLabelMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv6FlowLabelMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv6FlowLabelMetricTag to protobuf object *otg.PatternFlowIpv6FlowLabelMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6FlowLabelMetricTag
	// setMsg unmarshals PatternFlowIpv6FlowLabelMetricTag from protobuf object *otg.PatternFlowIpv6FlowLabelMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6FlowLabelMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6FlowLabelMetricTag
	// validate validates PatternFlowIpv6FlowLabelMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6FlowLabelMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv6FlowLabelMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv6FlowLabelMetricTag
	SetName(value string) PatternFlowIpv6FlowLabelMetricTag
	// Offset returns uint32, set in PatternFlowIpv6FlowLabelMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv6FlowLabelMetricTag
	SetOffset(value uint32) PatternFlowIpv6FlowLabelMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv6FlowLabelMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv6FlowLabelMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv6FlowLabelMetricTag
	SetLength(value uint32) PatternFlowIpv6FlowLabelMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv6FlowLabelMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv6FlowLabelMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv6FlowLabelMetricTag object
func (obj *patternFlowIpv6FlowLabelMetricTag) SetName(value string) PatternFlowIpv6FlowLabelMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6FlowLabelMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6FlowLabelMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv6FlowLabelMetricTag object
func (obj *patternFlowIpv6FlowLabelMetricTag) SetOffset(value uint32) PatternFlowIpv6FlowLabelMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6FlowLabelMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6FlowLabelMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv6FlowLabelMetricTag object
func (obj *patternFlowIpv6FlowLabelMetricTag) SetLength(value uint32) PatternFlowIpv6FlowLabelMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv6FlowLabelMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv6FlowLabelMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 19 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabelMetricTag.Offset <= 19 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 20 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv6FlowLabelMetricTag.Length <= 20 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv6FlowLabelMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(20)
	}

}
