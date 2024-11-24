package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoTypeMetricTag *****
type patternFlowIcmpv6EchoTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoTypeMetricTag
	marshaller   marshalPatternFlowIcmpv6EchoTypeMetricTag
	unMarshaller unMarshalPatternFlowIcmpv6EchoTypeMetricTag
}

func NewPatternFlowIcmpv6EchoTypeMetricTag() PatternFlowIcmpv6EchoTypeMetricTag {
	obj := patternFlowIcmpv6EchoTypeMetricTag{obj: &otg.PatternFlowIcmpv6EchoTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) msg() *otg.PatternFlowIcmpv6EchoTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) setMsg(msg *otg.PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoTypeMetricTag struct {
	obj *patternFlowIcmpv6EchoTypeMetricTag
}

type marshalPatternFlowIcmpv6EchoTypeMetricTag interface {
	// ToProto marshals PatternFlowIcmpv6EchoTypeMetricTag to protobuf object *otg.PatternFlowIcmpv6EchoTypeMetricTag
	ToProto() (*otg.PatternFlowIcmpv6EchoTypeMetricTag, error)
	// ToPbText marshals PatternFlowIcmpv6EchoTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoTypeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpv6EchoTypeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoTypeMetricTag struct {
	obj *patternFlowIcmpv6EchoTypeMetricTag
}

type unMarshalPatternFlowIcmpv6EchoTypeMetricTag interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoTypeMetricTag from protobuf object *otg.PatternFlowIcmpv6EchoTypeMetricTag
	FromProto(msg *otg.PatternFlowIcmpv6EchoTypeMetricTag) (PatternFlowIcmpv6EchoTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) Marshal() marshalPatternFlowIcmpv6EchoTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) Unmarshal() unMarshalPatternFlowIcmpv6EchoTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoTypeMetricTag) ToProto() (*otg.PatternFlowIcmpv6EchoTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoTypeMetricTag) FromProto(msg *otg.PatternFlowIcmpv6EchoTypeMetricTag) (PatternFlowIcmpv6EchoTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoTypeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpv6EchoTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) Clone() (PatternFlowIcmpv6EchoTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoTypeMetricTag()
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

// PatternFlowIcmpv6EchoTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIcmpv6EchoTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoTypeMetricTag to protobuf object *otg.PatternFlowIcmpv6EchoTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoTypeMetricTag
	// setMsg unmarshals PatternFlowIcmpv6EchoTypeMetricTag from protobuf object *otg.PatternFlowIcmpv6EchoTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoTypeMetricTag
	// validate validates PatternFlowIcmpv6EchoTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIcmpv6EchoTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIcmpv6EchoTypeMetricTag
	SetName(value string) PatternFlowIcmpv6EchoTypeMetricTag
	// Offset returns uint32, set in PatternFlowIcmpv6EchoTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIcmpv6EchoTypeMetricTag
	SetOffset(value uint32) PatternFlowIcmpv6EchoTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIcmpv6EchoTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIcmpv6EchoTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIcmpv6EchoTypeMetricTag
	SetLength(value uint32) PatternFlowIcmpv6EchoTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowIcmpv6EchoTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIcmpv6EchoTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIcmpv6EchoTypeMetricTag object
func (obj *patternFlowIcmpv6EchoTypeMetricTag) SetName(value string) PatternFlowIcmpv6EchoTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpv6EchoTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpv6EchoTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIcmpv6EchoTypeMetricTag object
func (obj *patternFlowIcmpv6EchoTypeMetricTag) SetOffset(value uint32) PatternFlowIcmpv6EchoTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpv6EchoTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpv6EchoTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIcmpv6EchoTypeMetricTag object
func (obj *patternFlowIcmpv6EchoTypeMetricTag) SetLength(value uint32) PatternFlowIcmpv6EchoTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIcmpv6EchoTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoTypeMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIcmpv6EchoTypeMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIcmpv6EchoTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
