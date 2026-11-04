package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmMaidCfm *****
type flowCfmMaidCfm struct {
	validation
	obj               *otg.FlowCfmMaidCfm
	marshaller        marshalFlowCfmMaidCfm
	unMarshaller      unMarshalFlowCfmMaidCfm
	mdNameHolder      FlowCfmMDNameFormat
	shortMaNameHolder FlowCfmShortMANameFormat
}

func NewFlowCfmMaidCfm() FlowCfmMaidCfm {
	obj := flowCfmMaidCfm{obj: &otg.FlowCfmMaidCfm{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmMaidCfm) msg() *otg.FlowCfmMaidCfm {
	return obj.obj
}

func (obj *flowCfmMaidCfm) setMsg(msg *otg.FlowCfmMaidCfm) FlowCfmMaidCfm {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmMaidCfm struct {
	obj *flowCfmMaidCfm
}

type marshalFlowCfmMaidCfm interface {
	// ToProto marshals FlowCfmMaidCfm to protobuf object *otg.FlowCfmMaidCfm
	ToProto() (*otg.FlowCfmMaidCfm, error)
	// ToPbText marshals FlowCfmMaidCfm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmMaidCfm to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmMaidCfm to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmMaidCfm struct {
	obj *flowCfmMaidCfm
}

type unMarshalFlowCfmMaidCfm interface {
	// FromProto unmarshals FlowCfmMaidCfm from protobuf object *otg.FlowCfmMaidCfm
	FromProto(msg *otg.FlowCfmMaidCfm) (FlowCfmMaidCfm, error)
	// FromPbText unmarshals FlowCfmMaidCfm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmMaidCfm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmMaidCfm from JSON text
	FromJson(value string) error
}

func (obj *flowCfmMaidCfm) Marshal() marshalFlowCfmMaidCfm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmMaidCfm{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmMaidCfm) Unmarshal() unMarshalFlowCfmMaidCfm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmMaidCfm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmMaidCfm) ToProto() (*otg.FlowCfmMaidCfm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmMaidCfm) FromProto(msg *otg.FlowCfmMaidCfm) (FlowCfmMaidCfm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmMaidCfm) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmMaidCfm) FromPbText(value string) error {
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

func (m *marshalflowCfmMaidCfm) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmMaidCfm) FromYaml(value string) error {
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

func (m *marshalflowCfmMaidCfm) ToJson() (string, error) {
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

func (m *unMarshalflowCfmMaidCfm) FromJson(value string) error {
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

func (obj *flowCfmMaidCfm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmMaidCfm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmMaidCfm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmMaidCfm) Clone() (FlowCfmMaidCfm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmMaidCfm()
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

func (obj *flowCfmMaidCfm) setNil() {
	obj.mdNameHolder = nil
	obj.shortMaNameHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmMaidCfm is iEEE 802.1Q-2018 MAID encoding: MD Name (optional) + Short MA Name, zero-padded to fill 48 bytes in the CCM PDU. The choice selects the Short MA Name Format as defined in IEEE 802.1Q-2018 Table 21-20 and provides a typed input for each format.
type FlowCfmMaidCfm interface {
	Validation
	// msg marshals FlowCfmMaidCfm to protobuf object *otg.FlowCfmMaidCfm
	// and doesn't set defaults
	msg() *otg.FlowCfmMaidCfm
	// setMsg unmarshals FlowCfmMaidCfm from protobuf object *otg.FlowCfmMaidCfm
	// and doesn't set defaults
	setMsg(*otg.FlowCfmMaidCfm) FlowCfmMaidCfm
	// provides marshal interface
	Marshal() marshalFlowCfmMaidCfm
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmMaidCfm
	// validate validates FlowCfmMaidCfm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmMaidCfm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MdName returns FlowCfmMDNameFormat, set in FlowCfmMaidCfm.
	// FlowCfmMDNameFormat is maintenance domain name encoding. The choice selects the MD Name Format as defined in IEEE 802.1Q-2018 Table 21-19 and provides a typed input for each format. Length is auto-computed from the value unless overridden.
	MdName() FlowCfmMDNameFormat
	// SetMdName assigns FlowCfmMDNameFormat provided by user to FlowCfmMaidCfm.
	// FlowCfmMDNameFormat is maintenance domain name encoding. The choice selects the MD Name Format as defined in IEEE 802.1Q-2018 Table 21-19 and provides a typed input for each format. Length is auto-computed from the value unless overridden.
	SetMdName(value FlowCfmMDNameFormat) FlowCfmMaidCfm
	// HasMdName checks if MdName has been set in FlowCfmMaidCfm
	HasMdName() bool
	// ShortMaName returns FlowCfmShortMANameFormat, set in FlowCfmMaidCfm.
	// FlowCfmShortMANameFormat is short MA Name encoding (IEEE 802.1Q-2018 Table 21-20). Choose one of four formats and provide its typed value. Length is auto-computed from the value unless overridden.
	ShortMaName() FlowCfmShortMANameFormat
	// SetShortMaName assigns FlowCfmShortMANameFormat provided by user to FlowCfmMaidCfm.
	// FlowCfmShortMANameFormat is short MA Name encoding (IEEE 802.1Q-2018 Table 21-20). Choose one of four formats and provide its typed value. Length is auto-computed from the value unless overridden.
	SetShortMaName(value FlowCfmShortMANameFormat) FlowCfmMaidCfm
	// HasShortMaName checks if ShortMaName has been set in FlowCfmMaidCfm
	HasShortMaName() bool
	setNil()
}

// Maintenance domain name encoding. When present, the MD Name field is included in the MAID using the selected format (IEEE 802.1Q-2018 Table 21-19 Formats 2-4). When absent, Format 1 (No MD Name) is encoded and the MAID begins directly with the Short MA Name.
// MdName returns a FlowCfmMDNameFormat
func (obj *flowCfmMaidCfm) MdName() FlowCfmMDNameFormat {
	if obj.obj.MdName == nil {
		obj.obj.MdName = NewFlowCfmMDNameFormat().msg()
	}
	if obj.mdNameHolder == nil {
		obj.mdNameHolder = &flowCfmMDNameFormat{obj: obj.obj.MdName}
	}
	return obj.mdNameHolder
}

// Maintenance domain name encoding. When present, the MD Name field is included in the MAID using the selected format (IEEE 802.1Q-2018 Table 21-19 Formats 2-4). When absent, Format 1 (No MD Name) is encoded and the MAID begins directly with the Short MA Name.
// MdName returns a FlowCfmMDNameFormat
func (obj *flowCfmMaidCfm) HasMdName() bool {
	return obj.obj.MdName != nil
}

// Maintenance domain name encoding. When present, the MD Name field is included in the MAID using the selected format (IEEE 802.1Q-2018 Table 21-19 Formats 2-4). When absent, Format 1 (No MD Name) is encoded and the MAID begins directly with the Short MA Name.
// SetMdName sets the FlowCfmMDNameFormat value in the FlowCfmMaidCfm object
func (obj *flowCfmMaidCfm) SetMdName(value FlowCfmMDNameFormat) FlowCfmMaidCfm {

	obj.mdNameHolder = nil
	obj.obj.MdName = value.msg()

	return obj
}

// Short MA Name encoding. The Short MA Name field is mandatory in the MAID and identifies the Maintenance Association within its Maintenance Domain. The selected format determines the on-wire encoding per IEEE 802.1Q-2018 Table 21-20 (Formats 1-4).
// ShortMaName returns a FlowCfmShortMANameFormat
func (obj *flowCfmMaidCfm) ShortMaName() FlowCfmShortMANameFormat {
	if obj.obj.ShortMaName == nil {
		obj.obj.ShortMaName = NewFlowCfmShortMANameFormat().msg()
	}
	if obj.shortMaNameHolder == nil {
		obj.shortMaNameHolder = &flowCfmShortMANameFormat{obj: obj.obj.ShortMaName}
	}
	return obj.shortMaNameHolder
}

// Short MA Name encoding. The Short MA Name field is mandatory in the MAID and identifies the Maintenance Association within its Maintenance Domain. The selected format determines the on-wire encoding per IEEE 802.1Q-2018 Table 21-20 (Formats 1-4).
// ShortMaName returns a FlowCfmShortMANameFormat
func (obj *flowCfmMaidCfm) HasShortMaName() bool {
	return obj.obj.ShortMaName != nil
}

// Short MA Name encoding. The Short MA Name field is mandatory in the MAID and identifies the Maintenance Association within its Maintenance Domain. The selected format determines the on-wire encoding per IEEE 802.1Q-2018 Table 21-20 (Formats 1-4).
// SetShortMaName sets the FlowCfmShortMANameFormat value in the FlowCfmMaidCfm object
func (obj *flowCfmMaidCfm) SetShortMaName(value FlowCfmShortMANameFormat) FlowCfmMaidCfm {

	obj.shortMaNameHolder = nil
	obj.obj.ShortMaName = value.msg()

	return obj
}

func (obj *flowCfmMaidCfm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MdName != nil {

		obj.MdName().validateObj(vObj, set_default)
	}

	if obj.obj.ShortMaName != nil {

		obj.ShortMaName().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmMaidCfm) setDefault() {

}
