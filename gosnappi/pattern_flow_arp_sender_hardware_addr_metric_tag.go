package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpSenderHardwareAddrMetricTag *****
type patternFlowArpSenderHardwareAddrMetricTag struct {
	validation
	obj          *otg.PatternFlowArpSenderHardwareAddrMetricTag
	marshaller   marshalPatternFlowArpSenderHardwareAddrMetricTag
	unMarshaller unMarshalPatternFlowArpSenderHardwareAddrMetricTag
}

func NewPatternFlowArpSenderHardwareAddrMetricTag() PatternFlowArpSenderHardwareAddrMetricTag {
	obj := patternFlowArpSenderHardwareAddrMetricTag{obj: &otg.PatternFlowArpSenderHardwareAddrMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) msg() *otg.PatternFlowArpSenderHardwareAddrMetricTag {
	return obj.obj
}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) setMsg(msg *otg.PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpSenderHardwareAddrMetricTag struct {
	obj *patternFlowArpSenderHardwareAddrMetricTag
}

type marshalPatternFlowArpSenderHardwareAddrMetricTag interface {
	// ToProto marshals PatternFlowArpSenderHardwareAddrMetricTag to protobuf object *otg.PatternFlowArpSenderHardwareAddrMetricTag
	ToProto() (*otg.PatternFlowArpSenderHardwareAddrMetricTag, error)
	// ToPbText marshals PatternFlowArpSenderHardwareAddrMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpSenderHardwareAddrMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpSenderHardwareAddrMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpSenderHardwareAddrMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpSenderHardwareAddrMetricTag struct {
	obj *patternFlowArpSenderHardwareAddrMetricTag
}

type unMarshalPatternFlowArpSenderHardwareAddrMetricTag interface {
	// FromProto unmarshals PatternFlowArpSenderHardwareAddrMetricTag from protobuf object *otg.PatternFlowArpSenderHardwareAddrMetricTag
	FromProto(msg *otg.PatternFlowArpSenderHardwareAddrMetricTag) (PatternFlowArpSenderHardwareAddrMetricTag, error)
	// FromPbText unmarshals PatternFlowArpSenderHardwareAddrMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpSenderHardwareAddrMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpSenderHardwareAddrMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) Marshal() marshalPatternFlowArpSenderHardwareAddrMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpSenderHardwareAddrMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) Unmarshal() unMarshalPatternFlowArpSenderHardwareAddrMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpSenderHardwareAddrMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpSenderHardwareAddrMetricTag) ToProto() (*otg.PatternFlowArpSenderHardwareAddrMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpSenderHardwareAddrMetricTag) FromProto(msg *otg.PatternFlowArpSenderHardwareAddrMetricTag) (PatternFlowArpSenderHardwareAddrMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpSenderHardwareAddrMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddrMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpSenderHardwareAddrMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddrMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpSenderHardwareAddrMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpSenderHardwareAddrMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddrMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpSenderHardwareAddrMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) Clone() (PatternFlowArpSenderHardwareAddrMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpSenderHardwareAddrMetricTag()
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

// PatternFlowArpSenderHardwareAddrMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpSenderHardwareAddrMetricTag interface {
	Validation
	// msg marshals PatternFlowArpSenderHardwareAddrMetricTag to protobuf object *otg.PatternFlowArpSenderHardwareAddrMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpSenderHardwareAddrMetricTag
	// setMsg unmarshals PatternFlowArpSenderHardwareAddrMetricTag from protobuf object *otg.PatternFlowArpSenderHardwareAddrMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpSenderHardwareAddrMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpSenderHardwareAddrMetricTag
	// validate validates PatternFlowArpSenderHardwareAddrMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpSenderHardwareAddrMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpSenderHardwareAddrMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpSenderHardwareAddrMetricTag
	SetName(value string) PatternFlowArpSenderHardwareAddrMetricTag
	// Offset returns uint32, set in PatternFlowArpSenderHardwareAddrMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpSenderHardwareAddrMetricTag
	SetOffset(value uint32) PatternFlowArpSenderHardwareAddrMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpSenderHardwareAddrMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpSenderHardwareAddrMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpSenderHardwareAddrMetricTag
	SetLength(value uint32) PatternFlowArpSenderHardwareAddrMetricTag
	// HasLength checks if Length has been set in PatternFlowArpSenderHardwareAddrMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpSenderHardwareAddrMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpSenderHardwareAddrMetricTag object
func (obj *patternFlowArpSenderHardwareAddrMetricTag) SetName(value string) PatternFlowArpSenderHardwareAddrMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpSenderHardwareAddrMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpSenderHardwareAddrMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpSenderHardwareAddrMetricTag object
func (obj *patternFlowArpSenderHardwareAddrMetricTag) SetOffset(value uint32) PatternFlowArpSenderHardwareAddrMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpSenderHardwareAddrMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpSenderHardwareAddrMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpSenderHardwareAddrMetricTag object
func (obj *patternFlowArpSenderHardwareAddrMetricTag) SetLength(value uint32) PatternFlowArpSenderHardwareAddrMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpSenderHardwareAddrMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpSenderHardwareAddrMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpSenderHardwareAddrMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpSenderHardwareAddrMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}
