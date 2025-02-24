package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpMetric *****
type rsvpMetric struct {
	validation
	obj          *otg.RsvpMetric
	marshaller   marshalRsvpMetric
	unMarshaller unMarshalRsvpMetric
}

func NewRsvpMetric() RsvpMetric {
	obj := rsvpMetric{obj: &otg.RsvpMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpMetric) msg() *otg.RsvpMetric {
	return obj.obj
}

func (obj *rsvpMetric) setMsg(msg *otg.RsvpMetric) RsvpMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpMetric struct {
	obj *rsvpMetric
}

type marshalRsvpMetric interface {
	// ToProto marshals RsvpMetric to protobuf object *otg.RsvpMetric
	ToProto() (*otg.RsvpMetric, error)
	// ToPbText marshals RsvpMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals RsvpMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrsvpMetric struct {
	obj *rsvpMetric
}

type unMarshalRsvpMetric interface {
	// FromProto unmarshals RsvpMetric from protobuf object *otg.RsvpMetric
	FromProto(msg *otg.RsvpMetric) (RsvpMetric, error)
	// FromPbText unmarshals RsvpMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpMetric from JSON text
	FromJson(value string) error
}

func (obj *rsvpMetric) Marshal() marshalRsvpMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpMetric) Unmarshal() unMarshalRsvpMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpMetric) ToProto() (*otg.RsvpMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpMetric) FromProto(msg *otg.RsvpMetric) (RsvpMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpMetric) ToPbText() (string, error) {
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

func (m *unMarshalrsvpMetric) FromPbText(value string) error {
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

func (m *marshalrsvpMetric) ToYaml() (string, error) {
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

func (m *unMarshalrsvpMetric) FromYaml(value string) error {
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

func (m *marshalrsvpMetric) ToJsonRaw() (string, error) {
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

func (m *marshalrsvpMetric) ToJson() (string, error) {
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

func (m *unMarshalrsvpMetric) FromJson(value string) error {
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

func (obj *rsvpMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpMetric) Clone() (RsvpMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpMetric()
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

// RsvpMetric is rSVP-TE per router statistics information.
type RsvpMetric interface {
	Validation
	// msg marshals RsvpMetric to protobuf object *otg.RsvpMetric
	// and doesn't set defaults
	msg() *otg.RsvpMetric
	// setMsg unmarshals RsvpMetric from protobuf object *otg.RsvpMetric
	// and doesn't set defaults
	setMsg(*otg.RsvpMetric) RsvpMetric
	// provides marshal interface
	Marshal() marshalRsvpMetric
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpMetric
	// validate validates RsvpMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in RsvpMetric.
	Name() string
	// SetName assigns string provided by user to RsvpMetric
	SetName(value string) RsvpMetric
	// HasName checks if Name has been set in RsvpMetric
	HasName() bool
	// IngressP2PLspsConfigured returns uint32, set in RsvpMetric.
	IngressP2PLspsConfigured() uint32
	// SetIngressP2PLspsConfigured assigns uint32 provided by user to RsvpMetric
	SetIngressP2PLspsConfigured(value uint32) RsvpMetric
	// HasIngressP2PLspsConfigured checks if IngressP2PLspsConfigured has been set in RsvpMetric
	HasIngressP2PLspsConfigured() bool
	// IngressP2PLspsUp returns uint32, set in RsvpMetric.
	IngressP2PLspsUp() uint32
	// SetIngressP2PLspsUp assigns uint32 provided by user to RsvpMetric
	SetIngressP2PLspsUp(value uint32) RsvpMetric
	// HasIngressP2PLspsUp checks if IngressP2PLspsUp has been set in RsvpMetric
	HasIngressP2PLspsUp() bool
	// EgressP2PLspsUp returns uint32, set in RsvpMetric.
	EgressP2PLspsUp() uint32
	// SetEgressP2PLspsUp assigns uint32 provided by user to RsvpMetric
	SetEgressP2PLspsUp(value uint32) RsvpMetric
	// HasEgressP2PLspsUp checks if EgressP2PLspsUp has been set in RsvpMetric
	HasEgressP2PLspsUp() bool
	// LspFlapCount returns uint64, set in RsvpMetric.
	LspFlapCount() uint64
	// SetLspFlapCount assigns uint64 provided by user to RsvpMetric
	SetLspFlapCount(value uint64) RsvpMetric
	// HasLspFlapCount checks if LspFlapCount has been set in RsvpMetric
	HasLspFlapCount() bool
	// PathsTx returns uint64, set in RsvpMetric.
	PathsTx() uint64
	// SetPathsTx assigns uint64 provided by user to RsvpMetric
	SetPathsTx(value uint64) RsvpMetric
	// HasPathsTx checks if PathsTx has been set in RsvpMetric
	HasPathsTx() bool
	// PathsRx returns uint64, set in RsvpMetric.
	PathsRx() uint64
	// SetPathsRx assigns uint64 provided by user to RsvpMetric
	SetPathsRx(value uint64) RsvpMetric
	// HasPathsRx checks if PathsRx has been set in RsvpMetric
	HasPathsRx() bool
	// ResvsTx returns uint64, set in RsvpMetric.
	ResvsTx() uint64
	// SetResvsTx assigns uint64 provided by user to RsvpMetric
	SetResvsTx(value uint64) RsvpMetric
	// HasResvsTx checks if ResvsTx has been set in RsvpMetric
	HasResvsTx() bool
	// ResvsRx returns uint64, set in RsvpMetric.
	ResvsRx() uint64
	// SetResvsRx assigns uint64 provided by user to RsvpMetric
	SetResvsRx(value uint64) RsvpMetric
	// HasResvsRx checks if ResvsRx has been set in RsvpMetric
	HasResvsRx() bool
	// PathTearsTx returns uint64, set in RsvpMetric.
	PathTearsTx() uint64
	// SetPathTearsTx assigns uint64 provided by user to RsvpMetric
	SetPathTearsTx(value uint64) RsvpMetric
	// HasPathTearsTx checks if PathTearsTx has been set in RsvpMetric
	HasPathTearsTx() bool
	// PathTearsRx returns uint64, set in RsvpMetric.
	PathTearsRx() uint64
	// SetPathTearsRx assigns uint64 provided by user to RsvpMetric
	SetPathTearsRx(value uint64) RsvpMetric
	// HasPathTearsRx checks if PathTearsRx has been set in RsvpMetric
	HasPathTearsRx() bool
	// ResvTearsTx returns uint64, set in RsvpMetric.
	ResvTearsTx() uint64
	// SetResvTearsTx assigns uint64 provided by user to RsvpMetric
	SetResvTearsTx(value uint64) RsvpMetric
	// HasResvTearsTx checks if ResvTearsTx has been set in RsvpMetric
	HasResvTearsTx() bool
	// ResvTearsRx returns uint64, set in RsvpMetric.
	ResvTearsRx() uint64
	// SetResvTearsRx assigns uint64 provided by user to RsvpMetric
	SetResvTearsRx(value uint64) RsvpMetric
	// HasResvTearsRx checks if ResvTearsRx has been set in RsvpMetric
	HasResvTearsRx() bool
	// PathErrorsTx returns uint64, set in RsvpMetric.
	PathErrorsTx() uint64
	// SetPathErrorsTx assigns uint64 provided by user to RsvpMetric
	SetPathErrorsTx(value uint64) RsvpMetric
	// HasPathErrorsTx checks if PathErrorsTx has been set in RsvpMetric
	HasPathErrorsTx() bool
	// PathErrorsRx returns uint64, set in RsvpMetric.
	PathErrorsRx() uint64
	// SetPathErrorsRx assigns uint64 provided by user to RsvpMetric
	SetPathErrorsRx(value uint64) RsvpMetric
	// HasPathErrorsRx checks if PathErrorsRx has been set in RsvpMetric
	HasPathErrorsRx() bool
	// ResvErrorsTx returns uint64, set in RsvpMetric.
	ResvErrorsTx() uint64
	// SetResvErrorsTx assigns uint64 provided by user to RsvpMetric
	SetResvErrorsTx(value uint64) RsvpMetric
	// HasResvErrorsTx checks if ResvErrorsTx has been set in RsvpMetric
	HasResvErrorsTx() bool
	// ResvErrorsRx returns uint64, set in RsvpMetric.
	ResvErrorsRx() uint64
	// SetResvErrorsRx assigns uint64 provided by user to RsvpMetric
	SetResvErrorsRx(value uint64) RsvpMetric
	// HasResvErrorsRx checks if ResvErrorsRx has been set in RsvpMetric
	HasResvErrorsRx() bool
	// ResvConfTx returns uint64, set in RsvpMetric.
	ResvConfTx() uint64
	// SetResvConfTx assigns uint64 provided by user to RsvpMetric
	SetResvConfTx(value uint64) RsvpMetric
	// HasResvConfTx checks if ResvConfTx has been set in RsvpMetric
	HasResvConfTx() bool
	// ResvConfRx returns uint64, set in RsvpMetric.
	ResvConfRx() uint64
	// SetResvConfRx assigns uint64 provided by user to RsvpMetric
	SetResvConfRx(value uint64) RsvpMetric
	// HasResvConfRx checks if ResvConfRx has been set in RsvpMetric
	HasResvConfRx() bool
	// HellosTx returns uint64, set in RsvpMetric.
	HellosTx() uint64
	// SetHellosTx assigns uint64 provided by user to RsvpMetric
	SetHellosTx(value uint64) RsvpMetric
	// HasHellosTx checks if HellosTx has been set in RsvpMetric
	HasHellosTx() bool
	// HellosRx returns uint64, set in RsvpMetric.
	HellosRx() uint64
	// SetHellosRx assigns uint64 provided by user to RsvpMetric
	SetHellosRx(value uint64) RsvpMetric
	// HasHellosRx checks if HellosRx has been set in RsvpMetric
	HasHellosRx() bool
	// AcksTx returns uint64, set in RsvpMetric.
	AcksTx() uint64
	// SetAcksTx assigns uint64 provided by user to RsvpMetric
	SetAcksTx(value uint64) RsvpMetric
	// HasAcksTx checks if AcksTx has been set in RsvpMetric
	HasAcksTx() bool
	// AcksRx returns uint64, set in RsvpMetric.
	AcksRx() uint64
	// SetAcksRx assigns uint64 provided by user to RsvpMetric
	SetAcksRx(value uint64) RsvpMetric
	// HasAcksRx checks if AcksRx has been set in RsvpMetric
	HasAcksRx() bool
	// NacksTx returns uint64, set in RsvpMetric.
	NacksTx() uint64
	// SetNacksTx assigns uint64 provided by user to RsvpMetric
	SetNacksTx(value uint64) RsvpMetric
	// HasNacksTx checks if NacksTx has been set in RsvpMetric
	HasNacksTx() bool
	// NacksRx returns uint64, set in RsvpMetric.
	NacksRx() uint64
	// SetNacksRx assigns uint64 provided by user to RsvpMetric
	SetNacksRx(value uint64) RsvpMetric
	// HasNacksRx checks if NacksRx has been set in RsvpMetric
	HasNacksRx() bool
	// SrefreshTx returns uint64, set in RsvpMetric.
	SrefreshTx() uint64
	// SetSrefreshTx assigns uint64 provided by user to RsvpMetric
	SetSrefreshTx(value uint64) RsvpMetric
	// HasSrefreshTx checks if SrefreshTx has been set in RsvpMetric
	HasSrefreshTx() bool
	// SrefreshRx returns uint64, set in RsvpMetric.
	SrefreshRx() uint64
	// SetSrefreshRx assigns uint64 provided by user to RsvpMetric
	SetSrefreshRx(value uint64) RsvpMetric
	// HasSrefreshRx checks if SrefreshRx has been set in RsvpMetric
	HasSrefreshRx() bool
	// BundleTx returns uint64, set in RsvpMetric.
	BundleTx() uint64
	// SetBundleTx assigns uint64 provided by user to RsvpMetric
	SetBundleTx(value uint64) RsvpMetric
	// HasBundleTx checks if BundleTx has been set in RsvpMetric
	HasBundleTx() bool
	// BundleRx returns uint64, set in RsvpMetric.
	BundleRx() uint64
	// SetBundleRx assigns uint64 provided by user to RsvpMetric
	SetBundleRx(value uint64) RsvpMetric
	// HasBundleRx checks if BundleRx has been set in RsvpMetric
	HasBundleRx() bool
	// PathReevaluationRequestTx returns uint64, set in RsvpMetric.
	PathReevaluationRequestTx() uint64
	// SetPathReevaluationRequestTx assigns uint64 provided by user to RsvpMetric
	SetPathReevaluationRequestTx(value uint64) RsvpMetric
	// HasPathReevaluationRequestTx checks if PathReevaluationRequestTx has been set in RsvpMetric
	HasPathReevaluationRequestTx() bool
	// PathReoptimizations returns uint64, set in RsvpMetric.
	PathReoptimizations() uint64
	// SetPathReoptimizations assigns uint64 provided by user to RsvpMetric
	SetPathReoptimizations(value uint64) RsvpMetric
	// HasPathReoptimizations checks if PathReoptimizations has been set in RsvpMetric
	HasPathReoptimizations() bool
}

// The name of a configured RSVP router.
// Name returns a string
func (obj *rsvpMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured RSVP router.
// Name returns a string
func (obj *rsvpMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured RSVP router.
// SetName sets the string value in the RsvpMetric object
func (obj *rsvpMetric) SetName(value string) RsvpMetric {

	obj.obj.Name = &value
	return obj
}

// The number of ingress point-to-point LSPs configured or transiting through the RSVP router which have been initated from the test port.
// IngressP2PLspsConfigured returns a uint32
func (obj *rsvpMetric) IngressP2PLspsConfigured() uint32 {

	return *obj.obj.IngressP2PLspsConfigured

}

// The number of ingress point-to-point LSPs configured or transiting through the RSVP router which have been initated from the test port.
// IngressP2PLspsConfigured returns a uint32
func (obj *rsvpMetric) HasIngressP2PLspsConfigured() bool {
	return obj.obj.IngressP2PLspsConfigured != nil
}

// The number of ingress point-to-point LSPs configured or transiting through the RSVP router which have been initated from the test port.
// SetIngressP2PLspsConfigured sets the uint32 value in the RsvpMetric object
func (obj *rsvpMetric) SetIngressP2PLspsConfigured(value uint32) RsvpMetric {

	obj.obj.IngressP2PLspsConfigured = &value
	return obj
}

// The number of ingress point-to-point LSPs for which Resv has been received and is currently up.
// IngressP2PLspsUp returns a uint32
func (obj *rsvpMetric) IngressP2PLspsUp() uint32 {

	return *obj.obj.IngressP2PLspsUp

}

// The number of ingress point-to-point LSPs for which Resv has been received and is currently up.
// IngressP2PLspsUp returns a uint32
func (obj *rsvpMetric) HasIngressP2PLspsUp() bool {
	return obj.obj.IngressP2PLspsUp != nil
}

// The number of ingress point-to-point LSPs for which Resv has been received and is currently up.
// SetIngressP2PLspsUp sets the uint32 value in the RsvpMetric object
func (obj *rsvpMetric) SetIngressP2PLspsUp(value uint32) RsvpMetric {

	obj.obj.IngressP2PLspsUp = &value
	return obj
}

// The number of egress point-to-point LSPs for which Path requests were successfully processed and is currently up.
// EgressP2PLspsUp returns a uint32
func (obj *rsvpMetric) EgressP2PLspsUp() uint32 {

	return *obj.obj.EgressP2PLspsUp

}

// The number of egress point-to-point LSPs for which Path requests were successfully processed and is currently up.
// EgressP2PLspsUp returns a uint32
func (obj *rsvpMetric) HasEgressP2PLspsUp() bool {
	return obj.obj.EgressP2PLspsUp != nil
}

// The number of egress point-to-point LSPs for which Path requests were successfully processed and is currently up.
// SetEgressP2PLspsUp sets the uint32 value in the RsvpMetric object
func (obj *rsvpMetric) SetEgressP2PLspsUp(value uint32) RsvpMetric {

	obj.obj.EgressP2PLspsUp = &value
	return obj
}

// The number of times an LSP went from up to down state either because it timed out while waiting for Refreshes or  a PathTear or ResvTear message was received which caused the LSP to flap.
// LspFlapCount returns a uint64
func (obj *rsvpMetric) LspFlapCount() uint64 {

	return *obj.obj.LspFlapCount

}

// The number of times an LSP went from up to down state either because it timed out while waiting for Refreshes or  a PathTear or ResvTear message was received which caused the LSP to flap.
// LspFlapCount returns a uint64
func (obj *rsvpMetric) HasLspFlapCount() bool {
	return obj.obj.LspFlapCount != nil
}

// The number of times an LSP went from up to down state either because it timed out while waiting for Refreshes or  a PathTear or ResvTear message was received which caused the LSP to flap.
// SetLspFlapCount sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetLspFlapCount(value uint64) RsvpMetric {

	obj.obj.LspFlapCount = &value
	return obj
}

// The number of Path messages sent by this RSVP router.
// PathsTx returns a uint64
func (obj *rsvpMetric) PathsTx() uint64 {

	return *obj.obj.PathsTx

}

// The number of Path messages sent by this RSVP router.
// PathsTx returns a uint64
func (obj *rsvpMetric) HasPathsTx() bool {
	return obj.obj.PathsTx != nil
}

// The number of Path messages sent by this RSVP router.
// SetPathsTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetPathsTx(value uint64) RsvpMetric {

	obj.obj.PathsTx = &value
	return obj
}

// The number of Path messages received by this RSVP router.
// PathsRx returns a uint64
func (obj *rsvpMetric) PathsRx() uint64 {

	return *obj.obj.PathsRx

}

// The number of Path messages received by this RSVP router.
// PathsRx returns a uint64
func (obj *rsvpMetric) HasPathsRx() bool {
	return obj.obj.PathsRx != nil
}

// The number of Path messages received by this RSVP router.
// SetPathsRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetPathsRx(value uint64) RsvpMetric {

	obj.obj.PathsRx = &value
	return obj
}

// The number of Resv messages sent by this RSVP router.
// ResvsTx returns a uint64
func (obj *rsvpMetric) ResvsTx() uint64 {

	return *obj.obj.ResvsTx

}

// The number of Resv messages sent by this RSVP router.
// ResvsTx returns a uint64
func (obj *rsvpMetric) HasResvsTx() bool {
	return obj.obj.ResvsTx != nil
}

// The number of Resv messages sent by this RSVP router.
// SetResvsTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetResvsTx(value uint64) RsvpMetric {

	obj.obj.ResvsTx = &value
	return obj
}

// The number of Resv messages received by this RSVP router.
// ResvsRx returns a uint64
func (obj *rsvpMetric) ResvsRx() uint64 {

	return *obj.obj.ResvsRx

}

// The number of Resv messages received by this RSVP router.
// ResvsRx returns a uint64
func (obj *rsvpMetric) HasResvsRx() bool {
	return obj.obj.ResvsRx != nil
}

// The number of Resv messages received by this RSVP router.
// SetResvsRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetResvsRx(value uint64) RsvpMetric {

	obj.obj.ResvsRx = &value
	return obj
}

// The number of  Path Tear messages sent by this RSVP router.
// PathTearsTx returns a uint64
func (obj *rsvpMetric) PathTearsTx() uint64 {

	return *obj.obj.PathTearsTx

}

// The number of  Path Tear messages sent by this RSVP router.
// PathTearsTx returns a uint64
func (obj *rsvpMetric) HasPathTearsTx() bool {
	return obj.obj.PathTearsTx != nil
}

// The number of  Path Tear messages sent by this RSVP router.
// SetPathTearsTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetPathTearsTx(value uint64) RsvpMetric {

	obj.obj.PathTearsTx = &value
	return obj
}

// The number of Path Tear messages received by this RSVP router.
// PathTearsRx returns a uint64
func (obj *rsvpMetric) PathTearsRx() uint64 {

	return *obj.obj.PathTearsRx

}

// The number of Path Tear messages received by this RSVP router.
// PathTearsRx returns a uint64
func (obj *rsvpMetric) HasPathTearsRx() bool {
	return obj.obj.PathTearsRx != nil
}

// The number of Path Tear messages received by this RSVP router.
// SetPathTearsRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetPathTearsRx(value uint64) RsvpMetric {

	obj.obj.PathTearsRx = &value
	return obj
}

// The number of  Resv Tear messages sent by this RSVP router.
// ResvTearsTx returns a uint64
func (obj *rsvpMetric) ResvTearsTx() uint64 {

	return *obj.obj.ResvTearsTx

}

// The number of  Resv Tear messages sent by this RSVP router.
// ResvTearsTx returns a uint64
func (obj *rsvpMetric) HasResvTearsTx() bool {
	return obj.obj.ResvTearsTx != nil
}

// The number of  Resv Tear messages sent by this RSVP router.
// SetResvTearsTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetResvTearsTx(value uint64) RsvpMetric {

	obj.obj.ResvTearsTx = &value
	return obj
}

// The number of Resv Tear messages received by this RSVP router.
// ResvTearsRx returns a uint64
func (obj *rsvpMetric) ResvTearsRx() uint64 {

	return *obj.obj.ResvTearsRx

}

// The number of Resv Tear messages received by this RSVP router.
// ResvTearsRx returns a uint64
func (obj *rsvpMetric) HasResvTearsRx() bool {
	return obj.obj.ResvTearsRx != nil
}

// The number of Resv Tear messages received by this RSVP router.
// SetResvTearsRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetResvTearsRx(value uint64) RsvpMetric {

	obj.obj.ResvTearsRx = &value
	return obj
}

// The number of Path Error messages sent by this RSVP router.
// PathErrorsTx returns a uint64
func (obj *rsvpMetric) PathErrorsTx() uint64 {

	return *obj.obj.PathErrorsTx

}

// The number of Path Error messages sent by this RSVP router.
// PathErrorsTx returns a uint64
func (obj *rsvpMetric) HasPathErrorsTx() bool {
	return obj.obj.PathErrorsTx != nil
}

// The number of Path Error messages sent by this RSVP router.
// SetPathErrorsTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetPathErrorsTx(value uint64) RsvpMetric {

	obj.obj.PathErrorsTx = &value
	return obj
}

// The number of Path Error messages received by this RSVP router.
// PathErrorsRx returns a uint64
func (obj *rsvpMetric) PathErrorsRx() uint64 {

	return *obj.obj.PathErrorsRx

}

// The number of Path Error messages received by this RSVP router.
// PathErrorsRx returns a uint64
func (obj *rsvpMetric) HasPathErrorsRx() bool {
	return obj.obj.PathErrorsRx != nil
}

// The number of Path Error messages received by this RSVP router.
// SetPathErrorsRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetPathErrorsRx(value uint64) RsvpMetric {

	obj.obj.PathErrorsRx = &value
	return obj
}

// The number of Resv Error messages sent by this RSVP router.
// ResvErrorsTx returns a uint64
func (obj *rsvpMetric) ResvErrorsTx() uint64 {

	return *obj.obj.ResvErrorsTx

}

// The number of Resv Error messages sent by this RSVP router.
// ResvErrorsTx returns a uint64
func (obj *rsvpMetric) HasResvErrorsTx() bool {
	return obj.obj.ResvErrorsTx != nil
}

// The number of Resv Error messages sent by this RSVP router.
// SetResvErrorsTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetResvErrorsTx(value uint64) RsvpMetric {

	obj.obj.ResvErrorsTx = &value
	return obj
}

// The number of Resv Error messages received by this RSVP router.
// ResvErrorsRx returns a uint64
func (obj *rsvpMetric) ResvErrorsRx() uint64 {

	return *obj.obj.ResvErrorsRx

}

// The number of Resv Error messages received by this RSVP router.
// ResvErrorsRx returns a uint64
func (obj *rsvpMetric) HasResvErrorsRx() bool {
	return obj.obj.ResvErrorsRx != nil
}

// The number of Resv Error messages received by this RSVP router.
// SetResvErrorsRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetResvErrorsRx(value uint64) RsvpMetric {

	obj.obj.ResvErrorsRx = &value
	return obj
}

// The number of ResvConf messages sent by this RSVP router.
// ResvConfTx returns a uint64
func (obj *rsvpMetric) ResvConfTx() uint64 {

	return *obj.obj.ResvConfTx

}

// The number of ResvConf messages sent by this RSVP router.
// ResvConfTx returns a uint64
func (obj *rsvpMetric) HasResvConfTx() bool {
	return obj.obj.ResvConfTx != nil
}

// The number of ResvConf messages sent by this RSVP router.
// SetResvConfTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetResvConfTx(value uint64) RsvpMetric {

	obj.obj.ResvConfTx = &value
	return obj
}

// The number of ResvConf messages received by this RSVP router.
// ResvConfRx returns a uint64
func (obj *rsvpMetric) ResvConfRx() uint64 {

	return *obj.obj.ResvConfRx

}

// The number of ResvConf messages received by this RSVP router.
// ResvConfRx returns a uint64
func (obj *rsvpMetric) HasResvConfRx() bool {
	return obj.obj.ResvConfRx != nil
}

// The number of ResvConf messages received by this RSVP router.
// SetResvConfRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetResvConfRx(value uint64) RsvpMetric {

	obj.obj.ResvConfRx = &value
	return obj
}

// The number of Hello messages sent by this RSVP router.
// HellosTx returns a uint64
func (obj *rsvpMetric) HellosTx() uint64 {

	return *obj.obj.HellosTx

}

// The number of Hello messages sent by this RSVP router.
// HellosTx returns a uint64
func (obj *rsvpMetric) HasHellosTx() bool {
	return obj.obj.HellosTx != nil
}

// The number of Hello messages sent by this RSVP router.
// SetHellosTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetHellosTx(value uint64) RsvpMetric {

	obj.obj.HellosTx = &value
	return obj
}

// The number of Hello messages received by this RSVP router.
// HellosRx returns a uint64
func (obj *rsvpMetric) HellosRx() uint64 {

	return *obj.obj.HellosRx

}

// The number of Hello messages received by this RSVP router.
// HellosRx returns a uint64
func (obj *rsvpMetric) HasHellosRx() bool {
	return obj.obj.HellosRx != nil
}

// The number of Hello messages received by this RSVP router.
// SetHellosRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetHellosRx(value uint64) RsvpMetric {

	obj.obj.HellosRx = &value
	return obj
}

// The number of Ack messages sent by this RSVP router.
// AcksTx returns a uint64
func (obj *rsvpMetric) AcksTx() uint64 {

	return *obj.obj.AcksTx

}

// The number of Ack messages sent by this RSVP router.
// AcksTx returns a uint64
func (obj *rsvpMetric) HasAcksTx() bool {
	return obj.obj.AcksTx != nil
}

// The number of Ack messages sent by this RSVP router.
// SetAcksTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetAcksTx(value uint64) RsvpMetric {

	obj.obj.AcksTx = &value
	return obj
}

// The number of Ack messages received by this RSVP router.
// AcksRx returns a uint64
func (obj *rsvpMetric) AcksRx() uint64 {

	return *obj.obj.AcksRx

}

// The number of Ack messages received by this RSVP router.
// AcksRx returns a uint64
func (obj *rsvpMetric) HasAcksRx() bool {
	return obj.obj.AcksRx != nil
}

// The number of Ack messages received by this RSVP router.
// SetAcksRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetAcksRx(value uint64) RsvpMetric {

	obj.obj.AcksRx = &value
	return obj
}

// The number of Nack messages sent by this RSVP router.
// NacksTx returns a uint64
func (obj *rsvpMetric) NacksTx() uint64 {

	return *obj.obj.NacksTx

}

// The number of Nack messages sent by this RSVP router.
// NacksTx returns a uint64
func (obj *rsvpMetric) HasNacksTx() bool {
	return obj.obj.NacksTx != nil
}

// The number of Nack messages sent by this RSVP router.
// SetNacksTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetNacksTx(value uint64) RsvpMetric {

	obj.obj.NacksTx = &value
	return obj
}

// The number of Nack messages received by this RSVP router.
// NacksRx returns a uint64
func (obj *rsvpMetric) NacksRx() uint64 {

	return *obj.obj.NacksRx

}

// The number of Nack messages received by this RSVP router.
// NacksRx returns a uint64
func (obj *rsvpMetric) HasNacksRx() bool {
	return obj.obj.NacksRx != nil
}

// The number of Nack messages received by this RSVP router.
// SetNacksRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetNacksRx(value uint64) RsvpMetric {

	obj.obj.NacksRx = &value
	return obj
}

// The number of SRefresh messages sent by this RSVP router.
// SrefreshTx returns a uint64
func (obj *rsvpMetric) SrefreshTx() uint64 {

	return *obj.obj.SrefreshTx

}

// The number of SRefresh messages sent by this RSVP router.
// SrefreshTx returns a uint64
func (obj *rsvpMetric) HasSrefreshTx() bool {
	return obj.obj.SrefreshTx != nil
}

// The number of SRefresh messages sent by this RSVP router.
// SetSrefreshTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetSrefreshTx(value uint64) RsvpMetric {

	obj.obj.SrefreshTx = &value
	return obj
}

// The number of SRefresh messages received by this RSVP router.
// SrefreshRx returns a uint64
func (obj *rsvpMetric) SrefreshRx() uint64 {

	return *obj.obj.SrefreshRx

}

// The number of SRefresh messages received by this RSVP router.
// SrefreshRx returns a uint64
func (obj *rsvpMetric) HasSrefreshRx() bool {
	return obj.obj.SrefreshRx != nil
}

// The number of SRefresh messages received by this RSVP router.
// SetSrefreshRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetSrefreshRx(value uint64) RsvpMetric {

	obj.obj.SrefreshRx = &value
	return obj
}

// The number of Bundle messages sent by this RSVP router.
// BundleTx returns a uint64
func (obj *rsvpMetric) BundleTx() uint64 {

	return *obj.obj.BundleTx

}

// The number of Bundle messages sent by this RSVP router.
// BundleTx returns a uint64
func (obj *rsvpMetric) HasBundleTx() bool {
	return obj.obj.BundleTx != nil
}

// The number of Bundle messages sent by this RSVP router.
// SetBundleTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetBundleTx(value uint64) RsvpMetric {

	obj.obj.BundleTx = &value
	return obj
}

// The number of Bundle messages received by this RSVP router.
// BundleRx returns a uint64
func (obj *rsvpMetric) BundleRx() uint64 {

	return *obj.obj.BundleRx

}

// The number of Bundle messages received by this RSVP router.
// BundleRx returns a uint64
func (obj *rsvpMetric) HasBundleRx() bool {
	return obj.obj.BundleRx != nil
}

// The number of Bundle messages received by this RSVP router.
// SetBundleRx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetBundleRx(value uint64) RsvpMetric {

	obj.obj.BundleRx = &value
	return obj
}

// The number of Path messages with Path Re-evaluation Request enabled sent by this RSVP router.
// PathReevaluationRequestTx returns a uint64
func (obj *rsvpMetric) PathReevaluationRequestTx() uint64 {

	return *obj.obj.PathReevaluationRequestTx

}

// The number of Path messages with Path Re-evaluation Request enabled sent by this RSVP router.
// PathReevaluationRequestTx returns a uint64
func (obj *rsvpMetric) HasPathReevaluationRequestTx() bool {
	return obj.obj.PathReevaluationRequestTx != nil
}

// The number of Path messages with Path Re-evaluation Request enabled sent by this RSVP router.
// SetPathReevaluationRequestTx sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetPathReevaluationRequestTx(value uint64) RsvpMetric {

	obj.obj.PathReevaluationRequestTx = &value
	return obj
}

// The number of successfully completed Make-Before-Break operations on LSPs on this RSVP router.
// PathReoptimizations returns a uint64
func (obj *rsvpMetric) PathReoptimizations() uint64 {

	return *obj.obj.PathReoptimizations

}

// The number of successfully completed Make-Before-Break operations on LSPs on this RSVP router.
// PathReoptimizations returns a uint64
func (obj *rsvpMetric) HasPathReoptimizations() bool {
	return obj.obj.PathReoptimizations != nil
}

// The number of successfully completed Make-Before-Break operations on LSPs on this RSVP router.
// SetPathReoptimizations sets the uint64 value in the RsvpMetric object
func (obj *rsvpMetric) SetPathReoptimizations(value uint64) RsvpMetric {

	obj.obj.PathReoptimizations = &value
	return obj
}

func (obj *rsvpMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *rsvpMetric) setDefault() {

}
