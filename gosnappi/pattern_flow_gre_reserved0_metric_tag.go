package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreReserved0MetricTag *****
type patternFlowGreReserved0MetricTag struct {
	validation
	obj          *otg.PatternFlowGreReserved0MetricTag
	marshaller   marshalPatternFlowGreReserved0MetricTag
	unMarshaller unMarshalPatternFlowGreReserved0MetricTag
}

func NewPatternFlowGreReserved0MetricTag() PatternFlowGreReserved0MetricTag {
	obj := patternFlowGreReserved0MetricTag{obj: &otg.PatternFlowGreReserved0MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreReserved0MetricTag) msg() *otg.PatternFlowGreReserved0MetricTag {
	return obj.obj
}

func (obj *patternFlowGreReserved0MetricTag) setMsg(msg *otg.PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreReserved0MetricTag struct {
	obj *patternFlowGreReserved0MetricTag
}

type marshalPatternFlowGreReserved0MetricTag interface {
	// ToProto marshals PatternFlowGreReserved0MetricTag to protobuf object *otg.PatternFlowGreReserved0MetricTag
	ToProto() (*otg.PatternFlowGreReserved0MetricTag, error)
	// ToPbText marshals PatternFlowGreReserved0MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreReserved0MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreReserved0MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreReserved0MetricTag struct {
	obj *patternFlowGreReserved0MetricTag
}

type unMarshalPatternFlowGreReserved0MetricTag interface {
	// FromProto unmarshals PatternFlowGreReserved0MetricTag from protobuf object *otg.PatternFlowGreReserved0MetricTag
	FromProto(msg *otg.PatternFlowGreReserved0MetricTag) (PatternFlowGreReserved0MetricTag, error)
	// FromPbText unmarshals PatternFlowGreReserved0MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreReserved0MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreReserved0MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreReserved0MetricTag) Marshal() marshalPatternFlowGreReserved0MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreReserved0MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreReserved0MetricTag) Unmarshal() unMarshalPatternFlowGreReserved0MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreReserved0MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreReserved0MetricTag) ToProto() (*otg.PatternFlowGreReserved0MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreReserved0MetricTag) FromProto(msg *otg.PatternFlowGreReserved0MetricTag) (PatternFlowGreReserved0MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreReserved0MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreReserved0MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreReserved0MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0MetricTag) FromJson(value string) error {
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

func (obj *patternFlowGreReserved0MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved0MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved0MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreReserved0MetricTag) Clone() (PatternFlowGreReserved0MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreReserved0MetricTag()
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

// PatternFlowGreReserved0MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGreReserved0MetricTag interface {
	Validation
	// msg marshals PatternFlowGreReserved0MetricTag to protobuf object *otg.PatternFlowGreReserved0MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGreReserved0MetricTag
	// setMsg unmarshals PatternFlowGreReserved0MetricTag from protobuf object *otg.PatternFlowGreReserved0MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGreReserved0MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreReserved0MetricTag
	// validate validates PatternFlowGreReserved0MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreReserved0MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGreReserved0MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGreReserved0MetricTag
	SetName(value string) PatternFlowGreReserved0MetricTag
	// Offset returns uint32, set in PatternFlowGreReserved0MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGreReserved0MetricTag
	SetOffset(value uint32) PatternFlowGreReserved0MetricTag
	// HasOffset checks if Offset has been set in PatternFlowGreReserved0MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGreReserved0MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGreReserved0MetricTag
	SetLength(value uint32) PatternFlowGreReserved0MetricTag
	// HasLength checks if Length has been set in PatternFlowGreReserved0MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGreReserved0MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGreReserved0MetricTag object
func (obj *patternFlowGreReserved0MetricTag) SetName(value string) PatternFlowGreReserved0MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreReserved0MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreReserved0MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGreReserved0MetricTag object
func (obj *patternFlowGreReserved0MetricTag) SetOffset(value uint32) PatternFlowGreReserved0MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreReserved0MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreReserved0MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGreReserved0MetricTag object
func (obj *patternFlowGreReserved0MetricTag) SetLength(value uint32) PatternFlowGreReserved0MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGreReserved0MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGreReserved0MetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 11 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved0MetricTag.Offset <= 11 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 12 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGreReserved0MetricTag.Length <= 12 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGreReserved0MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(12)
	}

}
