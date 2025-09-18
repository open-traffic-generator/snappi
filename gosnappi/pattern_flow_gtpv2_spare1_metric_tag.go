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

// ***** PatternFlowGtpv2Spare1MetricTag *****
type patternFlowGtpv2Spare1MetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv2Spare1MetricTag
	marshaller   marshalPatternFlowGtpv2Spare1MetricTag
	unMarshaller unMarshalPatternFlowGtpv2Spare1MetricTag
}

func NewPatternFlowGtpv2Spare1MetricTag() PatternFlowGtpv2Spare1MetricTag {
	obj := patternFlowGtpv2Spare1MetricTag{obj: &otg.PatternFlowGtpv2Spare1MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2Spare1MetricTag) msg() *otg.PatternFlowGtpv2Spare1MetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv2Spare1MetricTag) setMsg(msg *otg.PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2Spare1MetricTag struct {
	obj *patternFlowGtpv2Spare1MetricTag
}

type marshalPatternFlowGtpv2Spare1MetricTag interface {
	// ToProto marshals PatternFlowGtpv2Spare1MetricTag to protobuf object *otg.PatternFlowGtpv2Spare1MetricTag
	ToProto() (*otg.PatternFlowGtpv2Spare1MetricTag, error)
	// ToPbText marshals PatternFlowGtpv2Spare1MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2Spare1MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2Spare1MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2Spare1MetricTag struct {
	obj *patternFlowGtpv2Spare1MetricTag
}

type unMarshalPatternFlowGtpv2Spare1MetricTag interface {
	// FromProto unmarshals PatternFlowGtpv2Spare1MetricTag from protobuf object *otg.PatternFlowGtpv2Spare1MetricTag
	FromProto(msg *otg.PatternFlowGtpv2Spare1MetricTag) (PatternFlowGtpv2Spare1MetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv2Spare1MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2Spare1MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2Spare1MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2Spare1MetricTag) Marshal() marshalPatternFlowGtpv2Spare1MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2Spare1MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2Spare1MetricTag) Unmarshal() unMarshalPatternFlowGtpv2Spare1MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2Spare1MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2Spare1MetricTag) ToProto() (*otg.PatternFlowGtpv2Spare1MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2Spare1MetricTag) FromProto(msg *otg.PatternFlowGtpv2Spare1MetricTag) (PatternFlowGtpv2Spare1MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2Spare1MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2Spare1MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2Spare1MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1MetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2Spare1MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Spare1MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Spare1MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2Spare1MetricTag) Clone() (PatternFlowGtpv2Spare1MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2Spare1MetricTag()
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

// PatternFlowGtpv2Spare1MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv2Spare1MetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv2Spare1MetricTag to protobuf object *otg.PatternFlowGtpv2Spare1MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2Spare1MetricTag
	// setMsg unmarshals PatternFlowGtpv2Spare1MetricTag from protobuf object *otg.PatternFlowGtpv2Spare1MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2Spare1MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2Spare1MetricTag
	// validate validates PatternFlowGtpv2Spare1MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2Spare1MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv2Spare1MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv2Spare1MetricTag
	SetName(value string) PatternFlowGtpv2Spare1MetricTag
	// Offset returns uint32, set in PatternFlowGtpv2Spare1MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv2Spare1MetricTag
	SetOffset(value uint32) PatternFlowGtpv2Spare1MetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv2Spare1MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv2Spare1MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv2Spare1MetricTag
	SetLength(value uint32) PatternFlowGtpv2Spare1MetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv2Spare1MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv2Spare1MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv2Spare1MetricTag object
func (obj *patternFlowGtpv2Spare1MetricTag) SetName(value string) PatternFlowGtpv2Spare1MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2Spare1MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2Spare1MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv2Spare1MetricTag object
func (obj *patternFlowGtpv2Spare1MetricTag) SetOffset(value uint32) PatternFlowGtpv2Spare1MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2Spare1MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2Spare1MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv2Spare1MetricTag object
func (obj *patternFlowGtpv2Spare1MetricTag) SetLength(value uint32) PatternFlowGtpv2Spare1MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv2Spare1MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv2Spare1MetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGtpv2Spare1MetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare1MetricTag.Offset <= 2 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv2Spare1MetricTag.Length <= 3 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv2Spare1MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(3)
	}

}
