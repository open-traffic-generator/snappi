package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpTargetProtocolAddrMetricTag *****
type patternFlowArpTargetProtocolAddrMetricTag struct {
	validation
	obj          *otg.PatternFlowArpTargetProtocolAddrMetricTag
	marshaller   marshalPatternFlowArpTargetProtocolAddrMetricTag
	unMarshaller unMarshalPatternFlowArpTargetProtocolAddrMetricTag
}

func NewPatternFlowArpTargetProtocolAddrMetricTag() PatternFlowArpTargetProtocolAddrMetricTag {
	obj := patternFlowArpTargetProtocolAddrMetricTag{obj: &otg.PatternFlowArpTargetProtocolAddrMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) msg() *otg.PatternFlowArpTargetProtocolAddrMetricTag {
	return obj.obj
}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) setMsg(msg *otg.PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpTargetProtocolAddrMetricTag struct {
	obj *patternFlowArpTargetProtocolAddrMetricTag
}

type marshalPatternFlowArpTargetProtocolAddrMetricTag interface {
	// ToProto marshals PatternFlowArpTargetProtocolAddrMetricTag to protobuf object *otg.PatternFlowArpTargetProtocolAddrMetricTag
	ToProto() (*otg.PatternFlowArpTargetProtocolAddrMetricTag, error)
	// ToPbText marshals PatternFlowArpTargetProtocolAddrMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpTargetProtocolAddrMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpTargetProtocolAddrMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpTargetProtocolAddrMetricTag struct {
	obj *patternFlowArpTargetProtocolAddrMetricTag
}

type unMarshalPatternFlowArpTargetProtocolAddrMetricTag interface {
	// FromProto unmarshals PatternFlowArpTargetProtocolAddrMetricTag from protobuf object *otg.PatternFlowArpTargetProtocolAddrMetricTag
	FromProto(msg *otg.PatternFlowArpTargetProtocolAddrMetricTag) (PatternFlowArpTargetProtocolAddrMetricTag, error)
	// FromPbText unmarshals PatternFlowArpTargetProtocolAddrMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpTargetProtocolAddrMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpTargetProtocolAddrMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) Marshal() marshalPatternFlowArpTargetProtocolAddrMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpTargetProtocolAddrMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) Unmarshal() unMarshalPatternFlowArpTargetProtocolAddrMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpTargetProtocolAddrMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpTargetProtocolAddrMetricTag) ToProto() (*otg.PatternFlowArpTargetProtocolAddrMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpTargetProtocolAddrMetricTag) FromProto(msg *otg.PatternFlowArpTargetProtocolAddrMetricTag) (PatternFlowArpTargetProtocolAddrMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpTargetProtocolAddrMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddrMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpTargetProtocolAddrMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddrMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpTargetProtocolAddrMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddrMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpTargetProtocolAddrMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) Clone() (PatternFlowArpTargetProtocolAddrMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpTargetProtocolAddrMetricTag()
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

// PatternFlowArpTargetProtocolAddrMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpTargetProtocolAddrMetricTag interface {
	Validation
	// msg marshals PatternFlowArpTargetProtocolAddrMetricTag to protobuf object *otg.PatternFlowArpTargetProtocolAddrMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpTargetProtocolAddrMetricTag
	// setMsg unmarshals PatternFlowArpTargetProtocolAddrMetricTag from protobuf object *otg.PatternFlowArpTargetProtocolAddrMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpTargetProtocolAddrMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpTargetProtocolAddrMetricTag
	// validate validates PatternFlowArpTargetProtocolAddrMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpTargetProtocolAddrMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpTargetProtocolAddrMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpTargetProtocolAddrMetricTag
	SetName(value string) PatternFlowArpTargetProtocolAddrMetricTag
	// Offset returns uint32, set in PatternFlowArpTargetProtocolAddrMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpTargetProtocolAddrMetricTag
	SetOffset(value uint32) PatternFlowArpTargetProtocolAddrMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpTargetProtocolAddrMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpTargetProtocolAddrMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpTargetProtocolAddrMetricTag
	SetLength(value uint32) PatternFlowArpTargetProtocolAddrMetricTag
	// HasLength checks if Length has been set in PatternFlowArpTargetProtocolAddrMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpTargetProtocolAddrMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpTargetProtocolAddrMetricTag object
func (obj *patternFlowArpTargetProtocolAddrMetricTag) SetName(value string) PatternFlowArpTargetProtocolAddrMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpTargetProtocolAddrMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpTargetProtocolAddrMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpTargetProtocolAddrMetricTag object
func (obj *patternFlowArpTargetProtocolAddrMetricTag) SetOffset(value uint32) PatternFlowArpTargetProtocolAddrMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpTargetProtocolAddrMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpTargetProtocolAddrMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpTargetProtocolAddrMetricTag object
func (obj *patternFlowArpTargetProtocolAddrMetricTag) SetLength(value uint32) PatternFlowArpTargetProtocolAddrMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpTargetProtocolAddrMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowArpTargetProtocolAddrMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpTargetProtocolAddrMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpTargetProtocolAddrMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpTargetProtocolAddrMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
