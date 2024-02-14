package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClass *****
type flowRSVPPathObjectsClass struct {
	validation
	obj                    *otg.FlowRSVPPathObjectsClass
	marshaller             marshalFlowRSVPPathObjectsClass
	unMarshaller           unMarshalFlowRSVPPathObjectsClass
	sessionHolder          FlowRSVPPathObjectsClassSession
	rsvpHopHolder          FlowRSVPPathObjectsClassRsvpHop
	timeValuesHolder       FlowRSVPPathObjectsClassTimeValues
	explicitRouteHolder    FlowRSVPPathObjectsClassExplicitRoute
	labelRequestHolder     FlowRSVPPathObjectsClassLabelRequest
	sessionAttributeHolder FlowRSVPPathObjectsClassSessionAttribute
	senderTemplateHolder   FlowRSVPPathObjectsClassSenderTemplate
	senderTspecHolder      FlowRSVPPathObjectsClassSenderTspec
	recordRouteHolder      FlowRSVPPathObjectsClassRecordRoute
	customHolder           FlowRSVPPathObjectsCustom
}

func NewFlowRSVPPathObjectsClass() FlowRSVPPathObjectsClass {
	obj := flowRSVPPathObjectsClass{obj: &otg.FlowRSVPPathObjectsClass{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClass) msg() *otg.FlowRSVPPathObjectsClass {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClass) setMsg(msg *otg.FlowRSVPPathObjectsClass) FlowRSVPPathObjectsClass {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClass struct {
	obj *flowRSVPPathObjectsClass
}

type marshalFlowRSVPPathObjectsClass interface {
	// ToProto marshals FlowRSVPPathObjectsClass to protobuf object *otg.FlowRSVPPathObjectsClass
	ToProto() (*otg.FlowRSVPPathObjectsClass, error)
	// ToPbText marshals FlowRSVPPathObjectsClass to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClass to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClass to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsClass struct {
	obj *flowRSVPPathObjectsClass
}

type unMarshalFlowRSVPPathObjectsClass interface {
	// FromProto unmarshals FlowRSVPPathObjectsClass from protobuf object *otg.FlowRSVPPathObjectsClass
	FromProto(msg *otg.FlowRSVPPathObjectsClass) (FlowRSVPPathObjectsClass, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClass from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClass from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClass from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClass) Marshal() marshalFlowRSVPPathObjectsClass {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClass{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClass) Unmarshal() unMarshalFlowRSVPPathObjectsClass {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClass{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClass) ToProto() (*otg.FlowRSVPPathObjectsClass, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClass) FromProto(msg *otg.FlowRSVPPathObjectsClass) (FlowRSVPPathObjectsClass, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClass) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClass) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClass) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClass) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClass) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClass) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClass) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClass) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClass) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClass) Clone() (FlowRSVPPathObjectsClass, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClass()
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

func (obj *flowRSVPPathObjectsClass) setNil() {
	obj.sessionHolder = nil
	obj.rsvpHopHolder = nil
	obj.timeValuesHolder = nil
	obj.explicitRouteHolder = nil
	obj.labelRequestHolder = nil
	obj.sessionAttributeHolder = nil
	obj.senderTemplateHolder = nil
	obj.senderTspecHolder = nil
	obj.recordRouteHolder = nil
	obj.customHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClass is the class number is used to identify the class of an object. Defined in https://www.iana.org/assignments/rsvp-parameters/rsvp-parameters.xhtml#rsvp-parameters-4 . Curently supported class numbers are for "Path" message type. "Path" message: Supported Class numbers and it's value: SESSION: 1, RSVP_HOP: 3, TIME_VALUES: 5, EXPLICIT_ROUTE: 20, LABEL_REQUEST: 19, SESSION_ATTRIBUTE: 207, SENDER_TEMPLATE: 11, SENDER_TSPEC: 12, RECORD_ROUTE: 21, Custom: User defined bytes based on class and c-types not supported in above options.
type FlowRSVPPathObjectsClass interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClass to protobuf object *otg.FlowRSVPPathObjectsClass
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClass
	// setMsg unmarshals FlowRSVPPathObjectsClass from protobuf object *otg.FlowRSVPPathObjectsClass
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClass) FlowRSVPPathObjectsClass
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClass
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClass
	// validate validates FlowRSVPPathObjectsClass
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClass, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsClassChoiceEnum, set in FlowRSVPPathObjectsClass
	Choice() FlowRSVPPathObjectsClassChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsClassChoiceEnum provided by user to FlowRSVPPathObjectsClass
	setChoice(value FlowRSVPPathObjectsClassChoiceEnum) FlowRSVPPathObjectsClass
	// Session returns FlowRSVPPathObjectsClassSession, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassSession is c-Type is specific to a class num.
	Session() FlowRSVPPathObjectsClassSession
	// SetSession assigns FlowRSVPPathObjectsClassSession provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassSession is c-Type is specific to a class num.
	SetSession(value FlowRSVPPathObjectsClassSession) FlowRSVPPathObjectsClass
	// HasSession checks if Session has been set in FlowRSVPPathObjectsClass
	HasSession() bool
	// RsvpHop returns FlowRSVPPathObjectsClassRsvpHop, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassRsvpHop is c-Type is specific to a class num.
	RsvpHop() FlowRSVPPathObjectsClassRsvpHop
	// SetRsvpHop assigns FlowRSVPPathObjectsClassRsvpHop provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassRsvpHop is c-Type is specific to a class num.
	SetRsvpHop(value FlowRSVPPathObjectsClassRsvpHop) FlowRSVPPathObjectsClass
	// HasRsvpHop checks if RsvpHop has been set in FlowRSVPPathObjectsClass
	HasRsvpHop() bool
	// TimeValues returns FlowRSVPPathObjectsClassTimeValues, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassTimeValues is c-Type is specific to a class num.
	TimeValues() FlowRSVPPathObjectsClassTimeValues
	// SetTimeValues assigns FlowRSVPPathObjectsClassTimeValues provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassTimeValues is c-Type is specific to a class num.
	SetTimeValues(value FlowRSVPPathObjectsClassTimeValues) FlowRSVPPathObjectsClass
	// HasTimeValues checks if TimeValues has been set in FlowRSVPPathObjectsClass
	HasTimeValues() bool
	// ExplicitRoute returns FlowRSVPPathObjectsClassExplicitRoute, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassExplicitRoute is c-Type is specific to a class num.
	ExplicitRoute() FlowRSVPPathObjectsClassExplicitRoute
	// SetExplicitRoute assigns FlowRSVPPathObjectsClassExplicitRoute provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassExplicitRoute is c-Type is specific to a class num.
	SetExplicitRoute(value FlowRSVPPathObjectsClassExplicitRoute) FlowRSVPPathObjectsClass
	// HasExplicitRoute checks if ExplicitRoute has been set in FlowRSVPPathObjectsClass
	HasExplicitRoute() bool
	// LabelRequest returns FlowRSVPPathObjectsClassLabelRequest, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassLabelRequest is c-Type is specific to a class num.
	LabelRequest() FlowRSVPPathObjectsClassLabelRequest
	// SetLabelRequest assigns FlowRSVPPathObjectsClassLabelRequest provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassLabelRequest is c-Type is specific to a class num.
	SetLabelRequest(value FlowRSVPPathObjectsClassLabelRequest) FlowRSVPPathObjectsClass
	// HasLabelRequest checks if LabelRequest has been set in FlowRSVPPathObjectsClass
	HasLabelRequest() bool
	// SessionAttribute returns FlowRSVPPathObjectsClassSessionAttribute, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassSessionAttribute is c-Type is specific to a class num.
	SessionAttribute() FlowRSVPPathObjectsClassSessionAttribute
	// SetSessionAttribute assigns FlowRSVPPathObjectsClassSessionAttribute provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassSessionAttribute is c-Type is specific to a class num.
	SetSessionAttribute(value FlowRSVPPathObjectsClassSessionAttribute) FlowRSVPPathObjectsClass
	// HasSessionAttribute checks if SessionAttribute has been set in FlowRSVPPathObjectsClass
	HasSessionAttribute() bool
	// SenderTemplate returns FlowRSVPPathObjectsClassSenderTemplate, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassSenderTemplate is c-Type is specific to a class num.
	SenderTemplate() FlowRSVPPathObjectsClassSenderTemplate
	// SetSenderTemplate assigns FlowRSVPPathObjectsClassSenderTemplate provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassSenderTemplate is c-Type is specific to a class num.
	SetSenderTemplate(value FlowRSVPPathObjectsClassSenderTemplate) FlowRSVPPathObjectsClass
	// HasSenderTemplate checks if SenderTemplate has been set in FlowRSVPPathObjectsClass
	HasSenderTemplate() bool
	// SenderTspec returns FlowRSVPPathObjectsClassSenderTspec, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassSenderTspec is c-Type is specific to a class num.
	SenderTspec() FlowRSVPPathObjectsClassSenderTspec
	// SetSenderTspec assigns FlowRSVPPathObjectsClassSenderTspec provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassSenderTspec is c-Type is specific to a class num.
	SetSenderTspec(value FlowRSVPPathObjectsClassSenderTspec) FlowRSVPPathObjectsClass
	// HasSenderTspec checks if SenderTspec has been set in FlowRSVPPathObjectsClass
	HasSenderTspec() bool
	// RecordRoute returns FlowRSVPPathObjectsClassRecordRoute, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassRecordRoute is c-Type is specific to a class num.
	RecordRoute() FlowRSVPPathObjectsClassRecordRoute
	// SetRecordRoute assigns FlowRSVPPathObjectsClassRecordRoute provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsClassRecordRoute is c-Type is specific to a class num.
	SetRecordRoute(value FlowRSVPPathObjectsClassRecordRoute) FlowRSVPPathObjectsClass
	// HasRecordRoute checks if RecordRoute has been set in FlowRSVPPathObjectsClass
	HasRecordRoute() bool
	// Custom returns FlowRSVPPathObjectsCustom, set in FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsCustom is custom packet header
	Custom() FlowRSVPPathObjectsCustom
	// SetCustom assigns FlowRSVPPathObjectsCustom provided by user to FlowRSVPPathObjectsClass.
	// FlowRSVPPathObjectsCustom is custom packet header
	SetCustom(value FlowRSVPPathObjectsCustom) FlowRSVPPathObjectsClass
	// HasCustom checks if Custom has been set in FlowRSVPPathObjectsClass
	HasCustom() bool
	setNil()
}

type FlowRSVPPathObjectsClassChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsClass
var FlowRSVPPathObjectsClassChoice = struct {
	SESSION           FlowRSVPPathObjectsClassChoiceEnum
	RSVP_HOP          FlowRSVPPathObjectsClassChoiceEnum
	TIME_VALUES       FlowRSVPPathObjectsClassChoiceEnum
	EXPLICIT_ROUTE    FlowRSVPPathObjectsClassChoiceEnum
	LABEL_REQUEST     FlowRSVPPathObjectsClassChoiceEnum
	SESSION_ATTRIBUTE FlowRSVPPathObjectsClassChoiceEnum
	SENDER_TEMPLATE   FlowRSVPPathObjectsClassChoiceEnum
	SENDER_TSPEC      FlowRSVPPathObjectsClassChoiceEnum
	RECORD_ROUTE      FlowRSVPPathObjectsClassChoiceEnum
	CUSTOM            FlowRSVPPathObjectsClassChoiceEnum
}{
	SESSION:           FlowRSVPPathObjectsClassChoiceEnum("session"),
	RSVP_HOP:          FlowRSVPPathObjectsClassChoiceEnum("rsvp_hop"),
	TIME_VALUES:       FlowRSVPPathObjectsClassChoiceEnum("time_values"),
	EXPLICIT_ROUTE:    FlowRSVPPathObjectsClassChoiceEnum("explicit_route"),
	LABEL_REQUEST:     FlowRSVPPathObjectsClassChoiceEnum("label_request"),
	SESSION_ATTRIBUTE: FlowRSVPPathObjectsClassChoiceEnum("session_attribute"),
	SENDER_TEMPLATE:   FlowRSVPPathObjectsClassChoiceEnum("sender_template"),
	SENDER_TSPEC:      FlowRSVPPathObjectsClassChoiceEnum("sender_tspec"),
	RECORD_ROUTE:      FlowRSVPPathObjectsClassChoiceEnum("record_route"),
	CUSTOM:            FlowRSVPPathObjectsClassChoiceEnum("custom"),
}

func (obj *flowRSVPPathObjectsClass) Choice() FlowRSVPPathObjectsClassChoiceEnum {
	return FlowRSVPPathObjectsClassChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *flowRSVPPathObjectsClass) setChoice(value FlowRSVPPathObjectsClassChoiceEnum) FlowRSVPPathObjectsClass {
	intValue, ok := otg.FlowRSVPPathObjectsClass_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsClassChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsClass_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.customHolder = nil
	obj.obj.RecordRoute = nil
	obj.recordRouteHolder = nil
	obj.obj.SenderTspec = nil
	obj.senderTspecHolder = nil
	obj.obj.SenderTemplate = nil
	obj.senderTemplateHolder = nil
	obj.obj.SessionAttribute = nil
	obj.sessionAttributeHolder = nil
	obj.obj.LabelRequest = nil
	obj.labelRequestHolder = nil
	obj.obj.ExplicitRoute = nil
	obj.explicitRouteHolder = nil
	obj.obj.TimeValues = nil
	obj.timeValuesHolder = nil
	obj.obj.RsvpHop = nil
	obj.rsvpHopHolder = nil
	obj.obj.Session = nil
	obj.sessionHolder = nil

	if value == FlowRSVPPathObjectsClassChoice.SESSION {
		obj.obj.Session = NewFlowRSVPPathObjectsClassSession().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.RSVP_HOP {
		obj.obj.RsvpHop = NewFlowRSVPPathObjectsClassRsvpHop().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.TIME_VALUES {
		obj.obj.TimeValues = NewFlowRSVPPathObjectsClassTimeValues().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.EXPLICIT_ROUTE {
		obj.obj.ExplicitRoute = NewFlowRSVPPathObjectsClassExplicitRoute().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.LABEL_REQUEST {
		obj.obj.LabelRequest = NewFlowRSVPPathObjectsClassLabelRequest().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.SESSION_ATTRIBUTE {
		obj.obj.SessionAttribute = NewFlowRSVPPathObjectsClassSessionAttribute().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.SENDER_TEMPLATE {
		obj.obj.SenderTemplate = NewFlowRSVPPathObjectsClassSenderTemplate().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.SENDER_TSPEC {
		obj.obj.SenderTspec = NewFlowRSVPPathObjectsClassSenderTspec().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.RECORD_ROUTE {
		obj.obj.RecordRoute = NewFlowRSVPPathObjectsClassRecordRoute().msg()
	}

	if value == FlowRSVPPathObjectsClassChoice.CUSTOM {
		obj.obj.Custom = NewFlowRSVPPathObjectsCustom().msg()
	}

	return obj
}

// description is TBD
// Session returns a FlowRSVPPathObjectsClassSession
func (obj *flowRSVPPathObjectsClass) Session() FlowRSVPPathObjectsClassSession {
	if obj.obj.Session == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.SESSION)
	}
	if obj.sessionHolder == nil {
		obj.sessionHolder = &flowRSVPPathObjectsClassSession{obj: obj.obj.Session}
	}
	return obj.sessionHolder
}

// description is TBD
// Session returns a FlowRSVPPathObjectsClassSession
func (obj *flowRSVPPathObjectsClass) HasSession() bool {
	return obj.obj.Session != nil
}

// description is TBD
// SetSession sets the FlowRSVPPathObjectsClassSession value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetSession(value FlowRSVPPathObjectsClassSession) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.SESSION)
	obj.sessionHolder = nil
	obj.obj.Session = value.msg()

	return obj
}

// description is TBD
// RsvpHop returns a FlowRSVPPathObjectsClassRsvpHop
func (obj *flowRSVPPathObjectsClass) RsvpHop() FlowRSVPPathObjectsClassRsvpHop {
	if obj.obj.RsvpHop == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.RSVP_HOP)
	}
	if obj.rsvpHopHolder == nil {
		obj.rsvpHopHolder = &flowRSVPPathObjectsClassRsvpHop{obj: obj.obj.RsvpHop}
	}
	return obj.rsvpHopHolder
}

// description is TBD
// RsvpHop returns a FlowRSVPPathObjectsClassRsvpHop
func (obj *flowRSVPPathObjectsClass) HasRsvpHop() bool {
	return obj.obj.RsvpHop != nil
}

// description is TBD
// SetRsvpHop sets the FlowRSVPPathObjectsClassRsvpHop value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetRsvpHop(value FlowRSVPPathObjectsClassRsvpHop) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.RSVP_HOP)
	obj.rsvpHopHolder = nil
	obj.obj.RsvpHop = value.msg()

	return obj
}

// description is TBD
// TimeValues returns a FlowRSVPPathObjectsClassTimeValues
func (obj *flowRSVPPathObjectsClass) TimeValues() FlowRSVPPathObjectsClassTimeValues {
	if obj.obj.TimeValues == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.TIME_VALUES)
	}
	if obj.timeValuesHolder == nil {
		obj.timeValuesHolder = &flowRSVPPathObjectsClassTimeValues{obj: obj.obj.TimeValues}
	}
	return obj.timeValuesHolder
}

// description is TBD
// TimeValues returns a FlowRSVPPathObjectsClassTimeValues
func (obj *flowRSVPPathObjectsClass) HasTimeValues() bool {
	return obj.obj.TimeValues != nil
}

// description is TBD
// SetTimeValues sets the FlowRSVPPathObjectsClassTimeValues value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetTimeValues(value FlowRSVPPathObjectsClassTimeValues) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.TIME_VALUES)
	obj.timeValuesHolder = nil
	obj.obj.TimeValues = value.msg()

	return obj
}

// description is TBD
// ExplicitRoute returns a FlowRSVPPathObjectsClassExplicitRoute
func (obj *flowRSVPPathObjectsClass) ExplicitRoute() FlowRSVPPathObjectsClassExplicitRoute {
	if obj.obj.ExplicitRoute == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.EXPLICIT_ROUTE)
	}
	if obj.explicitRouteHolder == nil {
		obj.explicitRouteHolder = &flowRSVPPathObjectsClassExplicitRoute{obj: obj.obj.ExplicitRoute}
	}
	return obj.explicitRouteHolder
}

// description is TBD
// ExplicitRoute returns a FlowRSVPPathObjectsClassExplicitRoute
func (obj *flowRSVPPathObjectsClass) HasExplicitRoute() bool {
	return obj.obj.ExplicitRoute != nil
}

// description is TBD
// SetExplicitRoute sets the FlowRSVPPathObjectsClassExplicitRoute value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetExplicitRoute(value FlowRSVPPathObjectsClassExplicitRoute) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.EXPLICIT_ROUTE)
	obj.explicitRouteHolder = nil
	obj.obj.ExplicitRoute = value.msg()

	return obj
}

// description is TBD
// LabelRequest returns a FlowRSVPPathObjectsClassLabelRequest
func (obj *flowRSVPPathObjectsClass) LabelRequest() FlowRSVPPathObjectsClassLabelRequest {
	if obj.obj.LabelRequest == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.LABEL_REQUEST)
	}
	if obj.labelRequestHolder == nil {
		obj.labelRequestHolder = &flowRSVPPathObjectsClassLabelRequest{obj: obj.obj.LabelRequest}
	}
	return obj.labelRequestHolder
}

// description is TBD
// LabelRequest returns a FlowRSVPPathObjectsClassLabelRequest
func (obj *flowRSVPPathObjectsClass) HasLabelRequest() bool {
	return obj.obj.LabelRequest != nil
}

// description is TBD
// SetLabelRequest sets the FlowRSVPPathObjectsClassLabelRequest value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetLabelRequest(value FlowRSVPPathObjectsClassLabelRequest) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.LABEL_REQUEST)
	obj.labelRequestHolder = nil
	obj.obj.LabelRequest = value.msg()

	return obj
}

// description is TBD
// SessionAttribute returns a FlowRSVPPathObjectsClassSessionAttribute
func (obj *flowRSVPPathObjectsClass) SessionAttribute() FlowRSVPPathObjectsClassSessionAttribute {
	if obj.obj.SessionAttribute == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.SESSION_ATTRIBUTE)
	}
	if obj.sessionAttributeHolder == nil {
		obj.sessionAttributeHolder = &flowRSVPPathObjectsClassSessionAttribute{obj: obj.obj.SessionAttribute}
	}
	return obj.sessionAttributeHolder
}

// description is TBD
// SessionAttribute returns a FlowRSVPPathObjectsClassSessionAttribute
func (obj *flowRSVPPathObjectsClass) HasSessionAttribute() bool {
	return obj.obj.SessionAttribute != nil
}

// description is TBD
// SetSessionAttribute sets the FlowRSVPPathObjectsClassSessionAttribute value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetSessionAttribute(value FlowRSVPPathObjectsClassSessionAttribute) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.SESSION_ATTRIBUTE)
	obj.sessionAttributeHolder = nil
	obj.obj.SessionAttribute = value.msg()

	return obj
}

// description is TBD
// SenderTemplate returns a FlowRSVPPathObjectsClassSenderTemplate
func (obj *flowRSVPPathObjectsClass) SenderTemplate() FlowRSVPPathObjectsClassSenderTemplate {
	if obj.obj.SenderTemplate == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.SENDER_TEMPLATE)
	}
	if obj.senderTemplateHolder == nil {
		obj.senderTemplateHolder = &flowRSVPPathObjectsClassSenderTemplate{obj: obj.obj.SenderTemplate}
	}
	return obj.senderTemplateHolder
}

// description is TBD
// SenderTemplate returns a FlowRSVPPathObjectsClassSenderTemplate
func (obj *flowRSVPPathObjectsClass) HasSenderTemplate() bool {
	return obj.obj.SenderTemplate != nil
}

// description is TBD
// SetSenderTemplate sets the FlowRSVPPathObjectsClassSenderTemplate value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetSenderTemplate(value FlowRSVPPathObjectsClassSenderTemplate) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.SENDER_TEMPLATE)
	obj.senderTemplateHolder = nil
	obj.obj.SenderTemplate = value.msg()

	return obj
}

// description is TBD
// SenderTspec returns a FlowRSVPPathObjectsClassSenderTspec
func (obj *flowRSVPPathObjectsClass) SenderTspec() FlowRSVPPathObjectsClassSenderTspec {
	if obj.obj.SenderTspec == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.SENDER_TSPEC)
	}
	if obj.senderTspecHolder == nil {
		obj.senderTspecHolder = &flowRSVPPathObjectsClassSenderTspec{obj: obj.obj.SenderTspec}
	}
	return obj.senderTspecHolder
}

// description is TBD
// SenderTspec returns a FlowRSVPPathObjectsClassSenderTspec
func (obj *flowRSVPPathObjectsClass) HasSenderTspec() bool {
	return obj.obj.SenderTspec != nil
}

// description is TBD
// SetSenderTspec sets the FlowRSVPPathObjectsClassSenderTspec value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetSenderTspec(value FlowRSVPPathObjectsClassSenderTspec) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.SENDER_TSPEC)
	obj.senderTspecHolder = nil
	obj.obj.SenderTspec = value.msg()

	return obj
}

// description is TBD
// RecordRoute returns a FlowRSVPPathObjectsClassRecordRoute
func (obj *flowRSVPPathObjectsClass) RecordRoute() FlowRSVPPathObjectsClassRecordRoute {
	if obj.obj.RecordRoute == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.RECORD_ROUTE)
	}
	if obj.recordRouteHolder == nil {
		obj.recordRouteHolder = &flowRSVPPathObjectsClassRecordRoute{obj: obj.obj.RecordRoute}
	}
	return obj.recordRouteHolder
}

// description is TBD
// RecordRoute returns a FlowRSVPPathObjectsClassRecordRoute
func (obj *flowRSVPPathObjectsClass) HasRecordRoute() bool {
	return obj.obj.RecordRoute != nil
}

// description is TBD
// SetRecordRoute sets the FlowRSVPPathObjectsClassRecordRoute value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetRecordRoute(value FlowRSVPPathObjectsClassRecordRoute) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.RECORD_ROUTE)
	obj.recordRouteHolder = nil
	obj.obj.RecordRoute = value.msg()

	return obj
}

// description is TBD
// Custom returns a FlowRSVPPathObjectsCustom
func (obj *flowRSVPPathObjectsClass) Custom() FlowRSVPPathObjectsCustom {
	if obj.obj.Custom == nil {
		obj.setChoice(FlowRSVPPathObjectsClassChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &flowRSVPPathObjectsCustom{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// description is TBD
// Custom returns a FlowRSVPPathObjectsCustom
func (obj *flowRSVPPathObjectsClass) HasCustom() bool {
	return obj.obj.Custom != nil
}

// description is TBD
// SetCustom sets the FlowRSVPPathObjectsCustom value in the FlowRSVPPathObjectsClass object
func (obj *flowRSVPPathObjectsClass) SetCustom(value FlowRSVPPathObjectsCustom) FlowRSVPPathObjectsClass {
	obj.setChoice(FlowRSVPPathObjectsClassChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClass) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface FlowRSVPPathObjectsClass")
	}

	if obj.obj.Session != nil {

		obj.Session().validateObj(vObj, set_default)
	}

	if obj.obj.RsvpHop != nil {

		obj.RsvpHop().validateObj(vObj, set_default)
	}

	if obj.obj.TimeValues != nil {

		obj.TimeValues().validateObj(vObj, set_default)
	}

	if obj.obj.ExplicitRoute != nil {

		obj.ExplicitRoute().validateObj(vObj, set_default)
	}

	if obj.obj.LabelRequest != nil {

		obj.LabelRequest().validateObj(vObj, set_default)
	}

	if obj.obj.SessionAttribute != nil {

		obj.SessionAttribute().validateObj(vObj, set_default)
	}

	if obj.obj.SenderTemplate != nil {

		obj.SenderTemplate().validateObj(vObj, set_default)
	}

	if obj.obj.SenderTspec != nil {

		obj.SenderTspec().validateObj(vObj, set_default)
	}

	if obj.obj.RecordRoute != nil {

		obj.RecordRoute().validateObj(vObj, set_default)
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsClass) setDefault() {

}
