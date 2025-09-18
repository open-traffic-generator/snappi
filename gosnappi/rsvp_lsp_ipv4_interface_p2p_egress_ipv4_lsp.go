package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpLspIpv4InterfaceP2PEgressIpv4Lsp *****
type rsvpLspIpv4InterfaceP2PEgressIpv4Lsp struct {
	validation
	obj          *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	marshaller   marshalRsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	unMarshaller unMarshalRsvpLspIpv4InterfaceP2PEgressIpv4Lsp
}

func NewRsvpLspIpv4InterfaceP2PEgressIpv4Lsp() RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {
	obj := rsvpLspIpv4InterfaceP2PEgressIpv4Lsp{obj: &otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) msg() *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {
	return obj.obj
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) setMsg(msg *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp struct {
	obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp
}

type marshalRsvpLspIpv4InterfaceP2PEgressIpv4Lsp interface {
	// ToProto marshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp to protobuf object *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	ToProto() (*otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp, error)
	// ToPbText marshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp struct {
	obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp
}

type unMarshalRsvpLspIpv4InterfaceP2PEgressIpv4Lsp interface {
	// FromProto unmarshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp from protobuf object *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	FromProto(msg *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp) (RsvpLspIpv4InterfaceP2PEgressIpv4Lsp, error)
	// FromPbText unmarshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp from JSON text
	FromJson(value string) error
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) Marshal() marshalRsvpLspIpv4InterfaceP2PEgressIpv4Lsp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) Unmarshal() unMarshalRsvpLspIpv4InterfaceP2PEgressIpv4Lsp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp) ToProto() (*otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp) FromProto(msg *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp) (RsvpLspIpv4InterfaceP2PEgressIpv4Lsp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp) ToPbText() (string, error) {
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

func (m *unMarshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp) FromPbText(value string) error {
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

func (m *marshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp) ToYaml() (string, error) {
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

func (m *unMarshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp) FromYaml(value string) error {
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

func (m *marshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp) ToJson() (string, error) {
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

func (m *unMarshalrsvpLspIpv4InterfaceP2PEgressIpv4Lsp) FromJson(value string) error {
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

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) Clone() (RsvpLspIpv4InterfaceP2PEgressIpv4Lsp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpLspIpv4InterfaceP2PEgressIpv4Lsp()
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

// RsvpLspIpv4InterfaceP2PEgressIpv4Lsp is configuration for RSVP Egress Point-to-Point(P2P) IPv4 LSPs.
type RsvpLspIpv4InterfaceP2PEgressIpv4Lsp interface {
	Validation
	// msg marshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp to protobuf object *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// and doesn't set defaults
	msg() *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// setMsg unmarshals RsvpLspIpv4InterfaceP2PEgressIpv4Lsp from protobuf object *otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// and doesn't set defaults
	setMsg(*otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// provides marshal interface
	Marshal() marshalRsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// validate validates RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpLspIpv4InterfaceP2PEgressIpv4Lsp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.
	Name() string
	// SetName assigns string provided by user to RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	SetName(value string) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// RefreshInterval returns uint32, set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.
	RefreshInterval() uint32
	// SetRefreshInterval assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	SetRefreshInterval(value uint32) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// HasRefreshInterval checks if RefreshInterval has been set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	HasRefreshInterval() bool
	// TimeoutMultiplier returns uint32, set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.
	TimeoutMultiplier() uint32
	// SetTimeoutMultiplier assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	SetTimeoutMultiplier(value uint32) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// HasTimeoutMultiplier checks if TimeoutMultiplier has been set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	HasTimeoutMultiplier() bool
	// ReservationStyle returns RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum, set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	ReservationStyle() RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum
	// SetReservationStyle assigns RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum provided by user to RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	SetReservationStyle(value RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// HasReservationStyle checks if ReservationStyle has been set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	HasReservationStyle() bool
	// EnableFixedLabel returns bool, set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.
	EnableFixedLabel() bool
	// SetEnableFixedLabel assigns bool provided by user to RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	SetEnableFixedLabel(value bool) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// HasEnableFixedLabel checks if EnableFixedLabel has been set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	HasEnableFixedLabel() bool
	// FixedLabelValue returns uint32, set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.
	FixedLabelValue() uint32
	// SetFixedLabelValue assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	SetFixedLabelValue(value uint32) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// HasFixedLabelValue checks if FixedLabelValue has been set in RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	HasFixedLabelValue() bool
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the RsvpLspIpv4InterfaceP2PEgressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) SetName(value string) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {

	obj.obj.Name = &value
	return obj
}

// The time in seconds between successive transmissions of RESV Refreshes. The actual refresh interval is jittered by upto 50%. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// RefreshInterval returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) RefreshInterval() uint32 {

	return *obj.obj.RefreshInterval

}

// The time in seconds between successive transmissions of RESV Refreshes. The actual refresh interval is jittered by upto 50%. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// RefreshInterval returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) HasRefreshInterval() bool {
	return obj.obj.RefreshInterval != nil
}

// The time in seconds between successive transmissions of RESV Refreshes. The actual refresh interval is jittered by upto 50%. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// SetRefreshInterval sets the uint32 value in the RsvpLspIpv4InterfaceP2PEgressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) SetRefreshInterval(value uint32) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {

	obj.obj.RefreshInterval = &value
	return obj
}

// The number of missed PATH refreshes after which a recieving node should consider the LSP state to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// TimeoutMultiplier returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) TimeoutMultiplier() uint32 {

	return *obj.obj.TimeoutMultiplier

}

// The number of missed PATH refreshes after which a recieving node should consider the LSP state to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// TimeoutMultiplier returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) HasTimeoutMultiplier() bool {
	return obj.obj.TimeoutMultiplier != nil
}

// The number of missed PATH refreshes after which a recieving node should consider the LSP state to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// SetTimeoutMultiplier sets the uint32 value in the RsvpLspIpv4InterfaceP2PEgressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) SetTimeoutMultiplier(value uint32) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {

	obj.obj.TimeoutMultiplier = &value
	return obj
}

type RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum string

// Enum of ReservationStyle on RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
var RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyle = struct {
	SHARED_EXPLICIT RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum
	FIXED_FILTER    RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum
	AUTO            RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum
}{
	SHARED_EXPLICIT: RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum("shared_explicit"),
	FIXED_FILTER:    RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum("fixed_filter"),
	AUTO:            RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum("auto"),
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) ReservationStyle() RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum {
	return RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum(obj.obj.ReservationStyle.Enum().String())
}

// It determines how RSVP-TE enabled network devices set up reservations along the path between an end-to-end  QOS-enabled connection. If 'auto' is enabled, the style is chosen based on whether the incoming Path has 'SE Desired' flag set. Otherwise, the style is chosen based on the value selected for this attribute.
// ReservationStyle returns a string
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) HasReservationStyle() bool {
	return obj.obj.ReservationStyle != nil
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) SetReservationStyle(value RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {
	intValue, ok := otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp_ReservationStyle_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyleEnum", string(value)))
		return obj
	}
	enumValue := otg.RsvpLspIpv4InterfaceP2PEgressIpv4Lsp_ReservationStyle_Enum(intValue)
	obj.obj.ReservationStyle = &enumValue

	return obj
}

// If enabled, a specific fixed label will be advertised by the egress or tail end for all Path messages received by this egress.  This can be leveraged to advertise Explicit or Implicit null labels.
// EnableFixedLabel returns a bool
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) EnableFixedLabel() bool {

	return *obj.obj.EnableFixedLabel

}

// If enabled, a specific fixed label will be advertised by the egress or tail end for all Path messages received by this egress.  This can be leveraged to advertise Explicit or Implicit null labels.
// EnableFixedLabel returns a bool
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) HasEnableFixedLabel() bool {
	return obj.obj.EnableFixedLabel != nil
}

// If enabled, a specific fixed label will be advertised by the egress or tail end for all Path messages received by this egress.  This can be leveraged to advertise Explicit or Implicit null labels.
// SetEnableFixedLabel sets the bool value in the RsvpLspIpv4InterfaceP2PEgressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) SetEnableFixedLabel(value bool) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {

	obj.obj.EnableFixedLabel = &value
	return obj
}

// The fixed label value as advertised by egress in RESV message. Applicable only if 'fixed_label' is set to 'true'. Special values are '0 - IPv4 Explicit NULL', '2 - IPv6 Explicit NULL' and '3 - Implicit NULL'.   Outside of this, labels are expected to have a minimum value of 16.
// FixedLabelValue returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) FixedLabelValue() uint32 {

	return *obj.obj.FixedLabelValue

}

// The fixed label value as advertised by egress in RESV message. Applicable only if 'fixed_label' is set to 'true'. Special values are '0 - IPv4 Explicit NULL', '2 - IPv6 Explicit NULL' and '3 - Implicit NULL'.   Outside of this, labels are expected to have a minimum value of 16.
// FixedLabelValue returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) HasFixedLabelValue() bool {
	return obj.obj.FixedLabelValue != nil
}

// The fixed label value as advertised by egress in RESV message. Applicable only if 'fixed_label' is set to 'true'. Special values are '0 - IPv4 Explicit NULL', '2 - IPv6 Explicit NULL' and '3 - Implicit NULL'.   Outside of this, labels are expected to have a minimum value of 16.
// SetFixedLabelValue sets the uint32 value in the RsvpLspIpv4InterfaceP2PEgressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) SetFixedLabelValue(value uint32) RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {

	obj.obj.FixedLabelValue = &value
	return obj
}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface RsvpLspIpv4InterfaceP2PEgressIpv4Lsp")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.RefreshInterval != nil {

		if *obj.obj.RefreshInterval > 3600 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.RefreshInterval <= 3600 but Got %d", *obj.obj.RefreshInterval))
		}

	}

	if obj.obj.TimeoutMultiplier != nil {

		if *obj.obj.TimeoutMultiplier > 10 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.TimeoutMultiplier <= 10 but Got %d", *obj.obj.TimeoutMultiplier))
		}

	}

	if obj.obj.FixedLabelValue != nil {

		if *obj.obj.FixedLabelValue > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpLspIpv4InterfaceP2PEgressIpv4Lsp.FixedLabelValue <= 1048575 but Got %d", *obj.obj.FixedLabelValue))
		}

	}

}

func (obj *rsvpLspIpv4InterfaceP2PEgressIpv4Lsp) setDefault() {
	if obj.obj.RefreshInterval == nil {
		obj.SetRefreshInterval(30)
	}
	if obj.obj.TimeoutMultiplier == nil {
		obj.SetTimeoutMultiplier(3)
	}
	if obj.obj.ReservationStyle == nil {
		obj.SetReservationStyle(RsvpLspIpv4InterfaceP2PEgressIpv4LspReservationStyle.SHARED_EXPLICIT)

	}
	if obj.obj.EnableFixedLabel == nil {
		obj.SetEnableFixedLabel(false)
	}
	if obj.obj.FixedLabelValue == nil {
		obj.SetFixedLabelValue(0)
	}

}
