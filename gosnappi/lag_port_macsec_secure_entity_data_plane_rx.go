package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityDataPlaneRx *****
type lagPortMacsecSecureEntityDataPlaneRx struct {
	validation
	obj                  *otg.LagPortMacsecSecureEntityDataPlaneRx
	marshaller           marshalLagPortMacsecSecureEntityDataPlaneRx
	unMarshaller         unMarshalLagPortMacsecSecureEntityDataPlaneRx
	validateFramesHolder LagPortMacsecSecureEntityDataPlaneRxValidateFrames
}

func NewLagPortMacsecSecureEntityDataPlaneRx() LagPortMacsecSecureEntityDataPlaneRx {
	obj := lagPortMacsecSecureEntityDataPlaneRx{obj: &otg.LagPortMacsecSecureEntityDataPlaneRx{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) msg() *otg.LagPortMacsecSecureEntityDataPlaneRx {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) setMsg(msg *otg.LagPortMacsecSecureEntityDataPlaneRx) LagPortMacsecSecureEntityDataPlaneRx {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityDataPlaneRx struct {
	obj *lagPortMacsecSecureEntityDataPlaneRx
}

type marshalLagPortMacsecSecureEntityDataPlaneRx interface {
	// ToProto marshals LagPortMacsecSecureEntityDataPlaneRx to protobuf object *otg.LagPortMacsecSecureEntityDataPlaneRx
	ToProto() (*otg.LagPortMacsecSecureEntityDataPlaneRx, error)
	// ToPbText marshals LagPortMacsecSecureEntityDataPlaneRx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityDataPlaneRx to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityDataPlaneRx to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityDataPlaneRx struct {
	obj *lagPortMacsecSecureEntityDataPlaneRx
}

type unMarshalLagPortMacsecSecureEntityDataPlaneRx interface {
	// FromProto unmarshals LagPortMacsecSecureEntityDataPlaneRx from protobuf object *otg.LagPortMacsecSecureEntityDataPlaneRx
	FromProto(msg *otg.LagPortMacsecSecureEntityDataPlaneRx) (LagPortMacsecSecureEntityDataPlaneRx, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityDataPlaneRx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityDataPlaneRx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityDataPlaneRx from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) Marshal() marshalLagPortMacsecSecureEntityDataPlaneRx {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityDataPlaneRx{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlaneRx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityDataPlaneRx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityDataPlaneRx) ToProto() (*otg.LagPortMacsecSecureEntityDataPlaneRx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityDataPlaneRx) FromProto(msg *otg.LagPortMacsecSecureEntityDataPlaneRx) (LagPortMacsecSecureEntityDataPlaneRx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityDataPlaneRx) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneRx) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlaneRx) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneRx) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlaneRx) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneRx) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityDataPlaneRx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) Clone() (LagPortMacsecSecureEntityDataPlaneRx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityDataPlaneRx()
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

func (obj *lagPortMacsecSecureEntityDataPlaneRx) setNil() {
	obj.validateFramesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPortMacsecSecureEntityDataPlaneRx is a container for Rx settings of SecY.
type LagPortMacsecSecureEntityDataPlaneRx interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityDataPlaneRx to protobuf object *otg.LagPortMacsecSecureEntityDataPlaneRx
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityDataPlaneRx
	// setMsg unmarshals LagPortMacsecSecureEntityDataPlaneRx from protobuf object *otg.LagPortMacsecSecureEntityDataPlaneRx
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityDataPlaneRx) LagPortMacsecSecureEntityDataPlaneRx
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityDataPlaneRx
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlaneRx
	// validate validates LagPortMacsecSecureEntityDataPlaneRx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityDataPlaneRx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ReplayProtection returns bool, set in LagPortMacsecSecureEntityDataPlaneRx.
	ReplayProtection() bool
	// SetReplayProtection assigns bool provided by user to LagPortMacsecSecureEntityDataPlaneRx
	SetReplayProtection(value bool) LagPortMacsecSecureEntityDataPlaneRx
	// HasReplayProtection checks if ReplayProtection has been set in LagPortMacsecSecureEntityDataPlaneRx
	HasReplayProtection() bool
	// ReplayWindow returns uint32, set in LagPortMacsecSecureEntityDataPlaneRx.
	ReplayWindow() uint32
	// SetReplayWindow assigns uint32 provided by user to LagPortMacsecSecureEntityDataPlaneRx
	SetReplayWindow(value uint32) LagPortMacsecSecureEntityDataPlaneRx
	// HasReplayWindow checks if ReplayWindow has been set in LagPortMacsecSecureEntityDataPlaneRx
	HasReplayWindow() bool
	// ValidateFrames returns LagPortMacsecSecureEntityDataPlaneRxValidateFrames, set in LagPortMacsecSecureEntityDataPlaneRx.
	// LagPortMacsecSecureEntityDataPlaneRxValidateFrames is controls validation of received frames.
	ValidateFrames() LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	// SetValidateFrames assigns LagPortMacsecSecureEntityDataPlaneRxValidateFrames provided by user to LagPortMacsecSecureEntityDataPlaneRx.
	// LagPortMacsecSecureEntityDataPlaneRxValidateFrames is controls validation of received frames.
	SetValidateFrames(value LagPortMacsecSecureEntityDataPlaneRxValidateFrames) LagPortMacsecSecureEntityDataPlaneRx
	// HasValidateFrames checks if ValidateFrames has been set in LagPortMacsecSecureEntityDataPlaneRx
	HasValidateFrames() bool
	setNil()
}

// Enable replay protection or not.
// ReplayProtection returns a bool
func (obj *lagPortMacsecSecureEntityDataPlaneRx) ReplayProtection() bool {

	return *obj.obj.ReplayProtection

}

// Enable replay protection or not.
// ReplayProtection returns a bool
func (obj *lagPortMacsecSecureEntityDataPlaneRx) HasReplayProtection() bool {
	return obj.obj.ReplayProtection != nil
}

// Enable replay protection or not.
// SetReplayProtection sets the bool value in the LagPortMacsecSecureEntityDataPlaneRx object
func (obj *lagPortMacsecSecureEntityDataPlaneRx) SetReplayProtection(value bool) LagPortMacsecSecureEntityDataPlaneRx {

	obj.obj.ReplayProtection = &value
	return obj
}

// Replay window size.
// ReplayWindow returns a uint32
func (obj *lagPortMacsecSecureEntityDataPlaneRx) ReplayWindow() uint32 {

	return *obj.obj.ReplayWindow

}

// Replay window size.
// ReplayWindow returns a uint32
func (obj *lagPortMacsecSecureEntityDataPlaneRx) HasReplayWindow() bool {
	return obj.obj.ReplayWindow != nil
}

// Replay window size.
// SetReplayWindow sets the uint32 value in the LagPortMacsecSecureEntityDataPlaneRx object
func (obj *lagPortMacsecSecureEntityDataPlaneRx) SetReplayWindow(value uint32) LagPortMacsecSecureEntityDataPlaneRx {

	obj.obj.ReplayWindow = &value
	return obj
}

// description is TBD
// ValidateFrames returns a LagPortMacsecSecureEntityDataPlaneRxValidateFrames
func (obj *lagPortMacsecSecureEntityDataPlaneRx) ValidateFrames() LagPortMacsecSecureEntityDataPlaneRxValidateFrames {
	if obj.obj.ValidateFrames == nil {
		obj.obj.ValidateFrames = NewLagPortMacsecSecureEntityDataPlaneRxValidateFrames().msg()
	}
	if obj.validateFramesHolder == nil {
		obj.validateFramesHolder = &lagPortMacsecSecureEntityDataPlaneRxValidateFrames{obj: obj.obj.ValidateFrames}
	}
	return obj.validateFramesHolder
}

// description is TBD
// ValidateFrames returns a LagPortMacsecSecureEntityDataPlaneRxValidateFrames
func (obj *lagPortMacsecSecureEntityDataPlaneRx) HasValidateFrames() bool {
	return obj.obj.ValidateFrames != nil
}

// description is TBD
// SetValidateFrames sets the LagPortMacsecSecureEntityDataPlaneRxValidateFrames value in the LagPortMacsecSecureEntityDataPlaneRx object
func (obj *lagPortMacsecSecureEntityDataPlaneRx) SetValidateFrames(value LagPortMacsecSecureEntityDataPlaneRxValidateFrames) LagPortMacsecSecureEntityDataPlaneRx {

	obj.validateFramesHolder = nil
	obj.obj.ValidateFrames = value.msg()

	return obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ReplayWindow != nil {

		if *obj.obj.ReplayWindow < 1 || *obj.obj.ReplayWindow > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= LagPortMacsecSecureEntityDataPlaneRx.ReplayWindow <= 4294967295 but Got %d", *obj.obj.ReplayWindow))
		}

	}

	if obj.obj.ValidateFrames != nil {

		obj.ValidateFrames().validateObj(vObj, set_default)
	}

}

func (obj *lagPortMacsecSecureEntityDataPlaneRx) setDefault() {
	if obj.obj.ReplayProtection == nil {
		obj.SetReplayProtection(false)
	}
	if obj.obj.ReplayWindow == nil {
		obj.SetReplayWindow(1)
	}

}
