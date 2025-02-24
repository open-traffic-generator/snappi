package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseControlOpCodeMetricTag *****
type patternFlowEthernetPauseControlOpCodeMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetPauseControlOpCodeMetricTag
	marshaller   marshalPatternFlowEthernetPauseControlOpCodeMetricTag
	unMarshaller unMarshalPatternFlowEthernetPauseControlOpCodeMetricTag
}

func NewPatternFlowEthernetPauseControlOpCodeMetricTag() PatternFlowEthernetPauseControlOpCodeMetricTag {
	obj := patternFlowEthernetPauseControlOpCodeMetricTag{obj: &otg.PatternFlowEthernetPauseControlOpCodeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) msg() *otg.PatternFlowEthernetPauseControlOpCodeMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) setMsg(msg *otg.PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseControlOpCodeMetricTag struct {
	obj *patternFlowEthernetPauseControlOpCodeMetricTag
}

type marshalPatternFlowEthernetPauseControlOpCodeMetricTag interface {
	// ToProto marshals PatternFlowEthernetPauseControlOpCodeMetricTag to protobuf object *otg.PatternFlowEthernetPauseControlOpCodeMetricTag
	ToProto() (*otg.PatternFlowEthernetPauseControlOpCodeMetricTag, error)
	// ToPbText marshals PatternFlowEthernetPauseControlOpCodeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseControlOpCodeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseControlOpCodeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetPauseControlOpCodeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetPauseControlOpCodeMetricTag struct {
	obj *patternFlowEthernetPauseControlOpCodeMetricTag
}

type unMarshalPatternFlowEthernetPauseControlOpCodeMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetPauseControlOpCodeMetricTag from protobuf object *otg.PatternFlowEthernetPauseControlOpCodeMetricTag
	FromProto(msg *otg.PatternFlowEthernetPauseControlOpCodeMetricTag) (PatternFlowEthernetPauseControlOpCodeMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetPauseControlOpCodeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseControlOpCodeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseControlOpCodeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) Marshal() marshalPatternFlowEthernetPauseControlOpCodeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseControlOpCodeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) Unmarshal() unMarshalPatternFlowEthernetPauseControlOpCodeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseControlOpCodeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseControlOpCodeMetricTag) ToProto() (*otg.PatternFlowEthernetPauseControlOpCodeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseControlOpCodeMetricTag) FromProto(msg *otg.PatternFlowEthernetPauseControlOpCodeMetricTag) (PatternFlowEthernetPauseControlOpCodeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseControlOpCodeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCodeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseControlOpCodeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCodeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseControlOpCodeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowEthernetPauseControlOpCodeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCodeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) Clone() (PatternFlowEthernetPauseControlOpCodeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseControlOpCodeMetricTag()
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

// PatternFlowEthernetPauseControlOpCodeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetPauseControlOpCodeMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetPauseControlOpCodeMetricTag to protobuf object *otg.PatternFlowEthernetPauseControlOpCodeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseControlOpCodeMetricTag
	// setMsg unmarshals PatternFlowEthernetPauseControlOpCodeMetricTag from protobuf object *otg.PatternFlowEthernetPauseControlOpCodeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseControlOpCodeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseControlOpCodeMetricTag
	// validate validates PatternFlowEthernetPauseControlOpCodeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseControlOpCodeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetPauseControlOpCodeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetPauseControlOpCodeMetricTag
	SetName(value string) PatternFlowEthernetPauseControlOpCodeMetricTag
	// Offset returns uint32, set in PatternFlowEthernetPauseControlOpCodeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetPauseControlOpCodeMetricTag
	SetOffset(value uint32) PatternFlowEthernetPauseControlOpCodeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetPauseControlOpCodeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetPauseControlOpCodeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetPauseControlOpCodeMetricTag
	SetLength(value uint32) PatternFlowEthernetPauseControlOpCodeMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetPauseControlOpCodeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetPauseControlOpCodeMetricTag object
func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) SetName(value string) PatternFlowEthernetPauseControlOpCodeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetPauseControlOpCodeMetricTag object
func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) SetOffset(value uint32) PatternFlowEthernetPauseControlOpCodeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetPauseControlOpCodeMetricTag object
func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) SetLength(value uint32) PatternFlowEthernetPauseControlOpCodeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetPauseControlOpCodeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCodeMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetPauseControlOpCodeMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetPauseControlOpCodeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
