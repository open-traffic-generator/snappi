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

// ***** PatternFlowEthernetEtherTypeMetricTag *****
type patternFlowEthernetEtherTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetEtherTypeMetricTag
	marshaller   marshalPatternFlowEthernetEtherTypeMetricTag
	unMarshaller unMarshalPatternFlowEthernetEtherTypeMetricTag
}

func NewPatternFlowEthernetEtherTypeMetricTag() PatternFlowEthernetEtherTypeMetricTag {
	obj := patternFlowEthernetEtherTypeMetricTag{obj: &otg.PatternFlowEthernetEtherTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetEtherTypeMetricTag) msg() *otg.PatternFlowEthernetEtherTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetEtherTypeMetricTag) setMsg(msg *otg.PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetEtherTypeMetricTag struct {
	obj *patternFlowEthernetEtherTypeMetricTag
}

type marshalPatternFlowEthernetEtherTypeMetricTag interface {
	// ToProto marshals PatternFlowEthernetEtherTypeMetricTag to protobuf object *otg.PatternFlowEthernetEtherTypeMetricTag
	ToProto() (*otg.PatternFlowEthernetEtherTypeMetricTag, error)
	// ToPbText marshals PatternFlowEthernetEtherTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetEtherTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetEtherTypeMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetEtherTypeMetricTag struct {
	obj *patternFlowEthernetEtherTypeMetricTag
}

type unMarshalPatternFlowEthernetEtherTypeMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetEtherTypeMetricTag from protobuf object *otg.PatternFlowEthernetEtherTypeMetricTag
	FromProto(msg *otg.PatternFlowEthernetEtherTypeMetricTag) (PatternFlowEthernetEtherTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetEtherTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetEtherTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetEtherTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetEtherTypeMetricTag) Marshal() marshalPatternFlowEthernetEtherTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetEtherTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetEtherTypeMetricTag) Unmarshal() unMarshalPatternFlowEthernetEtherTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetEtherTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetEtherTypeMetricTag) ToProto() (*otg.PatternFlowEthernetEtherTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetEtherTypeMetricTag) FromProto(msg *otg.PatternFlowEthernetEtherTypeMetricTag) (PatternFlowEthernetEtherTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetEtherTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetEtherTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetEtherTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetEtherTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetEtherTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetEtherTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetEtherTypeMetricTag) Clone() (PatternFlowEthernetEtherTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetEtherTypeMetricTag()
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

// PatternFlowEthernetEtherTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetEtherTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetEtherTypeMetricTag to protobuf object *otg.PatternFlowEthernetEtherTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetEtherTypeMetricTag
	// setMsg unmarshals PatternFlowEthernetEtherTypeMetricTag from protobuf object *otg.PatternFlowEthernetEtherTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetEtherTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetEtherTypeMetricTag
	// validate validates PatternFlowEthernetEtherTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetEtherTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetEtherTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetEtherTypeMetricTag
	SetName(value string) PatternFlowEthernetEtherTypeMetricTag
	// Offset returns uint32, set in PatternFlowEthernetEtherTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetEtherTypeMetricTag
	SetOffset(value uint32) PatternFlowEthernetEtherTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetEtherTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetEtherTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetEtherTypeMetricTag
	SetLength(value uint32) PatternFlowEthernetEtherTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetEtherTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetEtherTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetEtherTypeMetricTag object
func (obj *patternFlowEthernetEtherTypeMetricTag) SetName(value string) PatternFlowEthernetEtherTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetEtherTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetEtherTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetEtherTypeMetricTag object
func (obj *patternFlowEthernetEtherTypeMetricTag) SetOffset(value uint32) PatternFlowEthernetEtherTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetEtherTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetEtherTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetEtherTypeMetricTag object
func (obj *patternFlowEthernetEtherTypeMetricTag) SetLength(value uint32) PatternFlowEthernetEtherTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetEtherTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetEtherTypeMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowEthernetEtherTypeMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherTypeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetEtherTypeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetEtherTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
