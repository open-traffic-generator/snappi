package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanFlagsMetricTag *****
type patternFlowVxlanFlagsMetricTag struct {
	validation
	obj          *otg.PatternFlowVxlanFlagsMetricTag
	marshaller   marshalPatternFlowVxlanFlagsMetricTag
	unMarshaller unMarshalPatternFlowVxlanFlagsMetricTag
}

func NewPatternFlowVxlanFlagsMetricTag() PatternFlowVxlanFlagsMetricTag {
	obj := patternFlowVxlanFlagsMetricTag{obj: &otg.PatternFlowVxlanFlagsMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanFlagsMetricTag) msg() *otg.PatternFlowVxlanFlagsMetricTag {
	return obj.obj
}

func (obj *patternFlowVxlanFlagsMetricTag) setMsg(msg *otg.PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanFlagsMetricTag struct {
	obj *patternFlowVxlanFlagsMetricTag
}

type marshalPatternFlowVxlanFlagsMetricTag interface {
	// ToProto marshals PatternFlowVxlanFlagsMetricTag to protobuf object *otg.PatternFlowVxlanFlagsMetricTag
	ToProto() (*otg.PatternFlowVxlanFlagsMetricTag, error)
	// ToPbText marshals PatternFlowVxlanFlagsMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanFlagsMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanFlagsMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVxlanFlagsMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVxlanFlagsMetricTag struct {
	obj *patternFlowVxlanFlagsMetricTag
}

type unMarshalPatternFlowVxlanFlagsMetricTag interface {
	// FromProto unmarshals PatternFlowVxlanFlagsMetricTag from protobuf object *otg.PatternFlowVxlanFlagsMetricTag
	FromProto(msg *otg.PatternFlowVxlanFlagsMetricTag) (PatternFlowVxlanFlagsMetricTag, error)
	// FromPbText unmarshals PatternFlowVxlanFlagsMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanFlagsMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanFlagsMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanFlagsMetricTag) Marshal() marshalPatternFlowVxlanFlagsMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanFlagsMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanFlagsMetricTag) Unmarshal() unMarshalPatternFlowVxlanFlagsMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanFlagsMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanFlagsMetricTag) ToProto() (*otg.PatternFlowVxlanFlagsMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanFlagsMetricTag) FromProto(msg *otg.PatternFlowVxlanFlagsMetricTag) (PatternFlowVxlanFlagsMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanFlagsMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlagsMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanFlagsMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlagsMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanFlagsMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVxlanFlagsMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlagsMetricTag) FromJson(value string) error {
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

func (obj *patternFlowVxlanFlagsMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanFlagsMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanFlagsMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanFlagsMetricTag) Clone() (PatternFlowVxlanFlagsMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanFlagsMetricTag()
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

// PatternFlowVxlanFlagsMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowVxlanFlagsMetricTag interface {
	Validation
	// msg marshals PatternFlowVxlanFlagsMetricTag to protobuf object *otg.PatternFlowVxlanFlagsMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanFlagsMetricTag
	// setMsg unmarshals PatternFlowVxlanFlagsMetricTag from protobuf object *otg.PatternFlowVxlanFlagsMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanFlagsMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanFlagsMetricTag
	// validate validates PatternFlowVxlanFlagsMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanFlagsMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowVxlanFlagsMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowVxlanFlagsMetricTag
	SetName(value string) PatternFlowVxlanFlagsMetricTag
	// Offset returns uint32, set in PatternFlowVxlanFlagsMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowVxlanFlagsMetricTag
	SetOffset(value uint32) PatternFlowVxlanFlagsMetricTag
	// HasOffset checks if Offset has been set in PatternFlowVxlanFlagsMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowVxlanFlagsMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowVxlanFlagsMetricTag
	SetLength(value uint32) PatternFlowVxlanFlagsMetricTag
	// HasLength checks if Length has been set in PatternFlowVxlanFlagsMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowVxlanFlagsMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowVxlanFlagsMetricTag object
func (obj *patternFlowVxlanFlagsMetricTag) SetName(value string) PatternFlowVxlanFlagsMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVxlanFlagsMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVxlanFlagsMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowVxlanFlagsMetricTag object
func (obj *patternFlowVxlanFlagsMetricTag) SetOffset(value uint32) PatternFlowVxlanFlagsMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVxlanFlagsMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVxlanFlagsMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowVxlanFlagsMetricTag object
func (obj *patternFlowVxlanFlagsMetricTag) SetLength(value uint32) PatternFlowVxlanFlagsMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowVxlanFlagsMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowVxlanFlagsMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanFlagsMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowVxlanFlagsMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowVxlanFlagsMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
