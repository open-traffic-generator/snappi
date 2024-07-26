package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosUnusedMetricTag *****
type patternFlowIpv4TosUnusedMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4TosUnusedMetricTag
	marshaller   marshalPatternFlowIpv4TosUnusedMetricTag
	unMarshaller unMarshalPatternFlowIpv4TosUnusedMetricTag
}

func NewPatternFlowIpv4TosUnusedMetricTag() PatternFlowIpv4TosUnusedMetricTag {
	obj := patternFlowIpv4TosUnusedMetricTag{obj: &otg.PatternFlowIpv4TosUnusedMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosUnusedMetricTag) msg() *otg.PatternFlowIpv4TosUnusedMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4TosUnusedMetricTag) setMsg(msg *otg.PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosUnusedMetricTag struct {
	obj *patternFlowIpv4TosUnusedMetricTag
}

type marshalPatternFlowIpv4TosUnusedMetricTag interface {
	// ToProto marshals PatternFlowIpv4TosUnusedMetricTag to protobuf object *otg.PatternFlowIpv4TosUnusedMetricTag
	ToProto() (*otg.PatternFlowIpv4TosUnusedMetricTag, error)
	// ToPbText marshals PatternFlowIpv4TosUnusedMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosUnusedMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosUnusedMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosUnusedMetricTag struct {
	obj *patternFlowIpv4TosUnusedMetricTag
}

type unMarshalPatternFlowIpv4TosUnusedMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4TosUnusedMetricTag from protobuf object *otg.PatternFlowIpv4TosUnusedMetricTag
	FromProto(msg *otg.PatternFlowIpv4TosUnusedMetricTag) (PatternFlowIpv4TosUnusedMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4TosUnusedMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosUnusedMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosUnusedMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosUnusedMetricTag) Marshal() marshalPatternFlowIpv4TosUnusedMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosUnusedMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosUnusedMetricTag) Unmarshal() unMarshalPatternFlowIpv4TosUnusedMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosUnusedMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosUnusedMetricTag) ToProto() (*otg.PatternFlowIpv4TosUnusedMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosUnusedMetricTag) FromProto(msg *otg.PatternFlowIpv4TosUnusedMetricTag) (PatternFlowIpv4TosUnusedMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosUnusedMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnusedMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosUnusedMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnusedMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosUnusedMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnusedMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosUnusedMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosUnusedMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosUnusedMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosUnusedMetricTag) Clone() (PatternFlowIpv4TosUnusedMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosUnusedMetricTag()
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

// PatternFlowIpv4TosUnusedMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4TosUnusedMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4TosUnusedMetricTag to protobuf object *otg.PatternFlowIpv4TosUnusedMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosUnusedMetricTag
	// setMsg unmarshals PatternFlowIpv4TosUnusedMetricTag from protobuf object *otg.PatternFlowIpv4TosUnusedMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosUnusedMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosUnusedMetricTag
	// validate validates PatternFlowIpv4TosUnusedMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosUnusedMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4TosUnusedMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4TosUnusedMetricTag
	SetName(value string) PatternFlowIpv4TosUnusedMetricTag
	// Offset returns uint32, set in PatternFlowIpv4TosUnusedMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4TosUnusedMetricTag
	SetOffset(value uint32) PatternFlowIpv4TosUnusedMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4TosUnusedMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4TosUnusedMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4TosUnusedMetricTag
	SetLength(value uint32) PatternFlowIpv4TosUnusedMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4TosUnusedMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4TosUnusedMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4TosUnusedMetricTag object
func (obj *patternFlowIpv4TosUnusedMetricTag) SetName(value string) PatternFlowIpv4TosUnusedMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosUnusedMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TosUnusedMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4TosUnusedMetricTag object
func (obj *patternFlowIpv4TosUnusedMetricTag) SetOffset(value uint32) PatternFlowIpv4TosUnusedMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosUnusedMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TosUnusedMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4TosUnusedMetricTag object
func (obj *patternFlowIpv4TosUnusedMetricTag) SetLength(value uint32) PatternFlowIpv4TosUnusedMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4TosUnusedMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4TosUnusedMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosUnusedMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4TosUnusedMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4TosUnusedMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
