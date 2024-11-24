package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoCodeMetricTag *****
type patternFlowIcmpv6EchoCodeMetricTag struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoCodeMetricTag
	marshaller   marshalPatternFlowIcmpv6EchoCodeMetricTag
	unMarshaller unMarshalPatternFlowIcmpv6EchoCodeMetricTag
}

func NewPatternFlowIcmpv6EchoCodeMetricTag() PatternFlowIcmpv6EchoCodeMetricTag {
	obj := patternFlowIcmpv6EchoCodeMetricTag{obj: &otg.PatternFlowIcmpv6EchoCodeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) msg() *otg.PatternFlowIcmpv6EchoCodeMetricTag {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) setMsg(msg *otg.PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoCodeMetricTag struct {
	obj *patternFlowIcmpv6EchoCodeMetricTag
}

type marshalPatternFlowIcmpv6EchoCodeMetricTag interface {
	// ToProto marshals PatternFlowIcmpv6EchoCodeMetricTag to protobuf object *otg.PatternFlowIcmpv6EchoCodeMetricTag
	ToProto() (*otg.PatternFlowIcmpv6EchoCodeMetricTag, error)
	// ToPbText marshals PatternFlowIcmpv6EchoCodeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoCodeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoCodeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpv6EchoCodeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoCodeMetricTag struct {
	obj *patternFlowIcmpv6EchoCodeMetricTag
}

type unMarshalPatternFlowIcmpv6EchoCodeMetricTag interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoCodeMetricTag from protobuf object *otg.PatternFlowIcmpv6EchoCodeMetricTag
	FromProto(msg *otg.PatternFlowIcmpv6EchoCodeMetricTag) (PatternFlowIcmpv6EchoCodeMetricTag, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoCodeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoCodeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoCodeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) Marshal() marshalPatternFlowIcmpv6EchoCodeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoCodeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) Unmarshal() unMarshalPatternFlowIcmpv6EchoCodeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoCodeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoCodeMetricTag) ToProto() (*otg.PatternFlowIcmpv6EchoCodeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoCodeMetricTag) FromProto(msg *otg.PatternFlowIcmpv6EchoCodeMetricTag) (PatternFlowIcmpv6EchoCodeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoCodeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCodeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoCodeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCodeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoCodeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpv6EchoCodeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCodeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoCodeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) Clone() (PatternFlowIcmpv6EchoCodeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoCodeMetricTag()
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

// PatternFlowIcmpv6EchoCodeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIcmpv6EchoCodeMetricTag interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoCodeMetricTag to protobuf object *otg.PatternFlowIcmpv6EchoCodeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoCodeMetricTag
	// setMsg unmarshals PatternFlowIcmpv6EchoCodeMetricTag from protobuf object *otg.PatternFlowIcmpv6EchoCodeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoCodeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoCodeMetricTag
	// validate validates PatternFlowIcmpv6EchoCodeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoCodeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIcmpv6EchoCodeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIcmpv6EchoCodeMetricTag
	SetName(value string) PatternFlowIcmpv6EchoCodeMetricTag
	// Offset returns uint32, set in PatternFlowIcmpv6EchoCodeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIcmpv6EchoCodeMetricTag
	SetOffset(value uint32) PatternFlowIcmpv6EchoCodeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIcmpv6EchoCodeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIcmpv6EchoCodeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIcmpv6EchoCodeMetricTag
	SetLength(value uint32) PatternFlowIcmpv6EchoCodeMetricTag
	// HasLength checks if Length has been set in PatternFlowIcmpv6EchoCodeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIcmpv6EchoCodeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIcmpv6EchoCodeMetricTag object
func (obj *patternFlowIcmpv6EchoCodeMetricTag) SetName(value string) PatternFlowIcmpv6EchoCodeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpv6EchoCodeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpv6EchoCodeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIcmpv6EchoCodeMetricTag object
func (obj *patternFlowIcmpv6EchoCodeMetricTag) SetOffset(value uint32) PatternFlowIcmpv6EchoCodeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpv6EchoCodeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpv6EchoCodeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIcmpv6EchoCodeMetricTag object
func (obj *patternFlowIcmpv6EchoCodeMetricTag) SetLength(value uint32) PatternFlowIcmpv6EchoCodeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIcmpv6EchoCodeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCodeMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIcmpv6EchoCodeMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIcmpv6EchoCodeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
