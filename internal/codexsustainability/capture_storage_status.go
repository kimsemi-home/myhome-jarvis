package codexsustainability

func withStorageArchiveCapture(
	status CaptureStatus,
	capture storageArchiveCapture,
) CaptureStatus {
	status.StorageArchiveState = capture.State
	status.StorageCacheHitCount = capture.CacheHitCount
	status.StorageCacheMissCount = capture.CacheMissCount
	status.StorageCacheSavingsUnits = capture.CacheSavingsUnits
	return status
}
