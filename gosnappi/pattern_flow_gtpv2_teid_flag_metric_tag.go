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

// ***** PatternFlowGtpv2TeidFlagMetricTag *****
type patternFlowGtpv2TeidFlagMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv2TeidFlagMetricTag
	marshaller   marshalPatternFlowGtpv2TeidFlagMetricTag
	unMarshaller unMarshalPatternFlowGtpv2TeidFlagMetricTag
}

func NewPatternFlowGtpv2TeidFlagMetricTag() PatternFlowGtpv2TeidFlagMetricTag {
	obj := patternFlowGtpv2TeidFlagMetricTag{obj: &otg.PatternFlowGtpv2TeidFlagMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2TeidFlagMetricTag) msg() *otg.PatternFlowGtpv2TeidFlagMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv2TeidFlagMetricTag) setMsg(msg *otg.PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2TeidFlagMetricTag struct {
	obj *patternFlowGtpv2TeidFlagMetricTag
}

type marshalPatternFlowGtpv2TeidFlagMetricTag interface {
	// ToProto marshals PatternFlowGtpv2TeidFlagMetricTag to protobuf object *otg.PatternFlowGtpv2TeidFlagMetricTag
	ToProto() (*otg.PatternFlowGtpv2TeidFlagMetricTag, error)
	// ToPbText marshals PatternFlowGtpv2TeidFlagMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2TeidFlagMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2TeidFlagMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2TeidFlagMetricTag struct {
	obj *patternFlowGtpv2TeidFlagMetricTag
}

type unMarshalPatternFlowGtpv2TeidFlagMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv2TeidFlagMetricTag from protobuf object *otg.PatternFlowGtpv2TeidFlagMetricTag
	FromProto(msg *otg.PatternFlowGtpv2TeidFlagMetricTag) (PatternFlowGtpv2TeidFlagMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv2TeidFlagMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2TeidFlagMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2TeidFlagMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2TeidFlagMetricTag) Marshal() marshalPatternFlowGtpv2TeidFlagMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2TeidFlagMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2TeidFlagMetricTag) Unmarshal() unMarshalPatternFlowGtpv2TeidFlagMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2TeidFlagMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2TeidFlagMetricTag) ToProto() (*otg.PatternFlowGtpv2TeidFlagMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2TeidFlagMetricTag) FromProto(msg *otg.PatternFlowGtpv2TeidFlagMetricTag) (PatternFlowGtpv2TeidFlagMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2TeidFlagMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlagMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidFlagMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlagMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidFlagMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlagMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2TeidFlagMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidFlagMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidFlagMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2TeidFlagMetricTag) Clone() (PatternFlowGtpv2TeidFlagMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2TeidFlagMetricTag()
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

// PatternFlowGtpv2TeidFlagMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv2TeidFlagMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv2TeidFlagMetricTag to protobuf object *otg.PatternFlowGtpv2TeidFlagMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2TeidFlagMetricTag
	// setMsg unmarshals PatternFlowGtpv2TeidFlagMetricTag from protobuf object *otg.PatternFlowGtpv2TeidFlagMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2TeidFlagMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2TeidFlagMetricTag
	// validate validates PatternFlowGtpv2TeidFlagMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2TeidFlagMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv2TeidFlagMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv2TeidFlagMetricTag
	SetName(value string) PatternFlowGtpv2TeidFlagMetricTag
	// Offset returns uint32, set in PatternFlowGtpv2TeidFlagMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv2TeidFlagMetricTag
	SetOffset(value uint32) PatternFlowGtpv2TeidFlagMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv2TeidFlagMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv2TeidFlagMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv2TeidFlagMetricTag
	SetLength(value uint32) PatternFlowGtpv2TeidFlagMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv2TeidFlagMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv2TeidFlagMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv2TeidFlagMetricTag object
func (obj *patternFlowGtpv2TeidFlagMetricTag) SetName(value string) PatternFlowGtpv2TeidFlagMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2TeidFlagMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2TeidFlagMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv2TeidFlagMetricTag object
func (obj *patternFlowGtpv2TeidFlagMetricTag) SetOffset(value uint32) PatternFlowGtpv2TeidFlagMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2TeidFlagMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2TeidFlagMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv2TeidFlagMetricTag object
func (obj *patternFlowGtpv2TeidFlagMetricTag) SetLength(value uint32) PatternFlowGtpv2TeidFlagMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv2TeidFlagMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv2TeidFlagMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGtpv2TeidFlagMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlagMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv2TeidFlagMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv2TeidFlagMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
