package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanIdMetricTag *****
type patternFlowVlanIdMetricTag struct {
	validation
	obj          *otg.PatternFlowVlanIdMetricTag
	marshaller   marshalPatternFlowVlanIdMetricTag
	unMarshaller unMarshalPatternFlowVlanIdMetricTag
}

func NewPatternFlowVlanIdMetricTag() PatternFlowVlanIdMetricTag {
	obj := patternFlowVlanIdMetricTag{obj: &otg.PatternFlowVlanIdMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanIdMetricTag) msg() *otg.PatternFlowVlanIdMetricTag {
	return obj.obj
}

func (obj *patternFlowVlanIdMetricTag) setMsg(msg *otg.PatternFlowVlanIdMetricTag) PatternFlowVlanIdMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanIdMetricTag struct {
	obj *patternFlowVlanIdMetricTag
}

type marshalPatternFlowVlanIdMetricTag interface {
	// ToProto marshals PatternFlowVlanIdMetricTag to protobuf object *otg.PatternFlowVlanIdMetricTag
	ToProto() (*otg.PatternFlowVlanIdMetricTag, error)
	// ToPbText marshals PatternFlowVlanIdMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanIdMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanIdMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVlanIdMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVlanIdMetricTag struct {
	obj *patternFlowVlanIdMetricTag
}

type unMarshalPatternFlowVlanIdMetricTag interface {
	// FromProto unmarshals PatternFlowVlanIdMetricTag from protobuf object *otg.PatternFlowVlanIdMetricTag
	FromProto(msg *otg.PatternFlowVlanIdMetricTag) (PatternFlowVlanIdMetricTag, error)
	// FromPbText unmarshals PatternFlowVlanIdMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanIdMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanIdMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanIdMetricTag) Marshal() marshalPatternFlowVlanIdMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanIdMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanIdMetricTag) Unmarshal() unMarshalPatternFlowVlanIdMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanIdMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanIdMetricTag) ToProto() (*otg.PatternFlowVlanIdMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanIdMetricTag) FromProto(msg *otg.PatternFlowVlanIdMetricTag) (PatternFlowVlanIdMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanIdMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanIdMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanIdMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanIdMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanIdMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVlanIdMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanIdMetricTag) FromJson(value string) error {
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

func (obj *patternFlowVlanIdMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanIdMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanIdMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanIdMetricTag) Clone() (PatternFlowVlanIdMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanIdMetricTag()
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

// PatternFlowVlanIdMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowVlanIdMetricTag interface {
	Validation
	// msg marshals PatternFlowVlanIdMetricTag to protobuf object *otg.PatternFlowVlanIdMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanIdMetricTag
	// setMsg unmarshals PatternFlowVlanIdMetricTag from protobuf object *otg.PatternFlowVlanIdMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanIdMetricTag) PatternFlowVlanIdMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowVlanIdMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanIdMetricTag
	// validate validates PatternFlowVlanIdMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanIdMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowVlanIdMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowVlanIdMetricTag
	SetName(value string) PatternFlowVlanIdMetricTag
	// Offset returns uint32, set in PatternFlowVlanIdMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowVlanIdMetricTag
	SetOffset(value uint32) PatternFlowVlanIdMetricTag
	// HasOffset checks if Offset has been set in PatternFlowVlanIdMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowVlanIdMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowVlanIdMetricTag
	SetLength(value uint32) PatternFlowVlanIdMetricTag
	// HasLength checks if Length has been set in PatternFlowVlanIdMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowVlanIdMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowVlanIdMetricTag object
func (obj *patternFlowVlanIdMetricTag) SetName(value string) PatternFlowVlanIdMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVlanIdMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVlanIdMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowVlanIdMetricTag object
func (obj *patternFlowVlanIdMetricTag) SetOffset(value uint32) PatternFlowVlanIdMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVlanIdMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVlanIdMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowVlanIdMetricTag object
func (obj *patternFlowVlanIdMetricTag) SetLength(value uint32) PatternFlowVlanIdMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowVlanIdMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowVlanIdMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 11 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanIdMetricTag.Offset <= 11 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 12 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowVlanIdMetricTag.Length <= 12 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowVlanIdMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(12)
	}

}
