package pointer

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func String(source *string, def string) string {
	if source != nil {
		return *source
	}
	return def
}
