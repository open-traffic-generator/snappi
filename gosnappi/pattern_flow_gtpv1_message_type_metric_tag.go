package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1MessageTypeMetricTag *****
type patternFlowGtpv1MessageTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1MessageTypeMetricTag
	marshaller   marshalPatternFlowGtpv1MessageTypeMetricTag
	unMarshaller unMarshalPatternFlowGtpv1MessageTypeMetricTag
}

func NewPatternFlowGtpv1MessageTypeMetricTag() PatternFlowGtpv1MessageTypeMetricTag {
	obj := patternFlowGtpv1MessageTypeMetricTag{obj: &otg.PatternFlowGtpv1MessageTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1MessageTypeMetricTag) msg() *otg.PatternFlowGtpv1MessageTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageTypeMetricTag) setMsg(msg *otg.PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1MessageTypeMetricTag struct {
	obj *patternFlowGtpv1MessageTypeMetricTag
}

type marshalPatternFlowGtpv1MessageTypeMetricTag interface {
	// ToProto marshals PatternFlowGtpv1MessageTypeMetricTag to protobuf object *otg.PatternFlowGtpv1MessageTypeMetricTag
	ToProto() (*otg.PatternFlowGtpv1MessageTypeMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1MessageTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1MessageTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1MessageTypeMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1MessageTypeMetricTag struct {
	obj *patternFlowGtpv1MessageTypeMetricTag
}

type unMarshalPatternFlowGtpv1MessageTypeMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1MessageTypeMetricTag from protobuf object *otg.PatternFlowGtpv1MessageTypeMetricTag
	FromProto(msg *otg.PatternFlowGtpv1MessageTypeMetricTag) (PatternFlowGtpv1MessageTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1MessageTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1MessageTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1MessageTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1MessageTypeMetricTag) Marshal() marshalPatternFlowGtpv1MessageTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1MessageTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1MessageTypeMetricTag) Unmarshal() unMarshalPatternFlowGtpv1MessageTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1MessageTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1MessageTypeMetricTag) ToProto() (*otg.PatternFlowGtpv1MessageTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1MessageTypeMetricTag) FromProto(msg *otg.PatternFlowGtpv1MessageTypeMetricTag) (PatternFlowGtpv1MessageTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1MessageTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1MessageTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1MessageTypeMetricTag) Clone() (PatternFlowGtpv1MessageTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1MessageTypeMetricTag()
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

// PatternFlowGtpv1MessageTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1MessageTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1MessageTypeMetricTag to protobuf object *otg.PatternFlowGtpv1MessageTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1MessageTypeMetricTag
	// setMsg unmarshals PatternFlowGtpv1MessageTypeMetricTag from protobuf object *otg.PatternFlowGtpv1MessageTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1MessageTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1MessageTypeMetricTag
	// validate validates PatternFlowGtpv1MessageTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1MessageTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1MessageTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1MessageTypeMetricTag
	SetName(value string) PatternFlowGtpv1MessageTypeMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1MessageTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1MessageTypeMetricTag
	SetOffset(value uint32) PatternFlowGtpv1MessageTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1MessageTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1MessageTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1MessageTypeMetricTag
	SetLength(value uint32) PatternFlowGtpv1MessageTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1MessageTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1MessageTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1MessageTypeMetricTag object
func (obj *patternFlowGtpv1MessageTypeMetricTag) SetName(value string) PatternFlowGtpv1MessageTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1MessageTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1MessageTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1MessageTypeMetricTag object
func (obj *patternFlowGtpv1MessageTypeMetricTag) SetOffset(value uint32) PatternFlowGtpv1MessageTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1MessageTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1MessageTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1MessageTypeMetricTag object
func (obj *patternFlowGtpv1MessageTypeMetricTag) SetLength(value uint32) PatternFlowGtpv1MessageTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1MessageTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1MessageTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageTypeMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1MessageTypeMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1MessageTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
