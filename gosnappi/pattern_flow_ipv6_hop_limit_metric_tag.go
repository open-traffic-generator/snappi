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

// ***** PatternFlowIpv6HopLimitMetricTag *****
type patternFlowIpv6HopLimitMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv6HopLimitMetricTag
	marshaller   marshalPatternFlowIpv6HopLimitMetricTag
	unMarshaller unMarshalPatternFlowIpv6HopLimitMetricTag
}

func NewPatternFlowIpv6HopLimitMetricTag() PatternFlowIpv6HopLimitMetricTag {
	obj := patternFlowIpv6HopLimitMetricTag{obj: &otg.PatternFlowIpv6HopLimitMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6HopLimitMetricTag) msg() *otg.PatternFlowIpv6HopLimitMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv6HopLimitMetricTag) setMsg(msg *otg.PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6HopLimitMetricTag struct {
	obj *patternFlowIpv6HopLimitMetricTag
}

type marshalPatternFlowIpv6HopLimitMetricTag interface {
	// ToProto marshals PatternFlowIpv6HopLimitMetricTag to protobuf object *otg.PatternFlowIpv6HopLimitMetricTag
	ToProto() (*otg.PatternFlowIpv6HopLimitMetricTag, error)
	// ToPbText marshals PatternFlowIpv6HopLimitMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6HopLimitMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6HopLimitMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6HopLimitMetricTag struct {
	obj *patternFlowIpv6HopLimitMetricTag
}

type unMarshalPatternFlowIpv6HopLimitMetricTag interface {
	// FromProto unmarshals PatternFlowIpv6HopLimitMetricTag from protobuf object *otg.PatternFlowIpv6HopLimitMetricTag
	FromProto(msg *otg.PatternFlowIpv6HopLimitMetricTag) (PatternFlowIpv6HopLimitMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv6HopLimitMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6HopLimitMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6HopLimitMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6HopLimitMetricTag) Marshal() marshalPatternFlowIpv6HopLimitMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6HopLimitMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6HopLimitMetricTag) Unmarshal() unMarshalPatternFlowIpv6HopLimitMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6HopLimitMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6HopLimitMetricTag) ToProto() (*otg.PatternFlowIpv6HopLimitMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6HopLimitMetricTag) FromProto(msg *otg.PatternFlowIpv6HopLimitMetricTag) (PatternFlowIpv6HopLimitMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6HopLimitMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimitMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6HopLimitMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimitMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6HopLimitMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimitMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6HopLimitMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6HopLimitMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6HopLimitMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6HopLimitMetricTag) Clone() (PatternFlowIpv6HopLimitMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6HopLimitMetricTag()
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

// PatternFlowIpv6HopLimitMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv6HopLimitMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv6HopLimitMetricTag to protobuf object *otg.PatternFlowIpv6HopLimitMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6HopLimitMetricTag
	// setMsg unmarshals PatternFlowIpv6HopLimitMetricTag from protobuf object *otg.PatternFlowIpv6HopLimitMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6HopLimitMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6HopLimitMetricTag
	// validate validates PatternFlowIpv6HopLimitMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6HopLimitMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv6HopLimitMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv6HopLimitMetricTag
	SetName(value string) PatternFlowIpv6HopLimitMetricTag
	// Offset returns uint32, set in PatternFlowIpv6HopLimitMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv6HopLimitMetricTag
	SetOffset(value uint32) PatternFlowIpv6HopLimitMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv6HopLimitMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv6HopLimitMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv6HopLimitMetricTag
	SetLength(value uint32) PatternFlowIpv6HopLimitMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv6HopLimitMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv6HopLimitMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv6HopLimitMetricTag object
func (obj *patternFlowIpv6HopLimitMetricTag) SetName(value string) PatternFlowIpv6HopLimitMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6HopLimitMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv6HopLimitMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv6HopLimitMetricTag object
func (obj *patternFlowIpv6HopLimitMetricTag) SetOffset(value uint32) PatternFlowIpv6HopLimitMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6HopLimitMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv6HopLimitMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv6HopLimitMetricTag object
func (obj *patternFlowIpv6HopLimitMetricTag) SetLength(value uint32) PatternFlowIpv6HopLimitMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv6HopLimitMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv6HopLimitMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowIpv6HopLimitMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6HopLimitMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv6HopLimitMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv6HopLimitMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
