package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2TeidMetricTag *****
type patternFlowGtpv2TeidMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv2TeidMetricTag
	marshaller   marshalPatternFlowGtpv2TeidMetricTag
	unMarshaller unMarshalPatternFlowGtpv2TeidMetricTag
}

func NewPatternFlowGtpv2TeidMetricTag() PatternFlowGtpv2TeidMetricTag {
	obj := patternFlowGtpv2TeidMetricTag{obj: &otg.PatternFlowGtpv2TeidMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2TeidMetricTag) msg() *otg.PatternFlowGtpv2TeidMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv2TeidMetricTag) setMsg(msg *otg.PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2TeidMetricTag struct {
	obj *patternFlowGtpv2TeidMetricTag
}

type marshalPatternFlowGtpv2TeidMetricTag interface {
	// ToProto marshals PatternFlowGtpv2TeidMetricTag to protobuf object *otg.PatternFlowGtpv2TeidMetricTag
	ToProto() (*otg.PatternFlowGtpv2TeidMetricTag, error)
	// ToPbText marshals PatternFlowGtpv2TeidMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2TeidMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2TeidMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2TeidMetricTag struct {
	obj *patternFlowGtpv2TeidMetricTag
}

type unMarshalPatternFlowGtpv2TeidMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv2TeidMetricTag from protobuf object *otg.PatternFlowGtpv2TeidMetricTag
	FromProto(msg *otg.PatternFlowGtpv2TeidMetricTag) (PatternFlowGtpv2TeidMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv2TeidMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2TeidMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2TeidMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2TeidMetricTag) Marshal() marshalPatternFlowGtpv2TeidMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2TeidMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2TeidMetricTag) Unmarshal() unMarshalPatternFlowGtpv2TeidMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2TeidMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2TeidMetricTag) ToProto() (*otg.PatternFlowGtpv2TeidMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2TeidMetricTag) FromProto(msg *otg.PatternFlowGtpv2TeidMetricTag) (PatternFlowGtpv2TeidMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2TeidMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2TeidMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2TeidMetricTag) Clone() (PatternFlowGtpv2TeidMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2TeidMetricTag()
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

// PatternFlowGtpv2TeidMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv2TeidMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv2TeidMetricTag to protobuf object *otg.PatternFlowGtpv2TeidMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2TeidMetricTag
	// setMsg unmarshals PatternFlowGtpv2TeidMetricTag from protobuf object *otg.PatternFlowGtpv2TeidMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2TeidMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2TeidMetricTag
	// validate validates PatternFlowGtpv2TeidMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2TeidMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv2TeidMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv2TeidMetricTag
	SetName(value string) PatternFlowGtpv2TeidMetricTag
	// Offset returns uint32, set in PatternFlowGtpv2TeidMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv2TeidMetricTag
	SetOffset(value uint32) PatternFlowGtpv2TeidMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv2TeidMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv2TeidMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv2TeidMetricTag
	SetLength(value uint32) PatternFlowGtpv2TeidMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv2TeidMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv2TeidMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv2TeidMetricTag object
func (obj *patternFlowGtpv2TeidMetricTag) SetName(value string) PatternFlowGtpv2TeidMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2TeidMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2TeidMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv2TeidMetricTag object
func (obj *patternFlowGtpv2TeidMetricTag) SetOffset(value uint32) PatternFlowGtpv2TeidMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2TeidMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2TeidMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv2TeidMetricTag object
func (obj *patternFlowGtpv2TeidMetricTag) SetLength(value uint32) PatternFlowGtpv2TeidMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv2TeidMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv2TeidMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv2TeidMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv2TeidMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
