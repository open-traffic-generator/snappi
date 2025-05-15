package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseEtherTypeMetricTag *****
type patternFlowEthernetPauseEtherTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetPauseEtherTypeMetricTag
	marshaller   marshalPatternFlowEthernetPauseEtherTypeMetricTag
	unMarshaller unMarshalPatternFlowEthernetPauseEtherTypeMetricTag
}

func NewPatternFlowEthernetPauseEtherTypeMetricTag() PatternFlowEthernetPauseEtherTypeMetricTag {
	obj := patternFlowEthernetPauseEtherTypeMetricTag{obj: &otg.PatternFlowEthernetPauseEtherTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) msg() *otg.PatternFlowEthernetPauseEtherTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) setMsg(msg *otg.PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseEtherTypeMetricTag struct {
	obj *patternFlowEthernetPauseEtherTypeMetricTag
}

type marshalPatternFlowEthernetPauseEtherTypeMetricTag interface {
	// ToProto marshals PatternFlowEthernetPauseEtherTypeMetricTag to protobuf object *otg.PatternFlowEthernetPauseEtherTypeMetricTag
	ToProto() (*otg.PatternFlowEthernetPauseEtherTypeMetricTag, error)
	// ToPbText marshals PatternFlowEthernetPauseEtherTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseEtherTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseEtherTypeMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseEtherTypeMetricTag struct {
	obj *patternFlowEthernetPauseEtherTypeMetricTag
}

type unMarshalPatternFlowEthernetPauseEtherTypeMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetPauseEtherTypeMetricTag from protobuf object *otg.PatternFlowEthernetPauseEtherTypeMetricTag
	FromProto(msg *otg.PatternFlowEthernetPauseEtherTypeMetricTag) (PatternFlowEthernetPauseEtherTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetPauseEtherTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseEtherTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseEtherTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) Marshal() marshalPatternFlowEthernetPauseEtherTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseEtherTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) Unmarshal() unMarshalPatternFlowEthernetPauseEtherTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseEtherTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseEtherTypeMetricTag) ToProto() (*otg.PatternFlowEthernetPauseEtherTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseEtherTypeMetricTag) FromProto(msg *otg.PatternFlowEthernetPauseEtherTypeMetricTag) (PatternFlowEthernetPauseEtherTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseEtherTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseEtherTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseEtherTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) Clone() (PatternFlowEthernetPauseEtherTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseEtherTypeMetricTag()
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

// PatternFlowEthernetPauseEtherTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetPauseEtherTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetPauseEtherTypeMetricTag to protobuf object *otg.PatternFlowEthernetPauseEtherTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseEtherTypeMetricTag
	// setMsg unmarshals PatternFlowEthernetPauseEtherTypeMetricTag from protobuf object *otg.PatternFlowEthernetPauseEtherTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseEtherTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseEtherTypeMetricTag
	// validate validates PatternFlowEthernetPauseEtherTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseEtherTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetPauseEtherTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetPauseEtherTypeMetricTag
	SetName(value string) PatternFlowEthernetPauseEtherTypeMetricTag
	// Offset returns uint32, set in PatternFlowEthernetPauseEtherTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetPauseEtherTypeMetricTag
	SetOffset(value uint32) PatternFlowEthernetPauseEtherTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetPauseEtherTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetPauseEtherTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetPauseEtherTypeMetricTag
	SetLength(value uint32) PatternFlowEthernetPauseEtherTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetPauseEtherTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetPauseEtherTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetPauseEtherTypeMetricTag object
func (obj *patternFlowEthernetPauseEtherTypeMetricTag) SetName(value string) PatternFlowEthernetPauseEtherTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetPauseEtherTypeMetricTag object
func (obj *patternFlowEthernetPauseEtherTypeMetricTag) SetOffset(value uint32) PatternFlowEthernetPauseEtherTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetPauseEtherTypeMetricTag object
func (obj *patternFlowEthernetPauseEtherTypeMetricTag) SetLength(value uint32) PatternFlowEthernetPauseEtherTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetPauseEtherTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherTypeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetPauseEtherTypeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetPauseEtherTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
