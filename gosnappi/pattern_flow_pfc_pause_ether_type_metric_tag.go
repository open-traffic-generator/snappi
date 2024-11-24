package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseEtherTypeMetricTag *****
type patternFlowPfcPauseEtherTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPauseEtherTypeMetricTag
	marshaller   marshalPatternFlowPfcPauseEtherTypeMetricTag
	unMarshaller unMarshalPatternFlowPfcPauseEtherTypeMetricTag
}

func NewPatternFlowPfcPauseEtherTypeMetricTag() PatternFlowPfcPauseEtherTypeMetricTag {
	obj := patternFlowPfcPauseEtherTypeMetricTag{obj: &otg.PatternFlowPfcPauseEtherTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) msg() *otg.PatternFlowPfcPauseEtherTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) setMsg(msg *otg.PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseEtherTypeMetricTag struct {
	obj *patternFlowPfcPauseEtherTypeMetricTag
}

type marshalPatternFlowPfcPauseEtherTypeMetricTag interface {
	// ToProto marshals PatternFlowPfcPauseEtherTypeMetricTag to protobuf object *otg.PatternFlowPfcPauseEtherTypeMetricTag
	ToProto() (*otg.PatternFlowPfcPauseEtherTypeMetricTag, error)
	// ToPbText marshals PatternFlowPfcPauseEtherTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseEtherTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseEtherTypeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPauseEtherTypeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPauseEtherTypeMetricTag struct {
	obj *patternFlowPfcPauseEtherTypeMetricTag
}

type unMarshalPatternFlowPfcPauseEtherTypeMetricTag interface {
	// FromProto unmarshals PatternFlowPfcPauseEtherTypeMetricTag from protobuf object *otg.PatternFlowPfcPauseEtherTypeMetricTag
	FromProto(msg *otg.PatternFlowPfcPauseEtherTypeMetricTag) (PatternFlowPfcPauseEtherTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPauseEtherTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseEtherTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseEtherTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) Marshal() marshalPatternFlowPfcPauseEtherTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseEtherTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) Unmarshal() unMarshalPatternFlowPfcPauseEtherTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseEtherTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseEtherTypeMetricTag) ToProto() (*otg.PatternFlowPfcPauseEtherTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseEtherTypeMetricTag) FromProto(msg *otg.PatternFlowPfcPauseEtherTypeMetricTag) (PatternFlowPfcPauseEtherTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseEtherTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseEtherTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseEtherTypeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPauseEtherTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseEtherTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) Clone() (PatternFlowPfcPauseEtherTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseEtherTypeMetricTag()
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

// PatternFlowPfcPauseEtherTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPauseEtherTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPauseEtherTypeMetricTag to protobuf object *otg.PatternFlowPfcPauseEtherTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseEtherTypeMetricTag
	// setMsg unmarshals PatternFlowPfcPauseEtherTypeMetricTag from protobuf object *otg.PatternFlowPfcPauseEtherTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseEtherTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseEtherTypeMetricTag
	// validate validates PatternFlowPfcPauseEtherTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseEtherTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPauseEtherTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPauseEtherTypeMetricTag
	SetName(value string) PatternFlowPfcPauseEtherTypeMetricTag
	// Offset returns uint32, set in PatternFlowPfcPauseEtherTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPauseEtherTypeMetricTag
	SetOffset(value uint32) PatternFlowPfcPauseEtherTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPauseEtherTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPauseEtherTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPauseEtherTypeMetricTag
	SetLength(value uint32) PatternFlowPfcPauseEtherTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPauseEtherTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPauseEtherTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPauseEtherTypeMetricTag object
func (obj *patternFlowPfcPauseEtherTypeMetricTag) SetName(value string) PatternFlowPfcPauseEtherTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseEtherTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseEtherTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPauseEtherTypeMetricTag object
func (obj *patternFlowPfcPauseEtherTypeMetricTag) SetOffset(value uint32) PatternFlowPfcPauseEtherTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseEtherTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseEtherTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPauseEtherTypeMetricTag object
func (obj *patternFlowPfcPauseEtherTypeMetricTag) SetLength(value uint32) PatternFlowPfcPauseEtherTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPauseEtherTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseEtherTypeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPauseEtherTypeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPauseEtherTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
