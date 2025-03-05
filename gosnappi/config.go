package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Config *****
type config struct {
	validation
	obj            *otg.Config
	marshaller     marshalConfig
	unMarshaller   unMarshalConfig
	portsHolder    ConfigPortIter
	lagsHolder     ConfigLagIter
	layer1Holder   ConfigLayer1Iter
	capturesHolder ConfigCaptureIter
	devicesHolder  ConfigDeviceIter
	flowsHolder    ConfigFlowIter
	eventsHolder   Event
	optionsHolder  ConfigOptions
	lldpHolder     ConfigLldpIter
}

func NewConfig() Config {
	obj := config{obj: &otg.Config{}}
	obj.setDefault()
	return &obj
}

func (obj *config) msg() *otg.Config {
	return obj.obj
}

func (obj *config) setMsg(msg *otg.Config) Config {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfig struct {
	obj *config
}

type marshalConfig interface {
	// ToProto marshals Config to protobuf object *otg.Config
	ToProto() (*otg.Config, error)
	// ToPbText marshals Config to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Config to YAML text
	ToYaml() (string, error)
	// ToJson marshals Config to JSON text
	ToJson() (string, error)
}

type unMarshalconfig struct {
	obj *config
}

type unMarshalConfig interface {
	// FromProto unmarshals Config from protobuf object *otg.Config
	FromProto(msg *otg.Config) (Config, error)
	// FromPbText unmarshals Config from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Config from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Config from JSON text
	FromJson(value string) error
}

func (obj *config) Marshal() marshalConfig {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfig{obj: obj}
	}
	return obj.marshaller
}

func (obj *config) Unmarshal() unMarshalConfig {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfig{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfig) ToProto() (*otg.Config, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfig) FromProto(msg *otg.Config) (Config, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfig) ToPbText() (string, error) {
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

func (m *unMarshalconfig) FromPbText(value string) error {
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

func (m *marshalconfig) ToYaml() (string, error) {
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

func (m *unMarshalconfig) FromYaml(value string) error {
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

func (m *marshalconfig) ToJson() (string, error) {
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

func (m *unMarshalconfig) FromJson(value string) error {
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

func (obj *config) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *config) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *config) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *config) Clone() (Config, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfig()
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

func (obj *config) setNil() {
	obj.portsHolder = nil
	obj.lagsHolder = nil
	obj.layer1Holder = nil
	obj.capturesHolder = nil
	obj.devicesHolder = nil
	obj.flowsHolder = nil
	obj.eventsHolder = nil
	obj.optionsHolder = nil
	obj.lldpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Config is a container for all models that are part of the configuration.
type Config interface {
	Validation
	// msg marshals Config to protobuf object *otg.Config
	// and doesn't set defaults
	msg() *otg.Config
	// setMsg unmarshals Config from protobuf object *otg.Config
	// and doesn't set defaults
	setMsg(*otg.Config) Config
	// provides marshal interface
	Marshal() marshalConfig
	// provides unmarshal interface
	Unmarshal() unMarshalConfig
	// validate validates Config
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Config, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ports returns ConfigPortIterIter, set in Config
	Ports() ConfigPortIter
	// Lags returns ConfigLagIterIter, set in Config
	Lags() ConfigLagIter
	// Layer1 returns ConfigLayer1IterIter, set in Config
	Layer1() ConfigLayer1Iter
	// Captures returns ConfigCaptureIterIter, set in Config
	Captures() ConfigCaptureIter
	// Devices returns ConfigDeviceIterIter, set in Config
	Devices() ConfigDeviceIter
	// Flows returns ConfigFlowIterIter, set in Config
	Flows() ConfigFlowIter
	// Events returns Event, set in Config.
	// Event is under Review: Event configuration is currently under review for pending exploration on use cases.
	//
	// The optional container for event configuration.
	// Both cp_events.enable and dp_events.enable must be explicitly set to true to get
	// control_plane_data_plane_convergence_us metric values for convergence metrics.
	Events() Event
	// SetEvents assigns Event provided by user to Config.
	// Event is under Review: Event configuration is currently under review for pending exploration on use cases.
	//
	// The optional container for event configuration.
	// Both cp_events.enable and dp_events.enable must be explicitly set to true to get
	// control_plane_data_plane_convergence_us metric values for convergence metrics.
	SetEvents(value Event) Config
	// HasEvents checks if Events has been set in Config
	HasEvents() bool
	// Options returns ConfigOptions, set in Config.
	// ConfigOptions is global configuration options.
	Options() ConfigOptions
	// SetOptions assigns ConfigOptions provided by user to Config.
	// ConfigOptions is global configuration options.
	SetOptions(value ConfigOptions) Config
	// HasOptions checks if Options has been set in Config
	HasOptions() bool
	// Lldp returns ConfigLldpIterIter, set in Config
	Lldp() ConfigLldpIter
	setNil()
}

// The ports that will be configured on the traffic generator.
// Ports returns a []Port
func (obj *config) Ports() ConfigPortIter {
	if len(obj.obj.Ports) == 0 {
		obj.obj.Ports = []*otg.Port{}
	}
	if obj.portsHolder == nil {
		obj.portsHolder = newConfigPortIter(&obj.obj.Ports).setMsg(obj)
	}
	return obj.portsHolder
}

type configPortIter struct {
	obj       *config
	portSlice []Port
	fieldPtr  *[]*otg.Port
}

func newConfigPortIter(ptr *[]*otg.Port) ConfigPortIter {
	return &configPortIter{fieldPtr: ptr}
}

type ConfigPortIter interface {
	setMsg(*config) ConfigPortIter
	Items() []Port
	Add() Port
	Append(items ...Port) ConfigPortIter
	Set(index int, newObj Port) ConfigPortIter
	Clear() ConfigPortIter
	clearHolderSlice() ConfigPortIter
	appendHolderSlice(item Port) ConfigPortIter
}

func (obj *configPortIter) setMsg(msg *config) ConfigPortIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&port{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configPortIter) Items() []Port {
	return obj.portSlice
}

func (obj *configPortIter) Add() Port {
	newObj := &otg.Port{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &port{obj: newObj}
	newLibObj.setDefault()
	obj.portSlice = append(obj.portSlice, newLibObj)
	return newLibObj
}

func (obj *configPortIter) Append(items ...Port) ConfigPortIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.portSlice = append(obj.portSlice, item)
	}
	return obj
}

func (obj *configPortIter) Set(index int, newObj Port) ConfigPortIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.portSlice[index] = newObj
	return obj
}
func (obj *configPortIter) Clear() ConfigPortIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Port{}
		obj.portSlice = []Port{}
	}
	return obj
}
func (obj *configPortIter) clearHolderSlice() ConfigPortIter {
	if len(obj.portSlice) > 0 {
		obj.portSlice = []Port{}
	}
	return obj
}
func (obj *configPortIter) appendHolderSlice(item Port) ConfigPortIter {
	obj.portSlice = append(obj.portSlice, item)
	return obj
}

// The LAGs that will be configured on the traffic generator.
// Lags returns a []Lag
func (obj *config) Lags() ConfigLagIter {
	if len(obj.obj.Lags) == 0 {
		obj.obj.Lags = []*otg.Lag{}
	}
	if obj.lagsHolder == nil {
		obj.lagsHolder = newConfigLagIter(&obj.obj.Lags).setMsg(obj)
	}
	return obj.lagsHolder
}

type configLagIter struct {
	obj      *config
	lagSlice []Lag
	fieldPtr *[]*otg.Lag
}

func newConfigLagIter(ptr *[]*otg.Lag) ConfigLagIter {
	return &configLagIter{fieldPtr: ptr}
}

type ConfigLagIter interface {
	setMsg(*config) ConfigLagIter
	Items() []Lag
	Add() Lag
	Append(items ...Lag) ConfigLagIter
	Set(index int, newObj Lag) ConfigLagIter
	Clear() ConfigLagIter
	clearHolderSlice() ConfigLagIter
	appendHolderSlice(item Lag) ConfigLagIter
}

func (obj *configLagIter) setMsg(msg *config) ConfigLagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configLagIter) Items() []Lag {
	return obj.lagSlice
}

func (obj *configLagIter) Add() Lag {
	newObj := &otg.Lag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lag{obj: newObj}
	newLibObj.setDefault()
	obj.lagSlice = append(obj.lagSlice, newLibObj)
	return newLibObj
}

func (obj *configLagIter) Append(items ...Lag) ConfigLagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lagSlice = append(obj.lagSlice, item)
	}
	return obj
}

func (obj *configLagIter) Set(index int, newObj Lag) ConfigLagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lagSlice[index] = newObj
	return obj
}
func (obj *configLagIter) Clear() ConfigLagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Lag{}
		obj.lagSlice = []Lag{}
	}
	return obj
}
func (obj *configLagIter) clearHolderSlice() ConfigLagIter {
	if len(obj.lagSlice) > 0 {
		obj.lagSlice = []Lag{}
	}
	return obj
}
func (obj *configLagIter) appendHolderSlice(item Lag) ConfigLagIter {
	obj.lagSlice = append(obj.lagSlice, item)
	return obj
}

// The layer1 settings that will be configured on the traffic generator.
// Since layer1 settings usually vary across variety of test ports, these
// most likely won't be portable.
// Layer1 returns a []Layer1
func (obj *config) Layer1() ConfigLayer1Iter {
	if len(obj.obj.Layer1) == 0 {
		obj.obj.Layer1 = []*otg.Layer1{}
	}
	if obj.layer1Holder == nil {
		obj.layer1Holder = newConfigLayer1Iter(&obj.obj.Layer1).setMsg(obj)
	}
	return obj.layer1Holder
}

type configLayer1Iter struct {
	obj         *config
	layer1Slice []Layer1
	fieldPtr    *[]*otg.Layer1
}

func newConfigLayer1Iter(ptr *[]*otg.Layer1) ConfigLayer1Iter {
	return &configLayer1Iter{fieldPtr: ptr}
}

type ConfigLayer1Iter interface {
	setMsg(*config) ConfigLayer1Iter
	Items() []Layer1
	Add() Layer1
	Append(items ...Layer1) ConfigLayer1Iter
	Set(index int, newObj Layer1) ConfigLayer1Iter
	Clear() ConfigLayer1Iter
	clearHolderSlice() ConfigLayer1Iter
	appendHolderSlice(item Layer1) ConfigLayer1Iter
}

func (obj *configLayer1Iter) setMsg(msg *config) ConfigLayer1Iter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&layer1{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configLayer1Iter) Items() []Layer1 {
	return obj.layer1Slice
}

func (obj *configLayer1Iter) Add() Layer1 {
	newObj := &otg.Layer1{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &layer1{obj: newObj}
	newLibObj.setDefault()
	obj.layer1Slice = append(obj.layer1Slice, newLibObj)
	return newLibObj
}

func (obj *configLayer1Iter) Append(items ...Layer1) ConfigLayer1Iter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.layer1Slice = append(obj.layer1Slice, item)
	}
	return obj
}

func (obj *configLayer1Iter) Set(index int, newObj Layer1) ConfigLayer1Iter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.layer1Slice[index] = newObj
	return obj
}
func (obj *configLayer1Iter) Clear() ConfigLayer1Iter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Layer1{}
		obj.layer1Slice = []Layer1{}
	}
	return obj
}
func (obj *configLayer1Iter) clearHolderSlice() ConfigLayer1Iter {
	if len(obj.layer1Slice) > 0 {
		obj.layer1Slice = []Layer1{}
	}
	return obj
}
func (obj *configLayer1Iter) appendHolderSlice(item Layer1) ConfigLayer1Iter {
	obj.layer1Slice = append(obj.layer1Slice, item)
	return obj
}

// The capture settings that will be configured on the traffic generator.
// Captures returns a []Capture
func (obj *config) Captures() ConfigCaptureIter {
	if len(obj.obj.Captures) == 0 {
		obj.obj.Captures = []*otg.Capture{}
	}
	if obj.capturesHolder == nil {
		obj.capturesHolder = newConfigCaptureIter(&obj.obj.Captures).setMsg(obj)
	}
	return obj.capturesHolder
}

type configCaptureIter struct {
	obj          *config
	captureSlice []Capture
	fieldPtr     *[]*otg.Capture
}

func newConfigCaptureIter(ptr *[]*otg.Capture) ConfigCaptureIter {
	return &configCaptureIter{fieldPtr: ptr}
}

type ConfigCaptureIter interface {
	setMsg(*config) ConfigCaptureIter
	Items() []Capture
	Add() Capture
	Append(items ...Capture) ConfigCaptureIter
	Set(index int, newObj Capture) ConfigCaptureIter
	Clear() ConfigCaptureIter
	clearHolderSlice() ConfigCaptureIter
	appendHolderSlice(item Capture) ConfigCaptureIter
}

func (obj *configCaptureIter) setMsg(msg *config) ConfigCaptureIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&capture{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configCaptureIter) Items() []Capture {
	return obj.captureSlice
}

func (obj *configCaptureIter) Add() Capture {
	newObj := &otg.Capture{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &capture{obj: newObj}
	newLibObj.setDefault()
	obj.captureSlice = append(obj.captureSlice, newLibObj)
	return newLibObj
}

func (obj *configCaptureIter) Append(items ...Capture) ConfigCaptureIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.captureSlice = append(obj.captureSlice, item)
	}
	return obj
}

func (obj *configCaptureIter) Set(index int, newObj Capture) ConfigCaptureIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.captureSlice[index] = newObj
	return obj
}
func (obj *configCaptureIter) Clear() ConfigCaptureIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Capture{}
		obj.captureSlice = []Capture{}
	}
	return obj
}
func (obj *configCaptureIter) clearHolderSlice() ConfigCaptureIter {
	if len(obj.captureSlice) > 0 {
		obj.captureSlice = []Capture{}
	}
	return obj
}
func (obj *configCaptureIter) appendHolderSlice(item Capture) ConfigCaptureIter {
	obj.captureSlice = append(obj.captureSlice, item)
	return obj
}

// The emulated devices that will be configured on the traffic generator.
// Each device contains configurations for network interfaces and
// protocols running on top of those interfaces.
// Devices returns a []Device
func (obj *config) Devices() ConfigDeviceIter {
	if len(obj.obj.Devices) == 0 {
		obj.obj.Devices = []*otg.Device{}
	}
	if obj.devicesHolder == nil {
		obj.devicesHolder = newConfigDeviceIter(&obj.obj.Devices).setMsg(obj)
	}
	return obj.devicesHolder
}

type configDeviceIter struct {
	obj         *config
	deviceSlice []Device
	fieldPtr    *[]*otg.Device
}

func newConfigDeviceIter(ptr *[]*otg.Device) ConfigDeviceIter {
	return &configDeviceIter{fieldPtr: ptr}
}

type ConfigDeviceIter interface {
	setMsg(*config) ConfigDeviceIter
	Items() []Device
	Add() Device
	Append(items ...Device) ConfigDeviceIter
	Set(index int, newObj Device) ConfigDeviceIter
	Clear() ConfigDeviceIter
	clearHolderSlice() ConfigDeviceIter
	appendHolderSlice(item Device) ConfigDeviceIter
}

func (obj *configDeviceIter) setMsg(msg *config) ConfigDeviceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&device{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configDeviceIter) Items() []Device {
	return obj.deviceSlice
}

func (obj *configDeviceIter) Add() Device {
	newObj := &otg.Device{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &device{obj: newObj}
	newLibObj.setDefault()
	obj.deviceSlice = append(obj.deviceSlice, newLibObj)
	return newLibObj
}

func (obj *configDeviceIter) Append(items ...Device) ConfigDeviceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceSlice = append(obj.deviceSlice, item)
	}
	return obj
}

func (obj *configDeviceIter) Set(index int, newObj Device) ConfigDeviceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceSlice[index] = newObj
	return obj
}
func (obj *configDeviceIter) Clear() ConfigDeviceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Device{}
		obj.deviceSlice = []Device{}
	}
	return obj
}
func (obj *configDeviceIter) clearHolderSlice() ConfigDeviceIter {
	if len(obj.deviceSlice) > 0 {
		obj.deviceSlice = []Device{}
	}
	return obj
}
func (obj *configDeviceIter) appendHolderSlice(item Device) ConfigDeviceIter {
	obj.deviceSlice = append(obj.deviceSlice, item)
	return obj
}

// The flows that will be configured on the traffic generator.
// Flows returns a []Flow
func (obj *config) Flows() ConfigFlowIter {
	if len(obj.obj.Flows) == 0 {
		obj.obj.Flows = []*otg.Flow{}
	}
	if obj.flowsHolder == nil {
		obj.flowsHolder = newConfigFlowIter(&obj.obj.Flows).setMsg(obj)
	}
	return obj.flowsHolder
}

type configFlowIter struct {
	obj       *config
	flowSlice []Flow
	fieldPtr  *[]*otg.Flow
}

func newConfigFlowIter(ptr *[]*otg.Flow) ConfigFlowIter {
	return &configFlowIter{fieldPtr: ptr}
}

type ConfigFlowIter interface {
	setMsg(*config) ConfigFlowIter
	Items() []Flow
	Add() Flow
	Append(items ...Flow) ConfigFlowIter
	Set(index int, newObj Flow) ConfigFlowIter
	Clear() ConfigFlowIter
	clearHolderSlice() ConfigFlowIter
	appendHolderSlice(item Flow) ConfigFlowIter
}

func (obj *configFlowIter) setMsg(msg *config) ConfigFlowIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flow{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configFlowIter) Items() []Flow {
	return obj.flowSlice
}

func (obj *configFlowIter) Add() Flow {
	newObj := &otg.Flow{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flow{obj: newObj}
	newLibObj.setDefault()
	obj.flowSlice = append(obj.flowSlice, newLibObj)
	return newLibObj
}

func (obj *configFlowIter) Append(items ...Flow) ConfigFlowIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowSlice = append(obj.flowSlice, item)
	}
	return obj
}

func (obj *configFlowIter) Set(index int, newObj Flow) ConfigFlowIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowSlice[index] = newObj
	return obj
}
func (obj *configFlowIter) Clear() ConfigFlowIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Flow{}
		obj.flowSlice = []Flow{}
	}
	return obj
}
func (obj *configFlowIter) clearHolderSlice() ConfigFlowIter {
	if len(obj.flowSlice) > 0 {
		obj.flowSlice = []Flow{}
	}
	return obj
}
func (obj *configFlowIter) appendHolderSlice(item Flow) ConfigFlowIter {
	obj.flowSlice = append(obj.flowSlice, item)
	return obj
}

// description is TBD
// Events returns a Event
func (obj *config) Events() Event {
	if obj.obj.Events == nil {
		obj.obj.Events = NewEvent().msg()
	}
	if obj.eventsHolder == nil {
		obj.eventsHolder = &event{obj: obj.obj.Events}
	}
	return obj.eventsHolder
}

// description is TBD
// Events returns a Event
func (obj *config) HasEvents() bool {
	return obj.obj.Events != nil
}

// description is TBD
// SetEvents sets the Event value in the Config object
func (obj *config) SetEvents(value Event) Config {

	obj.eventsHolder = nil
	obj.obj.Events = value.msg()

	return obj
}

// description is TBD
// Options returns a ConfigOptions
func (obj *config) Options() ConfigOptions {
	if obj.obj.Options == nil {
		obj.obj.Options = NewConfigOptions().msg()
	}
	if obj.optionsHolder == nil {
		obj.optionsHolder = &configOptions{obj: obj.obj.Options}
	}
	return obj.optionsHolder
}

// description is TBD
// Options returns a ConfigOptions
func (obj *config) HasOptions() bool {
	return obj.obj.Options != nil
}

// description is TBD
// SetOptions sets the ConfigOptions value in the Config object
func (obj *config) SetOptions(value ConfigOptions) Config {

	obj.optionsHolder = nil
	obj.obj.Options = value.msg()

	return obj
}

// LLDP protocol that will be configured on traffic generator.
// Lldp returns a []Lldp
func (obj *config) Lldp() ConfigLldpIter {
	if len(obj.obj.Lldp) == 0 {
		obj.obj.Lldp = []*otg.Lldp{}
	}
	if obj.lldpHolder == nil {
		obj.lldpHolder = newConfigLldpIter(&obj.obj.Lldp).setMsg(obj)
	}
	return obj.lldpHolder
}

type configLldpIter struct {
	obj       *config
	lldpSlice []Lldp
	fieldPtr  *[]*otg.Lldp
}

func newConfigLldpIter(ptr *[]*otg.Lldp) ConfigLldpIter {
	return &configLldpIter{fieldPtr: ptr}
}

type ConfigLldpIter interface {
	setMsg(*config) ConfigLldpIter
	Items() []Lldp
	Add() Lldp
	Append(items ...Lldp) ConfigLldpIter
	Set(index int, newObj Lldp) ConfigLldpIter
	Clear() ConfigLldpIter
	clearHolderSlice() ConfigLldpIter
	appendHolderSlice(item Lldp) ConfigLldpIter
}

func (obj *configLldpIter) setMsg(msg *config) ConfigLldpIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lldp{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configLldpIter) Items() []Lldp {
	return obj.lldpSlice
}

func (obj *configLldpIter) Add() Lldp {
	newObj := &otg.Lldp{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lldp{obj: newObj}
	newLibObj.setDefault()
	obj.lldpSlice = append(obj.lldpSlice, newLibObj)
	return newLibObj
}

func (obj *configLldpIter) Append(items ...Lldp) ConfigLldpIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lldpSlice = append(obj.lldpSlice, item)
	}
	return obj
}

func (obj *configLldpIter) Set(index int, newObj Lldp) ConfigLldpIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lldpSlice[index] = newObj
	return obj
}
func (obj *configLldpIter) Clear() ConfigLldpIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Lldp{}
		obj.lldpSlice = []Lldp{}
	}
	return obj
}
func (obj *configLldpIter) clearHolderSlice() ConfigLldpIter {
	if len(obj.lldpSlice) > 0 {
		obj.lldpSlice = []Lldp{}
	}
	return obj
}
func (obj *configLldpIter) appendHolderSlice(item Lldp) ConfigLldpIter {
	obj.lldpSlice = append(obj.lldpSlice, item)
	return obj
}

func (obj *config) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ports) != 0 {

		if set_default {
			obj.Ports().clearHolderSlice()
			for _, item := range obj.obj.Ports {
				obj.Ports().appendHolderSlice(&port{obj: item})
			}
		}
		for _, item := range obj.Ports().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Lags) != 0 {

		if set_default {
			obj.Lags().clearHolderSlice()
			for _, item := range obj.obj.Lags {
				obj.Lags().appendHolderSlice(&lag{obj: item})
			}
		}
		for _, item := range obj.Lags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Layer1) != 0 {

		if set_default {
			obj.Layer1().clearHolderSlice()
			for _, item := range obj.obj.Layer1 {
				obj.Layer1().appendHolderSlice(&layer1{obj: item})
			}
		}
		for _, item := range obj.Layer1().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Captures) != 0 {

		if set_default {
			obj.Captures().clearHolderSlice()
			for _, item := range obj.obj.Captures {
				obj.Captures().appendHolderSlice(&capture{obj: item})
			}
		}
		for _, item := range obj.Captures().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Devices) != 0 {

		if set_default {
			obj.Devices().clearHolderSlice()
			for _, item := range obj.obj.Devices {
				obj.Devices().appendHolderSlice(&device{obj: item})
			}
		}
		for _, item := range obj.Devices().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Flows) != 0 {

		if set_default {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				obj.Flows().appendHolderSlice(&flow{obj: item})
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Events != nil {

		obj.Events().validateObj(vObj, set_default)
	}

	if obj.obj.Options != nil {

		obj.Options().validateObj(vObj, set_default)
	}

	if len(obj.obj.Lldp) != 0 {

		if set_default {
			obj.Lldp().clearHolderSlice()
			for _, item := range obj.obj.Lldp {
				obj.Lldp().appendHolderSlice(&lldp{obj: item})
			}
		}
		for _, item := range obj.Lldp().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *config) setDefault() {

}
