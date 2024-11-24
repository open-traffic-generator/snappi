package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseTimeMetricTag *****
type patternFlowEthernetPauseTimeMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetPauseTimeMetricTag
	marshaller   marshalPatternFlowEthernetPauseTimeMetricTag
	unMarshaller unMarshalPatternFlowEthernetPauseTimeMetricTag
}

func NewPatternFlowEthernetPauseTimeMetricTag() PatternFlowEthernetPauseTimeMetricTag {
	obj := patternFlowEthernetPauseTimeMetricTag{obj: &otg.PatternFlowEthernetPauseTimeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseTimeMetricTag) msg() *otg.PatternFlowEthernetPauseTimeMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetPauseTimeMetricTag) setMsg(msg *otg.PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseTimeMetricTag struct {
	obj *patternFlowEthernetPauseTimeMetricTag
}

type marshalPatternFlowEthernetPauseTimeMetricTag interface {
	// ToProto marshals PatternFlowEthernetPauseTimeMetricTag to protobuf object *otg.PatternFlowEthernetPauseTimeMetricTag
	ToProto() (*otg.PatternFlowEthernetPauseTimeMetricTag, error)
	// ToPbText marshals PatternFlowEthernetPauseTimeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseTimeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseTimeMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseTimeMetricTag struct {
	obj *patternFlowEthernetPauseTimeMetricTag
}

type unMarshalPatternFlowEthernetPauseTimeMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetPauseTimeMetricTag from protobuf object *otg.PatternFlowEthernetPauseTimeMetricTag
	FromProto(msg *otg.PatternFlowEthernetPauseTimeMetricTag) (PatternFlowEthernetPauseTimeMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetPauseTimeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseTimeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseTimeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseTimeMetricTag) Marshal() marshalPatternFlowEthernetPauseTimeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseTimeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseTimeMetricTag) Unmarshal() unMarshalPatternFlowEthernetPauseTimeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseTimeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseTimeMetricTag) ToProto() (*otg.PatternFlowEthernetPauseTimeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseTimeMetricTag) FromProto(msg *otg.PatternFlowEthernetPauseTimeMetricTag) (PatternFlowEthernetPauseTimeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseTimeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTimeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseTimeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTimeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseTimeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTimeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseTimeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseTimeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseTimeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseTimeMetricTag) Clone() (PatternFlowEthernetPauseTimeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseTimeMetricTag()
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

// PatternFlowEthernetPauseTimeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetPauseTimeMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetPauseTimeMetricTag to protobuf object *otg.PatternFlowEthernetPauseTimeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseTimeMetricTag
	// setMsg unmarshals PatternFlowEthernetPauseTimeMetricTag from protobuf object *otg.PatternFlowEthernetPauseTimeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseTimeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseTimeMetricTag
	// validate validates PatternFlowEthernetPauseTimeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseTimeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetPauseTimeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetPauseTimeMetricTag
	SetName(value string) PatternFlowEthernetPauseTimeMetricTag
	// Offset returns uint32, set in PatternFlowEthernetPauseTimeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetPauseTimeMetricTag
	SetOffset(value uint32) PatternFlowEthernetPauseTimeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetPauseTimeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetPauseTimeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetPauseTimeMetricTag
	SetLength(value uint32) PatternFlowEthernetPauseTimeMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetPauseTimeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetPauseTimeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetPauseTimeMetricTag object
func (obj *patternFlowEthernetPauseTimeMetricTag) SetName(value string) PatternFlowEthernetPauseTimeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseTimeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseTimeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetPauseTimeMetricTag object
func (obj *patternFlowEthernetPauseTimeMetricTag) SetOffset(value uint32) PatternFlowEthernetPauseTimeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseTimeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseTimeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetPauseTimeMetricTag object
func (obj *patternFlowEthernetPauseTimeMetricTag) SetLength(value uint32) PatternFlowEthernetPauseTimeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetPauseTimeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetPauseTimeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseTimeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetPauseTimeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetPauseTimeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
