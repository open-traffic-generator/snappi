package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpHardwareTypeMetricTag *****
type patternFlowArpHardwareTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowArpHardwareTypeMetricTag
	marshaller   marshalPatternFlowArpHardwareTypeMetricTag
	unMarshaller unMarshalPatternFlowArpHardwareTypeMetricTag
}

func NewPatternFlowArpHardwareTypeMetricTag() PatternFlowArpHardwareTypeMetricTag {
	obj := patternFlowArpHardwareTypeMetricTag{obj: &otg.PatternFlowArpHardwareTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpHardwareTypeMetricTag) msg() *otg.PatternFlowArpHardwareTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowArpHardwareTypeMetricTag) setMsg(msg *otg.PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpHardwareTypeMetricTag struct {
	obj *patternFlowArpHardwareTypeMetricTag
}

type marshalPatternFlowArpHardwareTypeMetricTag interface {
	// ToProto marshals PatternFlowArpHardwareTypeMetricTag to protobuf object *otg.PatternFlowArpHardwareTypeMetricTag
	ToProto() (*otg.PatternFlowArpHardwareTypeMetricTag, error)
	// ToPbText marshals PatternFlowArpHardwareTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpHardwareTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpHardwareTypeMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpHardwareTypeMetricTag struct {
	obj *patternFlowArpHardwareTypeMetricTag
}

type unMarshalPatternFlowArpHardwareTypeMetricTag interface {
	// FromProto unmarshals PatternFlowArpHardwareTypeMetricTag from protobuf object *otg.PatternFlowArpHardwareTypeMetricTag
	FromProto(msg *otg.PatternFlowArpHardwareTypeMetricTag) (PatternFlowArpHardwareTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowArpHardwareTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpHardwareTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpHardwareTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpHardwareTypeMetricTag) Marshal() marshalPatternFlowArpHardwareTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpHardwareTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpHardwareTypeMetricTag) Unmarshal() unMarshalPatternFlowArpHardwareTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpHardwareTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpHardwareTypeMetricTag) ToProto() (*otg.PatternFlowArpHardwareTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpHardwareTypeMetricTag) FromProto(msg *otg.PatternFlowArpHardwareTypeMetricTag) (PatternFlowArpHardwareTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpHardwareTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpHardwareTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpHardwareTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpHardwareTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpHardwareTypeMetricTag) Clone() (PatternFlowArpHardwareTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpHardwareTypeMetricTag()
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

// PatternFlowArpHardwareTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpHardwareTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowArpHardwareTypeMetricTag to protobuf object *otg.PatternFlowArpHardwareTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpHardwareTypeMetricTag
	// setMsg unmarshals PatternFlowArpHardwareTypeMetricTag from protobuf object *otg.PatternFlowArpHardwareTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpHardwareTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpHardwareTypeMetricTag
	// validate validates PatternFlowArpHardwareTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpHardwareTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpHardwareTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpHardwareTypeMetricTag
	SetName(value string) PatternFlowArpHardwareTypeMetricTag
	// Offset returns uint32, set in PatternFlowArpHardwareTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpHardwareTypeMetricTag
	SetOffset(value uint32) PatternFlowArpHardwareTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpHardwareTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpHardwareTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpHardwareTypeMetricTag
	SetLength(value uint32) PatternFlowArpHardwareTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowArpHardwareTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpHardwareTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpHardwareTypeMetricTag object
func (obj *patternFlowArpHardwareTypeMetricTag) SetName(value string) PatternFlowArpHardwareTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpHardwareTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpHardwareTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpHardwareTypeMetricTag object
func (obj *patternFlowArpHardwareTypeMetricTag) SetOffset(value uint32) PatternFlowArpHardwareTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpHardwareTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpHardwareTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpHardwareTypeMetricTag object
func (obj *patternFlowArpHardwareTypeMetricTag) SetLength(value uint32) PatternFlowArpHardwareTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpHardwareTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpHardwareTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareTypeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpHardwareTypeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpHardwareTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
