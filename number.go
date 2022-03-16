package pointer

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func UInt8(source *uint8, def uint8) uint8 {
	if source != nil {
		return *source
	}
	return def
}

func Int8(source *int8, def int8) int8 {
	if source != nil {
		return *source
	}
	return def
}

func Int(source *int, def int) int {
	if source != nil {
		return *source
	}
	return def
}

func UInt(source *uint, def uint) uint {
	if source != nil {
		return *source
	}
	return def
}

func UInt32(source *uint32, def uint32) uint32 {
	if source != nil {
		return *source
	}
	return def
}

func Int32(source *int32, def int32) int32 {
	if source != nil {
		return *source
	}
	return def
}

func UInt64(source *uint64, def uint64) uint64 {
	if source != nil {
		return *source
	}
	return def
}

func Int64(source *int64, def int64) int64 {
	if source != nil {
		return *source
	}
	return def
}

func Float32(source *float32, def float32) float32 {
	if source != nil {
		return *source
	}
	return def
}

func Float64(source *float64, def float64) float64 {
	if source != nil {
		return *source
	}
	return def
}
