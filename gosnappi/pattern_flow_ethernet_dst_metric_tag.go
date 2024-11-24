package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetDstMetricTag *****
type patternFlowEthernetDstMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetDstMetricTag
	marshaller   marshalPatternFlowEthernetDstMetricTag
	unMarshaller unMarshalPatternFlowEthernetDstMetricTag
}

func NewPatternFlowEthernetDstMetricTag() PatternFlowEthernetDstMetricTag {
	obj := patternFlowEthernetDstMetricTag{obj: &otg.PatternFlowEthernetDstMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetDstMetricTag) msg() *otg.PatternFlowEthernetDstMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetDstMetricTag) setMsg(msg *otg.PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetDstMetricTag struct {
	obj *patternFlowEthernetDstMetricTag
}

type marshalPatternFlowEthernetDstMetricTag interface {
	// ToProto marshals PatternFlowEthernetDstMetricTag to protobuf object *otg.PatternFlowEthernetDstMetricTag
	ToProto() (*otg.PatternFlowEthernetDstMetricTag, error)
	// ToPbText marshals PatternFlowEthernetDstMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetDstMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetDstMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetDstMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetDstMetricTag struct {
	obj *patternFlowEthernetDstMetricTag
}

type unMarshalPatternFlowEthernetDstMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetDstMetricTag from protobuf object *otg.PatternFlowEthernetDstMetricTag
	FromProto(msg *otg.PatternFlowEthernetDstMetricTag) (PatternFlowEthernetDstMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetDstMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetDstMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetDstMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetDstMetricTag) Marshal() marshalPatternFlowEthernetDstMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetDstMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetDstMetricTag) Unmarshal() unMarshalPatternFlowEthernetDstMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetDstMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetDstMetricTag) ToProto() (*otg.PatternFlowEthernetDstMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetDstMetricTag) FromProto(msg *otg.PatternFlowEthernetDstMetricTag) (PatternFlowEthernetDstMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetDstMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDstMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetDstMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDstMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetDstMetricTag) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowEthernetDstMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDstMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetDstMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetDstMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetDstMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetDstMetricTag) Clone() (PatternFlowEthernetDstMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetDstMetricTag()
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

// PatternFlowEthernetDstMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetDstMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetDstMetricTag to protobuf object *otg.PatternFlowEthernetDstMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetDstMetricTag
	// setMsg unmarshals PatternFlowEthernetDstMetricTag from protobuf object *otg.PatternFlowEthernetDstMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetDstMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetDstMetricTag
	// validate validates PatternFlowEthernetDstMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetDstMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetDstMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetDstMetricTag
	SetName(value string) PatternFlowEthernetDstMetricTag
	// Offset returns uint32, set in PatternFlowEthernetDstMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetDstMetricTag
	SetOffset(value uint32) PatternFlowEthernetDstMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetDstMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetDstMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetDstMetricTag
	SetLength(value uint32) PatternFlowEthernetDstMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetDstMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetDstMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetDstMetricTag object
func (obj *patternFlowEthernetDstMetricTag) SetName(value string) PatternFlowEthernetDstMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetDstMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetDstMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetDstMetricTag object
func (obj *patternFlowEthernetDstMetricTag) SetOffset(value uint32) PatternFlowEthernetDstMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetDstMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetDstMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetDstMetricTag object
func (obj *patternFlowEthernetDstMetricTag) SetLength(value uint32) PatternFlowEthernetDstMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetDstMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetDstMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetDstMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetDstMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetDstMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}
