package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpOperationMetricTag *****
type patternFlowArpOperationMetricTag struct {
	validation
	obj          *otg.PatternFlowArpOperationMetricTag
	marshaller   marshalPatternFlowArpOperationMetricTag
	unMarshaller unMarshalPatternFlowArpOperationMetricTag
}

func NewPatternFlowArpOperationMetricTag() PatternFlowArpOperationMetricTag {
	obj := patternFlowArpOperationMetricTag{obj: &otg.PatternFlowArpOperationMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpOperationMetricTag) msg() *otg.PatternFlowArpOperationMetricTag {
	return obj.obj
}

func (obj *patternFlowArpOperationMetricTag) setMsg(msg *otg.PatternFlowArpOperationMetricTag) PatternFlowArpOperationMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpOperationMetricTag struct {
	obj *patternFlowArpOperationMetricTag
}

type marshalPatternFlowArpOperationMetricTag interface {
	// ToProto marshals PatternFlowArpOperationMetricTag to protobuf object *otg.PatternFlowArpOperationMetricTag
	ToProto() (*otg.PatternFlowArpOperationMetricTag, error)
	// ToPbText marshals PatternFlowArpOperationMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpOperationMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpOperationMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpOperationMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpOperationMetricTag struct {
	obj *patternFlowArpOperationMetricTag
}

type unMarshalPatternFlowArpOperationMetricTag interface {
	// FromProto unmarshals PatternFlowArpOperationMetricTag from protobuf object *otg.PatternFlowArpOperationMetricTag
	FromProto(msg *otg.PatternFlowArpOperationMetricTag) (PatternFlowArpOperationMetricTag, error)
	// FromPbText unmarshals PatternFlowArpOperationMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpOperationMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpOperationMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpOperationMetricTag) Marshal() marshalPatternFlowArpOperationMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpOperationMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpOperationMetricTag) Unmarshal() unMarshalPatternFlowArpOperationMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpOperationMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpOperationMetricTag) ToProto() (*otg.PatternFlowArpOperationMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpOperationMetricTag) FromProto(msg *otg.PatternFlowArpOperationMetricTag) (PatternFlowArpOperationMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpOperationMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpOperationMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpOperationMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpOperationMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpOperationMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpOperationMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpOperationMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpOperationMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpOperationMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpOperationMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpOperationMetricTag) Clone() (PatternFlowArpOperationMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpOperationMetricTag()
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

// PatternFlowArpOperationMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpOperationMetricTag interface {
	Validation
	// msg marshals PatternFlowArpOperationMetricTag to protobuf object *otg.PatternFlowArpOperationMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpOperationMetricTag
	// setMsg unmarshals PatternFlowArpOperationMetricTag from protobuf object *otg.PatternFlowArpOperationMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpOperationMetricTag) PatternFlowArpOperationMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpOperationMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpOperationMetricTag
	// validate validates PatternFlowArpOperationMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpOperationMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpOperationMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpOperationMetricTag
	SetName(value string) PatternFlowArpOperationMetricTag
	// Offset returns uint32, set in PatternFlowArpOperationMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpOperationMetricTag
	SetOffset(value uint32) PatternFlowArpOperationMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpOperationMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpOperationMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpOperationMetricTag
	SetLength(value uint32) PatternFlowArpOperationMetricTag
	// HasLength checks if Length has been set in PatternFlowArpOperationMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpOperationMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpOperationMetricTag object
func (obj *patternFlowArpOperationMetricTag) SetName(value string) PatternFlowArpOperationMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpOperationMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpOperationMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpOperationMetricTag object
func (obj *patternFlowArpOperationMetricTag) SetOffset(value uint32) PatternFlowArpOperationMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpOperationMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpOperationMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpOperationMetricTag object
func (obj *patternFlowArpOperationMetricTag) SetLength(value uint32) PatternFlowArpOperationMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpOperationMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpOperationMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpOperationMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpOperationMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpOperationMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
