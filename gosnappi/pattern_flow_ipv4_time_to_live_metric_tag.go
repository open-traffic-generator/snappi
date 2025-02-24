package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TimeToLiveMetricTag *****
type patternFlowIpv4TimeToLiveMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4TimeToLiveMetricTag
	marshaller   marshalPatternFlowIpv4TimeToLiveMetricTag
	unMarshaller unMarshalPatternFlowIpv4TimeToLiveMetricTag
}

func NewPatternFlowIpv4TimeToLiveMetricTag() PatternFlowIpv4TimeToLiveMetricTag {
	obj := patternFlowIpv4TimeToLiveMetricTag{obj: &otg.PatternFlowIpv4TimeToLiveMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TimeToLiveMetricTag) msg() *otg.PatternFlowIpv4TimeToLiveMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4TimeToLiveMetricTag) setMsg(msg *otg.PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLiveMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TimeToLiveMetricTag struct {
	obj *patternFlowIpv4TimeToLiveMetricTag
}

type marshalPatternFlowIpv4TimeToLiveMetricTag interface {
	// ToProto marshals PatternFlowIpv4TimeToLiveMetricTag to protobuf object *otg.PatternFlowIpv4TimeToLiveMetricTag
	ToProto() (*otg.PatternFlowIpv4TimeToLiveMetricTag, error)
	// ToPbText marshals PatternFlowIpv4TimeToLiveMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TimeToLiveMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TimeToLiveMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4TimeToLiveMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4TimeToLiveMetricTag struct {
	obj *patternFlowIpv4TimeToLiveMetricTag
}

type unMarshalPatternFlowIpv4TimeToLiveMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4TimeToLiveMetricTag from protobuf object *otg.PatternFlowIpv4TimeToLiveMetricTag
	FromProto(msg *otg.PatternFlowIpv4TimeToLiveMetricTag) (PatternFlowIpv4TimeToLiveMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4TimeToLiveMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TimeToLiveMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TimeToLiveMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TimeToLiveMetricTag) Marshal() marshalPatternFlowIpv4TimeToLiveMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TimeToLiveMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TimeToLiveMetricTag) Unmarshal() unMarshalPatternFlowIpv4TimeToLiveMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TimeToLiveMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TimeToLiveMetricTag) ToProto() (*otg.PatternFlowIpv4TimeToLiveMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TimeToLiveMetricTag) FromProto(msg *otg.PatternFlowIpv4TimeToLiveMetricTag) (PatternFlowIpv4TimeToLiveMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TimeToLiveMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLiveMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TimeToLiveMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLiveMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TimeToLiveMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4TimeToLiveMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLiveMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4TimeToLiveMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TimeToLiveMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TimeToLiveMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TimeToLiveMetricTag) Clone() (PatternFlowIpv4TimeToLiveMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TimeToLiveMetricTag()
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

// PatternFlowIpv4TimeToLiveMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4TimeToLiveMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4TimeToLiveMetricTag to protobuf object *otg.PatternFlowIpv4TimeToLiveMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TimeToLiveMetricTag
	// setMsg unmarshals PatternFlowIpv4TimeToLiveMetricTag from protobuf object *otg.PatternFlowIpv4TimeToLiveMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLiveMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TimeToLiveMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TimeToLiveMetricTag
	// validate validates PatternFlowIpv4TimeToLiveMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TimeToLiveMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4TimeToLiveMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4TimeToLiveMetricTag
	SetName(value string) PatternFlowIpv4TimeToLiveMetricTag
	// Offset returns uint32, set in PatternFlowIpv4TimeToLiveMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4TimeToLiveMetricTag
	SetOffset(value uint32) PatternFlowIpv4TimeToLiveMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4TimeToLiveMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4TimeToLiveMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4TimeToLiveMetricTag
	SetLength(value uint32) PatternFlowIpv4TimeToLiveMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4TimeToLiveMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4TimeToLiveMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4TimeToLiveMetricTag object
func (obj *patternFlowIpv4TimeToLiveMetricTag) SetName(value string) PatternFlowIpv4TimeToLiveMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TimeToLiveMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4TimeToLiveMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4TimeToLiveMetricTag object
func (obj *patternFlowIpv4TimeToLiveMetricTag) SetOffset(value uint32) PatternFlowIpv4TimeToLiveMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TimeToLiveMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4TimeToLiveMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4TimeToLiveMetricTag object
func (obj *patternFlowIpv4TimeToLiveMetricTag) SetLength(value uint32) PatternFlowIpv4TimeToLiveMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4TimeToLiveMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4TimeToLiveMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TimeToLiveMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4TimeToLiveMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4TimeToLiveMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
