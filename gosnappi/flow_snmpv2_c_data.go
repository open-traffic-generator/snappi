package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSnmpv2CData *****
type flowSnmpv2CData struct {
	validation
	obj                  *otg.FlowSnmpv2CData
	marshaller           marshalFlowSnmpv2CData
	unMarshaller         unMarshalFlowSnmpv2CData
	getRequestHolder     FlowSnmpv2CPDU
	getNextRequestHolder FlowSnmpv2CPDU
	responseHolder       FlowSnmpv2CPDU
	setRequestHolder     FlowSnmpv2CPDU
	getBulkRequestHolder FlowSnmpv2CBulkPDU
	informRequestHolder  FlowSnmpv2CPDU
	snmpv2TrapHolder     FlowSnmpv2CPDU
	reportHolder         FlowSnmpv2CPDU
}

func NewFlowSnmpv2CData() FlowSnmpv2CData {
	obj := flowSnmpv2CData{obj: &otg.FlowSnmpv2CData{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSnmpv2CData) msg() *otg.FlowSnmpv2CData {
	return obj.obj
}

func (obj *flowSnmpv2CData) setMsg(msg *otg.FlowSnmpv2CData) FlowSnmpv2CData {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSnmpv2CData struct {
	obj *flowSnmpv2CData
}

type marshalFlowSnmpv2CData interface {
	// ToProto marshals FlowSnmpv2CData to protobuf object *otg.FlowSnmpv2CData
	ToProto() (*otg.FlowSnmpv2CData, error)
	// ToPbText marshals FlowSnmpv2CData to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSnmpv2CData to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSnmpv2CData to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowSnmpv2CData to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowSnmpv2CData struct {
	obj *flowSnmpv2CData
}

type unMarshalFlowSnmpv2CData interface {
	// FromProto unmarshals FlowSnmpv2CData from protobuf object *otg.FlowSnmpv2CData
	FromProto(msg *otg.FlowSnmpv2CData) (FlowSnmpv2CData, error)
	// FromPbText unmarshals FlowSnmpv2CData from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSnmpv2CData from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSnmpv2CData from JSON text
	FromJson(value string) error
}

func (obj *flowSnmpv2CData) Marshal() marshalFlowSnmpv2CData {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSnmpv2CData{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSnmpv2CData) Unmarshal() unMarshalFlowSnmpv2CData {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSnmpv2CData{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSnmpv2CData) ToProto() (*otg.FlowSnmpv2CData, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSnmpv2CData) FromProto(msg *otg.FlowSnmpv2CData) (FlowSnmpv2CData, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSnmpv2CData) ToPbText() (string, error) {
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

func (m *unMarshalflowSnmpv2CData) FromPbText(value string) error {
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

func (m *marshalflowSnmpv2CData) ToYaml() (string, error) {
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

func (m *unMarshalflowSnmpv2CData) FromYaml(value string) error {
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

func (m *marshalflowSnmpv2CData) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowSnmpv2CData) ToJson() (string, error) {
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

func (m *unMarshalflowSnmpv2CData) FromJson(value string) error {
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

func (obj *flowSnmpv2CData) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSnmpv2CData) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSnmpv2CData) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSnmpv2CData) Clone() (FlowSnmpv2CData, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSnmpv2CData()
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

func (obj *flowSnmpv2CData) setNil() {
	obj.getRequestHolder = nil
	obj.getNextRequestHolder = nil
	obj.responseHolder = nil
	obj.setRequestHolder = nil
	obj.getBulkRequestHolder = nil
	obj.informRequestHolder = nil
	obj.snmpv2TrapHolder = nil
	obj.reportHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowSnmpv2CData is this contains the body of the SNMPv2C message.
//
// - Encoding of subsequent fields follow ASN.1 specification.
// Refer: http://www.itu.int/ITU-T/asn1
type FlowSnmpv2CData interface {
	Validation
	// msg marshals FlowSnmpv2CData to protobuf object *otg.FlowSnmpv2CData
	// and doesn't set defaults
	msg() *otg.FlowSnmpv2CData
	// setMsg unmarshals FlowSnmpv2CData from protobuf object *otg.FlowSnmpv2CData
	// and doesn't set defaults
	setMsg(*otg.FlowSnmpv2CData) FlowSnmpv2CData
	// provides marshal interface
	Marshal() marshalFlowSnmpv2CData
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSnmpv2CData
	// validate validates FlowSnmpv2CData
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSnmpv2CData, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowSnmpv2CDataChoiceEnum, set in FlowSnmpv2CData
	Choice() FlowSnmpv2CDataChoiceEnum
	// setChoice assigns FlowSnmpv2CDataChoiceEnum provided by user to FlowSnmpv2CData
	setChoice(value FlowSnmpv2CDataChoiceEnum) FlowSnmpv2CData
	// GetRequest returns FlowSnmpv2CPDU, set in FlowSnmpv2CData.
	GetRequest() FlowSnmpv2CPDU
	// SetGetRequest assigns FlowSnmpv2CPDU provided by user to FlowSnmpv2CData.
	SetGetRequest(value FlowSnmpv2CPDU) FlowSnmpv2CData
	// HasGetRequest checks if GetRequest has been set in FlowSnmpv2CData
	HasGetRequest() bool
	// GetNextRequest returns FlowSnmpv2CPDU, set in FlowSnmpv2CData.
	GetNextRequest() FlowSnmpv2CPDU
	// SetGetNextRequest assigns FlowSnmpv2CPDU provided by user to FlowSnmpv2CData.
	SetGetNextRequest(value FlowSnmpv2CPDU) FlowSnmpv2CData
	// HasGetNextRequest checks if GetNextRequest has been set in FlowSnmpv2CData
	HasGetNextRequest() bool
	// Response returns FlowSnmpv2CPDU, set in FlowSnmpv2CData.
	Response() FlowSnmpv2CPDU
	// SetResponse assigns FlowSnmpv2CPDU provided by user to FlowSnmpv2CData.
	SetResponse(value FlowSnmpv2CPDU) FlowSnmpv2CData
	// HasResponse checks if Response has been set in FlowSnmpv2CData
	HasResponse() bool
	// SetRequest returns FlowSnmpv2CPDU, set in FlowSnmpv2CData.
	SetRequest() FlowSnmpv2CPDU
	// SetSetRequest assigns FlowSnmpv2CPDU provided by user to FlowSnmpv2CData.
	SetSetRequest(value FlowSnmpv2CPDU) FlowSnmpv2CData
	// HasSetRequest checks if SetRequest has been set in FlowSnmpv2CData
	HasSetRequest() bool
	// GetBulkRequest returns FlowSnmpv2CBulkPDU, set in FlowSnmpv2CData.
	GetBulkRequest() FlowSnmpv2CBulkPDU
	// SetGetBulkRequest assigns FlowSnmpv2CBulkPDU provided by user to FlowSnmpv2CData.
	SetGetBulkRequest(value FlowSnmpv2CBulkPDU) FlowSnmpv2CData
	// HasGetBulkRequest checks if GetBulkRequest has been set in FlowSnmpv2CData
	HasGetBulkRequest() bool
	// InformRequest returns FlowSnmpv2CPDU, set in FlowSnmpv2CData.
	InformRequest() FlowSnmpv2CPDU
	// SetInformRequest assigns FlowSnmpv2CPDU provided by user to FlowSnmpv2CData.
	SetInformRequest(value FlowSnmpv2CPDU) FlowSnmpv2CData
	// HasInformRequest checks if InformRequest has been set in FlowSnmpv2CData
	HasInformRequest() bool
	// Snmpv2Trap returns FlowSnmpv2CPDU, set in FlowSnmpv2CData.
	Snmpv2Trap() FlowSnmpv2CPDU
	// SetSnmpv2Trap assigns FlowSnmpv2CPDU provided by user to FlowSnmpv2CData.
	SetSnmpv2Trap(value FlowSnmpv2CPDU) FlowSnmpv2CData
	// HasSnmpv2Trap checks if Snmpv2Trap has been set in FlowSnmpv2CData
	HasSnmpv2Trap() bool
	// Report returns FlowSnmpv2CPDU, set in FlowSnmpv2CData.
	Report() FlowSnmpv2CPDU
	// SetReport assigns FlowSnmpv2CPDU provided by user to FlowSnmpv2CData.
	SetReport(value FlowSnmpv2CPDU) FlowSnmpv2CData
	// HasReport checks if Report has been set in FlowSnmpv2CData
	HasReport() bool
	setNil()
}

type FlowSnmpv2CDataChoiceEnum string

// Enum of Choice on FlowSnmpv2CData
var FlowSnmpv2CDataChoice = struct {
	GET_REQUEST      FlowSnmpv2CDataChoiceEnum
	GET_NEXT_REQUEST FlowSnmpv2CDataChoiceEnum
	RESPONSE         FlowSnmpv2CDataChoiceEnum
	SET_REQUEST      FlowSnmpv2CDataChoiceEnum
	GET_BULK_REQUEST FlowSnmpv2CDataChoiceEnum
	INFORM_REQUEST   FlowSnmpv2CDataChoiceEnum
	SNMPV2_TRAP      FlowSnmpv2CDataChoiceEnum
	REPORT           FlowSnmpv2CDataChoiceEnum
}{
	GET_REQUEST:      FlowSnmpv2CDataChoiceEnum("get_request"),
	GET_NEXT_REQUEST: FlowSnmpv2CDataChoiceEnum("get_next_request"),
	RESPONSE:         FlowSnmpv2CDataChoiceEnum("response"),
	SET_REQUEST:      FlowSnmpv2CDataChoiceEnum("set_request"),
	GET_BULK_REQUEST: FlowSnmpv2CDataChoiceEnum("get_bulk_request"),
	INFORM_REQUEST:   FlowSnmpv2CDataChoiceEnum("inform_request"),
	SNMPV2_TRAP:      FlowSnmpv2CDataChoiceEnum("snmpv2_trap"),
	REPORT:           FlowSnmpv2CDataChoiceEnum("report"),
}

func (obj *flowSnmpv2CData) Choice() FlowSnmpv2CDataChoiceEnum {
	return FlowSnmpv2CDataChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *flowSnmpv2CData) setChoice(value FlowSnmpv2CDataChoiceEnum) FlowSnmpv2CData {
	intValue, ok := otg.FlowSnmpv2CData_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowSnmpv2CDataChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowSnmpv2CData_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Report = nil
	obj.reportHolder = nil
	obj.obj.Snmpv2Trap = nil
	obj.snmpv2TrapHolder = nil
	obj.obj.InformRequest = nil
	obj.informRequestHolder = nil
	obj.obj.GetBulkRequest = nil
	obj.getBulkRequestHolder = nil
	obj.obj.SetRequest = nil
	obj.setRequestHolder = nil
	obj.obj.Response = nil
	obj.responseHolder = nil
	obj.obj.GetNextRequest = nil
	obj.getNextRequestHolder = nil
	obj.obj.GetRequest = nil
	obj.getRequestHolder = nil

	if value == FlowSnmpv2CDataChoice.GET_REQUEST {
		obj.obj.GetRequest = NewFlowSnmpv2CPDU().msg()
	}

	if value == FlowSnmpv2CDataChoice.GET_NEXT_REQUEST {
		obj.obj.GetNextRequest = NewFlowSnmpv2CPDU().msg()
	}

	if value == FlowSnmpv2CDataChoice.RESPONSE {
		obj.obj.Response = NewFlowSnmpv2CPDU().msg()
	}

	if value == FlowSnmpv2CDataChoice.SET_REQUEST {
		obj.obj.SetRequest = NewFlowSnmpv2CPDU().msg()
	}

	if value == FlowSnmpv2CDataChoice.GET_BULK_REQUEST {
		obj.obj.GetBulkRequest = NewFlowSnmpv2CBulkPDU().msg()
	}

	if value == FlowSnmpv2CDataChoice.INFORM_REQUEST {
		obj.obj.InformRequest = NewFlowSnmpv2CPDU().msg()
	}

	if value == FlowSnmpv2CDataChoice.SNMPV2_TRAP {
		obj.obj.Snmpv2Trap = NewFlowSnmpv2CPDU().msg()
	}

	if value == FlowSnmpv2CDataChoice.REPORT {
		obj.obj.Report = NewFlowSnmpv2CPDU().msg()
	}

	return obj
}

// description is TBD
// GetRequest returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) GetRequest() FlowSnmpv2CPDU {
	if obj.obj.GetRequest == nil {
		obj.setChoice(FlowSnmpv2CDataChoice.GET_REQUEST)
	}
	if obj.getRequestHolder == nil {
		obj.getRequestHolder = &flowSnmpv2CPDU{obj: obj.obj.GetRequest}
	}
	return obj.getRequestHolder
}

// description is TBD
// GetRequest returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) HasGetRequest() bool {
	return obj.obj.GetRequest != nil
}

// description is TBD
// SetGetRequest sets the FlowSnmpv2CPDU value in the FlowSnmpv2CData object
func (obj *flowSnmpv2CData) SetGetRequest(value FlowSnmpv2CPDU) FlowSnmpv2CData {
	obj.setChoice(FlowSnmpv2CDataChoice.GET_REQUEST)
	obj.getRequestHolder = nil
	obj.obj.GetRequest = value.msg()

	return obj
}

// description is TBD
// GetNextRequest returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) GetNextRequest() FlowSnmpv2CPDU {
	if obj.obj.GetNextRequest == nil {
		obj.setChoice(FlowSnmpv2CDataChoice.GET_NEXT_REQUEST)
	}
	if obj.getNextRequestHolder == nil {
		obj.getNextRequestHolder = &flowSnmpv2CPDU{obj: obj.obj.GetNextRequest}
	}
	return obj.getNextRequestHolder
}

// description is TBD
// GetNextRequest returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) HasGetNextRequest() bool {
	return obj.obj.GetNextRequest != nil
}

// description is TBD
// SetGetNextRequest sets the FlowSnmpv2CPDU value in the FlowSnmpv2CData object
func (obj *flowSnmpv2CData) SetGetNextRequest(value FlowSnmpv2CPDU) FlowSnmpv2CData {
	obj.setChoice(FlowSnmpv2CDataChoice.GET_NEXT_REQUEST)
	obj.getNextRequestHolder = nil
	obj.obj.GetNextRequest = value.msg()

	return obj
}

// description is TBD
// Response returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) Response() FlowSnmpv2CPDU {
	if obj.obj.Response == nil {
		obj.setChoice(FlowSnmpv2CDataChoice.RESPONSE)
	}
	if obj.responseHolder == nil {
		obj.responseHolder = &flowSnmpv2CPDU{obj: obj.obj.Response}
	}
	return obj.responseHolder
}

// description is TBD
// Response returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) HasResponse() bool {
	return obj.obj.Response != nil
}

// description is TBD
// SetResponse sets the FlowSnmpv2CPDU value in the FlowSnmpv2CData object
func (obj *flowSnmpv2CData) SetResponse(value FlowSnmpv2CPDU) FlowSnmpv2CData {
	obj.setChoice(FlowSnmpv2CDataChoice.RESPONSE)
	obj.responseHolder = nil
	obj.obj.Response = value.msg()

	return obj
}

// description is TBD
// SetRequest returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) SetRequest() FlowSnmpv2CPDU {
	if obj.obj.SetRequest == nil {
		obj.setChoice(FlowSnmpv2CDataChoice.SET_REQUEST)
	}
	if obj.setRequestHolder == nil {
		obj.setRequestHolder = &flowSnmpv2CPDU{obj: obj.obj.SetRequest}
	}
	return obj.setRequestHolder
}

// description is TBD
// SetRequest returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) HasSetRequest() bool {
	return obj.obj.SetRequest != nil
}

// description is TBD
// SetSetRequest sets the FlowSnmpv2CPDU value in the FlowSnmpv2CData object
func (obj *flowSnmpv2CData) SetSetRequest(value FlowSnmpv2CPDU) FlowSnmpv2CData {
	obj.setChoice(FlowSnmpv2CDataChoice.SET_REQUEST)
	obj.setRequestHolder = nil
	obj.obj.SetRequest = value.msg()

	return obj
}

// description is TBD
// GetBulkRequest returns a FlowSnmpv2CBulkPDU
func (obj *flowSnmpv2CData) GetBulkRequest() FlowSnmpv2CBulkPDU {
	if obj.obj.GetBulkRequest == nil {
		obj.setChoice(FlowSnmpv2CDataChoice.GET_BULK_REQUEST)
	}
	if obj.getBulkRequestHolder == nil {
		obj.getBulkRequestHolder = &flowSnmpv2CBulkPDU{obj: obj.obj.GetBulkRequest}
	}
	return obj.getBulkRequestHolder
}

// description is TBD
// GetBulkRequest returns a FlowSnmpv2CBulkPDU
func (obj *flowSnmpv2CData) HasGetBulkRequest() bool {
	return obj.obj.GetBulkRequest != nil
}

// description is TBD
// SetGetBulkRequest sets the FlowSnmpv2CBulkPDU value in the FlowSnmpv2CData object
func (obj *flowSnmpv2CData) SetGetBulkRequest(value FlowSnmpv2CBulkPDU) FlowSnmpv2CData {
	obj.setChoice(FlowSnmpv2CDataChoice.GET_BULK_REQUEST)
	obj.getBulkRequestHolder = nil
	obj.obj.GetBulkRequest = value.msg()

	return obj
}

// description is TBD
// InformRequest returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) InformRequest() FlowSnmpv2CPDU {
	if obj.obj.InformRequest == nil {
		obj.setChoice(FlowSnmpv2CDataChoice.INFORM_REQUEST)
	}
	if obj.informRequestHolder == nil {
		obj.informRequestHolder = &flowSnmpv2CPDU{obj: obj.obj.InformRequest}
	}
	return obj.informRequestHolder
}

// description is TBD
// InformRequest returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) HasInformRequest() bool {
	return obj.obj.InformRequest != nil
}

// description is TBD
// SetInformRequest sets the FlowSnmpv2CPDU value in the FlowSnmpv2CData object
func (obj *flowSnmpv2CData) SetInformRequest(value FlowSnmpv2CPDU) FlowSnmpv2CData {
	obj.setChoice(FlowSnmpv2CDataChoice.INFORM_REQUEST)
	obj.informRequestHolder = nil
	obj.obj.InformRequest = value.msg()

	return obj
}

// description is TBD
// Snmpv2Trap returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) Snmpv2Trap() FlowSnmpv2CPDU {
	if obj.obj.Snmpv2Trap == nil {
		obj.setChoice(FlowSnmpv2CDataChoice.SNMPV2_TRAP)
	}
	if obj.snmpv2TrapHolder == nil {
		obj.snmpv2TrapHolder = &flowSnmpv2CPDU{obj: obj.obj.Snmpv2Trap}
	}
	return obj.snmpv2TrapHolder
}

// description is TBD
// Snmpv2Trap returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) HasSnmpv2Trap() bool {
	return obj.obj.Snmpv2Trap != nil
}

// description is TBD
// SetSnmpv2Trap sets the FlowSnmpv2CPDU value in the FlowSnmpv2CData object
func (obj *flowSnmpv2CData) SetSnmpv2Trap(value FlowSnmpv2CPDU) FlowSnmpv2CData {
	obj.setChoice(FlowSnmpv2CDataChoice.SNMPV2_TRAP)
	obj.snmpv2TrapHolder = nil
	obj.obj.Snmpv2Trap = value.msg()

	return obj
}

// description is TBD
// Report returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) Report() FlowSnmpv2CPDU {
	if obj.obj.Report == nil {
		obj.setChoice(FlowSnmpv2CDataChoice.REPORT)
	}
	if obj.reportHolder == nil {
		obj.reportHolder = &flowSnmpv2CPDU{obj: obj.obj.Report}
	}
	return obj.reportHolder
}

// description is TBD
// Report returns a FlowSnmpv2CPDU
func (obj *flowSnmpv2CData) HasReport() bool {
	return obj.obj.Report != nil
}

// description is TBD
// SetReport sets the FlowSnmpv2CPDU value in the FlowSnmpv2CData object
func (obj *flowSnmpv2CData) SetReport(value FlowSnmpv2CPDU) FlowSnmpv2CData {
	obj.setChoice(FlowSnmpv2CDataChoice.REPORT)
	obj.reportHolder = nil
	obj.obj.Report = value.msg()

	return obj
}

func (obj *flowSnmpv2CData) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface FlowSnmpv2CData")
	}

	if obj.obj.GetRequest != nil {

		obj.GetRequest().validateObj(vObj, set_default)
	}

	if obj.obj.GetNextRequest != nil {

		obj.GetNextRequest().validateObj(vObj, set_default)
	}

	if obj.obj.Response != nil {

		obj.Response().validateObj(vObj, set_default)
	}

	if obj.obj.SetRequest != nil {

		obj.SetRequest().validateObj(vObj, set_default)
	}

	if obj.obj.GetBulkRequest != nil {

		obj.GetBulkRequest().validateObj(vObj, set_default)
	}

	if obj.obj.InformRequest != nil {

		obj.InformRequest().validateObj(vObj, set_default)
	}

	if obj.obj.Snmpv2Trap != nil {

		obj.Snmpv2Trap().validateObj(vObj, set_default)
	}

	if obj.obj.Report != nil {

		obj.Report().validateObj(vObj, set_default)
	}

}

func (obj *flowSnmpv2CData) setDefault() {
	var choices_set int = 0
	var choice FlowSnmpv2CDataChoiceEnum

	if obj.obj.GetRequest != nil {
		choices_set += 1
		choice = FlowSnmpv2CDataChoice.GET_REQUEST
	}

	if obj.obj.GetNextRequest != nil {
		choices_set += 1
		choice = FlowSnmpv2CDataChoice.GET_NEXT_REQUEST
	}

	if obj.obj.Response != nil {
		choices_set += 1
		choice = FlowSnmpv2CDataChoice.RESPONSE
	}

	if obj.obj.SetRequest != nil {
		choices_set += 1
		choice = FlowSnmpv2CDataChoice.SET_REQUEST
	}

	if obj.obj.GetBulkRequest != nil {
		choices_set += 1
		choice = FlowSnmpv2CDataChoice.GET_BULK_REQUEST
	}

	if obj.obj.InformRequest != nil {
		choices_set += 1
		choice = FlowSnmpv2CDataChoice.INFORM_REQUEST
	}

	if obj.obj.Snmpv2Trap != nil {
		choices_set += 1
		choice = FlowSnmpv2CDataChoice.SNMPV2_TRAP
	}

	if obj.obj.Report != nil {
		choices_set += 1
		choice = FlowSnmpv2CDataChoice.REPORT
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowSnmpv2CData")
			}
		} else {
			intVal := otg.FlowSnmpv2CData_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowSnmpv2CData_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
