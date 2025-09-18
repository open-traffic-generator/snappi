package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass4MetricTag *****
type patternFlowPfcPausePauseClass4MetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass4MetricTag
	marshaller   marshalPatternFlowPfcPausePauseClass4MetricTag
	unMarshaller unMarshalPatternFlowPfcPausePauseClass4MetricTag
}

func NewPatternFlowPfcPausePauseClass4MetricTag() PatternFlowPfcPausePauseClass4MetricTag {
	obj := patternFlowPfcPausePauseClass4MetricTag{obj: &otg.PatternFlowPfcPausePauseClass4MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass4MetricTag) msg() *otg.PatternFlowPfcPausePauseClass4MetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass4MetricTag) setMsg(msg *otg.PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass4MetricTag struct {
	obj *patternFlowPfcPausePauseClass4MetricTag
}

type marshalPatternFlowPfcPausePauseClass4MetricTag interface {
	// ToProto marshals PatternFlowPfcPausePauseClass4MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass4MetricTag
	ToProto() (*otg.PatternFlowPfcPausePauseClass4MetricTag, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass4MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass4MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass4MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass4MetricTag struct {
	obj *patternFlowPfcPausePauseClass4MetricTag
}

type unMarshalPatternFlowPfcPausePauseClass4MetricTag interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass4MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass4MetricTag
	FromProto(msg *otg.PatternFlowPfcPausePauseClass4MetricTag) (PatternFlowPfcPausePauseClass4MetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass4MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass4MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass4MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass4MetricTag) Marshal() marshalPatternFlowPfcPausePauseClass4MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass4MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass4MetricTag) Unmarshal() unMarshalPatternFlowPfcPausePauseClass4MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass4MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass4MetricTag) ToProto() (*otg.PatternFlowPfcPausePauseClass4MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass4MetricTag) FromProto(msg *otg.PatternFlowPfcPausePauseClass4MetricTag) (PatternFlowPfcPausePauseClass4MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass4MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass4MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass4MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4MetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass4MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass4MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass4MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass4MetricTag) Clone() (PatternFlowPfcPausePauseClass4MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass4MetricTag()
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

// PatternFlowPfcPausePauseClass4MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPausePauseClass4MetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass4MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass4MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass4MetricTag
	// setMsg unmarshals PatternFlowPfcPausePauseClass4MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass4MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass4MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass4MetricTag
	// validate validates PatternFlowPfcPausePauseClass4MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass4MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPausePauseClass4MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPausePauseClass4MetricTag
	SetName(value string) PatternFlowPfcPausePauseClass4MetricTag
	// Offset returns uint32, set in PatternFlowPfcPausePauseClass4MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPausePauseClass4MetricTag
	SetOffset(value uint32) PatternFlowPfcPausePauseClass4MetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPausePauseClass4MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPausePauseClass4MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPausePauseClass4MetricTag
	SetLength(value uint32) PatternFlowPfcPausePauseClass4MetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPausePauseClass4MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPausePauseClass4MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPausePauseClass4MetricTag object
func (obj *patternFlowPfcPausePauseClass4MetricTag) SetName(value string) PatternFlowPfcPausePauseClass4MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass4MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass4MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPausePauseClass4MetricTag object
func (obj *patternFlowPfcPausePauseClass4MetricTag) SetOffset(value uint32) PatternFlowPfcPausePauseClass4MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass4MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass4MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPausePauseClass4MetricTag object
func (obj *patternFlowPfcPausePauseClass4MetricTag) SetLength(value uint32) PatternFlowPfcPausePauseClass4MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass4MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPausePauseClass4MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4MetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPausePauseClass4MetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass4MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
