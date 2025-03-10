package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2Metric *****
type ospfv2Metric struct {
	validation
	obj          *otg.Ospfv2Metric
	marshaller   marshalOspfv2Metric
	unMarshaller unMarshalOspfv2Metric
}

func NewOspfv2Metric() Ospfv2Metric {
	obj := ospfv2Metric{obj: &otg.Ospfv2Metric{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2Metric) msg() *otg.Ospfv2Metric {
	return obj.obj
}

func (obj *ospfv2Metric) setMsg(msg *otg.Ospfv2Metric) Ospfv2Metric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2Metric struct {
	obj *ospfv2Metric
}

type marshalOspfv2Metric interface {
	// ToProto marshals Ospfv2Metric to protobuf object *otg.Ospfv2Metric
	ToProto() (*otg.Ospfv2Metric, error)
	// ToPbText marshals Ospfv2Metric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2Metric to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2Metric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2Metric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2Metric struct {
	obj *ospfv2Metric
}

type unMarshalOspfv2Metric interface {
	// FromProto unmarshals Ospfv2Metric from protobuf object *otg.Ospfv2Metric
	FromProto(msg *otg.Ospfv2Metric) (Ospfv2Metric, error)
	// FromPbText unmarshals Ospfv2Metric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2Metric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2Metric from JSON text
	FromJson(value string) error
}

func (obj *ospfv2Metric) Marshal() marshalOspfv2Metric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2Metric{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2Metric) Unmarshal() unMarshalOspfv2Metric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2Metric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2Metric) ToProto() (*otg.Ospfv2Metric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2Metric) FromProto(msg *otg.Ospfv2Metric) (Ospfv2Metric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2Metric) ToPbText() (string, error) {
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

func (m *unMarshalospfv2Metric) FromPbText(value string) error {
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

func (m *marshalospfv2Metric) ToYaml() (string, error) {
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

func (m *unMarshalospfv2Metric) FromYaml(value string) error {
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

func (m *marshalospfv2Metric) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2Metric) ToJson() (string, error) {
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

func (m *unMarshalospfv2Metric) FromJson(value string) error {
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

func (obj *ospfv2Metric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2Metric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2Metric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2Metric) Clone() (Ospfv2Metric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2Metric()
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

// Ospfv2Metric is oSPFv2 per router statistics information.
type Ospfv2Metric interface {
	Validation
	// msg marshals Ospfv2Metric to protobuf object *otg.Ospfv2Metric
	// and doesn't set defaults
	msg() *otg.Ospfv2Metric
	// setMsg unmarshals Ospfv2Metric from protobuf object *otg.Ospfv2Metric
	// and doesn't set defaults
	setMsg(*otg.Ospfv2Metric) Ospfv2Metric
	// provides marshal interface
	Marshal() marshalOspfv2Metric
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2Metric
	// validate validates Ospfv2Metric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2Metric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Ospfv2Metric.
	Name() string
	// SetName assigns string provided by user to Ospfv2Metric
	SetName(value string) Ospfv2Metric
	// HasName checks if Name has been set in Ospfv2Metric
	HasName() bool
	// FullStateCount returns uint64, set in Ospfv2Metric.
	FullStateCount() uint64
	// SetFullStateCount assigns uint64 provided by user to Ospfv2Metric
	SetFullStateCount(value uint64) Ospfv2Metric
	// HasFullStateCount checks if FullStateCount has been set in Ospfv2Metric
	HasFullStateCount() bool
	// DownStateCount returns uint64, set in Ospfv2Metric.
	DownStateCount() uint64
	// SetDownStateCount assigns uint64 provided by user to Ospfv2Metric
	SetDownStateCount(value uint64) Ospfv2Metric
	// HasDownStateCount checks if DownStateCount has been set in Ospfv2Metric
	HasDownStateCount() bool
	// SessionsFlap returns uint64, set in Ospfv2Metric.
	SessionsFlap() uint64
	// SetSessionsFlap assigns uint64 provided by user to Ospfv2Metric
	SetSessionsFlap(value uint64) Ospfv2Metric
	// HasSessionsFlap checks if SessionsFlap has been set in Ospfv2Metric
	HasSessionsFlap() bool
	// HellosSent returns uint64, set in Ospfv2Metric.
	HellosSent() uint64
	// SetHellosSent assigns uint64 provided by user to Ospfv2Metric
	SetHellosSent(value uint64) Ospfv2Metric
	// HasHellosSent checks if HellosSent has been set in Ospfv2Metric
	HasHellosSent() bool
	// HellosReceived returns uint64, set in Ospfv2Metric.
	HellosReceived() uint64
	// SetHellosReceived assigns uint64 provided by user to Ospfv2Metric
	SetHellosReceived(value uint64) Ospfv2Metric
	// HasHellosReceived checks if HellosReceived has been set in Ospfv2Metric
	HasHellosReceived() bool
	// DbdSent returns uint64, set in Ospfv2Metric.
	DbdSent() uint64
	// SetDbdSent assigns uint64 provided by user to Ospfv2Metric
	SetDbdSent(value uint64) Ospfv2Metric
	// HasDbdSent checks if DbdSent has been set in Ospfv2Metric
	HasDbdSent() bool
	// DbdReceived returns uint64, set in Ospfv2Metric.
	DbdReceived() uint64
	// SetDbdReceived assigns uint64 provided by user to Ospfv2Metric
	SetDbdReceived(value uint64) Ospfv2Metric
	// HasDbdReceived checks if DbdReceived has been set in Ospfv2Metric
	HasDbdReceived() bool
	// LsRequestSent returns uint64, set in Ospfv2Metric.
	LsRequestSent() uint64
	// SetLsRequestSent assigns uint64 provided by user to Ospfv2Metric
	SetLsRequestSent(value uint64) Ospfv2Metric
	// HasLsRequestSent checks if LsRequestSent has been set in Ospfv2Metric
	HasLsRequestSent() bool
	// LsRequestReceived returns uint64, set in Ospfv2Metric.
	LsRequestReceived() uint64
	// SetLsRequestReceived assigns uint64 provided by user to Ospfv2Metric
	SetLsRequestReceived(value uint64) Ospfv2Metric
	// HasLsRequestReceived checks if LsRequestReceived has been set in Ospfv2Metric
	HasLsRequestReceived() bool
	// LsUpdateSent returns uint64, set in Ospfv2Metric.
	LsUpdateSent() uint64
	// SetLsUpdateSent assigns uint64 provided by user to Ospfv2Metric
	SetLsUpdateSent(value uint64) Ospfv2Metric
	// HasLsUpdateSent checks if LsUpdateSent has been set in Ospfv2Metric
	HasLsUpdateSent() bool
	// LsUpdateReceived returns uint64, set in Ospfv2Metric.
	LsUpdateReceived() uint64
	// SetLsUpdateReceived assigns uint64 provided by user to Ospfv2Metric
	SetLsUpdateReceived(value uint64) Ospfv2Metric
	// HasLsUpdateReceived checks if LsUpdateReceived has been set in Ospfv2Metric
	HasLsUpdateReceived() bool
	// LsAckSent returns uint64, set in Ospfv2Metric.
	LsAckSent() uint64
	// SetLsAckSent assigns uint64 provided by user to Ospfv2Metric
	SetLsAckSent(value uint64) Ospfv2Metric
	// HasLsAckSent checks if LsAckSent has been set in Ospfv2Metric
	HasLsAckSent() bool
	// LsAckReceived returns uint64, set in Ospfv2Metric.
	LsAckReceived() uint64
	// SetLsAckReceived assigns uint64 provided by user to Ospfv2Metric
	SetLsAckReceived(value uint64) Ospfv2Metric
	// HasLsAckReceived checks if LsAckReceived has been set in Ospfv2Metric
	HasLsAckReceived() bool
	// LsaSent returns uint64, set in Ospfv2Metric.
	LsaSent() uint64
	// SetLsaSent assigns uint64 provided by user to Ospfv2Metric
	SetLsaSent(value uint64) Ospfv2Metric
	// HasLsaSent checks if LsaSent has been set in Ospfv2Metric
	HasLsaSent() bool
	// LsaReceived returns uint64, set in Ospfv2Metric.
	LsaReceived() uint64
	// SetLsaReceived assigns uint64 provided by user to Ospfv2Metric
	SetLsaReceived(value uint64) Ospfv2Metric
	// HasLsaReceived checks if LsaReceived has been set in Ospfv2Metric
	HasLsaReceived() bool
	// LsaAckSent returns uint64, set in Ospfv2Metric.
	LsaAckSent() uint64
	// SetLsaAckSent assigns uint64 provided by user to Ospfv2Metric
	SetLsaAckSent(value uint64) Ospfv2Metric
	// HasLsaAckSent checks if LsaAckSent has been set in Ospfv2Metric
	HasLsaAckSent() bool
	// LsaAckReceived returns uint64, set in Ospfv2Metric.
	LsaAckReceived() uint64
	// SetLsaAckReceived assigns uint64 provided by user to Ospfv2Metric
	SetLsaAckReceived(value uint64) Ospfv2Metric
	// HasLsaAckReceived checks if LsaAckReceived has been set in Ospfv2Metric
	HasLsaAckReceived() bool
	// RouterLsaSent returns uint64, set in Ospfv2Metric.
	RouterLsaSent() uint64
	// SetRouterLsaSent assigns uint64 provided by user to Ospfv2Metric
	SetRouterLsaSent(value uint64) Ospfv2Metric
	// HasRouterLsaSent checks if RouterLsaSent has been set in Ospfv2Metric
	HasRouterLsaSent() bool
	// RouterLsaReceived returns uint64, set in Ospfv2Metric.
	RouterLsaReceived() uint64
	// SetRouterLsaReceived assigns uint64 provided by user to Ospfv2Metric
	SetRouterLsaReceived(value uint64) Ospfv2Metric
	// HasRouterLsaReceived checks if RouterLsaReceived has been set in Ospfv2Metric
	HasRouterLsaReceived() bool
	// NetworkLsaSent returns uint64, set in Ospfv2Metric.
	NetworkLsaSent() uint64
	// SetNetworkLsaSent assigns uint64 provided by user to Ospfv2Metric
	SetNetworkLsaSent(value uint64) Ospfv2Metric
	// HasNetworkLsaSent checks if NetworkLsaSent has been set in Ospfv2Metric
	HasNetworkLsaSent() bool
	// NetworkLsaReceived returns uint64, set in Ospfv2Metric.
	NetworkLsaReceived() uint64
	// SetNetworkLsaReceived assigns uint64 provided by user to Ospfv2Metric
	SetNetworkLsaReceived(value uint64) Ospfv2Metric
	// HasNetworkLsaReceived checks if NetworkLsaReceived has been set in Ospfv2Metric
	HasNetworkLsaReceived() bool
	// SummaryLsaSent returns uint64, set in Ospfv2Metric.
	SummaryLsaSent() uint64
	// SetSummaryLsaSent assigns uint64 provided by user to Ospfv2Metric
	SetSummaryLsaSent(value uint64) Ospfv2Metric
	// HasSummaryLsaSent checks if SummaryLsaSent has been set in Ospfv2Metric
	HasSummaryLsaSent() bool
	// SummaryLsaReceived returns uint64, set in Ospfv2Metric.
	SummaryLsaReceived() uint64
	// SetSummaryLsaReceived assigns uint64 provided by user to Ospfv2Metric
	SetSummaryLsaReceived(value uint64) Ospfv2Metric
	// HasSummaryLsaReceived checks if SummaryLsaReceived has been set in Ospfv2Metric
	HasSummaryLsaReceived() bool
	// ExternalLsaSent returns uint64, set in Ospfv2Metric.
	ExternalLsaSent() uint64
	// SetExternalLsaSent assigns uint64 provided by user to Ospfv2Metric
	SetExternalLsaSent(value uint64) Ospfv2Metric
	// HasExternalLsaSent checks if ExternalLsaSent has been set in Ospfv2Metric
	HasExternalLsaSent() bool
	// ExternalLsaReceived returns uint64, set in Ospfv2Metric.
	ExternalLsaReceived() uint64
	// SetExternalLsaReceived assigns uint64 provided by user to Ospfv2Metric
	SetExternalLsaReceived(value uint64) Ospfv2Metric
	// HasExternalLsaReceived checks if ExternalLsaReceived has been set in Ospfv2Metric
	HasExternalLsaReceived() bool
	// NssaLsaSent returns uint64, set in Ospfv2Metric.
	NssaLsaSent() uint64
	// SetNssaLsaSent assigns uint64 provided by user to Ospfv2Metric
	SetNssaLsaSent(value uint64) Ospfv2Metric
	// HasNssaLsaSent checks if NssaLsaSent has been set in Ospfv2Metric
	HasNssaLsaSent() bool
	// NssaLsaReceived returns uint64, set in Ospfv2Metric.
	NssaLsaReceived() uint64
	// SetNssaLsaReceived assigns uint64 provided by user to Ospfv2Metric
	SetNssaLsaReceived(value uint64) Ospfv2Metric
	// HasNssaLsaReceived checks if NssaLsaReceived has been set in Ospfv2Metric
	HasNssaLsaReceived() bool
	// OpaqueLocalSent returns uint64, set in Ospfv2Metric.
	OpaqueLocalSent() uint64
	// SetOpaqueLocalSent assigns uint64 provided by user to Ospfv2Metric
	SetOpaqueLocalSent(value uint64) Ospfv2Metric
	// HasOpaqueLocalSent checks if OpaqueLocalSent has been set in Ospfv2Metric
	HasOpaqueLocalSent() bool
	// OpaqueLocalReceived returns uint64, set in Ospfv2Metric.
	OpaqueLocalReceived() uint64
	// SetOpaqueLocalReceived assigns uint64 provided by user to Ospfv2Metric
	SetOpaqueLocalReceived(value uint64) Ospfv2Metric
	// HasOpaqueLocalReceived checks if OpaqueLocalReceived has been set in Ospfv2Metric
	HasOpaqueLocalReceived() bool
	// OpaqueAreaSent returns uint64, set in Ospfv2Metric.
	OpaqueAreaSent() uint64
	// SetOpaqueAreaSent assigns uint64 provided by user to Ospfv2Metric
	SetOpaqueAreaSent(value uint64) Ospfv2Metric
	// HasOpaqueAreaSent checks if OpaqueAreaSent has been set in Ospfv2Metric
	HasOpaqueAreaSent() bool
	// OpaqueAreaReceived returns uint64, set in Ospfv2Metric.
	OpaqueAreaReceived() uint64
	// SetOpaqueAreaReceived assigns uint64 provided by user to Ospfv2Metric
	SetOpaqueAreaReceived(value uint64) Ospfv2Metric
	// HasOpaqueAreaReceived checks if OpaqueAreaReceived has been set in Ospfv2Metric
	HasOpaqueAreaReceived() bool
	// OpaqueDomainSent returns uint64, set in Ospfv2Metric.
	OpaqueDomainSent() uint64
	// SetOpaqueDomainSent assigns uint64 provided by user to Ospfv2Metric
	SetOpaqueDomainSent(value uint64) Ospfv2Metric
	// HasOpaqueDomainSent checks if OpaqueDomainSent has been set in Ospfv2Metric
	HasOpaqueDomainSent() bool
	// OpaqueDomainReceived returns uint64, set in Ospfv2Metric.
	OpaqueDomainReceived() uint64
	// SetOpaqueDomainReceived assigns uint64 provided by user to Ospfv2Metric
	SetOpaqueDomainReceived(value uint64) Ospfv2Metric
	// HasOpaqueDomainReceived checks if OpaqueDomainReceived has been set in Ospfv2Metric
	HasOpaqueDomainReceived() bool
}

// The name of a configured OSPFv2 router.
// Name returns a string
func (obj *ospfv2Metric) Name() string {

	return *obj.obj.Name

}

// The name of a configured OSPFv2 router.
// Name returns a string
func (obj *ospfv2Metric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured OSPFv2 router.
// SetName sets the string value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetName(value string) Ospfv2Metric {

	obj.obj.Name = &value
	return obj
}

// The number of OSPFv2 sessions in up state.
// FullStateCount returns a uint64
func (obj *ospfv2Metric) FullStateCount() uint64 {

	return *obj.obj.FullStateCount

}

// The number of OSPFv2 sessions in up state.
// FullStateCount returns a uint64
func (obj *ospfv2Metric) HasFullStateCount() bool {
	return obj.obj.FullStateCount != nil
}

// The number of OSPFv2 sessions in up state.
// SetFullStateCount sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetFullStateCount(value uint64) Ospfv2Metric {

	obj.obj.FullStateCount = &value
	return obj
}

// The number of OSPFv2 sessions in down state.
// DownStateCount returns a uint64
func (obj *ospfv2Metric) DownStateCount() uint64 {

	return *obj.obj.DownStateCount

}

// The number of OSPFv2 sessions in down state.
// DownStateCount returns a uint64
func (obj *ospfv2Metric) HasDownStateCount() bool {
	return obj.obj.DownStateCount != nil
}

// The number of OSPFv2 sessions in down state.
// SetDownStateCount sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetDownStateCount(value uint64) Ospfv2Metric {

	obj.obj.DownStateCount = &value
	return obj
}

// The number of change of OSPFv2 sessions from up to down state.
// SessionsFlap returns a uint64
func (obj *ospfv2Metric) SessionsFlap() uint64 {

	return *obj.obj.SessionsFlap

}

// The number of change of OSPFv2 sessions from up to down state.
// SessionsFlap returns a uint64
func (obj *ospfv2Metric) HasSessionsFlap() bool {
	return obj.obj.SessionsFlap != nil
}

// The number of change of OSPFv2 sessions from up to down state.
// SetSessionsFlap sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetSessionsFlap(value uint64) Ospfv2Metric {

	obj.obj.SessionsFlap = &value
	return obj
}

// The number of OSPFv2 Hello messages transmitted.
// HellosSent returns a uint64
func (obj *ospfv2Metric) HellosSent() uint64 {

	return *obj.obj.HellosSent

}

// The number of OSPFv2 Hello messages transmitted.
// HellosSent returns a uint64
func (obj *ospfv2Metric) HasHellosSent() bool {
	return obj.obj.HellosSent != nil
}

// The number of OSPFv2 Hello messages transmitted.
// SetHellosSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetHellosSent(value uint64) Ospfv2Metric {

	obj.obj.HellosSent = &value
	return obj
}

// The number of OSPFv2 Hello messages received.
// HellosReceived returns a uint64
func (obj *ospfv2Metric) HellosReceived() uint64 {

	return *obj.obj.HellosReceived

}

// The number of OSPFv2 Hello messages received.
// HellosReceived returns a uint64
func (obj *ospfv2Metric) HasHellosReceived() bool {
	return obj.obj.HellosReceived != nil
}

// The number of OSPFv2 Hello messages received.
// SetHellosReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetHellosReceived(value uint64) Ospfv2Metric {

	obj.obj.HellosReceived = &value
	return obj
}

// The number of OSPFv2 Database Description (DBD) messages transmitted.
// DbdSent returns a uint64
func (obj *ospfv2Metric) DbdSent() uint64 {

	return *obj.obj.DbdSent

}

// The number of OSPFv2 Database Description (DBD) messages transmitted.
// DbdSent returns a uint64
func (obj *ospfv2Metric) HasDbdSent() bool {
	return obj.obj.DbdSent != nil
}

// The number of OSPFv2 Database Description (DBD) messages transmitted.
// SetDbdSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetDbdSent(value uint64) Ospfv2Metric {

	obj.obj.DbdSent = &value
	return obj
}

// The number of OSPFv2 Database Description (DBD) messages received.
// DbdReceived returns a uint64
func (obj *ospfv2Metric) DbdReceived() uint64 {

	return *obj.obj.DbdReceived

}

// The number of OSPFv2 Database Description (DBD) messages received.
// DbdReceived returns a uint64
func (obj *ospfv2Metric) HasDbdReceived() bool {
	return obj.obj.DbdReceived != nil
}

// The number of OSPFv2 Database Description (DBD) messages received.
// SetDbdReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetDbdReceived(value uint64) Ospfv2Metric {

	obj.obj.DbdReceived = &value
	return obj
}

// The number of OSPFv2 LinkState (LS) Request messages transmitted.
// LsRequestSent returns a uint64
func (obj *ospfv2Metric) LsRequestSent() uint64 {

	return *obj.obj.LsRequestSent

}

// The number of OSPFv2 LinkState (LS) Request messages transmitted.
// LsRequestSent returns a uint64
func (obj *ospfv2Metric) HasLsRequestSent() bool {
	return obj.obj.LsRequestSent != nil
}

// The number of OSPFv2 LinkState (LS) Request messages transmitted.
// SetLsRequestSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsRequestSent(value uint64) Ospfv2Metric {

	obj.obj.LsRequestSent = &value
	return obj
}

// The number of OSPFv2 LinkState (LS) Request messages received.
// LsRequestReceived returns a uint64
func (obj *ospfv2Metric) LsRequestReceived() uint64 {

	return *obj.obj.LsRequestReceived

}

// The number of OSPFv2 LinkState (LS) Request messages received.
// LsRequestReceived returns a uint64
func (obj *ospfv2Metric) HasLsRequestReceived() bool {
	return obj.obj.LsRequestReceived != nil
}

// The number of OSPFv2 LinkState (LS) Request messages received.
// SetLsRequestReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsRequestReceived(value uint64) Ospfv2Metric {

	obj.obj.LsRequestReceived = &value
	return obj
}

// The number of OSPFv2 LinkState (LS) Update messages transmitted.
// LsUpdateSent returns a uint64
func (obj *ospfv2Metric) LsUpdateSent() uint64 {

	return *obj.obj.LsUpdateSent

}

// The number of OSPFv2 LinkState (LS) Update messages transmitted.
// LsUpdateSent returns a uint64
func (obj *ospfv2Metric) HasLsUpdateSent() bool {
	return obj.obj.LsUpdateSent != nil
}

// The number of OSPFv2 LinkState (LS) Update messages transmitted.
// SetLsUpdateSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsUpdateSent(value uint64) Ospfv2Metric {

	obj.obj.LsUpdateSent = &value
	return obj
}

// The number of OSPFv2 LinkState (LS) Update messages received.
// LsUpdateReceived returns a uint64
func (obj *ospfv2Metric) LsUpdateReceived() uint64 {

	return *obj.obj.LsUpdateReceived

}

// The number of OSPFv2 LinkState (LS) Update messages received.
// LsUpdateReceived returns a uint64
func (obj *ospfv2Metric) HasLsUpdateReceived() bool {
	return obj.obj.LsUpdateReceived != nil
}

// The number of OSPFv2 LinkState (LS) Update messages received.
// SetLsUpdateReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsUpdateReceived(value uint64) Ospfv2Metric {

	obj.obj.LsUpdateReceived = &value
	return obj
}

// The number of OSPFv2 LinkState (LS) Acknowledgement messages transmitted.
// LsAckSent returns a uint64
func (obj *ospfv2Metric) LsAckSent() uint64 {

	return *obj.obj.LsAckSent

}

// The number of OSPFv2 LinkState (LS) Acknowledgement messages transmitted.
// LsAckSent returns a uint64
func (obj *ospfv2Metric) HasLsAckSent() bool {
	return obj.obj.LsAckSent != nil
}

// The number of OSPFv2 LinkState (LS) Acknowledgement messages transmitted.
// SetLsAckSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsAckSent(value uint64) Ospfv2Metric {

	obj.obj.LsAckSent = &value
	return obj
}

// The number of OSPFv2 LinkState (LS) Acknowledgement messages received.
// LsAckReceived returns a uint64
func (obj *ospfv2Metric) LsAckReceived() uint64 {

	return *obj.obj.LsAckReceived

}

// The number of OSPFv2 LinkState (LS) Acknowledgement messages received.
// LsAckReceived returns a uint64
func (obj *ospfv2Metric) HasLsAckReceived() bool {
	return obj.obj.LsAckReceived != nil
}

// The number of OSPFv2 LinkState (LS) Acknowledgement messages received.
// SetLsAckReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsAckReceived(value uint64) Ospfv2Metric {

	obj.obj.LsAckReceived = &value
	return obj
}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages transmitted.
// LsaSent returns a uint64
func (obj *ospfv2Metric) LsaSent() uint64 {

	return *obj.obj.LsaSent

}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages transmitted.
// LsaSent returns a uint64
func (obj *ospfv2Metric) HasLsaSent() bool {
	return obj.obj.LsaSent != nil
}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages transmitted.
// SetLsaSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsaSent(value uint64) Ospfv2Metric {

	obj.obj.LsaSent = &value
	return obj
}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages received.
// LsaReceived returns a uint64
func (obj *ospfv2Metric) LsaReceived() uint64 {

	return *obj.obj.LsaReceived

}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages received.
// LsaReceived returns a uint64
func (obj *ospfv2Metric) HasLsaReceived() bool {
	return obj.obj.LsaReceived != nil
}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages received.
// SetLsaReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsaReceived(value uint64) Ospfv2Metric {

	obj.obj.LsaReceived = &value
	return obj
}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages acknowledged.
// LsaAckSent returns a uint64
func (obj *ospfv2Metric) LsaAckSent() uint64 {

	return *obj.obj.LsaAckSent

}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages acknowledged.
// LsaAckSent returns a uint64
func (obj *ospfv2Metric) HasLsaAckSent() bool {
	return obj.obj.LsaAckSent != nil
}

// The total number of OSPFv2 LinkState Advertisement (LSA) messages acknowledged.
// SetLsaAckSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsaAckSent(value uint64) Ospfv2Metric {

	obj.obj.LsaAckSent = &value
	return obj
}

// The total number of OSPFv2 LinkState Advertisement (LSA) acknowledge messages received .
// LsaAckReceived returns a uint64
func (obj *ospfv2Metric) LsaAckReceived() uint64 {

	return *obj.obj.LsaAckReceived

}

// The total number of OSPFv2 LinkState Advertisement (LSA) acknowledge messages received .
// LsaAckReceived returns a uint64
func (obj *ospfv2Metric) HasLsaAckReceived() bool {
	return obj.obj.LsaAckReceived != nil
}

// The total number of OSPFv2 LinkState Advertisement (LSA) acknowledge messages received .
// SetLsaAckReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetLsaAckReceived(value uint64) Ospfv2Metric {

	obj.obj.LsaAckReceived = &value
	return obj
}

// The number of OSPFv2 Router (Type 1) LSAs transmitted.
// RouterLsaSent returns a uint64
func (obj *ospfv2Metric) RouterLsaSent() uint64 {

	return *obj.obj.RouterLsaSent

}

// The number of OSPFv2 Router (Type 1) LSAs transmitted.
// RouterLsaSent returns a uint64
func (obj *ospfv2Metric) HasRouterLsaSent() bool {
	return obj.obj.RouterLsaSent != nil
}

// The number of OSPFv2 Router (Type 1) LSAs transmitted.
// SetRouterLsaSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetRouterLsaSent(value uint64) Ospfv2Metric {

	obj.obj.RouterLsaSent = &value
	return obj
}

// The number of OSPFv2 Router (Type 1) LSAs received.
// RouterLsaReceived returns a uint64
func (obj *ospfv2Metric) RouterLsaReceived() uint64 {

	return *obj.obj.RouterLsaReceived

}

// The number of OSPFv2 Router (Type 1) LSAs received.
// RouterLsaReceived returns a uint64
func (obj *ospfv2Metric) HasRouterLsaReceived() bool {
	return obj.obj.RouterLsaReceived != nil
}

// The number of OSPFv2 Router (Type 1) LSAs received.
// SetRouterLsaReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetRouterLsaReceived(value uint64) Ospfv2Metric {

	obj.obj.RouterLsaReceived = &value
	return obj
}

// The number of OSPFv2 Network (Type 2) LSAs transmitted.
// NetworkLsaSent returns a uint64
func (obj *ospfv2Metric) NetworkLsaSent() uint64 {

	return *obj.obj.NetworkLsaSent

}

// The number of OSPFv2 Network (Type 2) LSAs transmitted.
// NetworkLsaSent returns a uint64
func (obj *ospfv2Metric) HasNetworkLsaSent() bool {
	return obj.obj.NetworkLsaSent != nil
}

// The number of OSPFv2 Network (Type 2) LSAs transmitted.
// SetNetworkLsaSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetNetworkLsaSent(value uint64) Ospfv2Metric {

	obj.obj.NetworkLsaSent = &value
	return obj
}

// The number of OSPFv2 Network (Type 2) LSAs transmitted.
// NetworkLsaReceived returns a uint64
func (obj *ospfv2Metric) NetworkLsaReceived() uint64 {

	return *obj.obj.NetworkLsaReceived

}

// The number of OSPFv2 Network (Type 2) LSAs transmitted.
// NetworkLsaReceived returns a uint64
func (obj *ospfv2Metric) HasNetworkLsaReceived() bool {
	return obj.obj.NetworkLsaReceived != nil
}

// The number of OSPFv2 Network (Type 2) LSAs transmitted.
// SetNetworkLsaReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetNetworkLsaReceived(value uint64) Ospfv2Metric {

	obj.obj.NetworkLsaReceived = &value
	return obj
}

// The number of OSPFv2 Summary IP (Type 3) LSAs transmitted.
// SummaryLsaSent returns a uint64
func (obj *ospfv2Metric) SummaryLsaSent() uint64 {

	return *obj.obj.SummaryLsaSent

}

// The number of OSPFv2 Summary IP (Type 3) LSAs transmitted.
// SummaryLsaSent returns a uint64
func (obj *ospfv2Metric) HasSummaryLsaSent() bool {
	return obj.obj.SummaryLsaSent != nil
}

// The number of OSPFv2 Summary IP (Type 3) LSAs transmitted.
// SetSummaryLsaSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetSummaryLsaSent(value uint64) Ospfv2Metric {

	obj.obj.SummaryLsaSent = &value
	return obj
}

// The number of OSPFv2 Summary IP (Type 3) LSA received.
// SummaryLsaReceived returns a uint64
func (obj *ospfv2Metric) SummaryLsaReceived() uint64 {

	return *obj.obj.SummaryLsaReceived

}

// The number of OSPFv2 Summary IP (Type 3) LSA received.
// SummaryLsaReceived returns a uint64
func (obj *ospfv2Metric) HasSummaryLsaReceived() bool {
	return obj.obj.SummaryLsaReceived != nil
}

// The number of OSPFv2 Summary IP (Type 3) LSA received.
// SetSummaryLsaReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetSummaryLsaReceived(value uint64) Ospfv2Metric {

	obj.obj.SummaryLsaReceived = &value
	return obj
}

// The number of OSPFv2 External (Type 5) LSAs transmitted.
// ExternalLsaSent returns a uint64
func (obj *ospfv2Metric) ExternalLsaSent() uint64 {

	return *obj.obj.ExternalLsaSent

}

// The number of OSPFv2 External (Type 5) LSAs transmitted.
// ExternalLsaSent returns a uint64
func (obj *ospfv2Metric) HasExternalLsaSent() bool {
	return obj.obj.ExternalLsaSent != nil
}

// The number of OSPFv2 External (Type 5) LSAs transmitted.
// SetExternalLsaSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetExternalLsaSent(value uint64) Ospfv2Metric {

	obj.obj.ExternalLsaSent = &value
	return obj
}

// The number of OSPFv2 External (Type 5) LSAs received.
// ExternalLsaReceived returns a uint64
func (obj *ospfv2Metric) ExternalLsaReceived() uint64 {

	return *obj.obj.ExternalLsaReceived

}

// The number of OSPFv2 External (Type 5) LSAs received.
// ExternalLsaReceived returns a uint64
func (obj *ospfv2Metric) HasExternalLsaReceived() bool {
	return obj.obj.ExternalLsaReceived != nil
}

// The number of OSPFv2 External (Type 5) LSAs received.
// SetExternalLsaReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetExternalLsaReceived(value uint64) Ospfv2Metric {

	obj.obj.ExternalLsaReceived = &value
	return obj
}

// The number of OSPFv2 NSSA (Type 7) LSAs transmitted.
// NssaLsaSent returns a uint64
func (obj *ospfv2Metric) NssaLsaSent() uint64 {

	return *obj.obj.NssaLsaSent

}

// The number of OSPFv2 NSSA (Type 7) LSAs transmitted.
// NssaLsaSent returns a uint64
func (obj *ospfv2Metric) HasNssaLsaSent() bool {
	return obj.obj.NssaLsaSent != nil
}

// The number of OSPFv2 NSSA (Type 7) LSAs transmitted.
// SetNssaLsaSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetNssaLsaSent(value uint64) Ospfv2Metric {

	obj.obj.NssaLsaSent = &value
	return obj
}

// The number of OSPFv2 NSSA (Type 7) LSAs received.
// NssaLsaReceived returns a uint64
func (obj *ospfv2Metric) NssaLsaReceived() uint64 {

	return *obj.obj.NssaLsaReceived

}

// The number of OSPFv2 NSSA (Type 7) LSAs received.
// NssaLsaReceived returns a uint64
func (obj *ospfv2Metric) HasNssaLsaReceived() bool {
	return obj.obj.NssaLsaReceived != nil
}

// The number of OSPFv2 NSSA (Type 7) LSAs received.
// SetNssaLsaReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetNssaLsaReceived(value uint64) Ospfv2Metric {

	obj.obj.NssaLsaReceived = &value
	return obj
}

// The number of OSPFv2 Opaque Local (Type 9) LSAs transmitted.
// OpaqueLocalSent returns a uint64
func (obj *ospfv2Metric) OpaqueLocalSent() uint64 {

	return *obj.obj.OpaqueLocalSent

}

// The number of OSPFv2 Opaque Local (Type 9) LSAs transmitted.
// OpaqueLocalSent returns a uint64
func (obj *ospfv2Metric) HasOpaqueLocalSent() bool {
	return obj.obj.OpaqueLocalSent != nil
}

// The number of OSPFv2 Opaque Local (Type 9) LSAs transmitted.
// SetOpaqueLocalSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetOpaqueLocalSent(value uint64) Ospfv2Metric {

	obj.obj.OpaqueLocalSent = &value
	return obj
}

// The number of OSPFv2 Opaque Local (Type 9) LSAs received.
// OpaqueLocalReceived returns a uint64
func (obj *ospfv2Metric) OpaqueLocalReceived() uint64 {

	return *obj.obj.OpaqueLocalReceived

}

// The number of OSPFv2 Opaque Local (Type 9) LSAs received.
// OpaqueLocalReceived returns a uint64
func (obj *ospfv2Metric) HasOpaqueLocalReceived() bool {
	return obj.obj.OpaqueLocalReceived != nil
}

// The number of OSPFv2 Opaque Local (Type 9) LSAs received.
// SetOpaqueLocalReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetOpaqueLocalReceived(value uint64) Ospfv2Metric {

	obj.obj.OpaqueLocalReceived = &value
	return obj
}

// The number of OSPF Opaque Area (Type 10) LSAs transmitted.
// OpaqueAreaSent returns a uint64
func (obj *ospfv2Metric) OpaqueAreaSent() uint64 {

	return *obj.obj.OpaqueAreaSent

}

// The number of OSPF Opaque Area (Type 10) LSAs transmitted.
// OpaqueAreaSent returns a uint64
func (obj *ospfv2Metric) HasOpaqueAreaSent() bool {
	return obj.obj.OpaqueAreaSent != nil
}

// The number of OSPF Opaque Area (Type 10) LSAs transmitted.
// SetOpaqueAreaSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetOpaqueAreaSent(value uint64) Ospfv2Metric {

	obj.obj.OpaqueAreaSent = &value
	return obj
}

// The number of OSPFv2 Opaque Area (Type 10) LSAs received.
// OpaqueAreaReceived returns a uint64
func (obj *ospfv2Metric) OpaqueAreaReceived() uint64 {

	return *obj.obj.OpaqueAreaReceived

}

// The number of OSPFv2 Opaque Area (Type 10) LSAs received.
// OpaqueAreaReceived returns a uint64
func (obj *ospfv2Metric) HasOpaqueAreaReceived() bool {
	return obj.obj.OpaqueAreaReceived != nil
}

// The number of OSPFv2 Opaque Area (Type 10) LSAs received.
// SetOpaqueAreaReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetOpaqueAreaReceived(value uint64) Ospfv2Metric {

	obj.obj.OpaqueAreaReceived = &value
	return obj
}

// The number of OSPFv2 Opaque Domain (Type 11) LSAs transmitted.
// OpaqueDomainSent returns a uint64
func (obj *ospfv2Metric) OpaqueDomainSent() uint64 {

	return *obj.obj.OpaqueDomainSent

}

// The number of OSPFv2 Opaque Domain (Type 11) LSAs transmitted.
// OpaqueDomainSent returns a uint64
func (obj *ospfv2Metric) HasOpaqueDomainSent() bool {
	return obj.obj.OpaqueDomainSent != nil
}

// The number of OSPFv2 Opaque Domain (Type 11) LSAs transmitted.
// SetOpaqueDomainSent sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetOpaqueDomainSent(value uint64) Ospfv2Metric {

	obj.obj.OpaqueDomainSent = &value
	return obj
}

// The number of OSPFv2 Opaque Domain (Type 11) LSAs received.
// OpaqueDomainReceived returns a uint64
func (obj *ospfv2Metric) OpaqueDomainReceived() uint64 {

	return *obj.obj.OpaqueDomainReceived

}

// The number of OSPFv2 Opaque Domain (Type 11) LSAs received.
// OpaqueDomainReceived returns a uint64
func (obj *ospfv2Metric) HasOpaqueDomainReceived() bool {
	return obj.obj.OpaqueDomainReceived != nil
}

// The number of OSPFv2 Opaque Domain (Type 11) LSAs received.
// SetOpaqueDomainReceived sets the uint64 value in the Ospfv2Metric object
func (obj *ospfv2Metric) SetOpaqueDomainReceived(value uint64) Ospfv2Metric {

	obj.obj.OpaqueDomainReceived = &value
	return obj
}

func (obj *ospfv2Metric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2Metric) setDefault() {

}
