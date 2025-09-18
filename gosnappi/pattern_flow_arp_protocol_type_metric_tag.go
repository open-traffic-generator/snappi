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

// ***** PatternFlowArpProtocolTypeMetricTag *****
type patternFlowArpProtocolTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowArpProtocolTypeMetricTag
	marshaller   marshalPatternFlowArpProtocolTypeMetricTag
	unMarshaller unMarshalPatternFlowArpProtocolTypeMetricTag
}

func NewPatternFlowArpProtocolTypeMetricTag() PatternFlowArpProtocolTypeMetricTag {
	obj := patternFlowArpProtocolTypeMetricTag{obj: &otg.PatternFlowArpProtocolTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpProtocolTypeMetricTag) msg() *otg.PatternFlowArpProtocolTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowArpProtocolTypeMetricTag) setMsg(msg *otg.PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpProtocolTypeMetricTag struct {
	obj *patternFlowArpProtocolTypeMetricTag
}

type marshalPatternFlowArpProtocolTypeMetricTag interface {
	// ToProto marshals PatternFlowArpProtocolTypeMetricTag to protobuf object *otg.PatternFlowArpProtocolTypeMetricTag
	ToProto() (*otg.PatternFlowArpProtocolTypeMetricTag, error)
	// ToPbText marshals PatternFlowArpProtocolTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpProtocolTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpProtocolTypeMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpProtocolTypeMetricTag struct {
	obj *patternFlowArpProtocolTypeMetricTag
}

type unMarshalPatternFlowArpProtocolTypeMetricTag interface {
	// FromProto unmarshals PatternFlowArpProtocolTypeMetricTag from protobuf object *otg.PatternFlowArpProtocolTypeMetricTag
	FromProto(msg *otg.PatternFlowArpProtocolTypeMetricTag) (PatternFlowArpProtocolTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowArpProtocolTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpProtocolTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpProtocolTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpProtocolTypeMetricTag) Marshal() marshalPatternFlowArpProtocolTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpProtocolTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpProtocolTypeMetricTag) Unmarshal() unMarshalPatternFlowArpProtocolTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpProtocolTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpProtocolTypeMetricTag) ToProto() (*otg.PatternFlowArpProtocolTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpProtocolTypeMetricTag) FromProto(msg *otg.PatternFlowArpProtocolTypeMetricTag) (PatternFlowArpProtocolTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpProtocolTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpProtocolTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpProtocolTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowArpProtocolTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpProtocolTypeMetricTag) Clone() (PatternFlowArpProtocolTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpProtocolTypeMetricTag()
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

// PatternFlowArpProtocolTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowArpProtocolTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowArpProtocolTypeMetricTag to protobuf object *otg.PatternFlowArpProtocolTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowArpProtocolTypeMetricTag
	// setMsg unmarshals PatternFlowArpProtocolTypeMetricTag from protobuf object *otg.PatternFlowArpProtocolTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowArpProtocolTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpProtocolTypeMetricTag
	// validate validates PatternFlowArpProtocolTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpProtocolTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowArpProtocolTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowArpProtocolTypeMetricTag
	SetName(value string) PatternFlowArpProtocolTypeMetricTag
	// Offset returns uint32, set in PatternFlowArpProtocolTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowArpProtocolTypeMetricTag
	SetOffset(value uint32) PatternFlowArpProtocolTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowArpProtocolTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowArpProtocolTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowArpProtocolTypeMetricTag
	SetLength(value uint32) PatternFlowArpProtocolTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowArpProtocolTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowArpProtocolTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowArpProtocolTypeMetricTag object
func (obj *patternFlowArpProtocolTypeMetricTag) SetName(value string) PatternFlowArpProtocolTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpProtocolTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowArpProtocolTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowArpProtocolTypeMetricTag object
func (obj *patternFlowArpProtocolTypeMetricTag) SetOffset(value uint32) PatternFlowArpProtocolTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpProtocolTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowArpProtocolTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowArpProtocolTypeMetricTag object
func (obj *patternFlowArpProtocolTypeMetricTag) SetLength(value uint32) PatternFlowArpProtocolTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowArpProtocolTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowArpProtocolTypeMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowArpProtocolTypeMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolTypeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowArpProtocolTypeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowArpProtocolTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
