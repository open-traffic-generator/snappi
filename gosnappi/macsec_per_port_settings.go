package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecPerPortSettings *****
type macsecPerPortSettings struct {
	validation
	obj                        *otg.MacsecPerPortSettings
	marshaller                 marshalMacsecPerPortSettings
	unMarshaller               unMarshalMacsecPerPortSettings
	hardwareAccelerationHolder MacsecHardwareAcceleration
	validateFramesHolder       MacsecValidateFrames
}

func NewMacsecPerPortSettings() MacsecPerPortSettings {
	obj := macsecPerPortSettings{obj: &otg.MacsecPerPortSettings{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecPerPortSettings) msg() *otg.MacsecPerPortSettings {
	return obj.obj
}

func (obj *macsecPerPortSettings) setMsg(msg *otg.MacsecPerPortSettings) MacsecPerPortSettings {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecPerPortSettings struct {
	obj *macsecPerPortSettings
}

type marshalMacsecPerPortSettings interface {
	// ToProto marshals MacsecPerPortSettings to protobuf object *otg.MacsecPerPortSettings
	ToProto() (*otg.MacsecPerPortSettings, error)
	// ToPbText marshals MacsecPerPortSettings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecPerPortSettings to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecPerPortSettings to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecPerPortSettings struct {
	obj *macsecPerPortSettings
}

type unMarshalMacsecPerPortSettings interface {
	// FromProto unmarshals MacsecPerPortSettings from protobuf object *otg.MacsecPerPortSettings
	FromProto(msg *otg.MacsecPerPortSettings) (MacsecPerPortSettings, error)
	// FromPbText unmarshals MacsecPerPortSettings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecPerPortSettings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecPerPortSettings from JSON text
	FromJson(value string) error
}

func (obj *macsecPerPortSettings) Marshal() marshalMacsecPerPortSettings {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecPerPortSettings{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecPerPortSettings) Unmarshal() unMarshalMacsecPerPortSettings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecPerPortSettings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecPerPortSettings) ToProto() (*otg.MacsecPerPortSettings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecPerPortSettings) FromProto(msg *otg.MacsecPerPortSettings) (MacsecPerPortSettings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecPerPortSettings) ToPbText() (string, error) {
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

func (m *unMarshalmacsecPerPortSettings) FromPbText(value string) error {
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

func (m *marshalmacsecPerPortSettings) ToYaml() (string, error) {
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

func (m *unMarshalmacsecPerPortSettings) FromYaml(value string) error {
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

func (m *marshalmacsecPerPortSettings) ToJson() (string, error) {
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

func (m *unMarshalmacsecPerPortSettings) FromJson(value string) error {
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

func (obj *macsecPerPortSettings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecPerPortSettings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecPerPortSettings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecPerPortSettings) Clone() (MacsecPerPortSettings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecPerPortSettings()
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

func (obj *macsecPerPortSettings) setNil() {
	obj.hardwareAccelerationHolder = nil
	obj.validateFramesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecPerPortSettings is mACsec hardware acceleration settings of a test port.
type MacsecPerPortSettings interface {
	Validation
	// msg marshals MacsecPerPortSettings to protobuf object *otg.MacsecPerPortSettings
	// and doesn't set defaults
	msg() *otg.MacsecPerPortSettings
	// setMsg unmarshals MacsecPerPortSettings from protobuf object *otg.MacsecPerPortSettings
	// and doesn't set defaults
	setMsg(*otg.MacsecPerPortSettings) MacsecPerPortSettings
	// provides marshal interface
	Marshal() marshalMacsecPerPortSettings
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecPerPortSettings
	// validate validates MacsecPerPortSettings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecPerPortSettings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HardwareAcceleration returns MacsecHardwareAcceleration, set in MacsecPerPortSettings.
	// MacsecHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
	HardwareAcceleration() MacsecHardwareAcceleration
	// SetHardwareAcceleration assigns MacsecHardwareAcceleration provided by user to MacsecPerPortSettings.
	// MacsecHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
	SetHardwareAcceleration(value MacsecHardwareAcceleration) MacsecPerPortSettings
	// HasHardwareAcceleration checks if HardwareAcceleration has been set in MacsecPerPortSettings
	HasHardwareAcceleration() bool
	// ValidateFrames returns MacsecValidateFrames, set in MacsecPerPortSettings.
	// MacsecValidateFrames is controls validation of received frames.
	ValidateFrames() MacsecValidateFrames
	// SetValidateFrames assigns MacsecValidateFrames provided by user to MacsecPerPortSettings.
	// MacsecValidateFrames is controls validation of received frames.
	SetValidateFrames(value MacsecValidateFrames) MacsecPerPortSettings
	// HasValidateFrames checks if ValidateFrames has been set in MacsecPerPortSettings
	HasValidateFrames() bool
	setNil()
}

// description is TBD
// HardwareAcceleration returns a MacsecHardwareAcceleration
func (obj *macsecPerPortSettings) HardwareAcceleration() MacsecHardwareAcceleration {
	if obj.obj.HardwareAcceleration == nil {
		obj.obj.HardwareAcceleration = NewMacsecHardwareAcceleration().msg()
	}
	if obj.hardwareAccelerationHolder == nil {
		obj.hardwareAccelerationHolder = &macsecHardwareAcceleration{obj: obj.obj.HardwareAcceleration}
	}
	return obj.hardwareAccelerationHolder
}

// description is TBD
// HardwareAcceleration returns a MacsecHardwareAcceleration
func (obj *macsecPerPortSettings) HasHardwareAcceleration() bool {
	return obj.obj.HardwareAcceleration != nil
}

// description is TBD
// SetHardwareAcceleration sets the MacsecHardwareAcceleration value in the MacsecPerPortSettings object
func (obj *macsecPerPortSettings) SetHardwareAcceleration(value MacsecHardwareAcceleration) MacsecPerPortSettings {

	obj.hardwareAccelerationHolder = nil
	obj.obj.HardwareAcceleration = value.msg()

	return obj
}

// description is TBD
// ValidateFrames returns a MacsecValidateFrames
func (obj *macsecPerPortSettings) ValidateFrames() MacsecValidateFrames {
	if obj.obj.ValidateFrames == nil {
		obj.obj.ValidateFrames = NewMacsecValidateFrames().msg()
	}
	if obj.validateFramesHolder == nil {
		obj.validateFramesHolder = &macsecValidateFrames{obj: obj.obj.ValidateFrames}
	}
	return obj.validateFramesHolder
}

// description is TBD
// ValidateFrames returns a MacsecValidateFrames
func (obj *macsecPerPortSettings) HasValidateFrames() bool {
	return obj.obj.ValidateFrames != nil
}

// description is TBD
// SetValidateFrames sets the MacsecValidateFrames value in the MacsecPerPortSettings object
func (obj *macsecPerPortSettings) SetValidateFrames(value MacsecValidateFrames) MacsecPerPortSettings {

	obj.validateFramesHolder = nil
	obj.obj.ValidateFrames = value.msg()

	return obj
}

func (obj *macsecPerPortSettings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.HardwareAcceleration != nil {

		obj.HardwareAcceleration().validateObj(vObj, set_default)
	}

	if obj.obj.ValidateFrames != nil {

		obj.ValidateFrames().validateObj(vObj, set_default)
	}

}

func (obj *macsecPerPortSettings) setDefault() {

}
