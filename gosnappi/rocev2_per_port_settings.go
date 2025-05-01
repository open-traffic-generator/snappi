package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2PerPortSettings *****
type rocev2PerPortSettings struct {
	validation
	obj                  *otg.Rocev2PerPortSettings
	marshaller           marshalRocev2PerPortSettings
	unMarshaller         unMarshalRocev2PerPortSettings
	cnpHolder            Rocev2CNP
	connectionTypeHolder Rocev2QPConnectionType
	dcqcnSettingsHolder  Rocev2DCQCN
}

func NewRocev2PerPortSettings() Rocev2PerPortSettings {
	obj := rocev2PerPortSettings{obj: &otg.Rocev2PerPortSettings{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2PerPortSettings) msg() *otg.Rocev2PerPortSettings {
	return obj.obj
}

func (obj *rocev2PerPortSettings) setMsg(msg *otg.Rocev2PerPortSettings) Rocev2PerPortSettings {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2PerPortSettings struct {
	obj *rocev2PerPortSettings
}

type marshalRocev2PerPortSettings interface {
	// ToProto marshals Rocev2PerPortSettings to protobuf object *otg.Rocev2PerPortSettings
	ToProto() (*otg.Rocev2PerPortSettings, error)
	// ToPbText marshals Rocev2PerPortSettings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2PerPortSettings to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2PerPortSettings to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2PerPortSettings struct {
	obj *rocev2PerPortSettings
}

type unMarshalRocev2PerPortSettings interface {
	// FromProto unmarshals Rocev2PerPortSettings from protobuf object *otg.Rocev2PerPortSettings
	FromProto(msg *otg.Rocev2PerPortSettings) (Rocev2PerPortSettings, error)
	// FromPbText unmarshals Rocev2PerPortSettings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2PerPortSettings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2PerPortSettings from JSON text
	FromJson(value string) error
}

func (obj *rocev2PerPortSettings) Marshal() marshalRocev2PerPortSettings {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2PerPortSettings{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2PerPortSettings) Unmarshal() unMarshalRocev2PerPortSettings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2PerPortSettings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2PerPortSettings) ToProto() (*otg.Rocev2PerPortSettings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2PerPortSettings) FromProto(msg *otg.Rocev2PerPortSettings) (Rocev2PerPortSettings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2PerPortSettings) ToPbText() (string, error) {
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

func (m *unMarshalrocev2PerPortSettings) FromPbText(value string) error {
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

func (m *marshalrocev2PerPortSettings) ToYaml() (string, error) {
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

func (m *unMarshalrocev2PerPortSettings) FromYaml(value string) error {
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

func (m *marshalrocev2PerPortSettings) ToJson() (string, error) {
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

func (m *unMarshalrocev2PerPortSettings) FromJson(value string) error {
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

func (obj *rocev2PerPortSettings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2PerPortSettings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2PerPortSettings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2PerPortSettings) Clone() (Rocev2PerPortSettings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2PerPortSettings()
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

func (obj *rocev2PerPortSettings) setNil() {
	obj.cnpHolder = nil
	obj.connectionTypeHolder = nil
	obj.dcqcnSettingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2PerPortSettings is data plane traffic flow configuration for a test port.
type Rocev2PerPortSettings interface {
	Validation
	// msg marshals Rocev2PerPortSettings to protobuf object *otg.Rocev2PerPortSettings
	// and doesn't set defaults
	msg() *otg.Rocev2PerPortSettings
	// setMsg unmarshals Rocev2PerPortSettings from protobuf object *otg.Rocev2PerPortSettings
	// and doesn't set defaults
	setMsg(*otg.Rocev2PerPortSettings) Rocev2PerPortSettings
	// provides marshal interface
	Marshal() marshalRocev2PerPortSettings
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2PerPortSettings
	// validate validates Rocev2PerPortSettings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2PerPortSettings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Cnp returns Rocev2CNP, set in Rocev2PerPortSettings.
	// Rocev2CNP is cNP parameters.
	Cnp() Rocev2CNP
	// SetCnp assigns Rocev2CNP provided by user to Rocev2PerPortSettings.
	// Rocev2CNP is cNP parameters.
	SetCnp(value Rocev2CNP) Rocev2PerPortSettings
	// HasCnp checks if Cnp has been set in Rocev2PerPortSettings
	HasCnp() bool
	// ConnectionType returns Rocev2QPConnectionType, set in Rocev2PerPortSettings.
	// Rocev2QPConnectionType is specifies the connection type for the QP, determining what and how the QP transfers data.
	ConnectionType() Rocev2QPConnectionType
	// SetConnectionType assigns Rocev2QPConnectionType provided by user to Rocev2PerPortSettings.
	// Rocev2QPConnectionType is specifies the connection type for the QP, determining what and how the QP transfers data.
	SetConnectionType(value Rocev2QPConnectionType) Rocev2PerPortSettings
	// HasConnectionType checks if ConnectionType has been set in Rocev2PerPortSettings
	HasConnectionType() bool
	// DcqcnSettings returns Rocev2DCQCN, set in Rocev2PerPortSettings.
	// Rocev2DCQCN is roCEv2 DCQCN Settings.
	DcqcnSettings() Rocev2DCQCN
	// SetDcqcnSettings assigns Rocev2DCQCN provided by user to Rocev2PerPortSettings.
	// Rocev2DCQCN is roCEv2 DCQCN Settings.
	SetDcqcnSettings(value Rocev2DCQCN) Rocev2PerPortSettings
	// HasDcqcnSettings checks if DcqcnSettings has been set in Rocev2PerPortSettings
	HasDcqcnSettings() bool
	setNil()
}

// description is TBD
// Cnp returns a Rocev2CNP
func (obj *rocev2PerPortSettings) Cnp() Rocev2CNP {
	if obj.obj.Cnp == nil {
		obj.obj.Cnp = NewRocev2CNP().msg()
	}
	if obj.cnpHolder == nil {
		obj.cnpHolder = &rocev2CNP{obj: obj.obj.Cnp}
	}
	return obj.cnpHolder
}

// description is TBD
// Cnp returns a Rocev2CNP
func (obj *rocev2PerPortSettings) HasCnp() bool {
	return obj.obj.Cnp != nil
}

// description is TBD
// SetCnp sets the Rocev2CNP value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetCnp(value Rocev2CNP) Rocev2PerPortSettings {

	obj.cnpHolder = nil
	obj.obj.Cnp = value.msg()

	return obj
}

// description is TBD
// ConnectionType returns a Rocev2QPConnectionType
func (obj *rocev2PerPortSettings) ConnectionType() Rocev2QPConnectionType {
	if obj.obj.ConnectionType == nil {
		obj.obj.ConnectionType = NewRocev2QPConnectionType().msg()
	}
	if obj.connectionTypeHolder == nil {
		obj.connectionTypeHolder = &rocev2QPConnectionType{obj: obj.obj.ConnectionType}
	}
	return obj.connectionTypeHolder
}

// description is TBD
// ConnectionType returns a Rocev2QPConnectionType
func (obj *rocev2PerPortSettings) HasConnectionType() bool {
	return obj.obj.ConnectionType != nil
}

// description is TBD
// SetConnectionType sets the Rocev2QPConnectionType value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetConnectionType(value Rocev2QPConnectionType) Rocev2PerPortSettings {

	obj.connectionTypeHolder = nil
	obj.obj.ConnectionType = value.msg()

	return obj
}

// description is TBD
// DcqcnSettings returns a Rocev2DCQCN
func (obj *rocev2PerPortSettings) DcqcnSettings() Rocev2DCQCN {
	if obj.obj.DcqcnSettings == nil {
		obj.obj.DcqcnSettings = NewRocev2DCQCN().msg()
	}
	if obj.dcqcnSettingsHolder == nil {
		obj.dcqcnSettingsHolder = &rocev2DCQCN{obj: obj.obj.DcqcnSettings}
	}
	return obj.dcqcnSettingsHolder
}

// description is TBD
// DcqcnSettings returns a Rocev2DCQCN
func (obj *rocev2PerPortSettings) HasDcqcnSettings() bool {
	return obj.obj.DcqcnSettings != nil
}

// description is TBD
// SetDcqcnSettings sets the Rocev2DCQCN value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetDcqcnSettings(value Rocev2DCQCN) Rocev2PerPortSettings {

	obj.dcqcnSettingsHolder = nil
	obj.obj.DcqcnSettings = value.msg()

	return obj
}

func (obj *rocev2PerPortSettings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Cnp != nil {

		obj.Cnp().validateObj(vObj, set_default)
	}

	if obj.obj.ConnectionType != nil {

		obj.ConnectionType().validateObj(vObj, set_default)
	}

	if obj.obj.DcqcnSettings != nil {

		obj.DcqcnSettings().validateObj(vObj, set_default)
	}

}

func (obj *rocev2PerPortSettings) setDefault() {

}
