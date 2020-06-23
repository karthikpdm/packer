// Code generated by protoc-gen-goext. DO NOT EDIT.

package devices

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *Device) SetId(v string) {
	m.Id = v
}

func (m *Device) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *Device) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Device) SetName(v string) {
	m.Name = v
}

func (m *Device) SetDescription(v string) {
	m.Description = v
}

func (m *Device) SetTopicAliases(v map[string]string) {
	m.TopicAliases = v
}

func (m *Device) SetStatus(v Device_Status) {
	m.Status = v
}

func (m *DeviceCertificate) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *DeviceCertificate) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *DeviceCertificate) SetCertificateData(v string) {
	m.CertificateData = v
}

func (m *DeviceCertificate) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *DevicePassword) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *DevicePassword) SetId(v string) {
	m.Id = v
}

func (m *DevicePassword) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}