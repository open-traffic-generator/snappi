package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpLspIpv4InterfaceP2PIngressIpv4Lsp *****
type rsvpLspIpv4InterfaceP2PIngressIpv4Lsp struct {
	validation
	obj                    *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	marshaller             marshalRsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	unMarshaller           unMarshalRsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	sessionAttributeHolder RsvpSessionAttribute
	tspecHolder            RsvpTspec
	fastRerouteHolder      RsvpFastReroute
	eroHolder              RsvpEro
}

func NewRsvpLspIpv4InterfaceP2PIngressIpv4Lsp() RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {
	obj := rsvpLspIpv4InterfaceP2PIngressIpv4Lsp{obj: &otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) msg() *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {
	return obj.obj
}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) setMsg(msg *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp struct {
	obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp
}

type marshalRsvpLspIpv4InterfaceP2PIngressIpv4Lsp interface {
	// ToProto marshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp to protobuf object *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	ToProto() (*otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp, error)
	// ToPbText marshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp struct {
	obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp
}

type unMarshalRsvpLspIpv4InterfaceP2PIngressIpv4Lsp interface {
	// FromProto unmarshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp from protobuf object *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	FromProto(msg *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) (RsvpLspIpv4InterfaceP2PIngressIpv4Lsp, error)
	// FromPbText unmarshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp from JSON text
	FromJson(value string) error
}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) Marshal() marshalRsvpLspIpv4InterfaceP2PIngressIpv4Lsp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) Unmarshal() unMarshalRsvpLspIpv4InterfaceP2PIngressIpv4Lsp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp) ToProto() (*otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp) FromProto(msg *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) (RsvpLspIpv4InterfaceP2PIngressIpv4Lsp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp) ToPbText() (string, error) {
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

func (m *unMarshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp) FromPbText(value string) error {
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

func (m *marshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp) ToYaml() (string, error) {
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

func (m *unMarshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp) FromYaml(value string) error {
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

func (m *marshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp) ToJson() (string, error) {
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

func (m *unMarshalrsvpLspIpv4InterfaceP2PIngressIpv4Lsp) FromJson(value string) error {
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

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) Clone() (RsvpLspIpv4InterfaceP2PIngressIpv4Lsp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpLspIpv4InterfaceP2PIngressIpv4Lsp()
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

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) setNil() {
	obj.sessionAttributeHolder = nil
	obj.tspecHolder = nil
	obj.fastRerouteHolder = nil
	obj.eroHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RsvpLspIpv4InterfaceP2PIngressIpv4Lsp is configuration for an RSVP Ingress point-to-point LSP.
type RsvpLspIpv4InterfaceP2PIngressIpv4Lsp interface {
	Validation
	// msg marshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp to protobuf object *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// and doesn't set defaults
	msg() *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// setMsg unmarshals RsvpLspIpv4InterfaceP2PIngressIpv4Lsp from protobuf object *otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// and doesn't set defaults
	setMsg(*otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// provides marshal interface
	Marshal() marshalRsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// validate validates RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpLspIpv4InterfaceP2PIngressIpv4Lsp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	Name() string
	// SetName assigns string provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	SetName(value string) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// RemoteAddress returns string, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	RemoteAddress() string
	// SetRemoteAddress assigns string provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	SetRemoteAddress(value string) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// TunnelId returns uint32, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	TunnelId() uint32
	// SetTunnelId assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	SetTunnelId(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasTunnelId checks if TunnelId has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasTunnelId() bool
	// LspId returns uint32, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	LspId() uint32
	// SetLspId assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	SetLspId(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasLspId checks if LspId has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasLspId() bool
	// RefreshInterval returns uint32, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	RefreshInterval() uint32
	// SetRefreshInterval assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	SetRefreshInterval(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasRefreshInterval checks if RefreshInterval has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasRefreshInterval() bool
	// TimeoutMultiplier returns uint32, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	TimeoutMultiplier() uint32
	// SetTimeoutMultiplier assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	SetTimeoutMultiplier(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasTimeoutMultiplier checks if TimeoutMultiplier has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasTimeoutMultiplier() bool
	// BackupLspId returns uint32, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	BackupLspId() uint32
	// SetBackupLspId assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	SetBackupLspId(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasBackupLspId checks if BackupLspId has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasBackupLspId() bool
	// LspSwitchoverDelay returns uint32, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	LspSwitchoverDelay() uint32
	// SetLspSwitchoverDelay assigns uint32 provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	SetLspSwitchoverDelay(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasLspSwitchoverDelay checks if LspSwitchoverDelay has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasLspSwitchoverDelay() bool
	// SessionAttribute returns RsvpSessionAttribute, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	// RsvpSessionAttribute is configuration for RSVP-TE SESSION_ATTRIBUTE object included in Path Messages as defined in RFC3209. The bandwidth_protection_desired and node_protection_desired flags are defined in RFC4090 (Fast Reroute).
	SessionAttribute() RsvpSessionAttribute
	// SetSessionAttribute assigns RsvpSessionAttribute provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	// RsvpSessionAttribute is configuration for RSVP-TE SESSION_ATTRIBUTE object included in Path Messages as defined in RFC3209. The bandwidth_protection_desired and node_protection_desired flags are defined in RFC4090 (Fast Reroute).
	SetSessionAttribute(value RsvpSessionAttribute) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasSessionAttribute checks if SessionAttribute has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasSessionAttribute() bool
	// Tspec returns RsvpTspec, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	// RsvpTspec is configuration for RSVP-TE TSPEC object included in Path Messages. The usage of these parameters is defined in RFC2215.
	Tspec() RsvpTspec
	// SetTspec assigns RsvpTspec provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	// RsvpTspec is configuration for RSVP-TE TSPEC object included in Path Messages. The usage of these parameters is defined in RFC2215.
	SetTspec(value RsvpTspec) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasTspec checks if Tspec has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasTspec() bool
	// FastReroute returns RsvpFastReroute, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	// RsvpFastReroute is configuration for the optional RSVP-TE FAST_REROUTE object included in Path Messages as defined in RFC4090.
	FastReroute() RsvpFastReroute
	// SetFastReroute assigns RsvpFastReroute provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	// RsvpFastReroute is configuration for the optional RSVP-TE FAST_REROUTE object included in Path Messages as defined in RFC4090.
	SetFastReroute(value RsvpFastReroute) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasFastReroute checks if FastReroute has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasFastReroute() bool
	// Ero returns RsvpEro, set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	// RsvpEro is configuration for the optional RSVP-TE explicit route object(ERO) object included in Path Messages.
	Ero() RsvpEro
	// SetEro assigns RsvpEro provided by user to RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.
	// RsvpEro is configuration for the optional RSVP-TE explicit route object(ERO) object included in Path Messages.
	SetEro(value RsvpEro) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	// HasEro checks if Ero has been set in RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	HasEro() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetName(value string) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.obj.Name = &value
	return obj
}

// IPv4 address of the remote endpoint of the LSP.
// RemoteAddress returns a string
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RemoteAddress() string {

	return *obj.obj.RemoteAddress

}

// IPv4 address of the remote endpoint of the LSP.
// SetRemoteAddress sets the string value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetRemoteAddress(value string) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.obj.RemoteAddress = &value
	return obj
}

// The Tunnel ID of the RSVP LSP. Carried in the SESSION object in Path Messages.
// TunnelId returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) TunnelId() uint32 {

	return *obj.obj.TunnelId

}

// The Tunnel ID of the RSVP LSP. Carried in the SESSION object in Path Messages.
// TunnelId returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasTunnelId() bool {
	return obj.obj.TunnelId != nil
}

// The Tunnel ID of the RSVP LSP. Carried in the SESSION object in Path Messages.
// SetTunnelId sets the uint32 value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetTunnelId(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.obj.TunnelId = &value
	return obj
}

// The LSP ID of the RSVP LSP. Carried in the SENDER_TEMPLATE object in Path Messages.
// LspId returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) LspId() uint32 {

	return *obj.obj.LspId

}

// The LSP ID of the RSVP LSP. Carried in the SENDER_TEMPLATE object in Path Messages.
// LspId returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasLspId() bool {
	return obj.obj.LspId != nil
}

// The LSP ID of the RSVP LSP. Carried in the SENDER_TEMPLATE object in Path Messages.
// SetLspId sets the uint32 value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetLspId(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.obj.LspId = &value
	return obj
}

// The time in seconds between successive transmissions of PATH Refreshes.  The actual refresh interval is jittered by upto 50%. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// RefreshInterval returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RefreshInterval() uint32 {

	return *obj.obj.RefreshInterval

}

// The time in seconds between successive transmissions of PATH Refreshes.  The actual refresh interval is jittered by upto 50%. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// RefreshInterval returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasRefreshInterval() bool {
	return obj.obj.RefreshInterval != nil
}

// The time in seconds between successive transmissions of PATH Refreshes.  The actual refresh interval is jittered by upto 50%. There is no specification specified maximum value. For clarity, setting the maximum to 1 hour.
// SetRefreshInterval sets the uint32 value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetRefreshInterval(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.obj.RefreshInterval = &value
	return obj
}

// The number of missed RESV refreshes after which a recieving node should consider the LSP state to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// TimeoutMultiplier returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) TimeoutMultiplier() uint32 {

	return *obj.obj.TimeoutMultiplier

}

// The number of missed RESV refreshes after which a recieving node should consider the LSP state to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// TimeoutMultiplier returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasTimeoutMultiplier() bool {
	return obj.obj.TimeoutMultiplier != nil
}

// The number of missed RESV refreshes after which a recieving node should consider the LSP state to have timed out. There is no specification specified maximum value. Setting the maximum allowed value to 10.
// SetTimeoutMultiplier sets the uint32 value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetTimeoutMultiplier(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.obj.TimeoutMultiplier = &value
	return obj
}

// The LSP id that will be used when creating a Make-Before-Break LSP when the active LSP is using lsp_id. If the active LSP on which Make-Before-Break is being done is using the backup_lsp_id, the new LSP created will toggle to  use the lsp_id instead.
// BackupLspId returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) BackupLspId() uint32 {

	return *obj.obj.BackupLspId

}

// The LSP id that will be used when creating a Make-Before-Break LSP when the active LSP is using lsp_id. If the active LSP on which Make-Before-Break is being done is using the backup_lsp_id, the new LSP created will toggle to  use the lsp_id instead.
// BackupLspId returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasBackupLspId() bool {
	return obj.obj.BackupLspId != nil
}

// The LSP id that will be used when creating a Make-Before-Break LSP when the active LSP is using lsp_id. If the active LSP on which Make-Before-Break is being done is using the backup_lsp_id, the new LSP created will toggle to  use the lsp_id instead.
// SetBackupLspId sets the uint32 value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetBackupLspId(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.obj.BackupLspId = &value
	return obj
}

// The amount of delay in milliseconds that an implementation should wait for before switching traffic to the new LSP created after  a Make-Before-Break is done on an LSP. The default value is 0 which means to switch immediately. An implementation should support a minimum delay value of at least 50ms . There is no specification specified maximum value. Setting maximum allowed value to 1 minute.  If a delay value is supplied which is lesser than the minimum delay value supported, a warning should be provided indicating that the minimum value of LSP switchover delay is automatically increased to the supported minimum value. This warning should be included in the list of warnings in the 'Response.Warning' attribute sent in the SetConfig 'Success' Response.
// LspSwitchoverDelay returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) LspSwitchoverDelay() uint32 {

	return *obj.obj.LspSwitchoverDelay

}

// The amount of delay in milliseconds that an implementation should wait for before switching traffic to the new LSP created after  a Make-Before-Break is done on an LSP. The default value is 0 which means to switch immediately. An implementation should support a minimum delay value of at least 50ms . There is no specification specified maximum value. Setting maximum allowed value to 1 minute.  If a delay value is supplied which is lesser than the minimum delay value supported, a warning should be provided indicating that the minimum value of LSP switchover delay is automatically increased to the supported minimum value. This warning should be included in the list of warnings in the 'Response.Warning' attribute sent in the SetConfig 'Success' Response.
// LspSwitchoverDelay returns a uint32
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasLspSwitchoverDelay() bool {
	return obj.obj.LspSwitchoverDelay != nil
}

// The amount of delay in milliseconds that an implementation should wait for before switching traffic to the new LSP created after  a Make-Before-Break is done on an LSP. The default value is 0 which means to switch immediately. An implementation should support a minimum delay value of at least 50ms . There is no specification specified maximum value. Setting maximum allowed value to 1 minute.  If a delay value is supplied which is lesser than the minimum delay value supported, a warning should be provided indicating that the minimum value of LSP switchover delay is automatically increased to the supported minimum value. This warning should be included in the list of warnings in the 'Response.Warning' attribute sent in the SetConfig 'Success' Response.
// SetLspSwitchoverDelay sets the uint32 value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetLspSwitchoverDelay(value uint32) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.obj.LspSwitchoverDelay = &value
	return obj
}

// This contains the values of the fields to be included in the SESSION_ATTRIBUTE object in the Path Message sent for the LSP.
// SessionAttribute returns a RsvpSessionAttribute
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SessionAttribute() RsvpSessionAttribute {
	if obj.obj.SessionAttribute == nil {
		obj.obj.SessionAttribute = NewRsvpSessionAttribute().msg()
	}
	if obj.sessionAttributeHolder == nil {
		obj.sessionAttributeHolder = &rsvpSessionAttribute{obj: obj.obj.SessionAttribute}
	}
	return obj.sessionAttributeHolder
}

// This contains the values of the fields to be included in the SESSION_ATTRIBUTE object in the Path Message sent for the LSP.
// SessionAttribute returns a RsvpSessionAttribute
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasSessionAttribute() bool {
	return obj.obj.SessionAttribute != nil
}

// This contains the values of the fields to be included in the SESSION_ATTRIBUTE object in the Path Message sent for the LSP.
// SetSessionAttribute sets the RsvpSessionAttribute value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetSessionAttribute(value RsvpSessionAttribute) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.sessionAttributeHolder = nil
	obj.obj.SessionAttribute = value.msg()

	return obj
}

// This contains the values of the fields to be included in the TSPEC object in the Path Message sent for the LSP.
// Tspec returns a RsvpTspec
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) Tspec() RsvpTspec {
	if obj.obj.Tspec == nil {
		obj.obj.Tspec = NewRsvpTspec().msg()
	}
	if obj.tspecHolder == nil {
		obj.tspecHolder = &rsvpTspec{obj: obj.obj.Tspec}
	}
	return obj.tspecHolder
}

// This contains the values of the fields to be included in the TSPEC object in the Path Message sent for the LSP.
// Tspec returns a RsvpTspec
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasTspec() bool {
	return obj.obj.Tspec != nil
}

// This contains the values of the fields to be included in the TSPEC object in the Path Message sent for the LSP.
// SetTspec sets the RsvpTspec value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetTspec(value RsvpTspec) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.tspecHolder = nil
	obj.obj.Tspec = value.msg()

	return obj
}

// This contains the values of the fields to be included in the FAST_REROUTE object in the Path Message sent for the LSP.
// This is an optional object . If this attribute is not included , the FAST_REROUTE object will not be included.
// FastReroute returns a RsvpFastReroute
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) FastReroute() RsvpFastReroute {
	if obj.obj.FastReroute == nil {
		obj.obj.FastReroute = NewRsvpFastReroute().msg()
	}
	if obj.fastRerouteHolder == nil {
		obj.fastRerouteHolder = &rsvpFastReroute{obj: obj.obj.FastReroute}
	}
	return obj.fastRerouteHolder
}

// This contains the values of the fields to be included in the FAST_REROUTE object in the Path Message sent for the LSP.
// This is an optional object . If this attribute is not included , the FAST_REROUTE object will not be included.
// FastReroute returns a RsvpFastReroute
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasFastReroute() bool {
	return obj.obj.FastReroute != nil
}

// This contains the values of the fields to be included in the FAST_REROUTE object in the Path Message sent for the LSP.
// This is an optional object . If this attribute is not included , the FAST_REROUTE object will not be included.
// SetFastReroute sets the RsvpFastReroute value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetFastReroute(value RsvpFastReroute) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.fastRerouteHolder = nil
	obj.obj.FastReroute = value.msg()

	return obj
}

// This contains the values of the fields to be included in the ERO object in the Path Message sent for the LSP.
// This is an optional object . If this attribute is not included , the ERO object will not be included.
// Ero returns a RsvpEro
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) Ero() RsvpEro {
	if obj.obj.Ero == nil {
		obj.obj.Ero = NewRsvpEro().msg()
	}
	if obj.eroHolder == nil {
		obj.eroHolder = &rsvpEro{obj: obj.obj.Ero}
	}
	return obj.eroHolder
}

// This contains the values of the fields to be included in the ERO object in the Path Message sent for the LSP.
// This is an optional object . If this attribute is not included , the ERO object will not be included.
// Ero returns a RsvpEro
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) HasEro() bool {
	return obj.obj.Ero != nil
}

// This contains the values of the fields to be included in the ERO object in the Path Message sent for the LSP.
// This is an optional object . If this attribute is not included , the ERO object will not be included.
// SetEro sets the RsvpEro value in the RsvpLspIpv4InterfaceP2PIngressIpv4Lsp object
func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) SetEro(value RsvpEro) RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {

	obj.eroHolder = nil
	obj.obj.Ero = value.msg()

	return obj
}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface RsvpLspIpv4InterfaceP2PIngressIpv4Lsp")
	}

	// RemoteAddress is required
	if obj.obj.RemoteAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RemoteAddress is required field on interface RsvpLspIpv4InterfaceP2PIngressIpv4Lsp")
	}
	if obj.obj.RemoteAddress != nil {

		err := obj.validateIpv4(obj.RemoteAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.RemoteAddress"))
		}

	}

	if obj.obj.TunnelId != nil {

		if *obj.obj.TunnelId < 1 || *obj.obj.TunnelId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.TunnelId <= 65535 but Got %d", *obj.obj.TunnelId))
		}

	}

	if obj.obj.LspId != nil {

		if *obj.obj.LspId < 1 || *obj.obj.LspId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.LspId <= 65535 but Got %d", *obj.obj.LspId))
		}

	}

	if obj.obj.RefreshInterval != nil {

		if *obj.obj.RefreshInterval < 1 || *obj.obj.RefreshInterval > 3600 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.RefreshInterval <= 3600 but Got %d", *obj.obj.RefreshInterval))
		}

	}

	if obj.obj.TimeoutMultiplier != nil {

		if *obj.obj.TimeoutMultiplier < 1 || *obj.obj.TimeoutMultiplier > 10 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.TimeoutMultiplier <= 10 but Got %d", *obj.obj.TimeoutMultiplier))
		}

	}

	if obj.obj.BackupLspId != nil {

		if *obj.obj.BackupLspId < 1 || *obj.obj.BackupLspId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.BackupLspId <= 65535 but Got %d", *obj.obj.BackupLspId))
		}

	}

	if obj.obj.LspSwitchoverDelay != nil {

		if *obj.obj.LspSwitchoverDelay > 60000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpLspIpv4InterfaceP2PIngressIpv4Lsp.LspSwitchoverDelay <= 60000 but Got %d", *obj.obj.LspSwitchoverDelay))
		}

	}

	if obj.obj.SessionAttribute != nil {

		obj.SessionAttribute().validateObj(vObj, set_default)
	}

	if obj.obj.Tspec != nil {

		obj.Tspec().validateObj(vObj, set_default)
	}

	if obj.obj.FastReroute != nil {

		obj.FastReroute().validateObj(vObj, set_default)
	}

	if obj.obj.Ero != nil {

		obj.Ero().validateObj(vObj, set_default)
	}

}

func (obj *rsvpLspIpv4InterfaceP2PIngressIpv4Lsp) setDefault() {
	if obj.obj.TunnelId == nil {
		obj.SetTunnelId(1)
	}
	if obj.obj.LspId == nil {
		obj.SetLspId(1)
	}
	if obj.obj.RefreshInterval == nil {
		obj.SetRefreshInterval(30)
	}
	if obj.obj.TimeoutMultiplier == nil {
		obj.SetTimeoutMultiplier(3)
	}
	if obj.obj.BackupLspId == nil {
		obj.SetBackupLspId(2)
	}
	if obj.obj.LspSwitchoverDelay == nil {
		obj.SetLspSwitchoverDelay(0)
	}

}
