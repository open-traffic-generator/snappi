package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6TrafficClassMetricTag *****
type patternFlowIpv6TrafficClassMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv6TrafficClassMetricTag
	marshaller   marshalPatternFlowIpv6TrafficClassMetricTag
	unMarshaller unMarshalPatternFlowIpv6TrafficClassMetricTag
}

func NewPatternFlowIpv6TrafficClassMetricTag() PatternFlowIpv6TrafficClassMetricTag {
	obj := patternFlowIpv6TrafficClassMetricTag{obj: &otg.PatternFlowIpv6TrafficClassMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6TrafficClassMetricTag) msg() *otg.PatternFlowIpv6TrafficClassMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv6TrafficClassMetricTag) setMsg(msg *otg.PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6TrafficClassMetricTag struct {
	obj *patternFlowIpv6TrafficClassMetricTag
}

type marshalPatternFlowIpv6TrafficClassMetricTag interface {
	// ToProto marshals PatternFlowIpv6TrafficClassMetricTag to protobuf object *otg.PatternFlowIpv6TrafficClassMetricTag
	ToProto() (*otg.PatternFlowIpv6TrafficClassMetricTag, error)
	// ToPbText marshals PatternFlowIpv6TrafficClassMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6TrafficClassMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6TrafficClassMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6TrafficClassMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6TrafficClassMetricTag struct {
	obj *patternFlowIpv6TrafficClassMetricTag
}

type unMarshalPatternFlowIpv6TrafficClassMetricTag interface {
	// FromProto unmarshals PatternFlowIpv6TrafficClassMetricTag from protobuf object *otg.PatternFlowIpv6TrafficClassMetricTag
	FromProto(msg *otg.PatternFlowIpv6TrafficClassMetricTag) (PatternFlowIpv6TrafficClassMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv6TrafficClassMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6TrafficClassMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6TrafficClassMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6TrafficClassMetricTag) Marshal() marshalPatternFlowIpv6TrafficClassMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6TrafficClassMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6TrafficClassMetricTag) Unmarshal() unMarshalPatternFlowIpv6TrafficClassMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6TrafficClassMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6TrafficClassMetricTag) ToProto() (*otg.PatternFlowIpv6TrafficClassMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6TrafficClassMetricTag) FromProto(msg *otg.PatternFlowIpv6TrafficClassMetricTag) (PatternFlowIpv6TrafficClassMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6TrafficClassMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClassMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6TrafficClassMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClassMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6TrafficClassMetricTag) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowIpv6TrafficClassMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClassMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6TrafficClassMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6TrafficClassMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6TrafficClassMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6TrafficClassMetricTag) Clone() (PatternFlowIpv6TrafficClassMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6TrafficClassMetricTag()
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

// PatternFlowIpv6TrafficClassMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv6TrafficClassMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv6TrafficClassMetricTag to protobuf object *otg.PatternFlowIpv6TrafficClassMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6TrafficClassMetricTag
	// setMsg unmarshals PatternFlowIpv6TrafficClassMetricTag from protobuf object *otg.PatternFlowIpv6TrafficClassMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6TrafficClassMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6TrafficClassMetricTag
	// validate validates PatternFlowIpv6TrafficClassMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6TrafficClassMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv6TrafficClassMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv6TrafficClassMetricTag
	SetName(value string) PatternFlowIpv6TrafficClassMetricTag
	// Offset returns uint32, set in PatternFlowIpv6TrafficClassMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv6TrafficClassMetricTag
	SetOffset(value uint32) PatternFlowIpv6TrafficClassMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv6TrafficClassMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv6TrafficClassMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv6TrafficClassMetricTag
	SetLength(value uint32) PatternFlowIpv6TrafficClassMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv6TrafficClassMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv6TrafficClassMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv6TrafficClassMetricTag object
func (obj *patternFlowIpv6TrafficClassMetricTag) SetName(value string) PatternFlowIpv6TrafficClassMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6TrafficClassMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6TrafficClassMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv6TrafficClassMetricTag object
func (obj *patternFlowIpv6TrafficClassMetricTag) SetOffset(value uint32) PatternFlowIpv6TrafficClassMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6TrafficClassMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6TrafficClassMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv6TrafficClassMetricTag object
func (obj *patternFlowIpv6TrafficClassMetricTag) SetLength(value uint32) PatternFlowIpv6TrafficClassMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv6TrafficClassMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv6TrafficClassMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6TrafficClassMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv6TrafficClassMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv6TrafficClassMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
