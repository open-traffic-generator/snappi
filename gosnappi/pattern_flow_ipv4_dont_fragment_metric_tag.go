package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DontFragmentMetricTag *****
type patternFlowIpv4DontFragmentMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4DontFragmentMetricTag
	marshaller   marshalPatternFlowIpv4DontFragmentMetricTag
	unMarshaller unMarshalPatternFlowIpv4DontFragmentMetricTag
}

func NewPatternFlowIpv4DontFragmentMetricTag() PatternFlowIpv4DontFragmentMetricTag {
	obj := patternFlowIpv4DontFragmentMetricTag{obj: &otg.PatternFlowIpv4DontFragmentMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DontFragmentMetricTag) msg() *otg.PatternFlowIpv4DontFragmentMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4DontFragmentMetricTag) setMsg(msg *otg.PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DontFragmentMetricTag struct {
	obj *patternFlowIpv4DontFragmentMetricTag
}

type marshalPatternFlowIpv4DontFragmentMetricTag interface {
	// ToProto marshals PatternFlowIpv4DontFragmentMetricTag to protobuf object *otg.PatternFlowIpv4DontFragmentMetricTag
	ToProto() (*otg.PatternFlowIpv4DontFragmentMetricTag, error)
	// ToPbText marshals PatternFlowIpv4DontFragmentMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DontFragmentMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DontFragmentMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4DontFragmentMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4DontFragmentMetricTag struct {
	obj *patternFlowIpv4DontFragmentMetricTag
}

type unMarshalPatternFlowIpv4DontFragmentMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4DontFragmentMetricTag from protobuf object *otg.PatternFlowIpv4DontFragmentMetricTag
	FromProto(msg *otg.PatternFlowIpv4DontFragmentMetricTag) (PatternFlowIpv4DontFragmentMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4DontFragmentMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DontFragmentMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DontFragmentMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DontFragmentMetricTag) Marshal() marshalPatternFlowIpv4DontFragmentMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DontFragmentMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DontFragmentMetricTag) Unmarshal() unMarshalPatternFlowIpv4DontFragmentMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DontFragmentMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DontFragmentMetricTag) ToProto() (*otg.PatternFlowIpv4DontFragmentMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DontFragmentMetricTag) FromProto(msg *otg.PatternFlowIpv4DontFragmentMetricTag) (PatternFlowIpv4DontFragmentMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DontFragmentMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragmentMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DontFragmentMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragmentMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DontFragmentMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4DontFragmentMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragmentMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4DontFragmentMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DontFragmentMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DontFragmentMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DontFragmentMetricTag) Clone() (PatternFlowIpv4DontFragmentMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DontFragmentMetricTag()
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

// PatternFlowIpv4DontFragmentMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4DontFragmentMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4DontFragmentMetricTag to protobuf object *otg.PatternFlowIpv4DontFragmentMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DontFragmentMetricTag
	// setMsg unmarshals PatternFlowIpv4DontFragmentMetricTag from protobuf object *otg.PatternFlowIpv4DontFragmentMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DontFragmentMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DontFragmentMetricTag
	// validate validates PatternFlowIpv4DontFragmentMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DontFragmentMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4DontFragmentMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4DontFragmentMetricTag
	SetName(value string) PatternFlowIpv4DontFragmentMetricTag
	// Offset returns uint32, set in PatternFlowIpv4DontFragmentMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4DontFragmentMetricTag
	SetOffset(value uint32) PatternFlowIpv4DontFragmentMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4DontFragmentMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4DontFragmentMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4DontFragmentMetricTag
	SetLength(value uint32) PatternFlowIpv4DontFragmentMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4DontFragmentMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4DontFragmentMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4DontFragmentMetricTag object
func (obj *patternFlowIpv4DontFragmentMetricTag) SetName(value string) PatternFlowIpv4DontFragmentMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4DontFragmentMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4DontFragmentMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4DontFragmentMetricTag object
func (obj *patternFlowIpv4DontFragmentMetricTag) SetOffset(value uint32) PatternFlowIpv4DontFragmentMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4DontFragmentMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4DontFragmentMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4DontFragmentMetricTag object
func (obj *patternFlowIpv4DontFragmentMetricTag) SetLength(value uint32) PatternFlowIpv4DontFragmentMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4DontFragmentMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4DontFragmentMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DontFragmentMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4DontFragmentMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4DontFragmentMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
