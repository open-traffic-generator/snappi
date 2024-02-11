package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1PnFlagMetricTag *****
type patternFlowGtpv1PnFlagMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1PnFlagMetricTag
	marshaller   marshalPatternFlowGtpv1PnFlagMetricTag
	unMarshaller unMarshalPatternFlowGtpv1PnFlagMetricTag
}

func NewPatternFlowGtpv1PnFlagMetricTag() PatternFlowGtpv1PnFlagMetricTag {
	obj := patternFlowGtpv1PnFlagMetricTag{obj: &otg.PatternFlowGtpv1PnFlagMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1PnFlagMetricTag) msg() *otg.PatternFlowGtpv1PnFlagMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1PnFlagMetricTag) setMsg(msg *otg.PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1PnFlagMetricTag struct {
	obj *patternFlowGtpv1PnFlagMetricTag
}

type marshalPatternFlowGtpv1PnFlagMetricTag interface {
	// ToProto marshals PatternFlowGtpv1PnFlagMetricTag to protobuf object *otg.PatternFlowGtpv1PnFlagMetricTag
	ToProto() (*otg.PatternFlowGtpv1PnFlagMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1PnFlagMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1PnFlagMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1PnFlagMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1PnFlagMetricTag struct {
	obj *patternFlowGtpv1PnFlagMetricTag
}

type unMarshalPatternFlowGtpv1PnFlagMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1PnFlagMetricTag from protobuf object *otg.PatternFlowGtpv1PnFlagMetricTag
	FromProto(msg *otg.PatternFlowGtpv1PnFlagMetricTag) (PatternFlowGtpv1PnFlagMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1PnFlagMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1PnFlagMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1PnFlagMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1PnFlagMetricTag) Marshal() marshalPatternFlowGtpv1PnFlagMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1PnFlagMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1PnFlagMetricTag) Unmarshal() unMarshalPatternFlowGtpv1PnFlagMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1PnFlagMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1PnFlagMetricTag) ToProto() (*otg.PatternFlowGtpv1PnFlagMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1PnFlagMetricTag) FromProto(msg *otg.PatternFlowGtpv1PnFlagMetricTag) (PatternFlowGtpv1PnFlagMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1PnFlagMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlagMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1PnFlagMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlagMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1PnFlagMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlagMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1PnFlagMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1PnFlagMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1PnFlagMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1PnFlagMetricTag) Clone() (PatternFlowGtpv1PnFlagMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1PnFlagMetricTag()
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

// PatternFlowGtpv1PnFlagMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1PnFlagMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1PnFlagMetricTag to protobuf object *otg.PatternFlowGtpv1PnFlagMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1PnFlagMetricTag
	// setMsg unmarshals PatternFlowGtpv1PnFlagMetricTag from protobuf object *otg.PatternFlowGtpv1PnFlagMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1PnFlagMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1PnFlagMetricTag
	// validate validates PatternFlowGtpv1PnFlagMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1PnFlagMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1PnFlagMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1PnFlagMetricTag
	SetName(value string) PatternFlowGtpv1PnFlagMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1PnFlagMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1PnFlagMetricTag
	SetOffset(value uint32) PatternFlowGtpv1PnFlagMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1PnFlagMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1PnFlagMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1PnFlagMetricTag
	SetLength(value uint32) PatternFlowGtpv1PnFlagMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1PnFlagMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1PnFlagMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1PnFlagMetricTag object
func (obj *patternFlowGtpv1PnFlagMetricTag) SetName(value string) PatternFlowGtpv1PnFlagMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1PnFlagMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1PnFlagMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1PnFlagMetricTag object
func (obj *patternFlowGtpv1PnFlagMetricTag) SetOffset(value uint32) PatternFlowGtpv1PnFlagMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1PnFlagMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1PnFlagMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1PnFlagMetricTag object
func (obj *patternFlowGtpv1PnFlagMetricTag) SetLength(value uint32) PatternFlowGtpv1PnFlagMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1PnFlagMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1PnFlagMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1PnFlagMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1PnFlagMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1PnFlagMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
