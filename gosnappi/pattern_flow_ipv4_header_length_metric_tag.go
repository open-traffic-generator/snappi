package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4HeaderLengthMetricTag *****
type patternFlowIpv4HeaderLengthMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4HeaderLengthMetricTag
	marshaller   marshalPatternFlowIpv4HeaderLengthMetricTag
	unMarshaller unMarshalPatternFlowIpv4HeaderLengthMetricTag
}

func NewPatternFlowIpv4HeaderLengthMetricTag() PatternFlowIpv4HeaderLengthMetricTag {
	obj := patternFlowIpv4HeaderLengthMetricTag{obj: &otg.PatternFlowIpv4HeaderLengthMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4HeaderLengthMetricTag) msg() *otg.PatternFlowIpv4HeaderLengthMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4HeaderLengthMetricTag) setMsg(msg *otg.PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4HeaderLengthMetricTag struct {
	obj *patternFlowIpv4HeaderLengthMetricTag
}

type marshalPatternFlowIpv4HeaderLengthMetricTag interface {
	// ToProto marshals PatternFlowIpv4HeaderLengthMetricTag to protobuf object *otg.PatternFlowIpv4HeaderLengthMetricTag
	ToProto() (*otg.PatternFlowIpv4HeaderLengthMetricTag, error)
	// ToPbText marshals PatternFlowIpv4HeaderLengthMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4HeaderLengthMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4HeaderLengthMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4HeaderLengthMetricTag struct {
	obj *patternFlowIpv4HeaderLengthMetricTag
}

type unMarshalPatternFlowIpv4HeaderLengthMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4HeaderLengthMetricTag from protobuf object *otg.PatternFlowIpv4HeaderLengthMetricTag
	FromProto(msg *otg.PatternFlowIpv4HeaderLengthMetricTag) (PatternFlowIpv4HeaderLengthMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4HeaderLengthMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4HeaderLengthMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4HeaderLengthMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4HeaderLengthMetricTag) Marshal() marshalPatternFlowIpv4HeaderLengthMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4HeaderLengthMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4HeaderLengthMetricTag) Unmarshal() unMarshalPatternFlowIpv4HeaderLengthMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4HeaderLengthMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4HeaderLengthMetricTag) ToProto() (*otg.PatternFlowIpv4HeaderLengthMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4HeaderLengthMetricTag) FromProto(msg *otg.PatternFlowIpv4HeaderLengthMetricTag) (PatternFlowIpv4HeaderLengthMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4HeaderLengthMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLengthMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4HeaderLengthMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLengthMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4HeaderLengthMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLengthMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4HeaderLengthMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4HeaderLengthMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4HeaderLengthMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4HeaderLengthMetricTag) Clone() (PatternFlowIpv4HeaderLengthMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4HeaderLengthMetricTag()
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

// PatternFlowIpv4HeaderLengthMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4HeaderLengthMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4HeaderLengthMetricTag to protobuf object *otg.PatternFlowIpv4HeaderLengthMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4HeaderLengthMetricTag
	// setMsg unmarshals PatternFlowIpv4HeaderLengthMetricTag from protobuf object *otg.PatternFlowIpv4HeaderLengthMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4HeaderLengthMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4HeaderLengthMetricTag
	// validate validates PatternFlowIpv4HeaderLengthMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4HeaderLengthMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4HeaderLengthMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4HeaderLengthMetricTag
	SetName(value string) PatternFlowIpv4HeaderLengthMetricTag
	// Offset returns uint32, set in PatternFlowIpv4HeaderLengthMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4HeaderLengthMetricTag
	SetOffset(value uint32) PatternFlowIpv4HeaderLengthMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4HeaderLengthMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4HeaderLengthMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4HeaderLengthMetricTag
	SetLength(value uint32) PatternFlowIpv4HeaderLengthMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4HeaderLengthMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4HeaderLengthMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4HeaderLengthMetricTag object
func (obj *patternFlowIpv4HeaderLengthMetricTag) SetName(value string) PatternFlowIpv4HeaderLengthMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4HeaderLengthMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4HeaderLengthMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4HeaderLengthMetricTag object
func (obj *patternFlowIpv4HeaderLengthMetricTag) SetOffset(value uint32) PatternFlowIpv4HeaderLengthMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4HeaderLengthMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4HeaderLengthMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4HeaderLengthMetricTag object
func (obj *patternFlowIpv4HeaderLengthMetricTag) SetLength(value uint32) PatternFlowIpv4HeaderLengthMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4HeaderLengthMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4HeaderLengthMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLengthMetricTag.Offset <= 3 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4HeaderLengthMetricTag.Length <= 4 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4HeaderLengthMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(4)
	}

}
