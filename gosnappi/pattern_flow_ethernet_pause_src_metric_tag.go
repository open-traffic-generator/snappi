package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseSrcMetricTag *****
type patternFlowEthernetPauseSrcMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetPauseSrcMetricTag
	marshaller   marshalPatternFlowEthernetPauseSrcMetricTag
	unMarshaller unMarshalPatternFlowEthernetPauseSrcMetricTag
}

func NewPatternFlowEthernetPauseSrcMetricTag() PatternFlowEthernetPauseSrcMetricTag {
	obj := patternFlowEthernetPauseSrcMetricTag{obj: &otg.PatternFlowEthernetPauseSrcMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseSrcMetricTag) msg() *otg.PatternFlowEthernetPauseSrcMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetPauseSrcMetricTag) setMsg(msg *otg.PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseSrcMetricTag struct {
	obj *patternFlowEthernetPauseSrcMetricTag
}

type marshalPatternFlowEthernetPauseSrcMetricTag interface {
	// ToProto marshals PatternFlowEthernetPauseSrcMetricTag to protobuf object *otg.PatternFlowEthernetPauseSrcMetricTag
	ToProto() (*otg.PatternFlowEthernetPauseSrcMetricTag, error)
	// ToPbText marshals PatternFlowEthernetPauseSrcMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseSrcMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseSrcMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseSrcMetricTag struct {
	obj *patternFlowEthernetPauseSrcMetricTag
}

type unMarshalPatternFlowEthernetPauseSrcMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetPauseSrcMetricTag from protobuf object *otg.PatternFlowEthernetPauseSrcMetricTag
	FromProto(msg *otg.PatternFlowEthernetPauseSrcMetricTag) (PatternFlowEthernetPauseSrcMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetPauseSrcMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseSrcMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseSrcMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseSrcMetricTag) Marshal() marshalPatternFlowEthernetPauseSrcMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseSrcMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseSrcMetricTag) Unmarshal() unMarshalPatternFlowEthernetPauseSrcMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseSrcMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseSrcMetricTag) ToProto() (*otg.PatternFlowEthernetPauseSrcMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseSrcMetricTag) FromProto(msg *otg.PatternFlowEthernetPauseSrcMetricTag) (PatternFlowEthernetPauseSrcMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseSrcMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrcMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseSrcMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrcMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseSrcMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrcMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseSrcMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseSrcMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseSrcMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseSrcMetricTag) Clone() (PatternFlowEthernetPauseSrcMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseSrcMetricTag()
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

// PatternFlowEthernetPauseSrcMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetPauseSrcMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetPauseSrcMetricTag to protobuf object *otg.PatternFlowEthernetPauseSrcMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseSrcMetricTag
	// setMsg unmarshals PatternFlowEthernetPauseSrcMetricTag from protobuf object *otg.PatternFlowEthernetPauseSrcMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseSrcMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseSrcMetricTag
	// validate validates PatternFlowEthernetPauseSrcMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseSrcMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetPauseSrcMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetPauseSrcMetricTag
	SetName(value string) PatternFlowEthernetPauseSrcMetricTag
	// Offset returns uint32, set in PatternFlowEthernetPauseSrcMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetPauseSrcMetricTag
	SetOffset(value uint32) PatternFlowEthernetPauseSrcMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetPauseSrcMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetPauseSrcMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetPauseSrcMetricTag
	SetLength(value uint32) PatternFlowEthernetPauseSrcMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetPauseSrcMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetPauseSrcMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetPauseSrcMetricTag object
func (obj *patternFlowEthernetPauseSrcMetricTag) SetName(value string) PatternFlowEthernetPauseSrcMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseSrcMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseSrcMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetPauseSrcMetricTag object
func (obj *patternFlowEthernetPauseSrcMetricTag) SetOffset(value uint32) PatternFlowEthernetPauseSrcMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseSrcMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseSrcMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetPauseSrcMetricTag object
func (obj *patternFlowEthernetPauseSrcMetricTag) SetLength(value uint32) PatternFlowEthernetPauseSrcMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetPauseSrcMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetPauseSrcMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseSrcMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetPauseSrcMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetPauseSrcMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}