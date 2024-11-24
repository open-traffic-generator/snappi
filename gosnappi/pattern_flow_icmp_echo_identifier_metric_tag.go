package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoIdentifierMetricTag *****
type patternFlowIcmpEchoIdentifierMetricTag struct {
	validation
	obj          *otg.PatternFlowIcmpEchoIdentifierMetricTag
	marshaller   marshalPatternFlowIcmpEchoIdentifierMetricTag
	unMarshaller unMarshalPatternFlowIcmpEchoIdentifierMetricTag
}

func NewPatternFlowIcmpEchoIdentifierMetricTag() PatternFlowIcmpEchoIdentifierMetricTag {
	obj := patternFlowIcmpEchoIdentifierMetricTag{obj: &otg.PatternFlowIcmpEchoIdentifierMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) msg() *otg.PatternFlowIcmpEchoIdentifierMetricTag {
	return obj.obj
}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) setMsg(msg *otg.PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoIdentifierMetricTag struct {
	obj *patternFlowIcmpEchoIdentifierMetricTag
}

type marshalPatternFlowIcmpEchoIdentifierMetricTag interface {
	// ToProto marshals PatternFlowIcmpEchoIdentifierMetricTag to protobuf object *otg.PatternFlowIcmpEchoIdentifierMetricTag
	ToProto() (*otg.PatternFlowIcmpEchoIdentifierMetricTag, error)
	// ToPbText marshals PatternFlowIcmpEchoIdentifierMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoIdentifierMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoIdentifierMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpEchoIdentifierMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpEchoIdentifierMetricTag struct {
	obj *patternFlowIcmpEchoIdentifierMetricTag
}

type unMarshalPatternFlowIcmpEchoIdentifierMetricTag interface {
	// FromProto unmarshals PatternFlowIcmpEchoIdentifierMetricTag from protobuf object *otg.PatternFlowIcmpEchoIdentifierMetricTag
	FromProto(msg *otg.PatternFlowIcmpEchoIdentifierMetricTag) (PatternFlowIcmpEchoIdentifierMetricTag, error)
	// FromPbText unmarshals PatternFlowIcmpEchoIdentifierMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoIdentifierMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoIdentifierMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) Marshal() marshalPatternFlowIcmpEchoIdentifierMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoIdentifierMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) Unmarshal() unMarshalPatternFlowIcmpEchoIdentifierMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoIdentifierMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoIdentifierMetricTag) ToProto() (*otg.PatternFlowIcmpEchoIdentifierMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoIdentifierMetricTag) FromProto(msg *otg.PatternFlowIcmpEchoIdentifierMetricTag) (PatternFlowIcmpEchoIdentifierMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoIdentifierMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifierMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoIdentifierMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifierMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoIdentifierMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpEchoIdentifierMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifierMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoIdentifierMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) Clone() (PatternFlowIcmpEchoIdentifierMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoIdentifierMetricTag()
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

// PatternFlowIcmpEchoIdentifierMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIcmpEchoIdentifierMetricTag interface {
	Validation
	// msg marshals PatternFlowIcmpEchoIdentifierMetricTag to protobuf object *otg.PatternFlowIcmpEchoIdentifierMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoIdentifierMetricTag
	// setMsg unmarshals PatternFlowIcmpEchoIdentifierMetricTag from protobuf object *otg.PatternFlowIcmpEchoIdentifierMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoIdentifierMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoIdentifierMetricTag
	// validate validates PatternFlowIcmpEchoIdentifierMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoIdentifierMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIcmpEchoIdentifierMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIcmpEchoIdentifierMetricTag
	SetName(value string) PatternFlowIcmpEchoIdentifierMetricTag
	// Offset returns uint32, set in PatternFlowIcmpEchoIdentifierMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIcmpEchoIdentifierMetricTag
	SetOffset(value uint32) PatternFlowIcmpEchoIdentifierMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIcmpEchoIdentifierMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIcmpEchoIdentifierMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIcmpEchoIdentifierMetricTag
	SetLength(value uint32) PatternFlowIcmpEchoIdentifierMetricTag
	// HasLength checks if Length has been set in PatternFlowIcmpEchoIdentifierMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIcmpEchoIdentifierMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIcmpEchoIdentifierMetricTag object
func (obj *patternFlowIcmpEchoIdentifierMetricTag) SetName(value string) PatternFlowIcmpEchoIdentifierMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpEchoIdentifierMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpEchoIdentifierMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIcmpEchoIdentifierMetricTag object
func (obj *patternFlowIcmpEchoIdentifierMetricTag) SetOffset(value uint32) PatternFlowIcmpEchoIdentifierMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpEchoIdentifierMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpEchoIdentifierMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIcmpEchoIdentifierMetricTag object
func (obj *patternFlowIcmpEchoIdentifierMetricTag) SetLength(value uint32) PatternFlowIcmpEchoIdentifierMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIcmpEchoIdentifierMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifierMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIcmpEchoIdentifierMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIcmpEchoIdentifierMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
