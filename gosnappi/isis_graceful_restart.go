package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisGracefulRestart *****
type isisGracefulRestart struct {
	validation
	obj          *otg.IsisGracefulRestart
	marshaller   marshalIsisGracefulRestart
	unMarshaller unMarshalIsisGracefulRestart
}

func NewIsisGracefulRestart() IsisGracefulRestart {
	obj := isisGracefulRestart{obj: &otg.IsisGracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *isisGracefulRestart) msg() *otg.IsisGracefulRestart {
	return obj.obj
}

func (obj *isisGracefulRestart) setMsg(msg *otg.IsisGracefulRestart) IsisGracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisGracefulRestart struct {
	obj *isisGracefulRestart
}

type marshalIsisGracefulRestart interface {
	// ToProto marshals IsisGracefulRestart to protobuf object *otg.IsisGracefulRestart
	ToProto() (*otg.IsisGracefulRestart, error)
	// ToPbText marshals IsisGracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisGracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisGracefulRestart to JSON text
	ToJson() (string, error)
}

type unMarshalisisGracefulRestart struct {
	obj *isisGracefulRestart
}

type unMarshalIsisGracefulRestart interface {
	// FromProto unmarshals IsisGracefulRestart from protobuf object *otg.IsisGracefulRestart
	FromProto(msg *otg.IsisGracefulRestart) (IsisGracefulRestart, error)
	// FromPbText unmarshals IsisGracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisGracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisGracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *isisGracefulRestart) Marshal() marshalIsisGracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisGracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisGracefulRestart) Unmarshal() unMarshalIsisGracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisGracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisGracefulRestart) ToProto() (*otg.IsisGracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisGracefulRestart) FromProto(msg *otg.IsisGracefulRestart) (IsisGracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisGracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalisisGracefulRestart) FromPbText(value string) error {
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

func (m *marshalisisGracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalisisGracefulRestart) FromYaml(value string) error {
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

func (m *marshalisisGracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalisisGracefulRestart) FromJson(value string) error {
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

func (obj *isisGracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisGracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisGracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisGracefulRestart) Clone() (IsisGracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisGracefulRestart()
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

// IsisGracefulRestart is contains IS-IS Graceful configuration parameters.
// The emulated IS-IS router that supports Graceful Restart capability MUST include Restart TLV in all transmitted IIHs.
// - If this router is to be acting as the "Helper" mode for a Restarting Router.
// - If this router is to be acting as the "Restarting" mode then "Restart" trigger should be initiated from set_control_action for "isis" under "protocol" choice.
// In that case, "helper_mode" can be configured as true or false for the "Helper" mode as well "Restarting" mode or only "Restaring" router respectively.
// Refrence: https://datatracker.ietf.org/doc/html/rfc8706
type IsisGracefulRestart interface {
	Validation
	// msg marshals IsisGracefulRestart to protobuf object *otg.IsisGracefulRestart
	// and doesn't set defaults
	msg() *otg.IsisGracefulRestart
	// setMsg unmarshals IsisGracefulRestart from protobuf object *otg.IsisGracefulRestart
	// and doesn't set defaults
	setMsg(*otg.IsisGracefulRestart) IsisGracefulRestart
	// provides marshal interface
	Marshal() marshalIsisGracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalIsisGracefulRestart
	// validate validates IsisGracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisGracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HelperMode returns bool, set in IsisGracefulRestart.
	HelperMode() bool
	// SetHelperMode assigns bool provided by user to IsisGracefulRestart
	SetHelperMode(value bool) IsisGracefulRestart
	// HasHelperMode checks if HelperMode has been set in IsisGracefulRestart
	HasHelperMode() bool
}

// The emulated IS-IS router is acting as the "Helper" for the DUT that is restarting.  It acknowledges the Restart TLV sent by the DUT by sending an IIH containing a Restart TLV with the RA (Restart Acknowledgment) bit set.
// HelperMode returns a bool
func (obj *isisGracefulRestart) HelperMode() bool {

	return *obj.obj.HelperMode

}

// The emulated IS-IS router is acting as the "Helper" for the DUT that is restarting.  It acknowledges the Restart TLV sent by the DUT by sending an IIH containing a Restart TLV with the RA (Restart Acknowledgment) bit set.
// HelperMode returns a bool
func (obj *isisGracefulRestart) HasHelperMode() bool {
	return obj.obj.HelperMode != nil
}

// The emulated IS-IS router is acting as the "Helper" for the DUT that is restarting.  It acknowledges the Restart TLV sent by the DUT by sending an IIH containing a Restart TLV with the RA (Restart Acknowledgment) bit set.
// SetHelperMode sets the bool value in the IsisGracefulRestart object
func (obj *isisGracefulRestart) SetHelperMode(value bool) IsisGracefulRestart {

	obj.obj.HelperMode = &value
	return obj
}

func (obj *isisGracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisGracefulRestart) setDefault() {
	if obj.obj.HelperMode == nil {
		obj.SetHelperMode(true)
	}

}
