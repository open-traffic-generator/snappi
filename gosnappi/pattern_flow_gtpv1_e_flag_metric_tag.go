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

// ***** PatternFlowGtpv1EFlagMetricTag *****
type patternFlowGtpv1EFlagMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1EFlagMetricTag
	marshaller   marshalPatternFlowGtpv1EFlagMetricTag
	unMarshaller unMarshalPatternFlowGtpv1EFlagMetricTag
}

func NewPatternFlowGtpv1EFlagMetricTag() PatternFlowGtpv1EFlagMetricTag {
	obj := patternFlowGtpv1EFlagMetricTag{obj: &otg.PatternFlowGtpv1EFlagMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1EFlagMetricTag) msg() *otg.PatternFlowGtpv1EFlagMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1EFlagMetricTag) setMsg(msg *otg.PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1EFlagMetricTag struct {
	obj *patternFlowGtpv1EFlagMetricTag
}

type marshalPatternFlowGtpv1EFlagMetricTag interface {
	// ToProto marshals PatternFlowGtpv1EFlagMetricTag to protobuf object *otg.PatternFlowGtpv1EFlagMetricTag
	ToProto() (*otg.PatternFlowGtpv1EFlagMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1EFlagMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1EFlagMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1EFlagMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1EFlagMetricTag struct {
	obj *patternFlowGtpv1EFlagMetricTag
}

type unMarshalPatternFlowGtpv1EFlagMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1EFlagMetricTag from protobuf object *otg.PatternFlowGtpv1EFlagMetricTag
	FromProto(msg *otg.PatternFlowGtpv1EFlagMetricTag) (PatternFlowGtpv1EFlagMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1EFlagMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1EFlagMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1EFlagMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1EFlagMetricTag) Marshal() marshalPatternFlowGtpv1EFlagMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1EFlagMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1EFlagMetricTag) Unmarshal() unMarshalPatternFlowGtpv1EFlagMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1EFlagMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1EFlagMetricTag) ToProto() (*otg.PatternFlowGtpv1EFlagMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1EFlagMetricTag) FromProto(msg *otg.PatternFlowGtpv1EFlagMetricTag) (PatternFlowGtpv1EFlagMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1EFlagMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlagMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1EFlagMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlagMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1EFlagMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlagMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1EFlagMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1EFlagMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1EFlagMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1EFlagMetricTag) Clone() (PatternFlowGtpv1EFlagMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1EFlagMetricTag()
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

// PatternFlowGtpv1EFlagMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1EFlagMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1EFlagMetricTag to protobuf object *otg.PatternFlowGtpv1EFlagMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1EFlagMetricTag
	// setMsg unmarshals PatternFlowGtpv1EFlagMetricTag from protobuf object *otg.PatternFlowGtpv1EFlagMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1EFlagMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1EFlagMetricTag
	// validate validates PatternFlowGtpv1EFlagMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1EFlagMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1EFlagMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1EFlagMetricTag
	SetName(value string) PatternFlowGtpv1EFlagMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1EFlagMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1EFlagMetricTag
	SetOffset(value uint32) PatternFlowGtpv1EFlagMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1EFlagMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1EFlagMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1EFlagMetricTag
	SetLength(value uint32) PatternFlowGtpv1EFlagMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1EFlagMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1EFlagMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1EFlagMetricTag object
func (obj *patternFlowGtpv1EFlagMetricTag) SetName(value string) PatternFlowGtpv1EFlagMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1EFlagMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1EFlagMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1EFlagMetricTag object
func (obj *patternFlowGtpv1EFlagMetricTag) SetOffset(value uint32) PatternFlowGtpv1EFlagMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1EFlagMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1EFlagMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1EFlagMetricTag object
func (obj *patternFlowGtpv1EFlagMetricTag) SetLength(value uint32) PatternFlowGtpv1EFlagMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1EFlagMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1EFlagMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGtpv1EFlagMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1EFlagMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1EFlagMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1EFlagMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
