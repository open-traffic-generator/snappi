package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpIpv4Interface *****
type rsvpIpv4Interface struct {
	validation
	obj          *otg.RsvpIpv4Interface
	marshaller   marshalRsvpIpv4Interface
	unMarshaller unMarshalRsvpIpv4Interface
}

func NewRsvpIpv4Interface() RsvpIpv4Interface {
	obj := rsvpIpv4Interface{obj: &otg.RsvpIpv4Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpIpv4Interface) msg() *otg.RsvpIpv4Interface {
	return obj.obj
}

func (obj *rsvpIpv4Interface) setMsg(msg *otg.RsvpIpv4Interface) RsvpIpv4Interface {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpIpv4Interface struct {
	obj *rsvpIpv4Interface
}

type marshalRsvpIpv4Interface interface {
	// ToProto marshals RsvpIpv4Interface to protobuf object *otg.RsvpIpv4Interface
	ToProto() (*otg.RsvpIpv4Interface, error)
	// ToPbText marshals RsvpIpv4Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpIpv4Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpIpv4Interface to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpIpv4Interface struct {
	obj *rsvpIpv4Interface
}

type unMarshalRsvpIpv4Interface interface {
	// FromProto unmarshals RsvpIpv4Interface from protobuf object *otg.RsvpIpv4Interface
	FromProto(msg *otg.RsvpIpv4Interface) (RsvpIpv4Interface, error)
	// FromPbText unmarshals RsvpIpv4Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpIpv4Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpIpv4Interface from JSON text
	FromJson(value string) error
}

func (obj *rsvpIpv4Interface) Marshal() marshalRsvpIpv4Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpIpv4Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpIpv4Interface) Unmarshal() unMarshalRsvpIpv4Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpIpv4Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpIpv4Interface) ToProto() (*otg.RsvpIpv4Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpIpv4Interface) FromProto(msg *otg.RsvpIpv4Interface) (RsvpIpv4Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpIpv4Interface) ToPbText() (string, error) {
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

func (m *unMarshalrsvpIpv4Interface) FromPbText(value string) error {
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

func (m *marshalrsvpIpv4Interface) ToYaml() (string, error) {
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

func (m *unMarshalrsvpIpv4Interface) FromYaml(value string) error {
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

func (m *marshalrsvpIpv4Interface) ToJson() (string, error) {
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

func (m *unMarshalrsvpIpv4Interface) FromJson(value string) error {
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

func (obj *rsvpIpv4Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpIpv4Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpIpv4Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpIpv4Interface) Clone() (RsvpIpv4Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpIpv4Interface()
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

// RsvpIpv4Interface is configuration for RSVP Interface.
type RsvpIpv4Interface interface {
	Validation
	// msg marshals RsvpIpv4Interface to protobuf object *otg.RsvpIpv4Interface
	// and doesn't set defaults
	msg() *otg.RsvpIpv4Interface
	// setMsg unmarshals RsvpIpv4Interface from protobuf object *otg.RsvpIpv4Interface
	// and doesn't set defaults
	setMsg(*otg.RsvpIpv4Interface) RsvpIpv4Interface
	// provides marshal interface
	Marshal() marshalRsvpIpv4Interface
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpIpv4Interface
	// validate validates RsvpIpv4Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpIpv4Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Name returns string, set in RsvpIpv4Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to RsvpIpv4Interface
	SetIpv4Name(value string) RsvpIpv4Interface
	// NeighborIp returns string, set in RsvpIpv4Interface.
	NeighborIp() string
	// SetNeighborIp assigns string provided by user to RsvpIpv4Interface
	SetNeighborIp(value string) RsvpIpv4Interface
	// LabelSpaceStart returns uint32, set in RsvpIpv4Interface.
	LabelSpaceStart() uint32
	// SetLabelSpaceStart assigns uint32 provided by user to RsvpIpv4Interface
	SetLabelSpaceStart(value uint32) RsvpIpv4Interface
	// HasLabelSpaceStart checks if LabelSpaceStart has been set in RsvpIpv4Interface
	HasLabelSpaceStart() bool
	// LabelSpaceEnd returns uint32, set in RsvpIpv4Interface.
	LabelSpaceEnd() uint32
	// SetLabelSpaceEnd assigns uint32 provided by user to RsvpIpv4Interface
	SetLabelSpaceEnd(value uint32) RsvpIpv4Interface
	// HasLabelSpaceEnd checks if LabelSpaceEnd has been set in RsvpIpv4Interface
	HasLabelSpaceEnd() bool
	// EnableRefreshReduction returns bool, set in RsvpIpv4Interface.
	EnableRefreshReduction() bool
	// SetEnableRefreshReduction assigns bool provided by user to RsvpIpv4Interface
	SetEnableRefreshReduction(value bool) RsvpIpv4Interface
	// HasEnableRefreshReduction checks if EnableRefreshReduction has been set in RsvpIpv4Interface
	HasEnableRefreshReduction() bool
	// SummaryRefreshInterval returns uint32, set in RsvpIpv4Interface.
	SummaryRefreshInterval() uint32
	// SetSummaryRefreshInterval assigns uint32 provided by user to RsvpIpv4Interface
	SetSummaryRefreshInterval(value uint32) RsvpIpv4Interface
	// HasSummaryRefreshInterval checks if SummaryRefreshInterval has been set in RsvpIpv4Interface
	HasSummaryRefreshInterval() bool
	// SendBundle returns bool, set in RsvpIpv4Interface.
	SendBundle() bool
	// SetSendBundle assigns bool provided by user to RsvpIpv4Interface
	SetSendBundle(value bool) RsvpIpv4Interface
	// HasSendBundle checks if SendBundle has been set in RsvpIpv4Interface
	HasSendBundle() bool
	// BundleThreshold returns uint32, set in RsvpIpv4Interface.
	BundleThreshold() uint32
	// SetBundleThreshold assigns uint32 provided by user to RsvpIpv4Interface
	SetBundleThreshold(value uint32) RsvpIpv4Interface
	// HasBundleThreshold checks if BundleThreshold has been set in RsvpIpv4Interface
	HasBundleThreshold() bool
	// EnableHello returns bool, set in RsvpIpv4Interface.
	EnableHello() bool
	// SetEnableHello assigns bool provided by user to RsvpIpv4Interface
	SetEnableHello(value bool) RsvpIpv4Interface
	// HasEnableHello checks if EnableHello has been set in RsvpIpv4Interface
	HasEnableHello() bool
	// HelloInterval returns uint32, set in RsvpIpv4Interface.
	HelloInterval() uint32
	// SetHelloInterval assigns uint32 provided by user to RsvpIpv4Interface
	SetHelloInterval(value uint32) RsvpIpv4Interface
	// HasHelloInterval checks if HelloInterval has been set in RsvpIpv4Interface
	HasHelloInterval() bool
	// TimeoutMultiplier returns uint32, set in RsvpIpv4Interface.
	TimeoutMultiplier() uint32
	// SetTimeoutMultiplier assigns uint32 provided by user to RsvpIpv4Interface
	SetTimeoutMultiplier(value uint32) RsvpIpv4Interface
	// HasTimeoutMultiplier checks if TimeoutMultiplier has been set in RsvpIpv4Interface
	HasTimeoutMultiplier() bool
}

// The globally unique name of the IPv4 interface connected to the DUT. This name must match the "name" field of the "ipv4_addresses" on top which this RSVP interface is configured.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *rsvpIpv4Interface) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The globally unique name of the IPv4 interface connected to the DUT. This name must match the "name" field of the "ipv4_addresses" on top which this RSVP interface is configured.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// SetIpv4Name sets the string value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetIpv4Name(value string) RsvpIpv4Interface {

	obj.obj.Ipv4Name = &value
	return obj
}

// IPv4 address of the RSVP neighbor on this interface.
// NeighborIp returns a string
func (obj *rsvpIpv4Interface) NeighborIp() string {

	return *obj.obj.NeighborIp

}

// IPv4 address of the RSVP neighbor on this interface.
// SetNeighborIp sets the string value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetNeighborIp(value string) RsvpIpv4Interface {

	obj.obj.NeighborIp = &value
	return obj
}

// The user-defined label space start value. The LSPs for which this router acts as a egress are assigned labels from this label pool.The"label_space_start" and "label_space_end" together defines this label-pool.
// LabelSpaceStart returns a uint32
func (obj *rsvpIpv4Interface) LabelSpaceStart() uint32 {

	return *obj.obj.LabelSpaceStart

}

// The user-defined label space start value. The LSPs for which this router acts as a egress are assigned labels from this label pool.The"label_space_start" and "label_space_end" together defines this label-pool.
// LabelSpaceStart returns a uint32
func (obj *rsvpIpv4Interface) HasLabelSpaceStart() bool {
	return obj.obj.LabelSpaceStart != nil
}

// The user-defined label space start value. The LSPs for which this router acts as a egress are assigned labels from this label pool.The"label_space_start" and "label_space_end" together defines this label-pool.
// SetLabelSpaceStart sets the uint32 value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetLabelSpaceStart(value uint32) RsvpIpv4Interface {

	obj.obj.LabelSpaceStart = &value
	return obj
}

// The user-defined label space end value.The last label value that can be assigned to the LSPs for which this router acts as egress.
// LabelSpaceEnd returns a uint32
func (obj *rsvpIpv4Interface) LabelSpaceEnd() uint32 {

	return *obj.obj.LabelSpaceEnd

}

// The user-defined label space end value.The last label value that can be assigned to the LSPs for which this router acts as egress.
// LabelSpaceEnd returns a uint32
func (obj *rsvpIpv4Interface) HasLabelSpaceEnd() bool {
	return obj.obj.LabelSpaceEnd != nil
}

// The user-defined label space end value.The last label value that can be assigned to the LSPs for which this router acts as egress.
// SetLabelSpaceEnd sets the uint32 value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetLabelSpaceEnd(value uint32) RsvpIpv4Interface {

	obj.obj.LabelSpaceEnd = &value
	return obj
}

// Enables sending of Refresh Reduction as described in RFC2961.
// EnableRefreshReduction returns a bool
func (obj *rsvpIpv4Interface) EnableRefreshReduction() bool {

	return *obj.obj.EnableRefreshReduction

}

// Enables sending of Refresh Reduction as described in RFC2961.
// EnableRefreshReduction returns a bool
func (obj *rsvpIpv4Interface) HasEnableRefreshReduction() bool {
	return obj.obj.EnableRefreshReduction != nil
}

// Enables sending of Refresh Reduction as described in RFC2961.
// SetEnableRefreshReduction sets the bool value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetEnableRefreshReduction(value bool) RsvpIpv4Interface {

	obj.obj.EnableRefreshReduction = &value
	return obj
}

// The number of seconds between transmissions of successive Summary Refreshes. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// SummaryRefreshInterval returns a uint32
func (obj *rsvpIpv4Interface) SummaryRefreshInterval() uint32 {

	return *obj.obj.SummaryRefreshInterval

}

// The number of seconds between transmissions of successive Summary Refreshes. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// SummaryRefreshInterval returns a uint32
func (obj *rsvpIpv4Interface) HasSummaryRefreshInterval() bool {
	return obj.obj.SummaryRefreshInterval != nil
}

// The number of seconds between transmissions of successive Summary Refreshes. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// SetSummaryRefreshInterval sets the uint32 value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetSummaryRefreshInterval(value uint32) RsvpIpv4Interface {

	obj.obj.SummaryRefreshInterval = &value
	return obj
}

// Enables aggregration of different RSVP messages within a single PDU.
// SendBundle returns a bool
func (obj *rsvpIpv4Interface) SendBundle() bool {

	return *obj.obj.SendBundle

}

// Enables aggregration of different RSVP messages within a single PDU.
// SendBundle returns a bool
func (obj *rsvpIpv4Interface) HasSendBundle() bool {
	return obj.obj.SendBundle != nil
}

// Enables aggregration of different RSVP messages within a single PDU.
// SetSendBundle sets the bool value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetSendBundle(value bool) RsvpIpv4Interface {

	obj.obj.SendBundle = &value
	return obj
}

// The number of milliseconds to wait after which RSVP will bundle different RSVP messages and transmit Bundle messages.
// BundleThreshold returns a uint32
func (obj *rsvpIpv4Interface) BundleThreshold() uint32 {

	return *obj.obj.BundleThreshold

}

// The number of milliseconds to wait after which RSVP will bundle different RSVP messages and transmit Bundle messages.
// BundleThreshold returns a uint32
func (obj *rsvpIpv4Interface) HasBundleThreshold() bool {
	return obj.obj.BundleThreshold != nil
}

// The number of milliseconds to wait after which RSVP will bundle different RSVP messages and transmit Bundle messages.
// SetBundleThreshold sets the uint32 value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetBundleThreshold(value uint32) RsvpIpv4Interface {

	obj.obj.BundleThreshold = &value
	return obj
}

// Enables sending of Hello Messages as per RFC3209.
// EnableHello returns a bool
func (obj *rsvpIpv4Interface) EnableHello() bool {

	return *obj.obj.EnableHello

}

// Enables sending of Hello Messages as per RFC3209.
// EnableHello returns a bool
func (obj *rsvpIpv4Interface) HasEnableHello() bool {
	return obj.obj.EnableHello != nil
}

// Enables sending of Hello Messages as per RFC3209.
// SetEnableHello sets the bool value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetEnableHello(value bool) RsvpIpv4Interface {

	obj.obj.EnableHello = &value
	return obj
}

// If enable_hello is set to 'true', this specifies the minimum hello interval in seconds at which successive Hello Messages  are sent as per RFC3209. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// HelloInterval returns a uint32
func (obj *rsvpIpv4Interface) HelloInterval() uint32 {

	return *obj.obj.HelloInterval

}

// If enable_hello is set to 'true', this specifies the minimum hello interval in seconds at which successive Hello Messages  are sent as per RFC3209. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// HelloInterval returns a uint32
func (obj *rsvpIpv4Interface) HasHelloInterval() bool {
	return obj.obj.HelloInterval != nil
}

// If enable_hello is set to 'true', this specifies the minimum hello interval in seconds at which successive Hello Messages  are sent as per RFC3209. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// SetHelloInterval sets the uint32 value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetHelloInterval(value uint32) RsvpIpv4Interface {

	obj.obj.HelloInterval = &value
	return obj
}

// The number of missed hellos after which the node should consider RSVP Neighbor to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// TimeoutMultiplier returns a uint32
func (obj *rsvpIpv4Interface) TimeoutMultiplier() uint32 {

	return *obj.obj.TimeoutMultiplier

}

// The number of missed hellos after which the node should consider RSVP Neighbor to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// TimeoutMultiplier returns a uint32
func (obj *rsvpIpv4Interface) HasTimeoutMultiplier() bool {
	return obj.obj.TimeoutMultiplier != nil
}

// The number of missed hellos after which the node should consider RSVP Neighbor to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// SetTimeoutMultiplier sets the uint32 value in the RsvpIpv4Interface object
func (obj *rsvpIpv4Interface) SetTimeoutMultiplier(value uint32) RsvpIpv4Interface {

	obj.obj.TimeoutMultiplier = &value
	return obj
}

func (obj *rsvpIpv4Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface RsvpIpv4Interface")
	}

	// NeighborIp is required
	if obj.obj.NeighborIp == nil {
		vObj.validationErrors = append(vObj.validationErrors, "NeighborIp is required field on interface RsvpIpv4Interface")
	}
	if obj.obj.NeighborIp != nil {

		err := obj.validateIpv4(obj.NeighborIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RsvpIpv4Interface.NeighborIp"))
		}

	}

	if obj.obj.LabelSpaceStart != nil {

		if *obj.obj.LabelSpaceStart > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpIpv4Interface.LabelSpaceStart <= 1048575 but Got %d", *obj.obj.LabelSpaceStart))
		}

	}

	if obj.obj.LabelSpaceEnd != nil {

		if *obj.obj.LabelSpaceEnd > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpIpv4Interface.LabelSpaceEnd <= 1048575 but Got %d", *obj.obj.LabelSpaceEnd))
		}

	}

	if obj.obj.SummaryRefreshInterval != nil {

		if *obj.obj.SummaryRefreshInterval > 3600 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpIpv4Interface.SummaryRefreshInterval <= 3600 but Got %d", *obj.obj.SummaryRefreshInterval))
		}

	}

	if obj.obj.BundleThreshold != nil {

		if *obj.obj.BundleThreshold > 1000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpIpv4Interface.BundleThreshold <= 1000 but Got %d", *obj.obj.BundleThreshold))
		}

	}

	if obj.obj.HelloInterval != nil {

		if *obj.obj.HelloInterval > 3600 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpIpv4Interface.HelloInterval <= 3600 but Got %d", *obj.obj.HelloInterval))
		}

	}

	if obj.obj.TimeoutMultiplier != nil {

		if *obj.obj.TimeoutMultiplier > 10 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpIpv4Interface.TimeoutMultiplier <= 10 but Got %d", *obj.obj.TimeoutMultiplier))
		}

	}

}

func (obj *rsvpIpv4Interface) setDefault() {
	if obj.obj.LabelSpaceStart == nil {
		obj.SetLabelSpaceStart(1000)
	}
	if obj.obj.LabelSpaceEnd == nil {
		obj.SetLabelSpaceEnd(100000)
	}
	if obj.obj.EnableRefreshReduction == nil {
		obj.SetEnableRefreshReduction(false)
	}
	if obj.obj.SummaryRefreshInterval == nil {
		obj.SetSummaryRefreshInterval(30)
	}
	if obj.obj.SendBundle == nil {
		obj.SetSendBundle(false)
	}
	if obj.obj.BundleThreshold == nil {
		obj.SetBundleThreshold(50)
	}
	if obj.obj.EnableHello == nil {
		obj.SetEnableHello(false)
	}
	if obj.obj.HelloInterval == nil {
		obj.SetHelloInterval(9)
	}
	if obj.obj.TimeoutMultiplier == nil {
		obj.SetTimeoutMultiplier(3)
	}

}
