package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsTrafficClassMetricTag *****
type patternFlowMplsTrafficClassMetricTag struct {
	validation
	obj          *otg.PatternFlowMplsTrafficClassMetricTag
	marshaller   marshalPatternFlowMplsTrafficClassMetricTag
	unMarshaller unMarshalPatternFlowMplsTrafficClassMetricTag
}

func NewPatternFlowMplsTrafficClassMetricTag() PatternFlowMplsTrafficClassMetricTag {
	obj := patternFlowMplsTrafficClassMetricTag{obj: &otg.PatternFlowMplsTrafficClassMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsTrafficClassMetricTag) msg() *otg.PatternFlowMplsTrafficClassMetricTag {
	return obj.obj
}

func (obj *patternFlowMplsTrafficClassMetricTag) setMsg(msg *otg.PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsTrafficClassMetricTag struct {
	obj *patternFlowMplsTrafficClassMetricTag
}

type marshalPatternFlowMplsTrafficClassMetricTag interface {
	// ToProto marshals PatternFlowMplsTrafficClassMetricTag to protobuf object *otg.PatternFlowMplsTrafficClassMetricTag
	ToProto() (*otg.PatternFlowMplsTrafficClassMetricTag, error)
	// ToPbText marshals PatternFlowMplsTrafficClassMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsTrafficClassMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsTrafficClassMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowMplsTrafficClassMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowMplsTrafficClassMetricTag struct {
	obj *patternFlowMplsTrafficClassMetricTag
}

type unMarshalPatternFlowMplsTrafficClassMetricTag interface {
	// FromProto unmarshals PatternFlowMplsTrafficClassMetricTag from protobuf object *otg.PatternFlowMplsTrafficClassMetricTag
	FromProto(msg *otg.PatternFlowMplsTrafficClassMetricTag) (PatternFlowMplsTrafficClassMetricTag, error)
	// FromPbText unmarshals PatternFlowMplsTrafficClassMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsTrafficClassMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsTrafficClassMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsTrafficClassMetricTag) Marshal() marshalPatternFlowMplsTrafficClassMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsTrafficClassMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsTrafficClassMetricTag) Unmarshal() unMarshalPatternFlowMplsTrafficClassMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsTrafficClassMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsTrafficClassMetricTag) ToProto() (*otg.PatternFlowMplsTrafficClassMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsTrafficClassMetricTag) FromProto(msg *otg.PatternFlowMplsTrafficClassMetricTag) (PatternFlowMplsTrafficClassMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsTrafficClassMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClassMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsTrafficClassMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClassMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsTrafficClassMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowMplsTrafficClassMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClassMetricTag) FromJson(value string) error {
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

func (obj *patternFlowMplsTrafficClassMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsTrafficClassMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsTrafficClassMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsTrafficClassMetricTag) Clone() (PatternFlowMplsTrafficClassMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsTrafficClassMetricTag()
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

// PatternFlowMplsTrafficClassMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowMplsTrafficClassMetricTag interface {
	Validation
	// msg marshals PatternFlowMplsTrafficClassMetricTag to protobuf object *otg.PatternFlowMplsTrafficClassMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsTrafficClassMetricTag
	// setMsg unmarshals PatternFlowMplsTrafficClassMetricTag from protobuf object *otg.PatternFlowMplsTrafficClassMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowMplsTrafficClassMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsTrafficClassMetricTag
	// validate validates PatternFlowMplsTrafficClassMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsTrafficClassMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowMplsTrafficClassMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowMplsTrafficClassMetricTag
	SetName(value string) PatternFlowMplsTrafficClassMetricTag
	// Offset returns uint32, set in PatternFlowMplsTrafficClassMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowMplsTrafficClassMetricTag
	SetOffset(value uint32) PatternFlowMplsTrafficClassMetricTag
	// HasOffset checks if Offset has been set in PatternFlowMplsTrafficClassMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowMplsTrafficClassMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowMplsTrafficClassMetricTag
	SetLength(value uint32) PatternFlowMplsTrafficClassMetricTag
	// HasLength checks if Length has been set in PatternFlowMplsTrafficClassMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowMplsTrafficClassMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowMplsTrafficClassMetricTag object
func (obj *patternFlowMplsTrafficClassMetricTag) SetName(value string) PatternFlowMplsTrafficClassMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowMplsTrafficClassMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowMplsTrafficClassMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowMplsTrafficClassMetricTag object
func (obj *patternFlowMplsTrafficClassMetricTag) SetOffset(value uint32) PatternFlowMplsTrafficClassMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowMplsTrafficClassMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowMplsTrafficClassMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowMplsTrafficClassMetricTag object
func (obj *patternFlowMplsTrafficClassMetricTag) SetLength(value uint32) PatternFlowMplsTrafficClassMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowMplsTrafficClassMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowMplsTrafficClassMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTrafficClassMetricTag.Offset <= 2 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowMplsTrafficClassMetricTag.Length <= 3 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowMplsTrafficClassMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(3)
	}

}
