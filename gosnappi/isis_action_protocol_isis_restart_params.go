package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisActionProtocolIsisRestartParams *****
type isisActionProtocolIsisRestartParams struct {
	validation
	obj          *otg.IsisActionProtocolIsisRestartParams
	marshaller   marshalIsisActionProtocolIsisRestartParams
	unMarshaller unMarshalIsisActionProtocolIsisRestartParams
}

func NewIsisActionProtocolIsisRestartParams() IsisActionProtocolIsisRestartParams {
	obj := isisActionProtocolIsisRestartParams{obj: &otg.IsisActionProtocolIsisRestartParams{}}
	obj.setDefault()
	return &obj
}

func (obj *isisActionProtocolIsisRestartParams) msg() *otg.IsisActionProtocolIsisRestartParams {
	return obj.obj
}

func (obj *isisActionProtocolIsisRestartParams) setMsg(msg *otg.IsisActionProtocolIsisRestartParams) IsisActionProtocolIsisRestartParams {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisActionProtocolIsisRestartParams struct {
	obj *isisActionProtocolIsisRestartParams
}

type marshalIsisActionProtocolIsisRestartParams interface {
	// ToProto marshals IsisActionProtocolIsisRestartParams to protobuf object *otg.IsisActionProtocolIsisRestartParams
	ToProto() (*otg.IsisActionProtocolIsisRestartParams, error)
	// ToPbText marshals IsisActionProtocolIsisRestartParams to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisActionProtocolIsisRestartParams to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisActionProtocolIsisRestartParams to JSON text
	ToJson() (string, error)
}

type unMarshalisisActionProtocolIsisRestartParams struct {
	obj *isisActionProtocolIsisRestartParams
}

type unMarshalIsisActionProtocolIsisRestartParams interface {
	// FromProto unmarshals IsisActionProtocolIsisRestartParams from protobuf object *otg.IsisActionProtocolIsisRestartParams
	FromProto(msg *otg.IsisActionProtocolIsisRestartParams) (IsisActionProtocolIsisRestartParams, error)
	// FromPbText unmarshals IsisActionProtocolIsisRestartParams from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisActionProtocolIsisRestartParams from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisActionProtocolIsisRestartParams from JSON text
	FromJson(value string) error
}

func (obj *isisActionProtocolIsisRestartParams) Marshal() marshalIsisActionProtocolIsisRestartParams {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisActionProtocolIsisRestartParams{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisActionProtocolIsisRestartParams) Unmarshal() unMarshalIsisActionProtocolIsisRestartParams {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisActionProtocolIsisRestartParams{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisActionProtocolIsisRestartParams) ToProto() (*otg.IsisActionProtocolIsisRestartParams, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisActionProtocolIsisRestartParams) FromProto(msg *otg.IsisActionProtocolIsisRestartParams) (IsisActionProtocolIsisRestartParams, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisActionProtocolIsisRestartParams) ToPbText() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisRestartParams) FromPbText(value string) error {
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

func (m *marshalisisActionProtocolIsisRestartParams) ToYaml() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisRestartParams) FromYaml(value string) error {
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

func (m *marshalisisActionProtocolIsisRestartParams) ToJson() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisRestartParams) FromJson(value string) error {
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

func (obj *isisActionProtocolIsisRestartParams) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisRestartParams) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisRestartParams) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisActionProtocolIsisRestartParams) Clone() (IsisActionProtocolIsisRestartParams, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisActionProtocolIsisRestartParams()
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

// IsisActionProtocolIsisRestartParams is configuration for IS-IS Graceful Restart
type IsisActionProtocolIsisRestartParams interface {
	Validation
	// msg marshals IsisActionProtocolIsisRestartParams to protobuf object *otg.IsisActionProtocolIsisRestartParams
	// and doesn't set defaults
	msg() *otg.IsisActionProtocolIsisRestartParams
	// setMsg unmarshals IsisActionProtocolIsisRestartParams from protobuf object *otg.IsisActionProtocolIsisRestartParams
	// and doesn't set defaults
	setMsg(*otg.IsisActionProtocolIsisRestartParams) IsisActionProtocolIsisRestartParams
	// provides marshal interface
	Marshal() marshalIsisActionProtocolIsisRestartParams
	// provides unmarshal interface
	Unmarshal() unMarshalIsisActionProtocolIsisRestartParams
	// validate validates IsisActionProtocolIsisRestartParams
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisActionProtocolIsisRestartParams, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in IsisActionProtocolIsisRestartParams.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to IsisActionProtocolIsisRestartParams
	SetRouterNames(value []string) IsisActionProtocolIsisRestartParams
	// RestartTime returns uint32, set in IsisActionProtocolIsisRestartParams.
	RestartTime() uint32
	// SetRestartTime assigns uint32 provided by user to IsisActionProtocolIsisRestartParams
	SetRestartTime(value uint32) IsisActionProtocolIsisRestartParams
	// HasRestartTime checks if RestartTime has been set in IsisActionProtocolIsisRestartParams
	HasRestartTime() bool
}

// The names of device objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *isisActionProtocolIsisRestartParams) RouterNames() []string {
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
// SetRouterNames sets the []string value in the IsisActionProtocolIsisRestartParams object
func (obj *isisActionProtocolIsisRestartParams) SetRouterNames(value []string) IsisActionProtocolIsisRestartParams {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart. This paramter is used in starting or restarting operation.
// RestartTime returns a uint32
func (obj *isisActionProtocolIsisRestartParams) RestartTime() uint32 {

	return *obj.obj.RestartTime

}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart. This paramter is used in starting or restarting operation.
// RestartTime returns a uint32
func (obj *isisActionProtocolIsisRestartParams) HasRestartTime() bool {
	return obj.obj.RestartTime != nil
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart. This paramter is used in starting or restarting operation.
// SetRestartTime sets the uint32 value in the IsisActionProtocolIsisRestartParams object
func (obj *isisActionProtocolIsisRestartParams) SetRestartTime(value uint32) IsisActionProtocolIsisRestartParams {

	obj.obj.RestartTime = &value
	return obj
}

func (obj *isisActionProtocolIsisRestartParams) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartTime != nil {

		if *obj.obj.RestartTime < 10 || *obj.obj.RestartTime > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= IsisActionProtocolIsisRestartParams.RestartTime <= 4096 but Got %d", *obj.obj.RestartTime))
		}

	}

}

func (obj *isisActionProtocolIsisRestartParams) setDefault() {
	if obj.obj.RestartTime == nil {
		obj.SetRestartTime(30)
	}

}
