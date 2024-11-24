package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathRecordRouteType1Label *****
type flowRSVPPathRecordRouteType1Label struct {
	validation
	obj          *otg.FlowRSVPPathRecordRouteType1Label
	marshaller   marshalFlowRSVPPathRecordRouteType1Label
	unMarshaller unMarshalFlowRSVPPathRecordRouteType1Label
	lengthHolder FlowRSVPRouteRecordLength
	flagsHolder  PatternFlowRSVPPathRecordRouteType1LabelFlags
	cTypeHolder  PatternFlowRSVPPathRecordRouteType1LabelCType
	labelHolder  FlowRSVPPathRecordRouteLabel
}

func NewFlowRSVPPathRecordRouteType1Label() FlowRSVPPathRecordRouteType1Label {
	obj := flowRSVPPathRecordRouteType1Label{obj: &otg.FlowRSVPPathRecordRouteType1Label{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathRecordRouteType1Label) msg() *otg.FlowRSVPPathRecordRouteType1Label {
	return obj.obj
}

func (obj *flowRSVPPathRecordRouteType1Label) setMsg(msg *otg.FlowRSVPPathRecordRouteType1Label) FlowRSVPPathRecordRouteType1Label {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathRecordRouteType1Label struct {
	obj *flowRSVPPathRecordRouteType1Label
}

type marshalFlowRSVPPathRecordRouteType1Label interface {
	// ToProto marshals FlowRSVPPathRecordRouteType1Label to protobuf object *otg.FlowRSVPPathRecordRouteType1Label
	ToProto() (*otg.FlowRSVPPathRecordRouteType1Label, error)
	// ToPbText marshals FlowRSVPPathRecordRouteType1Label to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathRecordRouteType1Label to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathRecordRouteType1Label to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathRecordRouteType1Label to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathRecordRouteType1Label struct {
	obj *flowRSVPPathRecordRouteType1Label
}

type unMarshalFlowRSVPPathRecordRouteType1Label interface {
	// FromProto unmarshals FlowRSVPPathRecordRouteType1Label from protobuf object *otg.FlowRSVPPathRecordRouteType1Label
	FromProto(msg *otg.FlowRSVPPathRecordRouteType1Label) (FlowRSVPPathRecordRouteType1Label, error)
	// FromPbText unmarshals FlowRSVPPathRecordRouteType1Label from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathRecordRouteType1Label from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathRecordRouteType1Label from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathRecordRouteType1Label) Marshal() marshalFlowRSVPPathRecordRouteType1Label {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathRecordRouteType1Label{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathRecordRouteType1Label) Unmarshal() unMarshalFlowRSVPPathRecordRouteType1Label {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathRecordRouteType1Label{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathRecordRouteType1Label) ToProto() (*otg.FlowRSVPPathRecordRouteType1Label, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathRecordRouteType1Label) FromProto(msg *otg.FlowRSVPPathRecordRouteType1Label) (FlowRSVPPathRecordRouteType1Label, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathRecordRouteType1Label) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1Label) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalflowRSVPPathRecordRouteType1Label) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1Label) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalflowRSVPPathRecordRouteType1Label) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathRecordRouteType1Label) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1Label) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowRSVPPathRecordRouteType1Label) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathRecordRouteType1Label) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathRecordRouteType1Label) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathRecordRouteType1Label) Clone() (FlowRSVPPathRecordRouteType1Label, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathRecordRouteType1Label()
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

func (obj *flowRSVPPathRecordRouteType1Label) setNil() {
	obj.lengthHolder = nil
	obj.flagsHolder = nil
	obj.cTypeHolder = nil
	obj.labelHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathRecordRouteType1Label is class = RECORD_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: Label, C-Type: 3
type FlowRSVPPathRecordRouteType1Label interface {
	Validation
	// msg marshals FlowRSVPPathRecordRouteType1Label to protobuf object *otg.FlowRSVPPathRecordRouteType1Label
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathRecordRouteType1Label
	// setMsg unmarshals FlowRSVPPathRecordRouteType1Label from protobuf object *otg.FlowRSVPPathRecordRouteType1Label
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathRecordRouteType1Label) FlowRSVPPathRecordRouteType1Label
	// provides marshal interface
	Marshal() marshalFlowRSVPPathRecordRouteType1Label
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathRecordRouteType1Label
	// validate validates FlowRSVPPathRecordRouteType1Label
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathRecordRouteType1Label, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPRouteRecordLength, set in FlowRSVPPathRecordRouteType1Label.
	// FlowRSVPRouteRecordLength is description is TBD
	Length() FlowRSVPRouteRecordLength
	// SetLength assigns FlowRSVPRouteRecordLength provided by user to FlowRSVPPathRecordRouteType1Label.
	// FlowRSVPRouteRecordLength is description is TBD
	SetLength(value FlowRSVPRouteRecordLength) FlowRSVPPathRecordRouteType1Label
	// HasLength checks if Length has been set in FlowRSVPPathRecordRouteType1Label
	HasLength() bool
	// Flags returns PatternFlowRSVPPathRecordRouteType1LabelFlags, set in FlowRSVPPathRecordRouteType1Label.
	// PatternFlowRSVPPathRecordRouteType1LabelFlags is 0x01 = Global label. This flag indicates that the label will be understood if received on any interface.
	Flags() PatternFlowRSVPPathRecordRouteType1LabelFlags
	// SetFlags assigns PatternFlowRSVPPathRecordRouteType1LabelFlags provided by user to FlowRSVPPathRecordRouteType1Label.
	// PatternFlowRSVPPathRecordRouteType1LabelFlags is 0x01 = Global label. This flag indicates that the label will be understood if received on any interface.
	SetFlags(value PatternFlowRSVPPathRecordRouteType1LabelFlags) FlowRSVPPathRecordRouteType1Label
	// HasFlags checks if Flags has been set in FlowRSVPPathRecordRouteType1Label
	HasFlags() bool
	// CType returns PatternFlowRSVPPathRecordRouteType1LabelCType, set in FlowRSVPPathRecordRouteType1Label.
	// PatternFlowRSVPPathRecordRouteType1LabelCType is the C-Type of the included Label Object. Copied from the Label object.
	CType() PatternFlowRSVPPathRecordRouteType1LabelCType
	// SetCType assigns PatternFlowRSVPPathRecordRouteType1LabelCType provided by user to FlowRSVPPathRecordRouteType1Label.
	// PatternFlowRSVPPathRecordRouteType1LabelCType is the C-Type of the included Label Object. Copied from the Label object.
	SetCType(value PatternFlowRSVPPathRecordRouteType1LabelCType) FlowRSVPPathRecordRouteType1Label
	// HasCType checks if CType has been set in FlowRSVPPathRecordRouteType1Label
	HasCType() bool
	// Label returns FlowRSVPPathRecordRouteLabel, set in FlowRSVPPathRecordRouteType1Label.
	// FlowRSVPPathRecordRouteLabel is description is TBD
	Label() FlowRSVPPathRecordRouteLabel
	// SetLabel assigns FlowRSVPPathRecordRouteLabel provided by user to FlowRSVPPathRecordRouteType1Label.
	// FlowRSVPPathRecordRouteLabel is description is TBD
	SetLabel(value FlowRSVPPathRecordRouteLabel) FlowRSVPPathRecordRouteType1Label
	// HasLabel checks if Label has been set in FlowRSVPPathRecordRouteType1Label
	HasLabel() bool
	setNil()
}

// The Length contains the total length of the subobject in bytes, including the Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// Length returns a FlowRSVPRouteRecordLength
func (obj *flowRSVPPathRecordRouteType1Label) Length() FlowRSVPRouteRecordLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowRSVPRouteRecordLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowRSVPRouteRecordLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// The Length contains the total length of the subobject in bytes, including the Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// Length returns a FlowRSVPRouteRecordLength
func (obj *flowRSVPPathRecordRouteType1Label) HasLength() bool {
	return obj.obj.Length != nil
}

// The Length contains the total length of the subobject in bytes, including the Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// SetLength sets the FlowRSVPRouteRecordLength value in the FlowRSVPPathRecordRouteType1Label object
func (obj *flowRSVPPathRecordRouteType1Label) SetLength(value FlowRSVPRouteRecordLength) FlowRSVPPathRecordRouteType1Label {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Flags returns a PatternFlowRSVPPathRecordRouteType1LabelFlags
func (obj *flowRSVPPathRecordRouteType1Label) Flags() PatternFlowRSVPPathRecordRouteType1LabelFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewPatternFlowRSVPPathRecordRouteType1LabelFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &patternFlowRSVPPathRecordRouteType1LabelFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a PatternFlowRSVPPathRecordRouteType1LabelFlags
func (obj *flowRSVPPathRecordRouteType1Label) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the PatternFlowRSVPPathRecordRouteType1LabelFlags value in the FlowRSVPPathRecordRouteType1Label object
func (obj *flowRSVPPathRecordRouteType1Label) SetFlags(value PatternFlowRSVPPathRecordRouteType1LabelFlags) FlowRSVPPathRecordRouteType1Label {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// CType returns a PatternFlowRSVPPathRecordRouteType1LabelCType
func (obj *flowRSVPPathRecordRouteType1Label) CType() PatternFlowRSVPPathRecordRouteType1LabelCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewPatternFlowRSVPPathRecordRouteType1LabelCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &patternFlowRSVPPathRecordRouteType1LabelCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a PatternFlowRSVPPathRecordRouteType1LabelCType
func (obj *flowRSVPPathRecordRouteType1Label) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the PatternFlowRSVPPathRecordRouteType1LabelCType value in the FlowRSVPPathRecordRouteType1Label object
func (obj *flowRSVPPathRecordRouteType1Label) SetCType(value PatternFlowRSVPPathRecordRouteType1LabelCType) FlowRSVPPathRecordRouteType1Label {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

// The contents of the Label Object. Copied from the Label Object.
// Label returns a FlowRSVPPathRecordRouteLabel
func (obj *flowRSVPPathRecordRouteType1Label) Label() FlowRSVPPathRecordRouteLabel {
	if obj.obj.Label == nil {
		obj.obj.Label = NewFlowRSVPPathRecordRouteLabel().msg()
	}
	if obj.labelHolder == nil {
		obj.labelHolder = &flowRSVPPathRecordRouteLabel{obj: obj.obj.Label}
	}
	return obj.labelHolder
}

// The contents of the Label Object. Copied from the Label Object.
// Label returns a FlowRSVPPathRecordRouteLabel
func (obj *flowRSVPPathRecordRouteType1Label) HasLabel() bool {
	return obj.obj.Label != nil
}

// The contents of the Label Object. Copied from the Label Object.
// SetLabel sets the FlowRSVPPathRecordRouteLabel value in the FlowRSVPPathRecordRouteType1Label object
func (obj *flowRSVPPathRecordRouteType1Label) SetLabel(value FlowRSVPPathRecordRouteLabel) FlowRSVPPathRecordRouteType1Label {

	obj.labelHolder = nil
	obj.obj.Label = value.msg()

	return obj
}

func (obj *flowRSVPPathRecordRouteType1Label) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.CType != nil {

		obj.CType().validateObj(vObj, set_default)
	}

	if obj.obj.Label != nil {

		obj.Label().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathRecordRouteType1Label) setDefault() {

}
