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

// ***** PatternFlowGtpv2SequenceNumberMetricTag *****
type patternFlowGtpv2SequenceNumberMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv2SequenceNumberMetricTag
	marshaller   marshalPatternFlowGtpv2SequenceNumberMetricTag
	unMarshaller unMarshalPatternFlowGtpv2SequenceNumberMetricTag
}

func NewPatternFlowGtpv2SequenceNumberMetricTag() PatternFlowGtpv2SequenceNumberMetricTag {
	obj := patternFlowGtpv2SequenceNumberMetricTag{obj: &otg.PatternFlowGtpv2SequenceNumberMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) msg() *otg.PatternFlowGtpv2SequenceNumberMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) setMsg(msg *otg.PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2SequenceNumberMetricTag struct {
	obj *patternFlowGtpv2SequenceNumberMetricTag
}

type marshalPatternFlowGtpv2SequenceNumberMetricTag interface {
	// ToProto marshals PatternFlowGtpv2SequenceNumberMetricTag to protobuf object *otg.PatternFlowGtpv2SequenceNumberMetricTag
	ToProto() (*otg.PatternFlowGtpv2SequenceNumberMetricTag, error)
	// ToPbText marshals PatternFlowGtpv2SequenceNumberMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2SequenceNumberMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2SequenceNumberMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2SequenceNumberMetricTag struct {
	obj *patternFlowGtpv2SequenceNumberMetricTag
}

type unMarshalPatternFlowGtpv2SequenceNumberMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv2SequenceNumberMetricTag from protobuf object *otg.PatternFlowGtpv2SequenceNumberMetricTag
	FromProto(msg *otg.PatternFlowGtpv2SequenceNumberMetricTag) (PatternFlowGtpv2SequenceNumberMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv2SequenceNumberMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2SequenceNumberMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2SequenceNumberMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) Marshal() marshalPatternFlowGtpv2SequenceNumberMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2SequenceNumberMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) Unmarshal() unMarshalPatternFlowGtpv2SequenceNumberMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2SequenceNumberMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2SequenceNumberMetricTag) ToProto() (*otg.PatternFlowGtpv2SequenceNumberMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2SequenceNumberMetricTag) FromProto(msg *otg.PatternFlowGtpv2SequenceNumberMetricTag) (PatternFlowGtpv2SequenceNumberMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2SequenceNumberMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumberMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2SequenceNumberMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumberMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2SequenceNumberMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumberMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2SequenceNumberMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) Clone() (PatternFlowGtpv2SequenceNumberMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2SequenceNumberMetricTag()
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

// PatternFlowGtpv2SequenceNumberMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv2SequenceNumberMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv2SequenceNumberMetricTag to protobuf object *otg.PatternFlowGtpv2SequenceNumberMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2SequenceNumberMetricTag
	// setMsg unmarshals PatternFlowGtpv2SequenceNumberMetricTag from protobuf object *otg.PatternFlowGtpv2SequenceNumberMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2SequenceNumberMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2SequenceNumberMetricTag
	// validate validates PatternFlowGtpv2SequenceNumberMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2SequenceNumberMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv2SequenceNumberMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv2SequenceNumberMetricTag
	SetName(value string) PatternFlowGtpv2SequenceNumberMetricTag
	// Offset returns uint32, set in PatternFlowGtpv2SequenceNumberMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv2SequenceNumberMetricTag
	SetOffset(value uint32) PatternFlowGtpv2SequenceNumberMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv2SequenceNumberMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv2SequenceNumberMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv2SequenceNumberMetricTag
	SetLength(value uint32) PatternFlowGtpv2SequenceNumberMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv2SequenceNumberMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv2SequenceNumberMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv2SequenceNumberMetricTag object
func (obj *patternFlowGtpv2SequenceNumberMetricTag) SetName(value string) PatternFlowGtpv2SequenceNumberMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2SequenceNumberMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2SequenceNumberMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv2SequenceNumberMetricTag object
func (obj *patternFlowGtpv2SequenceNumberMetricTag) SetOffset(value uint32) PatternFlowGtpv2SequenceNumberMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2SequenceNumberMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2SequenceNumberMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv2SequenceNumberMetricTag object
func (obj *patternFlowGtpv2SequenceNumberMetricTag) SetLength(value uint32) PatternFlowGtpv2SequenceNumberMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv2SequenceNumberMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGtpv2SequenceNumberMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 23 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumberMetricTag.Offset <= 23 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 24 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv2SequenceNumberMetricTag.Length <= 24 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv2SequenceNumberMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(24)
	}

}
