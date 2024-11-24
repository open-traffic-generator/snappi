package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4IdentificationMetricTag *****
type patternFlowIpv4IdentificationMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4IdentificationMetricTag
	marshaller   marshalPatternFlowIpv4IdentificationMetricTag
	unMarshaller unMarshalPatternFlowIpv4IdentificationMetricTag
}

func NewPatternFlowIpv4IdentificationMetricTag() PatternFlowIpv4IdentificationMetricTag {
	obj := patternFlowIpv4IdentificationMetricTag{obj: &otg.PatternFlowIpv4IdentificationMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4IdentificationMetricTag) msg() *otg.PatternFlowIpv4IdentificationMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4IdentificationMetricTag) setMsg(msg *otg.PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4IdentificationMetricTag struct {
	obj *patternFlowIpv4IdentificationMetricTag
}

type marshalPatternFlowIpv4IdentificationMetricTag interface {
	// ToProto marshals PatternFlowIpv4IdentificationMetricTag to protobuf object *otg.PatternFlowIpv4IdentificationMetricTag
	ToProto() (*otg.PatternFlowIpv4IdentificationMetricTag, error)
	// ToPbText marshals PatternFlowIpv4IdentificationMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4IdentificationMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4IdentificationMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4IdentificationMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4IdentificationMetricTag struct {
	obj *patternFlowIpv4IdentificationMetricTag
}

type unMarshalPatternFlowIpv4IdentificationMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4IdentificationMetricTag from protobuf object *otg.PatternFlowIpv4IdentificationMetricTag
	FromProto(msg *otg.PatternFlowIpv4IdentificationMetricTag) (PatternFlowIpv4IdentificationMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4IdentificationMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4IdentificationMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4IdentificationMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4IdentificationMetricTag) Marshal() marshalPatternFlowIpv4IdentificationMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4IdentificationMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4IdentificationMetricTag) Unmarshal() unMarshalPatternFlowIpv4IdentificationMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4IdentificationMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4IdentificationMetricTag) ToProto() (*otg.PatternFlowIpv4IdentificationMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4IdentificationMetricTag) FromProto(msg *otg.PatternFlowIpv4IdentificationMetricTag) (PatternFlowIpv4IdentificationMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4IdentificationMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4IdentificationMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4IdentificationMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4IdentificationMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4IdentificationMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4IdentificationMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4IdentificationMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4IdentificationMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4IdentificationMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4IdentificationMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4IdentificationMetricTag) Clone() (PatternFlowIpv4IdentificationMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4IdentificationMetricTag()
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

// PatternFlowIpv4IdentificationMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4IdentificationMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4IdentificationMetricTag to protobuf object *otg.PatternFlowIpv4IdentificationMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4IdentificationMetricTag
	// setMsg unmarshals PatternFlowIpv4IdentificationMetricTag from protobuf object *otg.PatternFlowIpv4IdentificationMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4IdentificationMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4IdentificationMetricTag
	// validate validates PatternFlowIpv4IdentificationMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4IdentificationMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4IdentificationMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4IdentificationMetricTag
	SetName(value string) PatternFlowIpv4IdentificationMetricTag
	// Offset returns uint32, set in PatternFlowIpv4IdentificationMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4IdentificationMetricTag
	SetOffset(value uint32) PatternFlowIpv4IdentificationMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4IdentificationMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4IdentificationMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4IdentificationMetricTag
	SetLength(value uint32) PatternFlowIpv4IdentificationMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4IdentificationMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4IdentificationMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4IdentificationMetricTag object
func (obj *patternFlowIpv4IdentificationMetricTag) SetName(value string) PatternFlowIpv4IdentificationMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4IdentificationMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4IdentificationMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4IdentificationMetricTag object
func (obj *patternFlowIpv4IdentificationMetricTag) SetOffset(value uint32) PatternFlowIpv4IdentificationMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4IdentificationMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4IdentificationMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4IdentificationMetricTag object
func (obj *patternFlowIpv4IdentificationMetricTag) SetLength(value uint32) PatternFlowIpv4IdentificationMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4IdentificationMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4IdentificationMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4IdentificationMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4IdentificationMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4IdentificationMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
