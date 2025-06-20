package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowGtpExtension *****
type flowGtpExtension struct {
	validation
	obj                       *otg.FlowGtpExtension
	marshaller                marshalFlowGtpExtension
	unMarshaller              unMarshalFlowGtpExtension
	extensionLengthHolder     PatternFlowGtpExtensionExtensionLength
	contentsHolder            PatternFlowGtpExtensionContents
	nextExtensionHeaderHolder PatternFlowGtpExtensionNextExtensionHeader
}

func NewFlowGtpExtension() FlowGtpExtension {
	obj := flowGtpExtension{obj: &otg.FlowGtpExtension{}}
	obj.setDefault()
	return &obj
}

func (obj *flowGtpExtension) msg() *otg.FlowGtpExtension {
	return obj.obj
}

func (obj *flowGtpExtension) setMsg(msg *otg.FlowGtpExtension) FlowGtpExtension {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowGtpExtension struct {
	obj *flowGtpExtension
}

type marshalFlowGtpExtension interface {
	// ToProto marshals FlowGtpExtension to protobuf object *otg.FlowGtpExtension
	ToProto() (*otg.FlowGtpExtension, error)
	// ToPbText marshals FlowGtpExtension to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowGtpExtension to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowGtpExtension to JSON text
	ToJson() (string, error)
}

type unMarshalflowGtpExtension struct {
	obj *flowGtpExtension
}

type unMarshalFlowGtpExtension interface {
	// FromProto unmarshals FlowGtpExtension from protobuf object *otg.FlowGtpExtension
	FromProto(msg *otg.FlowGtpExtension) (FlowGtpExtension, error)
	// FromPbText unmarshals FlowGtpExtension from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowGtpExtension from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowGtpExtension from JSON text
	FromJson(value string) error
}

func (obj *flowGtpExtension) Marshal() marshalFlowGtpExtension {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowGtpExtension{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowGtpExtension) Unmarshal() unMarshalFlowGtpExtension {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowGtpExtension{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowGtpExtension) ToProto() (*otg.FlowGtpExtension, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowGtpExtension) FromProto(msg *otg.FlowGtpExtension) (FlowGtpExtension, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowGtpExtension) ToPbText() (string, error) {
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

func (m *unMarshalflowGtpExtension) FromPbText(value string) error {
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

func (m *marshalflowGtpExtension) ToYaml() (string, error) {
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

func (m *unMarshalflowGtpExtension) FromYaml(value string) error {
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

func (m *marshalflowGtpExtension) ToJson() (string, error) {
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

func (m *unMarshalflowGtpExtension) FromJson(value string) error {
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

func (obj *flowGtpExtension) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowGtpExtension) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowGtpExtension) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowGtpExtension) Clone() (FlowGtpExtension, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowGtpExtension()
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

func (obj *flowGtpExtension) setNil() {
	obj.extensionLengthHolder = nil
	obj.contentsHolder = nil
	obj.nextExtensionHeaderHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowGtpExtension is description is TBD
type FlowGtpExtension interface {
	Validation
	// msg marshals FlowGtpExtension to protobuf object *otg.FlowGtpExtension
	// and doesn't set defaults
	msg() *otg.FlowGtpExtension
	// setMsg unmarshals FlowGtpExtension from protobuf object *otg.FlowGtpExtension
	// and doesn't set defaults
	setMsg(*otg.FlowGtpExtension) FlowGtpExtension
	// provides marshal interface
	Marshal() marshalFlowGtpExtension
	// provides unmarshal interface
	Unmarshal() unMarshalFlowGtpExtension
	// validate validates FlowGtpExtension
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowGtpExtension, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ExtensionLength returns PatternFlowGtpExtensionExtensionLength, set in FlowGtpExtension.
	// PatternFlowGtpExtensionExtensionLength is this field states the length of this extension header,  including the length, the contents, and the next extension header field, in 4-octet units, so the length of the extension must  always be a multiple of 4.
	ExtensionLength() PatternFlowGtpExtensionExtensionLength
	// SetExtensionLength assigns PatternFlowGtpExtensionExtensionLength provided by user to FlowGtpExtension.
	// PatternFlowGtpExtensionExtensionLength is this field states the length of this extension header,  including the length, the contents, and the next extension header field, in 4-octet units, so the length of the extension must  always be a multiple of 4.
	SetExtensionLength(value PatternFlowGtpExtensionExtensionLength) FlowGtpExtension
	// HasExtensionLength checks if ExtensionLength has been set in FlowGtpExtension
	HasExtensionLength() bool
	// Contents returns PatternFlowGtpExtensionContents, set in FlowGtpExtension.
	// PatternFlowGtpExtensionContents is the extension header contents
	Contents() PatternFlowGtpExtensionContents
	// SetContents assigns PatternFlowGtpExtensionContents provided by user to FlowGtpExtension.
	// PatternFlowGtpExtensionContents is the extension header contents
	SetContents(value PatternFlowGtpExtensionContents) FlowGtpExtension
	// HasContents checks if Contents has been set in FlowGtpExtension
	HasContents() bool
	// NextExtensionHeader returns PatternFlowGtpExtensionNextExtensionHeader, set in FlowGtpExtension.
	// PatternFlowGtpExtensionNextExtensionHeader is it states the type of the next extension, or 0 if no next  extension exists.  This permits chaining several next extension headers.
	NextExtensionHeader() PatternFlowGtpExtensionNextExtensionHeader
	// SetNextExtensionHeader assigns PatternFlowGtpExtensionNextExtensionHeader provided by user to FlowGtpExtension.
	// PatternFlowGtpExtensionNextExtensionHeader is it states the type of the next extension, or 0 if no next  extension exists.  This permits chaining several next extension headers.
	SetNextExtensionHeader(value PatternFlowGtpExtensionNextExtensionHeader) FlowGtpExtension
	// HasNextExtensionHeader checks if NextExtensionHeader has been set in FlowGtpExtension
	HasNextExtensionHeader() bool
	setNil()
}

// description is TBD
// ExtensionLength returns a PatternFlowGtpExtensionExtensionLength
func (obj *flowGtpExtension) ExtensionLength() PatternFlowGtpExtensionExtensionLength {
	if obj.obj.ExtensionLength == nil {
		obj.obj.ExtensionLength = NewPatternFlowGtpExtensionExtensionLength().msg()
	}
	if obj.extensionLengthHolder == nil {
		obj.extensionLengthHolder = &patternFlowGtpExtensionExtensionLength{obj: obj.obj.ExtensionLength}
	}
	return obj.extensionLengthHolder
}

// description is TBD
// ExtensionLength returns a PatternFlowGtpExtensionExtensionLength
func (obj *flowGtpExtension) HasExtensionLength() bool {
	return obj.obj.ExtensionLength != nil
}

// description is TBD
// SetExtensionLength sets the PatternFlowGtpExtensionExtensionLength value in the FlowGtpExtension object
func (obj *flowGtpExtension) SetExtensionLength(value PatternFlowGtpExtensionExtensionLength) FlowGtpExtension {

	obj.extensionLengthHolder = nil
	obj.obj.ExtensionLength = value.msg()

	return obj
}

// description is TBD
// Contents returns a PatternFlowGtpExtensionContents
func (obj *flowGtpExtension) Contents() PatternFlowGtpExtensionContents {
	if obj.obj.Contents == nil {
		obj.obj.Contents = NewPatternFlowGtpExtensionContents().msg()
	}
	if obj.contentsHolder == nil {
		obj.contentsHolder = &patternFlowGtpExtensionContents{obj: obj.obj.Contents}
	}
	return obj.contentsHolder
}

// description is TBD
// Contents returns a PatternFlowGtpExtensionContents
func (obj *flowGtpExtension) HasContents() bool {
	return obj.obj.Contents != nil
}

// description is TBD
// SetContents sets the PatternFlowGtpExtensionContents value in the FlowGtpExtension object
func (obj *flowGtpExtension) SetContents(value PatternFlowGtpExtensionContents) FlowGtpExtension {

	obj.contentsHolder = nil
	obj.obj.Contents = value.msg()

	return obj
}

// description is TBD
// NextExtensionHeader returns a PatternFlowGtpExtensionNextExtensionHeader
func (obj *flowGtpExtension) NextExtensionHeader() PatternFlowGtpExtensionNextExtensionHeader {
	if obj.obj.NextExtensionHeader == nil {
		obj.obj.NextExtensionHeader = NewPatternFlowGtpExtensionNextExtensionHeader().msg()
	}
	if obj.nextExtensionHeaderHolder == nil {
		obj.nextExtensionHeaderHolder = &patternFlowGtpExtensionNextExtensionHeader{obj: obj.obj.NextExtensionHeader}
	}
	return obj.nextExtensionHeaderHolder
}

// description is TBD
// NextExtensionHeader returns a PatternFlowGtpExtensionNextExtensionHeader
func (obj *flowGtpExtension) HasNextExtensionHeader() bool {
	return obj.obj.NextExtensionHeader != nil
}

// description is TBD
// SetNextExtensionHeader sets the PatternFlowGtpExtensionNextExtensionHeader value in the FlowGtpExtension object
func (obj *flowGtpExtension) SetNextExtensionHeader(value PatternFlowGtpExtensionNextExtensionHeader) FlowGtpExtension {

	obj.nextExtensionHeaderHolder = nil
	obj.obj.NextExtensionHeader = value.msg()

	return obj
}

func (obj *flowGtpExtension) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ExtensionLength != nil {

		obj.ExtensionLength().validateObj(vObj, set_default)
	}

	if obj.obj.Contents != nil {

		obj.Contents().validateObj(vObj, set_default)
	}

	if obj.obj.NextExtensionHeader != nil {

		obj.NextExtensionHeader().validateObj(vObj, set_default)
	}

}

func (obj *flowGtpExtension) setDefault() {

}
