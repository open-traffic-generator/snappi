package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TotalLengthMetricTag *****
type patternFlowIpv4TotalLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4TotalLengthMetricTag
	marshaller   marshalPatternFlowIpv4TotalLengthMetricTag
	unMarshaller unMarshalPatternFlowIpv4TotalLengthMetricTag
}

func NewPatternFlowIpv4TotalLengthMetricTag() PatternFlowIpv4TotalLengthMetricTag {
	obj := patternFlowIpv4TotalLengthMetricTag{obj: &otg.PatternFlowIpv4TotalLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TotalLengthMetricTag) msg() *otg.PatternFlowIpv4TotalLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4TotalLengthMetricTag) setMsg(msg *otg.PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TotalLengthMetricTag struct {
	obj *patternFlowIpv4TotalLengthMetricTag
}

type marshalPatternFlowIpv4TotalLengthMetricTag interface {
	// ToProto marshals PatternFlowIpv4TotalLengthMetricTag to protobuf object *otg.PatternFlowIpv4TotalLengthMetricTag
	ToProto() (*otg.PatternFlowIpv4TotalLengthMetricTag, error)
	// ToPbText marshals PatternFlowIpv4TotalLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TotalLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TotalLengthMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TotalLengthMetricTag struct {
	obj *patternFlowIpv4TotalLengthMetricTag
}

type unMarshalPatternFlowIpv4TotalLengthMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4TotalLengthMetricTag from protobuf object *otg.PatternFlowIpv4TotalLengthMetricTag
	FromProto(msg *otg.PatternFlowIpv4TotalLengthMetricTag) (PatternFlowIpv4TotalLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4TotalLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TotalLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TotalLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TotalLengthMetricTag) Marshal() marshalPatternFlowIpv4TotalLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TotalLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TotalLengthMetricTag) Unmarshal() unMarshalPatternFlowIpv4TotalLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TotalLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TotalLengthMetricTag) ToProto() (*otg.PatternFlowIpv4TotalLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TotalLengthMetricTag) FromProto(msg *otg.PatternFlowIpv4TotalLengthMetricTag) (PatternFlowIpv4TotalLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TotalLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TotalLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TotalLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4TotalLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TotalLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TotalLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TotalLengthMetricTag) Clone() (PatternFlowIpv4TotalLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TotalLengthMetricTag()
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

// PatternFlowIpv4TotalLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4TotalLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4TotalLengthMetricTag to protobuf object *otg.PatternFlowIpv4TotalLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TotalLengthMetricTag
	// setMsg unmarshals PatternFlowIpv4TotalLengthMetricTag from protobuf object *otg.PatternFlowIpv4TotalLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TotalLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TotalLengthMetricTag
	// validate validates PatternFlowIpv4TotalLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TotalLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4TotalLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4TotalLengthMetricTag
	SetName(value string) PatternFlowIpv4TotalLengthMetricTag
	// Offset returns uint32, set in PatternFlowIpv4TotalLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4TotalLengthMetricTag
	SetOffset(value uint32) PatternFlowIpv4TotalLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4TotalLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4TotalLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4TotalLengthMetricTag
	SetLength(value uint32) PatternFlowIpv4TotalLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4TotalLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4TotalLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4TotalLengthMetricTag object
func (obj *patternFlowIpv4TotalLengthMetricTag) SetName(value string) PatternFlowIpv4TotalLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TotalLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TotalLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4TotalLengthMetricTag object
func (obj *patternFlowIpv4TotalLengthMetricTag) SetOffset(value uint32) PatternFlowIpv4TotalLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TotalLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TotalLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4TotalLengthMetricTag object
func (obj *patternFlowIpv4TotalLengthMetricTag) SetLength(value uint32) PatternFlowIpv4TotalLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4TotalLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4TotalLengthMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLengthMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4TotalLengthMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4TotalLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
