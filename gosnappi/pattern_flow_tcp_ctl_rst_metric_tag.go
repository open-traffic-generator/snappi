package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlRstMetricTag *****
type patternFlowTcpCtlRstMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpCtlRstMetricTag
	marshaller   marshalPatternFlowTcpCtlRstMetricTag
	unMarshaller unMarshalPatternFlowTcpCtlRstMetricTag
}

func NewPatternFlowTcpCtlRstMetricTag() PatternFlowTcpCtlRstMetricTag {
	obj := patternFlowTcpCtlRstMetricTag{obj: &otg.PatternFlowTcpCtlRstMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlRstMetricTag) msg() *otg.PatternFlowTcpCtlRstMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpCtlRstMetricTag) setMsg(msg *otg.PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlRstMetricTag struct {
	obj *patternFlowTcpCtlRstMetricTag
}

type marshalPatternFlowTcpCtlRstMetricTag interface {
	// ToProto marshals PatternFlowTcpCtlRstMetricTag to protobuf object *otg.PatternFlowTcpCtlRstMetricTag
	ToProto() (*otg.PatternFlowTcpCtlRstMetricTag, error)
	// ToPbText marshals PatternFlowTcpCtlRstMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlRstMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlRstMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlRstMetricTag struct {
	obj *patternFlowTcpCtlRstMetricTag
}

type unMarshalPatternFlowTcpCtlRstMetricTag interface {
	// FromProto unmarshals PatternFlowTcpCtlRstMetricTag from protobuf object *otg.PatternFlowTcpCtlRstMetricTag
	FromProto(msg *otg.PatternFlowTcpCtlRstMetricTag) (PatternFlowTcpCtlRstMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpCtlRstMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlRstMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlRstMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlRstMetricTag) Marshal() marshalPatternFlowTcpCtlRstMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlRstMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlRstMetricTag) Unmarshal() unMarshalPatternFlowTcpCtlRstMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlRstMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlRstMetricTag) ToProto() (*otg.PatternFlowTcpCtlRstMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlRstMetricTag) FromProto(msg *otg.PatternFlowTcpCtlRstMetricTag) (PatternFlowTcpCtlRstMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlRstMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRstMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlRstMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRstMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlRstMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRstMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlRstMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlRstMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlRstMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlRstMetricTag) Clone() (PatternFlowTcpCtlRstMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlRstMetricTag()
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

// PatternFlowTcpCtlRstMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpCtlRstMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpCtlRstMetricTag to protobuf object *otg.PatternFlowTcpCtlRstMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlRstMetricTag
	// setMsg unmarshals PatternFlowTcpCtlRstMetricTag from protobuf object *otg.PatternFlowTcpCtlRstMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlRstMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlRstMetricTag
	// validate validates PatternFlowTcpCtlRstMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlRstMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpCtlRstMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpCtlRstMetricTag
	SetName(value string) PatternFlowTcpCtlRstMetricTag
	// Offset returns uint32, set in PatternFlowTcpCtlRstMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpCtlRstMetricTag
	SetOffset(value uint32) PatternFlowTcpCtlRstMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpCtlRstMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpCtlRstMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpCtlRstMetricTag
	SetLength(value uint32) PatternFlowTcpCtlRstMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpCtlRstMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpCtlRstMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpCtlRstMetricTag object
func (obj *patternFlowTcpCtlRstMetricTag) SetName(value string) PatternFlowTcpCtlRstMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlRstMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlRstMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpCtlRstMetricTag object
func (obj *patternFlowTcpCtlRstMetricTag) SetOffset(value uint32) PatternFlowTcpCtlRstMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlRstMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlRstMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpCtlRstMetricTag object
func (obj *patternFlowTcpCtlRstMetricTag) SetLength(value uint32) PatternFlowTcpCtlRstMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpCtlRstMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpCtlRstMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlRstMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpCtlRstMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpCtlRstMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
