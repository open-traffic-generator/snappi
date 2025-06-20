package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpSrcPortMetricTag *****
type patternFlowUdpSrcPortMetricTag struct {
	validation
	obj          *otg.PatternFlowUdpSrcPortMetricTag
	marshaller   marshalPatternFlowUdpSrcPortMetricTag
	unMarshaller unMarshalPatternFlowUdpSrcPortMetricTag
}

func NewPatternFlowUdpSrcPortMetricTag() PatternFlowUdpSrcPortMetricTag {
	obj := patternFlowUdpSrcPortMetricTag{obj: &otg.PatternFlowUdpSrcPortMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpSrcPortMetricTag) msg() *otg.PatternFlowUdpSrcPortMetricTag {
	return obj.obj
}

func (obj *patternFlowUdpSrcPortMetricTag) setMsg(msg *otg.PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpSrcPortMetricTag struct {
	obj *patternFlowUdpSrcPortMetricTag
}

type marshalPatternFlowUdpSrcPortMetricTag interface {
	// ToProto marshals PatternFlowUdpSrcPortMetricTag to protobuf object *otg.PatternFlowUdpSrcPortMetricTag
	ToProto() (*otg.PatternFlowUdpSrcPortMetricTag, error)
	// ToPbText marshals PatternFlowUdpSrcPortMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpSrcPortMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpSrcPortMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowUdpSrcPortMetricTag struct {
	obj *patternFlowUdpSrcPortMetricTag
}

type unMarshalPatternFlowUdpSrcPortMetricTag interface {
	// FromProto unmarshals PatternFlowUdpSrcPortMetricTag from protobuf object *otg.PatternFlowUdpSrcPortMetricTag
	FromProto(msg *otg.PatternFlowUdpSrcPortMetricTag) (PatternFlowUdpSrcPortMetricTag, error)
	// FromPbText unmarshals PatternFlowUdpSrcPortMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpSrcPortMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpSrcPortMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpSrcPortMetricTag) Marshal() marshalPatternFlowUdpSrcPortMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpSrcPortMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpSrcPortMetricTag) Unmarshal() unMarshalPatternFlowUdpSrcPortMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpSrcPortMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpSrcPortMetricTag) ToProto() (*otg.PatternFlowUdpSrcPortMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpSrcPortMetricTag) FromProto(msg *otg.PatternFlowUdpSrcPortMetricTag) (PatternFlowUdpSrcPortMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpSrcPortMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpSrcPortMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpSrcPortMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortMetricTag) FromJson(value string) error {
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

func (obj *patternFlowUdpSrcPortMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpSrcPortMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpSrcPortMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpSrcPortMetricTag) Clone() (PatternFlowUdpSrcPortMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpSrcPortMetricTag()
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

// PatternFlowUdpSrcPortMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowUdpSrcPortMetricTag interface {
	Validation
	// msg marshals PatternFlowUdpSrcPortMetricTag to protobuf object *otg.PatternFlowUdpSrcPortMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpSrcPortMetricTag
	// setMsg unmarshals PatternFlowUdpSrcPortMetricTag from protobuf object *otg.PatternFlowUdpSrcPortMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowUdpSrcPortMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpSrcPortMetricTag
	// validate validates PatternFlowUdpSrcPortMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpSrcPortMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowUdpSrcPortMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowUdpSrcPortMetricTag
	SetName(value string) PatternFlowUdpSrcPortMetricTag
	// Offset returns uint32, set in PatternFlowUdpSrcPortMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowUdpSrcPortMetricTag
	SetOffset(value uint32) PatternFlowUdpSrcPortMetricTag
	// HasOffset checks if Offset has been set in PatternFlowUdpSrcPortMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowUdpSrcPortMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowUdpSrcPortMetricTag
	SetLength(value uint32) PatternFlowUdpSrcPortMetricTag
	// HasLength checks if Length has been set in PatternFlowUdpSrcPortMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowUdpSrcPortMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowUdpSrcPortMetricTag object
func (obj *patternFlowUdpSrcPortMetricTag) SetName(value string) PatternFlowUdpSrcPortMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowUdpSrcPortMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowUdpSrcPortMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowUdpSrcPortMetricTag object
func (obj *patternFlowUdpSrcPortMetricTag) SetOffset(value uint32) PatternFlowUdpSrcPortMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowUdpSrcPortMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowUdpSrcPortMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowUdpSrcPortMetricTag object
func (obj *patternFlowUdpSrcPortMetricTag) SetLength(value uint32) PatternFlowUdpSrcPortMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowUdpSrcPortMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowUdpSrcPortMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPortMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowUdpSrcPortMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowUdpSrcPortMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
