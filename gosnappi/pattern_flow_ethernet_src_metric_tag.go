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

// ***** PatternFlowEthernetSrcMetricTag *****
type patternFlowEthernetSrcMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetSrcMetricTag
	marshaller   marshalPatternFlowEthernetSrcMetricTag
	unMarshaller unMarshalPatternFlowEthernetSrcMetricTag
}

func NewPatternFlowEthernetSrcMetricTag() PatternFlowEthernetSrcMetricTag {
	obj := patternFlowEthernetSrcMetricTag{obj: &otg.PatternFlowEthernetSrcMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetSrcMetricTag) msg() *otg.PatternFlowEthernetSrcMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetSrcMetricTag) setMsg(msg *otg.PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetSrcMetricTag struct {
	obj *patternFlowEthernetSrcMetricTag
}

type marshalPatternFlowEthernetSrcMetricTag interface {
	// ToProto marshals PatternFlowEthernetSrcMetricTag to protobuf object *otg.PatternFlowEthernetSrcMetricTag
	ToProto() (*otg.PatternFlowEthernetSrcMetricTag, error)
	// ToPbText marshals PatternFlowEthernetSrcMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetSrcMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetSrcMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetSrcMetricTag struct {
	obj *patternFlowEthernetSrcMetricTag
}

type unMarshalPatternFlowEthernetSrcMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetSrcMetricTag from protobuf object *otg.PatternFlowEthernetSrcMetricTag
	FromProto(msg *otg.PatternFlowEthernetSrcMetricTag) (PatternFlowEthernetSrcMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetSrcMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetSrcMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetSrcMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetSrcMetricTag) Marshal() marshalPatternFlowEthernetSrcMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetSrcMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetSrcMetricTag) Unmarshal() unMarshalPatternFlowEthernetSrcMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetSrcMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetSrcMetricTag) ToProto() (*otg.PatternFlowEthernetSrcMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetSrcMetricTag) FromProto(msg *otg.PatternFlowEthernetSrcMetricTag) (PatternFlowEthernetSrcMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetSrcMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrcMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetSrcMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrcMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetSrcMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrcMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetSrcMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetSrcMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetSrcMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetSrcMetricTag) Clone() (PatternFlowEthernetSrcMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetSrcMetricTag()
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

// PatternFlowEthernetSrcMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetSrcMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetSrcMetricTag to protobuf object *otg.PatternFlowEthernetSrcMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetSrcMetricTag
	// setMsg unmarshals PatternFlowEthernetSrcMetricTag from protobuf object *otg.PatternFlowEthernetSrcMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetSrcMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetSrcMetricTag
	// validate validates PatternFlowEthernetSrcMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetSrcMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetSrcMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetSrcMetricTag
	SetName(value string) PatternFlowEthernetSrcMetricTag
	// Offset returns uint32, set in PatternFlowEthernetSrcMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetSrcMetricTag
	SetOffset(value uint32) PatternFlowEthernetSrcMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetSrcMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetSrcMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetSrcMetricTag
	SetLength(value uint32) PatternFlowEthernetSrcMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetSrcMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetSrcMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetSrcMetricTag object
func (obj *patternFlowEthernetSrcMetricTag) SetName(value string) PatternFlowEthernetSrcMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetSrcMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetSrcMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetSrcMetricTag object
func (obj *patternFlowEthernetSrcMetricTag) SetOffset(value uint32) PatternFlowEthernetSrcMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetSrcMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetSrcMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetSrcMetricTag object
func (obj *patternFlowEthernetSrcMetricTag) SetLength(value uint32) PatternFlowEthernetSrcMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetSrcMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetSrcMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowEthernetSrcMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetSrcMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetSrcMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetSrcMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}
