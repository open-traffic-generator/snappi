package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseDstMetricTag *****
type patternFlowPfcPauseDstMetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPauseDstMetricTag
	marshaller   marshalPatternFlowPfcPauseDstMetricTag
	unMarshaller unMarshalPatternFlowPfcPauseDstMetricTag
}

func NewPatternFlowPfcPauseDstMetricTag() PatternFlowPfcPauseDstMetricTag {
	obj := patternFlowPfcPauseDstMetricTag{obj: &otg.PatternFlowPfcPauseDstMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseDstMetricTag) msg() *otg.PatternFlowPfcPauseDstMetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPauseDstMetricTag) setMsg(msg *otg.PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseDstMetricTag struct {
	obj *patternFlowPfcPauseDstMetricTag
}

type marshalPatternFlowPfcPauseDstMetricTag interface {
	// ToProto marshals PatternFlowPfcPauseDstMetricTag to protobuf object *otg.PatternFlowPfcPauseDstMetricTag
	ToProto() (*otg.PatternFlowPfcPauseDstMetricTag, error)
	// ToPbText marshals PatternFlowPfcPauseDstMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseDstMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseDstMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseDstMetricTag struct {
	obj *patternFlowPfcPauseDstMetricTag
}

type unMarshalPatternFlowPfcPauseDstMetricTag interface {
	// FromProto unmarshals PatternFlowPfcPauseDstMetricTag from protobuf object *otg.PatternFlowPfcPauseDstMetricTag
	FromProto(msg *otg.PatternFlowPfcPauseDstMetricTag) (PatternFlowPfcPauseDstMetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPauseDstMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseDstMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseDstMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseDstMetricTag) Marshal() marshalPatternFlowPfcPauseDstMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseDstMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseDstMetricTag) Unmarshal() unMarshalPatternFlowPfcPauseDstMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseDstMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseDstMetricTag) ToProto() (*otg.PatternFlowPfcPauseDstMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseDstMetricTag) FromProto(msg *otg.PatternFlowPfcPauseDstMetricTag) (PatternFlowPfcPauseDstMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseDstMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDstMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseDstMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDstMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseDstMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDstMetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseDstMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseDstMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseDstMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseDstMetricTag) Clone() (PatternFlowPfcPauseDstMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseDstMetricTag()
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

// PatternFlowPfcPauseDstMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPauseDstMetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPauseDstMetricTag to protobuf object *otg.PatternFlowPfcPauseDstMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseDstMetricTag
	// setMsg unmarshals PatternFlowPfcPauseDstMetricTag from protobuf object *otg.PatternFlowPfcPauseDstMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseDstMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseDstMetricTag
	// validate validates PatternFlowPfcPauseDstMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseDstMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPauseDstMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPauseDstMetricTag
	SetName(value string) PatternFlowPfcPauseDstMetricTag
	// Offset returns uint32, set in PatternFlowPfcPauseDstMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPauseDstMetricTag
	SetOffset(value uint32) PatternFlowPfcPauseDstMetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPauseDstMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPauseDstMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPauseDstMetricTag
	SetLength(value uint32) PatternFlowPfcPauseDstMetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPauseDstMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPauseDstMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPauseDstMetricTag object
func (obj *patternFlowPfcPauseDstMetricTag) SetName(value string) PatternFlowPfcPauseDstMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseDstMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseDstMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPauseDstMetricTag object
func (obj *patternFlowPfcPauseDstMetricTag) SetOffset(value uint32) PatternFlowPfcPauseDstMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseDstMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseDstMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPauseDstMetricTag object
func (obj *patternFlowPfcPauseDstMetricTag) SetLength(value uint32) PatternFlowPfcPauseDstMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPauseDstMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPauseDstMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseDstMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPauseDstMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPauseDstMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}
