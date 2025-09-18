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

// ***** PatternFlowArpTargetHardwareAddrMetricTag *****
type patternFlowArpTargetHardwareAddrMetricTag struct {
	validation
	obj          *otg.PatternFlowArpTargetHardwareAddrMetricTag
	marshaller   marshalPatternFlowArpTargetHardwareAddrMetricTag
	unMarshaller unMarshalPatternFlowArpTargetHardwareAddrMetricTag
}

func NewPatternFlowArpTargetHardwareAddrMetricTag() PatternFlowArpTargetHardwareAddrMetricTag {
	obj := patternFlowArpTargetHardwareAddrMetricTag{obj: &otg.PatternFlowArpTargetHardwareAddrMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) msg() *otg.PatternFlowArpTargetHardwareAddrMetricTag {
	return obj.obj
}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) setMsg(msg *otg.PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpTargetHardwareAddrMetricTag struct {
	obj *patternFlowArpTargetHardwareAddrMetricTag
}

type marshalPatternFlowArpTargetHardwareAddrMetricTag interface {
	// ToProto marshals PatternFlowArpTargetHardwareAddrMetricTag to protobuf object *otg.PatternFlowArpTargetHardwareAddrMetricTag
	ToProto() (*otg.PatternFlowArpTargetHardwareAddrMetricTag, error)
	// ToPbText marshals PatternFlowArpTargetHardwareAddrMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpTargetHardwareAddrMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpTargetHardwareAddrMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpTargetHardwareAddrMetricTag struct {
	obj *patternFlowArpTargetHardwareAddrMetricTag
}

type unMarshalPatternFlowArpTargetHardwareAddrMetricTag interface {
	// FromProto unmarshals PatternFlowArpTargetHardwareAddrMetricTag from protobuf object *otg.PatternFlowArpTargetHardwareAddrMetricTag
	FromProto(msg *otg.PatternFlowArpTargetHardwareAddrMetricTag) (PatternFlowArpTargetHardwareAddrMetricTag, error)
	// FromPbText unmarshals PatternFlowArpTargetHardwareAddrMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpTargetHardwareAddrMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpTargetHardwareAddrMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) Marshal() marshalPatternFlowArpTargetHardwareAddrMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpTargetHardwareAddrMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) Unmarshal() unMarshalPatternFlowArpTargetHardwareAddrMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpTargetHardwareAddrMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpTargetHardwareAddrMetricTag) ToProto() (*otg.PatternFlowArpTargetHardwareAddrMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpTargetHardwareAddrMetricTag) FromProto(msg *otg.PatternFlowArpTargetHardwareAddrMetricTag) (PatternFlowArpTargetHardwareAddrMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpTargetHardwareAddrMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddrMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpTargetHardwareAddrMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddrMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpTargetHardwareAddrMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddrMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpTargetHardwareAddrMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) Clone() (PatternFlowArpTargetHardwareAddrMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpTargetHardwareAddrMetricTag()
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

// PatternFlowArpTargetHardwareAddrMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpTargetHardwareAddrMetricTag interface {
	Validation
	// msg marshals PatternFlowArpTargetHardwareAddrMetricTag to protobuf object *otg.PatternFlowArpTargetHardwareAddrMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpTargetHardwareAddrMetricTag
	// setMsg unmarshals PatternFlowArpTargetHardwareAddrMetricTag from protobuf object *otg.PatternFlowArpTargetHardwareAddrMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpTargetHardwareAddrMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpTargetHardwareAddrMetricTag
	// validate validates PatternFlowArpTargetHardwareAddrMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpTargetHardwareAddrMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpTargetHardwareAddrMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpTargetHardwareAddrMetricTag
	SetName(value string) PatternFlowArpTargetHardwareAddrMetricTag
	// Offset returns uint32, set in PatternFlowArpTargetHardwareAddrMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpTargetHardwareAddrMetricTag
	SetOffset(value uint32) PatternFlowArpTargetHardwareAddrMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpTargetHardwareAddrMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpTargetHardwareAddrMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpTargetHardwareAddrMetricTag
	SetLength(value uint32) PatternFlowArpTargetHardwareAddrMetricTag
	// HasLength checks if Length has been set in PatternFlowArpTargetHardwareAddrMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpTargetHardwareAddrMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpTargetHardwareAddrMetricTag object
func (obj *patternFlowArpTargetHardwareAddrMetricTag) SetName(value string) PatternFlowArpTargetHardwareAddrMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpTargetHardwareAddrMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpTargetHardwareAddrMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpTargetHardwareAddrMetricTag object
func (obj *patternFlowArpTargetHardwareAddrMetricTag) SetOffset(value uint32) PatternFlowArpTargetHardwareAddrMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpTargetHardwareAddrMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpTargetHardwareAddrMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpTargetHardwareAddrMetricTag object
func (obj *patternFlowArpTargetHardwareAddrMetricTag) SetLength(value uint32) PatternFlowArpTargetHardwareAddrMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpTargetHardwareAddrMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowArpTargetHardwareAddrMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpTargetHardwareAddrMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpTargetHardwareAddrMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpTargetHardwareAddrMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}
