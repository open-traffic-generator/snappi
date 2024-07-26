package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpSeqNumMetricTag *****
type patternFlowTcpSeqNumMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpSeqNumMetricTag
	marshaller   marshalPatternFlowTcpSeqNumMetricTag
	unMarshaller unMarshalPatternFlowTcpSeqNumMetricTag
}

func NewPatternFlowTcpSeqNumMetricTag() PatternFlowTcpSeqNumMetricTag {
	obj := patternFlowTcpSeqNumMetricTag{obj: &otg.PatternFlowTcpSeqNumMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpSeqNumMetricTag) msg() *otg.PatternFlowTcpSeqNumMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpSeqNumMetricTag) setMsg(msg *otg.PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpSeqNumMetricTag struct {
	obj *patternFlowTcpSeqNumMetricTag
}

type marshalPatternFlowTcpSeqNumMetricTag interface {
	// ToProto marshals PatternFlowTcpSeqNumMetricTag to protobuf object *otg.PatternFlowTcpSeqNumMetricTag
	ToProto() (*otg.PatternFlowTcpSeqNumMetricTag, error)
	// ToPbText marshals PatternFlowTcpSeqNumMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpSeqNumMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpSeqNumMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpSeqNumMetricTag struct {
	obj *patternFlowTcpSeqNumMetricTag
}

type unMarshalPatternFlowTcpSeqNumMetricTag interface {
	// FromProto unmarshals PatternFlowTcpSeqNumMetricTag from protobuf object *otg.PatternFlowTcpSeqNumMetricTag
	FromProto(msg *otg.PatternFlowTcpSeqNumMetricTag) (PatternFlowTcpSeqNumMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpSeqNumMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpSeqNumMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpSeqNumMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpSeqNumMetricTag) Marshal() marshalPatternFlowTcpSeqNumMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpSeqNumMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpSeqNumMetricTag) Unmarshal() unMarshalPatternFlowTcpSeqNumMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpSeqNumMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpSeqNumMetricTag) ToProto() (*otg.PatternFlowTcpSeqNumMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpSeqNumMetricTag) FromProto(msg *otg.PatternFlowTcpSeqNumMetricTag) (PatternFlowTcpSeqNumMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpSeqNumMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNumMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpSeqNumMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNumMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpSeqNumMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNumMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpSeqNumMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpSeqNumMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpSeqNumMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpSeqNumMetricTag) Clone() (PatternFlowTcpSeqNumMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpSeqNumMetricTag()
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

// PatternFlowTcpSeqNumMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpSeqNumMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpSeqNumMetricTag to protobuf object *otg.PatternFlowTcpSeqNumMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpSeqNumMetricTag
	// setMsg unmarshals PatternFlowTcpSeqNumMetricTag from protobuf object *otg.PatternFlowTcpSeqNumMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpSeqNumMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpSeqNumMetricTag
	// validate validates PatternFlowTcpSeqNumMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpSeqNumMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpSeqNumMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpSeqNumMetricTag
	SetName(value string) PatternFlowTcpSeqNumMetricTag
	// Offset returns uint32, set in PatternFlowTcpSeqNumMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpSeqNumMetricTag
	SetOffset(value uint32) PatternFlowTcpSeqNumMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpSeqNumMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpSeqNumMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpSeqNumMetricTag
	SetLength(value uint32) PatternFlowTcpSeqNumMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpSeqNumMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpSeqNumMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpSeqNumMetricTag object
func (obj *patternFlowTcpSeqNumMetricTag) SetName(value string) PatternFlowTcpSeqNumMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpSeqNumMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpSeqNumMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpSeqNumMetricTag object
func (obj *patternFlowTcpSeqNumMetricTag) SetOffset(value uint32) PatternFlowTcpSeqNumMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpSeqNumMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpSeqNumMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpSeqNumMetricTag object
func (obj *patternFlowTcpSeqNumMetricTag) SetLength(value uint32) PatternFlowTcpSeqNumMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpSeqNumMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpSeqNumMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpSeqNumMetricTag.Offset <= 31 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpSeqNumMetricTag.Length <= 32 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpSeqNumMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(32)
	}

}
