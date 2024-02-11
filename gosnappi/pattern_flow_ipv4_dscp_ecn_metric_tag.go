package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DscpEcnMetricTag *****
type patternFlowIpv4DscpEcnMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4DscpEcnMetricTag
	marshaller   marshalPatternFlowIpv4DscpEcnMetricTag
	unMarshaller unMarshalPatternFlowIpv4DscpEcnMetricTag
}

func NewPatternFlowIpv4DscpEcnMetricTag() PatternFlowIpv4DscpEcnMetricTag {
	obj := patternFlowIpv4DscpEcnMetricTag{obj: &otg.PatternFlowIpv4DscpEcnMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DscpEcnMetricTag) msg() *otg.PatternFlowIpv4DscpEcnMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4DscpEcnMetricTag) setMsg(msg *otg.PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DscpEcnMetricTag struct {
	obj *patternFlowIpv4DscpEcnMetricTag
}

type marshalPatternFlowIpv4DscpEcnMetricTag interface {
	// ToProto marshals PatternFlowIpv4DscpEcnMetricTag to protobuf object *otg.PatternFlowIpv4DscpEcnMetricTag
	ToProto() (*otg.PatternFlowIpv4DscpEcnMetricTag, error)
	// ToPbText marshals PatternFlowIpv4DscpEcnMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DscpEcnMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DscpEcnMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DscpEcnMetricTag struct {
	obj *patternFlowIpv4DscpEcnMetricTag
}

type unMarshalPatternFlowIpv4DscpEcnMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4DscpEcnMetricTag from protobuf object *otg.PatternFlowIpv4DscpEcnMetricTag
	FromProto(msg *otg.PatternFlowIpv4DscpEcnMetricTag) (PatternFlowIpv4DscpEcnMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4DscpEcnMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DscpEcnMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DscpEcnMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DscpEcnMetricTag) Marshal() marshalPatternFlowIpv4DscpEcnMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DscpEcnMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DscpEcnMetricTag) Unmarshal() unMarshalPatternFlowIpv4DscpEcnMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DscpEcnMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DscpEcnMetricTag) ToProto() (*otg.PatternFlowIpv4DscpEcnMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DscpEcnMetricTag) FromProto(msg *otg.PatternFlowIpv4DscpEcnMetricTag) (PatternFlowIpv4DscpEcnMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DscpEcnMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcnMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DscpEcnMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcnMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DscpEcnMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcnMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4DscpEcnMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpEcnMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpEcnMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DscpEcnMetricTag) Clone() (PatternFlowIpv4DscpEcnMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DscpEcnMetricTag()
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

// PatternFlowIpv4DscpEcnMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4DscpEcnMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4DscpEcnMetricTag to protobuf object *otg.PatternFlowIpv4DscpEcnMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DscpEcnMetricTag
	// setMsg unmarshals PatternFlowIpv4DscpEcnMetricTag from protobuf object *otg.PatternFlowIpv4DscpEcnMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DscpEcnMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DscpEcnMetricTag
	// validate validates PatternFlowIpv4DscpEcnMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DscpEcnMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4DscpEcnMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4DscpEcnMetricTag
	SetName(value string) PatternFlowIpv4DscpEcnMetricTag
	// Offset returns uint32, set in PatternFlowIpv4DscpEcnMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4DscpEcnMetricTag
	SetOffset(value uint32) PatternFlowIpv4DscpEcnMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4DscpEcnMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4DscpEcnMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4DscpEcnMetricTag
	SetLength(value uint32) PatternFlowIpv4DscpEcnMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4DscpEcnMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4DscpEcnMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4DscpEcnMetricTag object
func (obj *patternFlowIpv4DscpEcnMetricTag) SetName(value string) PatternFlowIpv4DscpEcnMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4DscpEcnMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4DscpEcnMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4DscpEcnMetricTag object
func (obj *patternFlowIpv4DscpEcnMetricTag) SetOffset(value uint32) PatternFlowIpv4DscpEcnMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4DscpEcnMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4DscpEcnMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4DscpEcnMetricTag object
func (obj *patternFlowIpv4DscpEcnMetricTag) SetLength(value uint32) PatternFlowIpv4DscpEcnMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4DscpEcnMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4DscpEcnMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpEcnMetricTag.Offset <= 1 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4DscpEcnMetricTag.Length <= 2 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4DscpEcnMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(2)
	}

}
