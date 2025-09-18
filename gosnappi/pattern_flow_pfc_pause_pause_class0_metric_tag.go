package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass0MetricTag *****
type patternFlowPfcPausePauseClass0MetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass0MetricTag
	marshaller   marshalPatternFlowPfcPausePauseClass0MetricTag
	unMarshaller unMarshalPatternFlowPfcPausePauseClass0MetricTag
}

func NewPatternFlowPfcPausePauseClass0MetricTag() PatternFlowPfcPausePauseClass0MetricTag {
	obj := patternFlowPfcPausePauseClass0MetricTag{obj: &otg.PatternFlowPfcPausePauseClass0MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass0MetricTag) msg() *otg.PatternFlowPfcPausePauseClass0MetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass0MetricTag) setMsg(msg *otg.PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass0MetricTag struct {
	obj *patternFlowPfcPausePauseClass0MetricTag
}

type marshalPatternFlowPfcPausePauseClass0MetricTag interface {
	// ToProto marshals PatternFlowPfcPausePauseClass0MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass0MetricTag
	ToProto() (*otg.PatternFlowPfcPausePauseClass0MetricTag, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass0MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass0MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass0MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass0MetricTag struct {
	obj *patternFlowPfcPausePauseClass0MetricTag
}

type unMarshalPatternFlowPfcPausePauseClass0MetricTag interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass0MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass0MetricTag
	FromProto(msg *otg.PatternFlowPfcPausePauseClass0MetricTag) (PatternFlowPfcPausePauseClass0MetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass0MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass0MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass0MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass0MetricTag) Marshal() marshalPatternFlowPfcPausePauseClass0MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass0MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass0MetricTag) Unmarshal() unMarshalPatternFlowPfcPausePauseClass0MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass0MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass0MetricTag) ToProto() (*otg.PatternFlowPfcPausePauseClass0MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass0MetricTag) FromProto(msg *otg.PatternFlowPfcPausePauseClass0MetricTag) (PatternFlowPfcPausePauseClass0MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass0MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass0MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass0MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0MetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass0MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass0MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass0MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass0MetricTag) Clone() (PatternFlowPfcPausePauseClass0MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass0MetricTag()
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

// PatternFlowPfcPausePauseClass0MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPausePauseClass0MetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass0MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass0MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass0MetricTag
	// setMsg unmarshals PatternFlowPfcPausePauseClass0MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass0MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass0MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass0MetricTag
	// validate validates PatternFlowPfcPausePauseClass0MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass0MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPausePauseClass0MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPausePauseClass0MetricTag
	SetName(value string) PatternFlowPfcPausePauseClass0MetricTag
	// Offset returns uint32, set in PatternFlowPfcPausePauseClass0MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPausePauseClass0MetricTag
	SetOffset(value uint32) PatternFlowPfcPausePauseClass0MetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPausePauseClass0MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPausePauseClass0MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPausePauseClass0MetricTag
	SetLength(value uint32) PatternFlowPfcPausePauseClass0MetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPausePauseClass0MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPausePauseClass0MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPausePauseClass0MetricTag object
func (obj *patternFlowPfcPausePauseClass0MetricTag) SetName(value string) PatternFlowPfcPausePauseClass0MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass0MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass0MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPausePauseClass0MetricTag object
func (obj *patternFlowPfcPausePauseClass0MetricTag) SetOffset(value uint32) PatternFlowPfcPausePauseClass0MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass0MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass0MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPausePauseClass0MetricTag object
func (obj *patternFlowPfcPausePauseClass0MetricTag) SetLength(value uint32) PatternFlowPfcPausePauseClass0MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass0MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPausePauseClass0MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0MetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPausePauseClass0MetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass0MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
