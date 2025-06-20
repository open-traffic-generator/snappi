package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpDstPortMetricTag *****
type patternFlowUdpDstPortMetricTag struct {
	validation
	obj          *otg.PatternFlowUdpDstPortMetricTag
	marshaller   marshalPatternFlowUdpDstPortMetricTag
	unMarshaller unMarshalPatternFlowUdpDstPortMetricTag
}

func NewPatternFlowUdpDstPortMetricTag() PatternFlowUdpDstPortMetricTag {
	obj := patternFlowUdpDstPortMetricTag{obj: &otg.PatternFlowUdpDstPortMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpDstPortMetricTag) msg() *otg.PatternFlowUdpDstPortMetricTag {
	return obj.obj
}

func (obj *patternFlowUdpDstPortMetricTag) setMsg(msg *otg.PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpDstPortMetricTag struct {
	obj *patternFlowUdpDstPortMetricTag
}

type marshalPatternFlowUdpDstPortMetricTag interface {
	// ToProto marshals PatternFlowUdpDstPortMetricTag to protobuf object *otg.PatternFlowUdpDstPortMetricTag
	ToProto() (*otg.PatternFlowUdpDstPortMetricTag, error)
	// ToPbText marshals PatternFlowUdpDstPortMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpDstPortMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpDstPortMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowUdpDstPortMetricTag struct {
	obj *patternFlowUdpDstPortMetricTag
}

type unMarshalPatternFlowUdpDstPortMetricTag interface {
	// FromProto unmarshals PatternFlowUdpDstPortMetricTag from protobuf object *otg.PatternFlowUdpDstPortMetricTag
	FromProto(msg *otg.PatternFlowUdpDstPortMetricTag) (PatternFlowUdpDstPortMetricTag, error)
	// FromPbText unmarshals PatternFlowUdpDstPortMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpDstPortMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpDstPortMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpDstPortMetricTag) Marshal() marshalPatternFlowUdpDstPortMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpDstPortMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpDstPortMetricTag) Unmarshal() unMarshalPatternFlowUdpDstPortMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpDstPortMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpDstPortMetricTag) ToProto() (*otg.PatternFlowUdpDstPortMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpDstPortMetricTag) FromProto(msg *otg.PatternFlowUdpDstPortMetricTag) (PatternFlowUdpDstPortMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpDstPortMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpDstPortMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpDstPortMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortMetricTag) FromJson(value string) error {
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

func (obj *patternFlowUdpDstPortMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpDstPortMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpDstPortMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpDstPortMetricTag) Clone() (PatternFlowUdpDstPortMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpDstPortMetricTag()
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

// PatternFlowUdpDstPortMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowUdpDstPortMetricTag interface {
	Validation
	// msg marshals PatternFlowUdpDstPortMetricTag to protobuf object *otg.PatternFlowUdpDstPortMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpDstPortMetricTag
	// setMsg unmarshals PatternFlowUdpDstPortMetricTag from protobuf object *otg.PatternFlowUdpDstPortMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowUdpDstPortMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpDstPortMetricTag
	// validate validates PatternFlowUdpDstPortMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpDstPortMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowUdpDstPortMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowUdpDstPortMetricTag
	SetName(value string) PatternFlowUdpDstPortMetricTag
	// Offset returns uint32, set in PatternFlowUdpDstPortMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowUdpDstPortMetricTag
	SetOffset(value uint32) PatternFlowUdpDstPortMetricTag
	// HasOffset checks if Offset has been set in PatternFlowUdpDstPortMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowUdpDstPortMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowUdpDstPortMetricTag
	SetLength(value uint32) PatternFlowUdpDstPortMetricTag
	// HasLength checks if Length has been set in PatternFlowUdpDstPortMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowUdpDstPortMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowUdpDstPortMetricTag object
func (obj *patternFlowUdpDstPortMetricTag) SetName(value string) PatternFlowUdpDstPortMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowUdpDstPortMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowUdpDstPortMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowUdpDstPortMetricTag object
func (obj *patternFlowUdpDstPortMetricTag) SetOffset(value uint32) PatternFlowUdpDstPortMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowUdpDstPortMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowUdpDstPortMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowUdpDstPortMetricTag object
func (obj *patternFlowUdpDstPortMetricTag) SetLength(value uint32) PatternFlowUdpDstPortMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowUdpDstPortMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowUdpDstPortMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPortMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowUdpDstPortMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowUdpDstPortMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
