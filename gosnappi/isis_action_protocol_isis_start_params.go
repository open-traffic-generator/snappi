package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisActionProtocolIsisStartParams *****
type isisActionProtocolIsisStartParams struct {
	validation
	obj          *otg.IsisActionProtocolIsisStartParams
	marshaller   marshalIsisActionProtocolIsisStartParams
	unMarshaller unMarshalIsisActionProtocolIsisStartParams
}

func NewIsisActionProtocolIsisStartParams() IsisActionProtocolIsisStartParams {
	obj := isisActionProtocolIsisStartParams{obj: &otg.IsisActionProtocolIsisStartParams{}}
	obj.setDefault()
	return &obj
}

func (obj *isisActionProtocolIsisStartParams) msg() *otg.IsisActionProtocolIsisStartParams {
	return obj.obj
}

func (obj *isisActionProtocolIsisStartParams) setMsg(msg *otg.IsisActionProtocolIsisStartParams) IsisActionProtocolIsisStartParams {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisActionProtocolIsisStartParams struct {
	obj *isisActionProtocolIsisStartParams
}

type marshalIsisActionProtocolIsisStartParams interface {
	// ToProto marshals IsisActionProtocolIsisStartParams to protobuf object *otg.IsisActionProtocolIsisStartParams
	ToProto() (*otg.IsisActionProtocolIsisStartParams, error)
	// ToPbText marshals IsisActionProtocolIsisStartParams to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisActionProtocolIsisStartParams to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisActionProtocolIsisStartParams to JSON text
	ToJson() (string, error)
}

type unMarshalisisActionProtocolIsisStartParams struct {
	obj *isisActionProtocolIsisStartParams
}

type unMarshalIsisActionProtocolIsisStartParams interface {
	// FromProto unmarshals IsisActionProtocolIsisStartParams from protobuf object *otg.IsisActionProtocolIsisStartParams
	FromProto(msg *otg.IsisActionProtocolIsisStartParams) (IsisActionProtocolIsisStartParams, error)
	// FromPbText unmarshals IsisActionProtocolIsisStartParams from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisActionProtocolIsisStartParams from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisActionProtocolIsisStartParams from JSON text
	FromJson(value string) error
}

func (obj *isisActionProtocolIsisStartParams) Marshal() marshalIsisActionProtocolIsisStartParams {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisActionProtocolIsisStartParams{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisActionProtocolIsisStartParams) Unmarshal() unMarshalIsisActionProtocolIsisStartParams {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisActionProtocolIsisStartParams{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisActionProtocolIsisStartParams) ToProto() (*otg.IsisActionProtocolIsisStartParams, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisActionProtocolIsisStartParams) FromProto(msg *otg.IsisActionProtocolIsisStartParams) (IsisActionProtocolIsisStartParams, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisActionProtocolIsisStartParams) ToPbText() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisStartParams) FromPbText(value string) error {
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

func (m *marshalisisActionProtocolIsisStartParams) ToYaml() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisStartParams) FromYaml(value string) error {
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

func (m *marshalisisActionProtocolIsisStartParams) ToJson() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisStartParams) FromJson(value string) error {
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

func (obj *isisActionProtocolIsisStartParams) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisStartParams) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisStartParams) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisActionProtocolIsisStartParams) Clone() (IsisActionProtocolIsisStartParams, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisActionProtocolIsisStartParams()
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

// IsisActionProtocolIsisStartParams is configuration for IS-IS Graceful Restart
type IsisActionProtocolIsisStartParams interface {
	Validation
	// msg marshals IsisActionProtocolIsisStartParams to protobuf object *otg.IsisActionProtocolIsisStartParams
	// and doesn't set defaults
	msg() *otg.IsisActionProtocolIsisStartParams
	// setMsg unmarshals IsisActionProtocolIsisStartParams from protobuf object *otg.IsisActionProtocolIsisStartParams
	// and doesn't set defaults
	setMsg(*otg.IsisActionProtocolIsisStartParams) IsisActionProtocolIsisStartParams
	// provides marshal interface
	Marshal() marshalIsisActionProtocolIsisStartParams
	// provides unmarshal interface
	Unmarshal() unMarshalIsisActionProtocolIsisStartParams
	// validate validates IsisActionProtocolIsisStartParams
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisActionProtocolIsisStartParams, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in IsisActionProtocolIsisStartParams.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to IsisActionProtocolIsisStartParams
	SetRouterNames(value []string) IsisActionProtocolIsisStartParams
	// RestartTime returns uint32, set in IsisActionProtocolIsisStartParams.
	RestartTime() uint32
	// SetRestartTime assigns uint32 provided by user to IsisActionProtocolIsisStartParams
	SetRestartTime(value uint32) IsisActionProtocolIsisStartParams
	// HasRestartTime checks if RestartTime has been set in IsisActionProtocolIsisStartParams
	HasRestartTime() bool
}

// The names of device objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *isisActionProtocolIsisStartParams) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of device objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetRouterNames sets the []string value in the IsisActionProtocolIsisStartParams object
func (obj *isisActionProtocolIsisStartParams) SetRouterNames(value []string) IsisActionProtocolIsisStartParams {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart. This paramter is used in starting or restarting operation.
// RestartTime returns a uint32
func (obj *isisActionProtocolIsisStartParams) RestartTime() uint32 {

	return *obj.obj.RestartTime

}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart. This paramter is used in starting or restarting operation.
// RestartTime returns a uint32
func (obj *isisActionProtocolIsisStartParams) HasRestartTime() bool {
	return obj.obj.RestartTime != nil
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart. This paramter is used in starting or restarting operation.
// SetRestartTime sets the uint32 value in the IsisActionProtocolIsisStartParams object
func (obj *isisActionProtocolIsisStartParams) SetRestartTime(value uint32) IsisActionProtocolIsisStartParams {

	obj.obj.RestartTime = &value
	return obj
}

func (obj *isisActionProtocolIsisStartParams) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartTime != nil {

		if *obj.obj.RestartTime < 10 || *obj.obj.RestartTime > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= IsisActionProtocolIsisStartParams.RestartTime <= 4096 but Got %d", *obj.obj.RestartTime))
		}

	}

}

func (obj *isisActionProtocolIsisStartParams) setDefault() {
	if obj.obj.RestartTime == nil {
		obj.SetRestartTime(30)
	}

}
