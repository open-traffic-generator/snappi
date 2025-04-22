package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3Metric *****
type ospfv3Metric struct {
	validation
	obj          *otg.Ospfv3Metric
	marshaller   marshalOspfv3Metric
	unMarshaller unMarshalOspfv3Metric
}

func NewOspfv3Metric() Ospfv3Metric {
	obj := ospfv3Metric{obj: &otg.Ospfv3Metric{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3Metric) msg() *otg.Ospfv3Metric {
	return obj.obj
}

func (obj *ospfv3Metric) setMsg(msg *otg.Ospfv3Metric) Ospfv3Metric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3Metric struct {
	obj *ospfv3Metric
}

type marshalOspfv3Metric interface {
	// ToProto marshals Ospfv3Metric to protobuf object *otg.Ospfv3Metric
	ToProto() (*otg.Ospfv3Metric, error)
	// ToPbText marshals Ospfv3Metric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3Metric to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3Metric to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3Metric struct {
	obj *ospfv3Metric
}

type unMarshalOspfv3Metric interface {
	// FromProto unmarshals Ospfv3Metric from protobuf object *otg.Ospfv3Metric
	FromProto(msg *otg.Ospfv3Metric) (Ospfv3Metric, error)
	// FromPbText unmarshals Ospfv3Metric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3Metric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3Metric from JSON text
	FromJson(value string) error
}

func (obj *ospfv3Metric) Marshal() marshalOspfv3Metric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3Metric{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3Metric) Unmarshal() unMarshalOspfv3Metric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3Metric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3Metric) ToProto() (*otg.Ospfv3Metric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3Metric) FromProto(msg *otg.Ospfv3Metric) (Ospfv3Metric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3Metric) ToPbText() (string, error) {
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

func (m *unMarshalospfv3Metric) FromPbText(value string) error {
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

func (m *marshalospfv3Metric) ToYaml() (string, error) {
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

func (m *unMarshalospfv3Metric) FromYaml(value string) error {
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

func (m *marshalospfv3Metric) ToJson() (string, error) {
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

func (m *unMarshalospfv3Metric) FromJson(value string) error {
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

func (obj *ospfv3Metric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3Metric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3Metric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3Metric) Clone() (Ospfv3Metric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3Metric()
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

// Ospfv3Metric is oSPFv3 per router statistics information.
type Ospfv3Metric interface {
	Validation
	// msg marshals Ospfv3Metric to protobuf object *otg.Ospfv3Metric
	// and doesn't set defaults
	msg() *otg.Ospfv3Metric
	// setMsg unmarshals Ospfv3Metric from protobuf object *otg.Ospfv3Metric
	// and doesn't set defaults
	setMsg(*otg.Ospfv3Metric) Ospfv3Metric
	// provides marshal interface
	Marshal() marshalOspfv3Metric
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3Metric
	// validate validates Ospfv3Metric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3Metric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Ospfv3Metric.
	Name() string
	// SetName assigns string provided by user to Ospfv3Metric
	SetName(value string) Ospfv3Metric
	// HasName checks if Name has been set in Ospfv3Metric
	HasName() bool
	// FullStateCount returns uint64, set in Ospfv3Metric.
	FullStateCount() uint64
	// SetFullStateCount assigns uint64 provided by user to Ospfv3Metric
	SetFullStateCount(value uint64) Ospfv3Metric
	// HasFullStateCount checks if FullStateCount has been set in Ospfv3Metric
	HasFullStateCount() bool
	// DownStateCount returns uint64, set in Ospfv3Metric.
	DownStateCount() uint64
	// SetDownStateCount assigns uint64 provided by user to Ospfv3Metric
	SetDownStateCount(value uint64) Ospfv3Metric
	// HasDownStateCount checks if DownStateCount has been set in Ospfv3Metric
	HasDownStateCount() bool
	// SessionsFlap returns uint64, set in Ospfv3Metric.
	SessionsFlap() uint64
	// SetSessionsFlap assigns uint64 provided by user to Ospfv3Metric
	SetSessionsFlap(value uint64) Ospfv3Metric
	// HasSessionsFlap checks if SessionsFlap has been set in Ospfv3Metric
	HasSessionsFlap() bool
	// HellosSent returns uint64, set in Ospfv3Metric.
	HellosSent() uint64
	// SetHellosSent assigns uint64 provided by user to Ospfv3Metric
	SetHellosSent(value uint64) Ospfv3Metric
	// HasHellosSent checks if HellosSent has been set in Ospfv3Metric
	HasHellosSent() bool
	// HellosReceived returns uint64, set in Ospfv3Metric.
	HellosReceived() uint64
	// SetHellosReceived assigns uint64 provided by user to Ospfv3Metric
	SetHellosReceived(value uint64) Ospfv3Metric
	// HasHellosReceived checks if HellosReceived has been set in Ospfv3Metric
	HasHellosReceived() bool
	// DbdSent returns uint64, set in Ospfv3Metric.
	DbdSent() uint64
	// SetDbdSent assigns uint64 provided by user to Ospfv3Metric
	SetDbdSent(value uint64) Ospfv3Metric
	// HasDbdSent checks if DbdSent has been set in Ospfv3Metric
	HasDbdSent() bool
	// DbdReceived returns uint64, set in Ospfv3Metric.
	DbdReceived() uint64
	// SetDbdReceived assigns uint64 provided by user to Ospfv3Metric
	SetDbdReceived(value uint64) Ospfv3Metric
	// HasDbdReceived checks if DbdReceived has been set in Ospfv3Metric
	HasDbdReceived() bool
	// LsRequestSent returns uint64, set in Ospfv3Metric.
	LsRequestSent() uint64
	// SetLsRequestSent assigns uint64 provided by user to Ospfv3Metric
	SetLsRequestSent(value uint64) Ospfv3Metric
	// HasLsRequestSent checks if LsRequestSent has been set in Ospfv3Metric
	HasLsRequestSent() bool
	// LsRequestReceived returns uint64, set in Ospfv3Metric.
	LsRequestReceived() uint64
	// SetLsRequestReceived assigns uint64 provided by user to Ospfv3Metric
	SetLsRequestReceived(value uint64) Ospfv3Metric
	// HasLsRequestReceived checks if LsRequestReceived has been set in Ospfv3Metric
	HasLsRequestReceived() bool
	// LsUpdateSent returns uint64, set in Ospfv3Metric.
	LsUpdateSent() uint64
	// SetLsUpdateSent assigns uint64 provided by user to Ospfv3Metric
	SetLsUpdateSent(value uint64) Ospfv3Metric
	// HasLsUpdateSent checks if LsUpdateSent has been set in Ospfv3Metric
	HasLsUpdateSent() bool
	// LsUpdateReceived returns uint64, set in Ospfv3Metric.
	LsUpdateReceived() uint64
	// SetLsUpdateReceived assigns uint64 provided by user to Ospfv3Metric
	SetLsUpdateReceived(value uint64) Ospfv3Metric
	// HasLsUpdateReceived checks if LsUpdateReceived has been set in Ospfv3Metric
	HasLsUpdateReceived() bool
	// LsAckSent returns uint64, set in Ospfv3Metric.
	LsAckSent() uint64
	// SetLsAckSent assigns uint64 provided by user to Ospfv3Metric
	SetLsAckSent(value uint64) Ospfv3Metric
	// HasLsAckSent checks if LsAckSent has been set in Ospfv3Metric
	HasLsAckSent() bool
	// LsAckReceived returns uint64, set in Ospfv3Metric.
	LsAckReceived() uint64
	// SetLsAckReceived assigns uint64 provided by user to Ospfv3Metric
	SetLsAckReceived(value uint64) Ospfv3Metric
	// HasLsAckReceived checks if LsAckReceived has been set in Ospfv3Metric
	HasLsAckReceived() bool
	// LsaSent returns uint64, set in Ospfv3Metric.
	LsaSent() uint64
	// SetLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetLsaSent(value uint64) Ospfv3Metric
	// HasLsaSent checks if LsaSent has been set in Ospfv3Metric
	HasLsaSent() bool
	// LsaReceived returns uint64, set in Ospfv3Metric.
	LsaReceived() uint64
	// SetLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetLsaReceived(value uint64) Ospfv3Metric
	// HasLsaReceived checks if LsaReceived has been set in Ospfv3Metric
	HasLsaReceived() bool
	// RouterLsaSent returns uint64, set in Ospfv3Metric.
	RouterLsaSent() uint64
	// SetRouterLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetRouterLsaSent(value uint64) Ospfv3Metric
	// HasRouterLsaSent checks if RouterLsaSent has been set in Ospfv3Metric
	HasRouterLsaSent() bool
	// RouterLsaReceived returns uint64, set in Ospfv3Metric.
	RouterLsaReceived() uint64
	// SetRouterLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetRouterLsaReceived(value uint64) Ospfv3Metric
	// HasRouterLsaReceived checks if RouterLsaReceived has been set in Ospfv3Metric
	HasRouterLsaReceived() bool
	// NetworkLsaSent returns uint64, set in Ospfv3Metric.
	NetworkLsaSent() uint64
	// SetNetworkLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetNetworkLsaSent(value uint64) Ospfv3Metric
	// HasNetworkLsaSent checks if NetworkLsaSent has been set in Ospfv3Metric
	HasNetworkLsaSent() bool
	// NetworkLsaReceived returns uint64, set in Ospfv3Metric.
	NetworkLsaReceived() uint64
	// SetNetworkLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetNetworkLsaReceived(value uint64) Ospfv3Metric
	// HasNetworkLsaReceived checks if NetworkLsaReceived has been set in Ospfv3Metric
	HasNetworkLsaReceived() bool
	// InterAreaPrefixLsaSent returns uint64, set in Ospfv3Metric.
	InterAreaPrefixLsaSent() uint64
	// SetInterAreaPrefixLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetInterAreaPrefixLsaSent(value uint64) Ospfv3Metric
	// HasInterAreaPrefixLsaSent checks if InterAreaPrefixLsaSent has been set in Ospfv3Metric
	HasInterAreaPrefixLsaSent() bool
	// InterAreaPrefixLsaReceived returns uint64, set in Ospfv3Metric.
	InterAreaPrefixLsaReceived() uint64
	// SetInterAreaPrefixLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetInterAreaPrefixLsaReceived(value uint64) Ospfv3Metric
	// HasInterAreaPrefixLsaReceived checks if InterAreaPrefixLsaReceived has been set in Ospfv3Metric
	HasInterAreaPrefixLsaReceived() bool
	// InterAreaRouterLsaSent returns uint64, set in Ospfv3Metric.
	InterAreaRouterLsaSent() uint64
	// SetInterAreaRouterLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetInterAreaRouterLsaSent(value uint64) Ospfv3Metric
	// HasInterAreaRouterLsaSent checks if InterAreaRouterLsaSent has been set in Ospfv3Metric
	HasInterAreaRouterLsaSent() bool
	// InterAreaRouterLsaReceived returns uint64, set in Ospfv3Metric.
	InterAreaRouterLsaReceived() uint64
	// SetInterAreaRouterLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetInterAreaRouterLsaReceived(value uint64) Ospfv3Metric
	// HasInterAreaRouterLsaReceived checks if InterAreaRouterLsaReceived has been set in Ospfv3Metric
	HasInterAreaRouterLsaReceived() bool
	// ExternalLsaSent returns uint64, set in Ospfv3Metric.
	ExternalLsaSent() uint64
	// SetExternalLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetExternalLsaSent(value uint64) Ospfv3Metric
	// HasExternalLsaSent checks if ExternalLsaSent has been set in Ospfv3Metric
	HasExternalLsaSent() bool
	// ExternalLsaReceived returns uint64, set in Ospfv3Metric.
	ExternalLsaReceived() uint64
	// SetExternalLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetExternalLsaReceived(value uint64) Ospfv3Metric
	// HasExternalLsaReceived checks if ExternalLsaReceived has been set in Ospfv3Metric
	HasExternalLsaReceived() bool
	// NssaLsaSent returns uint64, set in Ospfv3Metric.
	NssaLsaSent() uint64
	// SetNssaLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetNssaLsaSent(value uint64) Ospfv3Metric
	// HasNssaLsaSent checks if NssaLsaSent has been set in Ospfv3Metric
	HasNssaLsaSent() bool
	// NssaLsaReceived returns uint64, set in Ospfv3Metric.
	NssaLsaReceived() uint64
	// SetNssaLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetNssaLsaReceived(value uint64) Ospfv3Metric
	// HasNssaLsaReceived checks if NssaLsaReceived has been set in Ospfv3Metric
	HasNssaLsaReceived() bool
	// LinkLsaSent returns uint64, set in Ospfv3Metric.
	LinkLsaSent() uint64
	// SetLinkLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetLinkLsaSent(value uint64) Ospfv3Metric
	// HasLinkLsaSent checks if LinkLsaSent has been set in Ospfv3Metric
	HasLinkLsaSent() bool
	// LinkLsaReceived returns uint64, set in Ospfv3Metric.
	LinkLsaReceived() uint64
	// SetLinkLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetLinkLsaReceived(value uint64) Ospfv3Metric
	// HasLinkLsaReceived checks if LinkLsaReceived has been set in Ospfv3Metric
	HasLinkLsaReceived() bool
	// IntraAreaPrefixLsaSent returns uint64, set in Ospfv3Metric.
	IntraAreaPrefixLsaSent() uint64
	// SetIntraAreaPrefixLsaSent assigns uint64 provided by user to Ospfv3Metric
	SetIntraAreaPrefixLsaSent(value uint64) Ospfv3Metric
	// HasIntraAreaPrefixLsaSent checks if IntraAreaPrefixLsaSent has been set in Ospfv3Metric
	HasIntraAreaPrefixLsaSent() bool
	// IntraAreaPrefixLsaReceived returns uint64, set in Ospfv3Metric.
	IntraAreaPrefixLsaReceived() uint64
	// SetIntraAreaPrefixLsaReceived assigns uint64 provided by user to Ospfv3Metric
	SetIntraAreaPrefixLsaReceived(value uint64) Ospfv3Metric
	// HasIntraAreaPrefixLsaReceived checks if IntraAreaPrefixLsaReceived has been set in Ospfv3Metric
	HasIntraAreaPrefixLsaReceived() bool
}

// The name of a configured OSPFv3 router.
// Name returns a string
func (obj *ospfv3Metric) Name() string {

	return *obj.obj.Name

}

// The name of a configured OSPFv3 router.
// Name returns a string
func (obj *ospfv3Metric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured OSPFv3 router.
// SetName sets the string value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetName(value string) Ospfv3Metric {

	obj.obj.Name = &value
	return obj
}

// The number of OSPFv3 sessions in up state.
// FullStateCount returns a uint64
func (obj *ospfv3Metric) FullStateCount() uint64 {

	return *obj.obj.FullStateCount

}

// The number of OSPFv3 sessions in up state.
// FullStateCount returns a uint64
func (obj *ospfv3Metric) HasFullStateCount() bool {
	return obj.obj.FullStateCount != nil
}

// The number of OSPFv3 sessions in up state.
// SetFullStateCount sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetFullStateCount(value uint64) Ospfv3Metric {

	obj.obj.FullStateCount = &value
	return obj
}

// The number of OSPFv3 sessions in down state.
// DownStateCount returns a uint64
func (obj *ospfv3Metric) DownStateCount() uint64 {

	return *obj.obj.DownStateCount

}

// The number of OSPFv3 sessions in down state.
// DownStateCount returns a uint64
func (obj *ospfv3Metric) HasDownStateCount() bool {
	return obj.obj.DownStateCount != nil
}

// The number of OSPFv3 sessions in down state.
// SetDownStateCount sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetDownStateCount(value uint64) Ospfv3Metric {

	obj.obj.DownStateCount = &value
	return obj
}

// The number of change of OSPFv3 sessions from up to down state.
// SessionsFlap returns a uint64
func (obj *ospfv3Metric) SessionsFlap() uint64 {

	return *obj.obj.SessionsFlap

}

// The number of change of OSPFv3 sessions from up to down state.
// SessionsFlap returns a uint64
func (obj *ospfv3Metric) HasSessionsFlap() bool {
	return obj.obj.SessionsFlap != nil
}

// The number of change of OSPFv3 sessions from up to down state.
// SetSessionsFlap sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetSessionsFlap(value uint64) Ospfv3Metric {

	obj.obj.SessionsFlap = &value
	return obj
}

// The number of OSPFv3 Hello messages transmitted.
// HellosSent returns a uint64
func (obj *ospfv3Metric) HellosSent() uint64 {

	return *obj.obj.HellosSent

}

// The number of OSPFv3 Hello messages transmitted.
// HellosSent returns a uint64
func (obj *ospfv3Metric) HasHellosSent() bool {
	return obj.obj.HellosSent != nil
}

// The number of OSPFv3 Hello messages transmitted.
// SetHellosSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetHellosSent(value uint64) Ospfv3Metric {

	obj.obj.HellosSent = &value
	return obj
}

// The number of OSPFv3 Hello messages received.
// HellosReceived returns a uint64
func (obj *ospfv3Metric) HellosReceived() uint64 {

	return *obj.obj.HellosReceived

}

// The number of OSPFv3 Hello messages received.
// HellosReceived returns a uint64
func (obj *ospfv3Metric) HasHellosReceived() bool {
	return obj.obj.HellosReceived != nil
}

// The number of OSPFv3 Hello messages received.
// SetHellosReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetHellosReceived(value uint64) Ospfv3Metric {

	obj.obj.HellosReceived = &value
	return obj
}

// The number of OSPFv3 Database Description (DBD) messages transmitted.
// DbdSent returns a uint64
func (obj *ospfv3Metric) DbdSent() uint64 {

	return *obj.obj.DbdSent

}

// The number of OSPFv3 Database Description (DBD) messages transmitted.
// DbdSent returns a uint64
func (obj *ospfv3Metric) HasDbdSent() bool {
	return obj.obj.DbdSent != nil
}

// The number of OSPFv3 Database Description (DBD) messages transmitted.
// SetDbdSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetDbdSent(value uint64) Ospfv3Metric {

	obj.obj.DbdSent = &value
	return obj
}

// The number of OSPFv3 Database Description (DBD) messages received.
// DbdReceived returns a uint64
func (obj *ospfv3Metric) DbdReceived() uint64 {

	return *obj.obj.DbdReceived

}

// The number of OSPFv3 Database Description (DBD) messages received.
// DbdReceived returns a uint64
func (obj *ospfv3Metric) HasDbdReceived() bool {
	return obj.obj.DbdReceived != nil
}

// The number of OSPFv3 Database Description (DBD) messages received.
// SetDbdReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetDbdReceived(value uint64) Ospfv3Metric {

	obj.obj.DbdReceived = &value
	return obj
}

// The number of OSPFv3 LinkState (LS) Request messages transmitted.
// LsRequestSent returns a uint64
func (obj *ospfv3Metric) LsRequestSent() uint64 {

	return *obj.obj.LsRequestSent

}

// The number of OSPFv3 LinkState (LS) Request messages transmitted.
// LsRequestSent returns a uint64
func (obj *ospfv3Metric) HasLsRequestSent() bool {
	return obj.obj.LsRequestSent != nil
}

// The number of OSPFv3 LinkState (LS) Request messages transmitted.
// SetLsRequestSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLsRequestSent(value uint64) Ospfv3Metric {

	obj.obj.LsRequestSent = &value
	return obj
}

// The number of OSPFv3 LinkState (LS) Request messages received.
// LsRequestReceived returns a uint64
func (obj *ospfv3Metric) LsRequestReceived() uint64 {

	return *obj.obj.LsRequestReceived

}

// The number of OSPFv3 LinkState (LS) Request messages received.
// LsRequestReceived returns a uint64
func (obj *ospfv3Metric) HasLsRequestReceived() bool {
	return obj.obj.LsRequestReceived != nil
}

// The number of OSPFv3 LinkState (LS) Request messages received.
// SetLsRequestReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLsRequestReceived(value uint64) Ospfv3Metric {

	obj.obj.LsRequestReceived = &value
	return obj
}

// The number of OSPFv3 LinkState (LS) Update messages transmitted.
// LsUpdateSent returns a uint64
func (obj *ospfv3Metric) LsUpdateSent() uint64 {

	return *obj.obj.LsUpdateSent

}

// The number of OSPFv3 LinkState (LS) Update messages transmitted.
// LsUpdateSent returns a uint64
func (obj *ospfv3Metric) HasLsUpdateSent() bool {
	return obj.obj.LsUpdateSent != nil
}

// The number of OSPFv3 LinkState (LS) Update messages transmitted.
// SetLsUpdateSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLsUpdateSent(value uint64) Ospfv3Metric {

	obj.obj.LsUpdateSent = &value
	return obj
}

// The number of OSPFv3 LinkState (LS) Update messages received.
// LsUpdateReceived returns a uint64
func (obj *ospfv3Metric) LsUpdateReceived() uint64 {

	return *obj.obj.LsUpdateReceived

}

// The number of OSPFv3 LinkState (LS) Update messages received.
// LsUpdateReceived returns a uint64
func (obj *ospfv3Metric) HasLsUpdateReceived() bool {
	return obj.obj.LsUpdateReceived != nil
}

// The number of OSPFv3 LinkState (LS) Update messages received.
// SetLsUpdateReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLsUpdateReceived(value uint64) Ospfv3Metric {

	obj.obj.LsUpdateReceived = &value
	return obj
}

// The number of OSPFv3 LinkState (LS) Acknowledgement messages transmitted.
// LsAckSent returns a uint64
func (obj *ospfv3Metric) LsAckSent() uint64 {

	return *obj.obj.LsAckSent

}

// The number of OSPFv3 LinkState (LS) Acknowledgement messages transmitted.
// LsAckSent returns a uint64
func (obj *ospfv3Metric) HasLsAckSent() bool {
	return obj.obj.LsAckSent != nil
}

// The number of OSPFv3 LinkState (LS) Acknowledgement messages transmitted.
// SetLsAckSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLsAckSent(value uint64) Ospfv3Metric {

	obj.obj.LsAckSent = &value
	return obj
}

// The number of OSPFv3 LinkState (LS) Acknowledgement messages received.
// LsAckReceived returns a uint64
func (obj *ospfv3Metric) LsAckReceived() uint64 {

	return *obj.obj.LsAckReceived

}

// The number of OSPFv3 LinkState (LS) Acknowledgement messages received.
// LsAckReceived returns a uint64
func (obj *ospfv3Metric) HasLsAckReceived() bool {
	return obj.obj.LsAckReceived != nil
}

// The number of OSPFv3 LinkState (LS) Acknowledgement messages received.
// SetLsAckReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLsAckReceived(value uint64) Ospfv3Metric {

	obj.obj.LsAckReceived = &value
	return obj
}

// The total number of OSPFv3 LinkState Advertisement (LSA) messages transmitted.
// LsaSent returns a uint64
func (obj *ospfv3Metric) LsaSent() uint64 {

	return *obj.obj.LsaSent

}

// The total number of OSPFv3 LinkState Advertisement (LSA) messages transmitted.
// LsaSent returns a uint64
func (obj *ospfv3Metric) HasLsaSent() bool {
	return obj.obj.LsaSent != nil
}

// The total number of OSPFv3 LinkState Advertisement (LSA) messages transmitted.
// SetLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLsaSent(value uint64) Ospfv3Metric {

	obj.obj.LsaSent = &value
	return obj
}

// The total number of OSPFv3 LinkState Advertisement (LSA) messages received.
// LsaReceived returns a uint64
func (obj *ospfv3Metric) LsaReceived() uint64 {

	return *obj.obj.LsaReceived

}

// The total number of OSPFv3 LinkState Advertisement (LSA) messages received.
// LsaReceived returns a uint64
func (obj *ospfv3Metric) HasLsaReceived() bool {
	return obj.obj.LsaReceived != nil
}

// The total number of OSPFv3 LinkState Advertisement (LSA) messages received.
// SetLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.LsaReceived = &value
	return obj
}

// The number of OSPFv3 Router (Type 1) LSAs transmitted.
// RouterLsaSent returns a uint64
func (obj *ospfv3Metric) RouterLsaSent() uint64 {

	return *obj.obj.RouterLsaSent

}

// The number of OSPFv3 Router (Type 1) LSAs transmitted.
// RouterLsaSent returns a uint64
func (obj *ospfv3Metric) HasRouterLsaSent() bool {
	return obj.obj.RouterLsaSent != nil
}

// The number of OSPFv3 Router (Type 1) LSAs transmitted.
// SetRouterLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetRouterLsaSent(value uint64) Ospfv3Metric {

	obj.obj.RouterLsaSent = &value
	return obj
}

// The number of OSPFv3 Router (Type 1) LSAs received.
// RouterLsaReceived returns a uint64
func (obj *ospfv3Metric) RouterLsaReceived() uint64 {

	return *obj.obj.RouterLsaReceived

}

// The number of OSPFv3 Router (Type 1) LSAs received.
// RouterLsaReceived returns a uint64
func (obj *ospfv3Metric) HasRouterLsaReceived() bool {
	return obj.obj.RouterLsaReceived != nil
}

// The number of OSPFv3 Router (Type 1) LSAs received.
// SetRouterLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetRouterLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.RouterLsaReceived = &value
	return obj
}

// The number of OSPFv3 Network (Type 2) LSAs transmitted.
// NetworkLsaSent returns a uint64
func (obj *ospfv3Metric) NetworkLsaSent() uint64 {

	return *obj.obj.NetworkLsaSent

}

// The number of OSPFv3 Network (Type 2) LSAs transmitted.
// NetworkLsaSent returns a uint64
func (obj *ospfv3Metric) HasNetworkLsaSent() bool {
	return obj.obj.NetworkLsaSent != nil
}

// The number of OSPFv3 Network (Type 2) LSAs transmitted.
// SetNetworkLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetNetworkLsaSent(value uint64) Ospfv3Metric {

	obj.obj.NetworkLsaSent = &value
	return obj
}

// The number of OSPFv3 Network (Type 2) LSAs received.
// NetworkLsaReceived returns a uint64
func (obj *ospfv3Metric) NetworkLsaReceived() uint64 {

	return *obj.obj.NetworkLsaReceived

}

// The number of OSPFv3 Network (Type 2) LSAs received.
// NetworkLsaReceived returns a uint64
func (obj *ospfv3Metric) HasNetworkLsaReceived() bool {
	return obj.obj.NetworkLsaReceived != nil
}

// The number of OSPFv3 Network (Type 2) LSAs received.
// SetNetworkLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetNetworkLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.NetworkLsaReceived = &value
	return obj
}

// The number of OSPFv3 Inter-Area-Prefix (Type 3) LSAs transmitted.
// InterAreaPrefixLsaSent returns a uint64
func (obj *ospfv3Metric) InterAreaPrefixLsaSent() uint64 {

	return *obj.obj.InterAreaPrefixLsaSent

}

// The number of OSPFv3 Inter-Area-Prefix (Type 3) LSAs transmitted.
// InterAreaPrefixLsaSent returns a uint64
func (obj *ospfv3Metric) HasInterAreaPrefixLsaSent() bool {
	return obj.obj.InterAreaPrefixLsaSent != nil
}

// The number of OSPFv3 Inter-Area-Prefix (Type 3) LSAs transmitted.
// SetInterAreaPrefixLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetInterAreaPrefixLsaSent(value uint64) Ospfv3Metric {

	obj.obj.InterAreaPrefixLsaSent = &value
	return obj
}

// The number of OSPFv3 Inter-Area-Prefix (Type 3) LSAs received.
// InterAreaPrefixLsaReceived returns a uint64
func (obj *ospfv3Metric) InterAreaPrefixLsaReceived() uint64 {

	return *obj.obj.InterAreaPrefixLsaReceived

}

// The number of OSPFv3 Inter-Area-Prefix (Type 3) LSAs received.
// InterAreaPrefixLsaReceived returns a uint64
func (obj *ospfv3Metric) HasInterAreaPrefixLsaReceived() bool {
	return obj.obj.InterAreaPrefixLsaReceived != nil
}

// The number of OSPFv3 Inter-Area-Prefix (Type 3) LSAs received.
// SetInterAreaPrefixLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetInterAreaPrefixLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.InterAreaPrefixLsaReceived = &value
	return obj
}

// The number of OSPFv3 Inter-Area-Router (Type 4) LSAs transmitted.
// InterAreaRouterLsaSent returns a uint64
func (obj *ospfv3Metric) InterAreaRouterLsaSent() uint64 {

	return *obj.obj.InterAreaRouterLsaSent

}

// The number of OSPFv3 Inter-Area-Router (Type 4) LSAs transmitted.
// InterAreaRouterLsaSent returns a uint64
func (obj *ospfv3Metric) HasInterAreaRouterLsaSent() bool {
	return obj.obj.InterAreaRouterLsaSent != nil
}

// The number of OSPFv3 Inter-Area-Router (Type 4) LSAs transmitted.
// SetInterAreaRouterLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetInterAreaRouterLsaSent(value uint64) Ospfv3Metric {

	obj.obj.InterAreaRouterLsaSent = &value
	return obj
}

// The number of OSPFv3 Inter-Area-Router (Type 4) LSAs received.
// InterAreaRouterLsaReceived returns a uint64
func (obj *ospfv3Metric) InterAreaRouterLsaReceived() uint64 {

	return *obj.obj.InterAreaRouterLsaReceived

}

// The number of OSPFv3 Inter-Area-Router (Type 4) LSAs received.
// InterAreaRouterLsaReceived returns a uint64
func (obj *ospfv3Metric) HasInterAreaRouterLsaReceived() bool {
	return obj.obj.InterAreaRouterLsaReceived != nil
}

// The number of OSPFv3 Inter-Area-Router (Type 4) LSAs received.
// SetInterAreaRouterLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetInterAreaRouterLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.InterAreaRouterLsaReceived = &value
	return obj
}

// The number of OSPFv3 External (Type 5) LSAs transmitted.
// ExternalLsaSent returns a uint64
func (obj *ospfv3Metric) ExternalLsaSent() uint64 {

	return *obj.obj.ExternalLsaSent

}

// The number of OSPFv3 External (Type 5) LSAs transmitted.
// ExternalLsaSent returns a uint64
func (obj *ospfv3Metric) HasExternalLsaSent() bool {
	return obj.obj.ExternalLsaSent != nil
}

// The number of OSPFv3 External (Type 5) LSAs transmitted.
// SetExternalLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetExternalLsaSent(value uint64) Ospfv3Metric {

	obj.obj.ExternalLsaSent = &value
	return obj
}

// The number of OSPFv3 External (Type 5) LSAs received.
// ExternalLsaReceived returns a uint64
func (obj *ospfv3Metric) ExternalLsaReceived() uint64 {

	return *obj.obj.ExternalLsaReceived

}

// The number of OSPFv3 External (Type 5) LSAs received.
// ExternalLsaReceived returns a uint64
func (obj *ospfv3Metric) HasExternalLsaReceived() bool {
	return obj.obj.ExternalLsaReceived != nil
}

// The number of OSPFv3 External (Type 5) LSAs received.
// SetExternalLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetExternalLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.ExternalLsaReceived = &value
	return obj
}

// The number of OSPFv3 NSSA (Type 7) LSAs transmitted.
// NssaLsaSent returns a uint64
func (obj *ospfv3Metric) NssaLsaSent() uint64 {

	return *obj.obj.NssaLsaSent

}

// The number of OSPFv3 NSSA (Type 7) LSAs transmitted.
// NssaLsaSent returns a uint64
func (obj *ospfv3Metric) HasNssaLsaSent() bool {
	return obj.obj.NssaLsaSent != nil
}

// The number of OSPFv3 NSSA (Type 7) LSAs transmitted.
// SetNssaLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetNssaLsaSent(value uint64) Ospfv3Metric {

	obj.obj.NssaLsaSent = &value
	return obj
}

// The number of OSPFv3 NSSA (Type 7) LSAs received.
// NssaLsaReceived returns a uint64
func (obj *ospfv3Metric) NssaLsaReceived() uint64 {

	return *obj.obj.NssaLsaReceived

}

// The number of OSPFv3 NSSA (Type 7) LSAs received.
// NssaLsaReceived returns a uint64
func (obj *ospfv3Metric) HasNssaLsaReceived() bool {
	return obj.obj.NssaLsaReceived != nil
}

// The number of OSPFv3 NSSA (Type 7) LSAs received.
// SetNssaLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetNssaLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.NssaLsaReceived = &value
	return obj
}

// The number of OSPFv3 Link (Type 8) LSAs transmitted.
// LinkLsaSent returns a uint64
func (obj *ospfv3Metric) LinkLsaSent() uint64 {

	return *obj.obj.LinkLsaSent

}

// The number of OSPFv3 Link (Type 8) LSAs transmitted.
// LinkLsaSent returns a uint64
func (obj *ospfv3Metric) HasLinkLsaSent() bool {
	return obj.obj.LinkLsaSent != nil
}

// The number of OSPFv3 Link (Type 8) LSAs transmitted.
// SetLinkLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLinkLsaSent(value uint64) Ospfv3Metric {

	obj.obj.LinkLsaSent = &value
	return obj
}

// The number of OSPFv3 Link (Type 8) LSAs received.
// LinkLsaReceived returns a uint64
func (obj *ospfv3Metric) LinkLsaReceived() uint64 {

	return *obj.obj.LinkLsaReceived

}

// The number of OSPFv3 Link (Type 8) LSAs received.
// LinkLsaReceived returns a uint64
func (obj *ospfv3Metric) HasLinkLsaReceived() bool {
	return obj.obj.LinkLsaReceived != nil
}

// The number of OSPFv3 Link (Type 8) LSAs received.
// SetLinkLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetLinkLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.LinkLsaReceived = &value
	return obj
}

// The number of OSPFv3 Intra-Area-Prefix (Type 9) LSAs transmitted.
// IntraAreaPrefixLsaSent returns a uint64
func (obj *ospfv3Metric) IntraAreaPrefixLsaSent() uint64 {

	return *obj.obj.IntraAreaPrefixLsaSent

}

// The number of OSPFv3 Intra-Area-Prefix (Type 9) LSAs transmitted.
// IntraAreaPrefixLsaSent returns a uint64
func (obj *ospfv3Metric) HasIntraAreaPrefixLsaSent() bool {
	return obj.obj.IntraAreaPrefixLsaSent != nil
}

// The number of OSPFv3 Intra-Area-Prefix (Type 9) LSAs transmitted.
// SetIntraAreaPrefixLsaSent sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetIntraAreaPrefixLsaSent(value uint64) Ospfv3Metric {

	obj.obj.IntraAreaPrefixLsaSent = &value
	return obj
}

// The number of OSPFv3 Intra-Area-Prefix (Type 9) LSAs received.
// IntraAreaPrefixLsaReceived returns a uint64
func (obj *ospfv3Metric) IntraAreaPrefixLsaReceived() uint64 {

	return *obj.obj.IntraAreaPrefixLsaReceived

}

// The number of OSPFv3 Intra-Area-Prefix (Type 9) LSAs received.
// IntraAreaPrefixLsaReceived returns a uint64
func (obj *ospfv3Metric) HasIntraAreaPrefixLsaReceived() bool {
	return obj.obj.IntraAreaPrefixLsaReceived != nil
}

// The number of OSPFv3 Intra-Area-Prefix (Type 9) LSAs received.
// SetIntraAreaPrefixLsaReceived sets the uint64 value in the Ospfv3Metric object
func (obj *ospfv3Metric) SetIntraAreaPrefixLsaReceived(value uint64) Ospfv3Metric {

	obj.obj.IntraAreaPrefixLsaReceived = &value
	return obj
}

func (obj *ospfv3Metric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3Metric) setDefault() {

}
