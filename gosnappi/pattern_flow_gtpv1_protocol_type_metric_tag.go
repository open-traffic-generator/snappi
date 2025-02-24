package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1ProtocolTypeMetricTag *****
type patternFlowGtpv1ProtocolTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1ProtocolTypeMetricTag
	marshaller   marshalPatternFlowGtpv1ProtocolTypeMetricTag
	unMarshaller unMarshalPatternFlowGtpv1ProtocolTypeMetricTag
}

func NewPatternFlowGtpv1ProtocolTypeMetricTag() PatternFlowGtpv1ProtocolTypeMetricTag {
	obj := patternFlowGtpv1ProtocolTypeMetricTag{obj: &otg.PatternFlowGtpv1ProtocolTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) msg() *otg.PatternFlowGtpv1ProtocolTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) setMsg(msg *otg.PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1ProtocolTypeMetricTag struct {
	obj *patternFlowGtpv1ProtocolTypeMetricTag
}

type marshalPatternFlowGtpv1ProtocolTypeMetricTag interface {
	// ToProto marshals PatternFlowGtpv1ProtocolTypeMetricTag to protobuf object *otg.PatternFlowGtpv1ProtocolTypeMetricTag
	ToProto() (*otg.PatternFlowGtpv1ProtocolTypeMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1ProtocolTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1ProtocolTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1ProtocolTypeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1ProtocolTypeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1ProtocolTypeMetricTag struct {
	obj *patternFlowGtpv1ProtocolTypeMetricTag
}

type unMarshalPatternFlowGtpv1ProtocolTypeMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1ProtocolTypeMetricTag from protobuf object *otg.PatternFlowGtpv1ProtocolTypeMetricTag
	FromProto(msg *otg.PatternFlowGtpv1ProtocolTypeMetricTag) (PatternFlowGtpv1ProtocolTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1ProtocolTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1ProtocolTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1ProtocolTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) Marshal() marshalPatternFlowGtpv1ProtocolTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1ProtocolTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) Unmarshal() unMarshalPatternFlowGtpv1ProtocolTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1ProtocolTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1ProtocolTypeMetricTag) ToProto() (*otg.PatternFlowGtpv1ProtocolTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1ProtocolTypeMetricTag) FromProto(msg *otg.PatternFlowGtpv1ProtocolTypeMetricTag) (PatternFlowGtpv1ProtocolTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1ProtocolTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1ProtocolTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1ProtocolTypeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1ProtocolTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) Clone() (PatternFlowGtpv1ProtocolTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1ProtocolTypeMetricTag()
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

// PatternFlowGtpv1ProtocolTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1ProtocolTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1ProtocolTypeMetricTag to protobuf object *otg.PatternFlowGtpv1ProtocolTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1ProtocolTypeMetricTag
	// setMsg unmarshals PatternFlowGtpv1ProtocolTypeMetricTag from protobuf object *otg.PatternFlowGtpv1ProtocolTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1ProtocolTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1ProtocolTypeMetricTag
	// validate validates PatternFlowGtpv1ProtocolTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1ProtocolTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1ProtocolTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1ProtocolTypeMetricTag
	SetName(value string) PatternFlowGtpv1ProtocolTypeMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1ProtocolTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1ProtocolTypeMetricTag
	SetOffset(value uint32) PatternFlowGtpv1ProtocolTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1ProtocolTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1ProtocolTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1ProtocolTypeMetricTag
	SetLength(value uint32) PatternFlowGtpv1ProtocolTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1ProtocolTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1ProtocolTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1ProtocolTypeMetricTag object
func (obj *patternFlowGtpv1ProtocolTypeMetricTag) SetName(value string) PatternFlowGtpv1ProtocolTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1ProtocolTypeMetricTag object
func (obj *patternFlowGtpv1ProtocolTypeMetricTag) SetOffset(value uint32) PatternFlowGtpv1ProtocolTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1ProtocolTypeMetricTag object
func (obj *patternFlowGtpv1ProtocolTypeMetricTag) SetLength(value uint32) PatternFlowGtpv1ProtocolTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1ProtocolTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolTypeMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1ProtocolTypeMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1ProtocolTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
