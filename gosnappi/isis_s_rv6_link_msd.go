package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6LinkMsd *****
type isisSRv6LinkMsd struct {
	validation
	obj          *otg.IsisSRv6LinkMsd
	marshaller   marshalIsisSRv6LinkMsd
	unMarshaller unMarshalIsisSRv6LinkMsd
}

func NewIsisSRv6LinkMsd() IsisSRv6LinkMsd {
	obj := isisSRv6LinkMsd{obj: &otg.IsisSRv6LinkMsd{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6LinkMsd) msg() *otg.IsisSRv6LinkMsd {
	return obj.obj
}

func (obj *isisSRv6LinkMsd) setMsg(msg *otg.IsisSRv6LinkMsd) IsisSRv6LinkMsd {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6LinkMsd struct {
	obj *isisSRv6LinkMsd
}

type marshalIsisSRv6LinkMsd interface {
	// ToProto marshals IsisSRv6LinkMsd to protobuf object *otg.IsisSRv6LinkMsd
	ToProto() (*otg.IsisSRv6LinkMsd, error)
	// ToPbText marshals IsisSRv6LinkMsd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6LinkMsd to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6LinkMsd to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6LinkMsd struct {
	obj *isisSRv6LinkMsd
}

type unMarshalIsisSRv6LinkMsd interface {
	// FromProto unmarshals IsisSRv6LinkMsd from protobuf object *otg.IsisSRv6LinkMsd
	FromProto(msg *otg.IsisSRv6LinkMsd) (IsisSRv6LinkMsd, error)
	// FromPbText unmarshals IsisSRv6LinkMsd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6LinkMsd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6LinkMsd from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6LinkMsd) Marshal() marshalIsisSRv6LinkMsd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6LinkMsd{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6LinkMsd) Unmarshal() unMarshalIsisSRv6LinkMsd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6LinkMsd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6LinkMsd) ToProto() (*otg.IsisSRv6LinkMsd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6LinkMsd) FromProto(msg *otg.IsisSRv6LinkMsd) (IsisSRv6LinkMsd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6LinkMsd) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6LinkMsd) FromPbText(value string) error {
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

func (m *marshalisisSRv6LinkMsd) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6LinkMsd) FromYaml(value string) error {
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

func (m *marshalisisSRv6LinkMsd) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6LinkMsd) FromJson(value string) error {
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

func (obj *isisSRv6LinkMsd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6LinkMsd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6LinkMsd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6LinkMsd) Clone() (IsisSRv6LinkMsd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6LinkMsd()
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

// IsisSRv6LinkMsd is link-level SRv6 Maximum SID Depth (MSD) capabilities advertised per IS-IS interface. Signals the SRv6 SRH processing capacity specific to this link, which may differ from the node-level MSD values. A non-zero value causes the corresponding MSD sub-TLV to be included in the neighbor TLV advertisement; a value of 0 (default) suppresses it. This follows RFC 8491 where MSD is a list of (MSD-Type, MSD-Value) tuples. Reference: RFC 8491, RFC 9352 Section 6.
type IsisSRv6LinkMsd interface {
	Validation
	// msg marshals IsisSRv6LinkMsd to protobuf object *otg.IsisSRv6LinkMsd
	// and doesn't set defaults
	msg() *otg.IsisSRv6LinkMsd
	// setMsg unmarshals IsisSRv6LinkMsd from protobuf object *otg.IsisSRv6LinkMsd
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6LinkMsd) IsisSRv6LinkMsd
	// provides marshal interface
	Marshal() marshalIsisSRv6LinkMsd
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6LinkMsd
	// validate validates IsisSRv6LinkMsd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6LinkMsd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MaxSl returns uint32, set in IsisSRv6LinkMsd.
	MaxSl() uint32
	// SetMaxSl assigns uint32 provided by user to IsisSRv6LinkMsd
	SetMaxSl(value uint32) IsisSRv6LinkMsd
	// HasMaxSl checks if MaxSl has been set in IsisSRv6LinkMsd
	HasMaxSl() bool
	// MaxEndPopSrh returns uint32, set in IsisSRv6LinkMsd.
	MaxEndPopSrh() uint32
	// SetMaxEndPopSrh assigns uint32 provided by user to IsisSRv6LinkMsd
	SetMaxEndPopSrh(value uint32) IsisSRv6LinkMsd
	// HasMaxEndPopSrh checks if MaxEndPopSrh has been set in IsisSRv6LinkMsd
	HasMaxEndPopSrh() bool
	// MaxTInsert returns uint32, set in IsisSRv6LinkMsd.
	MaxTInsert() uint32
	// SetMaxTInsert assigns uint32 provided by user to IsisSRv6LinkMsd
	SetMaxTInsert(value uint32) IsisSRv6LinkMsd
	// HasMaxTInsert checks if MaxTInsert has been set in IsisSRv6LinkMsd
	HasMaxTInsert() bool
	// MaxHEncaps returns uint32, set in IsisSRv6LinkMsd.
	MaxHEncaps() uint32
	// SetMaxHEncaps assigns uint32 provided by user to IsisSRv6LinkMsd
	SetMaxHEncaps(value uint32) IsisSRv6LinkMsd
	// HasMaxHEncaps checks if MaxHEncaps has been set in IsisSRv6LinkMsd
	HasMaxHEncaps() bool
	// MaxEndDSrh returns uint32, set in IsisSRv6LinkMsd.
	MaxEndDSrh() uint32
	// SetMaxEndDSrh assigns uint32 provided by user to IsisSRv6LinkMsd
	SetMaxEndDSrh(value uint32) IsisSRv6LinkMsd
	// HasMaxEndDSrh checks if MaxEndDSrh has been set in IsisSRv6LinkMsd
	HasMaxEndDSrh() bool
}

// Maximum value of the Segments Left (SL) field in an SRH that this interface can process (MSD-Type 41). A non-zero value is advertised as the Max SL MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// MaxSl returns a uint32
func (obj *isisSRv6LinkMsd) MaxSl() uint32 {

	return *obj.obj.MaxSl

}

// Maximum value of the Segments Left (SL) field in an SRH that this interface can process (MSD-Type 41). A non-zero value is advertised as the Max SL MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// MaxSl returns a uint32
func (obj *isisSRv6LinkMsd) HasMaxSl() bool {
	return obj.obj.MaxSl != nil
}

// Maximum value of the Segments Left (SL) field in an SRH that this interface can process (MSD-Type 41). A non-zero value is advertised as the Max SL MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// SetMaxSl sets the uint32 value in the IsisSRv6LinkMsd object
func (obj *isisSRv6LinkMsd) SetMaxSl(value uint32) IsisSRv6LinkMsd {

	obj.obj.MaxSl = &value
	return obj
}

// Maximum number of SIDs in the top SRH for PSP or USP End behaviors on this interface (MSD-Type 42). A non-zero value is advertised as the Max-End-Pop-SRH MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// MaxEndPopSrh returns a uint32
func (obj *isisSRv6LinkMsd) MaxEndPopSrh() uint32 {

	return *obj.obj.MaxEndPopSrh

}

// Maximum number of SIDs in the top SRH for PSP or USP End behaviors on this interface (MSD-Type 42). A non-zero value is advertised as the Max-End-Pop-SRH MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// MaxEndPopSrh returns a uint32
func (obj *isisSRv6LinkMsd) HasMaxEndPopSrh() bool {
	return obj.obj.MaxEndPopSrh != nil
}

// Maximum number of SIDs in the top SRH for PSP or USP End behaviors on this interface (MSD-Type 42). A non-zero value is advertised as the Max-End-Pop-SRH MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// SetMaxEndPopSrh sets the uint32 value in the IsisSRv6LinkMsd object
func (obj *isisSRv6LinkMsd) SetMaxEndPopSrh(value uint32) IsisSRv6LinkMsd {

	obj.obj.MaxEndPopSrh = &value
	return obj
}

// Maximum number of SIDs insertable via T.Insert behavior on this interface (MSD-Type 43, RFC 8491). A non-zero value is advertised as the Maximum T.Insert MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// MaxTInsert returns a uint32
func (obj *isisSRv6LinkMsd) MaxTInsert() uint32 {

	return *obj.obj.MaxTInsert

}

// Maximum number of SIDs insertable via T.Insert behavior on this interface (MSD-Type 43, RFC 8491). A non-zero value is advertised as the Maximum T.Insert MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// MaxTInsert returns a uint32
func (obj *isisSRv6LinkMsd) HasMaxTInsert() bool {
	return obj.obj.MaxTInsert != nil
}

// Maximum number of SIDs insertable via T.Insert behavior on this interface (MSD-Type 43, RFC 8491). A non-zero value is advertised as the Maximum T.Insert MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// SetMaxTInsert sets the uint32 value in the IsisSRv6LinkMsd object
func (obj *isisSRv6LinkMsd) SetMaxTInsert(value uint32) IsisSRv6LinkMsd {

	obj.obj.MaxTInsert = &value
	return obj
}

// Maximum number of SIDs for H.Encaps behavior on this interface (MSD-Type 44). A non-zero value is advertised as the Max-H-Encaps MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV. A value of 0 with E-flag set indicates IPinIP encapsulation without SRH.
// MaxHEncaps returns a uint32
func (obj *isisSRv6LinkMsd) MaxHEncaps() uint32 {

	return *obj.obj.MaxHEncaps

}

// Maximum number of SIDs for H.Encaps behavior on this interface (MSD-Type 44). A non-zero value is advertised as the Max-H-Encaps MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV. A value of 0 with E-flag set indicates IPinIP encapsulation without SRH.
// MaxHEncaps returns a uint32
func (obj *isisSRv6LinkMsd) HasMaxHEncaps() bool {
	return obj.obj.MaxHEncaps != nil
}

// Maximum number of SIDs for H.Encaps behavior on this interface (MSD-Type 44). A non-zero value is advertised as the Max-H-Encaps MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV. A value of 0 with E-flag set indicates IPinIP encapsulation without SRH.
// SetMaxHEncaps sets the uint32 value in the IsisSRv6LinkMsd object
func (obj *isisSRv6LinkMsd) SetMaxHEncaps(value uint32) IsisSRv6LinkMsd {

	obj.obj.MaxHEncaps = &value
	return obj
}

// Maximum number of SIDs in the SRH for decapsulation behaviors (End.DT4/DT6/DT46/DX4/DX6) on this interface (MSD-Type 45). A non-zero value is advertised as the Max-End-D-SRH MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// MaxEndDSrh returns a uint32
func (obj *isisSRv6LinkMsd) MaxEndDSrh() uint32 {

	return *obj.obj.MaxEndDSrh

}

// Maximum number of SIDs in the SRH for decapsulation behaviors (End.DT4/DT6/DT46/DX4/DX6) on this interface (MSD-Type 45). A non-zero value is advertised as the Max-End-D-SRH MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// MaxEndDSrh returns a uint32
func (obj *isisSRv6LinkMsd) HasMaxEndDSrh() bool {
	return obj.obj.MaxEndDSrh != nil
}

// Maximum number of SIDs in the SRH for decapsulation behaviors (End.DT4/DT6/DT46/DX4/DX6) on this interface (MSD-Type 45). A non-zero value is advertised as the Max-End-D-SRH MSD sub-TLV at the link level; a value of 0 (default) suppresses the sub-TLV.
// SetMaxEndDSrh sets the uint32 value in the IsisSRv6LinkMsd object
func (obj *isisSRv6LinkMsd) SetMaxEndDSrh(value uint32) IsisSRv6LinkMsd {

	obj.obj.MaxEndDSrh = &value
	return obj
}

func (obj *isisSRv6LinkMsd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxSl != nil {

		if *obj.obj.MaxSl > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6LinkMsd.MaxSl <= 255 but Got %d", *obj.obj.MaxSl))
		}

	}

	if obj.obj.MaxEndPopSrh != nil {

		if *obj.obj.MaxEndPopSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6LinkMsd.MaxEndPopSrh <= 255 but Got %d", *obj.obj.MaxEndPopSrh))
		}

	}

	if obj.obj.MaxTInsert != nil {

		if *obj.obj.MaxTInsert > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6LinkMsd.MaxTInsert <= 255 but Got %d", *obj.obj.MaxTInsert))
		}

	}

	if obj.obj.MaxHEncaps != nil {

		if *obj.obj.MaxHEncaps > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6LinkMsd.MaxHEncaps <= 255 but Got %d", *obj.obj.MaxHEncaps))
		}

	}

	if obj.obj.MaxEndDSrh != nil {

		if *obj.obj.MaxEndDSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6LinkMsd.MaxEndDSrh <= 255 but Got %d", *obj.obj.MaxEndDSrh))
		}

	}

}

func (obj *isisSRv6LinkMsd) setDefault() {
	if obj.obj.MaxSl == nil {
		obj.SetMaxSl(0)
	}
	if obj.obj.MaxEndPopSrh == nil {
		obj.SetMaxEndPopSrh(0)
	}
	if obj.obj.MaxTInsert == nil {
		obj.SetMaxTInsert(0)
	}
	if obj.obj.MaxHEncaps == nil {
		obj.SetMaxHEncaps(0)
	}
	if obj.obj.MaxEndDSrh == nil {
		obj.SetMaxEndDSrh(0)
	}

}
