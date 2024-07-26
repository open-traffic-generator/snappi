package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsLabelMetricTag *****
type patternFlowMplsLabelMetricTag struct {
	validation
	obj          *otg.PatternFlowMplsLabelMetricTag
	marshaller   marshalPatternFlowMplsLabelMetricTag
	unMarshaller unMarshalPatternFlowMplsLabelMetricTag
}

func NewPatternFlowMplsLabelMetricTag() PatternFlowMplsLabelMetricTag {
	obj := patternFlowMplsLabelMetricTag{obj: &otg.PatternFlowMplsLabelMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsLabelMetricTag) msg() *otg.PatternFlowMplsLabelMetricTag {
	return obj.obj
}

func (obj *patternFlowMplsLabelMetricTag) setMsg(msg *otg.PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsLabelMetricTag struct {
	obj *patternFlowMplsLabelMetricTag
}

type marshalPatternFlowMplsLabelMetricTag interface {
	// ToProto marshals PatternFlowMplsLabelMetricTag to protobuf object *otg.PatternFlowMplsLabelMetricTag
	ToProto() (*otg.PatternFlowMplsLabelMetricTag, error)
	// ToPbText marshals PatternFlowMplsLabelMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsLabelMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsLabelMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowMplsLabelMetricTag struct {
	obj *patternFlowMplsLabelMetricTag
}

type unMarshalPatternFlowMplsLabelMetricTag interface {
	// FromProto unmarshals PatternFlowMplsLabelMetricTag from protobuf object *otg.PatternFlowMplsLabelMetricTag
	FromProto(msg *otg.PatternFlowMplsLabelMetricTag) (PatternFlowMplsLabelMetricTag, error)
	// FromPbText unmarshals PatternFlowMplsLabelMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsLabelMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsLabelMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsLabelMetricTag) Marshal() marshalPatternFlowMplsLabelMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsLabelMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsLabelMetricTag) Unmarshal() unMarshalPatternFlowMplsLabelMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsLabelMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsLabelMetricTag) ToProto() (*otg.PatternFlowMplsLabelMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsLabelMetricTag) FromProto(msg *otg.PatternFlowMplsLabelMetricTag) (PatternFlowMplsLabelMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsLabelMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabelMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsLabelMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabelMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsLabelMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabelMetricTag) FromJson(value string) error {
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

func (obj *patternFlowMplsLabelMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsLabelMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsLabelMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsLabelMetricTag) Clone() (PatternFlowMplsLabelMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsLabelMetricTag()
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

// PatternFlowMplsLabelMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowMplsLabelMetricTag interface {
	Validation
	// msg marshals PatternFlowMplsLabelMetricTag to protobuf object *otg.PatternFlowMplsLabelMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsLabelMetricTag
	// setMsg unmarshals PatternFlowMplsLabelMetricTag from protobuf object *otg.PatternFlowMplsLabelMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowMplsLabelMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsLabelMetricTag
	// validate validates PatternFlowMplsLabelMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsLabelMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowMplsLabelMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowMplsLabelMetricTag
	SetName(value string) PatternFlowMplsLabelMetricTag
	// Offset returns uint32, set in PatternFlowMplsLabelMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowMplsLabelMetricTag
	SetOffset(value uint32) PatternFlowMplsLabelMetricTag
	// HasOffset checks if Offset has been set in PatternFlowMplsLabelMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowMplsLabelMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowMplsLabelMetricTag
	SetLength(value uint32) PatternFlowMplsLabelMetricTag
	// HasLength checks if Length has been set in PatternFlowMplsLabelMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowMplsLabelMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowMplsLabelMetricTag object
func (obj *patternFlowMplsLabelMetricTag) SetName(value string) PatternFlowMplsLabelMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowMplsLabelMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowMplsLabelMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowMplsLabelMetricTag object
func (obj *patternFlowMplsLabelMetricTag) SetOffset(value uint32) PatternFlowMplsLabelMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowMplsLabelMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowMplsLabelMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowMplsLabelMetricTag object
func (obj *patternFlowMplsLabelMetricTag) SetLength(value uint32) PatternFlowMplsLabelMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowMplsLabelMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowMplsLabelMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 19 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsLabelMetricTag.Offset <= 19 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 20 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowMplsLabelMetricTag.Length <= 20 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowMplsLabelMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(20)
	}

}
