package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlSynMetricTag *****
type patternFlowTcpCtlSynMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpCtlSynMetricTag
	marshaller   marshalPatternFlowTcpCtlSynMetricTag
	unMarshaller unMarshalPatternFlowTcpCtlSynMetricTag
}

func NewPatternFlowTcpCtlSynMetricTag() PatternFlowTcpCtlSynMetricTag {
	obj := patternFlowTcpCtlSynMetricTag{obj: &otg.PatternFlowTcpCtlSynMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlSynMetricTag) msg() *otg.PatternFlowTcpCtlSynMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpCtlSynMetricTag) setMsg(msg *otg.PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlSynMetricTag struct {
	obj *patternFlowTcpCtlSynMetricTag
}

type marshalPatternFlowTcpCtlSynMetricTag interface {
	// ToProto marshals PatternFlowTcpCtlSynMetricTag to protobuf object *otg.PatternFlowTcpCtlSynMetricTag
	ToProto() (*otg.PatternFlowTcpCtlSynMetricTag, error)
	// ToPbText marshals PatternFlowTcpCtlSynMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlSynMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlSynMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpCtlSynMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpCtlSynMetricTag struct {
	obj *patternFlowTcpCtlSynMetricTag
}

type unMarshalPatternFlowTcpCtlSynMetricTag interface {
	// FromProto unmarshals PatternFlowTcpCtlSynMetricTag from protobuf object *otg.PatternFlowTcpCtlSynMetricTag
	FromProto(msg *otg.PatternFlowTcpCtlSynMetricTag) (PatternFlowTcpCtlSynMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpCtlSynMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlSynMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlSynMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlSynMetricTag) Marshal() marshalPatternFlowTcpCtlSynMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlSynMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlSynMetricTag) Unmarshal() unMarshalPatternFlowTcpCtlSynMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlSynMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlSynMetricTag) ToProto() (*otg.PatternFlowTcpCtlSynMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlSynMetricTag) FromProto(msg *otg.PatternFlowTcpCtlSynMetricTag) (PatternFlowTcpCtlSynMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlSynMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSynMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlSynMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSynMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlSynMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpCtlSynMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSynMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlSynMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlSynMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlSynMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlSynMetricTag) Clone() (PatternFlowTcpCtlSynMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlSynMetricTag()
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

// PatternFlowTcpCtlSynMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpCtlSynMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpCtlSynMetricTag to protobuf object *otg.PatternFlowTcpCtlSynMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlSynMetricTag
	// setMsg unmarshals PatternFlowTcpCtlSynMetricTag from protobuf object *otg.PatternFlowTcpCtlSynMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlSynMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlSynMetricTag
	// validate validates PatternFlowTcpCtlSynMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlSynMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpCtlSynMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpCtlSynMetricTag
	SetName(value string) PatternFlowTcpCtlSynMetricTag
	// Offset returns uint32, set in PatternFlowTcpCtlSynMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpCtlSynMetricTag
	SetOffset(value uint32) PatternFlowTcpCtlSynMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpCtlSynMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpCtlSynMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpCtlSynMetricTag
	SetLength(value uint32) PatternFlowTcpCtlSynMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpCtlSynMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpCtlSynMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpCtlSynMetricTag object
func (obj *patternFlowTcpCtlSynMetricTag) SetName(value string) PatternFlowTcpCtlSynMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlSynMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlSynMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpCtlSynMetricTag object
func (obj *patternFlowTcpCtlSynMetricTag) SetOffset(value uint32) PatternFlowTcpCtlSynMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlSynMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlSynMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpCtlSynMetricTag object
func (obj *patternFlowTcpCtlSynMetricTag) SetLength(value uint32) PatternFlowTcpCtlSynMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpCtlSynMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpCtlSynMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlSynMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpCtlSynMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpCtlSynMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
