package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4ClientParams *****
type dhcpv4ClientParams struct {
	validation
	obj          *otg.Dhcpv4ClientParams
	marshaller   marshalDhcpv4ClientParams
	unMarshaller unMarshalDhcpv4ClientParams
}

func NewDhcpv4ClientParams() Dhcpv4ClientParams {
	obj := dhcpv4ClientParams{obj: &otg.Dhcpv4ClientParams{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4ClientParams) msg() *otg.Dhcpv4ClientParams {
	return obj.obj
}

func (obj *dhcpv4ClientParams) setMsg(msg *otg.Dhcpv4ClientParams) Dhcpv4ClientParams {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4ClientParams struct {
	obj *dhcpv4ClientParams
}

type marshalDhcpv4ClientParams interface {
	// ToProto marshals Dhcpv4ClientParams to protobuf object *otg.Dhcpv4ClientParams
	ToProto() (*otg.Dhcpv4ClientParams, error)
	// ToPbText marshals Dhcpv4ClientParams to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4ClientParams to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4ClientParams to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv4ClientParams to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv4ClientParams struct {
	obj *dhcpv4ClientParams
}

type unMarshalDhcpv4ClientParams interface {
	// FromProto unmarshals Dhcpv4ClientParams from protobuf object *otg.Dhcpv4ClientParams
	FromProto(msg *otg.Dhcpv4ClientParams) (Dhcpv4ClientParams, error)
	// FromPbText unmarshals Dhcpv4ClientParams from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4ClientParams from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4ClientParams from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4ClientParams) Marshal() marshalDhcpv4ClientParams {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4ClientParams{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4ClientParams) Unmarshal() unMarshalDhcpv4ClientParams {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4ClientParams{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4ClientParams) ToProto() (*otg.Dhcpv4ClientParams, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4ClientParams) FromProto(msg *otg.Dhcpv4ClientParams) (Dhcpv4ClientParams, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4ClientParams) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4ClientParams) FromPbText(value string) error {
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

func (m *marshaldhcpv4ClientParams) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4ClientParams) FromYaml(value string) error {
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

func (m *marshaldhcpv4ClientParams) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv4ClientParams) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4ClientParams) FromJson(value string) error {
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

func (obj *dhcpv4ClientParams) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4ClientParams) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4ClientParams) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4ClientParams) Clone() (Dhcpv4ClientParams, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4ClientParams()
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

// Dhcpv4ClientParams is configuration Parameter request list by emulated DHCPv4 Client.
type Dhcpv4ClientParams interface {
	Validation
	// msg marshals Dhcpv4ClientParams to protobuf object *otg.Dhcpv4ClientParams
	// and doesn't set defaults
	msg() *otg.Dhcpv4ClientParams
	// setMsg unmarshals Dhcpv4ClientParams from protobuf object *otg.Dhcpv4ClientParams
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4ClientParams) Dhcpv4ClientParams
	// provides marshal interface
	Marshal() marshalDhcpv4ClientParams
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4ClientParams
	// validate validates Dhcpv4ClientParams
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4ClientParams, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SubnetMask returns bool, set in Dhcpv4ClientParams.
	SubnetMask() bool
	// SetSubnetMask assigns bool provided by user to Dhcpv4ClientParams
	SetSubnetMask(value bool) Dhcpv4ClientParams
	// HasSubnetMask checks if SubnetMask has been set in Dhcpv4ClientParams
	HasSubnetMask() bool
	// Router returns bool, set in Dhcpv4ClientParams.
	Router() bool
	// SetRouter assigns bool provided by user to Dhcpv4ClientParams
	SetRouter(value bool) Dhcpv4ClientParams
	// HasRouter checks if Router has been set in Dhcpv4ClientParams
	HasRouter() bool
	// RenewalTimer returns bool, set in Dhcpv4ClientParams.
	RenewalTimer() bool
	// SetRenewalTimer assigns bool provided by user to Dhcpv4ClientParams
	SetRenewalTimer(value bool) Dhcpv4ClientParams
	// HasRenewalTimer checks if RenewalTimer has been set in Dhcpv4ClientParams
	HasRenewalTimer() bool
	// RebindingTimer returns bool, set in Dhcpv4ClientParams.
	RebindingTimer() bool
	// SetRebindingTimer assigns bool provided by user to Dhcpv4ClientParams
	SetRebindingTimer(value bool) Dhcpv4ClientParams
	// HasRebindingTimer checks if RebindingTimer has been set in Dhcpv4ClientParams
	HasRebindingTimer() bool
}

// Request for the subnet mask option specifies the client's subnet mask as per RFC950.
// SubnetMask returns a bool
func (obj *dhcpv4ClientParams) SubnetMask() bool {

	return *obj.obj.SubnetMask

}

// Request for the subnet mask option specifies the client's subnet mask as per RFC950.
// SubnetMask returns a bool
func (obj *dhcpv4ClientParams) HasSubnetMask() bool {
	return obj.obj.SubnetMask != nil
}

// Request for the subnet mask option specifies the client's subnet mask as per RFC950.
// SetSubnetMask sets the bool value in the Dhcpv4ClientParams object
func (obj *dhcpv4ClientParams) SetSubnetMask(value bool) Dhcpv4ClientParams {

	obj.obj.SubnetMask = &value
	return obj
}

// Request for the router option that specifies a list of IP addresses for routers on the client's subnet.
// Router returns a bool
func (obj *dhcpv4ClientParams) Router() bool {

	return *obj.obj.Router

}

// Request for the router option that specifies a list of IP addresses for routers on the client's subnet.
// Router returns a bool
func (obj *dhcpv4ClientParams) HasRouter() bool {
	return obj.obj.Router != nil
}

// Request for the router option that specifies a list of IP addresses for routers on the client's subnet.
// SetRouter sets the bool value in the Dhcpv4ClientParams object
func (obj *dhcpv4ClientParams) SetRouter(value bool) Dhcpv4ClientParams {

	obj.obj.Router = &value
	return obj
}

// Request for the renewal timer, T1. When the timer expires, the client transitions from the BOUND state to the RENEWING state.
// RenewalTimer returns a bool
func (obj *dhcpv4ClientParams) RenewalTimer() bool {

	return *obj.obj.RenewalTimer

}

// Request for the renewal timer, T1. When the timer expires, the client transitions from the BOUND state to the RENEWING state.
// RenewalTimer returns a bool
func (obj *dhcpv4ClientParams) HasRenewalTimer() bool {
	return obj.obj.RenewalTimer != nil
}

// Request for the renewal timer, T1. When the timer expires, the client transitions from the BOUND state to the RENEWING state.
// SetRenewalTimer sets the bool value in the Dhcpv4ClientParams object
func (obj *dhcpv4ClientParams) SetRenewalTimer(value bool) Dhcpv4ClientParams {

	obj.obj.RenewalTimer = &value
	return obj
}

// Request for the rebinding timer (T2). When expires, the client transitions to the REBINDING state.
// RebindingTimer returns a bool
func (obj *dhcpv4ClientParams) RebindingTimer() bool {

	return *obj.obj.RebindingTimer

}

// Request for the rebinding timer (T2). When expires, the client transitions to the REBINDING state.
// RebindingTimer returns a bool
func (obj *dhcpv4ClientParams) HasRebindingTimer() bool {
	return obj.obj.RebindingTimer != nil
}

// Request for the rebinding timer (T2). When expires, the client transitions to the REBINDING state.
// SetRebindingTimer sets the bool value in the Dhcpv4ClientParams object
func (obj *dhcpv4ClientParams) SetRebindingTimer(value bool) Dhcpv4ClientParams {

	obj.obj.RebindingTimer = &value
	return obj
}

func (obj *dhcpv4ClientParams) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv4ClientParams) setDefault() {
	if obj.obj.SubnetMask == nil {
		obj.SetSubnetMask(true)
	}
	if obj.obj.Router == nil {
		obj.SetRouter(true)
	}
	if obj.obj.RenewalTimer == nil {
		obj.SetRenewalTimer(false)
	}
	if obj.obj.RebindingTimer == nil {
		obj.SetRebindingTimer(false)
	}

}
