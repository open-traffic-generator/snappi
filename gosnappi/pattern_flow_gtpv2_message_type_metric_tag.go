package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2MessageTypeMetricTag *****
type patternFlowGtpv2MessageTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv2MessageTypeMetricTag
	marshaller   marshalPatternFlowGtpv2MessageTypeMetricTag
	unMarshaller unMarshalPatternFlowGtpv2MessageTypeMetricTag
}

func NewPatternFlowGtpv2MessageTypeMetricTag() PatternFlowGtpv2MessageTypeMetricTag {
	obj := patternFlowGtpv2MessageTypeMetricTag{obj: &otg.PatternFlowGtpv2MessageTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2MessageTypeMetricTag) msg() *otg.PatternFlowGtpv2MessageTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageTypeMetricTag) setMsg(msg *otg.PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2MessageTypeMetricTag struct {
	obj *patternFlowGtpv2MessageTypeMetricTag
}

type marshalPatternFlowGtpv2MessageTypeMetricTag interface {
	// ToProto marshals PatternFlowGtpv2MessageTypeMetricTag to protobuf object *otg.PatternFlowGtpv2MessageTypeMetricTag
	ToProto() (*otg.PatternFlowGtpv2MessageTypeMetricTag, error)
	// ToPbText marshals PatternFlowGtpv2MessageTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2MessageTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2MessageTypeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2MessageTypeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2MessageTypeMetricTag struct {
	obj *patternFlowGtpv2MessageTypeMetricTag
}

type unMarshalPatternFlowGtpv2MessageTypeMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv2MessageTypeMetricTag from protobuf object *otg.PatternFlowGtpv2MessageTypeMetricTag
	FromProto(msg *otg.PatternFlowGtpv2MessageTypeMetricTag) (PatternFlowGtpv2MessageTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv2MessageTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2MessageTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2MessageTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2MessageTypeMetricTag) Marshal() marshalPatternFlowGtpv2MessageTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2MessageTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2MessageTypeMetricTag) Unmarshal() unMarshalPatternFlowGtpv2MessageTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2MessageTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2MessageTypeMetricTag) ToProto() (*otg.PatternFlowGtpv2MessageTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2MessageTypeMetricTag) FromProto(msg *otg.PatternFlowGtpv2MessageTypeMetricTag) (PatternFlowGtpv2MessageTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2MessageTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageTypeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2MessageTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2MessageTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2MessageTypeMetricTag) Clone() (PatternFlowGtpv2MessageTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2MessageTypeMetricTag()
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

// PatternFlowGtpv2MessageTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv2MessageTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv2MessageTypeMetricTag to protobuf object *otg.PatternFlowGtpv2MessageTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2MessageTypeMetricTag
	// setMsg unmarshals PatternFlowGtpv2MessageTypeMetricTag from protobuf object *otg.PatternFlowGtpv2MessageTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2MessageTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2MessageTypeMetricTag
	// validate validates PatternFlowGtpv2MessageTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2MessageTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv2MessageTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv2MessageTypeMetricTag
	SetName(value string) PatternFlowGtpv2MessageTypeMetricTag
	// Offset returns uint32, set in PatternFlowGtpv2MessageTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv2MessageTypeMetricTag
	SetOffset(value uint32) PatternFlowGtpv2MessageTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv2MessageTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv2MessageTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv2MessageTypeMetricTag
	SetLength(value uint32) PatternFlowGtpv2MessageTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv2MessageTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv2MessageTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv2MessageTypeMetricTag object
func (obj *patternFlowGtpv2MessageTypeMetricTag) SetName(value string) PatternFlowGtpv2MessageTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2MessageTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2MessageTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv2MessageTypeMetricTag object
func (obj *patternFlowGtpv2MessageTypeMetricTag) SetOffset(value uint32) PatternFlowGtpv2MessageTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2MessageTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2MessageTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv2MessageTypeMetricTag object
func (obj *patternFlowGtpv2MessageTypeMetricTag) SetLength(value uint32) PatternFlowGtpv2MessageTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv2MessageTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv2MessageTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageTypeMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv2MessageTypeMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv2MessageTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
