package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosDelayMetricTag *****
type patternFlowIpv4TosDelayMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4TosDelayMetricTag
	marshaller   marshalPatternFlowIpv4TosDelayMetricTag
	unMarshaller unMarshalPatternFlowIpv4TosDelayMetricTag
}

func NewPatternFlowIpv4TosDelayMetricTag() PatternFlowIpv4TosDelayMetricTag {
	obj := patternFlowIpv4TosDelayMetricTag{obj: &otg.PatternFlowIpv4TosDelayMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosDelayMetricTag) msg() *otg.PatternFlowIpv4TosDelayMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4TosDelayMetricTag) setMsg(msg *otg.PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosDelayMetricTag struct {
	obj *patternFlowIpv4TosDelayMetricTag
}

type marshalPatternFlowIpv4TosDelayMetricTag interface {
	// ToProto marshals PatternFlowIpv4TosDelayMetricTag to protobuf object *otg.PatternFlowIpv4TosDelayMetricTag
	ToProto() (*otg.PatternFlowIpv4TosDelayMetricTag, error)
	// ToPbText marshals PatternFlowIpv4TosDelayMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosDelayMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosDelayMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosDelayMetricTag struct {
	obj *patternFlowIpv4TosDelayMetricTag
}

type unMarshalPatternFlowIpv4TosDelayMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4TosDelayMetricTag from protobuf object *otg.PatternFlowIpv4TosDelayMetricTag
	FromProto(msg *otg.PatternFlowIpv4TosDelayMetricTag) (PatternFlowIpv4TosDelayMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4TosDelayMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosDelayMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosDelayMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosDelayMetricTag) Marshal() marshalPatternFlowIpv4TosDelayMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosDelayMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosDelayMetricTag) Unmarshal() unMarshalPatternFlowIpv4TosDelayMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosDelayMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosDelayMetricTag) ToProto() (*otg.PatternFlowIpv4TosDelayMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosDelayMetricTag) FromProto(msg *otg.PatternFlowIpv4TosDelayMetricTag) (PatternFlowIpv4TosDelayMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosDelayMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelayMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosDelayMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelayMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosDelayMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelayMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosDelayMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosDelayMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosDelayMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosDelayMetricTag) Clone() (PatternFlowIpv4TosDelayMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosDelayMetricTag()
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

// PatternFlowIpv4TosDelayMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4TosDelayMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4TosDelayMetricTag to protobuf object *otg.PatternFlowIpv4TosDelayMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosDelayMetricTag
	// setMsg unmarshals PatternFlowIpv4TosDelayMetricTag from protobuf object *otg.PatternFlowIpv4TosDelayMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosDelayMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosDelayMetricTag
	// validate validates PatternFlowIpv4TosDelayMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosDelayMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4TosDelayMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4TosDelayMetricTag
	SetName(value string) PatternFlowIpv4TosDelayMetricTag
	// Offset returns uint32, set in PatternFlowIpv4TosDelayMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4TosDelayMetricTag
	SetOffset(value uint32) PatternFlowIpv4TosDelayMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4TosDelayMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4TosDelayMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4TosDelayMetricTag
	SetLength(value uint32) PatternFlowIpv4TosDelayMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4TosDelayMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4TosDelayMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4TosDelayMetricTag object
func (obj *patternFlowIpv4TosDelayMetricTag) SetName(value string) PatternFlowIpv4TosDelayMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosDelayMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosDelayMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4TosDelayMetricTag object
func (obj *patternFlowIpv4TosDelayMetricTag) SetOffset(value uint32) PatternFlowIpv4TosDelayMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosDelayMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosDelayMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4TosDelayMetricTag object
func (obj *patternFlowIpv4TosDelayMetricTag) SetLength(value uint32) PatternFlowIpv4TosDelayMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4TosDelayMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4TosDelayMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosDelayMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4TosDelayMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4TosDelayMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
