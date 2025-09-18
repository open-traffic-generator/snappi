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

// ***** PatternFlowIgmpv1UnusedMetricTag *****
type patternFlowIgmpv1UnusedMetricTag struct {
	validation
	obj          *otg.PatternFlowIgmpv1UnusedMetricTag
	marshaller   marshalPatternFlowIgmpv1UnusedMetricTag
	unMarshaller unMarshalPatternFlowIgmpv1UnusedMetricTag
}

func NewPatternFlowIgmpv1UnusedMetricTag() PatternFlowIgmpv1UnusedMetricTag {
	obj := patternFlowIgmpv1UnusedMetricTag{obj: &otg.PatternFlowIgmpv1UnusedMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1UnusedMetricTag) msg() *otg.PatternFlowIgmpv1UnusedMetricTag {
	return obj.obj
}

func (obj *patternFlowIgmpv1UnusedMetricTag) setMsg(msg *otg.PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1UnusedMetricTag struct {
	obj *patternFlowIgmpv1UnusedMetricTag
}

type marshalPatternFlowIgmpv1UnusedMetricTag interface {
	// ToProto marshals PatternFlowIgmpv1UnusedMetricTag to protobuf object *otg.PatternFlowIgmpv1UnusedMetricTag
	ToProto() (*otg.PatternFlowIgmpv1UnusedMetricTag, error)
	// ToPbText marshals PatternFlowIgmpv1UnusedMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1UnusedMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1UnusedMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv1UnusedMetricTag struct {
	obj *patternFlowIgmpv1UnusedMetricTag
}

type unMarshalPatternFlowIgmpv1UnusedMetricTag interface {
	// FromProto unmarshals PatternFlowIgmpv1UnusedMetricTag from protobuf object *otg.PatternFlowIgmpv1UnusedMetricTag
	FromProto(msg *otg.PatternFlowIgmpv1UnusedMetricTag) (PatternFlowIgmpv1UnusedMetricTag, error)
	// FromPbText unmarshals PatternFlowIgmpv1UnusedMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1UnusedMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1UnusedMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1UnusedMetricTag) Marshal() marshalPatternFlowIgmpv1UnusedMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1UnusedMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1UnusedMetricTag) Unmarshal() unMarshalPatternFlowIgmpv1UnusedMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1UnusedMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1UnusedMetricTag) ToProto() (*otg.PatternFlowIgmpv1UnusedMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1UnusedMetricTag) FromProto(msg *otg.PatternFlowIgmpv1UnusedMetricTag) (PatternFlowIgmpv1UnusedMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1UnusedMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1UnusedMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1UnusedMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1UnusedMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1UnusedMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1UnusedMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1UnusedMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1UnusedMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1UnusedMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1UnusedMetricTag) Clone() (PatternFlowIgmpv1UnusedMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1UnusedMetricTag()
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

// PatternFlowIgmpv1UnusedMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIgmpv1UnusedMetricTag interface {
	Validation
	// msg marshals PatternFlowIgmpv1UnusedMetricTag to protobuf object *otg.PatternFlowIgmpv1UnusedMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1UnusedMetricTag
	// setMsg unmarshals PatternFlowIgmpv1UnusedMetricTag from protobuf object *otg.PatternFlowIgmpv1UnusedMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1UnusedMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1UnusedMetricTag
	// validate validates PatternFlowIgmpv1UnusedMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1UnusedMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIgmpv1UnusedMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIgmpv1UnusedMetricTag
	SetName(value string) PatternFlowIgmpv1UnusedMetricTag
	// Offset returns uint32, set in PatternFlowIgmpv1UnusedMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIgmpv1UnusedMetricTag
	SetOffset(value uint32) PatternFlowIgmpv1UnusedMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIgmpv1UnusedMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIgmpv1UnusedMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIgmpv1UnusedMetricTag
	SetLength(value uint32) PatternFlowIgmpv1UnusedMetricTag
	// HasLength checks if Length has been set in PatternFlowIgmpv1UnusedMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIgmpv1UnusedMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIgmpv1UnusedMetricTag object
func (obj *patternFlowIgmpv1UnusedMetricTag) SetName(value string) PatternFlowIgmpv1UnusedMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIgmpv1UnusedMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIgmpv1UnusedMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIgmpv1UnusedMetricTag object
func (obj *patternFlowIgmpv1UnusedMetricTag) SetOffset(value uint32) PatternFlowIgmpv1UnusedMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIgmpv1UnusedMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIgmpv1UnusedMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIgmpv1UnusedMetricTag object
func (obj *patternFlowIgmpv1UnusedMetricTag) SetLength(value uint32) PatternFlowIgmpv1UnusedMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIgmpv1UnusedMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIgmpv1UnusedMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowIgmpv1UnusedMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1UnusedMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIgmpv1UnusedMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIgmpv1UnusedMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
