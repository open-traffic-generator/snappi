package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanCfiMetricTag *****
type patternFlowVlanCfiMetricTag struct {
	validation
	obj          *otg.PatternFlowVlanCfiMetricTag
	marshaller   marshalPatternFlowVlanCfiMetricTag
	unMarshaller unMarshalPatternFlowVlanCfiMetricTag
}

func NewPatternFlowVlanCfiMetricTag() PatternFlowVlanCfiMetricTag {
	obj := patternFlowVlanCfiMetricTag{obj: &otg.PatternFlowVlanCfiMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanCfiMetricTag) msg() *otg.PatternFlowVlanCfiMetricTag {
	return obj.obj
}

func (obj *patternFlowVlanCfiMetricTag) setMsg(msg *otg.PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanCfiMetricTag struct {
	obj *patternFlowVlanCfiMetricTag
}

type marshalPatternFlowVlanCfiMetricTag interface {
	// ToProto marshals PatternFlowVlanCfiMetricTag to protobuf object *otg.PatternFlowVlanCfiMetricTag
	ToProto() (*otg.PatternFlowVlanCfiMetricTag, error)
	// ToPbText marshals PatternFlowVlanCfiMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanCfiMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanCfiMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVlanCfiMetricTag struct {
	obj *patternFlowVlanCfiMetricTag
}

type unMarshalPatternFlowVlanCfiMetricTag interface {
	// FromProto unmarshals PatternFlowVlanCfiMetricTag from protobuf object *otg.PatternFlowVlanCfiMetricTag
	FromProto(msg *otg.PatternFlowVlanCfiMetricTag) (PatternFlowVlanCfiMetricTag, error)
	// FromPbText unmarshals PatternFlowVlanCfiMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanCfiMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanCfiMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanCfiMetricTag) Marshal() marshalPatternFlowVlanCfiMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanCfiMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanCfiMetricTag) Unmarshal() unMarshalPatternFlowVlanCfiMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanCfiMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanCfiMetricTag) ToProto() (*otg.PatternFlowVlanCfiMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanCfiMetricTag) FromProto(msg *otg.PatternFlowVlanCfiMetricTag) (PatternFlowVlanCfiMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanCfiMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfiMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanCfiMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfiMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanCfiMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfiMetricTag) FromJson(value string) error {
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

func (obj *patternFlowVlanCfiMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanCfiMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanCfiMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanCfiMetricTag) Clone() (PatternFlowVlanCfiMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanCfiMetricTag()
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

// PatternFlowVlanCfiMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowVlanCfiMetricTag interface {
	Validation
	// msg marshals PatternFlowVlanCfiMetricTag to protobuf object *otg.PatternFlowVlanCfiMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanCfiMetricTag
	// setMsg unmarshals PatternFlowVlanCfiMetricTag from protobuf object *otg.PatternFlowVlanCfiMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowVlanCfiMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanCfiMetricTag
	// validate validates PatternFlowVlanCfiMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanCfiMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowVlanCfiMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowVlanCfiMetricTag
	SetName(value string) PatternFlowVlanCfiMetricTag
	// Offset returns uint32, set in PatternFlowVlanCfiMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowVlanCfiMetricTag
	SetOffset(value uint32) PatternFlowVlanCfiMetricTag
	// HasOffset checks if Offset has been set in PatternFlowVlanCfiMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowVlanCfiMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowVlanCfiMetricTag
	SetLength(value uint32) PatternFlowVlanCfiMetricTag
	// HasLength checks if Length has been set in PatternFlowVlanCfiMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowVlanCfiMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowVlanCfiMetricTag object
func (obj *patternFlowVlanCfiMetricTag) SetName(value string) PatternFlowVlanCfiMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVlanCfiMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVlanCfiMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowVlanCfiMetricTag object
func (obj *patternFlowVlanCfiMetricTag) SetOffset(value uint32) PatternFlowVlanCfiMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVlanCfiMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVlanCfiMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowVlanCfiMetricTag object
func (obj *patternFlowVlanCfiMetricTag) SetLength(value uint32) PatternFlowVlanCfiMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowVlanCfiMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowVlanCfiMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanCfiMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowVlanCfiMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowVlanCfiMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
