package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmMDNameMacAddr *****
type flowCfmMDNameMacAddr struct {
	validation
	obj          *otg.FlowCfmMDNameMacAddr
	marshaller   marshalFlowCfmMDNameMacAddr
	unMarshaller unMarshalFlowCfmMDNameMacAddr
}

func NewFlowCfmMDNameMacAddr() FlowCfmMDNameMacAddr {
	obj := flowCfmMDNameMacAddr{obj: &otg.FlowCfmMDNameMacAddr{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmMDNameMacAddr) msg() *otg.FlowCfmMDNameMacAddr {
	return obj.obj
}

func (obj *flowCfmMDNameMacAddr) setMsg(msg *otg.FlowCfmMDNameMacAddr) FlowCfmMDNameMacAddr {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmMDNameMacAddr struct {
	obj *flowCfmMDNameMacAddr
}

type marshalFlowCfmMDNameMacAddr interface {
	// ToProto marshals FlowCfmMDNameMacAddr to protobuf object *otg.FlowCfmMDNameMacAddr
	ToProto() (*otg.FlowCfmMDNameMacAddr, error)
	// ToPbText marshals FlowCfmMDNameMacAddr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmMDNameMacAddr to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmMDNameMacAddr to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmMDNameMacAddr struct {
	obj *flowCfmMDNameMacAddr
}

type unMarshalFlowCfmMDNameMacAddr interface {
	// FromProto unmarshals FlowCfmMDNameMacAddr from protobuf object *otg.FlowCfmMDNameMacAddr
	FromProto(msg *otg.FlowCfmMDNameMacAddr) (FlowCfmMDNameMacAddr, error)
	// FromPbText unmarshals FlowCfmMDNameMacAddr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmMDNameMacAddr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmMDNameMacAddr from JSON text
	FromJson(value string) error
}

func (obj *flowCfmMDNameMacAddr) Marshal() marshalFlowCfmMDNameMacAddr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmMDNameMacAddr{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmMDNameMacAddr) Unmarshal() unMarshalFlowCfmMDNameMacAddr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmMDNameMacAddr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmMDNameMacAddr) ToProto() (*otg.FlowCfmMDNameMacAddr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmMDNameMacAddr) FromProto(msg *otg.FlowCfmMDNameMacAddr) (FlowCfmMDNameMacAddr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmMDNameMacAddr) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmMDNameMacAddr) FromPbText(value string) error {
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

func (m *marshalflowCfmMDNameMacAddr) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmMDNameMacAddr) FromYaml(value string) error {
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

func (m *marshalflowCfmMDNameMacAddr) ToJson() (string, error) {
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

func (m *unMarshalflowCfmMDNameMacAddr) FromJson(value string) error {
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

func (obj *flowCfmMDNameMacAddr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmMDNameMacAddr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmMDNameMacAddr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmMDNameMacAddr) Clone() (FlowCfmMDNameMacAddr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmMDNameMacAddr()
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

// FlowCfmMDNameMacAddr is mD Name Format 3 encoding as defined in IEEE 802.1Q-2018 Table 21-19. Consists of a 6-octet MAC address followed by a 2-octet unsigned integer in network byte order. Total wire length is always 8 octets.
type FlowCfmMDNameMacAddr interface {
	Validation
	// msg marshals FlowCfmMDNameMacAddr to protobuf object *otg.FlowCfmMDNameMacAddr
	// and doesn't set defaults
	msg() *otg.FlowCfmMDNameMacAddr
	// setMsg unmarshals FlowCfmMDNameMacAddr from protobuf object *otg.FlowCfmMDNameMacAddr
	// and doesn't set defaults
	setMsg(*otg.FlowCfmMDNameMacAddr) FlowCfmMDNameMacAddr
	// provides marshal interface
	Marshal() marshalFlowCfmMDNameMacAddr
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmMDNameMacAddr
	// validate validates FlowCfmMDNameMacAddr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmMDNameMacAddr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Mac returns string, set in FlowCfmMDNameMacAddr.
	Mac() string
	// SetMac assigns string provided by user to FlowCfmMDNameMacAddr
	SetMac(value string) FlowCfmMDNameMacAddr
	// HasMac checks if Mac has been set in FlowCfmMDNameMacAddr
	HasMac() bool
	// Id returns uint32, set in FlowCfmMDNameMacAddr.
	Id() uint32
	// SetId assigns uint32 provided by user to FlowCfmMDNameMacAddr
	SetId(value uint32) FlowCfmMDNameMacAddr
	// HasId checks if Id has been set in FlowCfmMDNameMacAddr
	HasId() bool
}

// 6-octet MAC address component, e.g. "00:11:22:33:44:55".
// Mac returns a string
func (obj *flowCfmMDNameMacAddr) Mac() string {

	return *obj.obj.Mac

}

// 6-octet MAC address component, e.g. "00:11:22:33:44:55".
// Mac returns a string
func (obj *flowCfmMDNameMacAddr) HasMac() bool {
	return obj.obj.Mac != nil
}

// 6-octet MAC address component, e.g. "00:11:22:33:44:55".
// SetMac sets the string value in the FlowCfmMDNameMacAddr object
func (obj *flowCfmMDNameMacAddr) SetMac(value string) FlowCfmMDNameMacAddr {

	obj.obj.Mac = &value
	return obj
}

// Locally assigned 2-octet unsigned integer in network byte order that distinguishes multiple maintenance domains sharing the same MAC address (IEEE 802.1Q-2018 Table 21-19).
// Id returns a uint32
func (obj *flowCfmMDNameMacAddr) Id() uint32 {

	return *obj.obj.Id

}

// Locally assigned 2-octet unsigned integer in network byte order that distinguishes multiple maintenance domains sharing the same MAC address (IEEE 802.1Q-2018 Table 21-19).
// Id returns a uint32
func (obj *flowCfmMDNameMacAddr) HasId() bool {
	return obj.obj.Id != nil
}

// Locally assigned 2-octet unsigned integer in network byte order that distinguishes multiple maintenance domains sharing the same MAC address (IEEE 802.1Q-2018 Table 21-19).
// SetId sets the uint32 value in the FlowCfmMDNameMacAddr object
func (obj *flowCfmMDNameMacAddr) SetId(value uint32) FlowCfmMDNameMacAddr {

	obj.obj.Id = &value
	return obj
}

func (obj *flowCfmMDNameMacAddr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Mac != nil {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmMDNameMacAddr.Mac"))
		}

	}

	if obj.obj.Id != nil {

		if *obj.obj.Id > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowCfmMDNameMacAddr.Id <= 65535 but Got %d", *obj.obj.Id))
		}

	}

}

func (obj *flowCfmMDNameMacAddr) setDefault() {
	if obj.obj.Mac == nil {
		obj.SetMac("00:00:00:00:00:00")
	}
	if obj.obj.Id == nil {
		obj.SetId(0)
	}

}
