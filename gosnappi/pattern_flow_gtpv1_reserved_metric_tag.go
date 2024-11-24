package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1ReservedMetricTag *****
type patternFlowGtpv1ReservedMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1ReservedMetricTag
	marshaller   marshalPatternFlowGtpv1ReservedMetricTag
	unMarshaller unMarshalPatternFlowGtpv1ReservedMetricTag
}

func NewPatternFlowGtpv1ReservedMetricTag() PatternFlowGtpv1ReservedMetricTag {
	obj := patternFlowGtpv1ReservedMetricTag{obj: &otg.PatternFlowGtpv1ReservedMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1ReservedMetricTag) msg() *otg.PatternFlowGtpv1ReservedMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1ReservedMetricTag) setMsg(msg *otg.PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1ReservedMetricTag struct {
	obj *patternFlowGtpv1ReservedMetricTag
}

type marshalPatternFlowGtpv1ReservedMetricTag interface {
	// ToProto marshals PatternFlowGtpv1ReservedMetricTag to protobuf object *otg.PatternFlowGtpv1ReservedMetricTag
	ToProto() (*otg.PatternFlowGtpv1ReservedMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1ReservedMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1ReservedMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1ReservedMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1ReservedMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1ReservedMetricTag struct {
	obj *patternFlowGtpv1ReservedMetricTag
}

type unMarshalPatternFlowGtpv1ReservedMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1ReservedMetricTag from protobuf object *otg.PatternFlowGtpv1ReservedMetricTag
	FromProto(msg *otg.PatternFlowGtpv1ReservedMetricTag) (PatternFlowGtpv1ReservedMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1ReservedMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1ReservedMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1ReservedMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1ReservedMetricTag) Marshal() marshalPatternFlowGtpv1ReservedMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1ReservedMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1ReservedMetricTag) Unmarshal() unMarshalPatternFlowGtpv1ReservedMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1ReservedMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1ReservedMetricTag) ToProto() (*otg.PatternFlowGtpv1ReservedMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1ReservedMetricTag) FromProto(msg *otg.PatternFlowGtpv1ReservedMetricTag) (PatternFlowGtpv1ReservedMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1ReservedMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ReservedMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1ReservedMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ReservedMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1ReservedMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1ReservedMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ReservedMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1ReservedMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ReservedMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ReservedMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1ReservedMetricTag) Clone() (PatternFlowGtpv1ReservedMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1ReservedMetricTag()
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

// PatternFlowGtpv1ReservedMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1ReservedMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1ReservedMetricTag to protobuf object *otg.PatternFlowGtpv1ReservedMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1ReservedMetricTag
	// setMsg unmarshals PatternFlowGtpv1ReservedMetricTag from protobuf object *otg.PatternFlowGtpv1ReservedMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1ReservedMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1ReservedMetricTag
	// validate validates PatternFlowGtpv1ReservedMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1ReservedMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1ReservedMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1ReservedMetricTag
	SetName(value string) PatternFlowGtpv1ReservedMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1ReservedMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1ReservedMetricTag
	SetOffset(value uint32) PatternFlowGtpv1ReservedMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1ReservedMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1ReservedMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1ReservedMetricTag
	SetLength(value uint32) PatternFlowGtpv1ReservedMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1ReservedMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1ReservedMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1ReservedMetricTag object
func (obj *patternFlowGtpv1ReservedMetricTag) SetName(value string) PatternFlowGtpv1ReservedMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1ReservedMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1ReservedMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1ReservedMetricTag object
func (obj *patternFlowGtpv1ReservedMetricTag) SetOffset(value uint32) PatternFlowGtpv1ReservedMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1ReservedMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1ReservedMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1ReservedMetricTag object
func (obj *patternFlowGtpv1ReservedMetricTag) SetLength(value uint32) PatternFlowGtpv1ReservedMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1ReservedMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1ReservedMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ReservedMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1ReservedMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1ReservedMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
