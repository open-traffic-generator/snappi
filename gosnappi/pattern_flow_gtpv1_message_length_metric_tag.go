package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1MessageLengthMetricTag *****
type patternFlowGtpv1MessageLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1MessageLengthMetricTag
	marshaller   marshalPatternFlowGtpv1MessageLengthMetricTag
	unMarshaller unMarshalPatternFlowGtpv1MessageLengthMetricTag
}

func NewPatternFlowGtpv1MessageLengthMetricTag() PatternFlowGtpv1MessageLengthMetricTag {
	obj := patternFlowGtpv1MessageLengthMetricTag{obj: &otg.PatternFlowGtpv1MessageLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1MessageLengthMetricTag) msg() *otg.PatternFlowGtpv1MessageLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageLengthMetricTag) setMsg(msg *otg.PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1MessageLengthMetricTag struct {
	obj *patternFlowGtpv1MessageLengthMetricTag
}

type marshalPatternFlowGtpv1MessageLengthMetricTag interface {
	// ToProto marshals PatternFlowGtpv1MessageLengthMetricTag to protobuf object *otg.PatternFlowGtpv1MessageLengthMetricTag
	ToProto() (*otg.PatternFlowGtpv1MessageLengthMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1MessageLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1MessageLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1MessageLengthMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1MessageLengthMetricTag struct {
	obj *patternFlowGtpv1MessageLengthMetricTag
}

type unMarshalPatternFlowGtpv1MessageLengthMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1MessageLengthMetricTag from protobuf object *otg.PatternFlowGtpv1MessageLengthMetricTag
	FromProto(msg *otg.PatternFlowGtpv1MessageLengthMetricTag) (PatternFlowGtpv1MessageLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1MessageLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1MessageLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1MessageLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1MessageLengthMetricTag) Marshal() marshalPatternFlowGtpv1MessageLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1MessageLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1MessageLengthMetricTag) Unmarshal() unMarshalPatternFlowGtpv1MessageLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1MessageLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1MessageLengthMetricTag) ToProto() (*otg.PatternFlowGtpv1MessageLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1MessageLengthMetricTag) FromProto(msg *otg.PatternFlowGtpv1MessageLengthMetricTag) (PatternFlowGtpv1MessageLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1MessageLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1MessageLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1MessageLengthMetricTag) Clone() (PatternFlowGtpv1MessageLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1MessageLengthMetricTag()
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

// PatternFlowGtpv1MessageLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1MessageLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1MessageLengthMetricTag to protobuf object *otg.PatternFlowGtpv1MessageLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1MessageLengthMetricTag
	// setMsg unmarshals PatternFlowGtpv1MessageLengthMetricTag from protobuf object *otg.PatternFlowGtpv1MessageLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1MessageLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1MessageLengthMetricTag
	// validate validates PatternFlowGtpv1MessageLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1MessageLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1MessageLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1MessageLengthMetricTag
	SetName(value string) PatternFlowGtpv1MessageLengthMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1MessageLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1MessageLengthMetricTag
	SetOffset(value uint32) PatternFlowGtpv1MessageLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1MessageLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1MessageLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1MessageLengthMetricTag
	SetLength(value uint32) PatternFlowGtpv1MessageLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1MessageLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1MessageLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1MessageLengthMetricTag object
func (obj *patternFlowGtpv1MessageLengthMetricTag) SetName(value string) PatternFlowGtpv1MessageLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1MessageLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1MessageLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1MessageLengthMetricTag object
func (obj *patternFlowGtpv1MessageLengthMetricTag) SetOffset(value uint32) PatternFlowGtpv1MessageLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1MessageLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1MessageLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1MessageLengthMetricTag object
func (obj *patternFlowGtpv1MessageLengthMetricTag) SetLength(value uint32) PatternFlowGtpv1MessageLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1MessageLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1MessageLengthMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGtpv1MessageLengthMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageLengthMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1MessageLengthMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1MessageLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
