package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoTypeMetricTag *****
type patternFlowIcmpEchoTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowIcmpEchoTypeMetricTag
	marshaller   marshalPatternFlowIcmpEchoTypeMetricTag
	unMarshaller unMarshalPatternFlowIcmpEchoTypeMetricTag
}

func NewPatternFlowIcmpEchoTypeMetricTag() PatternFlowIcmpEchoTypeMetricTag {
	obj := patternFlowIcmpEchoTypeMetricTag{obj: &otg.PatternFlowIcmpEchoTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoTypeMetricTag) msg() *otg.PatternFlowIcmpEchoTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowIcmpEchoTypeMetricTag) setMsg(msg *otg.PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoTypeMetricTag struct {
	obj *patternFlowIcmpEchoTypeMetricTag
}

type marshalPatternFlowIcmpEchoTypeMetricTag interface {
	// ToProto marshals PatternFlowIcmpEchoTypeMetricTag to protobuf object *otg.PatternFlowIcmpEchoTypeMetricTag
	ToProto() (*otg.PatternFlowIcmpEchoTypeMetricTag, error)
	// ToPbText marshals PatternFlowIcmpEchoTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoTypeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpEchoTypeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpEchoTypeMetricTag struct {
	obj *patternFlowIcmpEchoTypeMetricTag
}

type unMarshalPatternFlowIcmpEchoTypeMetricTag interface {
	// FromProto unmarshals PatternFlowIcmpEchoTypeMetricTag from protobuf object *otg.PatternFlowIcmpEchoTypeMetricTag
	FromProto(msg *otg.PatternFlowIcmpEchoTypeMetricTag) (PatternFlowIcmpEchoTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowIcmpEchoTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoTypeMetricTag) Marshal() marshalPatternFlowIcmpEchoTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoTypeMetricTag) Unmarshal() unMarshalPatternFlowIcmpEchoTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoTypeMetricTag) ToProto() (*otg.PatternFlowIcmpEchoTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoTypeMetricTag) FromProto(msg *otg.PatternFlowIcmpEchoTypeMetricTag) (PatternFlowIcmpEchoTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoTypeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpEchoTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoTypeMetricTag) Clone() (PatternFlowIcmpEchoTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoTypeMetricTag()
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

// PatternFlowIcmpEchoTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIcmpEchoTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowIcmpEchoTypeMetricTag to protobuf object *otg.PatternFlowIcmpEchoTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoTypeMetricTag
	// setMsg unmarshals PatternFlowIcmpEchoTypeMetricTag from protobuf object *otg.PatternFlowIcmpEchoTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoTypeMetricTag
	// validate validates PatternFlowIcmpEchoTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIcmpEchoTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIcmpEchoTypeMetricTag
	SetName(value string) PatternFlowIcmpEchoTypeMetricTag
	// Offset returns uint32, set in PatternFlowIcmpEchoTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIcmpEchoTypeMetricTag
	SetOffset(value uint32) PatternFlowIcmpEchoTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIcmpEchoTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIcmpEchoTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIcmpEchoTypeMetricTag
	SetLength(value uint32) PatternFlowIcmpEchoTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowIcmpEchoTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIcmpEchoTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIcmpEchoTypeMetricTag object
func (obj *patternFlowIcmpEchoTypeMetricTag) SetName(value string) PatternFlowIcmpEchoTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpEchoTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpEchoTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIcmpEchoTypeMetricTag object
func (obj *patternFlowIcmpEchoTypeMetricTag) SetOffset(value uint32) PatternFlowIcmpEchoTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpEchoTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpEchoTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIcmpEchoTypeMetricTag object
func (obj *patternFlowIcmpEchoTypeMetricTag) SetLength(value uint32) PatternFlowIcmpEchoTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIcmpEchoTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIcmpEchoTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoTypeMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIcmpEchoTypeMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIcmpEchoTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
