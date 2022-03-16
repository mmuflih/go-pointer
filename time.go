package ptr

import "time"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func Time(source *time.Time, def time.Time) time.Time {
	if source != nil {
		return *source
	}
	return def
}
