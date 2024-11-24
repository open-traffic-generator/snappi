package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass2MetricTag *****
type patternFlowPfcPausePauseClass2MetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass2MetricTag
	marshaller   marshalPatternFlowPfcPausePauseClass2MetricTag
	unMarshaller unMarshalPatternFlowPfcPausePauseClass2MetricTag
}

func NewPatternFlowPfcPausePauseClass2MetricTag() PatternFlowPfcPausePauseClass2MetricTag {
	obj := patternFlowPfcPausePauseClass2MetricTag{obj: &otg.PatternFlowPfcPausePauseClass2MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass2MetricTag) msg() *otg.PatternFlowPfcPausePauseClass2MetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass2MetricTag) setMsg(msg *otg.PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass2MetricTag struct {
	obj *patternFlowPfcPausePauseClass2MetricTag
}

type marshalPatternFlowPfcPausePauseClass2MetricTag interface {
	// ToProto marshals PatternFlowPfcPausePauseClass2MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass2MetricTag
	ToProto() (*otg.PatternFlowPfcPausePauseClass2MetricTag, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass2MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass2MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass2MetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPausePauseClass2MetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass2MetricTag struct {
	obj *patternFlowPfcPausePauseClass2MetricTag
}

type unMarshalPatternFlowPfcPausePauseClass2MetricTag interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass2MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass2MetricTag
	FromProto(msg *otg.PatternFlowPfcPausePauseClass2MetricTag) (PatternFlowPfcPausePauseClass2MetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass2MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass2MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass2MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass2MetricTag) Marshal() marshalPatternFlowPfcPausePauseClass2MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass2MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass2MetricTag) Unmarshal() unMarshalPatternFlowPfcPausePauseClass2MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass2MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass2MetricTag) ToProto() (*otg.PatternFlowPfcPausePauseClass2MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass2MetricTag) FromProto(msg *otg.PatternFlowPfcPausePauseClass2MetricTag) (PatternFlowPfcPausePauseClass2MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass2MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass2MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass2MetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPausePauseClass2MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2MetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass2MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass2MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass2MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass2MetricTag) Clone() (PatternFlowPfcPausePauseClass2MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass2MetricTag()
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

// PatternFlowPfcPausePauseClass2MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPausePauseClass2MetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass2MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass2MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass2MetricTag
	// setMsg unmarshals PatternFlowPfcPausePauseClass2MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass2MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass2MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass2MetricTag
	// validate validates PatternFlowPfcPausePauseClass2MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass2MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPausePauseClass2MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPausePauseClass2MetricTag
	SetName(value string) PatternFlowPfcPausePauseClass2MetricTag
	// Offset returns uint32, set in PatternFlowPfcPausePauseClass2MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPausePauseClass2MetricTag
	SetOffset(value uint32) PatternFlowPfcPausePauseClass2MetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPausePauseClass2MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPausePauseClass2MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPausePauseClass2MetricTag
	SetLength(value uint32) PatternFlowPfcPausePauseClass2MetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPausePauseClass2MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPausePauseClass2MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPausePauseClass2MetricTag object
func (obj *patternFlowPfcPausePauseClass2MetricTag) SetName(value string) PatternFlowPfcPausePauseClass2MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass2MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass2MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPausePauseClass2MetricTag object
func (obj *patternFlowPfcPausePauseClass2MetricTag) SetOffset(value uint32) PatternFlowPfcPausePauseClass2MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass2MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass2MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPausePauseClass2MetricTag object
func (obj *patternFlowPfcPausePauseClass2MetricTag) SetLength(value uint32) PatternFlowPfcPausePauseClass2MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass2MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPausePauseClass2MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2MetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPausePauseClass2MetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass2MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
