package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseDstMetricTag *****
type patternFlowEthernetPauseDstMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetPauseDstMetricTag
	marshaller   marshalPatternFlowEthernetPauseDstMetricTag
	unMarshaller unMarshalPatternFlowEthernetPauseDstMetricTag
}

func NewPatternFlowEthernetPauseDstMetricTag() PatternFlowEthernetPauseDstMetricTag {
	obj := patternFlowEthernetPauseDstMetricTag{obj: &otg.PatternFlowEthernetPauseDstMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseDstMetricTag) msg() *otg.PatternFlowEthernetPauseDstMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetPauseDstMetricTag) setMsg(msg *otg.PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseDstMetricTag struct {
	obj *patternFlowEthernetPauseDstMetricTag
}

type marshalPatternFlowEthernetPauseDstMetricTag interface {
	// ToProto marshals PatternFlowEthernetPauseDstMetricTag to protobuf object *otg.PatternFlowEthernetPauseDstMetricTag
	ToProto() (*otg.PatternFlowEthernetPauseDstMetricTag, error)
	// ToPbText marshals PatternFlowEthernetPauseDstMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseDstMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseDstMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseDstMetricTag struct {
	obj *patternFlowEthernetPauseDstMetricTag
}

type unMarshalPatternFlowEthernetPauseDstMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetPauseDstMetricTag from protobuf object *otg.PatternFlowEthernetPauseDstMetricTag
	FromProto(msg *otg.PatternFlowEthernetPauseDstMetricTag) (PatternFlowEthernetPauseDstMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetPauseDstMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseDstMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseDstMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseDstMetricTag) Marshal() marshalPatternFlowEthernetPauseDstMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseDstMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseDstMetricTag) Unmarshal() unMarshalPatternFlowEthernetPauseDstMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseDstMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseDstMetricTag) ToProto() (*otg.PatternFlowEthernetPauseDstMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseDstMetricTag) FromProto(msg *otg.PatternFlowEthernetPauseDstMetricTag) (PatternFlowEthernetPauseDstMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseDstMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDstMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseDstMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDstMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseDstMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDstMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseDstMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseDstMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseDstMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseDstMetricTag) Clone() (PatternFlowEthernetPauseDstMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseDstMetricTag()
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

// PatternFlowEthernetPauseDstMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetPauseDstMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetPauseDstMetricTag to protobuf object *otg.PatternFlowEthernetPauseDstMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseDstMetricTag
	// setMsg unmarshals PatternFlowEthernetPauseDstMetricTag from protobuf object *otg.PatternFlowEthernetPauseDstMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseDstMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseDstMetricTag
	// validate validates PatternFlowEthernetPauseDstMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseDstMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetPauseDstMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetPauseDstMetricTag
	SetName(value string) PatternFlowEthernetPauseDstMetricTag
	// Offset returns uint32, set in PatternFlowEthernetPauseDstMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetPauseDstMetricTag
	SetOffset(value uint32) PatternFlowEthernetPauseDstMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetPauseDstMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetPauseDstMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetPauseDstMetricTag
	SetLength(value uint32) PatternFlowEthernetPauseDstMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetPauseDstMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetPauseDstMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetPauseDstMetricTag object
func (obj *patternFlowEthernetPauseDstMetricTag) SetName(value string) PatternFlowEthernetPauseDstMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseDstMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseDstMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetPauseDstMetricTag object
func (obj *patternFlowEthernetPauseDstMetricTag) SetOffset(value uint32) PatternFlowEthernetPauseDstMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseDstMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseDstMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetPauseDstMetricTag object
func (obj *patternFlowEthernetPauseDstMetricTag) SetLength(value uint32) PatternFlowEthernetPauseDstMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetPauseDstMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetPauseDstMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseDstMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetPauseDstMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetPauseDstMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}
