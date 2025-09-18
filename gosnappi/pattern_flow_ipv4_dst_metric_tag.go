package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DstMetricTag *****
type patternFlowIpv4DstMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4DstMetricTag
	marshaller   marshalPatternFlowIpv4DstMetricTag
	unMarshaller unMarshalPatternFlowIpv4DstMetricTag
}

func NewPatternFlowIpv4DstMetricTag() PatternFlowIpv4DstMetricTag {
	obj := patternFlowIpv4DstMetricTag{obj: &otg.PatternFlowIpv4DstMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DstMetricTag) msg() *otg.PatternFlowIpv4DstMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4DstMetricTag) setMsg(msg *otg.PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DstMetricTag struct {
	obj *patternFlowIpv4DstMetricTag
}

type marshalPatternFlowIpv4DstMetricTag interface {
	// ToProto marshals PatternFlowIpv4DstMetricTag to protobuf object *otg.PatternFlowIpv4DstMetricTag
	ToProto() (*otg.PatternFlowIpv4DstMetricTag, error)
	// ToPbText marshals PatternFlowIpv4DstMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DstMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DstMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DstMetricTag struct {
	obj *patternFlowIpv4DstMetricTag
}

type unMarshalPatternFlowIpv4DstMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4DstMetricTag from protobuf object *otg.PatternFlowIpv4DstMetricTag
	FromProto(msg *otg.PatternFlowIpv4DstMetricTag) (PatternFlowIpv4DstMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4DstMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DstMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DstMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DstMetricTag) Marshal() marshalPatternFlowIpv4DstMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DstMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DstMetricTag) Unmarshal() unMarshalPatternFlowIpv4DstMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DstMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DstMetricTag) ToProto() (*otg.PatternFlowIpv4DstMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DstMetricTag) FromProto(msg *otg.PatternFlowIpv4DstMetricTag) (PatternFlowIpv4DstMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DstMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DstMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DstMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4DstMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DstMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DstMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DstMetricTag) Clone() (PatternFlowIpv4DstMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DstMetricTag()
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

// PatternFlowIpv4DstMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4DstMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4DstMetricTag to protobuf object *otg.PatternFlowIpv4DstMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DstMetricTag
	// setMsg unmarshals PatternFlowIpv4DstMetricTag from protobuf object *otg.PatternFlowIpv4DstMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DstMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DstMetricTag
	// validate validates PatternFlowIpv4DstMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DstMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4DstMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4DstMetricTag
	SetName(value string) PatternFlowIpv4DstMetricTag
	// Offset returns uint32, set in PatternFlowIpv4DstMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4DstMetricTag
	SetOffset(value uint32) PatternFlowIpv4DstMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4DstMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4DstMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4DstMetricTag
	SetLength(value uint32) PatternFlowIpv4DstMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4DstMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4DstMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4DstMetricTag object
func (obj *patternFlowIpv4DstMetricTag) SetName(value string) PatternFlowIpv4DstMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4DstMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4DstMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4DstMetricTag object
func (obj *patternFlowIpv4DstMetricTag) SetOffset(value uint32) PatternFlowIpv4DstMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4DstMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4DstMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4DstMetricTag object
func (obj *patternFlowIpv4DstMetricTag) SetLength(value uint32) PatternFlowIpv4DstMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4DstMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4DstMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DstMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4DstMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4DstMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
