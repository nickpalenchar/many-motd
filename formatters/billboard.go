package formatters

import (
	"fmt"
	"many-motd/data"
	"strings"
)

// Billboard creates a hash (#) boarder around the
// message. If there is an author, it's right-justified
func Billboard(m data.Message, w int) string {

	hb := strings.Repeat("#", w+4)
	msg := wrap(m.Body, w)

	result := []string{
		hb,
		"#" + strings.Repeat(" ", w+2) + "#",
	}
	for _, line := range strings.Split(msg, "\n") {
		result = append(result, fmt.Sprintf("# %-*s #", w, line))
	}
	if m.Author != "" {
		result = append(result, fmt.Sprintf("# %*s #", w, "-- "+m.Author+" "))
	} else {
		result = append(result, "#"+strings.Repeat(" ", w+2)+"#")
	}
	result = append(result, hb)

	return strings.Join(result, "\n")
}
