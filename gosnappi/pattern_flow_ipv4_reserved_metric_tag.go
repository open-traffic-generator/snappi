package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4ReservedMetricTag *****
type patternFlowIpv4ReservedMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4ReservedMetricTag
	marshaller   marshalPatternFlowIpv4ReservedMetricTag
	unMarshaller unMarshalPatternFlowIpv4ReservedMetricTag
}

func NewPatternFlowIpv4ReservedMetricTag() PatternFlowIpv4ReservedMetricTag {
	obj := patternFlowIpv4ReservedMetricTag{obj: &otg.PatternFlowIpv4ReservedMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4ReservedMetricTag) msg() *otg.PatternFlowIpv4ReservedMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4ReservedMetricTag) setMsg(msg *otg.PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4ReservedMetricTag struct {
	obj *patternFlowIpv4ReservedMetricTag
}

type marshalPatternFlowIpv4ReservedMetricTag interface {
	// ToProto marshals PatternFlowIpv4ReservedMetricTag to protobuf object *otg.PatternFlowIpv4ReservedMetricTag
	ToProto() (*otg.PatternFlowIpv4ReservedMetricTag, error)
	// ToPbText marshals PatternFlowIpv4ReservedMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4ReservedMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4ReservedMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4ReservedMetricTag struct {
	obj *patternFlowIpv4ReservedMetricTag
}

type unMarshalPatternFlowIpv4ReservedMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4ReservedMetricTag from protobuf object *otg.PatternFlowIpv4ReservedMetricTag
	FromProto(msg *otg.PatternFlowIpv4ReservedMetricTag) (PatternFlowIpv4ReservedMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4ReservedMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4ReservedMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4ReservedMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4ReservedMetricTag) Marshal() marshalPatternFlowIpv4ReservedMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4ReservedMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4ReservedMetricTag) Unmarshal() unMarshalPatternFlowIpv4ReservedMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4ReservedMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4ReservedMetricTag) ToProto() (*otg.PatternFlowIpv4ReservedMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4ReservedMetricTag) FromProto(msg *otg.PatternFlowIpv4ReservedMetricTag) (PatternFlowIpv4ReservedMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4ReservedMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ReservedMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4ReservedMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ReservedMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4ReservedMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ReservedMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4ReservedMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4ReservedMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4ReservedMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4ReservedMetricTag) Clone() (PatternFlowIpv4ReservedMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4ReservedMetricTag()
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

// PatternFlowIpv4ReservedMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4ReservedMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4ReservedMetricTag to protobuf object *otg.PatternFlowIpv4ReservedMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4ReservedMetricTag
	// setMsg unmarshals PatternFlowIpv4ReservedMetricTag from protobuf object *otg.PatternFlowIpv4ReservedMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4ReservedMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4ReservedMetricTag
	// validate validates PatternFlowIpv4ReservedMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4ReservedMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4ReservedMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4ReservedMetricTag
	SetName(value string) PatternFlowIpv4ReservedMetricTag
	// Offset returns uint32, set in PatternFlowIpv4ReservedMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4ReservedMetricTag
	SetOffset(value uint32) PatternFlowIpv4ReservedMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4ReservedMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4ReservedMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4ReservedMetricTag
	SetLength(value uint32) PatternFlowIpv4ReservedMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4ReservedMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4ReservedMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4ReservedMetricTag object
func (obj *patternFlowIpv4ReservedMetricTag) SetName(value string) PatternFlowIpv4ReservedMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4ReservedMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4ReservedMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4ReservedMetricTag object
func (obj *patternFlowIpv4ReservedMetricTag) SetOffset(value uint32) PatternFlowIpv4ReservedMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4ReservedMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4ReservedMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4ReservedMetricTag object
func (obj *patternFlowIpv4ReservedMetricTag) SetLength(value uint32) PatternFlowIpv4ReservedMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4ReservedMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4ReservedMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4ReservedMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4ReservedMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4ReservedMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
