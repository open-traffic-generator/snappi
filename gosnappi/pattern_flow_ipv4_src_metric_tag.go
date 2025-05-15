package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4SrcMetricTag *****
type patternFlowIpv4SrcMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4SrcMetricTag
	marshaller   marshalPatternFlowIpv4SrcMetricTag
	unMarshaller unMarshalPatternFlowIpv4SrcMetricTag
}

func NewPatternFlowIpv4SrcMetricTag() PatternFlowIpv4SrcMetricTag {
	obj := patternFlowIpv4SrcMetricTag{obj: &otg.PatternFlowIpv4SrcMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4SrcMetricTag) msg() *otg.PatternFlowIpv4SrcMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4SrcMetricTag) setMsg(msg *otg.PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4SrcMetricTag struct {
	obj *patternFlowIpv4SrcMetricTag
}

type marshalPatternFlowIpv4SrcMetricTag interface {
	// ToProto marshals PatternFlowIpv4SrcMetricTag to protobuf object *otg.PatternFlowIpv4SrcMetricTag
	ToProto() (*otg.PatternFlowIpv4SrcMetricTag, error)
	// ToPbText marshals PatternFlowIpv4SrcMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4SrcMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4SrcMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4SrcMetricTag struct {
	obj *patternFlowIpv4SrcMetricTag
}

type unMarshalPatternFlowIpv4SrcMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4SrcMetricTag from protobuf object *otg.PatternFlowIpv4SrcMetricTag
	FromProto(msg *otg.PatternFlowIpv4SrcMetricTag) (PatternFlowIpv4SrcMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4SrcMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4SrcMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4SrcMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4SrcMetricTag) Marshal() marshalPatternFlowIpv4SrcMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4SrcMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4SrcMetricTag) Unmarshal() unMarshalPatternFlowIpv4SrcMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4SrcMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4SrcMetricTag) ToProto() (*otg.PatternFlowIpv4SrcMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4SrcMetricTag) FromProto(msg *otg.PatternFlowIpv4SrcMetricTag) (PatternFlowIpv4SrcMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4SrcMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4SrcMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4SrcMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4SrcMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4SrcMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4SrcMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4SrcMetricTag) Clone() (PatternFlowIpv4SrcMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4SrcMetricTag()
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

// PatternFlowIpv4SrcMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4SrcMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4SrcMetricTag to protobuf object *otg.PatternFlowIpv4SrcMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4SrcMetricTag
	// setMsg unmarshals PatternFlowIpv4SrcMetricTag from protobuf object *otg.PatternFlowIpv4SrcMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4SrcMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4SrcMetricTag
	// validate validates PatternFlowIpv4SrcMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4SrcMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4SrcMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4SrcMetricTag
	SetName(value string) PatternFlowIpv4SrcMetricTag
	// Offset returns uint32, set in PatternFlowIpv4SrcMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4SrcMetricTag
	SetOffset(value uint32) PatternFlowIpv4SrcMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4SrcMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4SrcMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4SrcMetricTag
	SetLength(value uint32) PatternFlowIpv4SrcMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4SrcMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4SrcMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4SrcMetricTag object
func (obj *patternFlowIpv4SrcMetricTag) SetName(value string) PatternFlowIpv4SrcMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4SrcMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4SrcMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4SrcMetricTag object
func (obj *patternFlowIpv4SrcMetricTag) SetOffset(value uint32) PatternFlowIpv4SrcMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4SrcMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4SrcMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4SrcMetricTag object
func (obj *patternFlowIpv4SrcMetricTag) SetLength(value uint32) PatternFlowIpv4SrcMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4SrcMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4SrcMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4SrcMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4SrcMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4SrcMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
