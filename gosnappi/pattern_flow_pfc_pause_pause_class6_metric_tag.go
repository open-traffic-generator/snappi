package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass6MetricTag *****
type patternFlowPfcPausePauseClass6MetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass6MetricTag
	marshaller   marshalPatternFlowPfcPausePauseClass6MetricTag
	unMarshaller unMarshalPatternFlowPfcPausePauseClass6MetricTag
}

func NewPatternFlowPfcPausePauseClass6MetricTag() PatternFlowPfcPausePauseClass6MetricTag {
	obj := patternFlowPfcPausePauseClass6MetricTag{obj: &otg.PatternFlowPfcPausePauseClass6MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass6MetricTag) msg() *otg.PatternFlowPfcPausePauseClass6MetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass6MetricTag) setMsg(msg *otg.PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass6MetricTag struct {
	obj *patternFlowPfcPausePauseClass6MetricTag
}

type marshalPatternFlowPfcPausePauseClass6MetricTag interface {
	// ToProto marshals PatternFlowPfcPausePauseClass6MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass6MetricTag
	ToProto() (*otg.PatternFlowPfcPausePauseClass6MetricTag, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass6MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass6MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass6MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass6MetricTag struct {
	obj *patternFlowPfcPausePauseClass6MetricTag
}

type unMarshalPatternFlowPfcPausePauseClass6MetricTag interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass6MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass6MetricTag
	FromProto(msg *otg.PatternFlowPfcPausePauseClass6MetricTag) (PatternFlowPfcPausePauseClass6MetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass6MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass6MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass6MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass6MetricTag) Marshal() marshalPatternFlowPfcPausePauseClass6MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass6MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass6MetricTag) Unmarshal() unMarshalPatternFlowPfcPausePauseClass6MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass6MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass6MetricTag) ToProto() (*otg.PatternFlowPfcPausePauseClass6MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass6MetricTag) FromProto(msg *otg.PatternFlowPfcPausePauseClass6MetricTag) (PatternFlowPfcPausePauseClass6MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass6MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass6MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass6MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6MetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass6MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass6MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass6MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass6MetricTag) Clone() (PatternFlowPfcPausePauseClass6MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass6MetricTag()
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

// PatternFlowPfcPausePauseClass6MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPausePauseClass6MetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass6MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass6MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass6MetricTag
	// setMsg unmarshals PatternFlowPfcPausePauseClass6MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass6MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass6MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass6MetricTag
	// validate validates PatternFlowPfcPausePauseClass6MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass6MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPausePauseClass6MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPausePauseClass6MetricTag
	SetName(value string) PatternFlowPfcPausePauseClass6MetricTag
	// Offset returns uint32, set in PatternFlowPfcPausePauseClass6MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPausePauseClass6MetricTag
	SetOffset(value uint32) PatternFlowPfcPausePauseClass6MetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPausePauseClass6MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPausePauseClass6MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPausePauseClass6MetricTag
	SetLength(value uint32) PatternFlowPfcPausePauseClass6MetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPausePauseClass6MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPausePauseClass6MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPausePauseClass6MetricTag object
func (obj *patternFlowPfcPausePauseClass6MetricTag) SetName(value string) PatternFlowPfcPausePauseClass6MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass6MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass6MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPausePauseClass6MetricTag object
func (obj *patternFlowPfcPausePauseClass6MetricTag) SetOffset(value uint32) PatternFlowPfcPausePauseClass6MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass6MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass6MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPausePauseClass6MetricTag object
func (obj *patternFlowPfcPausePauseClass6MetricTag) SetLength(value uint32) PatternFlowPfcPausePauseClass6MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass6MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPausePauseClass6MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6MetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPausePauseClass6MetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass6MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}