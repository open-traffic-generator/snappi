package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4PriorityRawMetricTag *****
type patternFlowIpv4PriorityRawMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4PriorityRawMetricTag
	marshaller   marshalPatternFlowIpv4PriorityRawMetricTag
	unMarshaller unMarshalPatternFlowIpv4PriorityRawMetricTag
}

func NewPatternFlowIpv4PriorityRawMetricTag() PatternFlowIpv4PriorityRawMetricTag {
	obj := patternFlowIpv4PriorityRawMetricTag{obj: &otg.PatternFlowIpv4PriorityRawMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4PriorityRawMetricTag) msg() *otg.PatternFlowIpv4PriorityRawMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4PriorityRawMetricTag) setMsg(msg *otg.PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4PriorityRawMetricTag struct {
	obj *patternFlowIpv4PriorityRawMetricTag
}

type marshalPatternFlowIpv4PriorityRawMetricTag interface {
	// ToProto marshals PatternFlowIpv4PriorityRawMetricTag to protobuf object *otg.PatternFlowIpv4PriorityRawMetricTag
	ToProto() (*otg.PatternFlowIpv4PriorityRawMetricTag, error)
	// ToPbText marshals PatternFlowIpv4PriorityRawMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4PriorityRawMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4PriorityRawMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4PriorityRawMetricTag struct {
	obj *patternFlowIpv4PriorityRawMetricTag
}

type unMarshalPatternFlowIpv4PriorityRawMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4PriorityRawMetricTag from protobuf object *otg.PatternFlowIpv4PriorityRawMetricTag
	FromProto(msg *otg.PatternFlowIpv4PriorityRawMetricTag) (PatternFlowIpv4PriorityRawMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4PriorityRawMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4PriorityRawMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4PriorityRawMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4PriorityRawMetricTag) Marshal() marshalPatternFlowIpv4PriorityRawMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4PriorityRawMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4PriorityRawMetricTag) Unmarshal() unMarshalPatternFlowIpv4PriorityRawMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4PriorityRawMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4PriorityRawMetricTag) ToProto() (*otg.PatternFlowIpv4PriorityRawMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4PriorityRawMetricTag) FromProto(msg *otg.PatternFlowIpv4PriorityRawMetricTag) (PatternFlowIpv4PriorityRawMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4PriorityRawMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRawMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4PriorityRawMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRawMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4PriorityRawMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRawMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4PriorityRawMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4PriorityRawMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4PriorityRawMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4PriorityRawMetricTag) Clone() (PatternFlowIpv4PriorityRawMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4PriorityRawMetricTag()
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

// PatternFlowIpv4PriorityRawMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4PriorityRawMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4PriorityRawMetricTag to protobuf object *otg.PatternFlowIpv4PriorityRawMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4PriorityRawMetricTag
	// setMsg unmarshals PatternFlowIpv4PriorityRawMetricTag from protobuf object *otg.PatternFlowIpv4PriorityRawMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4PriorityRawMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4PriorityRawMetricTag
	// validate validates PatternFlowIpv4PriorityRawMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4PriorityRawMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4PriorityRawMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4PriorityRawMetricTag
	SetName(value string) PatternFlowIpv4PriorityRawMetricTag
	// Offset returns uint32, set in PatternFlowIpv4PriorityRawMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4PriorityRawMetricTag
	SetOffset(value uint32) PatternFlowIpv4PriorityRawMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4PriorityRawMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4PriorityRawMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4PriorityRawMetricTag
	SetLength(value uint32) PatternFlowIpv4PriorityRawMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4PriorityRawMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4PriorityRawMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4PriorityRawMetricTag object
func (obj *patternFlowIpv4PriorityRawMetricTag) SetName(value string) PatternFlowIpv4PriorityRawMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4PriorityRawMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4PriorityRawMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4PriorityRawMetricTag object
func (obj *patternFlowIpv4PriorityRawMetricTag) SetOffset(value uint32) PatternFlowIpv4PriorityRawMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4PriorityRawMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4PriorityRawMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4PriorityRawMetricTag object
func (obj *patternFlowIpv4PriorityRawMetricTag) SetLength(value uint32) PatternFlowIpv4PriorityRawMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4PriorityRawMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4PriorityRawMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4PriorityRawMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4PriorityRawMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4PriorityRawMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
