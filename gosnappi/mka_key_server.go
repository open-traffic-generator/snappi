package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaKeyServer *****
type mkaKeyServer struct {
	validation
	obj          *otg.MkaKeyServer
	marshaller   marshalMkaKeyServer
	unMarshaller unMarshalMkaKeyServer
}

func NewMkaKeyServer() MkaKeyServer {
	obj := mkaKeyServer{obj: &otg.MkaKeyServer{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaKeyServer) msg() *otg.MkaKeyServer {
	return obj.obj
}

func (obj *mkaKeyServer) setMsg(msg *otg.MkaKeyServer) MkaKeyServer {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaKeyServer struct {
	obj *mkaKeyServer
}

type marshalMkaKeyServer interface {
	// ToProto marshals MkaKeyServer to protobuf object *otg.MkaKeyServer
	ToProto() (*otg.MkaKeyServer, error)
	// ToPbText marshals MkaKeyServer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaKeyServer to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaKeyServer to JSON text
	ToJson() (string, error)
}

type unMarshalmkaKeyServer struct {
	obj *mkaKeyServer
}

type unMarshalMkaKeyServer interface {
	// FromProto unmarshals MkaKeyServer from protobuf object *otg.MkaKeyServer
	FromProto(msg *otg.MkaKeyServer) (MkaKeyServer, error)
	// FromPbText unmarshals MkaKeyServer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaKeyServer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaKeyServer from JSON text
	FromJson(value string) error
}

func (obj *mkaKeyServer) Marshal() marshalMkaKeyServer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaKeyServer{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaKeyServer) Unmarshal() unMarshalMkaKeyServer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaKeyServer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaKeyServer) ToProto() (*otg.MkaKeyServer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaKeyServer) FromProto(msg *otg.MkaKeyServer) (MkaKeyServer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaKeyServer) ToPbText() (string, error) {
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

func (m *unMarshalmkaKeyServer) FromPbText(value string) error {
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

func (m *marshalmkaKeyServer) ToYaml() (string, error) {
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

func (m *unMarshalmkaKeyServer) FromYaml(value string) error {
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

func (m *marshalmkaKeyServer) ToJson() (string, error) {
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

func (m *unMarshalmkaKeyServer) FromJson(value string) error {
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

func (obj *mkaKeyServer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaKeyServer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaKeyServer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaKeyServer) Clone() (MkaKeyServer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaKeyServer()
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

// MkaKeyServer is key server attributes of a KaY.
type MkaKeyServer interface {
	Validation
	// msg marshals MkaKeyServer to protobuf object *otg.MkaKeyServer
	// and doesn't set defaults
	msg() *otg.MkaKeyServer
	// setMsg unmarshals MkaKeyServer from protobuf object *otg.MkaKeyServer
	// and doesn't set defaults
	setMsg(*otg.MkaKeyServer) MkaKeyServer
	// provides marshal interface
	Marshal() marshalMkaKeyServer
	// provides unmarshal interface
	Unmarshal() unMarshalMkaKeyServer
	// validate validates MkaKeyServer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaKeyServer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ConfidentialtyOffset returns MkaKeyServerConfidentialtyOffsetEnum, set in MkaKeyServer
	ConfidentialtyOffset() MkaKeyServerConfidentialtyOffsetEnum
	// SetConfidentialtyOffset assigns MkaKeyServerConfidentialtyOffsetEnum provided by user to MkaKeyServer
	SetConfidentialtyOffset(value MkaKeyServerConfidentialtyOffsetEnum) MkaKeyServer
	// HasConfidentialtyOffset checks if ConfidentialtyOffset has been set in MkaKeyServer
	HasConfidentialtyOffset() bool
	// CipherSuite returns MkaKeyServerCipherSuiteEnum, set in MkaKeyServer
	CipherSuite() MkaKeyServerCipherSuiteEnum
	// SetCipherSuite assigns MkaKeyServerCipherSuiteEnum provided by user to MkaKeyServer
	SetCipherSuite(value MkaKeyServerCipherSuiteEnum) MkaKeyServer
	// HasCipherSuite checks if CipherSuite has been set in MkaKeyServer
	HasCipherSuite() bool
	// StartingKeyNumber returns uint32, set in MkaKeyServer.
	StartingKeyNumber() uint32
	// SetStartingKeyNumber assigns uint32 provided by user to MkaKeyServer
	SetStartingKeyNumber(value uint32) MkaKeyServer
	// HasStartingKeyNumber checks if StartingKeyNumber has been set in MkaKeyServer
	HasStartingKeyNumber() bool
	// StartingDistributedAn returns uint32, set in MkaKeyServer.
	StartingDistributedAn() uint32
	// SetStartingDistributedAn assigns uint32 provided by user to MkaKeyServer
	SetStartingDistributedAn(value uint32) MkaKeyServer
	// HasStartingDistributedAn checks if StartingDistributedAn has been set in MkaKeyServer
	HasStartingDistributedAn() bool
	// RekeyThresholdPn returns string, set in MkaKeyServer.
	RekeyThresholdPn() string
	// SetRekeyThresholdPn assigns string provided by user to MkaKeyServer
	SetRekeyThresholdPn(value string) MkaKeyServer
	// HasRekeyThresholdPn checks if RekeyThresholdPn has been set in MkaKeyServer
	HasRekeyThresholdPn() bool
}

type MkaKeyServerConfidentialtyOffsetEnum string

// Enum of ConfidentialtyOffset on MkaKeyServer
var MkaKeyServerConfidentialtyOffset = struct {
	NO_CONFIDENTIALITY               MkaKeyServerConfidentialtyOffsetEnum
	NO_CONFIDENTIALITY_OFFSET        MkaKeyServerConfidentialtyOffsetEnum
	CONFIDENTIALITY_OFFSET_30_OCTETS MkaKeyServerConfidentialtyOffsetEnum
	CONFIDENTIALITY_OFFSET_50_OCTETS MkaKeyServerConfidentialtyOffsetEnum
}{
	NO_CONFIDENTIALITY:               MkaKeyServerConfidentialtyOffsetEnum("no_confidentiality"),
	NO_CONFIDENTIALITY_OFFSET:        MkaKeyServerConfidentialtyOffsetEnum("no_confidentiality_offset"),
	CONFIDENTIALITY_OFFSET_30_OCTETS: MkaKeyServerConfidentialtyOffsetEnum("confidentiality_offset_30_octets"),
	CONFIDENTIALITY_OFFSET_50_OCTETS: MkaKeyServerConfidentialtyOffsetEnum("confidentiality_offset_50_octets"),
}

func (obj *mkaKeyServer) ConfidentialtyOffset() MkaKeyServerConfidentialtyOffsetEnum {
	return MkaKeyServerConfidentialtyOffsetEnum(obj.obj.ConfidentialtyOffset.Enum().String())
}

// Confidentiality Offset.
// ConfidentialtyOffset returns a string
func (obj *mkaKeyServer) HasConfidentialtyOffset() bool {
	return obj.obj.ConfidentialtyOffset != nil
}

func (obj *mkaKeyServer) SetConfidentialtyOffset(value MkaKeyServerConfidentialtyOffsetEnum) MkaKeyServer {
	intValue, ok := otg.MkaKeyServer_ConfidentialtyOffset_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaKeyServerConfidentialtyOffsetEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaKeyServer_ConfidentialtyOffset_Enum(intValue)
	obj.obj.ConfidentialtyOffset = &enumValue

	return obj
}

type MkaKeyServerCipherSuiteEnum string

// Enum of CipherSuite on MkaKeyServer
var MkaKeyServerCipherSuite = struct {
	GCM_AES_128     MkaKeyServerCipherSuiteEnum
	GCM_AES_256     MkaKeyServerCipherSuiteEnum
	GCM_AES_XPN_128 MkaKeyServerCipherSuiteEnum
	GCM_AES_XPN_256 MkaKeyServerCipherSuiteEnum
}{
	GCM_AES_128:     MkaKeyServerCipherSuiteEnum("gcm_aes_128"),
	GCM_AES_256:     MkaKeyServerCipherSuiteEnum("gcm_aes_256"),
	GCM_AES_XPN_128: MkaKeyServerCipherSuiteEnum("gcm_aes_xpn_128"),
	GCM_AES_XPN_256: MkaKeyServerCipherSuiteEnum("gcm_aes_xpn_256"),
}

func (obj *mkaKeyServer) CipherSuite() MkaKeyServerCipherSuiteEnum {
	return MkaKeyServerCipherSuiteEnum(obj.obj.CipherSuite.Enum().String())
}

// The cipher suite. Choose one from GCM-AES-128 GCM-AES-256 GCM-AES-XPN-128 GCM-AES-XPN-256
// CipherSuite returns a string
func (obj *mkaKeyServer) HasCipherSuite() bool {
	return obj.obj.CipherSuite != nil
}

func (obj *mkaKeyServer) SetCipherSuite(value MkaKeyServerCipherSuiteEnum) MkaKeyServer {
	intValue, ok := otg.MkaKeyServer_CipherSuite_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaKeyServerCipherSuiteEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaKeyServer_CipherSuite_Enum(intValue)
	obj.obj.CipherSuite = &enumValue

	return obj
}

// Starting Key Number.
// StartingKeyNumber returns a uint32
func (obj *mkaKeyServer) StartingKeyNumber() uint32 {

	return *obj.obj.StartingKeyNumber

}

// Starting Key Number.
// StartingKeyNumber returns a uint32
func (obj *mkaKeyServer) HasStartingKeyNumber() bool {
	return obj.obj.StartingKeyNumber != nil
}

// Starting Key Number.
// SetStartingKeyNumber sets the uint32 value in the MkaKeyServer object
func (obj *mkaKeyServer) SetStartingKeyNumber(value uint32) MkaKeyServer {

	obj.obj.StartingKeyNumber = &value
	return obj
}

// Starting Distributed AN.
// StartingDistributedAn returns a uint32
func (obj *mkaKeyServer) StartingDistributedAn() uint32 {

	return *obj.obj.StartingDistributedAn

}

// Starting Distributed AN.
// StartingDistributedAn returns a uint32
func (obj *mkaKeyServer) HasStartingDistributedAn() bool {
	return obj.obj.StartingDistributedAn != nil
}

// Starting Distributed AN.
// SetStartingDistributedAn sets the uint32 value in the MkaKeyServer object
func (obj *mkaKeyServer) SetStartingDistributedAn(value uint32) MkaKeyServer {

	obj.obj.StartingDistributedAn = &value
	return obj
}

// Determines the PN/ XPN rekey threshold.
// RekeyThresholdPn returns a string
func (obj *mkaKeyServer) RekeyThresholdPn() string {

	return *obj.obj.RekeyThresholdPn

}

// Determines the PN/ XPN rekey threshold.
// RekeyThresholdPn returns a string
func (obj *mkaKeyServer) HasRekeyThresholdPn() bool {
	return obj.obj.RekeyThresholdPn != nil
}

// Determines the PN/ XPN rekey threshold.
// SetRekeyThresholdPn sets the string value in the MkaKeyServer object
func (obj *mkaKeyServer) SetRekeyThresholdPn(value string) MkaKeyServer {

	obj.obj.RekeyThresholdPn = &value
	return obj
}

func (obj *mkaKeyServer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartingKeyNumber != nil {

		if *obj.obj.StartingKeyNumber < 1 || *obj.obj.StartingKeyNumber > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MkaKeyServer.StartingKeyNumber <= 65535 but Got %d", *obj.obj.StartingKeyNumber))
		}

	}

	if obj.obj.StartingDistributedAn != nil {

		if *obj.obj.StartingDistributedAn > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaKeyServer.StartingDistributedAn <= 3 but Got %d", *obj.obj.StartingDistributedAn))
		}

	}

	if obj.obj.RekeyThresholdPn != nil {

		if len(*obj.obj.RekeyThresholdPn) < 4 || len(*obj.obj.RekeyThresholdPn) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"4 <= length of MkaKeyServer.RekeyThresholdPn <= 8 but Got %d",
					len(*obj.obj.RekeyThresholdPn)))
		}

	}

}

func (obj *mkaKeyServer) setDefault() {
	if obj.obj.ConfidentialtyOffset == nil {
		obj.SetConfidentialtyOffset(MkaKeyServerConfidentialtyOffset.NO_CONFIDENTIALITY_OFFSET)

	}
	if obj.obj.CipherSuite == nil {
		obj.SetCipherSuite(MkaKeyServerCipherSuite.GCM_AES_128)

	}
	if obj.obj.StartingKeyNumber == nil {
		obj.SetStartingKeyNumber(1)
	}
	if obj.obj.StartingDistributedAn == nil {
		obj.SetStartingDistributedAn(0)
	}
	if obj.obj.RekeyThresholdPn == nil {
		obj.SetRekeyThresholdPn("C0000000")
	}

}
