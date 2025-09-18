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

// ***** PatternFlowGtpv1SquenceNumberMetricTag *****
type patternFlowGtpv1SquenceNumberMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1SquenceNumberMetricTag
	marshaller   marshalPatternFlowGtpv1SquenceNumberMetricTag
	unMarshaller unMarshalPatternFlowGtpv1SquenceNumberMetricTag
}

func NewPatternFlowGtpv1SquenceNumberMetricTag() PatternFlowGtpv1SquenceNumberMetricTag {
	obj := patternFlowGtpv1SquenceNumberMetricTag{obj: &otg.PatternFlowGtpv1SquenceNumberMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) msg() *otg.PatternFlowGtpv1SquenceNumberMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) setMsg(msg *otg.PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1SquenceNumberMetricTag struct {
	obj *patternFlowGtpv1SquenceNumberMetricTag
}

type marshalPatternFlowGtpv1SquenceNumberMetricTag interface {
	// ToProto marshals PatternFlowGtpv1SquenceNumberMetricTag to protobuf object *otg.PatternFlowGtpv1SquenceNumberMetricTag
	ToProto() (*otg.PatternFlowGtpv1SquenceNumberMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1SquenceNumberMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1SquenceNumberMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1SquenceNumberMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1SquenceNumberMetricTag struct {
	obj *patternFlowGtpv1SquenceNumberMetricTag
}

type unMarshalPatternFlowGtpv1SquenceNumberMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1SquenceNumberMetricTag from protobuf object *otg.PatternFlowGtpv1SquenceNumberMetricTag
	FromProto(msg *otg.PatternFlowGtpv1SquenceNumberMetricTag) (PatternFlowGtpv1SquenceNumberMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1SquenceNumberMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1SquenceNumberMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1SquenceNumberMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) Marshal() marshalPatternFlowGtpv1SquenceNumberMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1SquenceNumberMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) Unmarshal() unMarshalPatternFlowGtpv1SquenceNumberMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1SquenceNumberMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1SquenceNumberMetricTag) ToProto() (*otg.PatternFlowGtpv1SquenceNumberMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1SquenceNumberMetricTag) FromProto(msg *otg.PatternFlowGtpv1SquenceNumberMetricTag) (PatternFlowGtpv1SquenceNumberMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1SquenceNumberMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumberMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1SquenceNumberMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumberMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1SquenceNumberMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumberMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1SquenceNumberMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) Clone() (PatternFlowGtpv1SquenceNumberMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1SquenceNumberMetricTag()
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

// PatternFlowGtpv1SquenceNumberMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1SquenceNumberMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1SquenceNumberMetricTag to protobuf object *otg.PatternFlowGtpv1SquenceNumberMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1SquenceNumberMetricTag
	// setMsg unmarshals PatternFlowGtpv1SquenceNumberMetricTag from protobuf object *otg.PatternFlowGtpv1SquenceNumberMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1SquenceNumberMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1SquenceNumberMetricTag
	// validate validates PatternFlowGtpv1SquenceNumberMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1SquenceNumberMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1SquenceNumberMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1SquenceNumberMetricTag
	SetName(value string) PatternFlowGtpv1SquenceNumberMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1SquenceNumberMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1SquenceNumberMetricTag
	SetOffset(value uint32) PatternFlowGtpv1SquenceNumberMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1SquenceNumberMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1SquenceNumberMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1SquenceNumberMetricTag
	SetLength(value uint32) PatternFlowGtpv1SquenceNumberMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1SquenceNumberMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1SquenceNumberMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1SquenceNumberMetricTag object
func (obj *patternFlowGtpv1SquenceNumberMetricTag) SetName(value string) PatternFlowGtpv1SquenceNumberMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1SquenceNumberMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1SquenceNumberMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1SquenceNumberMetricTag object
func (obj *patternFlowGtpv1SquenceNumberMetricTag) SetOffset(value uint32) PatternFlowGtpv1SquenceNumberMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1SquenceNumberMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1SquenceNumberMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1SquenceNumberMetricTag object
func (obj *patternFlowGtpv1SquenceNumberMetricTag) SetLength(value uint32) PatternFlowGtpv1SquenceNumberMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1SquenceNumberMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGtpv1SquenceNumberMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumberMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1SquenceNumberMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1SquenceNumberMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
