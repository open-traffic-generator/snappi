package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpHardwareLengthMetricTag *****
type patternFlowArpHardwareLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowArpHardwareLengthMetricTag
	marshaller   marshalPatternFlowArpHardwareLengthMetricTag
	unMarshaller unMarshalPatternFlowArpHardwareLengthMetricTag
}

func NewPatternFlowArpHardwareLengthMetricTag() PatternFlowArpHardwareLengthMetricTag {
	obj := patternFlowArpHardwareLengthMetricTag{obj: &otg.PatternFlowArpHardwareLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpHardwareLengthMetricTag) msg() *otg.PatternFlowArpHardwareLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowArpHardwareLengthMetricTag) setMsg(msg *otg.PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpHardwareLengthMetricTag struct {
	obj *patternFlowArpHardwareLengthMetricTag
}

type marshalPatternFlowArpHardwareLengthMetricTag interface {
	// ToProto marshals PatternFlowArpHardwareLengthMetricTag to protobuf object *otg.PatternFlowArpHardwareLengthMetricTag
	ToProto() (*otg.PatternFlowArpHardwareLengthMetricTag, error)
	// ToPbText marshals PatternFlowArpHardwareLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpHardwareLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpHardwareLengthMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpHardwareLengthMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpHardwareLengthMetricTag struct {
	obj *patternFlowArpHardwareLengthMetricTag
}

type unMarshalPatternFlowArpHardwareLengthMetricTag interface {
	// FromProto unmarshals PatternFlowArpHardwareLengthMetricTag from protobuf object *otg.PatternFlowArpHardwareLengthMetricTag
	FromProto(msg *otg.PatternFlowArpHardwareLengthMetricTag) (PatternFlowArpHardwareLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowArpHardwareLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpHardwareLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpHardwareLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpHardwareLengthMetricTag) Marshal() marshalPatternFlowArpHardwareLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpHardwareLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpHardwareLengthMetricTag) Unmarshal() unMarshalPatternFlowArpHardwareLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpHardwareLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpHardwareLengthMetricTag) ToProto() (*otg.PatternFlowArpHardwareLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpHardwareLengthMetricTag) FromProto(msg *otg.PatternFlowArpHardwareLengthMetricTag) (PatternFlowArpHardwareLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpHardwareLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpHardwareLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpHardwareLengthMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpHardwareLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpHardwareLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpHardwareLengthMetricTag) Clone() (PatternFlowArpHardwareLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpHardwareLengthMetricTag()
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

// PatternFlowArpHardwareLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpHardwareLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowArpHardwareLengthMetricTag to protobuf object *otg.PatternFlowArpHardwareLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpHardwareLengthMetricTag
	// setMsg unmarshals PatternFlowArpHardwareLengthMetricTag from protobuf object *otg.PatternFlowArpHardwareLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpHardwareLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpHardwareLengthMetricTag
	// validate validates PatternFlowArpHardwareLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpHardwareLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpHardwareLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpHardwareLengthMetricTag
	SetName(value string) PatternFlowArpHardwareLengthMetricTag
	// Offset returns uint32, set in PatternFlowArpHardwareLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpHardwareLengthMetricTag
	SetOffset(value uint32) PatternFlowArpHardwareLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpHardwareLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpHardwareLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpHardwareLengthMetricTag
	SetLength(value uint32) PatternFlowArpHardwareLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowArpHardwareLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpHardwareLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpHardwareLengthMetricTag object
func (obj *patternFlowArpHardwareLengthMetricTag) SetName(value string) PatternFlowArpHardwareLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpHardwareLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpHardwareLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpHardwareLengthMetricTag object
func (obj *patternFlowArpHardwareLengthMetricTag) SetOffset(value uint32) PatternFlowArpHardwareLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpHardwareLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpHardwareLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpHardwareLengthMetricTag object
func (obj *patternFlowArpHardwareLengthMetricTag) SetLength(value uint32) PatternFlowArpHardwareLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpHardwareLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpHardwareLengthMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareLengthMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpHardwareLengthMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpHardwareLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
