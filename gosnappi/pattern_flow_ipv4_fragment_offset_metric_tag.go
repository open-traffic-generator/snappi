package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4FragmentOffsetMetricTag *****
type patternFlowIpv4FragmentOffsetMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4FragmentOffsetMetricTag
	marshaller   marshalPatternFlowIpv4FragmentOffsetMetricTag
	unMarshaller unMarshalPatternFlowIpv4FragmentOffsetMetricTag
}

func NewPatternFlowIpv4FragmentOffsetMetricTag() PatternFlowIpv4FragmentOffsetMetricTag {
	obj := patternFlowIpv4FragmentOffsetMetricTag{obj: &otg.PatternFlowIpv4FragmentOffsetMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) msg() *otg.PatternFlowIpv4FragmentOffsetMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) setMsg(msg *otg.PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4FragmentOffsetMetricTag struct {
	obj *patternFlowIpv4FragmentOffsetMetricTag
}

type marshalPatternFlowIpv4FragmentOffsetMetricTag interface {
	// ToProto marshals PatternFlowIpv4FragmentOffsetMetricTag to protobuf object *otg.PatternFlowIpv4FragmentOffsetMetricTag
	ToProto() (*otg.PatternFlowIpv4FragmentOffsetMetricTag, error)
	// ToPbText marshals PatternFlowIpv4FragmentOffsetMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4FragmentOffsetMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4FragmentOffsetMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4FragmentOffsetMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4FragmentOffsetMetricTag struct {
	obj *patternFlowIpv4FragmentOffsetMetricTag
}

type unMarshalPatternFlowIpv4FragmentOffsetMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4FragmentOffsetMetricTag from protobuf object *otg.PatternFlowIpv4FragmentOffsetMetricTag
	FromProto(msg *otg.PatternFlowIpv4FragmentOffsetMetricTag) (PatternFlowIpv4FragmentOffsetMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4FragmentOffsetMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4FragmentOffsetMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4FragmentOffsetMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) Marshal() marshalPatternFlowIpv4FragmentOffsetMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4FragmentOffsetMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) Unmarshal() unMarshalPatternFlowIpv4FragmentOffsetMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4FragmentOffsetMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4FragmentOffsetMetricTag) ToProto() (*otg.PatternFlowIpv4FragmentOffsetMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4FragmentOffsetMetricTag) FromProto(msg *otg.PatternFlowIpv4FragmentOffsetMetricTag) (PatternFlowIpv4FragmentOffsetMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4FragmentOffsetMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffsetMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4FragmentOffsetMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffsetMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4FragmentOffsetMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4FragmentOffsetMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffsetMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4FragmentOffsetMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) Clone() (PatternFlowIpv4FragmentOffsetMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4FragmentOffsetMetricTag()
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

// PatternFlowIpv4FragmentOffsetMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4FragmentOffsetMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4FragmentOffsetMetricTag to protobuf object *otg.PatternFlowIpv4FragmentOffsetMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4FragmentOffsetMetricTag
	// setMsg unmarshals PatternFlowIpv4FragmentOffsetMetricTag from protobuf object *otg.PatternFlowIpv4FragmentOffsetMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4FragmentOffsetMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4FragmentOffsetMetricTag
	// validate validates PatternFlowIpv4FragmentOffsetMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4FragmentOffsetMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4FragmentOffsetMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4FragmentOffsetMetricTag
	SetName(value string) PatternFlowIpv4FragmentOffsetMetricTag
	// Offset returns uint32, set in PatternFlowIpv4FragmentOffsetMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4FragmentOffsetMetricTag
	SetOffset(value uint32) PatternFlowIpv4FragmentOffsetMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4FragmentOffsetMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4FragmentOffsetMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4FragmentOffsetMetricTag
	SetLength(value uint32) PatternFlowIpv4FragmentOffsetMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4FragmentOffsetMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4FragmentOffsetMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4FragmentOffsetMetricTag object
func (obj *patternFlowIpv4FragmentOffsetMetricTag) SetName(value string) PatternFlowIpv4FragmentOffsetMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4FragmentOffsetMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4FragmentOffsetMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4FragmentOffsetMetricTag object
func (obj *patternFlowIpv4FragmentOffsetMetricTag) SetOffset(value uint32) PatternFlowIpv4FragmentOffsetMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4FragmentOffsetMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4FragmentOffsetMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4FragmentOffsetMetricTag object
func (obj *patternFlowIpv4FragmentOffsetMetricTag) SetLength(value uint32) PatternFlowIpv4FragmentOffsetMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4FragmentOffsetMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffsetMetricTag.Offset <= 4 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 5 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4FragmentOffsetMetricTag.Length <= 5 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4FragmentOffsetMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(5)
	}

}
