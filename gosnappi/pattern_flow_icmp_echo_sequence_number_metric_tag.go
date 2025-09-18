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

// ***** PatternFlowIcmpEchoSequenceNumberMetricTag *****
type patternFlowIcmpEchoSequenceNumberMetricTag struct {
	validation
	obj          *otg.PatternFlowIcmpEchoSequenceNumberMetricTag
	marshaller   marshalPatternFlowIcmpEchoSequenceNumberMetricTag
	unMarshaller unMarshalPatternFlowIcmpEchoSequenceNumberMetricTag
}

func NewPatternFlowIcmpEchoSequenceNumberMetricTag() PatternFlowIcmpEchoSequenceNumberMetricTag {
	obj := patternFlowIcmpEchoSequenceNumberMetricTag{obj: &otg.PatternFlowIcmpEchoSequenceNumberMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) msg() *otg.PatternFlowIcmpEchoSequenceNumberMetricTag {
	return obj.obj
}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) setMsg(msg *otg.PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoSequenceNumberMetricTag struct {
	obj *patternFlowIcmpEchoSequenceNumberMetricTag
}

type marshalPatternFlowIcmpEchoSequenceNumberMetricTag interface {
	// ToProto marshals PatternFlowIcmpEchoSequenceNumberMetricTag to protobuf object *otg.PatternFlowIcmpEchoSequenceNumberMetricTag
	ToProto() (*otg.PatternFlowIcmpEchoSequenceNumberMetricTag, error)
	// ToPbText marshals PatternFlowIcmpEchoSequenceNumberMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoSequenceNumberMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoSequenceNumberMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpEchoSequenceNumberMetricTag struct {
	obj *patternFlowIcmpEchoSequenceNumberMetricTag
}

type unMarshalPatternFlowIcmpEchoSequenceNumberMetricTag interface {
	// FromProto unmarshals PatternFlowIcmpEchoSequenceNumberMetricTag from protobuf object *otg.PatternFlowIcmpEchoSequenceNumberMetricTag
	FromProto(msg *otg.PatternFlowIcmpEchoSequenceNumberMetricTag) (PatternFlowIcmpEchoSequenceNumberMetricTag, error)
	// FromPbText unmarshals PatternFlowIcmpEchoSequenceNumberMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoSequenceNumberMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoSequenceNumberMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) Marshal() marshalPatternFlowIcmpEchoSequenceNumberMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoSequenceNumberMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) Unmarshal() unMarshalPatternFlowIcmpEchoSequenceNumberMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoSequenceNumberMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoSequenceNumberMetricTag) ToProto() (*otg.PatternFlowIcmpEchoSequenceNumberMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoSequenceNumberMetricTag) FromProto(msg *otg.PatternFlowIcmpEchoSequenceNumberMetricTag) (PatternFlowIcmpEchoSequenceNumberMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoSequenceNumberMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumberMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoSequenceNumberMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumberMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoSequenceNumberMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumberMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) Clone() (PatternFlowIcmpEchoSequenceNumberMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoSequenceNumberMetricTag()
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

// PatternFlowIcmpEchoSequenceNumberMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIcmpEchoSequenceNumberMetricTag interface {
	Validation
	// msg marshals PatternFlowIcmpEchoSequenceNumberMetricTag to protobuf object *otg.PatternFlowIcmpEchoSequenceNumberMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoSequenceNumberMetricTag
	// setMsg unmarshals PatternFlowIcmpEchoSequenceNumberMetricTag from protobuf object *otg.PatternFlowIcmpEchoSequenceNumberMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoSequenceNumberMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoSequenceNumberMetricTag
	// validate validates PatternFlowIcmpEchoSequenceNumberMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoSequenceNumberMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIcmpEchoSequenceNumberMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIcmpEchoSequenceNumberMetricTag
	SetName(value string) PatternFlowIcmpEchoSequenceNumberMetricTag
	// Offset returns uint32, set in PatternFlowIcmpEchoSequenceNumberMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIcmpEchoSequenceNumberMetricTag
	SetOffset(value uint32) PatternFlowIcmpEchoSequenceNumberMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIcmpEchoSequenceNumberMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIcmpEchoSequenceNumberMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIcmpEchoSequenceNumberMetricTag
	SetLength(value uint32) PatternFlowIcmpEchoSequenceNumberMetricTag
	// HasLength checks if Length has been set in PatternFlowIcmpEchoSequenceNumberMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIcmpEchoSequenceNumberMetricTag object
func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) SetName(value string) PatternFlowIcmpEchoSequenceNumberMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIcmpEchoSequenceNumberMetricTag object
func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) SetOffset(value uint32) PatternFlowIcmpEchoSequenceNumberMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIcmpEchoSequenceNumberMetricTag object
func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) SetLength(value uint32) PatternFlowIcmpEchoSequenceNumberMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIcmpEchoSequenceNumberMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowIcmpEchoSequenceNumberMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumberMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIcmpEchoSequenceNumberMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIcmpEchoSequenceNumberMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
