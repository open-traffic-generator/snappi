package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseControlOpCodeMetricTag *****
type patternFlowPfcPauseControlOpCodeMetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPauseControlOpCodeMetricTag
	marshaller   marshalPatternFlowPfcPauseControlOpCodeMetricTag
	unMarshaller unMarshalPatternFlowPfcPauseControlOpCodeMetricTag
}

func NewPatternFlowPfcPauseControlOpCodeMetricTag() PatternFlowPfcPauseControlOpCodeMetricTag {
	obj := patternFlowPfcPauseControlOpCodeMetricTag{obj: &otg.PatternFlowPfcPauseControlOpCodeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) msg() *otg.PatternFlowPfcPauseControlOpCodeMetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) setMsg(msg *otg.PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseControlOpCodeMetricTag struct {
	obj *patternFlowPfcPauseControlOpCodeMetricTag
}

type marshalPatternFlowPfcPauseControlOpCodeMetricTag interface {
	// ToProto marshals PatternFlowPfcPauseControlOpCodeMetricTag to protobuf object *otg.PatternFlowPfcPauseControlOpCodeMetricTag
	ToProto() (*otg.PatternFlowPfcPauseControlOpCodeMetricTag, error)
	// ToPbText marshals PatternFlowPfcPauseControlOpCodeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseControlOpCodeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseControlOpCodeMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseControlOpCodeMetricTag struct {
	obj *patternFlowPfcPauseControlOpCodeMetricTag
}

type unMarshalPatternFlowPfcPauseControlOpCodeMetricTag interface {
	// FromProto unmarshals PatternFlowPfcPauseControlOpCodeMetricTag from protobuf object *otg.PatternFlowPfcPauseControlOpCodeMetricTag
	FromProto(msg *otg.PatternFlowPfcPauseControlOpCodeMetricTag) (PatternFlowPfcPauseControlOpCodeMetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPauseControlOpCodeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseControlOpCodeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseControlOpCodeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) Marshal() marshalPatternFlowPfcPauseControlOpCodeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseControlOpCodeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) Unmarshal() unMarshalPatternFlowPfcPauseControlOpCodeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseControlOpCodeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseControlOpCodeMetricTag) ToProto() (*otg.PatternFlowPfcPauseControlOpCodeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseControlOpCodeMetricTag) FromProto(msg *otg.PatternFlowPfcPauseControlOpCodeMetricTag) (PatternFlowPfcPauseControlOpCodeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseControlOpCodeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCodeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseControlOpCodeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCodeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseControlOpCodeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCodeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) Clone() (PatternFlowPfcPauseControlOpCodeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseControlOpCodeMetricTag()
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

// PatternFlowPfcPauseControlOpCodeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPauseControlOpCodeMetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPauseControlOpCodeMetricTag to protobuf object *otg.PatternFlowPfcPauseControlOpCodeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseControlOpCodeMetricTag
	// setMsg unmarshals PatternFlowPfcPauseControlOpCodeMetricTag from protobuf object *otg.PatternFlowPfcPauseControlOpCodeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseControlOpCodeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseControlOpCodeMetricTag
	// validate validates PatternFlowPfcPauseControlOpCodeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseControlOpCodeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPauseControlOpCodeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPauseControlOpCodeMetricTag
	SetName(value string) PatternFlowPfcPauseControlOpCodeMetricTag
	// Offset returns uint32, set in PatternFlowPfcPauseControlOpCodeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPauseControlOpCodeMetricTag
	SetOffset(value uint32) PatternFlowPfcPauseControlOpCodeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPauseControlOpCodeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPauseControlOpCodeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPauseControlOpCodeMetricTag
	SetLength(value uint32) PatternFlowPfcPauseControlOpCodeMetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPauseControlOpCodeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPauseControlOpCodeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPauseControlOpCodeMetricTag object
func (obj *patternFlowPfcPauseControlOpCodeMetricTag) SetName(value string) PatternFlowPfcPauseControlOpCodeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPauseControlOpCodeMetricTag object
func (obj *patternFlowPfcPauseControlOpCodeMetricTag) SetOffset(value uint32) PatternFlowPfcPauseControlOpCodeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPauseControlOpCodeMetricTag object
func (obj *patternFlowPfcPauseControlOpCodeMetricTag) SetLength(value uint32) PatternFlowPfcPauseControlOpCodeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPauseControlOpCodeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCodeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPauseControlOpCodeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPauseControlOpCodeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
