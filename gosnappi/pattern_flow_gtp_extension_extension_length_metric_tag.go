package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpExtensionExtensionLengthMetricTag *****
type patternFlowGtpExtensionExtensionLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpExtensionExtensionLengthMetricTag
	marshaller   marshalPatternFlowGtpExtensionExtensionLengthMetricTag
	unMarshaller unMarshalPatternFlowGtpExtensionExtensionLengthMetricTag
}

func NewPatternFlowGtpExtensionExtensionLengthMetricTag() PatternFlowGtpExtensionExtensionLengthMetricTag {
	obj := patternFlowGtpExtensionExtensionLengthMetricTag{obj: &otg.PatternFlowGtpExtensionExtensionLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) msg() *otg.PatternFlowGtpExtensionExtensionLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) setMsg(msg *otg.PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionExtensionLengthMetricTag struct {
	obj *patternFlowGtpExtensionExtensionLengthMetricTag
}

type marshalPatternFlowGtpExtensionExtensionLengthMetricTag interface {
	// ToProto marshals PatternFlowGtpExtensionExtensionLengthMetricTag to protobuf object *otg.PatternFlowGtpExtensionExtensionLengthMetricTag
	ToProto() (*otg.PatternFlowGtpExtensionExtensionLengthMetricTag, error)
	// ToPbText marshals PatternFlowGtpExtensionExtensionLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionExtensionLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionExtensionLengthMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpExtensionExtensionLengthMetricTag struct {
	obj *patternFlowGtpExtensionExtensionLengthMetricTag
}

type unMarshalPatternFlowGtpExtensionExtensionLengthMetricTag interface {
	// FromProto unmarshals PatternFlowGtpExtensionExtensionLengthMetricTag from protobuf object *otg.PatternFlowGtpExtensionExtensionLengthMetricTag
	FromProto(msg *otg.PatternFlowGtpExtensionExtensionLengthMetricTag) (PatternFlowGtpExtensionExtensionLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpExtensionExtensionLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionExtensionLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionExtensionLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) Marshal() marshalPatternFlowGtpExtensionExtensionLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionExtensionLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) Unmarshal() unMarshalPatternFlowGtpExtensionExtensionLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionExtensionLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionExtensionLengthMetricTag) ToProto() (*otg.PatternFlowGtpExtensionExtensionLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionExtensionLengthMetricTag) FromProto(msg *otg.PatternFlowGtpExtensionExtensionLengthMetricTag) (PatternFlowGtpExtensionExtensionLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionExtensionLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionExtensionLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionExtensionLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) Clone() (PatternFlowGtpExtensionExtensionLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionExtensionLengthMetricTag()
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

// PatternFlowGtpExtensionExtensionLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpExtensionExtensionLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpExtensionExtensionLengthMetricTag to protobuf object *otg.PatternFlowGtpExtensionExtensionLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionExtensionLengthMetricTag
	// setMsg unmarshals PatternFlowGtpExtensionExtensionLengthMetricTag from protobuf object *otg.PatternFlowGtpExtensionExtensionLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionExtensionLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionExtensionLengthMetricTag
	// validate validates PatternFlowGtpExtensionExtensionLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionExtensionLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpExtensionExtensionLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpExtensionExtensionLengthMetricTag
	SetName(value string) PatternFlowGtpExtensionExtensionLengthMetricTag
	// Offset returns uint32, set in PatternFlowGtpExtensionExtensionLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpExtensionExtensionLengthMetricTag
	SetOffset(value uint32) PatternFlowGtpExtensionExtensionLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpExtensionExtensionLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpExtensionExtensionLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpExtensionExtensionLengthMetricTag
	SetLength(value uint32) PatternFlowGtpExtensionExtensionLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpExtensionExtensionLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpExtensionExtensionLengthMetricTag object
func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) SetName(value string) PatternFlowGtpExtensionExtensionLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpExtensionExtensionLengthMetricTag object
func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) SetOffset(value uint32) PatternFlowGtpExtensionExtensionLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpExtensionExtensionLengthMetricTag object
func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) SetLength(value uint32) PatternFlowGtpExtensionExtensionLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpExtensionExtensionLengthMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLengthMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpExtensionExtensionLengthMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpExtensionExtensionLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
