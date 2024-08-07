package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisInterface *****
type isisInterface struct {
	validation
	obj                      *otg.IsisInterface
	marshaller               marshalIsisInterface
	unMarshaller             unMarshalIsisInterface
	l1SettingsHolder         IsisInterfaceLevel
	l2SettingsHolder         IsisInterfaceLevel
	multiTopologyIdsHolder   IsisInterfaceIsisMTIter
	trafficEngineeringHolder IsisInterfaceLinkStateTEIter
	authenticationHolder     IsisInterfaceAuthentication
	advancedHolder           IsisInterfaceAdvanced
	linkProtectionHolder     IsisInterfaceLinkProtection
}

func NewIsisInterface() IsisInterface {
	obj := isisInterface{obj: &otg.IsisInterface{}}
	obj.setDefault()
	return &obj
}

func (obj *isisInterface) msg() *otg.IsisInterface {
	return obj.obj
}

func (obj *isisInterface) setMsg(msg *otg.IsisInterface) IsisInterface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisInterface struct {
	obj *isisInterface
}

type marshalIsisInterface interface {
	// ToProto marshals IsisInterface to protobuf object *otg.IsisInterface
	ToProto() (*otg.IsisInterface, error)
	// ToPbText marshals IsisInterface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisInterface to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisInterface to JSON text
	ToJson() (string, error)
}

type unMarshalisisInterface struct {
	obj *isisInterface
}

type unMarshalIsisInterface interface {
	// FromProto unmarshals IsisInterface from protobuf object *otg.IsisInterface
	FromProto(msg *otg.IsisInterface) (IsisInterface, error)
	// FromPbText unmarshals IsisInterface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisInterface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisInterface from JSON text
	FromJson(value string) error
}

func (obj *isisInterface) Marshal() marshalIsisInterface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisInterface{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisInterface) Unmarshal() unMarshalIsisInterface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisInterface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisInterface) ToProto() (*otg.IsisInterface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisInterface) FromProto(msg *otg.IsisInterface) (IsisInterface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisInterface) ToPbText() (string, error) {
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

func (m *unMarshalisisInterface) FromPbText(value string) error {
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

func (m *marshalisisInterface) ToYaml() (string, error) {
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

func (m *unMarshalisisInterface) FromYaml(value string) error {
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

func (m *marshalisisInterface) ToJson() (string, error) {
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

func (m *unMarshalisisInterface) FromJson(value string) error {
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

func (obj *isisInterface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisInterface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisInterface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisInterface) Clone() (IsisInterface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisInterface()
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

func (obj *isisInterface) setNil() {
	obj.l1SettingsHolder = nil
	obj.l2SettingsHolder = nil
	obj.multiTopologyIdsHolder = nil
	obj.trafficEngineeringHolder = nil
	obj.authenticationHolder = nil
	obj.advancedHolder = nil
	obj.linkProtectionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisInterface is configuration for single ISIS interface.
type IsisInterface interface {
	Validation
	// msg marshals IsisInterface to protobuf object *otg.IsisInterface
	// and doesn't set defaults
	msg() *otg.IsisInterface
	// setMsg unmarshals IsisInterface from protobuf object *otg.IsisInterface
	// and doesn't set defaults
	setMsg(*otg.IsisInterface) IsisInterface
	// provides marshal interface
	Marshal() marshalIsisInterface
	// provides unmarshal interface
	Unmarshal() unMarshalIsisInterface
	// validate validates IsisInterface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisInterface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthName returns string, set in IsisInterface.
	EthName() string
	// SetEthName assigns string provided by user to IsisInterface
	SetEthName(value string) IsisInterface
	// Metric returns uint32, set in IsisInterface.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to IsisInterface
	SetMetric(value uint32) IsisInterface
	// HasMetric checks if Metric has been set in IsisInterface
	HasMetric() bool
	// NetworkType returns IsisInterfaceNetworkTypeEnum, set in IsisInterface
	NetworkType() IsisInterfaceNetworkTypeEnum
	// SetNetworkType assigns IsisInterfaceNetworkTypeEnum provided by user to IsisInterface
	SetNetworkType(value IsisInterfaceNetworkTypeEnum) IsisInterface
	// HasNetworkType checks if NetworkType has been set in IsisInterface
	HasNetworkType() bool
	// LevelType returns IsisInterfaceLevelTypeEnum, set in IsisInterface
	LevelType() IsisInterfaceLevelTypeEnum
	// SetLevelType assigns IsisInterfaceLevelTypeEnum provided by user to IsisInterface
	SetLevelType(value IsisInterfaceLevelTypeEnum) IsisInterface
	// HasLevelType checks if LevelType has been set in IsisInterface
	HasLevelType() bool
	// L1Settings returns IsisInterfaceLevel, set in IsisInterface.
	// IsisInterfaceLevel is configuration for the properties of Level 1 Hello.
	L1Settings() IsisInterfaceLevel
	// SetL1Settings assigns IsisInterfaceLevel provided by user to IsisInterface.
	// IsisInterfaceLevel is configuration for the properties of Level 1 Hello.
	SetL1Settings(value IsisInterfaceLevel) IsisInterface
	// HasL1Settings checks if L1Settings has been set in IsisInterface
	HasL1Settings() bool
	// L2Settings returns IsisInterfaceLevel, set in IsisInterface.
	// IsisInterfaceLevel is configuration for the properties of Level 1 Hello.
	L2Settings() IsisInterfaceLevel
	// SetL2Settings assigns IsisInterfaceLevel provided by user to IsisInterface.
	// IsisInterfaceLevel is configuration for the properties of Level 1 Hello.
	SetL2Settings(value IsisInterfaceLevel) IsisInterface
	// HasL2Settings checks if L2Settings has been set in IsisInterface
	HasL2Settings() bool
	// MultiTopologyIds returns IsisInterfaceIsisMTIterIter, set in IsisInterface
	MultiTopologyIds() IsisInterfaceIsisMTIter
	// TrafficEngineering returns IsisInterfaceLinkStateTEIterIter, set in IsisInterface
	TrafficEngineering() IsisInterfaceLinkStateTEIter
	// Authentication returns IsisInterfaceAuthentication, set in IsisInterface.
	// IsisInterfaceAuthentication is optional container for circuit authentication properties.
	Authentication() IsisInterfaceAuthentication
	// SetAuthentication assigns IsisInterfaceAuthentication provided by user to IsisInterface.
	// IsisInterfaceAuthentication is optional container for circuit authentication properties.
	SetAuthentication(value IsisInterfaceAuthentication) IsisInterface
	// HasAuthentication checks if Authentication has been set in IsisInterface
	HasAuthentication() bool
	// Advanced returns IsisInterfaceAdvanced, set in IsisInterface.
	// IsisInterfaceAdvanced is optional container for advanced interface properties.
	Advanced() IsisInterfaceAdvanced
	// SetAdvanced assigns IsisInterfaceAdvanced provided by user to IsisInterface.
	// IsisInterfaceAdvanced is optional container for advanced interface properties.
	SetAdvanced(value IsisInterfaceAdvanced) IsisInterface
	// HasAdvanced checks if Advanced has been set in IsisInterface
	HasAdvanced() bool
	// LinkProtection returns IsisInterfaceLinkProtection, set in IsisInterface.
	// IsisInterfaceLinkProtection is optional container for the link protection sub TLV (type 20).
	LinkProtection() IsisInterfaceLinkProtection
	// SetLinkProtection assigns IsisInterfaceLinkProtection provided by user to IsisInterface.
	// IsisInterfaceLinkProtection is optional container for the link protection sub TLV (type 20).
	SetLinkProtection(value IsisInterfaceLinkProtection) IsisInterface
	// HasLinkProtection checks if LinkProtection has been set in IsisInterface
	HasLinkProtection() bool
	// SrlgValues returns []uint32, set in IsisInterface.
	SrlgValues() []uint32
	// SetSrlgValues assigns []uint32 provided by user to IsisInterface
	SetSrlgValues(value []uint32) IsisInterface
	// Name returns string, set in IsisInterface.
	Name() string
	// SetName assigns string provided by user to IsisInterface
	SetName(value string) IsisInterface
	setNil()
}

// The unique name of the Ethernet interface on which ISIS is running. Two ISIS interfaces cannot share the same Ethernet.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthName returns a string
func (obj *isisInterface) EthName() string {

	return *obj.obj.EthName

}

// The unique name of the Ethernet interface on which ISIS is running. Two ISIS interfaces cannot share the same Ethernet.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthName sets the string value in the IsisInterface object
func (obj *isisInterface) SetEthName(value string) IsisInterface {

	obj.obj.EthName = &value
	return obj
}

// The default metric cost for the interface.
// Metric returns a uint32
func (obj *isisInterface) Metric() uint32 {

	return *obj.obj.Metric

}

// The default metric cost for the interface.
// Metric returns a uint32
func (obj *isisInterface) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The default metric cost for the interface.
// SetMetric sets the uint32 value in the IsisInterface object
func (obj *isisInterface) SetMetric(value uint32) IsisInterface {

	obj.obj.Metric = &value
	return obj
}

type IsisInterfaceNetworkTypeEnum string

// Enum of NetworkType on IsisInterface
var IsisInterfaceNetworkType = struct {
	BROADCAST      IsisInterfaceNetworkTypeEnum
	POINT_TO_POINT IsisInterfaceNetworkTypeEnum
}{
	BROADCAST:      IsisInterfaceNetworkTypeEnum("broadcast"),
	POINT_TO_POINT: IsisInterfaceNetworkTypeEnum("point_to_point"),
}

func (obj *isisInterface) NetworkType() IsisInterfaceNetworkTypeEnum {
	return IsisInterfaceNetworkTypeEnum(obj.obj.NetworkType.Enum().String())
}

// The type of network link.
// NetworkType returns a string
func (obj *isisInterface) HasNetworkType() bool {
	return obj.obj.NetworkType != nil
}

func (obj *isisInterface) SetNetworkType(value IsisInterfaceNetworkTypeEnum) IsisInterface {
	intValue, ok := otg.IsisInterface_NetworkType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisInterfaceNetworkTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisInterface_NetworkType_Enum(intValue)
	obj.obj.NetworkType = &enumValue

	return obj
}

type IsisInterfaceLevelTypeEnum string

// Enum of LevelType on IsisInterface
var IsisInterfaceLevelType = struct {
	LEVEL_1   IsisInterfaceLevelTypeEnum
	LEVEL_2   IsisInterfaceLevelTypeEnum
	LEVEL_1_2 IsisInterfaceLevelTypeEnum
}{
	LEVEL_1:   IsisInterfaceLevelTypeEnum("level_1"),
	LEVEL_2:   IsisInterfaceLevelTypeEnum("level_2"),
	LEVEL_1_2: IsisInterfaceLevelTypeEnum("level_1_2"),
}

func (obj *isisInterface) LevelType() IsisInterfaceLevelTypeEnum {
	return IsisInterfaceLevelTypeEnum(obj.obj.LevelType.Enum().String())
}

// This indicates whether this router is participating in Level-1 (L1),
// Level-2 (L2) or both L1 and L2 domains on this interface.
// LevelType returns a string
func (obj *isisInterface) HasLevelType() bool {
	return obj.obj.LevelType != nil
}

func (obj *isisInterface) SetLevelType(value IsisInterfaceLevelTypeEnum) IsisInterface {
	intValue, ok := otg.IsisInterface_LevelType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisInterfaceLevelTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisInterface_LevelType_Enum(intValue)
	obj.obj.LevelType = &enumValue

	return obj
}

// Settings of Level 1 Hello.
// L1Settings returns a IsisInterfaceLevel
func (obj *isisInterface) L1Settings() IsisInterfaceLevel {
	if obj.obj.L1Settings == nil {
		obj.obj.L1Settings = NewIsisInterfaceLevel().msg()
	}
	if obj.l1SettingsHolder == nil {
		obj.l1SettingsHolder = &isisInterfaceLevel{obj: obj.obj.L1Settings}
	}
	return obj.l1SettingsHolder
}

// Settings of Level 1 Hello.
// L1Settings returns a IsisInterfaceLevel
func (obj *isisInterface) HasL1Settings() bool {
	return obj.obj.L1Settings != nil
}

// Settings of Level 1 Hello.
// SetL1Settings sets the IsisInterfaceLevel value in the IsisInterface object
func (obj *isisInterface) SetL1Settings(value IsisInterfaceLevel) IsisInterface {

	obj.l1SettingsHolder = nil
	obj.obj.L1Settings = value.msg()

	return obj
}

// Settings of Level 2 Hello.
// L2Settings returns a IsisInterfaceLevel
func (obj *isisInterface) L2Settings() IsisInterfaceLevel {
	if obj.obj.L2Settings == nil {
		obj.obj.L2Settings = NewIsisInterfaceLevel().msg()
	}
	if obj.l2SettingsHolder == nil {
		obj.l2SettingsHolder = &isisInterfaceLevel{obj: obj.obj.L2Settings}
	}
	return obj.l2SettingsHolder
}

// Settings of Level 2 Hello.
// L2Settings returns a IsisInterfaceLevel
func (obj *isisInterface) HasL2Settings() bool {
	return obj.obj.L2Settings != nil
}

// Settings of Level 2 Hello.
// SetL2Settings sets the IsisInterfaceLevel value in the IsisInterface object
func (obj *isisInterface) SetL2Settings(value IsisInterfaceLevel) IsisInterface {

	obj.l2SettingsHolder = nil
	obj.obj.L2Settings = value.msg()

	return obj
}

// Contains the properties of multiple topologies.
// MultiTopologyIds returns a []IsisMT
func (obj *isisInterface) MultiTopologyIds() IsisInterfaceIsisMTIter {
	if len(obj.obj.MultiTopologyIds) == 0 {
		obj.obj.MultiTopologyIds = []*otg.IsisMT{}
	}
	if obj.multiTopologyIdsHolder == nil {
		obj.multiTopologyIdsHolder = newIsisInterfaceIsisMTIter(&obj.obj.MultiTopologyIds).setMsg(obj)
	}
	return obj.multiTopologyIdsHolder
}

type isisInterfaceIsisMTIter struct {
	obj         *isisInterface
	isisMTSlice []IsisMT
	fieldPtr    *[]*otg.IsisMT
}

func newIsisInterfaceIsisMTIter(ptr *[]*otg.IsisMT) IsisInterfaceIsisMTIter {
	return &isisInterfaceIsisMTIter{fieldPtr: ptr}
}

type IsisInterfaceIsisMTIter interface {
	setMsg(*isisInterface) IsisInterfaceIsisMTIter
	Items() []IsisMT
	Add() IsisMT
	Append(items ...IsisMT) IsisInterfaceIsisMTIter
	Set(index int, newObj IsisMT) IsisInterfaceIsisMTIter
	Clear() IsisInterfaceIsisMTIter
	clearHolderSlice() IsisInterfaceIsisMTIter
	appendHolderSlice(item IsisMT) IsisInterfaceIsisMTIter
}

func (obj *isisInterfaceIsisMTIter) setMsg(msg *isisInterface) IsisInterfaceIsisMTIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisMT{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisInterfaceIsisMTIter) Items() []IsisMT {
	return obj.isisMTSlice
}

func (obj *isisInterfaceIsisMTIter) Add() IsisMT {
	newObj := &otg.IsisMT{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisMT{obj: newObj}
	newLibObj.setDefault()
	obj.isisMTSlice = append(obj.isisMTSlice, newLibObj)
	return newLibObj
}

func (obj *isisInterfaceIsisMTIter) Append(items ...IsisMT) IsisInterfaceIsisMTIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisMTSlice = append(obj.isisMTSlice, item)
	}
	return obj
}

func (obj *isisInterfaceIsisMTIter) Set(index int, newObj IsisMT) IsisInterfaceIsisMTIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisMTSlice[index] = newObj
	return obj
}
func (obj *isisInterfaceIsisMTIter) Clear() IsisInterfaceIsisMTIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisMT{}
		obj.isisMTSlice = []IsisMT{}
	}
	return obj
}
func (obj *isisInterfaceIsisMTIter) clearHolderSlice() IsisInterfaceIsisMTIter {
	if len(obj.isisMTSlice) > 0 {
		obj.isisMTSlice = []IsisMT{}
	}
	return obj
}
func (obj *isisInterfaceIsisMTIter) appendHolderSlice(item IsisMT) IsisInterfaceIsisMTIter {
	obj.isisMTSlice = append(obj.isisMTSlice, item)
	return obj
}

// Contains a list of Traffic Engineering attributes.
// TrafficEngineering returns a []LinkStateTE
func (obj *isisInterface) TrafficEngineering() IsisInterfaceLinkStateTEIter {
	if len(obj.obj.TrafficEngineering) == 0 {
		obj.obj.TrafficEngineering = []*otg.LinkStateTE{}
	}
	if obj.trafficEngineeringHolder == nil {
		obj.trafficEngineeringHolder = newIsisInterfaceLinkStateTEIter(&obj.obj.TrafficEngineering).setMsg(obj)
	}
	return obj.trafficEngineeringHolder
}

type isisInterfaceLinkStateTEIter struct {
	obj              *isisInterface
	linkStateTESlice []LinkStateTE
	fieldPtr         *[]*otg.LinkStateTE
}

func newIsisInterfaceLinkStateTEIter(ptr *[]*otg.LinkStateTE) IsisInterfaceLinkStateTEIter {
	return &isisInterfaceLinkStateTEIter{fieldPtr: ptr}
}

type IsisInterfaceLinkStateTEIter interface {
	setMsg(*isisInterface) IsisInterfaceLinkStateTEIter
	Items() []LinkStateTE
	Add() LinkStateTE
	Append(items ...LinkStateTE) IsisInterfaceLinkStateTEIter
	Set(index int, newObj LinkStateTE) IsisInterfaceLinkStateTEIter
	Clear() IsisInterfaceLinkStateTEIter
	clearHolderSlice() IsisInterfaceLinkStateTEIter
	appendHolderSlice(item LinkStateTE) IsisInterfaceLinkStateTEIter
}

func (obj *isisInterfaceLinkStateTEIter) setMsg(msg *isisInterface) IsisInterfaceLinkStateTEIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&linkStateTE{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisInterfaceLinkStateTEIter) Items() []LinkStateTE {
	return obj.linkStateTESlice
}

func (obj *isisInterfaceLinkStateTEIter) Add() LinkStateTE {
	newObj := &otg.LinkStateTE{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &linkStateTE{obj: newObj}
	newLibObj.setDefault()
	obj.linkStateTESlice = append(obj.linkStateTESlice, newLibObj)
	return newLibObj
}

func (obj *isisInterfaceLinkStateTEIter) Append(items ...LinkStateTE) IsisInterfaceLinkStateTEIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.linkStateTESlice = append(obj.linkStateTESlice, item)
	}
	return obj
}

func (obj *isisInterfaceLinkStateTEIter) Set(index int, newObj LinkStateTE) IsisInterfaceLinkStateTEIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.linkStateTESlice[index] = newObj
	return obj
}
func (obj *isisInterfaceLinkStateTEIter) Clear() IsisInterfaceLinkStateTEIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LinkStateTE{}
		obj.linkStateTESlice = []LinkStateTE{}
	}
	return obj
}
func (obj *isisInterfaceLinkStateTEIter) clearHolderSlice() IsisInterfaceLinkStateTEIter {
	if len(obj.linkStateTESlice) > 0 {
		obj.linkStateTESlice = []LinkStateTE{}
	}
	return obj
}
func (obj *isisInterfaceLinkStateTEIter) appendHolderSlice(item LinkStateTE) IsisInterfaceLinkStateTEIter {
	obj.linkStateTESlice = append(obj.linkStateTESlice, item)
	return obj
}

// The Circuit authentication method used for the interfaces on this emulated ISIS v4/v6 router.
// Authentication returns a IsisInterfaceAuthentication
func (obj *isisInterface) Authentication() IsisInterfaceAuthentication {
	if obj.obj.Authentication == nil {
		obj.obj.Authentication = NewIsisInterfaceAuthentication().msg()
	}
	if obj.authenticationHolder == nil {
		obj.authenticationHolder = &isisInterfaceAuthentication{obj: obj.obj.Authentication}
	}
	return obj.authenticationHolder
}

// The Circuit authentication method used for the interfaces on this emulated ISIS v4/v6 router.
// Authentication returns a IsisInterfaceAuthentication
func (obj *isisInterface) HasAuthentication() bool {
	return obj.obj.Authentication != nil
}

// The Circuit authentication method used for the interfaces on this emulated ISIS v4/v6 router.
// SetAuthentication sets the IsisInterfaceAuthentication value in the IsisInterface object
func (obj *isisInterface) SetAuthentication(value IsisInterfaceAuthentication) IsisInterface {

	obj.authenticationHolder = nil
	obj.obj.Authentication = value.msg()

	return obj
}

// Optional container for advanced interface properties.
// Advanced returns a IsisInterfaceAdvanced
func (obj *isisInterface) Advanced() IsisInterfaceAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = NewIsisInterfaceAdvanced().msg()
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &isisInterfaceAdvanced{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// Optional container for advanced interface properties.
// Advanced returns a IsisInterfaceAdvanced
func (obj *isisInterface) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// Optional container for advanced interface properties.
// SetAdvanced sets the IsisInterfaceAdvanced value in the IsisInterface object
func (obj *isisInterface) SetAdvanced(value IsisInterfaceAdvanced) IsisInterface {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Link protection on the ISIS link between two interfaces.
// LinkProtection returns a IsisInterfaceLinkProtection
func (obj *isisInterface) LinkProtection() IsisInterfaceLinkProtection {
	if obj.obj.LinkProtection == nil {
		obj.obj.LinkProtection = NewIsisInterfaceLinkProtection().msg()
	}
	if obj.linkProtectionHolder == nil {
		obj.linkProtectionHolder = &isisInterfaceLinkProtection{obj: obj.obj.LinkProtection}
	}
	return obj.linkProtectionHolder
}

// Link protection on the ISIS link between two interfaces.
// LinkProtection returns a IsisInterfaceLinkProtection
func (obj *isisInterface) HasLinkProtection() bool {
	return obj.obj.LinkProtection != nil
}

// Link protection on the ISIS link between two interfaces.
// SetLinkProtection sets the IsisInterfaceLinkProtection value in the IsisInterface object
func (obj *isisInterface) SetLinkProtection(value IsisInterfaceLinkProtection) IsisInterface {

	obj.linkProtectionHolder = nil
	obj.obj.LinkProtection = value.msg()

	return obj
}

// This contains list of SRLG values for the link between two interfaces.
// SrlgValues returns a []uint32
func (obj *isisInterface) SrlgValues() []uint32 {
	if obj.obj.SrlgValues == nil {
		obj.obj.SrlgValues = make([]uint32, 0)
	}
	return obj.obj.SrlgValues
}

// This contains list of SRLG values for the link between two interfaces.
// SetSrlgValues sets the []uint32 value in the IsisInterface object
func (obj *isisInterface) SetSrlgValues(value []uint32) IsisInterface {

	if obj.obj.SrlgValues == nil {
		obj.obj.SrlgValues = make([]uint32, 0)
	}
	obj.obj.SrlgValues = value

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *isisInterface) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the IsisInterface object
func (obj *isisInterface) SetName(value string) IsisInterface {

	obj.obj.Name = &value
	return obj
}

func (obj *isisInterface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthName is required
	if obj.obj.EthName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthName is required field on interface IsisInterface")
	}

	if obj.obj.Metric != nil {

		if *obj.obj.Metric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisInterface.Metric <= 16777215 but Got %d", *obj.obj.Metric))
		}

	}

	if obj.obj.L1Settings != nil {

		obj.L1Settings().validateObj(vObj, set_default)
	}

	if obj.obj.L2Settings != nil {

		obj.L2Settings().validateObj(vObj, set_default)
	}

	if len(obj.obj.MultiTopologyIds) != 0 {

		if set_default {
			obj.MultiTopologyIds().clearHolderSlice()
			for _, item := range obj.obj.MultiTopologyIds {
				obj.MultiTopologyIds().appendHolderSlice(&isisMT{obj: item})
			}
		}
		for _, item := range obj.MultiTopologyIds().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.TrafficEngineering) != 0 {

		if set_default {
			obj.TrafficEngineering().clearHolderSlice()
			for _, item := range obj.obj.TrafficEngineering {
				obj.TrafficEngineering().appendHolderSlice(&linkStateTE{obj: item})
			}
		}
		for _, item := range obj.TrafficEngineering().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Authentication != nil {

		obj.Authentication().validateObj(vObj, set_default)
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

	if obj.obj.LinkProtection != nil {

		obj.LinkProtection().validateObj(vObj, set_default)
	}

	if obj.obj.SrlgValues != nil {

		for _, item := range obj.obj.SrlgValues {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= IsisInterface.SrlgValues <= 16777215 but Got %d", item))
			}

		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface IsisInterface")
	}
}

func (obj *isisInterface) setDefault() {
	if obj.obj.Metric == nil {
		obj.SetMetric(10)
	}
	if obj.obj.NetworkType == nil {
		obj.SetNetworkType(IsisInterfaceNetworkType.BROADCAST)

	}
	if obj.obj.LevelType == nil {
		obj.SetLevelType(IsisInterfaceLevelType.LEVEL_2)

	}

}
