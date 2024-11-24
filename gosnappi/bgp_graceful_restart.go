package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpGracefulRestart *****
type bgpGracefulRestart struct {
	validation
	obj          *otg.BgpGracefulRestart
	marshaller   marshalBgpGracefulRestart
	unMarshaller unMarshalBgpGracefulRestart
}

func NewBgpGracefulRestart() BgpGracefulRestart {
	obj := bgpGracefulRestart{obj: &otg.BgpGracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpGracefulRestart) msg() *otg.BgpGracefulRestart {
	return obj.obj
}

func (obj *bgpGracefulRestart) setMsg(msg *otg.BgpGracefulRestart) BgpGracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpGracefulRestart struct {
	obj *bgpGracefulRestart
}

type marshalBgpGracefulRestart interface {
	// ToProto marshals BgpGracefulRestart to protobuf object *otg.BgpGracefulRestart
	ToProto() (*otg.BgpGracefulRestart, error)
	// ToPbText marshals BgpGracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpGracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpGracefulRestart to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpGracefulRestart to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpGracefulRestart struct {
	obj *bgpGracefulRestart
}

type unMarshalBgpGracefulRestart interface {
	// FromProto unmarshals BgpGracefulRestart from protobuf object *otg.BgpGracefulRestart
	FromProto(msg *otg.BgpGracefulRestart) (BgpGracefulRestart, error)
	// FromPbText unmarshals BgpGracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpGracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpGracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *bgpGracefulRestart) Marshal() marshalBgpGracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpGracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpGracefulRestart) Unmarshal() unMarshalBgpGracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpGracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpGracefulRestart) ToProto() (*otg.BgpGracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpGracefulRestart) FromProto(msg *otg.BgpGracefulRestart) (BgpGracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpGracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalbgpGracefulRestart) FromPbText(value string) error {
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

func (m *marshalbgpGracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalbgpGracefulRestart) FromYaml(value string) error {
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

func (m *marshalbgpGracefulRestart) ToJsonRaw() (string, error) {
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

func (m *marshalbgpGracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalbgpGracefulRestart) FromJson(value string) error {
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

func (obj *bgpGracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpGracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpGracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpGracefulRestart) Clone() (BgpGracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpGracefulRestart()
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

// BgpGracefulRestart is the Graceful Restart Capability (RFC 4724) is a BGP capability that can be used by a BGP speaker to indicate its ability to preserve its forwarding state during BGP restart. The Graceful Restart (GR) capability is advertised in OPEN messages sent between BGP peers. After a BGP session has been established, and the initial routing update has been completed,  an End-of-RIB (Routing Information Base) marker is sent in an UPDATE message to convey information  about routing convergence.
type BgpGracefulRestart interface {
	Validation
	// msg marshals BgpGracefulRestart to protobuf object *otg.BgpGracefulRestart
	// and doesn't set defaults
	msg() *otg.BgpGracefulRestart
	// setMsg unmarshals BgpGracefulRestart from protobuf object *otg.BgpGracefulRestart
	// and doesn't set defaults
	setMsg(*otg.BgpGracefulRestart) BgpGracefulRestart
	// provides marshal interface
	Marshal() marshalBgpGracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalBgpGracefulRestart
	// validate validates BgpGracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpGracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnableGr returns bool, set in BgpGracefulRestart.
	EnableGr() bool
	// SetEnableGr assigns bool provided by user to BgpGracefulRestart
	SetEnableGr(value bool) BgpGracefulRestart
	// HasEnableGr checks if EnableGr has been set in BgpGracefulRestart
	HasEnableGr() bool
	// RestartTime returns uint32, set in BgpGracefulRestart.
	RestartTime() uint32
	// SetRestartTime assigns uint32 provided by user to BgpGracefulRestart
	SetRestartTime(value uint32) BgpGracefulRestart
	// HasRestartTime checks if RestartTime has been set in BgpGracefulRestart
	HasRestartTime() bool
	// EnableLlgr returns bool, set in BgpGracefulRestart.
	EnableLlgr() bool
	// SetEnableLlgr assigns bool provided by user to BgpGracefulRestart
	SetEnableLlgr(value bool) BgpGracefulRestart
	// HasEnableLlgr checks if EnableLlgr has been set in BgpGracefulRestart
	HasEnableLlgr() bool
	// StaleTime returns uint32, set in BgpGracefulRestart.
	StaleTime() uint32
	// SetStaleTime assigns uint32 provided by user to BgpGracefulRestart
	SetStaleTime(value uint32) BgpGracefulRestart
	// HasStaleTime checks if StaleTime has been set in BgpGracefulRestart
	HasStaleTime() bool
	// EnableNotification returns bool, set in BgpGracefulRestart.
	EnableNotification() bool
	// SetEnableNotification assigns bool provided by user to BgpGracefulRestart
	SetEnableNotification(value bool) BgpGracefulRestart
	// HasEnableNotification checks if EnableNotification has been set in BgpGracefulRestart
	HasEnableNotification() bool
}

// If enabled, Graceful Restart capability is advertised in BGP OPEN messages.
// EnableGr returns a bool
func (obj *bgpGracefulRestart) EnableGr() bool {

	return *obj.obj.EnableGr

}

// If enabled, Graceful Restart capability is advertised in BGP OPEN messages.
// EnableGr returns a bool
func (obj *bgpGracefulRestart) HasEnableGr() bool {
	return obj.obj.EnableGr != nil
}

// If enabled, Graceful Restart capability is advertised in BGP OPEN messages.
// SetEnableGr sets the bool value in the BgpGracefulRestart object
func (obj *bgpGracefulRestart) SetEnableGr(value bool) BgpGracefulRestart {

	obj.obj.EnableGr = &value
	return obj
}

// This is the estimated duration (in seconds) it will take for the BGP session to be re-established after a restart.  This can be used to speed up routing convergence by its peer in case the BGP speaker does not come back after a restart.
// RestartTime returns a uint32
func (obj *bgpGracefulRestart) RestartTime() uint32 {

	return *obj.obj.RestartTime

}

// This is the estimated duration (in seconds) it will take for the BGP session to be re-established after a restart.  This can be used to speed up routing convergence by its peer in case the BGP speaker does not come back after a restart.
// RestartTime returns a uint32
func (obj *bgpGracefulRestart) HasRestartTime() bool {
	return obj.obj.RestartTime != nil
}

// This is the estimated duration (in seconds) it will take for the BGP session to be re-established after a restart.  This can be used to speed up routing convergence by its peer in case the BGP speaker does not come back after a restart.
// SetRestartTime sets the uint32 value in the BgpGracefulRestart object
func (obj *bgpGracefulRestart) SetRestartTime(value uint32) BgpGracefulRestart {

	obj.obj.RestartTime = &value
	return obj
}

// If enabled, the "Long-lived Graceful Restart Capability", or "LLGR Capability"
// will be advertised.
// This capability MUST be advertised in conjunction with the Graceful Restart
// capability.
// EnableLlgr returns a bool
func (obj *bgpGracefulRestart) EnableLlgr() bool {

	return *obj.obj.EnableLlgr

}

// If enabled, the "Long-lived Graceful Restart Capability", or "LLGR Capability"
// will be advertised.
// This capability MUST be advertised in conjunction with the Graceful Restart
// capability.
// EnableLlgr returns a bool
func (obj *bgpGracefulRestart) HasEnableLlgr() bool {
	return obj.obj.EnableLlgr != nil
}

// If enabled, the "Long-lived Graceful Restart Capability", or "LLGR Capability"
// will be advertised.
// This capability MUST be advertised in conjunction with the Graceful Restart
// capability.
// SetEnableLlgr sets the bool value in the BgpGracefulRestart object
func (obj *bgpGracefulRestart) SetEnableLlgr(value bool) BgpGracefulRestart {

	obj.obj.EnableLlgr = &value
	return obj
}

// Duration (in seconds) specifying how long stale information (for the AFI/SAFI)
// may be retained. This is a three byte field and is applicable
// only if 'enable_llgr' is set to 'true'.
// StaleTime returns a uint32
func (obj *bgpGracefulRestart) StaleTime() uint32 {

	return *obj.obj.StaleTime

}

// Duration (in seconds) specifying how long stale information (for the AFI/SAFI)
// may be retained. This is a three byte field and is applicable
// only if 'enable_llgr' is set to 'true'.
// StaleTime returns a uint32
func (obj *bgpGracefulRestart) HasStaleTime() bool {
	return obj.obj.StaleTime != nil
}

// Duration (in seconds) specifying how long stale information (for the AFI/SAFI)
// may be retained. This is a three byte field and is applicable
// only if 'enable_llgr' is set to 'true'.
// SetStaleTime sets the uint32 value in the BgpGracefulRestart object
func (obj *bgpGracefulRestart) SetStaleTime(value uint32) BgpGracefulRestart {

	obj.obj.StaleTime = &value
	return obj
}

// If enabled, the N flag will be set in the Graceful Restart capability in the Open message.
// If both peers in a BGP connection has this enabled, Graceful Restart procedures are performed
// even if the peer goes down due to sending of a Notification Message as per RFC8538.
// EnableNotification returns a bool
func (obj *bgpGracefulRestart) EnableNotification() bool {

	return *obj.obj.EnableNotification

}

// If enabled, the N flag will be set in the Graceful Restart capability in the Open message.
// If both peers in a BGP connection has this enabled, Graceful Restart procedures are performed
// even if the peer goes down due to sending of a Notification Message as per RFC8538.
// EnableNotification returns a bool
func (obj *bgpGracefulRestart) HasEnableNotification() bool {
	return obj.obj.EnableNotification != nil
}

// If enabled, the N flag will be set in the Graceful Restart capability in the Open message.
// If both peers in a BGP connection has this enabled, Graceful Restart procedures are performed
// even if the peer goes down due to sending of a Notification Message as per RFC8538.
// SetEnableNotification sets the bool value in the BgpGracefulRestart object
func (obj *bgpGracefulRestart) SetEnableNotification(value bool) BgpGracefulRestart {

	obj.obj.EnableNotification = &value
	return obj
}

func (obj *bgpGracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartTime != nil {

		if *obj.obj.RestartTime > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpGracefulRestart.RestartTime <= 4096 but Got %d", *obj.obj.RestartTime))
		}

	}

	if obj.obj.StaleTime != nil {

		if *obj.obj.StaleTime > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpGracefulRestart.StaleTime <= 16777215 but Got %d", *obj.obj.StaleTime))
		}

	}

}

func (obj *bgpGracefulRestart) setDefault() {
	if obj.obj.EnableGr == nil {
		obj.SetEnableGr(false)
	}
	if obj.obj.RestartTime == nil {
		obj.SetRestartTime(45)
	}
	if obj.obj.EnableLlgr == nil {
		obj.SetEnableLlgr(false)
	}
	if obj.obj.StaleTime == nil {
		obj.SetStaleTime(10)
	}
	if obj.obj.EnableNotification == nil {
		obj.SetEnableNotification(true)
	}

}
