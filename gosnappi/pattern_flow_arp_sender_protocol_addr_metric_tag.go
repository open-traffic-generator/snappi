package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpSenderProtocolAddrMetricTag *****
type patternFlowArpSenderProtocolAddrMetricTag struct {
	validation
	obj          *otg.PatternFlowArpSenderProtocolAddrMetricTag
	marshaller   marshalPatternFlowArpSenderProtocolAddrMetricTag
	unMarshaller unMarshalPatternFlowArpSenderProtocolAddrMetricTag
}

func NewPatternFlowArpSenderProtocolAddrMetricTag() PatternFlowArpSenderProtocolAddrMetricTag {
	obj := patternFlowArpSenderProtocolAddrMetricTag{obj: &otg.PatternFlowArpSenderProtocolAddrMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) msg() *otg.PatternFlowArpSenderProtocolAddrMetricTag {
	return obj.obj
}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) setMsg(msg *otg.PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpSenderProtocolAddrMetricTag struct {
	obj *patternFlowArpSenderProtocolAddrMetricTag
}

type marshalPatternFlowArpSenderProtocolAddrMetricTag interface {
	// ToProto marshals PatternFlowArpSenderProtocolAddrMetricTag to protobuf object *otg.PatternFlowArpSenderProtocolAddrMetricTag
	ToProto() (*otg.PatternFlowArpSenderProtocolAddrMetricTag, error)
	// ToPbText marshals PatternFlowArpSenderProtocolAddrMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpSenderProtocolAddrMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpSenderProtocolAddrMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpSenderProtocolAddrMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpSenderProtocolAddrMetricTag struct {
	obj *patternFlowArpSenderProtocolAddrMetricTag
}

type unMarshalPatternFlowArpSenderProtocolAddrMetricTag interface {
	// FromProto unmarshals PatternFlowArpSenderProtocolAddrMetricTag from protobuf object *otg.PatternFlowArpSenderProtocolAddrMetricTag
	FromProto(msg *otg.PatternFlowArpSenderProtocolAddrMetricTag) (PatternFlowArpSenderProtocolAddrMetricTag, error)
	// FromPbText unmarshals PatternFlowArpSenderProtocolAddrMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpSenderProtocolAddrMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpSenderProtocolAddrMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) Marshal() marshalPatternFlowArpSenderProtocolAddrMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpSenderProtocolAddrMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) Unmarshal() unMarshalPatternFlowArpSenderProtocolAddrMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpSenderProtocolAddrMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpSenderProtocolAddrMetricTag) ToProto() (*otg.PatternFlowArpSenderProtocolAddrMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpSenderProtocolAddrMetricTag) FromProto(msg *otg.PatternFlowArpSenderProtocolAddrMetricTag) (PatternFlowArpSenderProtocolAddrMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpSenderProtocolAddrMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddrMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpSenderProtocolAddrMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddrMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpSenderProtocolAddrMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpSenderProtocolAddrMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddrMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpSenderProtocolAddrMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) Clone() (PatternFlowArpSenderProtocolAddrMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpSenderProtocolAddrMetricTag()
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

// PatternFlowArpSenderProtocolAddrMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpSenderProtocolAddrMetricTag interface {
	Validation
	// msg marshals PatternFlowArpSenderProtocolAddrMetricTag to protobuf object *otg.PatternFlowArpSenderProtocolAddrMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpSenderProtocolAddrMetricTag
	// setMsg unmarshals PatternFlowArpSenderProtocolAddrMetricTag from protobuf object *otg.PatternFlowArpSenderProtocolAddrMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpSenderProtocolAddrMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpSenderProtocolAddrMetricTag
	// validate validates PatternFlowArpSenderProtocolAddrMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpSenderProtocolAddrMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpSenderProtocolAddrMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpSenderProtocolAddrMetricTag
	SetName(value string) PatternFlowArpSenderProtocolAddrMetricTag
	// Offset returns uint32, set in PatternFlowArpSenderProtocolAddrMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpSenderProtocolAddrMetricTag
	SetOffset(value uint32) PatternFlowArpSenderProtocolAddrMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpSenderProtocolAddrMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpSenderProtocolAddrMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpSenderProtocolAddrMetricTag
	SetLength(value uint32) PatternFlowArpSenderProtocolAddrMetricTag
	// HasLength checks if Length has been set in PatternFlowArpSenderProtocolAddrMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpSenderProtocolAddrMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpSenderProtocolAddrMetricTag object
func (obj *patternFlowArpSenderProtocolAddrMetricTag) SetName(value string) PatternFlowArpSenderProtocolAddrMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpSenderProtocolAddrMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpSenderProtocolAddrMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpSenderProtocolAddrMetricTag object
func (obj *patternFlowArpSenderProtocolAddrMetricTag) SetOffset(value uint32) PatternFlowArpSenderProtocolAddrMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpSenderProtocolAddrMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpSenderProtocolAddrMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpSenderProtocolAddrMetricTag object
func (obj *patternFlowArpSenderProtocolAddrMetricTag) SetLength(value uint32) PatternFlowArpSenderProtocolAddrMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpSenderProtocolAddrMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpSenderProtocolAddrMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpSenderProtocolAddrMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpSenderProtocolAddrMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
