package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoSequenceNumberMetricTag *****
type patternFlowIcmpv6EchoSequenceNumberMetricTag struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag
	marshaller   marshalPatternFlowIcmpv6EchoSequenceNumberMetricTag
	unMarshaller unMarshalPatternFlowIcmpv6EchoSequenceNumberMetricTag
}

func NewPatternFlowIcmpv6EchoSequenceNumberMetricTag() PatternFlowIcmpv6EchoSequenceNumberMetricTag {
	obj := patternFlowIcmpv6EchoSequenceNumberMetricTag{obj: &otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) msg() *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) setMsg(msg *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoSequenceNumberMetricTag struct {
	obj *patternFlowIcmpv6EchoSequenceNumberMetricTag
}

type marshalPatternFlowIcmpv6EchoSequenceNumberMetricTag interface {
	// ToProto marshals PatternFlowIcmpv6EchoSequenceNumberMetricTag to protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag
	ToProto() (*otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag, error)
	// ToPbText marshals PatternFlowIcmpv6EchoSequenceNumberMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoSequenceNumberMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoSequenceNumberMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpv6EchoSequenceNumberMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoSequenceNumberMetricTag struct {
	obj *patternFlowIcmpv6EchoSequenceNumberMetricTag
}

type unMarshalPatternFlowIcmpv6EchoSequenceNumberMetricTag interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoSequenceNumberMetricTag from protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag
	FromProto(msg *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag) (PatternFlowIcmpv6EchoSequenceNumberMetricTag, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoSequenceNumberMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoSequenceNumberMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoSequenceNumberMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) Marshal() marshalPatternFlowIcmpv6EchoSequenceNumberMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoSequenceNumberMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) Unmarshal() unMarshalPatternFlowIcmpv6EchoSequenceNumberMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoSequenceNumberMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) ToProto() (*otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) FromProto(msg *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag) (PatternFlowIcmpv6EchoSequenceNumberMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumberMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) Clone() (PatternFlowIcmpv6EchoSequenceNumberMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoSequenceNumberMetricTag()
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

// PatternFlowIcmpv6EchoSequenceNumberMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIcmpv6EchoSequenceNumberMetricTag interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoSequenceNumberMetricTag to protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag
	// setMsg unmarshals PatternFlowIcmpv6EchoSequenceNumberMetricTag from protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoSequenceNumberMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoSequenceNumberMetricTag
	// validate validates PatternFlowIcmpv6EchoSequenceNumberMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoSequenceNumberMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIcmpv6EchoSequenceNumberMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIcmpv6EchoSequenceNumberMetricTag
	SetName(value string) PatternFlowIcmpv6EchoSequenceNumberMetricTag
	// Offset returns uint32, set in PatternFlowIcmpv6EchoSequenceNumberMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIcmpv6EchoSequenceNumberMetricTag
	SetOffset(value uint32) PatternFlowIcmpv6EchoSequenceNumberMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIcmpv6EchoSequenceNumberMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIcmpv6EchoSequenceNumberMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIcmpv6EchoSequenceNumberMetricTag
	SetLength(value uint32) PatternFlowIcmpv6EchoSequenceNumberMetricTag
	// HasLength checks if Length has been set in PatternFlowIcmpv6EchoSequenceNumberMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIcmpv6EchoSequenceNumberMetricTag object
func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) SetName(value string) PatternFlowIcmpv6EchoSequenceNumberMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIcmpv6EchoSequenceNumberMetricTag object
func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) SetOffset(value uint32) PatternFlowIcmpv6EchoSequenceNumberMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIcmpv6EchoSequenceNumberMetricTag object
func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) SetLength(value uint32) PatternFlowIcmpv6EchoSequenceNumberMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIcmpv6EchoSequenceNumberMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumberMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIcmpv6EchoSequenceNumberMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIcmpv6EchoSequenceNumberMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
