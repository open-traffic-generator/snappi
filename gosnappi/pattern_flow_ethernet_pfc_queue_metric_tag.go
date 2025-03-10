package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPfcQueueMetricTag *****
type patternFlowEthernetPfcQueueMetricTag struct {
	validation
	obj          *otg.PatternFlowEthernetPfcQueueMetricTag
	marshaller   marshalPatternFlowEthernetPfcQueueMetricTag
	unMarshaller unMarshalPatternFlowEthernetPfcQueueMetricTag
}

func NewPatternFlowEthernetPfcQueueMetricTag() PatternFlowEthernetPfcQueueMetricTag {
	obj := patternFlowEthernetPfcQueueMetricTag{obj: &otg.PatternFlowEthernetPfcQueueMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPfcQueueMetricTag) msg() *otg.PatternFlowEthernetPfcQueueMetricTag {
	return obj.obj
}

func (obj *patternFlowEthernetPfcQueueMetricTag) setMsg(msg *otg.PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueueMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPfcQueueMetricTag struct {
	obj *patternFlowEthernetPfcQueueMetricTag
}

type marshalPatternFlowEthernetPfcQueueMetricTag interface {
	// ToProto marshals PatternFlowEthernetPfcQueueMetricTag to protobuf object *otg.PatternFlowEthernetPfcQueueMetricTag
	ToProto() (*otg.PatternFlowEthernetPfcQueueMetricTag, error)
	// ToPbText marshals PatternFlowEthernetPfcQueueMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPfcQueueMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPfcQueueMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetPfcQueueMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetPfcQueueMetricTag struct {
	obj *patternFlowEthernetPfcQueueMetricTag
}

type unMarshalPatternFlowEthernetPfcQueueMetricTag interface {
	// FromProto unmarshals PatternFlowEthernetPfcQueueMetricTag from protobuf object *otg.PatternFlowEthernetPfcQueueMetricTag
	FromProto(msg *otg.PatternFlowEthernetPfcQueueMetricTag) (PatternFlowEthernetPfcQueueMetricTag, error)
	// FromPbText unmarshals PatternFlowEthernetPfcQueueMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPfcQueueMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPfcQueueMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPfcQueueMetricTag) Marshal() marshalPatternFlowEthernetPfcQueueMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPfcQueueMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPfcQueueMetricTag) Unmarshal() unMarshalPatternFlowEthernetPfcQueueMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPfcQueueMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPfcQueueMetricTag) ToProto() (*otg.PatternFlowEthernetPfcQueueMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPfcQueueMetricTag) FromProto(msg *otg.PatternFlowEthernetPfcQueueMetricTag) (PatternFlowEthernetPfcQueueMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPfcQueueMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueueMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPfcQueueMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueueMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPfcQueueMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowEthernetPfcQueueMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueueMetricTag) FromJson(value string) error {
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

func (obj *patternFlowEthernetPfcQueueMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPfcQueueMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPfcQueueMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPfcQueueMetricTag) Clone() (PatternFlowEthernetPfcQueueMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPfcQueueMetricTag()
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

// PatternFlowEthernetPfcQueueMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowEthernetPfcQueueMetricTag interface {
	Validation
	// msg marshals PatternFlowEthernetPfcQueueMetricTag to protobuf object *otg.PatternFlowEthernetPfcQueueMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPfcQueueMetricTag
	// setMsg unmarshals PatternFlowEthernetPfcQueueMetricTag from protobuf object *otg.PatternFlowEthernetPfcQueueMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueueMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPfcQueueMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPfcQueueMetricTag
	// validate validates PatternFlowEthernetPfcQueueMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPfcQueueMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowEthernetPfcQueueMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowEthernetPfcQueueMetricTag
	SetName(value string) PatternFlowEthernetPfcQueueMetricTag
	// Offset returns uint32, set in PatternFlowEthernetPfcQueueMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowEthernetPfcQueueMetricTag
	SetOffset(value uint32) PatternFlowEthernetPfcQueueMetricTag
	// HasOffset checks if Offset has been set in PatternFlowEthernetPfcQueueMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowEthernetPfcQueueMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowEthernetPfcQueueMetricTag
	SetLength(value uint32) PatternFlowEthernetPfcQueueMetricTag
	// HasLength checks if Length has been set in PatternFlowEthernetPfcQueueMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowEthernetPfcQueueMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowEthernetPfcQueueMetricTag object
func (obj *patternFlowEthernetPfcQueueMetricTag) SetName(value string) PatternFlowEthernetPfcQueueMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPfcQueueMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowEthernetPfcQueueMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowEthernetPfcQueueMetricTag object
func (obj *patternFlowEthernetPfcQueueMetricTag) SetOffset(value uint32) PatternFlowEthernetPfcQueueMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPfcQueueMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowEthernetPfcQueueMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowEthernetPfcQueueMetricTag object
func (obj *patternFlowEthernetPfcQueueMetricTag) SetLength(value uint32) PatternFlowEthernetPfcQueueMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowEthernetPfcQueueMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowEthernetPfcQueueMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPfcQueueMetricTag.Offset <= 2 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowEthernetPfcQueueMetricTag.Length <= 3 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowEthernetPfcQueueMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(3)
	}

}
