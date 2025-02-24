package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanReserved1MetricTag *****
type patternFlowVxlanReserved1MetricTag struct {
	validation
	obj          *otg.PatternFlowVxlanReserved1MetricTag
	marshaller   marshalPatternFlowVxlanReserved1MetricTag
	unMarshaller unMarshalPatternFlowVxlanReserved1MetricTag
}

func NewPatternFlowVxlanReserved1MetricTag() PatternFlowVxlanReserved1MetricTag {
	obj := patternFlowVxlanReserved1MetricTag{obj: &otg.PatternFlowVxlanReserved1MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanReserved1MetricTag) msg() *otg.PatternFlowVxlanReserved1MetricTag {
	return obj.obj
}

func (obj *patternFlowVxlanReserved1MetricTag) setMsg(msg *otg.PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanReserved1MetricTag struct {
	obj *patternFlowVxlanReserved1MetricTag
}

type marshalPatternFlowVxlanReserved1MetricTag interface {
	// ToProto marshals PatternFlowVxlanReserved1MetricTag to protobuf object *otg.PatternFlowVxlanReserved1MetricTag
	ToProto() (*otg.PatternFlowVxlanReserved1MetricTag, error)
	// ToPbText marshals PatternFlowVxlanReserved1MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanReserved1MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanReserved1MetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVxlanReserved1MetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVxlanReserved1MetricTag struct {
	obj *patternFlowVxlanReserved1MetricTag
}

type unMarshalPatternFlowVxlanReserved1MetricTag interface {
	// FromProto unmarshals PatternFlowVxlanReserved1MetricTag from protobuf object *otg.PatternFlowVxlanReserved1MetricTag
	FromProto(msg *otg.PatternFlowVxlanReserved1MetricTag) (PatternFlowVxlanReserved1MetricTag, error)
	// FromPbText unmarshals PatternFlowVxlanReserved1MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanReserved1MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanReserved1MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanReserved1MetricTag) Marshal() marshalPatternFlowVxlanReserved1MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanReserved1MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanReserved1MetricTag) Unmarshal() unMarshalPatternFlowVxlanReserved1MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanReserved1MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanReserved1MetricTag) ToProto() (*otg.PatternFlowVxlanReserved1MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanReserved1MetricTag) FromProto(msg *otg.PatternFlowVxlanReserved1MetricTag) (PatternFlowVxlanReserved1MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanReserved1MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanReserved1MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanReserved1MetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVxlanReserved1MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1MetricTag) FromJson(value string) error {
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

func (obj *patternFlowVxlanReserved1MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved1MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved1MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanReserved1MetricTag) Clone() (PatternFlowVxlanReserved1MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanReserved1MetricTag()
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

// PatternFlowVxlanReserved1MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowVxlanReserved1MetricTag interface {
	Validation
	// msg marshals PatternFlowVxlanReserved1MetricTag to protobuf object *otg.PatternFlowVxlanReserved1MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanReserved1MetricTag
	// setMsg unmarshals PatternFlowVxlanReserved1MetricTag from protobuf object *otg.PatternFlowVxlanReserved1MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanReserved1MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanReserved1MetricTag
	// validate validates PatternFlowVxlanReserved1MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanReserved1MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowVxlanReserved1MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowVxlanReserved1MetricTag
	SetName(value string) PatternFlowVxlanReserved1MetricTag
	// Offset returns uint32, set in PatternFlowVxlanReserved1MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowVxlanReserved1MetricTag
	SetOffset(value uint32) PatternFlowVxlanReserved1MetricTag
	// HasOffset checks if Offset has been set in PatternFlowVxlanReserved1MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowVxlanReserved1MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowVxlanReserved1MetricTag
	SetLength(value uint32) PatternFlowVxlanReserved1MetricTag
	// HasLength checks if Length has been set in PatternFlowVxlanReserved1MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowVxlanReserved1MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowVxlanReserved1MetricTag object
func (obj *patternFlowVxlanReserved1MetricTag) SetName(value string) PatternFlowVxlanReserved1MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVxlanReserved1MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVxlanReserved1MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowVxlanReserved1MetricTag object
func (obj *patternFlowVxlanReserved1MetricTag) SetOffset(value uint32) PatternFlowVxlanReserved1MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVxlanReserved1MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVxlanReserved1MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowVxlanReserved1MetricTag object
func (obj *patternFlowVxlanReserved1MetricTag) SetLength(value uint32) PatternFlowVxlanReserved1MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowVxlanReserved1MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowVxlanReserved1MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved1MetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowVxlanReserved1MetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowVxlanReserved1MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
