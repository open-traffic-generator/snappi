package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoCodeMetricTag *****
type patternFlowIcmpEchoCodeMetricTag struct {
	validation
	obj          *otg.PatternFlowIcmpEchoCodeMetricTag
	marshaller   marshalPatternFlowIcmpEchoCodeMetricTag
	unMarshaller unMarshalPatternFlowIcmpEchoCodeMetricTag
}

func NewPatternFlowIcmpEchoCodeMetricTag() PatternFlowIcmpEchoCodeMetricTag {
	obj := patternFlowIcmpEchoCodeMetricTag{obj: &otg.PatternFlowIcmpEchoCodeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoCodeMetricTag) msg() *otg.PatternFlowIcmpEchoCodeMetricTag {
	return obj.obj
}

func (obj *patternFlowIcmpEchoCodeMetricTag) setMsg(msg *otg.PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoCodeMetricTag struct {
	obj *patternFlowIcmpEchoCodeMetricTag
}

type marshalPatternFlowIcmpEchoCodeMetricTag interface {
	// ToProto marshals PatternFlowIcmpEchoCodeMetricTag to protobuf object *otg.PatternFlowIcmpEchoCodeMetricTag
	ToProto() (*otg.PatternFlowIcmpEchoCodeMetricTag, error)
	// ToPbText marshals PatternFlowIcmpEchoCodeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoCodeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoCodeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpEchoCodeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpEchoCodeMetricTag struct {
	obj *patternFlowIcmpEchoCodeMetricTag
}

type unMarshalPatternFlowIcmpEchoCodeMetricTag interface {
	// FromProto unmarshals PatternFlowIcmpEchoCodeMetricTag from protobuf object *otg.PatternFlowIcmpEchoCodeMetricTag
	FromProto(msg *otg.PatternFlowIcmpEchoCodeMetricTag) (PatternFlowIcmpEchoCodeMetricTag, error)
	// FromPbText unmarshals PatternFlowIcmpEchoCodeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoCodeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoCodeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoCodeMetricTag) Marshal() marshalPatternFlowIcmpEchoCodeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoCodeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoCodeMetricTag) Unmarshal() unMarshalPatternFlowIcmpEchoCodeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoCodeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoCodeMetricTag) ToProto() (*otg.PatternFlowIcmpEchoCodeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoCodeMetricTag) FromProto(msg *otg.PatternFlowIcmpEchoCodeMetricTag) (PatternFlowIcmpEchoCodeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoCodeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCodeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoCodeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCodeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoCodeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpEchoCodeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCodeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoCodeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoCodeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoCodeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoCodeMetricTag) Clone() (PatternFlowIcmpEchoCodeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoCodeMetricTag()
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

// PatternFlowIcmpEchoCodeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIcmpEchoCodeMetricTag interface {
	Validation
	// msg marshals PatternFlowIcmpEchoCodeMetricTag to protobuf object *otg.PatternFlowIcmpEchoCodeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoCodeMetricTag
	// setMsg unmarshals PatternFlowIcmpEchoCodeMetricTag from protobuf object *otg.PatternFlowIcmpEchoCodeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoCodeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoCodeMetricTag
	// validate validates PatternFlowIcmpEchoCodeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoCodeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIcmpEchoCodeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIcmpEchoCodeMetricTag
	SetName(value string) PatternFlowIcmpEchoCodeMetricTag
	// Offset returns uint32, set in PatternFlowIcmpEchoCodeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIcmpEchoCodeMetricTag
	SetOffset(value uint32) PatternFlowIcmpEchoCodeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIcmpEchoCodeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIcmpEchoCodeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIcmpEchoCodeMetricTag
	SetLength(value uint32) PatternFlowIcmpEchoCodeMetricTag
	// HasLength checks if Length has been set in PatternFlowIcmpEchoCodeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIcmpEchoCodeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIcmpEchoCodeMetricTag object
func (obj *patternFlowIcmpEchoCodeMetricTag) SetName(value string) PatternFlowIcmpEchoCodeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpEchoCodeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpEchoCodeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIcmpEchoCodeMetricTag object
func (obj *patternFlowIcmpEchoCodeMetricTag) SetOffset(value uint32) PatternFlowIcmpEchoCodeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpEchoCodeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpEchoCodeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIcmpEchoCodeMetricTag object
func (obj *patternFlowIcmpEchoCodeMetricTag) SetLength(value uint32) PatternFlowIcmpEchoCodeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIcmpEchoCodeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIcmpEchoCodeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoCodeMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIcmpEchoCodeMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIcmpEchoCodeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
