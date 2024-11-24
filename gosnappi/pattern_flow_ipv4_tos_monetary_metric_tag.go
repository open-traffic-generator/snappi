package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosMonetaryMetricTag *****
type patternFlowIpv4TosMonetaryMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4TosMonetaryMetricTag
	marshaller   marshalPatternFlowIpv4TosMonetaryMetricTag
	unMarshaller unMarshalPatternFlowIpv4TosMonetaryMetricTag
}

func NewPatternFlowIpv4TosMonetaryMetricTag() PatternFlowIpv4TosMonetaryMetricTag {
	obj := patternFlowIpv4TosMonetaryMetricTag{obj: &otg.PatternFlowIpv4TosMonetaryMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosMonetaryMetricTag) msg() *otg.PatternFlowIpv4TosMonetaryMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4TosMonetaryMetricTag) setMsg(msg *otg.PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosMonetaryMetricTag struct {
	obj *patternFlowIpv4TosMonetaryMetricTag
}

type marshalPatternFlowIpv4TosMonetaryMetricTag interface {
	// ToProto marshals PatternFlowIpv4TosMonetaryMetricTag to protobuf object *otg.PatternFlowIpv4TosMonetaryMetricTag
	ToProto() (*otg.PatternFlowIpv4TosMonetaryMetricTag, error)
	// ToPbText marshals PatternFlowIpv4TosMonetaryMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosMonetaryMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosMonetaryMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4TosMonetaryMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4TosMonetaryMetricTag struct {
	obj *patternFlowIpv4TosMonetaryMetricTag
}

type unMarshalPatternFlowIpv4TosMonetaryMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4TosMonetaryMetricTag from protobuf object *otg.PatternFlowIpv4TosMonetaryMetricTag
	FromProto(msg *otg.PatternFlowIpv4TosMonetaryMetricTag) (PatternFlowIpv4TosMonetaryMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4TosMonetaryMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosMonetaryMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosMonetaryMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosMonetaryMetricTag) Marshal() marshalPatternFlowIpv4TosMonetaryMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosMonetaryMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosMonetaryMetricTag) Unmarshal() unMarshalPatternFlowIpv4TosMonetaryMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosMonetaryMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosMonetaryMetricTag) ToProto() (*otg.PatternFlowIpv4TosMonetaryMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosMonetaryMetricTag) FromProto(msg *otg.PatternFlowIpv4TosMonetaryMetricTag) (PatternFlowIpv4TosMonetaryMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosMonetaryMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetaryMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosMonetaryMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetaryMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosMonetaryMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4TosMonetaryMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetaryMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosMonetaryMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosMonetaryMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosMonetaryMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosMonetaryMetricTag) Clone() (PatternFlowIpv4TosMonetaryMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosMonetaryMetricTag()
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

// PatternFlowIpv4TosMonetaryMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4TosMonetaryMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4TosMonetaryMetricTag to protobuf object *otg.PatternFlowIpv4TosMonetaryMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosMonetaryMetricTag
	// setMsg unmarshals PatternFlowIpv4TosMonetaryMetricTag from protobuf object *otg.PatternFlowIpv4TosMonetaryMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosMonetaryMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosMonetaryMetricTag
	// validate validates PatternFlowIpv4TosMonetaryMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosMonetaryMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4TosMonetaryMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4TosMonetaryMetricTag
	SetName(value string) PatternFlowIpv4TosMonetaryMetricTag
	// Offset returns uint32, set in PatternFlowIpv4TosMonetaryMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4TosMonetaryMetricTag
	SetOffset(value uint32) PatternFlowIpv4TosMonetaryMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4TosMonetaryMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4TosMonetaryMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4TosMonetaryMetricTag
	SetLength(value uint32) PatternFlowIpv4TosMonetaryMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4TosMonetaryMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4TosMonetaryMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4TosMonetaryMetricTag object
func (obj *patternFlowIpv4TosMonetaryMetricTag) SetName(value string) PatternFlowIpv4TosMonetaryMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosMonetaryMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosMonetaryMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4TosMonetaryMetricTag object
func (obj *patternFlowIpv4TosMonetaryMetricTag) SetOffset(value uint32) PatternFlowIpv4TosMonetaryMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosMonetaryMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosMonetaryMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4TosMonetaryMetricTag object
func (obj *patternFlowIpv4TosMonetaryMetricTag) SetLength(value uint32) PatternFlowIpv4TosMonetaryMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4TosMonetaryMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4TosMonetaryMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosMonetaryMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4TosMonetaryMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4TosMonetaryMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
