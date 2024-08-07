package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreReserved1MetricTag *****
type patternFlowGreReserved1MetricTag struct {
	validation
	obj          *otg.PatternFlowGreReserved1MetricTag
	marshaller   marshalPatternFlowGreReserved1MetricTag
	unMarshaller unMarshalPatternFlowGreReserved1MetricTag
}

func NewPatternFlowGreReserved1MetricTag() PatternFlowGreReserved1MetricTag {
	obj := patternFlowGreReserved1MetricTag{obj: &otg.PatternFlowGreReserved1MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreReserved1MetricTag) msg() *otg.PatternFlowGreReserved1MetricTag {
	return obj.obj
}

func (obj *patternFlowGreReserved1MetricTag) setMsg(msg *otg.PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreReserved1MetricTag struct {
	obj *patternFlowGreReserved1MetricTag
}

type marshalPatternFlowGreReserved1MetricTag interface {
	// ToProto marshals PatternFlowGreReserved1MetricTag to protobuf object *otg.PatternFlowGreReserved1MetricTag
	ToProto() (*otg.PatternFlowGreReserved1MetricTag, error)
	// ToPbText marshals PatternFlowGreReserved1MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreReserved1MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreReserved1MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreReserved1MetricTag struct {
	obj *patternFlowGreReserved1MetricTag
}

type unMarshalPatternFlowGreReserved1MetricTag interface {
	// FromProto unmarshals PatternFlowGreReserved1MetricTag from protobuf object *otg.PatternFlowGreReserved1MetricTag
	FromProto(msg *otg.PatternFlowGreReserved1MetricTag) (PatternFlowGreReserved1MetricTag, error)
	// FromPbText unmarshals PatternFlowGreReserved1MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreReserved1MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreReserved1MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreReserved1MetricTag) Marshal() marshalPatternFlowGreReserved1MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreReserved1MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreReserved1MetricTag) Unmarshal() unMarshalPatternFlowGreReserved1MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreReserved1MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreReserved1MetricTag) ToProto() (*otg.PatternFlowGreReserved1MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreReserved1MetricTag) FromProto(msg *otg.PatternFlowGreReserved1MetricTag) (PatternFlowGreReserved1MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreReserved1MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreReserved1MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreReserved1MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1MetricTag) FromJson(value string) error {
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

func (obj *patternFlowGreReserved1MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved1MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved1MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreReserved1MetricTag) Clone() (PatternFlowGreReserved1MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreReserved1MetricTag()
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

// PatternFlowGreReserved1MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGreReserved1MetricTag interface {
	Validation
	// msg marshals PatternFlowGreReserved1MetricTag to protobuf object *otg.PatternFlowGreReserved1MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGreReserved1MetricTag
	// setMsg unmarshals PatternFlowGreReserved1MetricTag from protobuf object *otg.PatternFlowGreReserved1MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGreReserved1MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreReserved1MetricTag
	// validate validates PatternFlowGreReserved1MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreReserved1MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGreReserved1MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGreReserved1MetricTag
	SetName(value string) PatternFlowGreReserved1MetricTag
	// Offset returns uint32, set in PatternFlowGreReserved1MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGreReserved1MetricTag
	SetOffset(value uint32) PatternFlowGreReserved1MetricTag
	// HasOffset checks if Offset has been set in PatternFlowGreReserved1MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGreReserved1MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGreReserved1MetricTag
	SetLength(value uint32) PatternFlowGreReserved1MetricTag
	// HasLength checks if Length has been set in PatternFlowGreReserved1MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGreReserved1MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGreReserved1MetricTag object
func (obj *patternFlowGreReserved1MetricTag) SetName(value string) PatternFlowGreReserved1MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreReserved1MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreReserved1MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGreReserved1MetricTag object
func (obj *patternFlowGreReserved1MetricTag) SetOffset(value uint32) PatternFlowGreReserved1MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreReserved1MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreReserved1MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGreReserved1MetricTag object
func (obj *patternFlowGreReserved1MetricTag) SetLength(value uint32) PatternFlowGreReserved1MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGreReserved1MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGreReserved1MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved1MetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGreReserved1MetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGreReserved1MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
