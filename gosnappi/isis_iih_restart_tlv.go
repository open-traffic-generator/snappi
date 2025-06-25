package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHRestartTlv *****
type isisIIHRestartTlv struct {
	validation
	obj          *otg.IsisIIHRestartTlv
	marshaller   marshalIsisIIHRestartTlv
	unMarshaller unMarshalIsisIIHRestartTlv
	flagsHolder  IsisIIHRestartFlags
}

func NewIsisIIHRestartTlv() IsisIIHRestartTlv {
	obj := isisIIHRestartTlv{obj: &otg.IsisIIHRestartTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHRestartTlv) msg() *otg.IsisIIHRestartTlv {
	return obj.obj
}

func (obj *isisIIHRestartTlv) setMsg(msg *otg.IsisIIHRestartTlv) IsisIIHRestartTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHRestartTlv struct {
	obj *isisIIHRestartTlv
}

type marshalIsisIIHRestartTlv interface {
	// ToProto marshals IsisIIHRestartTlv to protobuf object *otg.IsisIIHRestartTlv
	ToProto() (*otg.IsisIIHRestartTlv, error)
	// ToPbText marshals IsisIIHRestartTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHRestartTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHRestartTlv to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHRestartTlv struct {
	obj *isisIIHRestartTlv
}

type unMarshalIsisIIHRestartTlv interface {
	// FromProto unmarshals IsisIIHRestartTlv from protobuf object *otg.IsisIIHRestartTlv
	FromProto(msg *otg.IsisIIHRestartTlv) (IsisIIHRestartTlv, error)
	// FromPbText unmarshals IsisIIHRestartTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHRestartTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHRestartTlv from JSON text
	FromJson(value string) error
}

func (obj *isisIIHRestartTlv) Marshal() marshalIsisIIHRestartTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHRestartTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHRestartTlv) Unmarshal() unMarshalIsisIIHRestartTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHRestartTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHRestartTlv) ToProto() (*otg.IsisIIHRestartTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHRestartTlv) FromProto(msg *otg.IsisIIHRestartTlv) (IsisIIHRestartTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHRestartTlv) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHRestartTlv) FromPbText(value string) error {
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

func (m *marshalisisIIHRestartTlv) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHRestartTlv) FromYaml(value string) error {
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

func (m *marshalisisIIHRestartTlv) ToJson() (string, error) {
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

func (m *unMarshalisisIIHRestartTlv) FromJson(value string) error {
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

func (obj *isisIIHRestartTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHRestartTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHRestartTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHRestartTlv) Clone() (IsisIIHRestartTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHRestartTlv()
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

func (obj *isisIIHRestartTlv) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisIIHRestartTlv is container of Restart TLV in IIH PDU. Reference: https://datatracker.ietf.org/doc/html/rfc8706#name-restart-tlv
type IsisIIHRestartTlv interface {
	Validation
	// msg marshals IsisIIHRestartTlv to protobuf object *otg.IsisIIHRestartTlv
	// and doesn't set defaults
	msg() *otg.IsisIIHRestartTlv
	// setMsg unmarshals IsisIIHRestartTlv from protobuf object *otg.IsisIIHRestartTlv
	// and doesn't set defaults
	setMsg(*otg.IsisIIHRestartTlv) IsisIIHRestartTlv
	// provides marshal interface
	Marshal() marshalIsisIIHRestartTlv
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHRestartTlv
	// validate validates IsisIIHRestartTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHRestartTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns IsisIIHRestartFlags, set in IsisIIHRestartTlv.
	// IsisIIHRestartFlags is restarting flags in Restarting TLV in IIH PDU.
	Flags() IsisIIHRestartFlags
	// SetFlags assigns IsisIIHRestartFlags provided by user to IsisIIHRestartTlv.
	// IsisIIHRestartFlags is restarting flags in Restarting TLV in IIH PDU.
	SetFlags(value IsisIIHRestartFlags) IsisIIHRestartTlv
	// HasFlags checks if Flags has been set in IsisIIHRestartTlv
	HasFlags() bool
	// RemainingTime returns uint32, set in IsisIIHRestartTlv.
	RemainingTime() uint32
	// SetRemainingTime assigns uint32 provided by user to IsisIIHRestartTlv
	SetRemainingTime(value uint32) IsisIIHRestartTlv
	// HasRemainingTime checks if RemainingTime has been set in IsisIIHRestartTlv
	HasRemainingTime() bool
	// RestartingNeighborId returns string, set in IsisIIHRestartTlv.
	RestartingNeighborId() string
	// SetRestartingNeighborId assigns string provided by user to IsisIIHRestartTlv
	SetRestartingNeighborId(value string) IsisIIHRestartTlv
	// HasRestartingNeighborId checks if RestartingNeighborId has been set in IsisIIHRestartTlv
	HasRestartingNeighborId() bool
	setNil()
}

// One octet Restart flags in the Restart TLV.
// Flags returns a IsisIIHRestartFlags
func (obj *isisIIHRestartTlv) Flags() IsisIIHRestartFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewIsisIIHRestartFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &isisIIHRestartFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// One octet Restart flags in the Restart TLV.
// Flags returns a IsisIIHRestartFlags
func (obj *isisIIHRestartTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet Restart flags in the Restart TLV.
// SetFlags sets the IsisIIHRestartFlags value in the IsisIIHRestartTlv object
func (obj *isisIIHRestartTlv) SetFlags(value IsisIIHRestartFlags) IsisIIHRestartTlv {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// Remaining Holding Time (in seconds).
// RemainingTime returns a uint32
func (obj *isisIIHRestartTlv) RemainingTime() uint32 {

	return *obj.obj.RemainingTime

}

// Remaining Holding Time (in seconds).
// RemainingTime returns a uint32
func (obj *isisIIHRestartTlv) HasRemainingTime() bool {
	return obj.obj.RemainingTime != nil
}

// Remaining Holding Time (in seconds).
// SetRemainingTime sets the uint32 value in the IsisIIHRestartTlv object
func (obj *isisIIHRestartTlv) SetRemainingTime(value uint32) IsisIIHRestartTlv {

	obj.obj.RemainingTime = &value
	return obj
}

// Restarting Neighbor's System ID in hex format without "0x" at the beginning. e.g. '640000000001'
// RestartingNeighborId returns a string
func (obj *isisIIHRestartTlv) RestartingNeighborId() string {

	return *obj.obj.RestartingNeighborId

}

// Restarting Neighbor's System ID in hex format without "0x" at the beginning. e.g. '640000000001'
// RestartingNeighborId returns a string
func (obj *isisIIHRestartTlv) HasRestartingNeighborId() bool {
	return obj.obj.RestartingNeighborId != nil
}

// Restarting Neighbor's System ID in hex format without "0x" at the beginning. e.g. '640000000001'
// SetRestartingNeighborId sets the string value in the IsisIIHRestartTlv object
func (obj *isisIIHRestartTlv) SetRestartingNeighborId(value string) IsisIIHRestartTlv {

	obj.obj.RestartingNeighborId = &value
	return obj
}

func (obj *isisIIHRestartTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

}

func (obj *isisIIHRestartTlv) setDefault() {

}
