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

// ***** PatternFlowVxlanReserved0MetricTag *****
type patternFlowVxlanReserved0MetricTag struct {
	validation
	obj          *otg.PatternFlowVxlanReserved0MetricTag
	marshaller   marshalPatternFlowVxlanReserved0MetricTag
	unMarshaller unMarshalPatternFlowVxlanReserved0MetricTag
}

func NewPatternFlowVxlanReserved0MetricTag() PatternFlowVxlanReserved0MetricTag {
	obj := patternFlowVxlanReserved0MetricTag{obj: &otg.PatternFlowVxlanReserved0MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanReserved0MetricTag) msg() *otg.PatternFlowVxlanReserved0MetricTag {
	return obj.obj
}

func (obj *patternFlowVxlanReserved0MetricTag) setMsg(msg *otg.PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanReserved0MetricTag struct {
	obj *patternFlowVxlanReserved0MetricTag
}

type marshalPatternFlowVxlanReserved0MetricTag interface {
	// ToProto marshals PatternFlowVxlanReserved0MetricTag to protobuf object *otg.PatternFlowVxlanReserved0MetricTag
	ToProto() (*otg.PatternFlowVxlanReserved0MetricTag, error)
	// ToPbText marshals PatternFlowVxlanReserved0MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanReserved0MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanReserved0MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVxlanReserved0MetricTag struct {
	obj *patternFlowVxlanReserved0MetricTag
}

type unMarshalPatternFlowVxlanReserved0MetricTag interface {
	// FromProto unmarshals PatternFlowVxlanReserved0MetricTag from protobuf object *otg.PatternFlowVxlanReserved0MetricTag
	FromProto(msg *otg.PatternFlowVxlanReserved0MetricTag) (PatternFlowVxlanReserved0MetricTag, error)
	// FromPbText unmarshals PatternFlowVxlanReserved0MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanReserved0MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanReserved0MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanReserved0MetricTag) Marshal() marshalPatternFlowVxlanReserved0MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanReserved0MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanReserved0MetricTag) Unmarshal() unMarshalPatternFlowVxlanReserved0MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanReserved0MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanReserved0MetricTag) ToProto() (*otg.PatternFlowVxlanReserved0MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanReserved0MetricTag) FromProto(msg *otg.PatternFlowVxlanReserved0MetricTag) (PatternFlowVxlanReserved0MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanReserved0MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanReserved0MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanReserved0MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0MetricTag) FromJson(value string) error {
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

func (obj *patternFlowVxlanReserved0MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved0MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved0MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanReserved0MetricTag) Clone() (PatternFlowVxlanReserved0MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanReserved0MetricTag()
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

// PatternFlowVxlanReserved0MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowVxlanReserved0MetricTag interface {
	Validation
	// msg marshals PatternFlowVxlanReserved0MetricTag to protobuf object *otg.PatternFlowVxlanReserved0MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanReserved0MetricTag
	// setMsg unmarshals PatternFlowVxlanReserved0MetricTag from protobuf object *otg.PatternFlowVxlanReserved0MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanReserved0MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanReserved0MetricTag
	// validate validates PatternFlowVxlanReserved0MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanReserved0MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowVxlanReserved0MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowVxlanReserved0MetricTag
	SetName(value string) PatternFlowVxlanReserved0MetricTag
	// Offset returns uint32, set in PatternFlowVxlanReserved0MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowVxlanReserved0MetricTag
	SetOffset(value uint32) PatternFlowVxlanReserved0MetricTag
	// HasOffset checks if Offset has been set in PatternFlowVxlanReserved0MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowVxlanReserved0MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowVxlanReserved0MetricTag
	SetLength(value uint32) PatternFlowVxlanReserved0MetricTag
	// HasLength checks if Length has been set in PatternFlowVxlanReserved0MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowVxlanReserved0MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowVxlanReserved0MetricTag object
func (obj *patternFlowVxlanReserved0MetricTag) SetName(value string) PatternFlowVxlanReserved0MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVxlanReserved0MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVxlanReserved0MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowVxlanReserved0MetricTag object
func (obj *patternFlowVxlanReserved0MetricTag) SetOffset(value uint32) PatternFlowVxlanReserved0MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVxlanReserved0MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVxlanReserved0MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowVxlanReserved0MetricTag object
func (obj *patternFlowVxlanReserved0MetricTag) SetLength(value uint32) PatternFlowVxlanReserved0MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowVxlanReserved0MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowVxlanReserved0MetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowVxlanReserved0MetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 23 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved0MetricTag.Offset <= 23 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 24 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowVxlanReserved0MetricTag.Length <= 24 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowVxlanReserved0MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(24)
	}

}
