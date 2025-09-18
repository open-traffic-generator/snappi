package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsBottomOfStackMetricTag *****
type patternFlowMplsBottomOfStackMetricTag struct {
	validation
	obj          *otg.PatternFlowMplsBottomOfStackMetricTag
	marshaller   marshalPatternFlowMplsBottomOfStackMetricTag
	unMarshaller unMarshalPatternFlowMplsBottomOfStackMetricTag
}

func NewPatternFlowMplsBottomOfStackMetricTag() PatternFlowMplsBottomOfStackMetricTag {
	obj := patternFlowMplsBottomOfStackMetricTag{obj: &otg.PatternFlowMplsBottomOfStackMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsBottomOfStackMetricTag) msg() *otg.PatternFlowMplsBottomOfStackMetricTag {
	return obj.obj
}

func (obj *patternFlowMplsBottomOfStackMetricTag) setMsg(msg *otg.PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsBottomOfStackMetricTag struct {
	obj *patternFlowMplsBottomOfStackMetricTag
}

type marshalPatternFlowMplsBottomOfStackMetricTag interface {
	// ToProto marshals PatternFlowMplsBottomOfStackMetricTag to protobuf object *otg.PatternFlowMplsBottomOfStackMetricTag
	ToProto() (*otg.PatternFlowMplsBottomOfStackMetricTag, error)
	// ToPbText marshals PatternFlowMplsBottomOfStackMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsBottomOfStackMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsBottomOfStackMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowMplsBottomOfStackMetricTag struct {
	obj *patternFlowMplsBottomOfStackMetricTag
}

type unMarshalPatternFlowMplsBottomOfStackMetricTag interface {
	// FromProto unmarshals PatternFlowMplsBottomOfStackMetricTag from protobuf object *otg.PatternFlowMplsBottomOfStackMetricTag
	FromProto(msg *otg.PatternFlowMplsBottomOfStackMetricTag) (PatternFlowMplsBottomOfStackMetricTag, error)
	// FromPbText unmarshals PatternFlowMplsBottomOfStackMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsBottomOfStackMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsBottomOfStackMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsBottomOfStackMetricTag) Marshal() marshalPatternFlowMplsBottomOfStackMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsBottomOfStackMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsBottomOfStackMetricTag) Unmarshal() unMarshalPatternFlowMplsBottomOfStackMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsBottomOfStackMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsBottomOfStackMetricTag) ToProto() (*otg.PatternFlowMplsBottomOfStackMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsBottomOfStackMetricTag) FromProto(msg *otg.PatternFlowMplsBottomOfStackMetricTag) (PatternFlowMplsBottomOfStackMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsBottomOfStackMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStackMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsBottomOfStackMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStackMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsBottomOfStackMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStackMetricTag) FromJson(value string) error {
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

func (obj *patternFlowMplsBottomOfStackMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsBottomOfStackMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsBottomOfStackMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsBottomOfStackMetricTag) Clone() (PatternFlowMplsBottomOfStackMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsBottomOfStackMetricTag()
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

// PatternFlowMplsBottomOfStackMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowMplsBottomOfStackMetricTag interface {
	Validation
	// msg marshals PatternFlowMplsBottomOfStackMetricTag to protobuf object *otg.PatternFlowMplsBottomOfStackMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsBottomOfStackMetricTag
	// setMsg unmarshals PatternFlowMplsBottomOfStackMetricTag from protobuf object *otg.PatternFlowMplsBottomOfStackMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowMplsBottomOfStackMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsBottomOfStackMetricTag
	// validate validates PatternFlowMplsBottomOfStackMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsBottomOfStackMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowMplsBottomOfStackMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowMplsBottomOfStackMetricTag
	SetName(value string) PatternFlowMplsBottomOfStackMetricTag
	// Offset returns uint32, set in PatternFlowMplsBottomOfStackMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowMplsBottomOfStackMetricTag
	SetOffset(value uint32) PatternFlowMplsBottomOfStackMetricTag
	// HasOffset checks if Offset has been set in PatternFlowMplsBottomOfStackMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowMplsBottomOfStackMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowMplsBottomOfStackMetricTag
	SetLength(value uint32) PatternFlowMplsBottomOfStackMetricTag
	// HasLength checks if Length has been set in PatternFlowMplsBottomOfStackMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowMplsBottomOfStackMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowMplsBottomOfStackMetricTag object
func (obj *patternFlowMplsBottomOfStackMetricTag) SetName(value string) PatternFlowMplsBottomOfStackMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowMplsBottomOfStackMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowMplsBottomOfStackMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowMplsBottomOfStackMetricTag object
func (obj *patternFlowMplsBottomOfStackMetricTag) SetOffset(value uint32) PatternFlowMplsBottomOfStackMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowMplsBottomOfStackMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowMplsBottomOfStackMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowMplsBottomOfStackMetricTag object
func (obj *patternFlowMplsBottomOfStackMetricTag) SetLength(value uint32) PatternFlowMplsBottomOfStackMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowMplsBottomOfStackMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowMplsBottomOfStackMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStackMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowMplsBottomOfStackMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowMplsBottomOfStackMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
