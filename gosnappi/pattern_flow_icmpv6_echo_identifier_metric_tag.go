package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoIdentifierMetricTag *****
type patternFlowIcmpv6EchoIdentifierMetricTag struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoIdentifierMetricTag
	marshaller   marshalPatternFlowIcmpv6EchoIdentifierMetricTag
	unMarshaller unMarshalPatternFlowIcmpv6EchoIdentifierMetricTag
}

func NewPatternFlowIcmpv6EchoIdentifierMetricTag() PatternFlowIcmpv6EchoIdentifierMetricTag {
	obj := patternFlowIcmpv6EchoIdentifierMetricTag{obj: &otg.PatternFlowIcmpv6EchoIdentifierMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) msg() *otg.PatternFlowIcmpv6EchoIdentifierMetricTag {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) setMsg(msg *otg.PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoIdentifierMetricTag struct {
	obj *patternFlowIcmpv6EchoIdentifierMetricTag
}

type marshalPatternFlowIcmpv6EchoIdentifierMetricTag interface {
	// ToProto marshals PatternFlowIcmpv6EchoIdentifierMetricTag to protobuf object *otg.PatternFlowIcmpv6EchoIdentifierMetricTag
	ToProto() (*otg.PatternFlowIcmpv6EchoIdentifierMetricTag, error)
	// ToPbText marshals PatternFlowIcmpv6EchoIdentifierMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoIdentifierMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoIdentifierMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoIdentifierMetricTag struct {
	obj *patternFlowIcmpv6EchoIdentifierMetricTag
}

type unMarshalPatternFlowIcmpv6EchoIdentifierMetricTag interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoIdentifierMetricTag from protobuf object *otg.PatternFlowIcmpv6EchoIdentifierMetricTag
	FromProto(msg *otg.PatternFlowIcmpv6EchoIdentifierMetricTag) (PatternFlowIcmpv6EchoIdentifierMetricTag, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoIdentifierMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoIdentifierMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoIdentifierMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) Marshal() marshalPatternFlowIcmpv6EchoIdentifierMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoIdentifierMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) Unmarshal() unMarshalPatternFlowIcmpv6EchoIdentifierMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoIdentifierMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoIdentifierMetricTag) ToProto() (*otg.PatternFlowIcmpv6EchoIdentifierMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoIdentifierMetricTag) FromProto(msg *otg.PatternFlowIcmpv6EchoIdentifierMetricTag) (PatternFlowIcmpv6EchoIdentifierMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoIdentifierMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifierMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoIdentifierMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifierMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoIdentifierMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifierMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) Clone() (PatternFlowIcmpv6EchoIdentifierMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoIdentifierMetricTag()
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

// PatternFlowIcmpv6EchoIdentifierMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIcmpv6EchoIdentifierMetricTag interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoIdentifierMetricTag to protobuf object *otg.PatternFlowIcmpv6EchoIdentifierMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoIdentifierMetricTag
	// setMsg unmarshals PatternFlowIcmpv6EchoIdentifierMetricTag from protobuf object *otg.PatternFlowIcmpv6EchoIdentifierMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoIdentifierMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoIdentifierMetricTag
	// validate validates PatternFlowIcmpv6EchoIdentifierMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoIdentifierMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIcmpv6EchoIdentifierMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIcmpv6EchoIdentifierMetricTag
	SetName(value string) PatternFlowIcmpv6EchoIdentifierMetricTag
	// Offset returns uint32, set in PatternFlowIcmpv6EchoIdentifierMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIcmpv6EchoIdentifierMetricTag
	SetOffset(value uint32) PatternFlowIcmpv6EchoIdentifierMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIcmpv6EchoIdentifierMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIcmpv6EchoIdentifierMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIcmpv6EchoIdentifierMetricTag
	SetLength(value uint32) PatternFlowIcmpv6EchoIdentifierMetricTag
	// HasLength checks if Length has been set in PatternFlowIcmpv6EchoIdentifierMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIcmpv6EchoIdentifierMetricTag object
func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) SetName(value string) PatternFlowIcmpv6EchoIdentifierMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIcmpv6EchoIdentifierMetricTag object
func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) SetOffset(value uint32) PatternFlowIcmpv6EchoIdentifierMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIcmpv6EchoIdentifierMetricTag object
func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) SetLength(value uint32) PatternFlowIcmpv6EchoIdentifierMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIcmpv6EchoIdentifierMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifierMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIcmpv6EchoIdentifierMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIcmpv6EchoIdentifierMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
