package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6NextHeaderMetricTag *****
type patternFlowIpv6NextHeaderMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv6NextHeaderMetricTag
	marshaller   marshalPatternFlowIpv6NextHeaderMetricTag
	unMarshaller unMarshalPatternFlowIpv6NextHeaderMetricTag
}

func NewPatternFlowIpv6NextHeaderMetricTag() PatternFlowIpv6NextHeaderMetricTag {
	obj := patternFlowIpv6NextHeaderMetricTag{obj: &otg.PatternFlowIpv6NextHeaderMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6NextHeaderMetricTag) msg() *otg.PatternFlowIpv6NextHeaderMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv6NextHeaderMetricTag) setMsg(msg *otg.PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6NextHeaderMetricTag struct {
	obj *patternFlowIpv6NextHeaderMetricTag
}

type marshalPatternFlowIpv6NextHeaderMetricTag interface {
	// ToProto marshals PatternFlowIpv6NextHeaderMetricTag to protobuf object *otg.PatternFlowIpv6NextHeaderMetricTag
	ToProto() (*otg.PatternFlowIpv6NextHeaderMetricTag, error)
	// ToPbText marshals PatternFlowIpv6NextHeaderMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6NextHeaderMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6NextHeaderMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6NextHeaderMetricTag struct {
	obj *patternFlowIpv6NextHeaderMetricTag
}

type unMarshalPatternFlowIpv6NextHeaderMetricTag interface {
	// FromProto unmarshals PatternFlowIpv6NextHeaderMetricTag from protobuf object *otg.PatternFlowIpv6NextHeaderMetricTag
	FromProto(msg *otg.PatternFlowIpv6NextHeaderMetricTag) (PatternFlowIpv6NextHeaderMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv6NextHeaderMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6NextHeaderMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6NextHeaderMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6NextHeaderMetricTag) Marshal() marshalPatternFlowIpv6NextHeaderMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6NextHeaderMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6NextHeaderMetricTag) Unmarshal() unMarshalPatternFlowIpv6NextHeaderMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6NextHeaderMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6NextHeaderMetricTag) ToProto() (*otg.PatternFlowIpv6NextHeaderMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6NextHeaderMetricTag) FromProto(msg *otg.PatternFlowIpv6NextHeaderMetricTag) (PatternFlowIpv6NextHeaderMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6NextHeaderMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeaderMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6NextHeaderMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeaderMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6NextHeaderMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeaderMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6NextHeaderMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6NextHeaderMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6NextHeaderMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6NextHeaderMetricTag) Clone() (PatternFlowIpv6NextHeaderMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6NextHeaderMetricTag()
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

// PatternFlowIpv6NextHeaderMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv6NextHeaderMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv6NextHeaderMetricTag to protobuf object *otg.PatternFlowIpv6NextHeaderMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6NextHeaderMetricTag
	// setMsg unmarshals PatternFlowIpv6NextHeaderMetricTag from protobuf object *otg.PatternFlowIpv6NextHeaderMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6NextHeaderMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6NextHeaderMetricTag
	// validate validates PatternFlowIpv6NextHeaderMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6NextHeaderMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv6NextHeaderMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv6NextHeaderMetricTag
	SetName(value string) PatternFlowIpv6NextHeaderMetricTag
	// Offset returns uint32, set in PatternFlowIpv6NextHeaderMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv6NextHeaderMetricTag
	SetOffset(value uint32) PatternFlowIpv6NextHeaderMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv6NextHeaderMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv6NextHeaderMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv6NextHeaderMetricTag
	SetLength(value uint32) PatternFlowIpv6NextHeaderMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv6NextHeaderMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv6NextHeaderMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv6NextHeaderMetricTag object
func (obj *patternFlowIpv6NextHeaderMetricTag) SetName(value string) PatternFlowIpv6NextHeaderMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6NextHeaderMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6NextHeaderMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv6NextHeaderMetricTag object
func (obj *patternFlowIpv6NextHeaderMetricTag) SetOffset(value uint32) PatternFlowIpv6NextHeaderMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6NextHeaderMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6NextHeaderMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv6NextHeaderMetricTag object
func (obj *patternFlowIpv6NextHeaderMetricTag) SetLength(value uint32) PatternFlowIpv6NextHeaderMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv6NextHeaderMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv6NextHeaderMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeaderMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv6NextHeaderMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv6NextHeaderMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
