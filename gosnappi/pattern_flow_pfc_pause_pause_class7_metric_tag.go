package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass7MetricTag *****
type patternFlowPfcPausePauseClass7MetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass7MetricTag
	marshaller   marshalPatternFlowPfcPausePauseClass7MetricTag
	unMarshaller unMarshalPatternFlowPfcPausePauseClass7MetricTag
}

func NewPatternFlowPfcPausePauseClass7MetricTag() PatternFlowPfcPausePauseClass7MetricTag {
	obj := patternFlowPfcPausePauseClass7MetricTag{obj: &otg.PatternFlowPfcPausePauseClass7MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass7MetricTag) msg() *otg.PatternFlowPfcPausePauseClass7MetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass7MetricTag) setMsg(msg *otg.PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass7MetricTag struct {
	obj *patternFlowPfcPausePauseClass7MetricTag
}

type marshalPatternFlowPfcPausePauseClass7MetricTag interface {
	// ToProto marshals PatternFlowPfcPausePauseClass7MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass7MetricTag
	ToProto() (*otg.PatternFlowPfcPausePauseClass7MetricTag, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass7MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass7MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass7MetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPausePauseClass7MetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass7MetricTag struct {
	obj *patternFlowPfcPausePauseClass7MetricTag
}

type unMarshalPatternFlowPfcPausePauseClass7MetricTag interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass7MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass7MetricTag
	FromProto(msg *otg.PatternFlowPfcPausePauseClass7MetricTag) (PatternFlowPfcPausePauseClass7MetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass7MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass7MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass7MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass7MetricTag) Marshal() marshalPatternFlowPfcPausePauseClass7MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass7MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass7MetricTag) Unmarshal() unMarshalPatternFlowPfcPausePauseClass7MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass7MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass7MetricTag) ToProto() (*otg.PatternFlowPfcPausePauseClass7MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass7MetricTag) FromProto(msg *otg.PatternFlowPfcPausePauseClass7MetricTag) (PatternFlowPfcPausePauseClass7MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass7MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass7MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass7MetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPausePauseClass7MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7MetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass7MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass7MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass7MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass7MetricTag) Clone() (PatternFlowPfcPausePauseClass7MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass7MetricTag()
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

// PatternFlowPfcPausePauseClass7MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPausePauseClass7MetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass7MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass7MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass7MetricTag
	// setMsg unmarshals PatternFlowPfcPausePauseClass7MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass7MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass7MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass7MetricTag
	// validate validates PatternFlowPfcPausePauseClass7MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass7MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPausePauseClass7MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPausePauseClass7MetricTag
	SetName(value string) PatternFlowPfcPausePauseClass7MetricTag
	// Offset returns uint32, set in PatternFlowPfcPausePauseClass7MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPausePauseClass7MetricTag
	SetOffset(value uint32) PatternFlowPfcPausePauseClass7MetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPausePauseClass7MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPausePauseClass7MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPausePauseClass7MetricTag
	SetLength(value uint32) PatternFlowPfcPausePauseClass7MetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPausePauseClass7MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPausePauseClass7MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPausePauseClass7MetricTag object
func (obj *patternFlowPfcPausePauseClass7MetricTag) SetName(value string) PatternFlowPfcPausePauseClass7MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass7MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass7MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPausePauseClass7MetricTag object
func (obj *patternFlowPfcPausePauseClass7MetricTag) SetOffset(value uint32) PatternFlowPfcPausePauseClass7MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass7MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass7MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPausePauseClass7MetricTag object
func (obj *patternFlowPfcPausePauseClass7MetricTag) SetLength(value uint32) PatternFlowPfcPausePauseClass7MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass7MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPausePauseClass7MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7MetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPausePauseClass7MetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass7MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
