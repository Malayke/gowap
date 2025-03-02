// This file is generated by "./lib/proto/generate"

package proto

import (
	"github.com/ysmood/gson"
)

/*

SystemInfo

The SystemInfo domain defines methods and events for querying low-level system information.

*/

// SystemInfoGPUDevice Describes a single graphics processor (GPU).
type SystemInfoGPUDevice struct {

	// VendorID PCI ID of the GPU vendor, if available; 0 otherwise.
	VendorID float64 `json:"vendorId"`

	// DeviceID PCI ID of the GPU device, if available; 0 otherwise.
	DeviceID float64 `json:"deviceId"`

	// SubSysID (optional) Sub sys ID of the GPU, only available on Windows.
	SubSysID *float64 `json:"subSysId,omitempty"`

	// Revision (optional) Revision of the GPU, only available on Windows.
	Revision *float64 `json:"revision,omitempty"`

	// VendorString String description of the GPU vendor, if the PCI ID is not available.
	VendorString string `json:"vendorString"`

	// DeviceString String description of the GPU device, if the PCI ID is not available.
	DeviceString string `json:"deviceString"`

	// DriverVendor String description of the GPU driver vendor.
	DriverVendor string `json:"driverVendor"`

	// DriverVersion String description of the GPU driver version.
	DriverVersion string `json:"driverVersion"`
}

// SystemInfoSize Describes the width and height dimensions of an entity.
type SystemInfoSize struct {

	// Width Width in pixels.
	Width int `json:"width"`

	// Height Height in pixels.
	Height int `json:"height"`
}

// SystemInfoVideoDecodeAcceleratorCapability Describes a supported video decoding profile with its associated minimum and
// maximum resolutions.
type SystemInfoVideoDecodeAcceleratorCapability struct {

	// Profile Video codec profile that is supported, e.g. VP9 Profile 2.
	Profile string `json:"profile"`

	// MaxResolution Maximum video dimensions in pixels supported for this |profile|.
	MaxResolution *SystemInfoSize `json:"maxResolution"`

	// MinResolution Minimum video dimensions in pixels supported for this |profile|.
	MinResolution *SystemInfoSize `json:"minResolution"`
}

// SystemInfoVideoEncodeAcceleratorCapability Describes a supported video encoding profile with its associated maximum
// resolution and maximum framerate.
type SystemInfoVideoEncodeAcceleratorCapability struct {

	// Profile Video codec profile that is supported, e.g H264 Main.
	Profile string `json:"profile"`

	// MaxResolution Maximum video dimensions in pixels supported for this |profile|.
	MaxResolution *SystemInfoSize `json:"maxResolution"`

	// MaxFramerateNumerator Maximum encoding framerate in frames per second supported for this
	// |profile|, as fraction's numerator and denominator, e.g. 24/1 fps,
	// 24000/1001 fps, etc.
	MaxFramerateNumerator int `json:"maxFramerateNumerator"`

	// MaxFramerateDenominator ...
	MaxFramerateDenominator int `json:"maxFramerateDenominator"`
}

// SystemInfoSubsamplingFormat YUV subsampling type of the pixels of a given image.
type SystemInfoSubsamplingFormat string

const (
	// SystemInfoSubsamplingFormatYuv420 enum const
	SystemInfoSubsamplingFormatYuv420 SystemInfoSubsamplingFormat = "yuv420"

	// SystemInfoSubsamplingFormatYuv422 enum const
	SystemInfoSubsamplingFormatYuv422 SystemInfoSubsamplingFormat = "yuv422"

	// SystemInfoSubsamplingFormatYuv444 enum const
	SystemInfoSubsamplingFormatYuv444 SystemInfoSubsamplingFormat = "yuv444"
)

// SystemInfoImageType Image format of a given image.
type SystemInfoImageType string

const (
	// SystemInfoImageTypeJpeg enum const
	SystemInfoImageTypeJpeg SystemInfoImageType = "jpeg"

	// SystemInfoImageTypeWebp enum const
	SystemInfoImageTypeWebp SystemInfoImageType = "webp"

	// SystemInfoImageTypeUnknown enum const
	SystemInfoImageTypeUnknown SystemInfoImageType = "unknown"
)

// SystemInfoImageDecodeAcceleratorCapability Describes a supported image decoding profile with its associated minimum and
// maximum resolutions and subsampling.
type SystemInfoImageDecodeAcceleratorCapability struct {

	// ImageType Image coded, e.g. Jpeg.
	ImageType SystemInfoImageType `json:"imageType"`

	// MaxDimensions Maximum supported dimensions of the image in pixels.
	MaxDimensions *SystemInfoSize `json:"maxDimensions"`

	// MinDimensions Minimum supported dimensions of the image in pixels.
	MinDimensions *SystemInfoSize `json:"minDimensions"`

	// Subsamplings Optional array of supported subsampling formats, e.g. 4:2:0, if known.
	Subsamplings []SystemInfoSubsamplingFormat `json:"subsamplings"`
}

// SystemInfoGPUInfo Provides information about the GPU(s) on the system.
type SystemInfoGPUInfo struct {

	// Devices The graphics devices on the system. Element 0 is the primary GPU.
	Devices []*SystemInfoGPUDevice `json:"devices"`

	// AuxAttributes (optional) An optional dictionary of additional GPU related attributes.
	AuxAttributes map[string]gson.JSON `json:"auxAttributes,omitempty"`

	// FeatureStatus (optional) An optional dictionary of graphics features and their status.
	FeatureStatus map[string]gson.JSON `json:"featureStatus,omitempty"`

	// DriverBugWorkarounds An optional array of GPU driver bug workarounds.
	DriverBugWorkarounds []string `json:"driverBugWorkarounds"`

	// VideoDecoding Supported accelerated video decoding capabilities.
	VideoDecoding []*SystemInfoVideoDecodeAcceleratorCapability `json:"videoDecoding"`

	// VideoEncoding Supported accelerated video encoding capabilities.
	VideoEncoding []*SystemInfoVideoEncodeAcceleratorCapability `json:"videoEncoding"`

	// ImageDecoding Supported accelerated image decoding capabilities.
	ImageDecoding []*SystemInfoImageDecodeAcceleratorCapability `json:"imageDecoding"`
}

// SystemInfoProcessInfo Represents process info.
type SystemInfoProcessInfo struct {

	// Type Specifies process type.
	Type string `json:"type"`

	// ID Specifies process id.
	ID int `json:"id"`

	// CPUTime Specifies cumulative CPU usage in seconds across all threads of the
	// process since the process start.
	CPUTime float64 `json:"cpuTime"`
}

// SystemInfoGetInfo Returns information about the system.
type SystemInfoGetInfo struct {
}

// ProtoReq name
func (m SystemInfoGetInfo) ProtoReq() string { return "SystemInfo.getInfo" }

// Call the request
func (m SystemInfoGetInfo) Call(c Client) (*SystemInfoGetInfoResult, error) {
	var res SystemInfoGetInfoResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// SystemInfoGetInfoResult Returns information about the system.
type SystemInfoGetInfoResult struct {

	// Gpu Information about the GPUs on the system.
	Gpu *SystemInfoGPUInfo `json:"gpu"`

	// ModelName A platform-dependent description of the model of the machine. On Mac OS, this is, for
	// example, 'MacBookPro'. Will be the empty string if not supported.
	ModelName string `json:"modelName"`

	// ModelVersion A platform-dependent description of the version of the machine. On Mac OS, this is, for
	// example, '10.1'. Will be the empty string if not supported.
	ModelVersion string `json:"modelVersion"`

	// CommandLine The command line string used to launch the browser. Will be the empty string if not
	// supported.
	CommandLine string `json:"commandLine"`
}

// SystemInfoGetProcessInfo Returns information about all running processes.
type SystemInfoGetProcessInfo struct {
}

// ProtoReq name
func (m SystemInfoGetProcessInfo) ProtoReq() string { return "SystemInfo.getProcessInfo" }

// Call the request
func (m SystemInfoGetProcessInfo) Call(c Client) (*SystemInfoGetProcessInfoResult, error) {
	var res SystemInfoGetProcessInfoResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// SystemInfoGetProcessInfoResult Returns information about all running processes.
type SystemInfoGetProcessInfoResult struct {

	// ProcessInfo An array of process info blocks.
	ProcessInfo []*SystemInfoProcessInfo `json:"processInfo"`
}
