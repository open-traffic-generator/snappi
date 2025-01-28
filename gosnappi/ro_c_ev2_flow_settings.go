package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2FlowSettings *****
type roCEv2FlowSettings struct {
	validation
	obj             *otg.RoCEv2FlowSettings
	marshaller      marshalRoCEv2FlowSettings
	unMarshaller    unMarshalRoCEv2FlowSettings
	localEndHolder  RoCEv2FlowSettingsLocalEnd
	remoteEndHolder RoCEv2FlowSettingsRemoteEnd
}

func NewRoCEv2FlowSettings() RoCEv2FlowSettings {
	obj := roCEv2FlowSettings{obj: &otg.RoCEv2FlowSettings{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2FlowSettings) msg() *otg.RoCEv2FlowSettings {
	return obj.obj
}

func (obj *roCEv2FlowSettings) setMsg(msg *otg.RoCEv2FlowSettings) RoCEv2FlowSettings {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2FlowSettings struct {
	obj *roCEv2FlowSettings
}

type marshalRoCEv2FlowSettings interface {
	// ToProto marshals RoCEv2FlowSettings to protobuf object *otg.RoCEv2FlowSettings
	ToProto() (*otg.RoCEv2FlowSettings, error)
	// ToPbText marshals RoCEv2FlowSettings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2FlowSettings to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2FlowSettings to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2FlowSettings struct {
	obj *roCEv2FlowSettings
}

type unMarshalRoCEv2FlowSettings interface {
	// FromProto unmarshals RoCEv2FlowSettings from protobuf object *otg.RoCEv2FlowSettings
	FromProto(msg *otg.RoCEv2FlowSettings) (RoCEv2FlowSettings, error)
	// FromPbText unmarshals RoCEv2FlowSettings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2FlowSettings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2FlowSettings from JSON text
	FromJson(value string) error
}

func (obj *roCEv2FlowSettings) Marshal() marshalRoCEv2FlowSettings {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2FlowSettings{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2FlowSettings) Unmarshal() unMarshalRoCEv2FlowSettings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2FlowSettings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2FlowSettings) ToProto() (*otg.RoCEv2FlowSettings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2FlowSettings) FromProto(msg *otg.RoCEv2FlowSettings) (RoCEv2FlowSettings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2FlowSettings) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2FlowSettings) FromPbText(value string) error {
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

func (m *marshalroCEv2FlowSettings) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2FlowSettings) FromYaml(value string) error {
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

func (m *marshalroCEv2FlowSettings) ToJson() (string, error) {
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

func (m *unMarshalroCEv2FlowSettings) FromJson(value string) error {
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

func (obj *roCEv2FlowSettings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2FlowSettings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2FlowSettings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2FlowSettings) Clone() (RoCEv2FlowSettings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2FlowSettings()
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

func (obj *roCEv2FlowSettings) setNil() {
	obj.localEndHolder = nil
	obj.remoteEndHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RoCEv2FlowSettings is this section has two views, Local End and Remote End.
// Both views have same configurations. However, the remote and local peer IP addresses are interchanged.
// This configuration allows you to configure RDMA flow over the same QP number from same source and destination.
// Default value for commands at Remote peer is set to None.
// So that by default, this remote peer does not initiate any traffic flow.
type RoCEv2FlowSettings interface {
	Validation
	// msg marshals RoCEv2FlowSettings to protobuf object *otg.RoCEv2FlowSettings
	// and doesn't set defaults
	msg() *otg.RoCEv2FlowSettings
	// setMsg unmarshals RoCEv2FlowSettings from protobuf object *otg.RoCEv2FlowSettings
	// and doesn't set defaults
	setMsg(*otg.RoCEv2FlowSettings) RoCEv2FlowSettings
	// provides marshal interface
	Marshal() marshalRoCEv2FlowSettings
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2FlowSettings
	// validate validates RoCEv2FlowSettings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2FlowSettings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocalEnd returns RoCEv2FlowSettingsLocalEnd, set in RoCEv2FlowSettings.
	// RoCEv2FlowSettingsLocalEnd is local End Flow Settings
	LocalEnd() RoCEv2FlowSettingsLocalEnd
	// SetLocalEnd assigns RoCEv2FlowSettingsLocalEnd provided by user to RoCEv2FlowSettings.
	// RoCEv2FlowSettingsLocalEnd is local End Flow Settings
	SetLocalEnd(value RoCEv2FlowSettingsLocalEnd) RoCEv2FlowSettings
	// HasLocalEnd checks if LocalEnd has been set in RoCEv2FlowSettings
	HasLocalEnd() bool
	// RemoteEnd returns RoCEv2FlowSettingsRemoteEnd, set in RoCEv2FlowSettings.
	// RoCEv2FlowSettingsRemoteEnd is remote End Flow Settings
	RemoteEnd() RoCEv2FlowSettingsRemoteEnd
	// SetRemoteEnd assigns RoCEv2FlowSettingsRemoteEnd provided by user to RoCEv2FlowSettings.
	// RoCEv2FlowSettingsRemoteEnd is remote End Flow Settings
	SetRemoteEnd(value RoCEv2FlowSettingsRemoteEnd) RoCEv2FlowSettings
	// HasRemoteEnd checks if RemoteEnd has been set in RoCEv2FlowSettings
	HasRemoteEnd() bool
	setNil()
}

// description is TBD
// LocalEnd returns a RoCEv2FlowSettingsLocalEnd
func (obj *roCEv2FlowSettings) LocalEnd() RoCEv2FlowSettingsLocalEnd {
	if obj.obj.LocalEnd == nil {
		obj.obj.LocalEnd = NewRoCEv2FlowSettingsLocalEnd().msg()
	}
	if obj.localEndHolder == nil {
		obj.localEndHolder = &roCEv2FlowSettingsLocalEnd{obj: obj.obj.LocalEnd}
	}
	return obj.localEndHolder
}

// description is TBD
// LocalEnd returns a RoCEv2FlowSettingsLocalEnd
func (obj *roCEv2FlowSettings) HasLocalEnd() bool {
	return obj.obj.LocalEnd != nil
}

// description is TBD
// SetLocalEnd sets the RoCEv2FlowSettingsLocalEnd value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetLocalEnd(value RoCEv2FlowSettingsLocalEnd) RoCEv2FlowSettings {

	obj.localEndHolder = nil
	obj.obj.LocalEnd = value.msg()

	return obj
}

// description is TBD
// RemoteEnd returns a RoCEv2FlowSettingsRemoteEnd
func (obj *roCEv2FlowSettings) RemoteEnd() RoCEv2FlowSettingsRemoteEnd {
	if obj.obj.RemoteEnd == nil {
		obj.obj.RemoteEnd = NewRoCEv2FlowSettingsRemoteEnd().msg()
	}
	if obj.remoteEndHolder == nil {
		obj.remoteEndHolder = &roCEv2FlowSettingsRemoteEnd{obj: obj.obj.RemoteEnd}
	}
	return obj.remoteEndHolder
}

// description is TBD
// RemoteEnd returns a RoCEv2FlowSettingsRemoteEnd
func (obj *roCEv2FlowSettings) HasRemoteEnd() bool {
	return obj.obj.RemoteEnd != nil
}

// description is TBD
// SetRemoteEnd sets the RoCEv2FlowSettingsRemoteEnd value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetRemoteEnd(value RoCEv2FlowSettingsRemoteEnd) RoCEv2FlowSettings {

	obj.remoteEndHolder = nil
	obj.obj.RemoteEnd = value.msg()

	return obj
}

func (obj *roCEv2FlowSettings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocalEnd != nil {

		obj.LocalEnd().validateObj(vObj, set_default)
	}

	if obj.obj.RemoteEnd != nil {

		obj.RemoteEnd().validateObj(vObj, set_default)
	}

}

func (obj *roCEv2FlowSettings) setDefault() {

}
