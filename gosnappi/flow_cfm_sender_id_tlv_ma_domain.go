package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmSenderIdTlvMADomain *****
type flowCfmSenderIdTlvMADomain struct {
	validation
	obj                     *otg.FlowCfmSenderIdTlvMADomain
	marshaller              marshalFlowCfmSenderIdTlvMADomain
	unMarshaller            unMarshalFlowCfmSenderIdTlvMADomain
	lengthHolder            FlowCfmLength
	managementAddressHolder FlowCfmSenderIdTlvMgmtAddress
}

func NewFlowCfmSenderIdTlvMADomain() FlowCfmSenderIdTlvMADomain {
	obj := flowCfmSenderIdTlvMADomain{obj: &otg.FlowCfmSenderIdTlvMADomain{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmSenderIdTlvMADomain) msg() *otg.FlowCfmSenderIdTlvMADomain {
	return obj.obj
}

func (obj *flowCfmSenderIdTlvMADomain) setMsg(msg *otg.FlowCfmSenderIdTlvMADomain) FlowCfmSenderIdTlvMADomain {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmSenderIdTlvMADomain struct {
	obj *flowCfmSenderIdTlvMADomain
}

type marshalFlowCfmSenderIdTlvMADomain interface {
	// ToProto marshals FlowCfmSenderIdTlvMADomain to protobuf object *otg.FlowCfmSenderIdTlvMADomain
	ToProto() (*otg.FlowCfmSenderIdTlvMADomain, error)
	// ToPbText marshals FlowCfmSenderIdTlvMADomain to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmSenderIdTlvMADomain to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmSenderIdTlvMADomain to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmSenderIdTlvMADomain struct {
	obj *flowCfmSenderIdTlvMADomain
}

type unMarshalFlowCfmSenderIdTlvMADomain interface {
	// FromProto unmarshals FlowCfmSenderIdTlvMADomain from protobuf object *otg.FlowCfmSenderIdTlvMADomain
	FromProto(msg *otg.FlowCfmSenderIdTlvMADomain) (FlowCfmSenderIdTlvMADomain, error)
	// FromPbText unmarshals FlowCfmSenderIdTlvMADomain from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmSenderIdTlvMADomain from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmSenderIdTlvMADomain from JSON text
	FromJson(value string) error
}

func (obj *flowCfmSenderIdTlvMADomain) Marshal() marshalFlowCfmSenderIdTlvMADomain {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmSenderIdTlvMADomain{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmSenderIdTlvMADomain) Unmarshal() unMarshalFlowCfmSenderIdTlvMADomain {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmSenderIdTlvMADomain{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmSenderIdTlvMADomain) ToProto() (*otg.FlowCfmSenderIdTlvMADomain, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmSenderIdTlvMADomain) FromProto(msg *otg.FlowCfmSenderIdTlvMADomain) (FlowCfmSenderIdTlvMADomain, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmSenderIdTlvMADomain) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvMADomain) FromPbText(value string) error {
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

func (m *marshalflowCfmSenderIdTlvMADomain) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvMADomain) FromYaml(value string) error {
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

func (m *marshalflowCfmSenderIdTlvMADomain) ToJson() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvMADomain) FromJson(value string) error {
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

func (obj *flowCfmSenderIdTlvMADomain) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlvMADomain) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlvMADomain) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmSenderIdTlvMADomain) Clone() (FlowCfmSenderIdTlvMADomain, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmSenderIdTlvMADomain()
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

func (obj *flowCfmSenderIdTlvMADomain) setNil() {
	obj.lengthHolder = nil
	obj.managementAddressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmSenderIdTlvMADomain is sender ID TLV with Management Address Domain
type FlowCfmSenderIdTlvMADomain interface {
	Validation
	// msg marshals FlowCfmSenderIdTlvMADomain to protobuf object *otg.FlowCfmSenderIdTlvMADomain
	// and doesn't set defaults
	msg() *otg.FlowCfmSenderIdTlvMADomain
	// setMsg unmarshals FlowCfmSenderIdTlvMADomain from protobuf object *otg.FlowCfmSenderIdTlvMADomain
	// and doesn't set defaults
	setMsg(*otg.FlowCfmSenderIdTlvMADomain) FlowCfmSenderIdTlvMADomain
	// provides marshal interface
	Marshal() marshalFlowCfmSenderIdTlvMADomain
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmSenderIdTlvMADomain
	// validate validates FlowCfmSenderIdTlvMADomain
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmSenderIdTlvMADomain, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowCfmLength, set in FlowCfmSenderIdTlvMADomain.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	Length() FlowCfmLength
	// SetLength assigns FlowCfmLength provided by user to FlowCfmSenderIdTlvMADomain.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	SetLength(value FlowCfmLength) FlowCfmSenderIdTlvMADomain
	// HasLength checks if Length has been set in FlowCfmSenderIdTlvMADomain
	HasLength() bool
	// Domain returns string, set in FlowCfmSenderIdTlvMADomain.
	Domain() string
	// SetDomain assigns string provided by user to FlowCfmSenderIdTlvMADomain
	SetDomain(value string) FlowCfmSenderIdTlvMADomain
	// HasDomain checks if Domain has been set in FlowCfmSenderIdTlvMADomain
	HasDomain() bool
	// ManagementAddress returns FlowCfmSenderIdTlvMgmtAddress, set in FlowCfmSenderIdTlvMADomain.
	// FlowCfmSenderIdTlvMgmtAddress is sender ID TLV with Management Address
	ManagementAddress() FlowCfmSenderIdTlvMgmtAddress
	// SetManagementAddress assigns FlowCfmSenderIdTlvMgmtAddress provided by user to FlowCfmSenderIdTlvMADomain.
	// FlowCfmSenderIdTlvMgmtAddress is sender ID TLV with Management Address
	SetManagementAddress(value FlowCfmSenderIdTlvMgmtAddress) FlowCfmSenderIdTlvMADomain
	// HasManagementAddress checks if ManagementAddress has been set in FlowCfmSenderIdTlvMADomain
	HasManagementAddress() bool
	setNil()
}

// description is TBD
// Length returns a FlowCfmLength
func (obj *flowCfmSenderIdTlvMADomain) Length() FlowCfmLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowCfmLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowCfmLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a FlowCfmLength
func (obj *flowCfmSenderIdTlvMADomain) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the FlowCfmLength value in the FlowCfmSenderIdTlvMADomain object
func (obj *flowCfmSenderIdTlvMADomain) SetLength(value FlowCfmLength) FlowCfmSenderIdTlvMADomain {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// The domain is a full BER-encoded OID identifying the transport/protocol, CFM follows the SNMPv3 RFC 3411.
// Domain returns a string
func (obj *flowCfmSenderIdTlvMADomain) Domain() string {

	return *obj.obj.Domain

}

// The domain is a full BER-encoded OID identifying the transport/protocol, CFM follows the SNMPv3 RFC 3411.
// Domain returns a string
func (obj *flowCfmSenderIdTlvMADomain) HasDomain() bool {
	return obj.obj.Domain != nil
}

// The domain is a full BER-encoded OID identifying the transport/protocol, CFM follows the SNMPv3 RFC 3411.
// SetDomain sets the string value in the FlowCfmSenderIdTlvMADomain object
func (obj *flowCfmSenderIdTlvMADomain) SetDomain(value string) FlowCfmSenderIdTlvMADomain {

	obj.obj.Domain = &value
	return obj
}

// Management Address of the sender. When absent, the Management Address Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// ManagementAddress returns a FlowCfmSenderIdTlvMgmtAddress
func (obj *flowCfmSenderIdTlvMADomain) ManagementAddress() FlowCfmSenderIdTlvMgmtAddress {
	if obj.obj.ManagementAddress == nil {
		obj.obj.ManagementAddress = NewFlowCfmSenderIdTlvMgmtAddress().msg()
	}
	if obj.managementAddressHolder == nil {
		obj.managementAddressHolder = &flowCfmSenderIdTlvMgmtAddress{obj: obj.obj.ManagementAddress}
	}
	return obj.managementAddressHolder
}

// Management Address of the sender. When absent, the Management Address Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// ManagementAddress returns a FlowCfmSenderIdTlvMgmtAddress
func (obj *flowCfmSenderIdTlvMADomain) HasManagementAddress() bool {
	return obj.obj.ManagementAddress != nil
}

// Management Address of the sender. When absent, the Management Address Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// SetManagementAddress sets the FlowCfmSenderIdTlvMgmtAddress value in the FlowCfmSenderIdTlvMADomain object
func (obj *flowCfmSenderIdTlvMADomain) SetManagementAddress(value FlowCfmSenderIdTlvMgmtAddress) FlowCfmSenderIdTlvMADomain {

	obj.managementAddressHolder = nil
	obj.obj.ManagementAddress = value.msg()

	return obj
}

func (obj *flowCfmSenderIdTlvMADomain) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Domain != nil {

		err := obj.validateHex(obj.Domain())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmSenderIdTlvMADomain.Domain"))
		}

	}

	if obj.obj.ManagementAddress != nil {

		obj.ManagementAddress().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmSenderIdTlvMADomain) setDefault() {
	if obj.obj.Domain == nil {
		obj.SetDomain("00")
	}

}
