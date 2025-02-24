package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1SFlagMetricTag *****
type patternFlowGtpv1SFlagMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1SFlagMetricTag
	marshaller   marshalPatternFlowGtpv1SFlagMetricTag
	unMarshaller unMarshalPatternFlowGtpv1SFlagMetricTag
}

func NewPatternFlowGtpv1SFlagMetricTag() PatternFlowGtpv1SFlagMetricTag {
	obj := patternFlowGtpv1SFlagMetricTag{obj: &otg.PatternFlowGtpv1SFlagMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1SFlagMetricTag) msg() *otg.PatternFlowGtpv1SFlagMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1SFlagMetricTag) setMsg(msg *otg.PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1SFlagMetricTag struct {
	obj *patternFlowGtpv1SFlagMetricTag
}

type marshalPatternFlowGtpv1SFlagMetricTag interface {
	// ToProto marshals PatternFlowGtpv1SFlagMetricTag to protobuf object *otg.PatternFlowGtpv1SFlagMetricTag
	ToProto() (*otg.PatternFlowGtpv1SFlagMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1SFlagMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1SFlagMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1SFlagMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1SFlagMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1SFlagMetricTag struct {
	obj *patternFlowGtpv1SFlagMetricTag
}

type unMarshalPatternFlowGtpv1SFlagMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1SFlagMetricTag from protobuf object *otg.PatternFlowGtpv1SFlagMetricTag
	FromProto(msg *otg.PatternFlowGtpv1SFlagMetricTag) (PatternFlowGtpv1SFlagMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1SFlagMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1SFlagMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1SFlagMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1SFlagMetricTag) Marshal() marshalPatternFlowGtpv1SFlagMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1SFlagMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1SFlagMetricTag) Unmarshal() unMarshalPatternFlowGtpv1SFlagMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1SFlagMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1SFlagMetricTag) ToProto() (*otg.PatternFlowGtpv1SFlagMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1SFlagMetricTag) FromProto(msg *otg.PatternFlowGtpv1SFlagMetricTag) (PatternFlowGtpv1SFlagMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1SFlagMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlagMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1SFlagMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlagMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1SFlagMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1SFlagMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlagMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1SFlagMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SFlagMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SFlagMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1SFlagMetricTag) Clone() (PatternFlowGtpv1SFlagMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1SFlagMetricTag()
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

// PatternFlowGtpv1SFlagMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1SFlagMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1SFlagMetricTag to protobuf object *otg.PatternFlowGtpv1SFlagMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1SFlagMetricTag
	// setMsg unmarshals PatternFlowGtpv1SFlagMetricTag from protobuf object *otg.PatternFlowGtpv1SFlagMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1SFlagMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1SFlagMetricTag
	// validate validates PatternFlowGtpv1SFlagMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1SFlagMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1SFlagMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1SFlagMetricTag
	SetName(value string) PatternFlowGtpv1SFlagMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1SFlagMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1SFlagMetricTag
	SetOffset(value uint32) PatternFlowGtpv1SFlagMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1SFlagMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1SFlagMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1SFlagMetricTag
	SetLength(value uint32) PatternFlowGtpv1SFlagMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1SFlagMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1SFlagMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1SFlagMetricTag object
func (obj *patternFlowGtpv1SFlagMetricTag) SetName(value string) PatternFlowGtpv1SFlagMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1SFlagMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1SFlagMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1SFlagMetricTag object
func (obj *patternFlowGtpv1SFlagMetricTag) SetOffset(value uint32) PatternFlowGtpv1SFlagMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1SFlagMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1SFlagMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1SFlagMetricTag object
func (obj *patternFlowGtpv1SFlagMetricTag) SetLength(value uint32) PatternFlowGtpv1SFlagMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1SFlagMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1SFlagMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SFlagMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1SFlagMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1SFlagMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
