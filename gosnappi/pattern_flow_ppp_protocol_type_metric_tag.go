package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPppProtocolTypeMetricTag *****
type patternFlowPppProtocolTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowPppProtocolTypeMetricTag
	marshaller   marshalPatternFlowPppProtocolTypeMetricTag
	unMarshaller unMarshalPatternFlowPppProtocolTypeMetricTag
}

func NewPatternFlowPppProtocolTypeMetricTag() PatternFlowPppProtocolTypeMetricTag {
	obj := patternFlowPppProtocolTypeMetricTag{obj: &otg.PatternFlowPppProtocolTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppProtocolTypeMetricTag) msg() *otg.PatternFlowPppProtocolTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowPppProtocolTypeMetricTag) setMsg(msg *otg.PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppProtocolTypeMetricTag struct {
	obj *patternFlowPppProtocolTypeMetricTag
}

type marshalPatternFlowPppProtocolTypeMetricTag interface {
	// ToProto marshals PatternFlowPppProtocolTypeMetricTag to protobuf object *otg.PatternFlowPppProtocolTypeMetricTag
	ToProto() (*otg.PatternFlowPppProtocolTypeMetricTag, error)
	// ToPbText marshals PatternFlowPppProtocolTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppProtocolTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppProtocolTypeMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPppProtocolTypeMetricTag struct {
	obj *patternFlowPppProtocolTypeMetricTag
}

type unMarshalPatternFlowPppProtocolTypeMetricTag interface {
	// FromProto unmarshals PatternFlowPppProtocolTypeMetricTag from protobuf object *otg.PatternFlowPppProtocolTypeMetricTag
	FromProto(msg *otg.PatternFlowPppProtocolTypeMetricTag) (PatternFlowPppProtocolTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowPppProtocolTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppProtocolTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppProtocolTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppProtocolTypeMetricTag) Marshal() marshalPatternFlowPppProtocolTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppProtocolTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppProtocolTypeMetricTag) Unmarshal() unMarshalPatternFlowPppProtocolTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppProtocolTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppProtocolTypeMetricTag) ToProto() (*otg.PatternFlowPppProtocolTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppProtocolTypeMetricTag) FromProto(msg *otg.PatternFlowPppProtocolTypeMetricTag) (PatternFlowPppProtocolTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppProtocolTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppProtocolTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppProtocolTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowPppProtocolTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppProtocolTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppProtocolTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppProtocolTypeMetricTag) Clone() (PatternFlowPppProtocolTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppProtocolTypeMetricTag()
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

// PatternFlowPppProtocolTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPppProtocolTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowPppProtocolTypeMetricTag to protobuf object *otg.PatternFlowPppProtocolTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPppProtocolTypeMetricTag
	// setMsg unmarshals PatternFlowPppProtocolTypeMetricTag from protobuf object *otg.PatternFlowPppProtocolTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPppProtocolTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppProtocolTypeMetricTag
	// validate validates PatternFlowPppProtocolTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppProtocolTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPppProtocolTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPppProtocolTypeMetricTag
	SetName(value string) PatternFlowPppProtocolTypeMetricTag
	// Offset returns uint32, set in PatternFlowPppProtocolTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPppProtocolTypeMetricTag
	SetOffset(value uint32) PatternFlowPppProtocolTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowPppProtocolTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPppProtocolTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPppProtocolTypeMetricTag
	SetLength(value uint32) PatternFlowPppProtocolTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowPppProtocolTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPppProtocolTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPppProtocolTypeMetricTag object
func (obj *patternFlowPppProtocolTypeMetricTag) SetName(value string) PatternFlowPppProtocolTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPppProtocolTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPppProtocolTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPppProtocolTypeMetricTag object
func (obj *patternFlowPppProtocolTypeMetricTag) SetOffset(value uint32) PatternFlowPppProtocolTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPppProtocolTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPppProtocolTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPppProtocolTypeMetricTag object
func (obj *patternFlowPppProtocolTypeMetricTag) SetLength(value uint32) PatternFlowPppProtocolTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPppProtocolTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPppProtocolTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolTypeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPppProtocolTypeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPppProtocolTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
