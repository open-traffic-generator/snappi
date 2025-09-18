package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6DstMetricTag *****
type patternFlowIpv6DstMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv6DstMetricTag
	marshaller   marshalPatternFlowIpv6DstMetricTag
	unMarshaller unMarshalPatternFlowIpv6DstMetricTag
}

func NewPatternFlowIpv6DstMetricTag() PatternFlowIpv6DstMetricTag {
	obj := patternFlowIpv6DstMetricTag{obj: &otg.PatternFlowIpv6DstMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6DstMetricTag) msg() *otg.PatternFlowIpv6DstMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv6DstMetricTag) setMsg(msg *otg.PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6DstMetricTag struct {
	obj *patternFlowIpv6DstMetricTag
}

type marshalPatternFlowIpv6DstMetricTag interface {
	// ToProto marshals PatternFlowIpv6DstMetricTag to protobuf object *otg.PatternFlowIpv6DstMetricTag
	ToProto() (*otg.PatternFlowIpv6DstMetricTag, error)
	// ToPbText marshals PatternFlowIpv6DstMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6DstMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6DstMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6DstMetricTag struct {
	obj *patternFlowIpv6DstMetricTag
}

type unMarshalPatternFlowIpv6DstMetricTag interface {
	// FromProto unmarshals PatternFlowIpv6DstMetricTag from protobuf object *otg.PatternFlowIpv6DstMetricTag
	FromProto(msg *otg.PatternFlowIpv6DstMetricTag) (PatternFlowIpv6DstMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv6DstMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6DstMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6DstMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6DstMetricTag) Marshal() marshalPatternFlowIpv6DstMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6DstMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6DstMetricTag) Unmarshal() unMarshalPatternFlowIpv6DstMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6DstMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6DstMetricTag) ToProto() (*otg.PatternFlowIpv6DstMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6DstMetricTag) FromProto(msg *otg.PatternFlowIpv6DstMetricTag) (PatternFlowIpv6DstMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6DstMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6DstMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6DstMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6DstMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6DstMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6DstMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6DstMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6DstMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6DstMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6DstMetricTag) Clone() (PatternFlowIpv6DstMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6DstMetricTag()
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

// PatternFlowIpv6DstMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv6DstMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv6DstMetricTag to protobuf object *otg.PatternFlowIpv6DstMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6DstMetricTag
	// setMsg unmarshals PatternFlowIpv6DstMetricTag from protobuf object *otg.PatternFlowIpv6DstMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6DstMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6DstMetricTag
	// validate validates PatternFlowIpv6DstMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6DstMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv6DstMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv6DstMetricTag
	SetName(value string) PatternFlowIpv6DstMetricTag
	// Offset returns uint32, set in PatternFlowIpv6DstMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv6DstMetricTag
	SetOffset(value uint32) PatternFlowIpv6DstMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv6DstMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv6DstMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv6DstMetricTag
	SetLength(value uint32) PatternFlowIpv6DstMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv6DstMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv6DstMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv6DstMetricTag object
func (obj *patternFlowIpv6DstMetricTag) SetName(value string) PatternFlowIpv6DstMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6DstMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6DstMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv6DstMetricTag object
func (obj *patternFlowIpv6DstMetricTag) SetOffset(value uint32) PatternFlowIpv6DstMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6DstMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6DstMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv6DstMetricTag object
func (obj *patternFlowIpv6DstMetricTag) SetLength(value uint32) PatternFlowIpv6DstMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv6DstMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv6DstMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6DstMetricTag.Offset <= 127 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv6DstMetricTag.Length <= 128 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv6DstMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(128)
	}

}
