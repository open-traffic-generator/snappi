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

// ***** PatternFlowVxlanVniMetricTag *****
type patternFlowVxlanVniMetricTag struct {
	validation
	obj          *otg.PatternFlowVxlanVniMetricTag
	marshaller   marshalPatternFlowVxlanVniMetricTag
	unMarshaller unMarshalPatternFlowVxlanVniMetricTag
}

func NewPatternFlowVxlanVniMetricTag() PatternFlowVxlanVniMetricTag {
	obj := patternFlowVxlanVniMetricTag{obj: &otg.PatternFlowVxlanVniMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanVniMetricTag) msg() *otg.PatternFlowVxlanVniMetricTag {
	return obj.obj
}

func (obj *patternFlowVxlanVniMetricTag) setMsg(msg *otg.PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanVniMetricTag struct {
	obj *patternFlowVxlanVniMetricTag
}

type marshalPatternFlowVxlanVniMetricTag interface {
	// ToProto marshals PatternFlowVxlanVniMetricTag to protobuf object *otg.PatternFlowVxlanVniMetricTag
	ToProto() (*otg.PatternFlowVxlanVniMetricTag, error)
	// ToPbText marshals PatternFlowVxlanVniMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanVniMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanVniMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVxlanVniMetricTag struct {
	obj *patternFlowVxlanVniMetricTag
}

type unMarshalPatternFlowVxlanVniMetricTag interface {
	// FromProto unmarshals PatternFlowVxlanVniMetricTag from protobuf object *otg.PatternFlowVxlanVniMetricTag
	FromProto(msg *otg.PatternFlowVxlanVniMetricTag) (PatternFlowVxlanVniMetricTag, error)
	// FromPbText unmarshals PatternFlowVxlanVniMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanVniMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanVniMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanVniMetricTag) Marshal() marshalPatternFlowVxlanVniMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanVniMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanVniMetricTag) Unmarshal() unMarshalPatternFlowVxlanVniMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanVniMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanVniMetricTag) ToProto() (*otg.PatternFlowVxlanVniMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanVniMetricTag) FromProto(msg *otg.PatternFlowVxlanVniMetricTag) (PatternFlowVxlanVniMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanVniMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVniMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanVniMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVniMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanVniMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVniMetricTag) FromJson(value string) error {
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

func (obj *patternFlowVxlanVniMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanVniMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanVniMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanVniMetricTag) Clone() (PatternFlowVxlanVniMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanVniMetricTag()
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

// PatternFlowVxlanVniMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowVxlanVniMetricTag interface {
	Validation
	// msg marshals PatternFlowVxlanVniMetricTag to protobuf object *otg.PatternFlowVxlanVniMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanVniMetricTag
	// setMsg unmarshals PatternFlowVxlanVniMetricTag from protobuf object *otg.PatternFlowVxlanVniMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanVniMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanVniMetricTag
	// validate validates PatternFlowVxlanVniMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanVniMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowVxlanVniMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowVxlanVniMetricTag
	SetName(value string) PatternFlowVxlanVniMetricTag
	// Offset returns uint32, set in PatternFlowVxlanVniMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowVxlanVniMetricTag
	SetOffset(value uint32) PatternFlowVxlanVniMetricTag
	// HasOffset checks if Offset has been set in PatternFlowVxlanVniMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowVxlanVniMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowVxlanVniMetricTag
	SetLength(value uint32) PatternFlowVxlanVniMetricTag
	// HasLength checks if Length has been set in PatternFlowVxlanVniMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowVxlanVniMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowVxlanVniMetricTag object
func (obj *patternFlowVxlanVniMetricTag) SetName(value string) PatternFlowVxlanVniMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVxlanVniMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowVxlanVniMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowVxlanVniMetricTag object
func (obj *patternFlowVxlanVniMetricTag) SetOffset(value uint32) PatternFlowVxlanVniMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVxlanVniMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowVxlanVniMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowVxlanVniMetricTag object
func (obj *patternFlowVxlanVniMetricTag) SetLength(value uint32) PatternFlowVxlanVniMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowVxlanVniMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowVxlanVniMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowVxlanVniMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 23 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanVniMetricTag.Offset <= 23 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 24 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowVxlanVniMetricTag.Length <= 24 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowVxlanVniMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(24)
	}

}
