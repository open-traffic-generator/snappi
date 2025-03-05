package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityDataPlaneRx *****
type secureEntityDataPlaneRx struct {
	validation
	obj          *otg.SecureEntityDataPlaneRx
	marshaller   marshalSecureEntityDataPlaneRx
	unMarshaller unMarshalSecureEntityDataPlaneRx
}

func NewSecureEntityDataPlaneRx() SecureEntityDataPlaneRx {
	obj := secureEntityDataPlaneRx{obj: &otg.SecureEntityDataPlaneRx{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityDataPlaneRx) msg() *otg.SecureEntityDataPlaneRx {
	return obj.obj
}

func (obj *secureEntityDataPlaneRx) setMsg(msg *otg.SecureEntityDataPlaneRx) SecureEntityDataPlaneRx {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityDataPlaneRx struct {
	obj *secureEntityDataPlaneRx
}

type marshalSecureEntityDataPlaneRx interface {
	// ToProto marshals SecureEntityDataPlaneRx to protobuf object *otg.SecureEntityDataPlaneRx
	ToProto() (*otg.SecureEntityDataPlaneRx, error)
	// ToPbText marshals SecureEntityDataPlaneRx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityDataPlaneRx to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityDataPlaneRx to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityDataPlaneRx struct {
	obj *secureEntityDataPlaneRx
}

type unMarshalSecureEntityDataPlaneRx interface {
	// FromProto unmarshals SecureEntityDataPlaneRx from protobuf object *otg.SecureEntityDataPlaneRx
	FromProto(msg *otg.SecureEntityDataPlaneRx) (SecureEntityDataPlaneRx, error)
	// FromPbText unmarshals SecureEntityDataPlaneRx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityDataPlaneRx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityDataPlaneRx from JSON text
	FromJson(value string) error
}

func (obj *secureEntityDataPlaneRx) Marshal() marshalSecureEntityDataPlaneRx {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityDataPlaneRx{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityDataPlaneRx) Unmarshal() unMarshalSecureEntityDataPlaneRx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityDataPlaneRx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityDataPlaneRx) ToProto() (*otg.SecureEntityDataPlaneRx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityDataPlaneRx) FromProto(msg *otg.SecureEntityDataPlaneRx) (SecureEntityDataPlaneRx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityDataPlaneRx) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneRx) FromPbText(value string) error {
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

func (m *marshalsecureEntityDataPlaneRx) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneRx) FromYaml(value string) error {
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

func (m *marshalsecureEntityDataPlaneRx) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneRx) FromJson(value string) error {
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

func (obj *secureEntityDataPlaneRx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityDataPlaneRx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityDataPlaneRx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityDataPlaneRx) Clone() (SecureEntityDataPlaneRx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityDataPlaneRx()
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

// SecureEntityDataPlaneRx is a container for Rx settings of SecY.
type SecureEntityDataPlaneRx interface {
	Validation
	// msg marshals SecureEntityDataPlaneRx to protobuf object *otg.SecureEntityDataPlaneRx
	// and doesn't set defaults
	msg() *otg.SecureEntityDataPlaneRx
	// setMsg unmarshals SecureEntityDataPlaneRx from protobuf object *otg.SecureEntityDataPlaneRx
	// and doesn't set defaults
	setMsg(*otg.SecureEntityDataPlaneRx) SecureEntityDataPlaneRx
	// provides marshal interface
	Marshal() marshalSecureEntityDataPlaneRx
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityDataPlaneRx
	// validate validates SecureEntityDataPlaneRx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityDataPlaneRx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ReplayProtection returns bool, set in SecureEntityDataPlaneRx.
	ReplayProtection() bool
	// SetReplayProtection assigns bool provided by user to SecureEntityDataPlaneRx
	SetReplayProtection(value bool) SecureEntityDataPlaneRx
	// HasReplayProtection checks if ReplayProtection has been set in SecureEntityDataPlaneRx
	HasReplayProtection() bool
	// ReplayWindow returns uint32, set in SecureEntityDataPlaneRx.
	ReplayWindow() uint32
	// SetReplayWindow assigns uint32 provided by user to SecureEntityDataPlaneRx
	SetReplayWindow(value uint32) SecureEntityDataPlaneRx
	// HasReplayWindow checks if ReplayWindow has been set in SecureEntityDataPlaneRx
	HasReplayWindow() bool
}

// Enable replay protection on not.
// ReplayProtection returns a bool
func (obj *secureEntityDataPlaneRx) ReplayProtection() bool {

	return *obj.obj.ReplayProtection

}

// Enable replay protection on not.
// ReplayProtection returns a bool
func (obj *secureEntityDataPlaneRx) HasReplayProtection() bool {
	return obj.obj.ReplayProtection != nil
}

// Enable replay protection on not.
// SetReplayProtection sets the bool value in the SecureEntityDataPlaneRx object
func (obj *secureEntityDataPlaneRx) SetReplayProtection(value bool) SecureEntityDataPlaneRx {

	obj.obj.ReplayProtection = &value
	return obj
}

// Replay window size.
// ReplayWindow returns a uint32
func (obj *secureEntityDataPlaneRx) ReplayWindow() uint32 {

	return *obj.obj.ReplayWindow

}

// Replay window size.
// ReplayWindow returns a uint32
func (obj *secureEntityDataPlaneRx) HasReplayWindow() bool {
	return obj.obj.ReplayWindow != nil
}

// Replay window size.
// SetReplayWindow sets the uint32 value in the SecureEntityDataPlaneRx object
func (obj *secureEntityDataPlaneRx) SetReplayWindow(value uint32) SecureEntityDataPlaneRx {

	obj.obj.ReplayWindow = &value
	return obj
}

func (obj *secureEntityDataPlaneRx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ReplayWindow != nil {

		if *obj.obj.ReplayWindow < 1 || *obj.obj.ReplayWindow > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityDataPlaneRx.ReplayWindow <= 4294967295 but Got %d", *obj.obj.ReplayWindow))
		}

	}

}

func (obj *secureEntityDataPlaneRx) setDefault() {
	if obj.obj.ReplayProtection == nil {
		obj.SetReplayProtection(false)
	}
	if obj.obj.ReplayWindow == nil {
		obj.SetReplayWindow(1)
	}

}
