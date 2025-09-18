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

// ***** PatternFlowGtpv2PiggybackingFlagMetricTag *****
type patternFlowGtpv2PiggybackingFlagMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv2PiggybackingFlagMetricTag
	marshaller   marshalPatternFlowGtpv2PiggybackingFlagMetricTag
	unMarshaller unMarshalPatternFlowGtpv2PiggybackingFlagMetricTag
}

func NewPatternFlowGtpv2PiggybackingFlagMetricTag() PatternFlowGtpv2PiggybackingFlagMetricTag {
	obj := patternFlowGtpv2PiggybackingFlagMetricTag{obj: &otg.PatternFlowGtpv2PiggybackingFlagMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) msg() *otg.PatternFlowGtpv2PiggybackingFlagMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) setMsg(msg *otg.PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2PiggybackingFlagMetricTag struct {
	obj *patternFlowGtpv2PiggybackingFlagMetricTag
}

type marshalPatternFlowGtpv2PiggybackingFlagMetricTag interface {
	// ToProto marshals PatternFlowGtpv2PiggybackingFlagMetricTag to protobuf object *otg.PatternFlowGtpv2PiggybackingFlagMetricTag
	ToProto() (*otg.PatternFlowGtpv2PiggybackingFlagMetricTag, error)
	// ToPbText marshals PatternFlowGtpv2PiggybackingFlagMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2PiggybackingFlagMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2PiggybackingFlagMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2PiggybackingFlagMetricTag struct {
	obj *patternFlowGtpv2PiggybackingFlagMetricTag
}

type unMarshalPatternFlowGtpv2PiggybackingFlagMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv2PiggybackingFlagMetricTag from protobuf object *otg.PatternFlowGtpv2PiggybackingFlagMetricTag
	FromProto(msg *otg.PatternFlowGtpv2PiggybackingFlagMetricTag) (PatternFlowGtpv2PiggybackingFlagMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv2PiggybackingFlagMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2PiggybackingFlagMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2PiggybackingFlagMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) Marshal() marshalPatternFlowGtpv2PiggybackingFlagMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2PiggybackingFlagMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) Unmarshal() unMarshalPatternFlowGtpv2PiggybackingFlagMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2PiggybackingFlagMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2PiggybackingFlagMetricTag) ToProto() (*otg.PatternFlowGtpv2PiggybackingFlagMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2PiggybackingFlagMetricTag) FromProto(msg *otg.PatternFlowGtpv2PiggybackingFlagMetricTag) (PatternFlowGtpv2PiggybackingFlagMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2PiggybackingFlagMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlagMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2PiggybackingFlagMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlagMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2PiggybackingFlagMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlagMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) Clone() (PatternFlowGtpv2PiggybackingFlagMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2PiggybackingFlagMetricTag()
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

// PatternFlowGtpv2PiggybackingFlagMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv2PiggybackingFlagMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv2PiggybackingFlagMetricTag to protobuf object *otg.PatternFlowGtpv2PiggybackingFlagMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2PiggybackingFlagMetricTag
	// setMsg unmarshals PatternFlowGtpv2PiggybackingFlagMetricTag from protobuf object *otg.PatternFlowGtpv2PiggybackingFlagMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2PiggybackingFlagMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2PiggybackingFlagMetricTag
	// validate validates PatternFlowGtpv2PiggybackingFlagMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2PiggybackingFlagMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv2PiggybackingFlagMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv2PiggybackingFlagMetricTag
	SetName(value string) PatternFlowGtpv2PiggybackingFlagMetricTag
	// Offset returns uint32, set in PatternFlowGtpv2PiggybackingFlagMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv2PiggybackingFlagMetricTag
	SetOffset(value uint32) PatternFlowGtpv2PiggybackingFlagMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv2PiggybackingFlagMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv2PiggybackingFlagMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv2PiggybackingFlagMetricTag
	SetLength(value uint32) PatternFlowGtpv2PiggybackingFlagMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv2PiggybackingFlagMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv2PiggybackingFlagMetricTag object
func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) SetName(value string) PatternFlowGtpv2PiggybackingFlagMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv2PiggybackingFlagMetricTag object
func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) SetOffset(value uint32) PatternFlowGtpv2PiggybackingFlagMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv2PiggybackingFlagMetricTag object
func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) SetLength(value uint32) PatternFlowGtpv2PiggybackingFlagMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv2PiggybackingFlagMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGtpv2PiggybackingFlagMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlagMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv2PiggybackingFlagMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv2PiggybackingFlagMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
