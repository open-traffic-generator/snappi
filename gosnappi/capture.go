package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Capture *****
type capture struct {
	validation
	obj           *otg.Capture
	marshaller    marshalCapture
	unMarshaller  unMarshalCapture
	filtersHolder CaptureCaptureFilterIter
}

func NewCapture() Capture {
	obj := capture{obj: &otg.Capture{}}
	obj.setDefault()
	return &obj
}

func (obj *capture) msg() *otg.Capture {
	return obj.obj
}

func (obj *capture) setMsg(msg *otg.Capture) Capture {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcapture struct {
	obj *capture
}

type marshalCapture interface {
	// ToProto marshals Capture to protobuf object *otg.Capture
	ToProto() (*otg.Capture, error)
	// ToPbText marshals Capture to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Capture to YAML text
	ToYaml() (string, error)
	// ToJson marshals Capture to JSON text
	ToJson() (string, error)
}

type unMarshalcapture struct {
	obj *capture
}

type unMarshalCapture interface {
	// FromProto unmarshals Capture from protobuf object *otg.Capture
	FromProto(msg *otg.Capture) (Capture, error)
	// FromPbText unmarshals Capture from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Capture from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Capture from JSON text
	FromJson(value string) error
}

func (obj *capture) Marshal() marshalCapture {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcapture{obj: obj}
	}
	return obj.marshaller
}

func (obj *capture) Unmarshal() unMarshalCapture {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcapture{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcapture) ToProto() (*otg.Capture, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcapture) FromProto(msg *otg.Capture) (Capture, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcapture) ToPbText() (string, error) {
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

func (m *unMarshalcapture) FromPbText(value string) error {
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

func (m *marshalcapture) ToYaml() (string, error) {
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

func (m *unMarshalcapture) FromYaml(value string) error {
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

func (m *marshalcapture) ToJson() (string, error) {
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

func (m *unMarshalcapture) FromJson(value string) error {
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

func (obj *capture) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *capture) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *capture) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *capture) Clone() (Capture, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCapture()
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

func (obj *capture) setNil() {
	obj.filtersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Capture is configuration for capture settings.
type Capture interface {
	Validation
	// msg marshals Capture to protobuf object *otg.Capture
	// and doesn't set defaults
	msg() *otg.Capture
	// setMsg unmarshals Capture from protobuf object *otg.Capture
	// and doesn't set defaults
	setMsg(*otg.Capture) Capture
	// provides marshal interface
	Marshal() marshalCapture
	// provides unmarshal interface
	Unmarshal() unMarshalCapture
	// validate validates Capture
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Capture, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortNames returns []string, set in Capture.
	PortNames() []string
	// SetPortNames assigns []string provided by user to Capture
	SetPortNames(value []string) Capture
	// Filters returns CaptureCaptureFilterIterIter, set in Capture
	Filters() CaptureCaptureFilterIter
	// Overwrite returns bool, set in Capture.
	Overwrite() bool
	// SetOverwrite assigns bool provided by user to Capture
	SetOverwrite(value bool) Capture
	// HasOverwrite checks if Overwrite has been set in Capture
	HasOverwrite() bool
	// PacketSize returns uint32, set in Capture.
	PacketSize() uint32
	// SetPacketSize assigns uint32 provided by user to Capture
	SetPacketSize(value uint32) Capture
	// HasPacketSize checks if PacketSize has been set in Capture
	HasPacketSize() bool
	// Format returns CaptureFormatEnum, set in Capture
	Format() CaptureFormatEnum
	// SetFormat assigns CaptureFormatEnum provided by user to Capture
	SetFormat(value CaptureFormatEnum) Capture
	// HasFormat checks if Format has been set in Capture
	HasFormat() bool
	// Name returns string, set in Capture.
	Name() string
	// SetName assigns string provided by user to Capture
	SetName(value string) Capture
	setNil()
}

// The unique names of ports that the capture settings will apply to. Port_names cannot be duplicated between capture objects.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortNames returns a []string
func (obj *capture) PortNames() []string {
	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	return obj.obj.PortNames
}

// The unique names of ports that the capture settings will apply to. Port_names cannot be duplicated between capture objects.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortNames sets the []string value in the Capture object
func (obj *capture) SetPortNames(value []string) Capture {

	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	obj.obj.PortNames = value

	return obj
}

// A list of filters to apply to the capturing ports. If no filters are specified then all packets will be captured. A capture can have multiple filters. The number of filters supported is determined by the implementation which can be retrieved using the capabilities API.
// When multiple filters are specified the capture implementation  must && (and) all the filters.
// Filters returns a []CaptureFilter
func (obj *capture) Filters() CaptureCaptureFilterIter {
	if len(obj.obj.Filters) == 0 {
		obj.obj.Filters = []*otg.CaptureFilter{}
	}
	if obj.filtersHolder == nil {
		obj.filtersHolder = newCaptureCaptureFilterIter(&obj.obj.Filters).setMsg(obj)
	}
	return obj.filtersHolder
}

type captureCaptureFilterIter struct {
	obj                *capture
	captureFilterSlice []CaptureFilter
	fieldPtr           *[]*otg.CaptureFilter
}

func newCaptureCaptureFilterIter(ptr *[]*otg.CaptureFilter) CaptureCaptureFilterIter {
	return &captureCaptureFilterIter{fieldPtr: ptr}
}

type CaptureCaptureFilterIter interface {
	setMsg(*capture) CaptureCaptureFilterIter
	Items() []CaptureFilter
	Add() CaptureFilter
	Append(items ...CaptureFilter) CaptureCaptureFilterIter
	Set(index int, newObj CaptureFilter) CaptureCaptureFilterIter
	Clear() CaptureCaptureFilterIter
	clearHolderSlice() CaptureCaptureFilterIter
	appendHolderSlice(item CaptureFilter) CaptureCaptureFilterIter
}

func (obj *captureCaptureFilterIter) setMsg(msg *capture) CaptureCaptureFilterIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&captureFilter{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *captureCaptureFilterIter) Items() []CaptureFilter {
	return obj.captureFilterSlice
}

func (obj *captureCaptureFilterIter) Add() CaptureFilter {
	newObj := &otg.CaptureFilter{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &captureFilter{obj: newObj}
	newLibObj.setDefault()
	obj.captureFilterSlice = append(obj.captureFilterSlice, newLibObj)
	return newLibObj
}

func (obj *captureCaptureFilterIter) Append(items ...CaptureFilter) CaptureCaptureFilterIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.captureFilterSlice = append(obj.captureFilterSlice, item)
	}
	return obj
}

func (obj *captureCaptureFilterIter) Set(index int, newObj CaptureFilter) CaptureCaptureFilterIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.captureFilterSlice[index] = newObj
	return obj
}
func (obj *captureCaptureFilterIter) Clear() CaptureCaptureFilterIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.CaptureFilter{}
		obj.captureFilterSlice = []CaptureFilter{}
	}
	return obj
}
func (obj *captureCaptureFilterIter) clearHolderSlice() CaptureCaptureFilterIter {
	if len(obj.captureFilterSlice) > 0 {
		obj.captureFilterSlice = []CaptureFilter{}
	}
	return obj
}
func (obj *captureCaptureFilterIter) appendHolderSlice(item CaptureFilter) CaptureCaptureFilterIter {
	obj.captureFilterSlice = append(obj.captureFilterSlice, item)
	return obj
}

// Overwrite the capture buffer.
// Overwrite returns a bool
func (obj *capture) Overwrite() bool {

	return *obj.obj.Overwrite

}

// Overwrite the capture buffer.
// Overwrite returns a bool
func (obj *capture) HasOverwrite() bool {
	return obj.obj.Overwrite != nil
}

// Overwrite the capture buffer.
// SetOverwrite sets the bool value in the Capture object
func (obj *capture) SetOverwrite(value bool) Capture {

	obj.obj.Overwrite = &value
	return obj
}

// The maximum size of each captured packet. If no value is specified or it is null then the entire packet will be captured.
// PacketSize returns a uint32
func (obj *capture) PacketSize() uint32 {

	return *obj.obj.PacketSize

}

// The maximum size of each captured packet. If no value is specified or it is null then the entire packet will be captured.
// PacketSize returns a uint32
func (obj *capture) HasPacketSize() bool {
	return obj.obj.PacketSize != nil
}

// The maximum size of each captured packet. If no value is specified or it is null then the entire packet will be captured.
// SetPacketSize sets the uint32 value in the Capture object
func (obj *capture) SetPacketSize(value uint32) Capture {

	obj.obj.PacketSize = &value
	return obj
}

type CaptureFormatEnum string

// Enum of Format on Capture
var CaptureFormat = struct {
	PCAP   CaptureFormatEnum
	PCAPNG CaptureFormatEnum
}{
	PCAP:   CaptureFormatEnum("pcap"),
	PCAPNG: CaptureFormatEnum("pcapng"),
}

func (obj *capture) Format() CaptureFormatEnum {
	return CaptureFormatEnum(obj.obj.Format.Enum().String())
}

// The format of the capture file.
// Format returns a string
func (obj *capture) HasFormat() bool {
	return obj.obj.Format != nil
}

func (obj *capture) SetFormat(value CaptureFormatEnum) Capture {
	intValue, ok := otg.Capture_Format_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on CaptureFormatEnum", string(value)))
		return obj
	}
	enumValue := otg.Capture_Format_Enum(intValue)
	obj.obj.Format = &enumValue

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *capture) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Capture object
func (obj *capture) SetName(value string) Capture {

	obj.obj.Name = &value
	return obj
}

func (obj *capture) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Filters) != 0 {

		if set_default {
			obj.Filters().clearHolderSlice()
			for _, item := range obj.obj.Filters {
				obj.Filters().appendHolderSlice(&captureFilter{obj: item})
			}
		}
		for _, item := range obj.Filters().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.PacketSize != nil {

		if *obj.obj.PacketSize > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Capture.PacketSize <= 65535 but Got %d", *obj.obj.PacketSize))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Capture")
	}
}

func (obj *capture) setDefault() {
	if obj.obj.Overwrite == nil {
		obj.SetOverwrite(true)
	}
	if obj.obj.Format == nil {
		obj.SetFormat(CaptureFormat.PCAP)

	}

}
