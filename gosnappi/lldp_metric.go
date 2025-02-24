package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpMetric *****
type lldpMetric struct {
	validation
	obj          *otg.LldpMetric
	marshaller   marshalLldpMetric
	unMarshaller unMarshalLldpMetric
}

func NewLldpMetric() LldpMetric {
	obj := lldpMetric{obj: &otg.LldpMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpMetric) msg() *otg.LldpMetric {
	return obj.obj
}

func (obj *lldpMetric) setMsg(msg *otg.LldpMetric) LldpMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpMetric struct {
	obj *lldpMetric
}

type marshalLldpMetric interface {
	// ToProto marshals LldpMetric to protobuf object *otg.LldpMetric
	ToProto() (*otg.LldpMetric, error)
	// ToPbText marshals LldpMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpMetric struct {
	obj *lldpMetric
}

type unMarshalLldpMetric interface {
	// FromProto unmarshals LldpMetric from protobuf object *otg.LldpMetric
	FromProto(msg *otg.LldpMetric) (LldpMetric, error)
	// FromPbText unmarshals LldpMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpMetric from JSON text
	FromJson(value string) error
}

func (obj *lldpMetric) Marshal() marshalLldpMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpMetric) Unmarshal() unMarshalLldpMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpMetric) ToProto() (*otg.LldpMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpMetric) FromProto(msg *otg.LldpMetric) (LldpMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpMetric) ToPbText() (string, error) {
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

func (m *unMarshallldpMetric) FromPbText(value string) error {
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

func (m *marshallldpMetric) ToYaml() (string, error) {
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

func (m *unMarshallldpMetric) FromYaml(value string) error {
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

func (m *marshallldpMetric) ToJsonRaw() (string, error) {
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

func (m *marshallldpMetric) ToJson() (string, error) {
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

func (m *unMarshallldpMetric) FromJson(value string) error {
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

func (obj *lldpMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpMetric) Clone() (LldpMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpMetric()
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

// LldpMetric is lLDP per instance statistics information.
type LldpMetric interface {
	Validation
	// msg marshals LldpMetric to protobuf object *otg.LldpMetric
	// and doesn't set defaults
	msg() *otg.LldpMetric
	// setMsg unmarshals LldpMetric from protobuf object *otg.LldpMetric
	// and doesn't set defaults
	setMsg(*otg.LldpMetric) LldpMetric
	// provides marshal interface
	Marshal() marshalLldpMetric
	// provides unmarshal interface
	Unmarshal() unMarshalLldpMetric
	// validate validates LldpMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in LldpMetric.
	Name() string
	// SetName assigns string provided by user to LldpMetric
	SetName(value string) LldpMetric
	// HasName checks if Name has been set in LldpMetric
	HasName() bool
	// FramesRx returns uint64, set in LldpMetric.
	FramesRx() uint64
	// SetFramesRx assigns uint64 provided by user to LldpMetric
	SetFramesRx(value uint64) LldpMetric
	// HasFramesRx checks if FramesRx has been set in LldpMetric
	HasFramesRx() bool
	// FramesTx returns uint64, set in LldpMetric.
	FramesTx() uint64
	// SetFramesTx assigns uint64 provided by user to LldpMetric
	SetFramesTx(value uint64) LldpMetric
	// HasFramesTx checks if FramesTx has been set in LldpMetric
	HasFramesTx() bool
	// FramesErrorRx returns uint64, set in LldpMetric.
	FramesErrorRx() uint64
	// SetFramesErrorRx assigns uint64 provided by user to LldpMetric
	SetFramesErrorRx(value uint64) LldpMetric
	// HasFramesErrorRx checks if FramesErrorRx has been set in LldpMetric
	HasFramesErrorRx() bool
	// FramesDiscard returns uint64, set in LldpMetric.
	FramesDiscard() uint64
	// SetFramesDiscard assigns uint64 provided by user to LldpMetric
	SetFramesDiscard(value uint64) LldpMetric
	// HasFramesDiscard checks if FramesDiscard has been set in LldpMetric
	HasFramesDiscard() bool
	// TlvsDiscard returns uint64, set in LldpMetric.
	TlvsDiscard() uint64
	// SetTlvsDiscard assigns uint64 provided by user to LldpMetric
	SetTlvsDiscard(value uint64) LldpMetric
	// HasTlvsDiscard checks if TlvsDiscard has been set in LldpMetric
	HasTlvsDiscard() bool
	// TlvsUnknown returns uint64, set in LldpMetric.
	TlvsUnknown() uint64
	// SetTlvsUnknown assigns uint64 provided by user to LldpMetric
	SetTlvsUnknown(value uint64) LldpMetric
	// HasTlvsUnknown checks if TlvsUnknown has been set in LldpMetric
	HasTlvsUnknown() bool
}

// The name of the configured LLDP instance.
// Name returns a string
func (obj *lldpMetric) Name() string {

	return *obj.obj.Name

}

// The name of the configured LLDP instance.
// Name returns a string
func (obj *lldpMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of the configured LLDP instance.
// SetName sets the string value in the LldpMetric object
func (obj *lldpMetric) SetName(value string) LldpMetric {

	obj.obj.Name = &value
	return obj
}

// Number of LLDP frames received.
// FramesRx returns a uint64
func (obj *lldpMetric) FramesRx() uint64 {

	return *obj.obj.FramesRx

}

// Number of LLDP frames received.
// FramesRx returns a uint64
func (obj *lldpMetric) HasFramesRx() bool {
	return obj.obj.FramesRx != nil
}

// Number of LLDP frames received.
// SetFramesRx sets the uint64 value in the LldpMetric object
func (obj *lldpMetric) SetFramesRx(value uint64) LldpMetric {

	obj.obj.FramesRx = &value
	return obj
}

// Number of LLDP frames transmitted.
// FramesTx returns a uint64
func (obj *lldpMetric) FramesTx() uint64 {

	return *obj.obj.FramesTx

}

// Number of LLDP frames transmitted.
// FramesTx returns a uint64
func (obj *lldpMetric) HasFramesTx() bool {
	return obj.obj.FramesTx != nil
}

// Number of LLDP frames transmitted.
// SetFramesTx sets the uint64 value in the LldpMetric object
func (obj *lldpMetric) SetFramesTx(value uint64) LldpMetric {

	obj.obj.FramesTx = &value
	return obj
}

// Number of LLDP frames received with packet errors. This stat should be incremented based on statsFramesInErrorsTotal increment rule in section 10.3.2 of IEEE Std 802.1 AB-2005.
// FramesErrorRx returns a uint64
func (obj *lldpMetric) FramesErrorRx() uint64 {

	return *obj.obj.FramesErrorRx

}

// Number of LLDP frames received with packet errors. This stat should be incremented based on statsFramesInErrorsTotal increment rule in section 10.3.2 of IEEE Std 802.1 AB-2005.
// FramesErrorRx returns a uint64
func (obj *lldpMetric) HasFramesErrorRx() bool {
	return obj.obj.FramesErrorRx != nil
}

// Number of LLDP frames received with packet errors. This stat should be incremented based on statsFramesInErrorsTotal increment rule in section 10.3.2 of IEEE Std 802.1 AB-2005.
// SetFramesErrorRx sets the uint64 value in the LldpMetric object
func (obj *lldpMetric) SetFramesErrorRx(value uint64) LldpMetric {

	obj.obj.FramesErrorRx = &value
	return obj
}

// Number of LLDP frames received that are discarded. This stat should be incremented when one or more of the three mandatory  TLVs at the beginning of the LLDPDU is missing, out of order or  contains an out of range information string length. This stat should follow the validation rules in section 10.3.2 of IEEE  Std 802.1 AB-2005.
// FramesDiscard returns a uint64
func (obj *lldpMetric) FramesDiscard() uint64 {

	return *obj.obj.FramesDiscard

}

// Number of LLDP frames received that are discarded. This stat should be incremented when one or more of the three mandatory  TLVs at the beginning of the LLDPDU is missing, out of order or  contains an out of range information string length. This stat should follow the validation rules in section 10.3.2 of IEEE  Std 802.1 AB-2005.
// FramesDiscard returns a uint64
func (obj *lldpMetric) HasFramesDiscard() bool {
	return obj.obj.FramesDiscard != nil
}

// Number of LLDP frames received that are discarded. This stat should be incremented when one or more of the three mandatory  TLVs at the beginning of the LLDPDU is missing, out of order or  contains an out of range information string length. This stat should follow the validation rules in section 10.3.2 of IEEE  Std 802.1 AB-2005.
// SetFramesDiscard sets the uint64 value in the LldpMetric object
func (obj *lldpMetric) SetFramesDiscard(value uint64) LldpMetric {

	obj.obj.FramesDiscard = &value
	return obj
}

// Number of LLDP tlvs received that are discarded. If any TLV contains  an error condition specific for that particular TLV or if any TLV  extends past the physical end of the frame then these TLVs will be discarded.
// TlvsDiscard returns a uint64
func (obj *lldpMetric) TlvsDiscard() uint64 {

	return *obj.obj.TlvsDiscard

}

// Number of LLDP tlvs received that are discarded. If any TLV contains  an error condition specific for that particular TLV or if any TLV  extends past the physical end of the frame then these TLVs will be discarded.
// TlvsDiscard returns a uint64
func (obj *lldpMetric) HasTlvsDiscard() bool {
	return obj.obj.TlvsDiscard != nil
}

// Number of LLDP tlvs received that are discarded. If any TLV contains  an error condition specific for that particular TLV or if any TLV  extends past the physical end of the frame then these TLVs will be discarded.
// SetTlvsDiscard sets the uint64 value in the LldpMetric object
func (obj *lldpMetric) SetTlvsDiscard(value uint64) LldpMetric {

	obj.obj.TlvsDiscard = &value
	return obj
}

// Number of LLDP unknown tlvs received. If the OUI of the organizationlly specific TLV and/or organizationally defined subtype are not recognized,or if TLV type  value is in the range of reserved TLV types then these TLVs will be considered as  unknown TLVs.
// TlvsUnknown returns a uint64
func (obj *lldpMetric) TlvsUnknown() uint64 {

	return *obj.obj.TlvsUnknown

}

// Number of LLDP unknown tlvs received. If the OUI of the organizationlly specific TLV and/or organizationally defined subtype are not recognized,or if TLV type  value is in the range of reserved TLV types then these TLVs will be considered as  unknown TLVs.
// TlvsUnknown returns a uint64
func (obj *lldpMetric) HasTlvsUnknown() bool {
	return obj.obj.TlvsUnknown != nil
}

// Number of LLDP unknown tlvs received. If the OUI of the organizationlly specific TLV and/or organizationally defined subtype are not recognized,or if TLV type  value is in the range of reserved TLV types then these TLVs will be considered as  unknown TLVs.
// SetTlvsUnknown sets the uint64 value in the LldpMetric object
func (obj *lldpMetric) SetTlvsUnknown(value uint64) LldpMetric {

	obj.obj.TlvsUnknown = &value
	return obj
}

func (obj *lldpMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lldpMetric) setDefault() {

}
