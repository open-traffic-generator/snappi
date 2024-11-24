package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseSrcMetricTag *****
type patternFlowPfcPauseSrcMetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPauseSrcMetricTag
	marshaller   marshalPatternFlowPfcPauseSrcMetricTag
	unMarshaller unMarshalPatternFlowPfcPauseSrcMetricTag
}

func NewPatternFlowPfcPauseSrcMetricTag() PatternFlowPfcPauseSrcMetricTag {
	obj := patternFlowPfcPauseSrcMetricTag{obj: &otg.PatternFlowPfcPauseSrcMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseSrcMetricTag) msg() *otg.PatternFlowPfcPauseSrcMetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPauseSrcMetricTag) setMsg(msg *otg.PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseSrcMetricTag struct {
	obj *patternFlowPfcPauseSrcMetricTag
}

type marshalPatternFlowPfcPauseSrcMetricTag interface {
	// ToProto marshals PatternFlowPfcPauseSrcMetricTag to protobuf object *otg.PatternFlowPfcPauseSrcMetricTag
	ToProto() (*otg.PatternFlowPfcPauseSrcMetricTag, error)
	// ToPbText marshals PatternFlowPfcPauseSrcMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseSrcMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseSrcMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPauseSrcMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPauseSrcMetricTag struct {
	obj *patternFlowPfcPauseSrcMetricTag
}

type unMarshalPatternFlowPfcPauseSrcMetricTag interface {
	// FromProto unmarshals PatternFlowPfcPauseSrcMetricTag from protobuf object *otg.PatternFlowPfcPauseSrcMetricTag
	FromProto(msg *otg.PatternFlowPfcPauseSrcMetricTag) (PatternFlowPfcPauseSrcMetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPauseSrcMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseSrcMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseSrcMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseSrcMetricTag) Marshal() marshalPatternFlowPfcPauseSrcMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseSrcMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseSrcMetricTag) Unmarshal() unMarshalPatternFlowPfcPauseSrcMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseSrcMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseSrcMetricTag) ToProto() (*otg.PatternFlowPfcPauseSrcMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseSrcMetricTag) FromProto(msg *otg.PatternFlowPfcPauseSrcMetricTag) (PatternFlowPfcPauseSrcMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseSrcMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrcMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseSrcMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrcMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseSrcMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPauseSrcMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrcMetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseSrcMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseSrcMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseSrcMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseSrcMetricTag) Clone() (PatternFlowPfcPauseSrcMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseSrcMetricTag()
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

// PatternFlowPfcPauseSrcMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPauseSrcMetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPauseSrcMetricTag to protobuf object *otg.PatternFlowPfcPauseSrcMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseSrcMetricTag
	// setMsg unmarshals PatternFlowPfcPauseSrcMetricTag from protobuf object *otg.PatternFlowPfcPauseSrcMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseSrcMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseSrcMetricTag
	// validate validates PatternFlowPfcPauseSrcMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseSrcMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPauseSrcMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPauseSrcMetricTag
	SetName(value string) PatternFlowPfcPauseSrcMetricTag
	// Offset returns uint32, set in PatternFlowPfcPauseSrcMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPauseSrcMetricTag
	SetOffset(value uint32) PatternFlowPfcPauseSrcMetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPauseSrcMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPauseSrcMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPauseSrcMetricTag
	SetLength(value uint32) PatternFlowPfcPauseSrcMetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPauseSrcMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPauseSrcMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPauseSrcMetricTag object
func (obj *patternFlowPfcPauseSrcMetricTag) SetName(value string) PatternFlowPfcPauseSrcMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseSrcMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseSrcMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPauseSrcMetricTag object
func (obj *patternFlowPfcPauseSrcMetricTag) SetOffset(value uint32) PatternFlowPfcPauseSrcMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseSrcMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseSrcMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPauseSrcMetricTag object
func (obj *patternFlowPfcPauseSrcMetricTag) SetLength(value uint32) PatternFlowPfcPauseSrcMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPauseSrcMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPauseSrcMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseSrcMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPauseSrcMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPauseSrcMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}
