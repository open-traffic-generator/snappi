package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpLengthMetricTag *****
type patternFlowUdpLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowUdpLengthMetricTag
	marshaller   marshalPatternFlowUdpLengthMetricTag
	unMarshaller unMarshalPatternFlowUdpLengthMetricTag
}

func NewPatternFlowUdpLengthMetricTag() PatternFlowUdpLengthMetricTag {
	obj := patternFlowUdpLengthMetricTag{obj: &otg.PatternFlowUdpLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpLengthMetricTag) msg() *otg.PatternFlowUdpLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowUdpLengthMetricTag) setMsg(msg *otg.PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpLengthMetricTag struct {
	obj *patternFlowUdpLengthMetricTag
}

type marshalPatternFlowUdpLengthMetricTag interface {
	// ToProto marshals PatternFlowUdpLengthMetricTag to protobuf object *otg.PatternFlowUdpLengthMetricTag
	ToProto() (*otg.PatternFlowUdpLengthMetricTag, error)
	// ToPbText marshals PatternFlowUdpLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpLengthMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowUdpLengthMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowUdpLengthMetricTag struct {
	obj *patternFlowUdpLengthMetricTag
}

type unMarshalPatternFlowUdpLengthMetricTag interface {
	// FromProto unmarshals PatternFlowUdpLengthMetricTag from protobuf object *otg.PatternFlowUdpLengthMetricTag
	FromProto(msg *otg.PatternFlowUdpLengthMetricTag) (PatternFlowUdpLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowUdpLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpLengthMetricTag) Marshal() marshalPatternFlowUdpLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpLengthMetricTag) Unmarshal() unMarshalPatternFlowUdpLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpLengthMetricTag) ToProto() (*otg.PatternFlowUdpLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpLengthMetricTag) FromProto(msg *otg.PatternFlowUdpLengthMetricTag) (PatternFlowUdpLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpLengthMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowUdpLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowUdpLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpLengthMetricTag) Clone() (PatternFlowUdpLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpLengthMetricTag()
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

// PatternFlowUdpLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowUdpLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowUdpLengthMetricTag to protobuf object *otg.PatternFlowUdpLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpLengthMetricTag
	// setMsg unmarshals PatternFlowUdpLengthMetricTag from protobuf object *otg.PatternFlowUdpLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowUdpLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpLengthMetricTag
	// validate validates PatternFlowUdpLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowUdpLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowUdpLengthMetricTag
	SetName(value string) PatternFlowUdpLengthMetricTag
	// Offset returns uint32, set in PatternFlowUdpLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowUdpLengthMetricTag
	SetOffset(value uint32) PatternFlowUdpLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowUdpLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowUdpLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowUdpLengthMetricTag
	SetLength(value uint32) PatternFlowUdpLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowUdpLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowUdpLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowUdpLengthMetricTag object
func (obj *patternFlowUdpLengthMetricTag) SetName(value string) PatternFlowUdpLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowUdpLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowUdpLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowUdpLengthMetricTag object
func (obj *patternFlowUdpLengthMetricTag) SetOffset(value uint32) PatternFlowUdpLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowUdpLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowUdpLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowUdpLengthMetricTag object
func (obj *patternFlowUdpLengthMetricTag) SetLength(value uint32) PatternFlowUdpLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowUdpLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowUdpLengthMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpLengthMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowUdpLengthMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowUdpLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
