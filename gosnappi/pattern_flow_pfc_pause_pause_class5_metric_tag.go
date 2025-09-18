package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass5MetricTag *****
type patternFlowPfcPausePauseClass5MetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass5MetricTag
	marshaller   marshalPatternFlowPfcPausePauseClass5MetricTag
	unMarshaller unMarshalPatternFlowPfcPausePauseClass5MetricTag
}

func NewPatternFlowPfcPausePauseClass5MetricTag() PatternFlowPfcPausePauseClass5MetricTag {
	obj := patternFlowPfcPausePauseClass5MetricTag{obj: &otg.PatternFlowPfcPausePauseClass5MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass5MetricTag) msg() *otg.PatternFlowPfcPausePauseClass5MetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass5MetricTag) setMsg(msg *otg.PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass5MetricTag struct {
	obj *patternFlowPfcPausePauseClass5MetricTag
}

type marshalPatternFlowPfcPausePauseClass5MetricTag interface {
	// ToProto marshals PatternFlowPfcPausePauseClass5MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass5MetricTag
	ToProto() (*otg.PatternFlowPfcPausePauseClass5MetricTag, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass5MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass5MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass5MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass5MetricTag struct {
	obj *patternFlowPfcPausePauseClass5MetricTag
}

type unMarshalPatternFlowPfcPausePauseClass5MetricTag interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass5MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass5MetricTag
	FromProto(msg *otg.PatternFlowPfcPausePauseClass5MetricTag) (PatternFlowPfcPausePauseClass5MetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass5MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass5MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass5MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass5MetricTag) Marshal() marshalPatternFlowPfcPausePauseClass5MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass5MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass5MetricTag) Unmarshal() unMarshalPatternFlowPfcPausePauseClass5MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass5MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass5MetricTag) ToProto() (*otg.PatternFlowPfcPausePauseClass5MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass5MetricTag) FromProto(msg *otg.PatternFlowPfcPausePauseClass5MetricTag) (PatternFlowPfcPausePauseClass5MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass5MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass5MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass5MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5MetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass5MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass5MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass5MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass5MetricTag) Clone() (PatternFlowPfcPausePauseClass5MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass5MetricTag()
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

// PatternFlowPfcPausePauseClass5MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPausePauseClass5MetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass5MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass5MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass5MetricTag
	// setMsg unmarshals PatternFlowPfcPausePauseClass5MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass5MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass5MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass5MetricTag
	// validate validates PatternFlowPfcPausePauseClass5MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass5MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPausePauseClass5MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPausePauseClass5MetricTag
	SetName(value string) PatternFlowPfcPausePauseClass5MetricTag
	// Offset returns uint32, set in PatternFlowPfcPausePauseClass5MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPausePauseClass5MetricTag
	SetOffset(value uint32) PatternFlowPfcPausePauseClass5MetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPausePauseClass5MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPausePauseClass5MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPausePauseClass5MetricTag
	SetLength(value uint32) PatternFlowPfcPausePauseClass5MetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPausePauseClass5MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPausePauseClass5MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPausePauseClass5MetricTag object
func (obj *patternFlowPfcPausePauseClass5MetricTag) SetName(value string) PatternFlowPfcPausePauseClass5MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass5MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass5MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPausePauseClass5MetricTag object
func (obj *patternFlowPfcPausePauseClass5MetricTag) SetOffset(value uint32) PatternFlowPfcPausePauseClass5MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass5MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass5MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPausePauseClass5MetricTag object
func (obj *patternFlowPfcPausePauseClass5MetricTag) SetLength(value uint32) PatternFlowPfcPausePauseClass5MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass5MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPausePauseClass5MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5MetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPausePauseClass5MetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass5MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
