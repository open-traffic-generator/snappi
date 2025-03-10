package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanPriorityMetricTag *****
type patternFlowVlanPriorityMetricTag struct {
	validation
	obj          *otg.PatternFlowVlanPriorityMetricTag
	marshaller   marshalPatternFlowVlanPriorityMetricTag
	unMarshaller unMarshalPatternFlowVlanPriorityMetricTag
}

func NewPatternFlowVlanPriorityMetricTag() PatternFlowVlanPriorityMetricTag {
	obj := patternFlowVlanPriorityMetricTag{obj: &otg.PatternFlowVlanPriorityMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanPriorityMetricTag) msg() *otg.PatternFlowVlanPriorityMetricTag {
	return obj.obj
}

func (obj *patternFlowVlanPriorityMetricTag) setMsg(msg *otg.PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanPriorityMetricTag struct {
	obj *patternFlowVlanPriorityMetricTag
}

type marshalPatternFlowVlanPriorityMetricTag interface {
	// ToProto marshals PatternFlowVlanPriorityMetricTag to protobuf object *otg.PatternFlowVlanPriorityMetricTag
	ToProto() (*otg.PatternFlowVlanPriorityMetricTag, error)
	// ToPbText marshals PatternFlowVlanPriorityMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanPriorityMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanPriorityMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVlanPriorityMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVlanPriorityMetricTag struct {
	obj *patternFlowVlanPriorityMetricTag
}

type unMarshalPatternFlowVlanPriorityMetricTag interface {
	// FromProto unmarshals PatternFlowVlanPriorityMetricTag from protobuf object *otg.PatternFlowVlanPriorityMetricTag
	FromProto(msg *otg.PatternFlowVlanPriorityMetricTag) (PatternFlowVlanPriorityMetricTag, error)
	// FromPbText unmarshals PatternFlowVlanPriorityMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanPriorityMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanPriorityMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanPriorityMetricTag) Marshal() marshalPatternFlowVlanPriorityMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanPriorityMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanPriorityMetricTag) Unmarshal() unMarshalPatternFlowVlanPriorityMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanPriorityMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanPriorityMetricTag) ToProto() (*otg.PatternFlowVlanPriorityMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanPriorityMetricTag) FromProto(msg *otg.PatternFlowVlanPriorityMetricTag) (PatternFlowVlanPriorityMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanPriorityMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriorityMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanPriorityMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriorityMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanPriorityMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVlanPriorityMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriorityMetricTag) FromJson(value string) error {
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

func (obj *patternFlowVlanPriorityMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanPriorityMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanPriorityMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanPriorityMetricTag) Clone() (PatternFlowVlanPriorityMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanPriorityMetricTag()
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

// PatternFlowVlanPriorityMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowVlanPriorityMetricTag interface {
	Validation
	// msg marshals PatternFlowVlanPriorityMetricTag to protobuf object *otg.PatternFlowVlanPriorityMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanPriorityMetricTag
	// setMsg unmarshals PatternFlowVlanPriorityMetricTag from protobuf object *otg.PatternFlowVlanPriorityMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowVlanPriorityMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanPriorityMetricTag
	// validate validates PatternFlowVlanPriorityMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanPriorityMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowVlanPriorityMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowVlanPriorityMetricTag
	SetName(value string) PatternFlowVlanPriorityMetricTag
	// Offset returns uint32, set in PatternFlowVlanPriorityMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowVlanPriorityMetricTag
	SetOffset(value uint32) PatternFlowVlanPriorityMetricTag
	// HasOffset checks if Offset has been set in PatternFlowVlanPriorityMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowVlanPriorityMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowVlanPriorityMetricTag
	SetLength(value uint32) PatternFlowVlanPriorityMetricTag
	// HasLength checks if Length has been set in PatternFlowVlanPriorityMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowVlanPriorityMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowVlanPriorityMetricTag object
func (obj *patternFlowVlanPriorityMetricTag) SetName(value string) PatternFlowVlanPriorityMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVlanPriorityMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVlanPriorityMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowVlanPriorityMetricTag object
func (obj *patternFlowVlanPriorityMetricTag) SetOffset(value uint32) PatternFlowVlanPriorityMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVlanPriorityMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVlanPriorityMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowVlanPriorityMetricTag object
func (obj *patternFlowVlanPriorityMetricTag) SetLength(value uint32) PatternFlowVlanPriorityMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowVlanPriorityMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowVlanPriorityMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanPriorityMetricTag.Offset <= 2 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowVlanPriorityMetricTag.Length <= 3 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowVlanPriorityMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(3)
	}

}
