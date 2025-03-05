package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2FlowInfo *****
type roCEv2FlowInfo struct {
	validation
	obj          *otg.RoCEv2FlowInfo
	marshaller   marshalRoCEv2FlowInfo
	unMarshaller unMarshalRoCEv2FlowInfo
}

func NewRoCEv2FlowInfo() RoCEv2FlowInfo {
	obj := roCEv2FlowInfo{obj: &otg.RoCEv2FlowInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2FlowInfo) msg() *otg.RoCEv2FlowInfo {
	return obj.obj
}

func (obj *roCEv2FlowInfo) setMsg(msg *otg.RoCEv2FlowInfo) RoCEv2FlowInfo {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2FlowInfo struct {
	obj *roCEv2FlowInfo
}

type marshalRoCEv2FlowInfo interface {
	// ToProto marshals RoCEv2FlowInfo to protobuf object *otg.RoCEv2FlowInfo
	ToProto() (*otg.RoCEv2FlowInfo, error)
	// ToPbText marshals RoCEv2FlowInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2FlowInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2FlowInfo to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2FlowInfo struct {
	obj *roCEv2FlowInfo
}

type unMarshalRoCEv2FlowInfo interface {
	// FromProto unmarshals RoCEv2FlowInfo from protobuf object *otg.RoCEv2FlowInfo
	FromProto(msg *otg.RoCEv2FlowInfo) (RoCEv2FlowInfo, error)
	// FromPbText unmarshals RoCEv2FlowInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2FlowInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2FlowInfo from JSON text
	FromJson(value string) error
}

func (obj *roCEv2FlowInfo) Marshal() marshalRoCEv2FlowInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2FlowInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2FlowInfo) Unmarshal() unMarshalRoCEv2FlowInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2FlowInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2FlowInfo) ToProto() (*otg.RoCEv2FlowInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2FlowInfo) FromProto(msg *otg.RoCEv2FlowInfo) (RoCEv2FlowInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2FlowInfo) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2FlowInfo) FromPbText(value string) error {
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

func (m *marshalroCEv2FlowInfo) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2FlowInfo) FromYaml(value string) error {
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

func (m *marshalroCEv2FlowInfo) ToJson() (string, error) {
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

func (m *unMarshalroCEv2FlowInfo) FromJson(value string) error {
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

func (obj *roCEv2FlowInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2FlowInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2FlowInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2FlowInfo) Clone() (RoCEv2FlowInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2FlowInfo()
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

// RoCEv2FlowInfo is roCEv2 Flow Info, per flow details.
type RoCEv2FlowInfo interface {
	Validation
	// msg marshals RoCEv2FlowInfo to protobuf object *otg.RoCEv2FlowInfo
	// and doesn't set defaults
	msg() *otg.RoCEv2FlowInfo
	// setMsg unmarshals RoCEv2FlowInfo from protobuf object *otg.RoCEv2FlowInfo
	// and doesn't set defaults
	setMsg(*otg.RoCEv2FlowInfo) RoCEv2FlowInfo
	// provides marshal interface
	Marshal() marshalRoCEv2FlowInfo
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2FlowInfo
	// validate validates RoCEv2FlowInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2FlowInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in RoCEv2FlowInfo.
	Name() string
	// SetName assigns string provided by user to RoCEv2FlowInfo
	SetName(value string) RoCEv2FlowInfo
	// HasName checks if Name has been set in RoCEv2FlowInfo
	HasName() bool
	// TxEndpoints returns string, set in RoCEv2FlowInfo.
	TxEndpoints() string
	// SetTxEndpoints assigns string provided by user to RoCEv2FlowInfo
	SetTxEndpoints(value string) RoCEv2FlowInfo
	// HasTxEndpoints checks if TxEndpoints has been set in RoCEv2FlowInfo
	HasTxEndpoints() bool
	// RxEndpoints returns string, set in RoCEv2FlowInfo.
	RxEndpoints() string
	// SetRxEndpoints assigns string provided by user to RoCEv2FlowInfo
	SetRxEndpoints(value string) RoCEv2FlowInfo
	// HasRxEndpoints checks if RxEndpoints has been set in RoCEv2FlowInfo
	HasRxEndpoints() bool
	// Rate returns RoCEv2FlowInfoRateEnum, set in RoCEv2FlowInfo
	Rate() RoCEv2FlowInfoRateEnum
	// SetRate assigns RoCEv2FlowInfoRateEnum provided by user to RoCEv2FlowInfo
	SetRate(value RoCEv2FlowInfoRateEnum) RoCEv2FlowInfo
	// HasRate checks if Rate has been set in RoCEv2FlowInfo
	HasRate() bool
	// BurstMode returns RoCEv2FlowInfoBurstModeEnum, set in RoCEv2FlowInfo
	BurstMode() RoCEv2FlowInfoBurstModeEnum
	// SetBurstMode assigns RoCEv2FlowInfoBurstModeEnum provided by user to RoCEv2FlowInfo
	SetBurstMode(value RoCEv2FlowInfoBurstModeEnum) RoCEv2FlowInfo
	// HasBurstMode checks if BurstMode has been set in RoCEv2FlowInfo
	HasBurstMode() bool
	// BurstCount returns uint32, set in RoCEv2FlowInfo.
	BurstCount() uint32
	// SetBurstCount assigns uint32 provided by user to RoCEv2FlowInfo
	SetBurstCount(value uint32) RoCEv2FlowInfo
	// HasBurstCount checks if BurstCount has been set in RoCEv2FlowInfo
	HasBurstCount() bool
}

// Name of the flow group.
// Name returns a string
func (obj *roCEv2FlowInfo) Name() string {

	return *obj.obj.Name

}

// Name of the flow group.
// Name returns a string
func (obj *roCEv2FlowInfo) HasName() bool {
	return obj.obj.Name != nil
}

// Name of the flow group.
// SetName sets the string value in the RoCEv2FlowInfo object
func (obj *roCEv2FlowInfo) SetName(value string) RoCEv2FlowInfo {

	obj.obj.Name = &value
	return obj
}

// Tx Endpoint. Empty for RoCEv2.
// TxEndpoints returns a string
func (obj *roCEv2FlowInfo) TxEndpoints() string {

	return *obj.obj.TxEndpoints

}

// Tx Endpoint. Empty for RoCEv2.
// TxEndpoints returns a string
func (obj *roCEv2FlowInfo) HasTxEndpoints() bool {
	return obj.obj.TxEndpoints != nil
}

// Tx Endpoint. Empty for RoCEv2.
// SetTxEndpoints sets the string value in the RoCEv2FlowInfo object
func (obj *roCEv2FlowInfo) SetTxEndpoints(value string) RoCEv2FlowInfo {

	obj.obj.TxEndpoints = &value
	return obj
}

// Rx Endpoint. Empty for RoCEv2.
// RxEndpoints returns a string
func (obj *roCEv2FlowInfo) RxEndpoints() string {

	return *obj.obj.RxEndpoints

}

// Rx Endpoint. Empty for RoCEv2.
// RxEndpoints returns a string
func (obj *roCEv2FlowInfo) HasRxEndpoints() bool {
	return obj.obj.RxEndpoints != nil
}

// Rx Endpoint. Empty for RoCEv2.
// SetRxEndpoints sets the string value in the RoCEv2FlowInfo object
func (obj *roCEv2FlowInfo) SetRxEndpoints(value string) RoCEv2FlowInfo {

	obj.obj.RxEndpoints = &value
	return obj
}

type RoCEv2FlowInfoRateEnum string

// Enum of Rate on RoCEv2FlowInfo
var RoCEv2FlowInfoRate = struct {
	LINE_RATE  RoCEv2FlowInfoRateEnum
	BURST_MODE RoCEv2FlowInfoRateEnum
}{
	LINE_RATE:  RoCEv2FlowInfoRateEnum("line_rate"),
	BURST_MODE: RoCEv2FlowInfoRateEnum("burst_mode"),
}

func (obj *roCEv2FlowInfo) Rate() RoCEv2FlowInfoRateEnum {
	return RoCEv2FlowInfoRateEnum(obj.obj.Rate.Enum().String())
}

// Two valid choices are line_rate and burst_mode, if line_rate is selected, per_port_settings will be used to configure.
// Rate returns a string
func (obj *roCEv2FlowInfo) HasRate() bool {
	return obj.obj.Rate != nil
}

func (obj *roCEv2FlowInfo) SetRate(value RoCEv2FlowInfoRateEnum) RoCEv2FlowInfo {
	intValue, ok := otg.RoCEv2FlowInfo_Rate_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowInfoRateEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowInfo_Rate_Enum(intValue)
	obj.obj.Rate = &enumValue

	return obj
}

type RoCEv2FlowInfoBurstModeEnum string

// Enum of BurstMode on RoCEv2FlowInfo
var RoCEv2FlowInfoBurstMode = struct {
	CONTINUOUS RoCEv2FlowInfoBurstModeEnum
	FIXED      RoCEv2FlowInfoBurstModeEnum
}{
	CONTINUOUS: RoCEv2FlowInfoBurstModeEnum("Continuous"),
	FIXED:      RoCEv2FlowInfoBurstModeEnum("Fixed"),
}

func (obj *roCEv2FlowInfo) BurstMode() RoCEv2FlowInfoBurstModeEnum {
	return RoCEv2FlowInfoBurstModeEnum(obj.obj.BurstMode.Enum().String())
}

// Traffic Burst Mode to applied in RoCEv2 Traffic , Continuous or Fixed.
// BurstMode returns a string
func (obj *roCEv2FlowInfo) HasBurstMode() bool {
	return obj.obj.BurstMode != nil
}

func (obj *roCEv2FlowInfo) SetBurstMode(value RoCEv2FlowInfoBurstModeEnum) RoCEv2FlowInfo {
	intValue, ok := otg.RoCEv2FlowInfo_BurstMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowInfoBurstModeEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowInfo_BurstMode_Enum(intValue)
	obj.obj.BurstMode = &enumValue

	return obj
}

// Burst Count to applied in RoCEv2 Traffic item.
// BurstCount returns a uint32
func (obj *roCEv2FlowInfo) BurstCount() uint32 {

	return *obj.obj.BurstCount

}

// Burst Count to applied in RoCEv2 Traffic item.
// BurstCount returns a uint32
func (obj *roCEv2FlowInfo) HasBurstCount() bool {
	return obj.obj.BurstCount != nil
}

// Burst Count to applied in RoCEv2 Traffic item.
// SetBurstCount sets the uint32 value in the RoCEv2FlowInfo object
func (obj *roCEv2FlowInfo) SetBurstCount(value uint32) RoCEv2FlowInfo {

	obj.obj.BurstCount = &value
	return obj
}

func (obj *roCEv2FlowInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.BurstCount != nil {

		if *obj.obj.BurstCount > 16777216 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowInfo.BurstCount <= 16777216 but Got %d", *obj.obj.BurstCount))
		}

	}

}

func (obj *roCEv2FlowInfo) setDefault() {

}
