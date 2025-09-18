package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1TeidMetricTag *****
type patternFlowGtpv1TeidMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1TeidMetricTag
	marshaller   marshalPatternFlowGtpv1TeidMetricTag
	unMarshaller unMarshalPatternFlowGtpv1TeidMetricTag
}

func NewPatternFlowGtpv1TeidMetricTag() PatternFlowGtpv1TeidMetricTag {
	obj := patternFlowGtpv1TeidMetricTag{obj: &otg.PatternFlowGtpv1TeidMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1TeidMetricTag) msg() *otg.PatternFlowGtpv1TeidMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1TeidMetricTag) setMsg(msg *otg.PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1TeidMetricTag struct {
	obj *patternFlowGtpv1TeidMetricTag
}

type marshalPatternFlowGtpv1TeidMetricTag interface {
	// ToProto marshals PatternFlowGtpv1TeidMetricTag to protobuf object *otg.PatternFlowGtpv1TeidMetricTag
	ToProto() (*otg.PatternFlowGtpv1TeidMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1TeidMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1TeidMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1TeidMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1TeidMetricTag struct {
	obj *patternFlowGtpv1TeidMetricTag
}

type unMarshalPatternFlowGtpv1TeidMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1TeidMetricTag from protobuf object *otg.PatternFlowGtpv1TeidMetricTag
	FromProto(msg *otg.PatternFlowGtpv1TeidMetricTag) (PatternFlowGtpv1TeidMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1TeidMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1TeidMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1TeidMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1TeidMetricTag) Marshal() marshalPatternFlowGtpv1TeidMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1TeidMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1TeidMetricTag) Unmarshal() unMarshalPatternFlowGtpv1TeidMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1TeidMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1TeidMetricTag) ToProto() (*otg.PatternFlowGtpv1TeidMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1TeidMetricTag) FromProto(msg *otg.PatternFlowGtpv1TeidMetricTag) (PatternFlowGtpv1TeidMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1TeidMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1TeidMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1TeidMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1TeidMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1TeidMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1TeidMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1TeidMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1TeidMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1TeidMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1TeidMetricTag) Clone() (PatternFlowGtpv1TeidMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1TeidMetricTag()
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

// PatternFlowGtpv1TeidMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1TeidMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1TeidMetricTag to protobuf object *otg.PatternFlowGtpv1TeidMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1TeidMetricTag
	// setMsg unmarshals PatternFlowGtpv1TeidMetricTag from protobuf object *otg.PatternFlowGtpv1TeidMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1TeidMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1TeidMetricTag
	// validate validates PatternFlowGtpv1TeidMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1TeidMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1TeidMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1TeidMetricTag
	SetName(value string) PatternFlowGtpv1TeidMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1TeidMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1TeidMetricTag
	SetOffset(value uint32) PatternFlowGtpv1TeidMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1TeidMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1TeidMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1TeidMetricTag
	SetLength(value uint32) PatternFlowGtpv1TeidMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1TeidMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1TeidMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1TeidMetricTag object
func (obj *patternFlowGtpv1TeidMetricTag) SetName(value string) PatternFlowGtpv1TeidMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1TeidMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1TeidMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1TeidMetricTag object
func (obj *patternFlowGtpv1TeidMetricTag) SetOffset(value uint32) PatternFlowGtpv1TeidMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1TeidMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1TeidMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1TeidMetricTag object
func (obj *patternFlowGtpv1TeidMetricTag) SetLength(value uint32) PatternFlowGtpv1TeidMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1TeidMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1TeidMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1TeidMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1TeidMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1TeidMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
