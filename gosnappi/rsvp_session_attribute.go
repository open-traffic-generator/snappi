package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpSessionAttribute *****
type rsvpSessionAttribute struct {
	validation
	obj                      *otg.RsvpSessionAttribute
	marshaller               marshalRsvpSessionAttribute
	unMarshaller             unMarshalRsvpSessionAttribute
	resourceAffinitiesHolder RsvpResourceAffinities
}

func NewRsvpSessionAttribute() RsvpSessionAttribute {
	obj := rsvpSessionAttribute{obj: &otg.RsvpSessionAttribute{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpSessionAttribute) msg() *otg.RsvpSessionAttribute {
	return obj.obj
}

func (obj *rsvpSessionAttribute) setMsg(msg *otg.RsvpSessionAttribute) RsvpSessionAttribute {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpSessionAttribute struct {
	obj *rsvpSessionAttribute
}

type marshalRsvpSessionAttribute interface {
	// ToProto marshals RsvpSessionAttribute to protobuf object *otg.RsvpSessionAttribute
	ToProto() (*otg.RsvpSessionAttribute, error)
	// ToPbText marshals RsvpSessionAttribute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpSessionAttribute to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpSessionAttribute to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpSessionAttribute struct {
	obj *rsvpSessionAttribute
}

type unMarshalRsvpSessionAttribute interface {
	// FromProto unmarshals RsvpSessionAttribute from protobuf object *otg.RsvpSessionAttribute
	FromProto(msg *otg.RsvpSessionAttribute) (RsvpSessionAttribute, error)
	// FromPbText unmarshals RsvpSessionAttribute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpSessionAttribute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpSessionAttribute from JSON text
	FromJson(value string) error
}

func (obj *rsvpSessionAttribute) Marshal() marshalRsvpSessionAttribute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpSessionAttribute{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpSessionAttribute) Unmarshal() unMarshalRsvpSessionAttribute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpSessionAttribute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpSessionAttribute) ToProto() (*otg.RsvpSessionAttribute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpSessionAttribute) FromProto(msg *otg.RsvpSessionAttribute) (RsvpSessionAttribute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpSessionAttribute) ToPbText() (string, error) {
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

func (m *unMarshalrsvpSessionAttribute) FromPbText(value string) error {
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

func (m *marshalrsvpSessionAttribute) ToYaml() (string, error) {
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

func (m *unMarshalrsvpSessionAttribute) FromYaml(value string) error {
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

func (m *marshalrsvpSessionAttribute) ToJson() (string, error) {
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

func (m *unMarshalrsvpSessionAttribute) FromJson(value string) error {
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

func (obj *rsvpSessionAttribute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpSessionAttribute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpSessionAttribute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpSessionAttribute) Clone() (RsvpSessionAttribute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpSessionAttribute()
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

func (obj *rsvpSessionAttribute) setNil() {
	obj.resourceAffinitiesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RsvpSessionAttribute is configuration for RSVP-TE SESSION_ATTRIBUTE object included in Path Messages as defined in RFC3209. The bandwidth_protection_desired and node_protection_desired flags are defined in RFC4090 (Fast Reroute).
type RsvpSessionAttribute interface {
	Validation
	// msg marshals RsvpSessionAttribute to protobuf object *otg.RsvpSessionAttribute
	// and doesn't set defaults
	msg() *otg.RsvpSessionAttribute
	// setMsg unmarshals RsvpSessionAttribute from protobuf object *otg.RsvpSessionAttribute
	// and doesn't set defaults
	setMsg(*otg.RsvpSessionAttribute) RsvpSessionAttribute
	// provides marshal interface
	Marshal() marshalRsvpSessionAttribute
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpSessionAttribute
	// validate validates RsvpSessionAttribute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpSessionAttribute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AutoGenerateSessionName returns bool, set in RsvpSessionAttribute.
	AutoGenerateSessionName() bool
	// SetAutoGenerateSessionName assigns bool provided by user to RsvpSessionAttribute
	SetAutoGenerateSessionName(value bool) RsvpSessionAttribute
	// HasAutoGenerateSessionName checks if AutoGenerateSessionName has been set in RsvpSessionAttribute
	HasAutoGenerateSessionName() bool
	// SessionName returns string, set in RsvpSessionAttribute.
	SessionName() string
	// SetSessionName assigns string provided by user to RsvpSessionAttribute
	SetSessionName(value string) RsvpSessionAttribute
	// HasSessionName checks if SessionName has been set in RsvpSessionAttribute
	HasSessionName() bool
	// SetupPriority returns uint32, set in RsvpSessionAttribute.
	SetupPriority() uint32
	// SetSetupPriority assigns uint32 provided by user to RsvpSessionAttribute
	SetSetupPriority(value uint32) RsvpSessionAttribute
	// HasSetupPriority checks if SetupPriority has been set in RsvpSessionAttribute
	HasSetupPriority() bool
	// HoldingPriority returns uint32, set in RsvpSessionAttribute.
	HoldingPriority() uint32
	// SetHoldingPriority assigns uint32 provided by user to RsvpSessionAttribute
	SetHoldingPriority(value uint32) RsvpSessionAttribute
	// HasHoldingPriority checks if HoldingPriority has been set in RsvpSessionAttribute
	HasHoldingPriority() bool
	// LocalProtectionDesired returns bool, set in RsvpSessionAttribute.
	LocalProtectionDesired() bool
	// SetLocalProtectionDesired assigns bool provided by user to RsvpSessionAttribute
	SetLocalProtectionDesired(value bool) RsvpSessionAttribute
	// HasLocalProtectionDesired checks if LocalProtectionDesired has been set in RsvpSessionAttribute
	HasLocalProtectionDesired() bool
	// LabelRecordingDesired returns bool, set in RsvpSessionAttribute.
	LabelRecordingDesired() bool
	// SetLabelRecordingDesired assigns bool provided by user to RsvpSessionAttribute
	SetLabelRecordingDesired(value bool) RsvpSessionAttribute
	// HasLabelRecordingDesired checks if LabelRecordingDesired has been set in RsvpSessionAttribute
	HasLabelRecordingDesired() bool
	// SeStyleDesired returns bool, set in RsvpSessionAttribute.
	SeStyleDesired() bool
	// SetSeStyleDesired assigns bool provided by user to RsvpSessionAttribute
	SetSeStyleDesired(value bool) RsvpSessionAttribute
	// HasSeStyleDesired checks if SeStyleDesired has been set in RsvpSessionAttribute
	HasSeStyleDesired() bool
	// BandwidthProtectionDesired returns bool, set in RsvpSessionAttribute.
	BandwidthProtectionDesired() bool
	// SetBandwidthProtectionDesired assigns bool provided by user to RsvpSessionAttribute
	SetBandwidthProtectionDesired(value bool) RsvpSessionAttribute
	// HasBandwidthProtectionDesired checks if BandwidthProtectionDesired has been set in RsvpSessionAttribute
	HasBandwidthProtectionDesired() bool
	// NodeProtectionDesired returns bool, set in RsvpSessionAttribute.
	NodeProtectionDesired() bool
	// SetNodeProtectionDesired assigns bool provided by user to RsvpSessionAttribute
	SetNodeProtectionDesired(value bool) RsvpSessionAttribute
	// HasNodeProtectionDesired checks if NodeProtectionDesired has been set in RsvpSessionAttribute
	HasNodeProtectionDesired() bool
	// ResourceAffinities returns RsvpResourceAffinities, set in RsvpSessionAttribute.
	// RsvpResourceAffinities is this is an optional object. If included, the extended SESSION_ATTRIBUTE object is sent in the Path message containing
	// the additional fields included in this object. This contains a set of three bitmaps using which further constraints can be
	// set on the path calculated for the LSP based on the Admin Group settings in the IGP (e.g ISIS or OSPF interface).
	ResourceAffinities() RsvpResourceAffinities
	// SetResourceAffinities assigns RsvpResourceAffinities provided by user to RsvpSessionAttribute.
	// RsvpResourceAffinities is this is an optional object. If included, the extended SESSION_ATTRIBUTE object is sent in the Path message containing
	// the additional fields included in this object. This contains a set of three bitmaps using which further constraints can be
	// set on the path calculated for the LSP based on the Admin Group settings in the IGP (e.g ISIS or OSPF interface).
	SetResourceAffinities(value RsvpResourceAffinities) RsvpSessionAttribute
	// HasResourceAffinities checks if ResourceAffinities has been set in RsvpSessionAttribute
	HasResourceAffinities() bool
	setNil()
}

// If this is enabled, an auto-generated Session Name is included in the SESSION_ATTRIBUTE object in the Path Message for this LSP.
// AutoGenerateSessionName returns a bool
func (obj *rsvpSessionAttribute) AutoGenerateSessionName() bool {

	return *obj.obj.AutoGenerateSessionName

}

// If this is enabled, an auto-generated Session Name is included in the SESSION_ATTRIBUTE object in the Path Message for this LSP.
// AutoGenerateSessionName returns a bool
func (obj *rsvpSessionAttribute) HasAutoGenerateSessionName() bool {
	return obj.obj.AutoGenerateSessionName != nil
}

// If this is enabled, an auto-generated Session Name is included in the SESSION_ATTRIBUTE object in the Path Message for this LSP.
// SetAutoGenerateSessionName sets the bool value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetAutoGenerateSessionName(value bool) RsvpSessionAttribute {

	obj.obj.AutoGenerateSessionName = &value
	return obj
}

// If auto_generate_session_name is set to 'false', then the value of this field is used to fill the Session Name field of the SESSION_ATTRIBUTE object in the Path Message for this LSP. It is suggested to include the Local IP, Remote IP, Tunnel ID and LSP ID in the auto-generated Session Name to ensure uniqueness of the name in the test. The maximum length of session name is 254 bytes.
// SessionName returns a string
func (obj *rsvpSessionAttribute) SessionName() string {

	return *obj.obj.SessionName

}

// If auto_generate_session_name is set to 'false', then the value of this field is used to fill the Session Name field of the SESSION_ATTRIBUTE object in the Path Message for this LSP. It is suggested to include the Local IP, Remote IP, Tunnel ID and LSP ID in the auto-generated Session Name to ensure uniqueness of the name in the test. The maximum length of session name is 254 bytes.
// SessionName returns a string
func (obj *rsvpSessionAttribute) HasSessionName() bool {
	return obj.obj.SessionName != nil
}

// If auto_generate_session_name is set to 'false', then the value of this field is used to fill the Session Name field of the SESSION_ATTRIBUTE object in the Path Message for this LSP. It is suggested to include the Local IP, Remote IP, Tunnel ID and LSP ID in the auto-generated Session Name to ensure uniqueness of the name in the test. The maximum length of session name is 254 bytes.
// SetSessionName sets the string value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetSessionName(value string) RsvpSessionAttribute {

	obj.obj.SessionName = &value
	return obj
}

// Specifies the value of the Setup Priority field. This controls whether the LSP should pre-empt existing  LSP setup with certain Holding Priority if resource limitation is encountered when setting up the LSP. (e.g. bandwidth availability). The value 0 is the highest priority while 7 is the lowest.
// SetupPriority returns a uint32
func (obj *rsvpSessionAttribute) SetupPriority() uint32 {

	return *obj.obj.SetupPriority

}

// Specifies the value of the Setup Priority field. This controls whether the LSP should pre-empt existing  LSP setup with certain Holding Priority if resource limitation is encountered when setting up the LSP. (e.g. bandwidth availability). The value 0 is the highest priority while 7 is the lowest.
// SetupPriority returns a uint32
func (obj *rsvpSessionAttribute) HasSetupPriority() bool {
	return obj.obj.SetupPriority != nil
}

// Specifies the value of the Setup Priority field. This controls whether the LSP should pre-empt existing  LSP setup with certain Holding Priority if resource limitation is encountered when setting up the LSP. (e.g. bandwidth availability). The value 0 is the highest priority while 7 is the lowest.
// SetSetupPriority sets the uint32 value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetSetupPriority(value uint32) RsvpSessionAttribute {

	obj.obj.SetupPriority = &value
	return obj
}

// Specifies the value of the Holding Priority field. This controls whether a new LSP being created with certain Setup Priority should pre-empt this LSP if resource limitation is encountered when setting up the LSP. (e.g. bandwidth availability). The value 0 is the highest priority while 7 is the lowest.
// HoldingPriority returns a uint32
func (obj *rsvpSessionAttribute) HoldingPriority() uint32 {

	return *obj.obj.HoldingPriority

}

// Specifies the value of the Holding Priority field. This controls whether a new LSP being created with certain Setup Priority should pre-empt this LSP if resource limitation is encountered when setting up the LSP. (e.g. bandwidth availability). The value 0 is the highest priority while 7 is the lowest.
// HoldingPriority returns a uint32
func (obj *rsvpSessionAttribute) HasHoldingPriority() bool {
	return obj.obj.HoldingPriority != nil
}

// Specifies the value of the Holding Priority field. This controls whether a new LSP being created with certain Setup Priority should pre-empt this LSP if resource limitation is encountered when setting up the LSP. (e.g. bandwidth availability). The value 0 is the highest priority while 7 is the lowest.
// SetHoldingPriority sets the uint32 value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetHoldingPriority(value uint32) RsvpSessionAttribute {

	obj.obj.HoldingPriority = &value
	return obj
}

// This flag permits transit routers to use a local repair mechanism which may result in violation of the explicit route object.  When a fault is detected on an adjacent downstream link or node, a transit router can reroute traffic for fast service restoration.
// LocalProtectionDesired returns a bool
func (obj *rsvpSessionAttribute) LocalProtectionDesired() bool {

	return *obj.obj.LocalProtectionDesired

}

// This flag permits transit routers to use a local repair mechanism which may result in violation of the explicit route object.  When a fault is detected on an adjacent downstream link or node, a transit router can reroute traffic for fast service restoration.
// LocalProtectionDesired returns a bool
func (obj *rsvpSessionAttribute) HasLocalProtectionDesired() bool {
	return obj.obj.LocalProtectionDesired != nil
}

// This flag permits transit routers to use a local repair mechanism which may result in violation of the explicit route object.  When a fault is detected on an adjacent downstream link or node, a transit router can reroute traffic for fast service restoration.
// SetLocalProtectionDesired sets the bool value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetLocalProtectionDesired(value bool) RsvpSessionAttribute {

	obj.obj.LocalProtectionDesired = &value
	return obj
}

// This flag indicates that label information should be included when doing a route record.
// LabelRecordingDesired returns a bool
func (obj *rsvpSessionAttribute) LabelRecordingDesired() bool {

	return *obj.obj.LabelRecordingDesired

}

// This flag indicates that label information should be included when doing a route record.
// LabelRecordingDesired returns a bool
func (obj *rsvpSessionAttribute) HasLabelRecordingDesired() bool {
	return obj.obj.LabelRecordingDesired != nil
}

// This flag indicates that label information should be included when doing a route record.
// SetLabelRecordingDesired sets the bool value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetLabelRecordingDesired(value bool) RsvpSessionAttribute {

	obj.obj.LabelRecordingDesired = &value
	return obj
}

// This flag indicates that the tunnel ingress node may choose to reroute this tunnel without tearing it down. A tunnel egress node SHOULD use the Shared Explicit(SE) Style when responding with a Resv message.
// SeStyleDesired returns a bool
func (obj *rsvpSessionAttribute) SeStyleDesired() bool {

	return *obj.obj.SeStyleDesired

}

// This flag indicates that the tunnel ingress node may choose to reroute this tunnel without tearing it down. A tunnel egress node SHOULD use the Shared Explicit(SE) Style when responding with a Resv message.
// SeStyleDesired returns a bool
func (obj *rsvpSessionAttribute) HasSeStyleDesired() bool {
	return obj.obj.SeStyleDesired != nil
}

// This flag indicates that the tunnel ingress node may choose to reroute this tunnel without tearing it down. A tunnel egress node SHOULD use the Shared Explicit(SE) Style when responding with a Resv message.
// SetSeStyleDesired sets the bool value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetSeStyleDesired(value bool) RsvpSessionAttribute {

	obj.obj.SeStyleDesired = &value
	return obj
}

// This flag in the SESSION_ATTRIBUTE object in the Path Message indicates to the PLRs along the protected LSP path that a  backup path with a bandwidth guarantee is desired.   This bandwidth has to be guaranteed for the protected LSP, if no FAST_REROUTE object is included in the PATH message. If a FAST_REROUTE object is present in the Path message, then the bandwidth specified therein is to be guaranteed.
// BandwidthProtectionDesired returns a bool
func (obj *rsvpSessionAttribute) BandwidthProtectionDesired() bool {

	return *obj.obj.BandwidthProtectionDesired

}

// This flag in the SESSION_ATTRIBUTE object in the Path Message indicates to the PLRs along the protected LSP path that a  backup path with a bandwidth guarantee is desired.   This bandwidth has to be guaranteed for the protected LSP, if no FAST_REROUTE object is included in the PATH message. If a FAST_REROUTE object is present in the Path message, then the bandwidth specified therein is to be guaranteed.
// BandwidthProtectionDesired returns a bool
func (obj *rsvpSessionAttribute) HasBandwidthProtectionDesired() bool {
	return obj.obj.BandwidthProtectionDesired != nil
}

// This flag in the SESSION_ATTRIBUTE object in the Path Message indicates to the PLRs along the protected LSP path that a  backup path with a bandwidth guarantee is desired.   This bandwidth has to be guaranteed for the protected LSP, if no FAST_REROUTE object is included in the PATH message. If a FAST_REROUTE object is present in the Path message, then the bandwidth specified therein is to be guaranteed.
// SetBandwidthProtectionDesired sets the bool value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetBandwidthProtectionDesired(value bool) RsvpSessionAttribute {

	obj.obj.BandwidthProtectionDesired = &value
	return obj
}

// This flag in the SESSION_ATTRIBUTE object in the Path Message indicates to the PLRs along a protected LSP path that it is desired to have a backup path that bypasses at least the next node of the protected LSP.
// NodeProtectionDesired returns a bool
func (obj *rsvpSessionAttribute) NodeProtectionDesired() bool {

	return *obj.obj.NodeProtectionDesired

}

// This flag in the SESSION_ATTRIBUTE object in the Path Message indicates to the PLRs along a protected LSP path that it is desired to have a backup path that bypasses at least the next node of the protected LSP.
// NodeProtectionDesired returns a bool
func (obj *rsvpSessionAttribute) HasNodeProtectionDesired() bool {
	return obj.obj.NodeProtectionDesired != nil
}

// This flag in the SESSION_ATTRIBUTE object in the Path Message indicates to the PLRs along a protected LSP path that it is desired to have a backup path that bypasses at least the next node of the protected LSP.
// SetNodeProtectionDesired sets the bool value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetNodeProtectionDesired(value bool) RsvpSessionAttribute {

	obj.obj.NodeProtectionDesired = &value
	return obj
}

// This is an optional object. If included the extended SESSION_ATTRIBUTE object is sent in the Path message containing
// the additional fields included in this object. This contains a set of three bitmaps using which further constraints can be
// set on the path calculated for the LSP based on the Admin Group settings in the IGP (e.g ISIS or OSPF interface).
// ResourceAffinities returns a RsvpResourceAffinities
func (obj *rsvpSessionAttribute) ResourceAffinities() RsvpResourceAffinities {
	if obj.obj.ResourceAffinities == nil {
		obj.obj.ResourceAffinities = NewRsvpResourceAffinities().msg()
	}
	if obj.resourceAffinitiesHolder == nil {
		obj.resourceAffinitiesHolder = &rsvpResourceAffinities{obj: obj.obj.ResourceAffinities}
	}
	return obj.resourceAffinitiesHolder
}

// This is an optional object. If included the extended SESSION_ATTRIBUTE object is sent in the Path message containing
// the additional fields included in this object. This contains a set of three bitmaps using which further constraints can be
// set on the path calculated for the LSP based on the Admin Group settings in the IGP (e.g ISIS or OSPF interface).
// ResourceAffinities returns a RsvpResourceAffinities
func (obj *rsvpSessionAttribute) HasResourceAffinities() bool {
	return obj.obj.ResourceAffinities != nil
}

// This is an optional object. If included the extended SESSION_ATTRIBUTE object is sent in the Path message containing
// the additional fields included in this object. This contains a set of three bitmaps using which further constraints can be
// set on the path calculated for the LSP based on the Admin Group settings in the IGP (e.g ISIS or OSPF interface).
// SetResourceAffinities sets the RsvpResourceAffinities value in the RsvpSessionAttribute object
func (obj *rsvpSessionAttribute) SetResourceAffinities(value RsvpResourceAffinities) RsvpSessionAttribute {

	obj.resourceAffinitiesHolder = nil
	obj.obj.ResourceAffinities = value.msg()

	return obj
}

func (obj *rsvpSessionAttribute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SessionName != nil {

		if len(*obj.obj.SessionName) < 0 || len(*obj.obj.SessionName) > 254 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of RsvpSessionAttribute.SessionName <= 254 but Got %d",
					len(*obj.obj.SessionName)))
		}

	}

	if obj.obj.SetupPriority != nil {

		if *obj.obj.SetupPriority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpSessionAttribute.SetupPriority <= 7 but Got %d", *obj.obj.SetupPriority))
		}

	}

	if obj.obj.HoldingPriority != nil {

		if *obj.obj.HoldingPriority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpSessionAttribute.HoldingPriority <= 7 but Got %d", *obj.obj.HoldingPriority))
		}

	}

	if obj.obj.ResourceAffinities != nil {

		obj.ResourceAffinities().validateObj(vObj, set_default)
	}

}

func (obj *rsvpSessionAttribute) setDefault() {
	if obj.obj.AutoGenerateSessionName == nil {
		obj.SetAutoGenerateSessionName(true)
	}
	if obj.obj.SetupPriority == nil {
		obj.SetSetupPriority(7)
	}
	if obj.obj.HoldingPriority == nil {
		obj.SetHoldingPriority(7)
	}
	if obj.obj.LocalProtectionDesired == nil {
		obj.SetLocalProtectionDesired(false)
	}
	if obj.obj.LabelRecordingDesired == nil {
		obj.SetLabelRecordingDesired(false)
	}
	if obj.obj.SeStyleDesired == nil {
		obj.SetSeStyleDesired(false)
	}
	if obj.obj.BandwidthProtectionDesired == nil {
		obj.SetBandwidthProtectionDesired(false)
	}
	if obj.obj.NodeProtectionDesired == nil {
		obj.SetNodeProtectionDesired(false)
	}

}
